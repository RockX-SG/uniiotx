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

import "./Roles.sol";
import "../interfaces/IIOTXClear.sol";
import "../interfaces/IIOTXStake.sol";
import "../interfaces/ISystemStake.sol";

contract IOTXClear is IIOTXClear, Initializable, PausableUpgradeable, AccessControlUpgradeable, ReentrancyGuardUpgradeable {
    // ---Use libraries---
    using Address for address payable;

    // ---External dependencies---
    address public systemStake;

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
        require(amount > 0 && amount % debtAmountBase == 0, "Invalid debt amount");
        _;
    }

    modifier onlyDebtToken(uint[] calldata tokenIds) {
        for (uint i; i < tokenIds.length; i++) {
            (uint tokenAmt, , , ,) = ISystemStake(systemStake).bucketOf(tokenIds[i]);
            require(tokenAmt == debtAmountBase, "Invalid token amount");
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

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /**
     * @dev This function initializes the contract
     */
    function initialize(
        address _systemStake,
        address _iotxStake,
        address _oracle
    ) public initializer  {
        // Init
        __Pausable_init();
        __AccessControl_init();
        __ReentrancyGuard_init();

        // Roles
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _setupRole(ROLE_PAUSER, msg.sender);
        _setupRole(ROLE_STAKER, _iotxStake);
        _setupRole(ROLE_ORACLE, _oracle);

        // Collaborative contracts
        systemStake = _systemStake;

        // Debt management parameters
        debtAmountBase = IIOTXStake(_iotxStake).redeemAmountBase();
    }

    /**
     * ======================================================================================
     *
     * VIEW FUNCTIONS
     *
     * ======================================================================================
     */

    /**
     * @dev This function is the 'IERC721Receiver' implement for receiving redeemed/unlocked NFT transferred by IOTXStake contract
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
     * @return The user's total reward that is available for future claims.
     * @dev The returned value includes the pending reward that hasn't been accounted yet.
     */
    function getReward(address account) external view returns (uint) {
        return userInfos[account].reward + _calcPendingReward(account);
    }

    /**
     * ======================================================================================
     *
     * EXTERNAL FUNCTIONS
     *
     * ======================================================================================
     */

    /**
     * @dev The contract 'IOTXStake' calls this function upon the user's redeeming request.
     * This function queues the redeemed amount as debt, which can be paid by withdrawal in FIFO order.
     */
    function joinDebt(address account, uint amount) external whenNotPaused onlyDebtAmount(amount) onlyRole(ROLE_STAKER) {
        // Update current user reward
        _updateUserReward(account);

        // Record new user debt
        _enqueueDebt(account, amount);
    }

    function updateDelegates(uint[] calldata tokenIds, address delegate) external whenNotPaused onlyRole(ROLE_ORACLE) {
        ISystemStake(systemStake).changeDelegates(tokenIds, delegate);

        emit DelegatesUpdated(tokenIds, delegate);
    }

    function unstake(uint[] calldata tokenIds) external whenNotPaused onlyDebtToken(tokenIds) onlyRole(ROLE_ORACLE) {
        if (tokenIds.length > 0) ISystemStake(systemStake).unstake(tokenIds);
    }

    function payDebts(uint[] calldata tokenIds) external whenNotPaused onlyDebtToken(tokenIds) onlyRole(ROLE_ORACLE) {
        uint totalTokenCntToPay = tokenIds.length;
        require(totalTokenCntToPay > 0 && totalDebts >= totalTokenCntToPay*debtAmountBase, "Invalid total principal for debt payment");
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

    function claimPrincipals(uint amount, address recipient) external nonReentrant whenNotPaused {
        // Check principal
        UserInfo storage info = userInfos[msg.sender];
        require(info.principal >= amount, "Insufficient accounted principal");

        // Transfer principal
        payable(recipient).sendValue(amount);
        info.principal -= amount;
        accountedBalance -= amount;

        emit PrincipalClaimed(msg.sender, recipient, amount);
    }

    function claimRewards(uint amount, address recipient) external nonReentrant whenNotPaused {
         // Update reward
        _updateUserReward(msg.sender);

        // Check reward
        UserInfo storage info = userInfos[msg.sender];
        require(info.reward >= amount, "Insufficient accounted reward");

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

    function _enqueueDebt(address account, uint amount) internal {
        // Add a debt in FIFO order
        rearIndex += 1;
        iotxDebts[rearIndex] = Debt({account:account, amount:amount});

        // Update debt states
        userInfos[account].debt += amount;
        totalDebts += amount;

        emit DebtQueued(account, amount, rearIndex);
    }

    function _peekNextDebt() internal view returns (Debt memory firstDebt) {
        firstDebt = iotxDebts[headIndex+1];
    }

    function _dequeueDebt() internal returns (Debt memory debt) {
        require(!_isEmptyQueue(), "Empty queue");
        uint firstDebt = headIndex+1;
        debt = iotxDebts[firstDebt];
        delete iotxDebts[firstDebt];
        headIndex += 1;
    }

    function _isEmptyQueue() internal view returns (bool) {
        return rearIndex == headIndex;
    }

    function _payNextDebt(uint amountToPay, uint[] memory tokenIds) internal {
        // Retrieve next debt
        uint nextDebtIndex = headIndex+1;
        Debt storage nextDebt = iotxDebts[nextDebtIndex];
        address account = nextDebt.account;

        // Withdraw NFT to this contract
        ISystemStake(systemStake).withdraw(tokenIds, payable(address(this)));
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

    function _updateSharedReward() internal {
        if (address(this).balance > accountedBalance && totalDebts > 0) {
            uint incrReward = address(this).balance - accountedBalance;
            rewardRate += incrReward * MULTIPLIER / totalDebts;
            accountedBalance = address(this).balance;
        }
    }

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
