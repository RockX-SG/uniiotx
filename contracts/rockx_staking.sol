pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

contract RockXStaking is Initializable, PausableUpgradeable, AccessControlUpgradeable {
    // errors
    error ZeroDelegates();

    enum BucketType {
        BucketType1,
        BucketType2,
        BucketType3
    }

    // track iotex debts to return to async caller
    struct Debt {
        address account;
        uint256 amount;
    }

    ISystemStakingContract public iotexSystemStakingContract;  // IoTeX system staking contract
    IMintableContract public xIOTEXAddress;               // xIOTEX token address
    IRockXRedeem public redeemContract;          // redeeming contract for user to pull iotexs

    uint256 private totalPending;               // total pending IOTEXs awaiting to be staked
    uint256 private totalDebts;             // track current unpaid debts


    // FIFO of debts from redeemFromDelegates
    mapping(uint256=>Debt) private iotexDebts;
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
    function setIOTEXSystemStakingContract(address _iotexSystemStakingContract) external onlyRole(DEFAULT_ADMIN_ROLE) {
        iotexSystemStakingContract = _iotexSystemStakingContract;

        emit SystemStakingContractSet(_iotexSystemStakingContract);
    }

    /**
     * @dev set xIOTEX token contract address
     */
    function setXIOTEXContractAddress(address _xIOTEXAddress) external onlyRole(DEFAULT_ADMIN_ROLE) {
        xIOTEXAddress = _xIOTEXAddress;

        emit XIOTEXContractSet(_xIOTEXAddress);
    }

    /**
     * @dev set redeem contract
     */
    function setRedeemContract(address _redeemContract) external onlyRole(DEFAULT_ADMIN_ROLE) {
        redeemContract = _redeemContract;

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
     * @dev return pending iotexs
     */
    function getPendingIotexs() external view returns (uint256) { return totalPending; }

    /**
     * @dev return current debts
     */
    function getCurrentDebts() external view returns (uint256) { return totalDebts; }

    /**
     * @dev return debt of index
     */
    function checkDebt(uint256 index) external view returns (address account, uint256 amount) {
        Debt memory debt = iotexDebts[index];
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
     * @dev mint uniIOTEX with IOTEX
     */
    function deposit() external payable returns (uint256 minted) {
        amount = msg.value;
        require(amount > 0, "USR002");

        // TODO: to be optimized

        // mint uniIOTEX
        _mint();

        // stake
        _stake();

        // merge
        _merge();

    return toMint;
    }

    // TODO: to be modified
    function redeemFromDelegates(uint256 iotexsToRedeem, uint256 maxToBurn, uint256 deadline) external nonReentrant returns(uint256 burned) {
        require(block.timestamp < deadline, "USR001");
        require(iotexsToRedeem % DEPOSIT_SIZE == 0, "USR005");

        uint256 totalXIOTEX = IERC20(xIOTEXAddress).totalSupply();
        uint256 xIOTEXToBurn = totalXIOTEX * iotexsToRedeem / currentReserve(); // TODO:
        require(xIOTEXToBurn <= maxToBurn, "USR004");

        // NOTE: the following procdure must keep exchangeRatio invariant:
        // transfer xETH from sender & burn
        IERC20(xETHAddress).safeTransferFrom(msg.sender, address(this), xIOTEXToBurn);
        IMintableContract(xIOTEXAddress).burn(xIOTEXToBurn);

        // queue ether debts
        _enqueueDebt(msg.sender, iotexsToRedeem);

        // return burned
        return xIOTEXToBurn;
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
        iotexDebts[lastDebt] = Debt({account:account, amount:amount});

        // track user debts
        userDebts[account] += amount;
        // track total debts
        totalDebts += amount;

        // log
        emit DebtQueued(account, amount);
    }

    function _dequeueDebt() internal returns (Debt memory debt) {
        require(lastDebt >= firstDebt, "SYS022");  // non-empty queue
        debt = iotexDebts[firstDebt];
        delete iotexDebts[firstDebt];
        firstDebt += 1;
    }

    /**
     * @dev pay debts for a given amount
     */
    function _payDebts(uint256 total) internal returns(uint256 amountPaid) {
        require(address(redeemContract) != address(0x0), "SYS023");

        // iotexs to pay
        for (uint i=firstDebt;i<=lastDebt;i++) {
            if (total == 0) {
                break;
            }

            Debt storage debt = iotexDebts[i];

            // clean debts
            uint256 toPay = debt.amount <= total? debt.amount:total;
            debt.amount -= toPay;
            total -= toPay;
            userDebts[debt.account] -= toPay;
            amountPaid += toPay;

            // transfer money to debt contract
            IRockXRedeem(redeemContract).pay{value:toPay}(debt.account);

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
        xIOTEXAddress.mint(msg.sender, toMint);
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
            iotexSystemStakingContract.merge(tokenIds,stakeDuration);
            theFirst = new uint256[](1);
            theFirst[0] = tokenIds[0];
            _pushLockedTokens(BucketType.BucketType2, theFirst);
        }

        count = _lockedTokenIdCount(BucketType.BucketType2);
        if (count >= mergeThreshold) {
            tokenIds = _removeLockedTokenIds(BucketType.BucketType2, mergeThreshold);
            iotexSystemStakingContract.merge(tokenIds,stakeDuration);
            theFirst = new uint256[](1);
            theFirst[0] = tokenIds[0];
            _pushLockedTokens(BucketType.BucketType3, theFirst);
        }
    }

    function _requestStake(BucketType bucketType, uint256 amount, uint256 count, address delegate) private{
        if (count == 0) return;
        totalAmount = amount*count;
        if (count == 1) {
            tokenId = iotexSystemStakingContract.stake{value:totalAmount}(stakeDuration, delegate);
            _addLockedTokenId(BucketType.bucketType, tokenId);
        } else {
            firstTokenId = iotexSystemStakingContract.stake{value:totalAmount}(amount,stakeDuration, delegate, count);
            _addLockedTokenIds(BucketType.bucketType, firstTokenId, count);
        }
        totalPending -= totalAmount;
    }

    function _nextDelegate() private returns (address) {
        delegate = delegates[nextDelegateIndex];
        nextDelegateIndex = (nextDelegateIndex + 1) % delegates.length;
        return delegate;
    }

    function _lockedTokenIdCount(BucketType bucketType) private returns (uint256) {
        // TODO:
        return 0;
    }

    function _removeLockedTokenIds(BucketType bucketType, uint256 count) private returns (uint256[]) {
        // TODO:
        return 0;
    }

    function _addLockedTokenId(BucketType bucketType, uint256 tokenId) private {
        // TODO:
        for (i = 0; i < tokensIds.length; i++) {

        }
    }

    function _addLockedTokenIds(BucketType bucketType, uint256 firstTokenId, uint256 count) private {
        // TODO:
    }

    /**
     * ======================================================================================
     *
     * ROCKX SYSTEM EVENTS
     *
     * ======================================================================================
     */
    event DelegateStopped(uint256 stoppedCount);
    event DebtQueued(address creditor, uint256 amountEther);
    event SystemStakingContractSet(address addr);
    event XIOTEXContractSet(address addr);
    event DepositEvent(address indexed from, uint256 amount);
    event RedeemContractSet(address addr);
}