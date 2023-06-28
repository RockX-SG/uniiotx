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

    // Constants
    uint256 public defaultExchangeRatio = 1;
    uint256 private constant MULTIPLIER = 1e18;

    // State variables

    uint256 public accountedBalance; // TODO: Double check whether its value can be negative.


    // Stake variables

    // exchange ratio related variables
    // track user deposits & redeem (uniIOTX mint & burn)
    uint256 private totalPending;               // total pending IOTXs awaiting to be staked
    uint256 private totalStaked;            // track current staked ethers for delegates
    uint256 private totalDebts;             // track current unpaid debts // TODO: ignore it?

    // track revenue from validators to form exchange ratio
    uint256 private accountedUserRevenue;           // accounted shared user revenue
    uint256 private accountedManagerRevenue;        // accounted manager's revenue
    uint256 private rewardDebts;                    // check validatorStopped function // TODO: modify this comment

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
    event MintEvent(address user, uint256 minToMint, uint256 minted);
    event RedeemContractSet(address addr);


    // Errors
    error ZeroDelegates();
    error TransactionExpired(uint256 deadline, uint256 now);
    error InvalidRedeemAmount(uint256 redeemAmount, uint256 redeemBase);
    error ExchangeRatioMismatch(uint256 expectedAmount, uint256 gotAmount);

    // Modifiers // TODO: code reuse across smart contracts
    modifier onlyValidTransaction(uint256 deadline) {
        if (deadline <= block.timestamp) revert TransactionExpired(deadline, block.timestamp);
        _;
    }

    modifer notZeroMint() {
        if (msg.value == 0) {
            return;
        }
        _;
    }

    /**
     * ======================================================================================
     *
     * SYSTEM SETTINGS
     *
     * ======================================================================================
     */

    // TODO: What if a user accidentally sent some IOTXs to this contract?
    // These IOTXs will be treated as reward. How could we prevent that?
    receive() external payable { }

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
     * @dev Return exchange ratio of uniIOTX to IOTX, multiplied by 1e18
     */
    function exchangeRatio() external view returns (uint256 ratio) {
        uint256 uniIOTXAmount = uniIOTX.totalSupply();
        if (uniIOTXAmount == 0) {
            return 1 * MULTIPLIER;
        }

        ratio = currentReserve() * MULTIPLIER / uniIOTXAmount; // TODO: further consideration on the fractional part
    }

    /**
     * @dev returns current reserve of ethers
         */
    function currentReserve() public view returns(uint256) {
        return totalPending + totalStaked + accountedUserRevenue - totalDebts - rewardDebts;  // TODO: ignore total debts?
    }

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
    function deposit(uint256 minToMint, uint256 deadline) external payable onlyValidTransaction(deadline) returns (uint256 minted) {
        amount = msg.value;
        require(amount > 0, "USR002"); // TODO: don't need to check zero mint?
        // TODO: check msgValue ceiling

        // TODO: to be optimized

        // mint uniIOTX
        _mint(minToMint);

        // stake
        _stake();

        // merge
        _merge();

        return toMint;
    }

    // TODO: to be modified
    function redeem(uint256 amount, uint256 maxToBurn, uint256 deadline) external onlyValidTransaction(deadline) nonReentrant returns(uint256 burned) {
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

    // TODO: Consider whitelisting KYC users?
    // TODO: Give an explanation of param minToMint
    function _mint(uint256 minToMint) private notZeroMint returns (uint256 minted) {
        accountedBalance += msg.value;

        toMint = _convertIotxToUniIotx(msg.value);
        if (toMint < minToMint) revert ExchangeRatioMismatch(minToMint, toMint);
        uniIOTX.mint(msg.sender, toMint);
        minted = toMint;

        totalPending += msg.value;

        emit MintEvent(msg.sender, minToMint, minted);
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


    /**
     * @dev Calculates uniIOTX amount based on IOTX amount for mint and burn operation,
     * aiming to keep the exchange ratio invariant to avoid user arbitrage.
     * Reference: https://github.com/RockX-SG/stake/blob/main/doc/uniETH_ETH2_0_Liquid_Staking_Explained.pdf
     */
    function _convertIotxToUniIotx(uint256 amountIOTX) private pure returns (uint256 amountUniIOTX) {
        uint256 totalSupply = uniIOTX.totalSupply();
        uint256 currentReserve = currentReserve();
        amountUniIOTX = defaultExchangeRatio * amountIOTX

        if (currentReserve > 0) { // avert division overflow
            amountUniIOTX = totalSupply * amountIOTX / currentReserve; // TODO: further consideration on the fractional part
        }
 }

}