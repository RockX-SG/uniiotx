pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";

import "interfaces/IIOTXClear.sol"
import "interfaces/IUniIOTX.sol"
import "../interfaces/ISystemStake.sol"

contract IOTXStake is Initializable, PausableUpgradeable, AccessControlUpgradeable, IERC721Receiver {
    // External dependencies
    ISystemStake public systemStake;
    IUniIOTX public uniIOTX;
    IIOTXClear public iotxClear;

    // Type declarations
    enum BucketType {
        BucketType1,
        BucketType2,
        BucketType3
    }

    struct tokenIds {
        uint256 size;
        mapping(uint256 => uint256) data;   // startId => count
    }

    // Stake variables
    uint256 private totalPending;               // total pending IOTXs awaiting to be staked

    address[] private delegates;
    uint256 private nextDelegateIndex;

    uint256 private stakeAmount01 = 10000;
    uint256 private stakeAmount02 = 100000;
    uint256 private stakeAmount03 = 1000000;
    uint256 private stakeDuration = 3600*91;

    uint256 private mergeThreshold = 10;

    // Token ids
    mapping(BucketType => tokenIds) public lockedTokenIds;
    mapping(BucketType => tokenIds) public unlockedTokenIds;
    mapping(BucketType => tokenIds) public unstakedTokenIds;
    mapping(BucketType => tokenIds) public withdrawedTokenIds;

    // Events
    event DelegateStopped(uint256 stoppedCount);
    event SystemStakingContractSet(address addr);
    event XIOTXContractSet(address addr);
    event DepositEvent(address indexed from, uint256 amount);
    event RedeemContractSet(address addr);

    // Errors
    error ZeroDelegates();
    error TransactionExpired(uint256 deadline, uint256 now);
    error InvalidRedeemAmount(uint256 redeemAmount, uint256 redeemBase);

    /**
     * ======================================================================================
     *
     * SYSTEM SETTINGS
     *
     * ======================================================================================
     */

    /**
     * @dev pause the contract
     */
    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /**
     * @dev unpause the contract
     */
    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /**
     * @dev initialization address
     */
    function initialize() initializer public {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(PAUSER_ROLE, msg.sender);

//        // init default values
//        firstDebt = 1;
//        lastDebt = 0;
    }


    /**
     * @dev register a new batch of validators
     */
    function registerDelegates(address[] _delegates) external onlyRole(REGISTRY_ROLE) {
        uint256 len = _delegates.length;
        if (len == 0) revert ZeroDelegates();
        delete delegates;
        delete nextDelegateIndex;
        for (uint256 i = 0; i < len; i++) {
            delegates.push(_delegates[i]);
        }
    }

    /**
     * @dev set eth deposit contract address
     */
    function setSystemStakingContract(address _SystemStakingContract) external onlyRole(DEFAULT_ADMIN_ROLE) {
        systemStake = _SystemStakingContract;

        emit SystemStakingContractSet(_SystemStakingContract);
    }

    /**
     * @dev set xIOTX token contract address
     */
    function setXIOTXContractAddress(address _xIOTXAddress) external onlyRole(DEFAULT_ADMIN_ROLE) {
        uniIOTX = _xIOTXAddress;

        emit XIOTXContractSet(_xIOTXAddress);
    }

    /**
     * @dev set redeem contract
     */
    function setRedeemContract(address _redeemContract) external onlyRole(DEFAULT_ADMIN_ROLE) {
        iotxClear = _redeemContract;

        emit RedeemContractSet(_redeemContract);
    }

    /**
     * ======================================================================================
     *
     * VIEW FUNCTIONS
     *
     * ======================================================================================
     */

    /**
     * @dev return pending IOTXs
     */
    function getPendingIOTXs() external view returns (uint256) { return totalPending; }

    /**
     * @dev return number of registered delegates
     */
    function getDelegatesCount() external view returns (uint256) {
        return delegates.length;
    }

    /**
     * @dev return all registered delegates
     */
    function getAllDelegates() external view returns (address[]) {
        return delegates;
    }

    /**
     * ======================================================================================
     *
     * EXTERNAL FUNCTIONS
     *
     * ======================================================================================
     */

    /**
     * @dev IERC721Receiver implement for receiving staking NFT
     */
    function onERC721Received(
        address, // operator
        address, // from
        uint256, // tokenId
        bytes calldata // data
    ) external pure override returns (bytes4) {
        return this.onERC721Received.selector;
    }

    /**
     * @dev mint uniIOTX with IOTX
     */
    function deposit() external payable returns (uint256 minted) {
        amount = msg.value;
        require(amount > 0, "USR002"); // TODO: don't need to check zero mint?
        // TODO: check msgValue ceiling

        // TODO: to be optimized

        // mint uniIOTX
        _mint();

        // stake
        _stake();

        // merge
        _merge();

        return toMint;
    }

    // TODO: to be modified
    function redeem(uint256 amount, uint256 maxToBurn, uint256 deadline) external nonReentrant returns(uint256 burned) {
        now = block.timestamp;
        if (deadline <= now) revert TransactionExpired(deadline, now);
        if (amount % stakeAmount03 != 0) revert InvalidRedeemAmount(amount, stakeAmount03);
        // TODO: verify the rights/preconditiions of msg.sender to redeem ? including:
        // TODO: check amount ceiling according to clients deposit amount.

        uint256 totalXIOTX = uniIOTX.totalSupply();
        uint256 xIOTXToBurn = totalXIOTX * amount / currentReserve(); // TODO:
        require(xIOTXToBurn <= maxToBurn, "USR004");

        // NOTE: the following procdure must keep exchangeRatio invariant:
        // transfer uniIOTX from sender & burn
        // TODO: why transfer and burn, but not only burn?
        uniIOTX.safeTransferFrom(msg.sender, address(this), xIOTXToBurn);
        uniIOTX.burn(xIOTXToBurn);

        // Unlock NFT(s)
        count = amount / stakeAmount03;
        tokenIds = _removeLockedTokenIds(BucketType.BucketType3, count);
        systemStake.unlock(tokenIds);
        _addUnlockedTokenIds(tokenIds);

        // Transfer NFT owner
        for (i = 0; i < tokenIds.length; i++) {
            systemStake.transferFrom(address(this), address(iotxClear), tokenIds[i]);
        }

        // Join debt
        iotxClear.joinDebt(msg.sender, amount);

        // return burned
        return xIOTXToBurn;
    }

    function changeDelegates(uint256[] tokenId, address delegate) external whenNotPaused onlyRole(ORACLE_ROLE) {
        systemStake.changeDelegates(tokenId, delegate);
    }

    /**
     * ======================================================================================
     *
     * INTERNAL FUNCTIONS
     *
     * ======================================================================================
     */


    /**
     * ======================================================================================
     *
     * PRIVATE FUNCTIONS
     *
     * ======================================================================================
     */

    // todo: to be optimized
    function _mint() private {
        uint256 toMint = 1 * amount; // default exchange ratio 1:1
        uniIOTX.mint(msg.sender, toMint);
        totalPending += amount;
        emit DepositEvent(msg.sender, amount);
    }

    function _stake() private {
        delegate = _nextDelegate();

        count = totalPending / stakeAmount03;
        _requestStake(stakeAmount03, count, delegate);

        count = totalPending / stakeAmount02;
        _requestStake(stakeAmount02, count, delegate);

        count = totalPending / stakeAmount01;
        _requestStake(stakeAmount01, count, delegate);
    }

    function _merge() private {
        count = _lockedTokenIdCount(BucketType.BucketType1);
        if (count >= mergeThreshold) {
            tokenIds = _removeLockedTokenIds(BucketType.BucketType1, mergeThreshold);
            systemStake.merge(tokenIds,stakeDuration);
            theFirst = new uint256[](1);
            theFirst[0] = tokenIds[0];
            _pushLockedTokens(BucketType.BucketType2, theFirst);
        }

        count = _lockedTokenIdCount(BucketType.BucketType2);
        if (count >= mergeThreshold) {
            tokenIds = _removeLockedTokenIds(BucketType.BucketType2, mergeThreshold);
            systemStake.merge(tokenIds,stakeDuration);
            theFirst = new uint256[](1);
            theFirst[0] = tokenIds[0];
            _pushLockedTokens(BucketType.BucketType3, theFirst);
        }
    }

    function _requestStake(BucketType bucketType, uint256 amount, uint256 count, address delegate) private{
        if (count == 0) return;
        totalAmount = amount*count;
        if (count == 1) {
            tokenId = systemStake.stake{value:totalAmount}(stakeDuration, delegate);
            _addLockedTokenIds(BucketType.bucketType, tokenId,1);
        } else {
            startTokenId = systemStake.stake{value:totalAmount}(amount,stakeDuration, delegate, count);
            _addLockedTokenIds(BucketType.bucketType, startTokenId, count);
        }
        totalPending -= totalAmount;
    }

    function _nextDelegate() private returns (address) {
        delegate = delegates[nextDelegateIndex];
        nextDelegateIndex = (nextDelegateIndex + 1) % delegates.length;
        return delegate;
    }

    function _lockedTokenIdCount(BucketType bucketType) private returns (uint256) {
        return lockedTokenIds[bucketType].size;
    }

    function _removeLockedTokenIds(BucketType bucketType, uint256 count) private returns (uint256[]) {
//        delete lockedTokenIds[bucketType].ids[]
        // TODO:
        return 0;
    }

    function _addLockedTokenIds(BucketType bucketType, uint256 firstTokenId, uint256 count) private {
        lockedTokenIds[bucketType].data[firstTokenId] = count;
        lockedTokenIds[bucketType].size += count;
    }

    // TODO: consider move it to IOTXCLear.sol
    function _addUnlockedTokenIds(BucketType bucketType, uint256[] tokenIds) private {
        // todo:
    }

}