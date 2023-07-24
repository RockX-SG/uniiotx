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
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "@openzeppelin/contracts/utils/Address.sol";

import "./Roles.sol";
import "../interfaces/ISystemStake.sol";

contract IOTXClear is Initializable, PausableUpgradeable, AccessControlUpgradeable, ReentrancyGuardUpgradeable, IERC721Receiver {
    // ---Use libraries---
    using Address for address payable;

    // ---External dependencies---
    ISystemStake public systemStake;

    // ---Constants---
    uint public constant MULTIPLIER = 1e18;

    // ---Type declarations---

    struct UserInfo {
        uint debt;   // IOTX value a user requests to redeem but hasn't been withdrawn
        uint reward;  // The claimable reward which is distributed after redeeming requested
        uint rewardRate; // Latest rewardRate assigned upon reward update
    }

    struct Debt {
        address account;
        uint amount;
    }

    // ---State variables---
    uint public accountedBalance;

    uint public totalDebts;  // Current total unpaid debts caused by redeeming requests
    uint public rewardRate;    // Accumulated reward rate calculated against totalDebts, multiplied by MULTIPLIER


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
    event DebtAdded(address account, uint amount);
    event DebtPaid(address account, uint amount);
    event RewardClaimed(address claimer, address recipient, uint amount);

    /**
     * ======================================================================================
     *
     * SYSTEM SETTINGS
     *
     * ======================================================================================
     */

    /**
    * @notice This function is exclusively designed to receive staking rewards generated after the 'joinDebt' function is evoked.
     * Any IOTXs inadvertently sent to this contract will be considered as rewards.
     */
    receive() external payable { }

    /**
     * @dev initialization
     */
    function initialize(
        address _systemStakeAddress,
        address _iotxStakeAddress,
        address _oracleAddress
    ) public initializer  {
        __Pausable_init();
        __ReentrancyGuard_init();

        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ROLE_STAKE, _iotxStakeAddress);
        _grantRole(ROLE_ORACLE, _oracleAddress);

        systemStake = ISystemStake(_systemStakeAddress);
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
     * @dev Return the user's total reward that is available for future claims.
     * @notice The returned value includes the pending reward that hasn't been accounted yet.
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
     * @dev IOTXStake contract calls this function upon the user's redeeming request.
     * This function queues the redeemed amount as debt, which can be paid by withdrawal in FIFO order.
     */
    function joinDebt(address account, uint amount) external whenNotPaused onlyRole(ROLE_STAKE) {
        // Update current user reward
        _updateUserReward(account);

        // Record new user debt
        _addDebt(account, amount);
    }

    function updateDelegates(uint[] calldata tokenIds, address delegate) external whenNotPaused onlyRole(ROLE_ORACLE) {
        systemStake.changeDelegates(tokenIds, delegate);
    }

    function unstake(uint[] calldata tokenIds) external whenNotPaused onlyRole(ROLE_ORACLE) {
        if (tokenIds.length > 0) systemStake.unstake(tokenIds);
    }

    function payDebts(uint[] calldata tokenIds) external whenNotPaused onlyRole(ROLE_ORACLE) {
        for (uint i = 0; i < tokenIds.length; i++) {
            // Pop a debt in FIFO order
            Debt storage nextDebt = iotxDebts[headIndex+1];
            address account = nextDebt.account;

            // Validate NFT amount against the debt
            uint tokenId = tokenIds[i];
            (uint amount, , , ,) = systemStake.bucketOf(tokenId);
            require(amount == nextDebt.amount, "Debt amount mismatch");

            // Update current user reward
            _updateUserReward(account);

            // Settle the user's debt
            _payDebt(account, amount, tokenId);
        }
    }

    function claimRewards(uint amount, address recipient) external nonReentrant whenNotPaused {
         // Update reward
        _updateUserReward(msg.sender);

        // Check reward
        UserInfo storage info = userInfos[msg.sender];
        require(info.reward >= amount, "Insufficient reward");

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
        rearIndex += 1;
        iotxDebts[rearIndex] = Debt({account:account, amount:amount});

        // Update debt states
        userInfos[account].debt += amount;
        totalDebts += amount;

        emit DebtAdded(account, amount);
    }

    function _payDebt(address account, uint amount, uint tokenId) internal {
        // Withdraw NFT to user account
        systemStake.withdraw(tokenId, payable(account));

        // Update debt states
        userInfos[account].debt -= amount;
        totalDebts -= amount;
        delete iotxDebts[headIndex+1];
        headIndex += 1;

        emit DebtPaid(account, amount);
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
