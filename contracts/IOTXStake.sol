pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";

import "interfaces/IIOTXClear.sol"
import "interfaces/IUniIOTX.sol"
import "../IIOTXStake.sol"

contract IOTXStake is Initializable, PausableUpgradeable, AccessControlUpgradeable, IERC721Receiver {
    // track IOTX debts to return to async caller
    struct Debt {
        address account;
        uint256 amount
    }

    enum BucketType {
        BucketType1,
        BucketType2,
        BucketType3
    }

    IIOTXStake public iotxStake;
    IUniIOTX public uniIOTX;
    IIOTXClear public iotxClear;

    uint256 private totalPending;               // total pending IOTXs awaiting to be staked
    uint256 private totalDebts;             // track current unpaid debts

    // FIFO of debts from redeemFromDelegates
    mapping(uint256=>Debt) private IOTXDebts;
    uint256 private firstDebt;
    uint256 private lastDebt;
    mapping(address=>uint256) private userDebts;    // debts from user's perspective

    // delegates
    address[] private delegates;

    // next delegate index
    uint256 private nextDelegateIndex;

    // default stake params
    uint256 private stakeAmount01 = 10000;
    uint256 private stakeAmount02 = 100000;
    uint256 private stakeAmount03 = 1000000;
    uint256 private stakeDuration = 3600*91;

    uint256 private mergeThreshold = 10;

    // token ids
    struct tokenIds {
        uint256 size;
        // startId => count
        mapping(uint256 => uint256) data;
    }
    mapping(BucketType => tokenIds) public lockedTokenIds;
    mapping(BucketType => tokenIds) public unlockedTokenIds;
    mapping(BucketType => tokenIds) public unstakedTokenIds;
    mapping(BucketType => tokenIds) public withdrawedTokenIds;

    // Events
    event DelegateStopped(uint256 stoppedCount);
    event DebtQueued(address creditor, uint256 amountEther);
    event SystemStakingContractSet(address addr);
    event XIOTXContractSet(address addr);
    event DepositEvent(address indexed from, uint256 amount);
    event RedeemContractSet(address addr);

    // Errors
    error ZeroDelegates();

    /**
     * ======================================================================================
     *
     * SYSTEM SETTINGS, OPERATED VIA OWNER(DAO/TIMELOCK)
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

        // init default values
        firstDebt = 1;
        lastDebt = 0;
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
        iotxStake = _SystemStakingContract;

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
     * @dev return current debts
     */
    function getCurrentDebts() external view returns (uint256) { return totalDebts; }

    /**
     * @dev return debt of index
     */
    function checkDebt(uint256 index) external view returns (address account, uint256 amount) {
        Debt memory debt = IOTXDebts[index];
        return (debt.account, debt.amount);
    }

    /**
     * @dev return debt queue index
     */
    function getDebtQueue() external view returns (uint256 first, uint256 last) {
        return (firstDebt, lastDebt);
    }

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
        require(amount > 0, "USR002");

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
    function redeemFromDelegates(uint256 IOTXsToRedeem, uint256 maxToBurn, uint256 deadline) external nonReentrant returns(uint256 burned) {
        require(block.timestamp < deadline, "USR001");
        require(IOTXsToRedeem % DEPOSIT_SIZE == 0, "USR005");

        uint256 totalXIOTX = IERC20(uniIOTX).totalSupply();
        uint256 xIOTXToBurn = totalXIOTX * IOTXsToRedeem / currentReserve(); // TODO:
        require(xIOTXToBurn <= maxToBurn, "USR004");

        // NOTE: the following procdure must keep exchangeRatio invariant:
        // transfer xETH from sender & burn
        IERC20(xETHAddress).safeTransferFrom(msg.sender, address(this), xIOTXToBurn);
        IUniIOTX(uniIOTX).burn(xIOTXToBurn);

        // queue ether debts
        _enqueueDebt(msg.sender, IOTXsToRedeem);

        // return burned
        return xIOTXToBurn;
    }

    /**
     * ======================================================================================
     *
     * INTERNAL FUNCTIONS
     *
     * ======================================================================================
     */

    function _enqueueDebt(address account, uint256 amount) internal {
        // debt is paid in FIFO queue
        lastDebt += 1;
        IOTXDebts[lastDebt] = Debt({account:account, amount:amount});

        // track user debts
        userDebts[account] += amount;
        // track total debts
        totalDebts += amount;

        // log
        emit DebtQueued(account, amount);
    }

    function _dequeueDebt() internal returns (Debt memory debt) {
        require(lastDebt >= firstDebt, "SYS022");  // non-empty queue
        debt = IOTXDebts[firstDebt];
        delete IOTXDebts[firstDebt];
        firstDebt += 1;
    }

    /**
     * @dev pay debts for a given amount
     */
    function _payDebts(uint256 total) internal returns(uint256 amountPaid) {
        require(address(iotxClear) != address(0x0), "SYS023");

        // IOTXs to pay
        for (uint i=firstDebt;i<=lastDebt;i++) {
            if (total == 0) {
                break;
            }

            Debt storage debt = IOTXDebts[i];

            // clean debts
            uint256 toPay = debt.amount <= total? debt.amount:total;
            debt.amount -= toPay;
            total -= toPay;
            userDebts[debt.account] -= toPay;
            amountPaid += toPay;

            // transfer money to debt contract
            IIOTXClear(iotxClear).pay{value:toPay}(debt.account);

            // dequeue if cleared
            if (debt.amount == 0) {
                _dequeueDebt();
            }
        }

        totalDebts -= amountPaid;

        // track balance
        _balanceDecrease(amountPaid);
    }

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
            iotxStake.merge(tokenIds,stakeDuration);
            theFirst = new uint256[](1);
            theFirst[0] = tokenIds[0];
            _pushLockedTokens(BucketType.BucketType2, theFirst);
        }

        count = _lockedTokenIdCount(BucketType.BucketType2);
        if (count >= mergeThreshold) {
            tokenIds = _removeLockedTokenIds(BucketType.BucketType2, mergeThreshold);
            iotxStake.merge(tokenIds,stakeDuration);
            theFirst = new uint256[](1);
            theFirst[0] = tokenIds[0];
            _pushLockedTokens(BucketType.BucketType3, theFirst);
        }
    }

    function _requestStake(BucketType bucketType, uint256 amount, uint256 count, address delegate) private{
        if (count == 0) return;
        totalAmount = amount*count;
        if (count == 1) {
            tokenId = iotxStake.stake{value:totalAmount}(stakeDuration, delegate);
            _addLockedTokenIds(BucketType.bucketType, tokenId,1);
        } else {
            startTokenId = iotxStake.stake{value:totalAmount}(amount,stakeDuration, delegate, count);
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
}