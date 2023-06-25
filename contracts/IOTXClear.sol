// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "interfaces/IUniIOTX.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";

import "../interfaces/ISystemStake.sol"

contract IOTXClear is IIOTXClear, Initializable, PausableUpgradeable, ReentrancyGuardUpgradeable, OwnableUpgradeable, IERC721Receiver {
    using SafeERC20 for IERC20;
    using Address for address payable;

    // External dependencies
    ISystemStake public systemStake;

    // Type declarations
    struct Debt {
        address account;
        uint256 amount;
    }

    uint256 private totalDebts;             // track current unpaid debts

    // FIFO of debts from redeemFromDelegates
    mapping(uint256=>Debt) private IOTXDebts;
    uint256 private firstDebt;
    uint256 private lastDebt;
    mapping(address=>uint256) private userDebts;    // debts from user's perspective

    mapping(address=>uint256) private balances;
    uint256 private totalBalance;

    // Events
    event DebtQueued(address creditor, uint256 amountEther);
    event Paied(address account, uint256 amount);
    event Claimed(address indexed _from, address indexed _to, uint256 _value);
    event SystemStakingContractSet(address addr);

    /**
     * ======================================================================================
     *
     * SYSTEM SETTINGS
     *
     * ======================================================================================
     */

    /**
     * @dev initialization
     */
    function initialize() initializer public {
        __Ownable_init();
        __Pausable_init();
        __ReentrancyGuard_init();

        // init default values
        firstDebt = 1;
        lastDebt = 0;
    }

    /**
     * @dev set eth deposit contract address
         */
    function setSystemStakingContract(address _SystemStakingContract) external onlyRole(DEFAULT_ADMIN_ROLE) {
        systemStake = _SystemStakingContract;

        emit SystemStakingContractSet(_SystemStakingContract);
    }

    // some convenient method to help show their claimable in wallet
    function name() external pure returns (string memory) { return "RockX Claimable IOTX"; }
    function symbol() external pure returns (string memory) { return "redeemIOTX"; }
    function decimals() external pure returns (uint8) { return 18; }
    function totalSupply() external view returns (uint256) { return totalBalance; }
    function balanceOf(address account) external view returns(uint256) { return balances[account]; }

    /**
     * ======================================================================================
     *
     * VIEW FUNCTIONS
     *
     * ======================================================================================
     */

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
     * ======================================================================================
     *
     * INTERNAL FUNCTIONS
     *
     * ======================================================================================
     */

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
     * PUBLIC FUNCTIONS
     *
     * ======================================================================================
     */

    function joinDebt(address account, uint256 amount) public {
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

    /**
     * ======================================================================================
     *
     * EXTERNAL FUNCTIONS
     *
     * ======================================================================================
     */

    function changeDelegates(uint256[] tokenId, address delegate) external whenNotPaused onlyRole(ORACLE_ROLE) {
        systemStake.changeDelegates(tokenId, delegate);
    }

    function unstake(uint256[] calldata tokenIds) external whenNotPaused onlyRole(ORACLE_ROLE) {
        if (tokenIds.length > 0) systemStake.unstake(tokenIds);
    }

    function claim(address to, uint256 amount) public nonReentrant returns (bool success) {
        // check
        require(balances[msg.sender] >= amount, "INSUFFICIENT_BALANCE");

        // modify
        balances[msg.sender] -= amount;
        totalBalance -= amount;
        payable(to).sendValue(amount);

        // log
        emit Claimed(msg.sender, to, amount);
        return true;
    }

    /**
     * @dev pay debts from rockx staking contract
     */
    function payDebt(address account) external override payable {
        balances[account] += msg.value;
        totalBalance += msg.value;

        // log
        emit Paied(account, msg.value);
    }


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
    * ======================================================================================
    *
    * PRIVATE FUNCTIONS
    *
    * ======================================================================================
    */



}
