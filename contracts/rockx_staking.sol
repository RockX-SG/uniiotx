pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

contract RockXStaking is Initializable, PausableUpgradeable, AccessControlUpgradeable {
    // errors
    error ZeroDelegates();

    // track iotex debts to return to async caller
    struct Debt {
        address account;
        uint256 amount;
    }

    address public iotexSystemStakingContract;  // IoTeX system staking contract
    address public xIOTEXAddress;               // xIOTEX token address
    address public redeemContract;          // redeeming contract for user to pull iotexs

    uint256 private totalPending;               // total pending IOTEXs awaiting to be staked
    uint256 private totalDebts;             // track current unpaid debts


    // FIFO of debts from redeemFromDelegates
    mapping(uint256=>Debt) private iotexDebts;
    uint256 private firstDebt;
    uint256 private lastDebt;
    mapping(address=>uint256) private userDebts;    // debts from user's perspective

    // delegates
    address[] private delegateRegistry;

    // next delegate index
    uint256 private nextDelegateIndex;

    // default stake params
    uint256 private stakeAmount01 = 10000;
    uint256 private stakeAmount02 = 100000;
    uint256 private stakeAmount03 = 1000000;
    uint256 private stakeDuration = 3600*91;

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
    function registerDelegates(address[] delegates) external onlyRole(REGISTRY_ROLE) {
        if (delegates.length == 0) revert ZeroDelegates();
        delete delegates;
        uint256 len = delegates.length;
        for (uint256 i = 0; i < len; i++) {
            delegateRegistry.push(delegates[i]);
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
        return delegateRegistry.length;
    }

    /**
     * @dev return all registered delegates
     */
    function getAllDelegates() external view returns (address[]) {
        return delegateRegistry;
    }

    /**
     * ======================================================================================
     *
     * EXTERNAL FUNCTIONS
     *
     * ======================================================================================
     */

    /**
     * @dev mint xIOTEX with IOTEX
     */
    function mint() external payable returns (uint256 minted) {
        amount = msg.value;
        require(amount > 0, "USR002");

        // merge
        _merge();

        // TODO: to be optimized

        // mint xIOTEX
        uint256 toMint = 1 * amount; // default exchange ratio 1:1
        IMintableContract(xIOTEXAddress).mint(msg.sender, toMint);
        totalPending += amount;
        emit DepositEvent(msg.sender, amount);

        // stake
        count = totalPending / stakeAmount03;
        _stake(stakeAmount03, count);

        count = totalPending / stakeAmount02;
        _stake(stakeAmount02, count);

        count = totalPending / stakeAmount01;
        _stake(stakeAmount01, count);

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

    function _merge() private {
        // TODO: to be implemented
    }

    function _stake(uint256 amount, uint256 count) private {
        if (count == 0) return;
        // TODO: to be implemented
        totalPending -= amount*count;
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