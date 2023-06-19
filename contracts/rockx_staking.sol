pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

contract RockXStaking is Initializable, PausableUpgradeable, AccessControlUpgradeable {
    // errors
    error MintZero();

    address public iotexSystemStakingContract;  // IoTeX system staking contract
    address public xIOTEXAddress;               // xIOTEX token address
    address public redeemContract;          // redeeming contract for user to pull iotexs

    uint256 private totalPending;               // total pending IOTEXs awaiting to be staked
    uint256 private totalDebts;             // track current unpaid debts

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
        if  (amount == 0) revert MintZero();

        // merge
        _merge();

        // TODO: to be optimized

        // mint xIOTEX
        uint256 toMint = 1 * amount; // default exchange ratio 1:1
        IMintableContract(xIOTEXAddress).mint(msg.sender, toMint);
        totalPending += amount;
        emit DepositEvent(msg.sender, amount);

        // stake
        _stake();

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
        etherDebts[lastDebt] = Debt({account:account, amount:amount});

        // track user debts
        userDebts[account] += amount;
        // track total debts
        totalDebts += amount;

        // log
        emit DebtQueued(account, amount);
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

    function _stake() private {
        // TODO: to be implemented
    }

    /**
     * ======================================================================================
     *
     * ROCKX SYSTEM EVENTS
     *
     * ======================================================================================
     */
    event DebtQueued(address creditor, uint256 amountEther);
    event SystemStakingContractSet(address addr);
    event XIOTEXContractSet(address addr);
    event DepositEvent(address indexed from, uint256 amount);
    event RedeemContractSet(address addr);
}