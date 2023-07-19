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
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "./Roles.sol";
import "../interfaces/IIOTXClear.sol";
import "../interfaces/IUniIOTX.sol";
import "../interfaces/ISystemStake.sol";

contract IOTXStake is Initializable, PausableUpgradeable, AccessControlUpgradeable, IERC721Receiver, ReentrancyGuardUpgradeable {
    // ---External dependencies---
    ISystemStake public systemStake;
    IUniIOTX public uniIOTX;
    IIOTXClear public iotxClear;

    // ---Constants---
    uint public defaultExchangeRatio = 1;
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

    uint public recentReceived;

    // Exchange ratio related variables
    // Track user deposits & redeem (uniIOTX mint & burn)
    uint public totalPending;               // Total pending IOTXs awaiting to be staked
    uint public totalStaked;            // Track current staked iotxs for delegates

    uint public managerFeeShares; // Shares range: [0, 1000]

    // Track revenue from validators to form exchange ratio
    uint public accountedUserRevenue;           // Accounted shared user revenue
    uint public accountedManagerRevenue;        // Accounted manager's revenue
    uint public rewardDebts;

    // ---Events---
    event DelegateStopped(uint stoppedCount);
    event ManagerFeeSharesSet(uint shares);
    event Minted(address user, uint minted);
    event Redeemed(address user, uint burned, uint[] tokenIds);
    event Staked(uint firstTokenId, uint amount, address delegate, uint count);
    event Merged(uint[] tokenIds, uint amount);
    event BalanceSynced(uint diff);
    event RevenueAccounted(uint amount);
    event ManagerFeeWithdrawed(uint amount, uint minted, address recipient);

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
     * @notice This function is exclusively designed to receive staking rewards.
     * The 'deposit' function should be invoked whenever users wish to stake IOTXs.
     * Any IOTXs accidentally sent to this contract will be considered as rewards.
     */
    receive() external payable { }

    /**
     * @dev Pause the contract
     */
    function pause() public onlyRole(ROLE_PAUSE) {
        _pause();
    }

    /**
     * @dev Unpause the contract
     */
    function unpause() public onlyRole(ROLE_PAUSE) {
        _unpause();
    }

    /**
     * @dev Initialization address
     */
    function initialize(
        address _systemStakeAddress,
        address _uniIOTX,
        address _iotxClear,
        address _oracleAddress,
        uint _startAmount,
        uint _commonRatio,
        uint _sequenceLength,
        uint _stakeDuration
    ) public initializer {
        // Roles
        _grantRole(ROLE_FEE_MANAGER, msg.sender);
        _grantRole(ROLE_PAUSE, msg.sender);
        _grantRole(ROLE_ORACLE, _oracleAddress);

        // Collaborative contracts
        systemStake = ISystemStake(_systemStakeAddress);
        uniIOTX = IUniIOTX(_uniIOTX);
        iotxClear = IIOTXClear(_iotxClear);

        // Immutable staking variables
        startAmount = _startAmount;
        commonRatio = _commonRatio;
        sequenceLength = _sequenceLength;
        stakeDuration = _stakeDuration;
        redeemAmountBase = startAmount * (commonRatio ** (sequenceLength-1));
    }

    /**
     * @dev Set Manager's fee in range [0, 1000]
     */
    function setManagerFeeShares(uint shares) external onlyRole(ROLE_FEE_MANAGER)  {
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
     * @dev IERC721Receiver implement for receiving staking NFT
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
     * @dev Return exchange ratio of uniIOTX to IOTX, multiplied by 1e18
     */
    function exchangeRatio() external view returns (uint ratio) {
        uint uniIOTXAmount = uniIOTX.totalSupply();
        if (uniIOTXAmount == 0) {
            return 1 * MULTIPLIER;
        }

        ratio = currentReserve() * MULTIPLIER / uniIOTXAmount;
    }

    /**
     * @dev Returns current reserve of IOTXs
     */
    function currentReserve() public view returns(uint) {
        return totalPending + totalStaked + accountedUserRevenue - rewardDebts;
    }

    /**
     * @dev Return [i, j) slice of already redeemed/unlocked token id.
     */
    function getRedeemedTokenIds(uint i, uint j) external view returns (uint[] memory tokenIds) {
        if (i < j && j <= redeemedTokenCount) {
            uint[] memory tq = tokenQueues[sequenceLength-1];
            for (uint k = 0; k < j-i; k++) {
                tokenIds[k] = tq[i+k];
            }
        }
    }

    /**
     * @dev Return current staked token count at given staking level
     */
    function getStakedTokenCount(uint stakingLevel) external view returns (uint count) {
        require(stakingLevel < sequenceLength, "Staking level out of range");
        count = tokenQueues[stakingLevel].length;
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
    }

    function updateDelegates(uint[] calldata tokenIds, address delegate) external whenNotPaused onlyRole(ROLE_ORACLE) {
        systemStake.changeDelegates(tokenIds, delegate);
    }

    /**
     * @notice This function keeps the exchange ratio invariant to avoid user arbitrage.
     */
    function deposit(uint minToMint, uint deadline) external payable nonReentrant whenNotPaused onlyValidTransaction(deadline) returns (uint minted) {
        require(msg.value > 0, "Invalid deposit amount");

        minted = _mint(minToMint);
        _stakeAtTopLevel();
        _stakeAndMergeAtSubLevel();
    }

    function stake() external whenNotPaused onlyRole(ROLE_ORACLE) {
        _stakeAtTopLevel();
        _stakeAndMergeAtSubLevel();
    }

    /**
     * @notice This function keeps the exchange ratio invariant to avoid user arbitrage.
     * @param iotxsToRedeem The number of IOTXs to redeem must be a multiple of the accepted amount of redeeming base.
     */
    function redeem(uint iotxsToRedeem, uint maxToBurn, uint deadline) external nonReentrant onlyValidTransaction(deadline) returns(uint burned) {
        burned = _redeem(iotxsToRedeem, maxToBurn);
    }

    function updateReward() external onlyRole(ROLE_ORACLE) {
        if (_syncBalance()) {
            uint rewards = _calculateRewards();
            _distributeRewards(rewards);
            _autoCompound();
            recentReceived = 0;
        }
    }

    /**
     * @dev This function handles manager revenue in this way:
     * 1. Mint uniIOTXs to the given recipient based on the given IOTX amount;
     * 2. Shift the corresponding amount of accountedManagerRevenue to totalPending.
     */
    function withdrawManagerFee(uint amount, address recipient) external nonReentrant onlyRole(ROLE_FEE_MANAGER)  {
        require(amount <= accountedManagerRevenue, "Insufficient manager revenue");

        uint toMint = _convertIotxToUniIOTX(amount);
        uniIOTX.mint(recipient, toMint);

        accountedManagerRevenue -= amount;
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

    function _mint(uint minToMint) internal returns (uint minted) {
        accountedBalance += msg.value;

        uint toMint = _convertIotxToUniIOTX(msg.value);
        require(toMint >= minToMint, "Exchange ratio mismatch");
        uniIOTX.mint(msg.sender, toMint);
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
        for (uint level = sequenceLength-2; level >= 0 && level < sequenceLength-1;  ) {
            // Determine values of stake params
            (uint amount, uint count) = _getStakeAmountAndCount(level);
            if (count == 0) {
                if (level == 0) return;
                level--;
                continue;
            }

            uint stakedCount = tokenQueues[level].length;
            if ((count+stakedCount) >= commonRatio) {
                count = commonRatio-stakedCount;
            }

            // Perform stake
            _doStake(level, amount, count);

            // Merge tokens if possible
            if (count+stakedCount == commonRatio) _merge(level);

            // Handle remained amount
            (, count) = _getStakeAmountAndCount(level);
            if (count == 0) {
                if (level == 0) return;
                level--;
            }
        }
    }

    function _doStake(uint level, uint amount, uint count) internal {
        // Calculate total amount
        uint totalAmount = amount * count;

        // Call system stake service
        uint firstTokenId = systemStake.stake{value:totalAmount}(amount, stakeDuration, globalDelegate, count);

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
            systemStake.merge(tq, stakeDuration);

            // Move the merged tokens to upper queue
            uint[] storage tqUpper = tokenQueues[i+1];
            tqUpper.push(tq[0]);
            delete tokenQueues[i];

            emit Merged(tq, startAmount * (commonRatio**(i+1)));
        }
    }

    function _redeem(uint iotxsToRedeem, uint maxToBurn) internal returns(uint burned) {
        // Check redeem condition
        require(iotxsToRedeem >= redeemAmountBase && iotxsToRedeem % redeemAmountBase == 0, "Invalid redeem amount");

        // Burn uniIOTXs
        uint toBurn = _convertIotxToUniIOTX(msg.value);
        require(toBurn <= maxToBurn, "Exchange ratio mismatch");
        uniIOTX.burn(toBurn);
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
        systemStake.unlock(tokenIdsToUnlock);

        // Transfer unlocked tokens to IOTXClear contract
        for (uint i = 0; i < count; i++) {
            systemStake.safeTransferFrom(address(this), address(iotxClear), tokenIdsToUnlock[i]);
        }

        // Record corresponding amount of debt with IOTXClear contract
        iotxClear.joinDebt(msg.sender, iotxsToRedeem);

        emit Redeemed(msg.sender, burned, tokenIdsToUnlock);
    }

    /**
     * @dev Calculates uniIOTX amount based on IOTX amount for mint and burn operation,
     * aiming to keep the exchange ratio invariant to avoid user arbitrage.
     * Reference: https://github.com/RockX-SG/stake/blob/main/doc/uniETH_ETH2_0_Liquid_Staking_Explained.pdf
     */
    function _convertIotxToUniIOTX(uint amountIOTX) internal view returns (uint amountuniIOTX) {
        uint totalSupply = uniIOTX.totalSupply();
        uint _currentReserve = currentReserve();
        amountuniIOTX = defaultExchangeRatio * amountIOTX;

        if (_currentReserve > 0) { // avert division overflow
            amountuniIOTX = totalSupply * amountIOTX / _currentReserve;
        }
    }

    function _syncBalance() internal returns (bool changed) {
        uint thisBalance = address(this).balance;
        if (thisBalance > accountedBalance) {
            uint diff = thisBalance - accountedBalance;
            accountedBalance = thisBalance;
            recentReceived += diff;
            changed = true;

            emit BalanceSynced(diff);
        }
    }

    function _calculateRewards() internal view returns (uint) {
        return recentReceived;
    }

    function _distributeRewards(uint rewards) internal {
        uint fee = rewards * managerFeeShares / 1000;
        accountedManagerRevenue += fee;
        accountedUserRevenue += rewards - fee;

        emit RevenueAccounted(rewards);
    }

    function _autoCompound() internal {
        uint amount = accountedUserRevenue - rewardDebts;
        totalPending += amount;
        rewardDebts += amount;
    }
}

