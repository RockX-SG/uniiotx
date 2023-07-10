// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "interfaces/IUniIOTX.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";

import "interfaces/IIOTXStake.sol";
import "../interfaces/ISystemStake.sol";

contract IOTXClear is IIOTXClear, Initializable, PausableUpgradeable, ReentrancyGuardUpgradeable, OwnableUpgradeable, IERC721Receiver {
    using SafeERC20 for IERC20;
    using Address for address payable;

    // External dependencies
    ISystemStake public systemStake;
    IIOTXStake public iotxStake;

    // Constants
    uint public constant MULTIPLIER = 1e18;

    // Type declarations

    struct UserInfo {
        uint debt;   // IOTX value a user requests to redeem but hasn't been withdrawn
        uint reward;  // The claimable reward which is distributed after redeeming requested
        uint rewardRate; // Latest rewardRate assigned upon reward update
    }

    struct Debt {
        address account;
        uint amount;
    }

    // State variables
    uint public accountedBalance;

    uint public totalDebts;  // Current total unpaid debts caused by redeeming requests
    uint public rewardRate;    // Accumulated reward rate calculated against totalDebts, multiplied by MULTIPLIER


    // Simulating a FIFO queue of debts
    mapping(uint=>Debt) public iotxDebts;   // Index -> Debt
    uint public firstIndex;
    uint public lastIndex;

    // User infos
    mapping(address => UserInfo) public userInfos; // account -> info

    // Events
    event DebtAdded(address account, uint amount);
    event DebtPaid(address account, uint amount);
    event RewardClaimed(address claimer, address recipient, uint amount);

    // Errors
    error InsufficientReward(uint expected, uint available);
    error DebtAmountMismatched(uint toPayAmount, uint queuedAmount);

    /**
     * ======================================================================================
     *
     * SYSTEM SETTINGS
     *
     * ======================================================================================
     */

    // Todo: Tune it properly
    /**
     * @dev initialization
     */
    function initialize(
        address _iotxStakeAddress,
        address _systemStakeAddress,
    ) initializer public {
        __Ownable_init();
        __Pausable_init();
        __ReentrancyGuard_init();

        _grantRole(IOTX_STAKE_ROLE, _iotxStakeAddress);

        systemStake = _systemStakeAddress;
        iotxStake = _ _iotxStakeAddress;

        firstIndex = 1;
        lastIndex = 0;
    }

    /**
     * ======================================================================================
     *
     * VIEW FUNCTIONS
     *
     * ======================================================================================
     */

    /**
     * @dev IERC721Receiver implement for receiving redeemed/unlocked NFT transferred by IOTXStake contract
         */
    function onERC721Received(
        address, // operator
        address, // from
        uint, // tokenId
        bytes calldata // data
    ) external pure override returns (bytes4) {
            return this.onERC721Received.selector;
    }


    /**
     * @dev Return user reward which is available for future claim
     */
    function getReward(address acount) external returns (uint) {
        return userInfos[account].reward;
    }

    /**
     * ======================================================================================
     *
     * EXTERNAL FUNCTIONS
     *
     * ======================================================================================
     */

    /**
     * @dev IOTXStake contract calls this function upon the user's redeeming request.
     * This function queues the redeemed amount as debt, which can be paid by withdrawal in FIFO order.
     */
    function joinDebt(address account, uint amount) public shenNotPaused onlyRole(IOTX_STAKE_ROLE) {
        // Update current user reward
        _updateReward(account);

        // Record new user debt
        _addDebt(account, amount);
    }

    function updateDelegates(uint[] tokenIds, address delegate) external whenNotPaused onlyRole(ORACLE_ROLE) {
        systemStake.changeDelegates(tokenIds, delegate);
    }

    function unstake(uint[] calldata tokenIds) external whenNotPaused onlyRole(ORACLE_ROLE) {
        if (tokenIds.length > 0) systemStake.unstake(tokenIds);
    }

    // Todo: Maybe optimize the implementation, including introducing necessary validations.
    function withdraw(uint[] tokenIds) external whenNotPaused onlyRole(ORACLE_ROLE) {
        for (uint i = 0; i < tokenIds.length); i++ {
            address account = _payDebt(tokenIds[i]);
            _updateReward(account);
        }
    }

    /**
     * @dev This function accounts for new reward increments (if there are any),
     * and accumulate reward rate to the latest value.
     * Then it returns the updated claimable reward for the given account
     */
    function updateReward(address acount) external returns (uint) {
        _updateReward(account);
        return userInfos[account].reward;
    }

    function claimRewards(uint amount, address recipient) external nonReentrant whenNotPaused {
         // Update reward
        _updateReward(msg.sender);

        // Check reward
        UserInfo info = userInfos[msg.sender];
        if (info.reward < amount) revert InsufficientReward(amount, info.reward);

        // Transfer reward
        payable(recipient).sendValue(amount);
        info.reward -= amount;
        accountedBalance -= amount;

        emit RewardClaimed(msg.sender, recipient, amount);
    }

    /**
    * ======================================================================================
    *
    * INTERNAL FUNCTIONS
    *
    * ======================================================================================
    */

    function _addDebt(address account, uint amount) internal {
        // Add a debt in FIFO order
        lastIndex += 1;
        iotxDebts[lastIndex] = Debt({account:account, amount:amount});

        // Update debt states
        userInfos[account].debt += amount;
        totalDebts += amount;

        emit DebtAdded(account, amount);
    }

    function _payDebt(uint tokenId) internal returns (address account) {
        // Pick a debt in FIFO order
        Debt storage firstDebt = iotxDebts[firstIndex];
        account = firstDebt.account;

        // Validate NFT amount
        (amount, _, _, _, _,) = systemStake.bucketOf(tokenId);
        if (amount != firstDebt.amount) revert DebtAmountMismathced(amount, firstDebt.amount);

        // Withdraw NFT to user account
        systemStake.withdraw(tokenId, account);

        // Update debt states
        userInfos[account].debt -= amount;
        totalDebts -= amount;
        delete iotxDebts[firstIndex];
        firstIndex += 1;

        event DebtPaid(account, amount);
    }

    function _updateReward(address acount) internal {
        _updateReward();
        UserInfo storage info = userInfos[account];
        info.reward += (rewardRate - info.rewardRate) * info.debt / MULTIPLIER;
        info.rewardRate = rewardRate;
    }

    // Todo: Confirm whether consider an extra manager fee in IOTXClear contract (We also count it in IOTXStake contract).
    function _updateReward() internal {
        if (address(this).balance > accountedBalance && totalDebts > 0) {
            uint incrReward = address(this).balance - accountedBalance;
            rewardRate += incrReward * MULTIPLIER / totalDebts;
            accountedBalance = address(this).balance;
        }
    }
}
