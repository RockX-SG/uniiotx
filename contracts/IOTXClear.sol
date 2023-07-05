// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

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
    uint256 public constant MULTIPLIER = 1e18;

    // Type declarations

    struct UserInfo {
        uint256 accSharePoint; // share starting point
        uint256 debt;   // user's debt
        uint256 reward;  // user's reward which is available for claim
    }

    struct Debt {
        address account;
        uint256 amount;
    }

    // State variables
    uint256 public accountedBalance;

    uint256 public totalDebts;             // Track current unpaid debts

    // FIFO of debts
    mapping(uint256=>Debt) private iotxDebts;   // Index -> Debt
    uint256 public firstDebt;
    uint256 public lastDebt;

    mapping(address => UserInfo) public userInfos; // account -> info

    // Events
    event DebtJoined(address account, uint256 amount);
    event DebtLeft(address account, uint256 amount);
    event RewardClaimed(address claimer, address recipient, uint256 amount);

    // Errors
    error InsufficientReward(uint256 expected, uint256 available);
    error DebtAmountMismatched(uint256 toPayAmount, uint256 queuedAmount);

    /**
     * ======================================================================================
     *
     * CONFIGURATION FUNCTIONS
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

        firstDebt = 1;
        lastDebt = 0;
    }

    /**
     * ======================================================================================
     *
     * VIEW FUNCTIONS
     *
     * ======================================================================================
     */

    /**
     * @dev Return debt of index
     */
    function checkDebt(uint256 index) external view returns (address account, uint256 amount) {
        Debt memory debt = iotxDebts[index];
        return (debt.account, debt.amount);
    }

    /**
     * @dev Return debt queue index
     */
    function getDebtQueue() external view returns (uint256 first, uint256 last) {
        return (firstDebt, lastDebt);
    }

    /**
     * @dev IERC721Receiver implement for receiving redeemed/unlocked NFT transferred by IOTXStake contract
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
     * EXTERNAL FUNCTIONS
     *
     * ======================================================================================
     */

    function joinDebt(address account, uint256 amount) public shenNotPaused onlyRole(IOTX_STAKE_ROLE) {
        // Update current user reward
        _updateReward(account);

        // Record new user debt
        _enqueueDebt(account, amount);

        emit DebtJoined(account, amount);
    }

    function updateDelegates(uint256[] tokenIds, address delegate) external whenNotPaused onlyRole(ORACLE_ROLE) {
        systemStake.changeDelegates(tokenIds, delegate);
    }

    function unstake(uint256[] calldata tokenIds) external whenNotPaused onlyRole(ORACLE_ROLE) {
        if (tokenIds.length > 0) systemStake.unstake(tokenIds);
    }

    // Todo: Maybe optimize the implementation, including introducing necessary validations.
    function withdraw(uint256[] tokenIds) external whenNotPaused onlyRole(ORACLE_ROLE) {
        for (uint256 i = 0; i < tokenIds.length); i++ {
            address account = _payDebt(tokenIds[i]);
            _updateReward(account);
        }
    }

    function updateReward() external {
        _updateReward();
    }

    /**
     * @dev Return updated user reward which is available for a later claim
     */
    function updateReward(address acount) external returns (uint256) {
        _updateReward(account);
        return userInfos[account].reward;
    }

    function claimRewards(uint256 amount, address recipient) external nonReentrant whenNotPaused {
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

    function _enqueueDebt(address account, uint256 amount) internal {
        lastDebt += 1;
        iotxDebts[lastDebt] = Debt({account:account, amount:amount});

        userInfos[account].debt += amount;
        totalDebts += amount;
    }

    function _payDebt(uint256 tokenId) internal returns (address account) {
        // Pick a debt in FIFO order
        Debt storage firstDebt = iotxDebts[firstDebt];
        account = firstDebt.account;

        // Validate NFT amount
        (amount, _, _, _, _,) = systemStake.bucketOf(tokenId);
        if (amount != firstDebt.amount) revert DebtAmountMismathced(amount, firstDebt.amount);

        // Withdraw NFT to user account
        systemStake.withdraw(tokenId, account);

        // Update debt states
        userInfos[account].debt -= amount;
        totalDebts -= amount;
        delete iotxDebts[firstDebt];
        firstDebt += 1;

        event DebtLeft(account, amount);
    }

    function _updateReward(address acount) internal {
        _updateReward();
        UserInfo storage info = userInfos[account];
        info.reward += (accShare - info.accSharePoint) * info.debt / MULTIPLIER;
        info.accSharePoint = accShare;
    }

    // Todo: Confirm whether consider an extra manager fee in IOTXClear contract (We also count it in IOTXStake contract).
    function _updateReward() internal {
        if (address(this).balance > accountedBalance && totalDebts > 0) {
            uint256 incrReward = address(this).balance - accountedBalance;
            accShare += incrReward * MULTIPLIER / totalDebts;
            accountedBalance = address(this).balance;
        }
    }
}
