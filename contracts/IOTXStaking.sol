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
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "../interfaces/Roles.sol";
import "../interfaces/IIOTXClear.sol";
import "../interfaces/IUniIOTX.sol";
import "../interfaces/IIOTXStaking.sol";
import "../interfaces/ISystemStaking.sol";

contract IOTXStaking is IIOTXStaking, Initializable, PausableUpgradeable, AccessControlUpgradeable, ReentrancyGuardUpgradeable {
    // ---Use libraries---
    using SafeERC20 for IERC20;

    // ---External dependencies---
    address public systemStaking;
    address public uniIOTX;
    address public iotxClear;

    // ---Constants---
    uint public constant DEFAULT_EXCHANGE_RATIO = 1;
    uint public constant MULTIPLIER = 1e18;

    // ---State variables---

    address public globalDelegate;

    // The geometric sequence of the staking amount that can be staked onto the IoTeX network by our liquid staking protocol.
    // Every value populated in this sequence must be a valid bucket amount predefined by the IoTeX network.
    // Every value in this sequence corresponds to a unique indexed level in the range of [0, sequenceLength-1]
    uint public startAmount;
    uint public commonRatio;
    uint public sequenceLength;

    // Users can deposit any amount of IOTXs, but can only redeem amounts that are multiples of the redeemAmountBase.
    // This value is determined at contract initialization using the formula:
    // redeemAmountBase = startAmount * (commonRatio ** (sequenceLength-1))
    // Once set, it remains immutable.
    uint public redeemAmountBase;

    uint public stakeDuration;

    // Token queue map: The KEY corresponds to the bucket amount level defined as above;  the VALUE is a dynamic array of token IDs.
    mapping(uint => uint[]) public tokenQueues;

    // Redeemed token count. It can change only when redeeming operation performed.
    // Note: The staked token count at top staking level is: tokenQueues[sequenceLength-1].length - redeemedTokenCount
    uint public redeemedTokenCount;

    uint public accountedBalance;

    // Exchange ratio related variables
    // Track user deposits & redeem (uniIOTX mint & burn)
    uint public totalPending;               // Total pending IOTXs awaiting to be staked
    uint public totalStaked;            // Track current staked iotxs for delegates

    uint public managerFeeShares; // Shares range: [0, 1000]

    uint public accountedUserReward;           // Accounted shared user reward
    uint public accountedManagerReward;        // Accounted manager's reward

    // ---Events---
    event Minted(address user, uint minted);
    event Redeemed(address user, uint burned, uint[] tokenIds);
    event Staked(uint firstTokenId, uint amount, address delegate, uint count);
    event Merged(uint[] tokenIds, uint amount);
    event RewardUpdated(uint amount);
    event ManagerFeeSharesSet(uint shares);
    event ManagerFeeWithdrawed(uint amount, uint minted, address recipient);
    event GlobalDelegateSet(address delegate);
    event DelegatesUpdated(uint[] tokenIds, address delegate);

    // ---Modifiers---
    modifier onlyValidTransaction(uint deadline) {
        require(deadline > block.timestamp, "Transaction expired");
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

    /// @custom:oz-upgrades-unsafe-allow constructor
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
        _setupRole(ROLE_PAUSER, msg.sender);
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
            require(isActive, "inactive bucket type");
        }
    }

    /**
     * @dev This function sets manager's fee in range [0, 1000]
     */
    function setManagerFeeShares(uint shares) external onlyRole(DEFAULT_ADMIN_ROLE)  {
        require(shares <= 1000, "Manager fee shares out of range");
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
     * @return ratio The exchange ratio of uniIOTX to IOTX, multiplied by 'MULTIPLIER' (1e18).
     * @dev The factors that affect the returned result are the same as those of the 'currentReserve' function.
     */
    function exchangeRatio() external view returns (uint ratio) {
        return _exchangeRatio();
    }

    /**
     * @return The current reserve of IOTXs in our liquid staking protocol.
     * @dev The returned amount is determined by the following contributions:
     * 1. User deposits/stakes their principal.
     * 2. Rewards generated from delegation, excluding manager rewards, are added to the reserve.
     * 3. The manager fee is withdrawn and included in the totalPending amount.
     * 4. When users make a 'redeem' call, the totalStaked amount decreases.
     * 5. The Oracle service should regularly call the 'updateReward' function to account for new rewards.
     */
    function currentReserve() external view returns(uint) {
        return _currentReserve();
    }

    /**
     * @return An [i, j) slice of already redeemed/unlocked token id, which is indexed from 0 in this contract.
     * @param i, j The valid index values for i and j are determined by this conditional check: i < j && j <= redeemedTokenCount
     * @dev It recommended to check the value of 'redeemedTokenCount' beforehand to prevent the passed j from going out of range.
     */
    function getRedeemedTokenIds(uint i, uint j) external view returns (uint[] memory) {
        if (i < j && j <= redeemedTokenCount) {
            uint[] memory tq = tokenQueues[sequenceLength-1];
            uint[] memory tokenIds = new uint[](j-i);
            for (uint k = 0; k < j-i; k++) {
                tokenIds[k] = tq[i+k];
            }
            return tokenIds;
        }
        uint[] memory tokenIds = new uint[](0);
        return tokenIds;
    }

    /**
     * @return count The current staked token count of the specified token queue
     * @param tokenQueueIndex The token queue index falls within the range of [0, sequenceLength]
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
     * ======================================================================================
     *
     * EXTERNAL FUNCTIONS
     *
     * ======================================================================================
     */

    function setGlobalDelegate(address delegate) external whenNotPaused onlyRole(ROLE_ORACLE) {
        globalDelegate = delegate;

        emit GlobalDelegateSet(delegate);
    }

    function updateDelegates(uint[] calldata tokenIds, address delegate) external whenNotPaused onlyRole(ROLE_ORACLE) {
        ISystemStaking(systemStaking).changeDelegates(tokenIds, delegate);

        emit DelegatesUpdated(tokenIds, delegate);
    }

    /**
     * @dev This function keeps the exchange ratio invariant to avoid user arbitrage.
     */
    function deposit(uint deadline) external payable nonReentrant whenNotPaused onlyValidTransaction(deadline) returns (uint minted) {
        require(msg.value > 0, "Invalid deposit amount");

        minted = _mint();
        _stakeAtTopLevel();
        _stakeAndMergeAtSubLevel();
    }

    function stake() external whenNotPaused onlyRole(ROLE_ORACLE) {
        _stakeAtTopLevel();
        _stakeAndMergeAtSubLevel();
    }

    /**
     * @dev This function keeps the exchange ratio invariant to avoid user arbitrage.
     * @param iotxsToRedeem The number of IOTXs to redeem must be a multiple of the accepted amount of redeeming base.
     */
    function redeem(uint iotxsToRedeem, uint deadline) external nonReentrant onlyValidTransaction(deadline) returns(uint burned) {
        burned = _redeem(iotxsToRedeem);
    }

    /**
     * @dev This function distributes recently accrued rewards from delegates among users and the manager.
     * It also automatically reinvests the users' share into totalStaking for future stakes.
     */
    function updateReward() external onlyRole(ROLE_ORACLE) {
        uint reward = _syncReward();
        if ( reward > 0 ){
            uint userReward = _splitReward(reward);
            _compoundReward(userReward);
            emit RewardUpdated(reward);
        }
    }

    /**
     * @dev This function handles manager reward in this way:
     * 1. Mint uniIOTXs to the given recipient based on the given IOTX amount;
     * 2. Shift the corresponding amount of accountedManagerReward to totalPending.
     */
    function withdrawManagerFee(uint amount, address recipient) external nonReentrant onlyRole(ROLE_FEE_MANAGER)  {
        require(amount <= accountedManagerReward, "Insufficient accounted manager reward");

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

    function _mint() internal returns (uint minted) {
        accountedBalance += msg.value;

        uint toMint = _convertIotxToUniIOTX(msg.value);
        IUniIOTX(uniIOTX).mint(msg.sender, toMint);
        minted = toMint;

        totalPending += msg.value;

        emit Minted(msg.sender, minted);
    }

    function _stakeAtTopLevel() internal {
        // Determine values of stake params
        uint level = sequenceLength-1;
        (uint amount, uint count) = _getStakeAmountAndCount(level);
        if (count == 0) return;

        // Perform stake
        _doStake(level, amount, count);
    }

    function _stakeAndMergeAtSubLevel() internal {
        uint nextLevel = sequenceLength-2;
        while (totalPending >= startAmount) {
            nextLevel = _tryStake(nextLevel);
        }
    }

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

    function _getStakeAmountAndCount(uint level) internal view returns(uint amount, uint count) {
        amount = startAmount * (commonRatio**level);
        count = totalPending / amount;
    }

    function _merge(uint fromLevel) internal {
        for (uint i = fromLevel; i < sequenceLength-1; i++) {
            // Check merge condition
            uint[] storage tq = tokenQueues[i];
            if (tq.length < commonRatio) continue;

            // Call system merge service
            // All tokens will be merged into the first token in tokenIdsToMerge
            // Reference: https://github.com/iotexproject/iip13-contracts/blob/main/src/SystemStaking.sol#L302
            ISystemStaking(systemStaking).merge(tq, stakeDuration);

            // Move the merged tokens to upper queue
            uint[] storage tqUpper = tokenQueues[i+1];
            tqUpper.push(tq[0]);
            delete tokenQueues[i];

            emit Merged(tq, startAmount * (commonRatio**(i+1)));
        }
    }

    function _redeem(uint iotxsToRedeem) internal returns(uint burned) {
        // Check redeem condition
        require(iotxsToRedeem >= redeemAmountBase && iotxsToRedeem % redeemAmountBase == 0, "Invalid redeem amount");

        // Burn uniIOTXs
        uint toBurn = _convertIotxToUniIOTX(iotxsToRedeem);
        IERC20(uniIOTX).safeTransferFrom(msg.sender, address(this), toBurn);
        IUniIOTX(uniIOTX).burn(toBurn);
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
     * @dev This function calculates uniIOTX amount based on IOTX amount for mint and burn operation,
     * aiming to keep the exchange ratio invariant to avoid user arbitrage.
     * Reference: https://github.com/RockX-SG/stake/blob/main/doc/uniETH_ETH2_0_Liquid_Staking_Explained.pdf
     */
    function _convertIotxToUniIOTX(uint amountIOTX) internal view returns (uint uniIOTXAmount) {
        uint totalSupply = IUniIOTX(uniIOTX).totalSupply();
        uint currentReserveAmt = _currentReserve();
        uniIOTXAmount = DEFAULT_EXCHANGE_RATIO * amountIOTX;

        if (currentReserveAmt > 0) { // avert division overflow
            uniIOTXAmount = totalSupply * amountIOTX / currentReserveAmt;
        }
    }

    function _exchangeRatio() internal view returns (uint ratio) {
        uint uniIOTXAmount = IUniIOTX(uniIOTX).totalSupply();
        if (uniIOTXAmount == 0) {
            return DEFAULT_EXCHANGE_RATIO * MULTIPLIER;
        }
        ratio = _currentReserve() * MULTIPLIER / uniIOTXAmount;
    }

    function _currentReserve() internal view returns(uint) {
        return totalPending + totalStaked;
    }

    function _syncReward() internal returns (uint reward) {
        uint thisBalance = address(this).balance;
        uint diff = thisBalance - accountedBalance;
        accountedBalance = thisBalance;
        reward = diff;
    }

    function _splitReward(uint reward) internal returns (uint userReward) {
        uint fee = reward * managerFeeShares / 1000;
        accountedManagerReward += fee;
        userReward = reward - fee;
        accountedUserReward += userReward;
    }

    function _compoundReward(uint amount) internal {
        totalPending += amount;
    }
}
