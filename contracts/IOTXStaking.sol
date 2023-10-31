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

import "../interfaces/Roles.sol";
import "../interfaces/IIOTXClear.sol";
import "../interfaces/IUniIOTX.sol";
import "../interfaces/IIOTXStaking.sol";
import "../interfaces/ISystemStaking.sol";

contract IOTXStaking is IIOTXStaking, Initializable, PausableUpgradeable, AccessControlUpgradeable, ReentrancyGuardUpgradeable {
    address public systemStaking;
    address public uniIOTX;
    address public iotxClear;

    uint private constant DEFAULT_EXCHANGE_RATIO = 1;
    uint private constant MULTIPLIER = 1e18;

    /**
     * @dev The global delegate for upcoming deposit activities.
     */
    address public globalDelegate;

    /**
     * @dev The startAmount, commonRatio and sequenceLength collectively constitute a geometric sequence
     * for the staking amount that can be committed to the IoTeX network through our liquid staking protocol.
     * Each value in this sequence must be a valid bucket amount, as predefined by the IoTeX network.
     * Each value in this sequence corresponds to a unique indexed level within the range of [0, sequenceLength-1]
     * Once set, the sequence remains immutable.
     */
    uint public startAmount;
    uint public commonRatio;
    uint public sequenceLength;

    /**
     * @dev Users can deposit any amount of IOTXs, but can only redeem amounts that are in multiples of the redeemAmountBase.
     * This value is determined at contract initialization using the formula:
     * redeemAmountBase = startAmount * (commonRatio ** (sequenceLength-1))
     * Once set, it remains immutable.
     */
    uint public redeemAmountBase;

    /**
     * @dev The global stake duration for upcoming deposit activities.
     * This value is determined at contract initialization.
     * Once set, it remains immutable.
     */
    uint public stakeDuration;

    /**
     * @dev Token queue map for token ID management:
     * 1. The KEY corresponds to the bucket amount level defined as above.
     * 2. The VALUE is a dynamic array of token IDs.
     */
    mapping(uint => uint[]) private tokenQueues;

    /**
     * @dev The number of redeemed tokens can change only when a redeeming operation is performed.
     * In addition, the staked token count at the highest staking level can be calculated using this formula:
     * tokenQueues[sequenceLength-1].length - redeemedTokenCount
     */
    uint public redeemedTokenCount;

    /**
     * @dev The balance synchronized from this contract fluctuates due to several factors:
     * 1. It increases when users deposit IOTXs for liquid staking service.
     * 2. It increases when rewards are distributed by delegates.
     * 3. It decreases when pending IOTXs are staked with delegates.
     */
    uint public accountedBalance;

    /**
     * @dev The total pending IOTXs fluctuates due to several factors:
     * 1. It increases when users deposit IOTXs for liquid staking service.
     * 2. It increases when users' rewards are compounded..
     * 3. It increases when the manager fee is withdrawn.
     * 4. It decreases when pending IOTXs are staked with delegates.
     */
    uint public totalPending;

    /**
     * @dev The total staked IOTXs fluctuates due to several factors:
     * 1.It increases when pending IOTXs are staked with delegates.
     * 2. It decreases when users request to redeem IOTXs.
     */
    uint public totalStaked;

    /**
     * @dev The manager fee share ranges from 0 to 1000, as regulated by the default Admin.
     */
    uint public managerFeeShares;

    /**
     * @dev The accounted user rewards will increase when the rewards are divided.
     * The total is equivalent to the amount of automatic compounding.
     */
    uint public accountedUserReward;

    /**
     * @dev The accounted manger rewards will increase when the rewards are divided.
     * However, the quantity will decrease when the manager's fee is deducted.
     */
    uint public accountedManagerReward;

    event Minted(address user, uint minted);
    event Redeemed(address user, uint burned, uint[] tokenIds);
    event Staked(uint firstTokenId, uint amount, address delegate, uint count);
    event Merged(uint fromLevel, uint toLevel, uint amount);
    event RewardUpdated(uint amount);
    event ManagerFeeSharesSet(uint shares);
    event ManagerFeeWithdrawed(uint amount, uint minted, address recipient);
    event GlobalDelegateSet(address delegate);
    event DelegatesUpdated(uint[] tokenIds, address delegate);

    modifier onlyValidTransaction(uint deadline) {
        require(deadline > block.timestamp, "USR001");  // Transaction expired
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
     * @dev This function is exclusively designed to receive staking rewards.
     * The 'deposit' function should be invoked whenever users wish to stake IOTXs.
     * Any IOTXs accidentally sent to this contract will be considered as rewards.
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
     * @dev This function initializes the contract.
     */
    function initialize(
        address _systemStaking,
        address _uniIOTX,
        address _iotxClear,
        address _oracle,
        address _pauser,
        uint _startAmount,
        uint _commonRatio,
        uint _sequenceLength,
        uint _stakeDuration
    ) public initializer {
        // Init
        __Pausable_init();
        __AccessControl_init();
        __ReentrancyGuard_init();

        // Roles
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _setupRole(ROLE_FEE_MANAGER, msg.sender);
        _setupRole(ROLE_PAUSER, _pauser);
        _setupRole(ROLE_ORACLE, _oracle);

        // Collaborative contracts
        systemStaking = _systemStaking;
        uniIOTX = _uniIOTX;
        iotxClear = _iotxClear;

        // Immutable staking variables
        startAmount = _startAmount;
        commonRatio = _commonRatio;
        sequenceLength = _sequenceLength;
        stakeDuration = _stakeDuration;
        redeemAmountBase = startAmount * (commonRatio ** (sequenceLength-1));

        // Validate bucket type info
        for (uint level = 0; level < _sequenceLength; level++) {
            uint amount = _startAmount * (_commonRatio**level);
            bool isActive = ISystemStaking(systemStaking).isActiveBucketType(amount, _stakeDuration);
            require(isActive, "SYS001");  // Inactive bucket type
        }
    }

    /**
     * @dev This function sets manager's fee in range [0, 1000]
     */
    function setManagerFeeShares(uint shares) external onlyRole(DEFAULT_ADMIN_ROLE)  {
        require(shares <= 1000, "SYS002");  // Manager fee shares out of range
        managerFeeShares = shares;

        emit ManagerFeeSharesSet(shares);
    }


    /**
     * ======================================================================================
     *
     * VIEW FUNCTIONS
     *
     * ======================================================================================
     */

    /**
     * @dev This function is the 'IERC721Receiver' implement for receiving staking NFT
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
     * @dev The factors that affect the returned result are the same as those of the 'currentReserve' function.
     * @return ratio The current exchange ratio of uniIOTX to IOTX, multiplied by 1e18.
     */
    function exchangeRatio() external view returns (uint ratio) {
        return _exchangeRatio();
    }

    /**
     * @dev The returned amount is determined by the following contributions:
     * 1. User deposits/stakes their principal.
     * 2. Rewards generated from delegation, excluding manager rewards, are added to the reserve.
     * 3. The manager fee is withdrawn and included in the totalPending amount.
     * 4. When users make a 'redeem' call, the totalStaked amount decreases.
     * 5. The Oracle service should regularly call the 'updateReward' function to account for new rewards.
     * @return The current reserve of IOTXs in our liquid staking protocol.
     */
    function currentReserve() external view returns(uint) {
        return _currentReserve();
    }

    /**
     * @return The array of IOTX amounts that remain constant and active for each staking activity.
     */
    function getStakeAmounts() external view returns (uint[] memory) {
        uint[] memory amounts = new uint[](sequenceLength);
        for (uint level = 0; level < sequenceLength; level++) {
            uint amount = startAmount * (commonRatio**level);
            amounts[level] = amount;
        }
        return amounts;
    }

    /**
     * @return The global delegate that is utilized for upcoming staking requests.
     */
    function getGlobalDelegate() external view returns (address) {
        return globalDelegate;
    }

    /**
     * @notice The amount to be redeemed must be in multiples of the 'redeemAmountBase'.
     * @return The base amount that remains constant for each redeeming request.
     */
    function getRedeemAmountBase() external view returns (uint) {
        return redeemAmountBase;
    }

    /**
     * @return The stake duration in seconds that remains constant for each staking request.
     */
    function getStakeDuration() external view returns (uint) {
        return stakeDuration;
    }

    /**
     * @return The balance that has been synchronized and accounted for this contract.
     */
    function getAccountedBalance() external view returns (uint) {
        return accountedBalance;
    }

    /**
     * @return the pending reward that has not yet been synchronized and accounted for.
     */
    function getPendingReward() external view returns(uint) {
        return address(this).balance - accountedBalance;
    }

    /**
     * @return The total amount of IOTX awaiting staking, including the pending user reward.
     */
    function getTotalPending() external view returns (uint) {
        (uint pendingUserReward, ) = _calcPendingReward();
        return totalPending + pendingUserReward;
    }

    /**
     * @return The total amount of staked IOTX.
     */
    function getTotalStaked() external view returns (uint) {
        return totalStaked;
    }

    /**
     * @return The current manager fee shares.
     */
    function getManagerFeeShares() external view returns (uint) {
        return managerFeeShares;
    }

    /**
     * @return The amount of the user's shared reward, including the pending portion.
     */
    function getUserReward() external view returns (uint) {
        (uint userRewardIncr, ) = _calcPendingReward();
        return accountedUserReward + userRewardIncr;
    }

    /**
     * @return The amount of the manager's reward, including the pending portion.
     */
    function getManagerReward() external view returns (uint) {
        (, uint pendingManagerReward) = _calcPendingReward();
        return accountedManagerReward + pendingManagerReward;
    }

    /**
     * @return The amount of the user's shared reward that has been automatically compounded for upcoming staking requests.
     */
    function getCompoundedAmount() external view returns (uint) {
        return accountedUserReward;
    }

    /**
     * @dev If an invalid tokenQueueIndex is given, a zero value will be returned.
     * @param tokenQueueIndex The token queue index falls within the range of [0, sequenceLength).
     * @return The length of the queried token queue.
     */
    function getTokenQueueLength(uint tokenQueueIndex) external view returns (uint) {
        if (tokenQueueIndex < sequenceLength) {
            return tokenQueues[tokenQueueIndex].length;
        }
        return 0;
    }

    /**
     * @dev If an invalid tokenQueueIndex is given, a zero value will be returned.
     * @param tokenQueueIndex The token queue index falls within the range of [0, sequenceLength).
     * @return count The current staked token count of the specified token queue.
     */
    function getStakedTokenCount(uint tokenQueueIndex) external view returns (uint count) {
        if (tokenQueueIndex < sequenceLength) {
            uint queueLen = tokenQueues[tokenQueueIndex].length;
            if (tokenQueueIndex == sequenceLength-1) {
                count = queueLen - redeemedTokenCount;
            } else {
                count = queueLen;
            }
        }
    }

    /**
     * @return The number of redeemed token IDs.
     */
    function getRedeemedTokenCount() external view returns (uint) {
        return redeemedTokenCount;
    }

    /**
     * @dev If an invalid tokenQueueIndex and tokenIndex are given, a zero value will be returned.
     * @param tokenQueueIndex The token queue index falls within the range of [0, sequenceLength).
     * @param tokenIndex The token index is within the range of [0, tokenQueue.Length) in the specified token queue..
     * @return tokenId The queried token ID.
     */
    function getTokenId(uint tokenQueueIndex, uint tokenIndex) external view returns (uint tokenId) {
        if (tokenQueueIndex < sequenceLength && tokenIndex < tokenQueues[tokenQueueIndex].length) {
            return tokenQueues[tokenQueueIndex][tokenIndex];
        }
        return 0;
    }

    /**
     * @dev If an invalid tokenQueueIndex is given, an empty array will be returned.
     * @param tokenQueueIndex The token queue index falls within the range of [0, sequenceLength)
     * @return The staked token IDs for the specified index.
     */
    function getStakedTokenIds(uint tokenQueueIndex) external view returns (uint[] memory) {
        if (tokenQueueIndex < sequenceLength) {
            uint[] memory tq = tokenQueues[tokenQueueIndex];
            uint queueLen = tq.length;
            if (tokenQueueIndex == sequenceLength-1) {
                uint stakedCount = queueLen - redeemedTokenCount;
                uint[] memory tokenIds = new uint[](stakedCount);
                for (uint i = 0; i < stakedCount; i++) {
                    tokenIds[i] = tq[redeemedTokenCount+i];
                }
                return tokenIds;
            } else {
                return tq;
            }
        }
        return new uint[](0);
    }

    /**
     * @dev If param values are given, an empty array will be returned.
     * @param tokenQueueIndex The token queue index falls within the range of [0, sequenceLength)
     * @param i, j The index values for i and j should satisfy the following conditions:
     * 1. i < j && j < tokenQueues[tokenQueueIndex].length.
     * 2. i >= redeemedTokenCount if tokenQueueIndex == sequenceLength-1.
     * @return An [i, j) slice of staked token IDs for the specified index.
     */
    function getStakedTokenIdSlice(uint tokenQueueIndex, uint i, uint j) external view returns (uint[] memory) {
        if (tokenQueueIndex < sequenceLength && i < j) {
            uint[] memory tq = tokenQueues[tokenQueueIndex];
            uint[] memory tokenIds = new uint[](j-i);

            if (j <= tq.length) {
                for (uint k = 0; k < j-i; k++) {
                    tokenIds[k] = tq[i+k];
                }

                bool validI = true;
                if (tokenQueueIndex == sequenceLength-1 && i < redeemedTokenCount) {
                    validI = false;
                }

                if (validI) {
                    return tokenIds;
                }
            }
        }
        return new uint[](0);
    }

    /**
     * @return An array of already redeemed/unlocked token id, which is indexed from 0 in this contract.
     */
    function getRedeemedTokenIds() external view returns (uint[] memory) {
        uint[] memory tq = tokenQueues[sequenceLength-1];
        uint[] memory tokenIds = new uint[](redeemedTokenCount);
        for (uint i = 0; i < redeemedTokenCount; i++) {
            tokenIds[i] = tq[i];
        }
        return tokenIds;
    }

    /**
     * @dev If an invalid tokenQueueIndex is given, an empty array will be returned.
     * @dev It recommended to check the value of 'redeemedTokenCount' beforehand to prevent the passed j from going out of range.
     * @param i, j The valid index values for i and j are determined by this conditional check: i < j && j <= redeemedTokenCount
     * @return An [i, j) slice of already redeemed/unlocked token id, which is indexed from 0 in this contract.
     */
    function getRedeemedTokenIdSlice(uint i, uint j) external view returns (uint[] memory) {
        if (i < j && j <= redeemedTokenCount) {
            uint[] memory tq = tokenQueues[sequenceLength-1];
            uint[] memory tokenIds = new uint[](j-i);
            for (uint k = 0; k < j-i; k++) {
                tokenIds[k] = tq[i+k];
            }
            return tokenIds;
        }
        return new uint[](0);
    }


    /**
     * ======================================================================================
     *
     * EXTERNAL FUNCTIONS FOR ORACLE
     *
     * ======================================================================================
     */

    /**
     * @dev This function sets the global delegate for upcoming deposit activities.
     */
    function setGlobalDelegate(address delegate) external whenNotPaused onlyRole(ROLE_ORACLE) {
        globalDelegate = delegate;

        emit GlobalDelegateSet(delegate);
    }

    /**
     * @dev This function updates the delegates of token IDs.
     */
    function updateDelegates(uint[] calldata tokenIds, address delegate) external whenNotPaused onlyRole(ROLE_ORACLE) {
        ISystemStaking(systemStaking).changeDelegates(tokenIds, delegate);

        emit DelegatesUpdated(tokenIds, delegate);
    }

    /**
     * @dev This function stakes any pending IOTXs and merges staked buckets when conditions are fulfilled.
     * It synchronizes the most recently distributed reward before conducting the staking operation.
     */
    function stake() external whenNotPaused onlyRole(ROLE_ORACLE) {
        // Synchronize rewards
        _updateReward();

        // Proceed with the staking process.
        _stakeAtTopLevel();
        _stakeAndMergeAtSubLevel();
    }

    /**
     * @dev This function distributes recently accrued rewards from delegates among users and the manager.
     * It also automatically reinvests the users' share into totalPending for upcoming staking activities.
     */
    function updateReward() external onlyRole(ROLE_ORACLE) {
       _updateReward();
    }


    /**
     * ======================================================================================
     *
     * EXTERNAL FUNCTIONS FOR USERS
     *
     * ======================================================================================
     */

    /**
     * @dev This function mints uniIOTXs for the user, stakes any pending IOTXs and
     * merges staked buckets when conditions are fulfilled.
     * @return minted The quantity of minted uniIOTXs
     */
    function deposit(uint deadline) external payable nonReentrant whenNotPaused onlyValidTransaction(deadline) returns (uint minted) {
        // Verify the transferred value
        require(msg.value > 0, "USR002");  // Invalid deposit amount
        accountedBalance += msg.value;

        // Mint uniIOTX, keeping a consistent exchange ratio
        _updateReward();
        minted = _mint();

        // Proceed with the staking process.
        _stakeAtTopLevel();
        _stakeAndMergeAtSubLevel();
    }

    /**
     * @dev This function unlocks staked bucket(s) and subsequently calls the 'IOTXClear' contract to record the debt.
     * @param iotxsToRedeem The number of IOTXs to redeem must be a multiple of the accepted amount of redeeming base.
     * @return burned the quantity of burned uniIOTXs.
     */
    function redeem(uint iotxsToRedeem, uint deadline) external nonReentrant whenNotPaused onlyValidTransaction(deadline) returns(uint burned) {
        burned = _redeem(iotxsToRedeem);
    }


    /**
     * ======================================================================================
     *
     * EXTERNAL FUNCTIONS FOR FEE_MANAGER
     *
     * ======================================================================================
     */

    /**
     * @dev This function handles manager reward in this way:
     * 1. Mint uniIOTXs to the given recipient based on the given IOTX amount;
     * 2. Shift the corresponding amount of accountedManagerReward to totalPending.
     */
    function withdrawManagerFee(uint amount, address recipient) external nonReentrant onlyRole(ROLE_FEE_MANAGER)  {
        // Update the reward to help maintain a consistent exchange ratio
        _updateReward();

        require(amount <= accountedManagerReward, "USR006");  // Insufficient accounted manager reward

        uint toMint = _convertIotxToUniIOTX(amount);
        IUniIOTX(uniIOTX).mint(recipient, toMint);

        accountedManagerReward -= amount;
        totalPending += amount;

        emit ManagerFeeWithdrawed(amount, toMint, recipient);
    }


    /**
     * ======================================================================================
     *
     * INTERNAL FUNCTIONS
     *
     * ======================================================================================
     */

    /**
     * @dev This function mints uniIOTXs for the user based on their sent value and the latest exchange ratio.
     */
    function _mint() internal returns (uint minted) {
        uint toMint = _convertIotxToUniIOTX(msg.value);
        IUniIOTX(uniIOTX).mint(msg.sender, toMint);
        minted = toMint;

        totalPending += msg.value;

        emit Minted(msg.sender, minted);
    }

    /**
     * @dev This function stakes IOTXs at the highest level, utilizing the maximum staking amount,
     * without incorporating any merging behaviors.
     */
    function _stakeAtTopLevel() internal {
        // Determine values of stake params
        uint level = sequenceLength-1;
        (uint amount, uint count) = _getStakeAmountAndCount(level);
        if (count == 0) return;

        // Perform stake
        _doStake(level, amount, count);
    }

    /**
     * @dev This function stakes IOTXs at lower levels, using smaller staking amounts.
     * It may incorporate merging behaviors if conditions are met.
     */
    function _stakeAndMergeAtSubLevel() internal {
        uint nextLevel = sequenceLength-2;
        while (totalPending >= startAmount) {
            nextLevel = _tryStake(nextLevel);
        }
    }

    /**
     * @dev This function verifies if the pending IOTXs are adequate for staking at the specified level.
     * If adequate , it triggers the '_doStake' function to execute the actual staking.
     * Otherwise, it will return a lower try level.
     */
    function _tryStake(uint tryLevel) internal returns (uint nextLevel) {
        // Verify if there is an adequate total of pending IOTX and ascertain the values of stake parameters.
        (uint amount, uint count) = _getStakeAmountAndCount(tryLevel);
        if (count == 0) {
            return tryLevel-1;
        }

        uint stakedCount = tokenQueues[tryLevel].length;
        if ((count+stakedCount) >= commonRatio) {
            count = commonRatio-stakedCount;
        }

        // Perform stake
        _doStake(tryLevel, amount, count);

        // Merge tokens if possible
        if (count+stakedCount == commonRatio) _merge(tryLevel);

        return tryLevel;
    }

    /**
     * @dev This function carries out the actual staking activity.
     */
    function _doStake(uint level, uint amount, uint count) internal {
        // Calculate total amount
        uint totalAmount = amount * count;

        // Call system stake service
        uint firstTokenId = ISystemStaking(systemStaking).stake{value:totalAmount}(amount, stakeDuration, globalDelegate, count);

        // Record minted & staked tokens
        uint[] storage tq = tokenQueues[level];
        for (uint j = 0; j < count; j++) {
            tq.push(firstTokenId+j);
        }

        // Update fund status
        totalPending -= totalAmount;
        totalStaked  += totalAmount;
        accountedBalance -= totalAmount;

        emit Staked(firstTokenId, amount, globalDelegate, count);
    }

    /**
     * @dev This function computes and return the available staking amount and count at the specified level,
     * considering the current total pending amount.
     */
    function _getStakeAmountAndCount(uint level) internal view returns(uint amount, uint count) {
        amount = startAmount * (commonRatio**level);
        count = totalPending / amount;
    }

    /**
     * @dev This function combines lower-level tokens into upper-level ones.
     */
    function _merge(uint fromLevel) internal {
        uint steps;

        for (uint i = fromLevel; i < sequenceLength-1; i++) {
            // Check merge condition
            uint[] storage tq = tokenQueues[i];
            if (tq.length < commonRatio) break;

            // Call system merge service
            // All tokens will be merged into the first token in tokenIdsToMerge
            // Reference: https://github.com/iotexproject/iip13-contracts/blob/main/src/SystemStaking.sol#L302
            ISystemStaking(systemStaking).merge(tq, stakeDuration);

            // Move the merged tokens to upper queue
            uint[] storage tqUpper = tokenQueues[i+1];
            tqUpper.push(tq[0]);
            delete tokenQueues[i];

            steps++;
        }

        uint toLevel = fromLevel + steps;
        uint mergeAmount = startAmount * (commonRatio**toLevel);

        emit Merged(fromLevel, toLevel, mergeAmount);
    }

    /**
     * @dev This function unlocks staked buckets and triggers the 'IOTXClear' contract to record corresponding debts.
     */
    function _redeem(uint iotxsToRedeem) internal returns(uint burned) {
        // Check redeem condition
        require(iotxsToRedeem >= redeemAmountBase && iotxsToRedeem % redeemAmountBase == 0, "USR003");  // Invalid redeem amount

        // Update the reward to help maintain a consistent exchange ratio
        _updateReward();

        // Burn uniIOTXs
        uint toBurn = _convertIotxToUniIOTX(iotxsToRedeem);
        IUniIOTX(uniIOTX).burnFrom(msg.sender, toBurn);
        burned = toBurn;
        totalStaked -= iotxsToRedeem;

        // Extract tokens to unlock
        uint [] memory tq = tokenQueues[sequenceLength-1];
        uint count = iotxsToRedeem / redeemAmountBase;
        uint[] memory tokenIdsToUnlock = new uint[](count);
        for (uint i = 0; i < count; i++) {
            tokenIdsToUnlock[i] = tq[redeemedTokenCount];
            redeemedTokenCount++;
        }

        // Call system unlock service
        ISystemStaking(systemStaking).unlock(tokenIdsToUnlock);

        // Transfer unlocked tokens to IOTXClear contract
        for (uint i = 0; i < count; i++) {
            ISystemStaking(systemStaking).safeTransferFrom(address(this), address(iotxClear), tokenIdsToUnlock[i]);
        }

        // Record corresponding amount of debt with IOTXClear contract
        IIOTXClear(iotxClear).joinDebt(msg.sender, iotxsToRedeem);

        emit Redeemed(msg.sender, burned, tokenIdsToUnlock);
    }

    /**
     * @dev This function calculates uniIOTX amount based on IOTX amount for mint and burn operation after updating
     * rewards, aiming to keep the exchange ratio invariant to avoid user arbitrage.
     * Reference: https://github.com/RockX-SG/stake/blob/main/doc/uniETH_ETH2_0_Liquid_Staking_Explained.pdf
     */
    function _convertIotxToUniIOTX(uint amountIOTX) internal view returns (uint uniIOTXAmount) {
        uint totalSupply = IUniIOTX(uniIOTX).totalSupply();
        uint currentReserveAmt = _accountedReserve();
        uniIOTXAmount = DEFAULT_EXCHANGE_RATIO * amountIOTX;

        if (currentReserveAmt > 0) { // avert division overflow
            uniIOTXAmount = totalSupply * amountIOTX / currentReserveAmt;
        }
    }

    /**
     * @dev This function computes and returns the exchange ratio of uniIOTX to IOTX, multiplied by 1e18.
     */
    function _exchangeRatio() internal view returns (uint ratio) {
        uint uniIOTXAmount = IUniIOTX(uniIOTX).totalSupply();
        if (uniIOTXAmount == 0) {
            return DEFAULT_EXCHANGE_RATIO * MULTIPLIER;
        }
        ratio = _currentReserve() * MULTIPLIER / uniIOTXAmount;
    }

    /**
     * @dev This function computes and provides the current reserved IOTXs.
     */
    function _currentReserve() internal view returns(uint) {
        (uint pendingUserReward, ) = _calcPendingReward();
        return totalPending + totalStaked + pendingUserReward;
    }

    /**
     * @dev This function computes and provides the accounted reserved IOTXs.
     */
    function _accountedReserve() internal view returns(uint) {
        return totalPending + totalStaked;
    }

    /**
     * @dev This function distributes recently accrued rewards from delegates among users and the manager.
     * It also automatically reinvests the users' share into totalPending for upcoming staking activities.
     */
    function _updateReward() internal {
        uint reward = _syncReward();
        if ( reward > 0 ){
            uint userReward = _splitReward(reward);
            _compoundReward(userReward);
            emit RewardUpdated(reward);
        }
    }

    /**
     * @dev This function synchronizes the contract balance, calculating the increase in balance as the returned reward.
     */
    function _syncReward() internal returns (uint reward) {
        uint thisBalance = address(this).balance;
        uint diff = thisBalance - accountedBalance;
        accountedBalance = thisBalance;
        reward = diff;
    }

    /**
     * @dev This function divides the specified reward amount between users and the manager.
     */
    function _splitReward(uint reward) internal returns (uint userReward) {
        uint fee = reward * managerFeeShares / 1000;
        accountedManagerReward += fee;
        userReward = reward - fee;
        accountedUserReward += userReward;
    }

    /**
     * @dev This function incorporates the specified reward amount into the 'totalPending' value.
     */
    function _compoundReward(uint amount) internal {
        totalPending += amount;
    }

    /**
     * @dev This function calculates the pending reward that has not yet been synchronized and accounted for.
     */
    function _calcPendingReward() internal view returns(uint pendingUserReward, uint pendingManagerReward) {
        uint pendingReward = address(this).balance - accountedBalance;
        if (pendingReward > 0) {
            pendingManagerReward = pendingReward * managerFeeShares / 1000;
            return (pendingReward - pendingManagerReward, pendingManagerReward);
        }
        return (0, 0);
    }
}

