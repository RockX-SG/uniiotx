// SPDX-License-Identifier: MIT
/*
 * ==================================================================
 * Copyright (C) 2023 Altstake Technology Pte. Ltd. (RockX)
 * This code is free software; you can redistribute it
 * and/or modify it under the terms of the GNU General Public License as
 * published by the Free Software Foundation; either version 3 of the License,
 * or (at your option) any later version.
 * This code is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * You should have received a copy of the GNU General Public License.
 * If not, see <http://www.gnu.org/licenses/>
 * ==================================================================
 */

pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/utils/Address.sol";

import "../interfaces/Roles.sol";
import "../interfaces/IIOTXClear.sol";
import "../interfaces/IIOTXStaking.sol";
import "../interfaces/ISystemStaking.sol";

contract IOTXClear is IIOTXClear, Initializable, PausableUpgradeable, AccessControlUpgradeable, ReentrancyGuardUpgradeable {
    using Address for address payable;

    address public systemStaking;

    uint private constant MULTIPLIER = 1e18;

    /**
     * @dev A struct type for recording users' debt, principal and reward.
     */
    struct UserInfo {
        // The remaining amount of the user's debt. Its value fluctuates due to several factors:
        // 1. It increases when a user request to redeem IOTXs from the 'IOTXStake' contract.
        // 2. It decreases when a debt payment is made.
        uint debt;
        // The claimable amount of the users' principal. Its value fluctuates due to several factors:
        // 1. It increases when a debt payment is made.
        // 2. It decreases when the user claims principal.
        uint principal;
        // The claimable amount of the users' reward. Its value fluctuates due to several factors:
        // 1. It increases When the delegate distributes rewards.
        // 2. It decreases when the user claim a reward.
        uint reward;
        // Users' reward rate depends on the shared 'rewardRate'.
        // It is updated whenever '_updateUserReward' is called.
        uint rewardRate;
    }

    /**
     * @dev A struct type for managing debt items in a FIFO queue.
     */
    struct Debt {
        address account;
        uint amount;
    }

    /**
     * @dev The balance synchronized from this contract fluctuates due to several factors:
     * 1. It increases when the Oracle triggers debt payment.
     * 2. It increases when rewards are distributed by delegates.
     * 3. It decreases when users claim their rewards or principal..
     */
    uint public accountedBalance;

    /**
     * @dev The total debt fluctuates due to several factors:
     * 1. It increases when a new debt item is added.
     * 2. It decreases when a debt payment is made.
     */
    uint public totalDebts;

    /**
     * @dev The accumulated reward rate is influenced by the incremental reward and total debt.
     * It's calculated using the following formula: rewardRate += incrReward * 1e18 / totalDebts
     */
    uint public rewardRate;

    /**
     * @dev The permissible amount of new debt should be in multiples of the base debt amount.
     * This base value is determined at contract initialization.
     * Once set, it remains immutable.
     */
    uint public debtAmountBase;

    /**
     * @dev Simulating a First-In-First-Out (FIFO) queue of debts.
     *
     * A 'joinDebt' request will add a new debt item to the end of the queue.
     * A 'payDebts' request will settle the debt from the front end of the queue.
     *
     * Here are the key features of this FIFO queue:
     * 1. Each debt item is given a unique index.
     * 2. The index for adding a new debt: rearIndex + 1.
     * 3. The index for the next debt payment is: headIndex + 1.
     * 4. The total number of added debts is represented by: rearIndex.
     * 5. The total number of paid debts is represented by: headIndex.
     * 6. The total number of unpaid debts can be calculated as: rearIndex - headIndex.
     */
    mapping(uint=>Debt) private iotxDebts;
    uint private headIndex;
    uint private rearIndex;

    /**
     * @dev The map for user information management:
     * 1. The KEY is the user's account address.
     * 2. The VALUE is of type UserInfo.
     */
    mapping(address => UserInfo) private userInfos;

    event DebtQueued(address account, uint amount, uint debtIndex);
    event DebtPaid(address account, uint amount, uint debtIndex);
    event Claimed(address claimer, address recipient, uint amount);
    event PrincipalClaimed(address claimer, address recipient, uint amount);
    event RewardClaimed(address claimer, address recipient, uint amount);
    event DelegatesUpdated(uint[] tokenIds, address delegate);

    modifier onlyDebtAmount(uint amount) {
        require(amount > 0 && amount % debtAmountBase == 0, "SYS003");  // Invalid debt amount
        _;
    }

    modifier onlyDebtToken(uint[] calldata tokenIds) {
        for (uint i; i < tokenIds.length; i++) {
            (uint tokenAmt, , , ,) = ISystemStaking(systemStaking).bucketOf(tokenIds[i]);
            require(tokenAmt == debtAmountBase, "USR007");  // Invalid token amount for debt payment
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

    /**
     * @dev This function is exclusively designed to receive staking rewards generated after the 'joinDebt' function is evoked.
     * Any IOTXs inadvertently sent to this contract will be considered as rewards.
     */
    receive() external payable { }

    /**
     * @dev This function pauses the contract
     */
    function pause() public onlyRole(ROLE_PAUSER) {
        _pause();
    }

    /**
     * @dev This function unpauses the contract
     */
    function unpause() public onlyRole(ROLE_PAUSER) {
        _unpause();
    }

    /**
     * @dev @custom:oz-upgrades-unsafe-allow constructor
     */
    constructor() {
        _disableInitializers();
    }

    /**
     * @dev This function initializes the contract
     */
    function initialize(
        address _systemStaking,
        address _iotxStaking,
        address _oracle,
        address _pauser
    ) public initializer  {
        // Init
        __Pausable_init();
        __AccessControl_init();
        __ReentrancyGuard_init();

        // Roles
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _setupRole(ROLE_PAUSER, _pauser);
        _setupRole(ROLE_STAKER, _iotxStaking);
        _setupRole(ROLE_ORACLE, _oracle);

        // Collaborative contracts
        systemStaking = _systemStaking;

        // Debt management parameters
        debtAmountBase = IIOTXStaking(_iotxStaking).redeemAmountBase();
    }


    /**
     * ======================================================================================
     *
     * VIEW FUNCTIONS
     *
     * ======================================================================================
     */

    /**
     * @dev This function is the 'IERC721Receiver' implement for receiving redeemed/unlocked NFT transferred by IOTXStaking contract
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
     * @return The balance that has been synchronized and accounted for this contract.
     */
    function getAccountedBalance() external view returns (uint) {
        return accountedBalance;
    }

    /**
     * @return The total amount of outstanding debts.
     */
    function getTotalDebts() external view returns (uint) {
        return totalDebts;
    }

    /**
     * @return The reward rate, which is accumulated and shared among users.
     */
    function getRewardRate() external view returns (uint) {
        return rewardRate;
    }

    /**
     * @dev The permissible amount of new debt should be in multiples of the base debt amount.
     * @return The base amount for joining a new debt item.
     */
    function getDebtAmountBase() external view returns (uint) {
        return debtAmountBase;
    }

    /**
     * @return The user's balance, comprising both rewards and principal, is available for future claims.
     */
    function getBalance(address account) external view returns (uint) {
        return userInfos[account].reward + _calcPendingReward(account) + userInfos[account].principal;
    }

    /**
     * @dev The returned value includes the pending reward that hasn't been accounted yet.
     * @return The user's reward that is available for future claims.
     */
    function getReward(address account) external view returns (uint) {
        return userInfos[account].reward + _calcPendingReward(account);
    }

    /**
     * @dev The amount of the paid debt will be added to the user's principal.
     * @return The user's principal that is available for future claims.
     */
    function getPrincipal(address account) external view returns (uint) {
        return userInfos[account].principal;
    }

    /**
     * @dev The amount of the paid debt will be added to the user's principal.
     * @return The user's remaining debt that needs payment.
     */
    function getDebt(address account) external view returns (uint) {
        return userInfos[account].debt;
    }

    /**
     * @return The total number of debt items, encompassing both paid and outstanding payments.
     */
    function getTotalDebtItemCount() external view returns (uint) {
        return rearIndex;
    }

    /**
     * @return The number of debt items that have been paid.
     */
    function getPaidDebtItemCount() external view returns (uint) {
        return headIndex;
    }

    /**
     * @return The number of debt items awaiting payment.
     */
    function getUnpaidDebtItemCount() external view returns (uint) {
        return rearIndex-headIndex;
    }

    /**
     * @param unpaidDebtIndex The index of the unpaid debt item in the FIFO queue
     * should range between (paidDebtItemCount, totalDebtItemCount].
     * @return The queued unpaid debt item.
     */
    function getUnpaidDebtItem(uint unpaidDebtIndex) external view returns (Debt memory) {
        if (headIndex < unpaidDebtIndex && unpaidDebtIndex <= rearIndex) {
            return iotxDebts[unpaidDebtIndex];
        }
        return Debt({account:0x0000000000000000000000000000000000000000, amount:0});
    }

    /**
     * @return The total queued unpaid debt items.
     */
    function getUnpaidDebtItems() external view returns (Debt[] memory) {
        if (headIndex < rearIndex) {
            uint count = rearIndex - headIndex;
            Debt[] memory items = new Debt[](count);
            uint firstIndex = headIndex+1;
            for (uint i = 0; i < count; i++) {
                items[i] = iotxDebts[firstIndex+i];
            }
            return items;
        }
        return new Debt[](0);
    }

    /**
     * @return The user information of the given account.
     */
    function getUserInfo(address account) external view returns (UserInfo memory) {
        return userInfos[account];
    }

    /**
     * ======================================================================================
     *
     * EXTERNAL FUNCTIONS FOR STAKER
     *
     * ======================================================================================
     */

    /**
     * @dev The contract 'IOTXStaking' calls this function upon the user's redeeming request.
     * This function queues the redeemed amount as debt, which can be paid by withdrawal in FIFO order.
     */
    function joinDebt(address account, uint amount) external whenNotPaused onlyDebtAmount(amount) onlyRole(ROLE_STAKER) {
        // Update current user reward
        _updateUserReward(account);

        // Record new user debt
        _enqueueDebt(account, amount);
    }


    /**
     * ======================================================================================
     *
     * EXTERNAL FUNCTIONS FOR ORACLE
     *
     * ======================================================================================
     */

    /**
     * @dev This function updates the delegates of token IDs.
     */
    function updateDelegates(uint[] calldata tokenIds, address delegate) external whenNotPaused onlyRole(ROLE_ORACLE) {
        ISystemStaking(systemStaking).changeDelegates(tokenIds, delegate);

        emit DelegatesUpdated(tokenIds, delegate);
    }

    /**
     * @dev This function unstakes unlocked tokens, allowing them to be used for future debt payment.
     */
    function unstake(uint[] calldata tokenIds) external nonReentrant whenNotPaused onlyDebtToken(tokenIds) {
        if (tokenIds.length > 0) ISystemStaking(systemStaking).unstake(tokenIds);
    }

    /**
     * @dev This function withdraws the specified tokens for debt payment.
     */
    function payDebts(uint[] calldata tokenIds) external nonReentrant whenNotPaused onlyDebtToken(tokenIds) {
        uint totalTokenCntToPay = tokenIds.length;
        require(totalTokenCntToPay > 0 && totalDebts >= totalTokenCntToPay*debtAmountBase, "USR008");  // Invalid total principal for debt payment
        uint paidTokenCnt;
        while (paidTokenCnt < totalTokenCntToPay) {
            // Peek next debt
            Debt memory nextDebt = _peekNextDebt();

            // Update current user reward
            _updateUserReward(nextDebt.account);

            // Determine token IDs
            uint remainedTokenCntToPay = totalTokenCntToPay - paidTokenCnt;
            uint maxTokenCntToPay = nextDebt.amount / debtAmountBase;
            uint tokenCntToPay = (maxTokenCntToPay > remainedTokenCntToPay) ? remainedTokenCntToPay: maxTokenCntToPay;
            uint amountToPay = debtAmountBase * tokenCntToPay;

            uint[] memory tokenIdsToPay = new uint[](tokenCntToPay);
            for (uint i = 0; i < tokenCntToPay; i++) {
                tokenIdsToPay[i] = tokenIds[paidTokenCnt+i];
            }

            // Pay the user's debt
            _payNextDebt(amountToPay, tokenIdsToPay);
            paidTokenCnt += tokenCntToPay;
        }
    }


    /**
     * ======================================================================================
     *
     * EXTERNAL FUNCTIONS FOR USERS
     *
     * ======================================================================================
     */

    /**
     * @dev This function allows users to claim their claimable asset to the specified recipient.
     * The assets that the user can claim consist of rewards and principal.
     * The rewards will be subtracted first, followed by the principal.
     */
    function claim(uint amount, address recipient) external nonReentrant whenNotPaused {
        // Update reward
        _updateUserReward(msg.sender);

        // Check user quota
        UserInfo storage info = userInfos[msg.sender];
        require(info.principal + info.reward >= amount, "USR009");  // Insufficient accounted asset

        // Transfer asset
        if (info.reward >= amount) {
            info.reward -= amount;
        } else {
            info.principal -= amount - info.reward;
            info.reward = 0;
        }
        accountedBalance -= amount;
        payable(recipient).sendValue(amount);

        emit Claimed(msg.sender, recipient, amount);
    }

    /**
     * @dev This function allows users to claim their principals to the specified recipient.
     */
    function claimPrincipals(uint amount, address recipient) external nonReentrant whenNotPaused {
        // Check principal
        UserInfo storage info = userInfos[msg.sender];
        require(info.principal >= amount, "USR004");  // Insufficient accounted principal

        // Transfer principal
        info.principal -= amount;
        accountedBalance -= amount;
        payable(recipient).sendValue(amount);

        emit PrincipalClaimed(msg.sender, recipient, amount);
    }

    /**
     * @dev This function allows users to claim their rewards to the specified recipient.
     */
    function claimRewards(uint amount, address recipient) external nonReentrant whenNotPaused {
         // Update reward
        _updateUserReward(msg.sender);

        // Check reward
        UserInfo storage info = userInfos[msg.sender];
        require(info.reward >= amount, "USR005");  // Insufficient accounted reward

        // Transfer reward
        info.reward -= amount;
        accountedBalance -= amount;
        payable(recipient).sendValue(amount);

        emit RewardClaimed(msg.sender, recipient, amount);
    }


    /**
     * ======================================================================================
     *
     * INTERNAL FUNCTIONS
     *
     * ======================================================================================
     */

    /**
     * @dev This function appends a debt item to the FIFO queue.
     */
    function _enqueueDebt(address account, uint amount) internal {
        // Add a debt in FIFO order
        rearIndex += 1;
        iotxDebts[rearIndex] = Debt({account:account, amount:amount});

        // Update debt states
        userInfos[account].debt += amount;
        totalDebts += amount;

        emit DebtQueued(account, amount, rearIndex);
    }

    /**
     * @dev This function retrieves a debt item from the FIFO queue for the next payment,
     * without removing it from the queue.
     */
    function _peekNextDebt() internal view returns (Debt memory firstDebt) {
        firstDebt = iotxDebts[headIndex+1];
    }

    /**
     * @dev This function removes a debt item from the FIFO queue.
     */
    function _dequeueDebt() internal returns (Debt memory debt) {
        require(!_isEmptyQueue(), "SYS004");  // Empty queue
        uint firstDebt = headIndex+1;
        debt = iotxDebts[firstDebt];
        delete iotxDebts[firstDebt];
        headIndex += 1;
    }

    /**
     * @dev This function checks if the FIFO queue is empty.
     */
    function _isEmptyQueue() internal view returns (bool) {
        return rearIndex == headIndex;
    }

    /**
     * @dev This function pays the next debt based on the FIFO queue.
     */
    function _payNextDebt(uint amountToPay, uint[] memory tokenIds) internal {
        // Retrieve next debt
        uint nextDebtIndex = headIndex+1;
        Debt storage nextDebt = iotxDebts[nextDebtIndex];
        address account = nextDebt.account;

        // Withdraw NFT to this contract
        ISystemStaking(systemStaking).withdraw(tokenIds, payable(address(this)));
        accountedBalance += amountToPay;

        // Update debt states
        UserInfo storage userInfo = userInfos[account];
        userInfo.debt -= amountToPay;
        userInfo.principal += amountToPay;

        nextDebt.amount -= amountToPay;
        totalDebts -= amountToPay;

        // Remove debt entry if it has been fully paid
        if (nextDebt.amount == 0) {
            _dequeueDebt();
        }

        emit DebtPaid(account, amountToPay, nextDebtIndex);
    }

    /**
     * @dev This function updates the user's rewards according to their most recent debt
     */
    function _updateUserReward(address account) internal {
        _updateSharedReward();
        UserInfo storage info = userInfos[account];
        info.reward += (rewardRate - info.rewardRate) * info.debt / MULTIPLIER;
        info.rewardRate = rewardRate;
    }

    /**
     * @dev This function synchronizes this contract balance and updates the shared 'rewardRate'.
     */
    function _updateSharedReward() internal {
        // Sync shared pending rewards and update the global rewardRate
        // Note: The _calcIncrRewardRate function can bring an calculation accuracy loss.
        uint incrRewardRate = _calcIncrRewardRate();
        uint accountedReward;
        if (incrRewardRate > 0) {
            rewardRate += incrRewardRate;
            accountedReward = incrRewardRate * totalDebts / MULTIPLIER;
        }

        // Update the accountedBalance, considering the calculation accuracy loss
        // to ensure that piecemeal rewards do not accumulate over time into a significant amount.
        accountedBalance = accountedBalance + accountedReward;
    }

    /**
     * @dev This function calculates the user's pending reward that has not yet been synchronized and accounted for.
     */
    function _calcPendingReward(address account) internal view returns (uint) {
        // Calculate the global rewardRate
        uint globalRewardRate = rewardRate + _calcIncrRewardRate();

       // Calculate personal pending rewards
        UserInfo memory info = userInfos[account];
        if (globalRewardRate > info.rewardRate) {
            return (globalRewardRate - info.rewardRate) * info.debt / MULTIPLIER;
        }
        return 0;
    }

    /**
    * @dev This function calculates the pending reward rate that has not yet been synchronized and accounted for.
     */
    function _calcIncrRewardRate() internal view returns (uint) {
        // Sync shared pending rewards
        if (address(this).balance > accountedBalance && totalDebts > 0) {
            uint incrReward = address(this).balance - accountedBalance;
            // Note: The accuracy loss occurs in the calculation result for division rounds towards zero.
            // This means that uint(5) / uint(2) == uint(2).
            return incrReward * MULTIPLIER / totalDebts;
        }
        return 0;
    }
}
