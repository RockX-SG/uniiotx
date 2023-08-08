/*
 * ==================================================================
 * Copyright (C) 2023 Altstake Technology Pte. Ltd. (RockX)
 * CAUTION: THESE CODES HAVE NOT BEEN AUDITED
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
    // ---Use libraries---
    using Address for address payable;

    // ---External dependencies---
    address public systemStaking;

    // ---Constants---
    uint public constant MULTIPLIER = 1e18;

    // ---Type declarations---

    struct UserInfo {
        uint debt;   // IOTX value a user requests to redeem but hasn't been withdrawn
        uint principal;  // Claimable principal which is from debt payment
        uint reward;  // Claimable reward which is distributed after redeeming requested
        uint rewardRate;  // Latest rewardRate assigned upon reward update
    }

    struct Debt {
        address account;
        uint amount;
    }

    // ---State variables---
    uint public accountedBalance;

    uint public totalDebts;  // Current total unpaid debts caused by redeeming requests
    uint public rewardRate;    // Accumulated reward rate calculated against totalDebts, multiplied by MULTIPLIER


    uint public debtAmountBase;

    // Simulating a First-In-First-Out (FIFO) queue of debts
    // Index for adding new debt: rearIndex + 1
    // Index for the next debt payment: headIndex + 1
    // Total count of added debts: rearIndex
    // Total count of paid debts: headIndex
    // Total count of unpaid debts: rearIndex - headIndex
    mapping(uint=>Debt) public iotxDebts;   // Index -> Debt
    uint public headIndex;
    uint public rearIndex;

    // User infos
    mapping(address => UserInfo) public userInfos; // account -> info

    // ---Events---
    event DebtQueued(address account, uint amount, uint debtIndex);
    event DebtPaid(address account, uint amount, uint debtIndex);
    event PrincipalClaimed(address claimer, address recipient, uint amount);
    event RewardClaimed(address claimer, address recipient, uint amount);
    event DelegatesUpdated(uint[] tokenIds, address delegate);

    // ---Modifiers---
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
        address _oracle
    ) public initializer  {
        // Init
        __Pausable_init();
        __AccessControl_init();
        __ReentrancyGuard_init();

        // Roles
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _setupRole(ROLE_PAUSER, msg.sender);
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
    function unstake(uint[] calldata tokenIds) external whenNotPaused onlyDebtToken(tokenIds) onlyRole(ROLE_ORACLE) {
        if (tokenIds.length > 0) ISystemStaking(systemStaking).unstake(tokenIds);
    }

    /**
     * @dev This function withdraws the specified tokens for debt payment.
     */
    function payDebts(uint[] calldata tokenIds) external whenNotPaused onlyDebtToken(tokenIds) onlyRole(ROLE_ORACLE) {
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
     * @dev This function allows users to claim their principals to the specified recipient.
     */
    function claimPrincipals(uint amount, address recipient) external nonReentrant whenNotPaused {
        // Check principal
        UserInfo storage info = userInfos[msg.sender];
        require(info.principal >= amount, "USR004");  // Insufficient accounted principal

        // Transfer principal
        payable(recipient).sendValue(amount);
        info.principal -= amount;
        accountedBalance -= amount;

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
        if (address(this).balance > accountedBalance && totalDebts > 0) {
            uint incrReward = address(this).balance - accountedBalance;
            rewardRate += incrReward * MULTIPLIER / totalDebts;
            accountedBalance = address(this).balance;
        }
    }

    /**
     * @dev This function calculates the user's pending reward that has not yet been synchronized and accounted for.
     */
    function _calcPendingReward(address account) internal view returns (uint) {
        UserInfo memory info = userInfos[account];
        if (info.debt > 0 && address(this).balance > accountedBalance) {
            uint incrReward = address(this).balance - accountedBalance;
            uint _rewardRate = rewardRate + incrReward * MULTIPLIER / totalDebts;
            return (_rewardRate - info.rewardRate) * info.debt / MULTIPLIER;
        }
        return 0;
    }
}
