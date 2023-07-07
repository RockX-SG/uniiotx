pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";

import "interfaces/IIOTXClear.sol";
import "interfaces/IUniIOTX.sol";
import "../interfaces/ISystemStake.sol";
import "./DelegateManager.sol"

contract IOTXStake is Initializable, PausableUpgradeable, AccessControlUpgradeable, IERC721Receiver, DelegateManager {
    // Libraries
    using SafeERC20 for IERC20;

    // External dependencies
    ISystemStake public systemStakeContract;
    IUniIOTX public uniIOTXContract;
    IIOTXClear public iotxClearContract;

    // Constants
    uint256 public defaultExchangeRatio = 1;
    uint256 public constant MULTIPLIER = 1e18;

    // Type declarations

    // TopTokenQueue contains a dynamic-size queue, designed for managing tokens of the highest staking amount.
    // All tokens with less staking amount will be eventually merged into TopTokenQueue under proper conditions.
    // Besides, TopTokenQueue is the only container that provides claimable tokens.
    struct TopTokenQueue {
        // Designates where the next staked token should reside.
        // Its value increments continuously from 0.
        uint256 nextPushIndex;
        // This field stands for the start token queued that should be applied for the next redeem/unlock request.
        // Its value increments continuously from 0 and also equal to the total count of redeemed/unlocked tokens.
        // Tokens which are indexed less than this value can be applied for later unstake and withdraw requests.
        uint256 nextRedeemIndex;
        uint256 stakedCount;
        uint256[] tokenIds;
    }

    // SubTokenQueue contains a fixed-size queue, designed for managing tokens with less staking amount that can be merged upward level by level.
    // Its underlying array is actually a ring queue and only holds less than commonRatio*2 tokens because of the merging operation.
   struct SubTokenQueue {
        // Designates where the next staked token should reside.
        // Its value increments continuously from 0 to (commonRatio*2 - 1) again and again.
        uint256 nextPushIndex;
        // This field stands for the start token queued that should be applied for the next merging operation.
        // Its value increments continuously from 0 to (commonRatio*2 - 1) again and again.
        uint256 nextMergeIndex;
        uint256 stakedCount;
        uint256[] tokenIds;
    }

    // State variables

    // The geometric sequence of the staking amount that can be staked onto the IoTeX network by our liquid staking protocol.
    uint256 public immutable startAmount;
    uint256 public immutable commonRatio;
    uint256 public immutable sequenceLength;

    uint256 public immutable stakeDuration;

    // Token containers
    TopTokenQueue public topTokenQueue; // implicitly sequenceIndex == sequenceLength-1
    mapping(uint256 => SubTokenQueue) public subTokenQueues; // sequenceIndex => SubTokenQueue

    uint256 public accountedBalance; // TODO: Double check whether its value can be negative.

    uint256 public recentReceived;


    // Stake variables

    // Exchange ratio related variables
    // Track user deposits & redeem (uniIOTX mint & burn)
    uint256 public totalPending;               // Total pending IOTXs awaiting to be staked
    uint256 public totalStaked;            // Track current staked iotxs for delegates

    uint256 public managerFeeShares; // Shares range: [0, 1000]

    // Track revenue from validators to form exchange ratio
    uint256 public accountedUserRevenue;           // Accounted shared user revenue
    uint256 public accountedManagerRevenue;        // Accounted manager's revenue
    uint256 public rewardDebts;

    // Events
    event DelegateStopped(uint256 stoppedCount);
    event ManagerFeeSharesSet(uint256 shares);
    event Minted(address user, uint256 minted);
    event Redeemed(address user, uint256 burned, uint256[] tokenIds);
    event Staked(uint256 firstTokenId, uint256 amount, address delegate, uint256 count);
    event Merged(uint256[] tokenIds, uint256 amount);
    event BalanceSynced(uint256 diff);
    event RevenueAccounted(uint256 amount);
    event ManagerFeeWithdrawed(uint256 amount, uint256 minted, address recipient);


// Errors
    error ZeroDelegates();
    error TransactionExpired(uint256 deadline, uint256 now);
    error InvalidRedeemAmount(uint256 expectedAmount, uint256 allowedAmount);
    error ExchangeRatioMismatch(uint256 expectedAmount, uint256 gotAmount);
    error ManagerFeeSharesOutOfRange();
    error InsufficientManagerRevenue(uint256 withdrawAmount, uint256 availableAmount);

    // Modifiers // TODO: code reuse across smart contracts?
    modifier onlyValidTransaction(uint256 deadline) {
        if (deadline <= block.timestamp) revert TransactionExpired(deadline, block.timestamp);
        _;
    }

    modifier notZeroMint() {
        if (msg.value == 0) {
            return;
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

    // TODO: What if a user accidentally sent some IOTXs to this contract?
    // These IOTXs will be treated as reward. How could we prevent that?
    receive() external payable { }

    /**
     * @dev Pause the contract
     */
    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /**
     * @dev Unpause the contract
     */
    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    // Todo: Tune it properly, including initialize token container
    /**
     * @dev Initialization address
     */
    function initialize(
        address _systemStake,
        address _iotxClear,
        uint256[] _stakeAmountSequence,
        uint256 _globalStakeDuration
    ) initializer public {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(PAUSER_ROLE, msg.sender);

        systemStake = _systemStake;
        iotxClear = _iotxClear;

        stakeAmountSequence = _stakeAmountSequence // Todo: Validate amount
        globalStakeDuration = _globalStakeDuration
    }

    /**
     * @dev Set Manager's fee in range [0, 1000]
     */
    function setManagerFeeShares(uint256 shares) external onlyRole(DEFAULT_ADMIN_ROLE)  {
        if (shares > 1000) revert ManagerFeeSharesOutOfRange();
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
        uint256, // tokenId
        bytes calldata // data
    ) external pure override returns (bytes4) {
        return this.onERC721Received.selector;
    }

    /**
     * @dev Return exchange ratio of uniIOTX to IOTX, multiplied by 1e18
     */
    function exchangeRatio() external view returns (uint256 ratio) {
        uint256 uniIOTXAmount = uniIOTX.totalSupply();
        if (uniIOTXAmount == 0) {
            return 1 * MULTIPLIER;
        }

        ratio = currentReserve() * MULTIPLIER / uniIOTXAmount; // Todo: further consideration on the fractional part
    }

    /**
     * @dev Returns current reserve of IOTXs
     */
    function currentReserve() public view returns(uint256) {
        return totalPending + totalStaked + accountedUserRevenue - rewardDebts; // Todo: Second thought of the correctness, including potential total debts.
    }

    /**
     * @dev This is a handy function to return the total redeemed token ids count.
     * All already redeemed/unlocked token ids are indexed less than this count and can be applied for later unstake and withdraw requests.
     * The token id index equal to this count stands for the start token that should be applied for the next redeem/unlock request.
     */
    function getTotalRedeemedTokenIdsCount() external view returns (uint256) {
        return topTokenQueue.nextRedeemIndex;
    }

    /**
     * @dev Return [i, j) slice of already redeemed/unlocked token id.
     * These returned token ids can be applied for later unstake and withdraw requests.
     */
    function getRedeemedTokenIds(uint256 i, uint256 j) external view returns (uint256[] memory tokenIds) {
        if (i < j && j <= topTokenQueue.nextRedeemIndex) {
            for (uint256 k := 0; k < j-i; k++) {
                tokenIds[k] = topTokenQueue.tokenIds[i+k];
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

    // Todo: Give an explanation of param minToMint
    // Todo: Prove that
    /**
     * @notice This function keeps the exchange ratio invariant to avoid user arbitrage.
     */
    function deposit(uint256 minToMint, uint256 deadline) external payable nonReentrant whenNotPaused onlyValidTransaction(deadline) returns (uint256 minted) {
        minted = _mint(minToMint);
        if (_stake()) _merge();
    }

    function stake() external whenNotPaused onlyRole(ORACLE_ROLE) {
        if (_stake()) _merge();
    }

    // Todo: to be optimized
    // Todo: Give an explanation of param maxToBurn
    /**
     * @notice This function keeps the exchange ratio invariant to avoid user arbitrage.
     * @param iotxsToRedeem The number of IOTXs to redeem must be a multiple of the accepted amount of redeeming base.
     */
    function redeem(uint256 iotxsToRedeem, uint256 maxToBurn, uint256 deadline) external nonReentrant onlyValidTransaction(deadline) returns(uint256 burned) {
        burned = _redeem(iotxsToRedeem, maxToBurn);
    }

    // Todo: reconsideration on economic params.
    functon updateReward() external onlyRole(ORACLE_ROLE) { // Todo: Disable onlyRole check?
        if (_syncBalance()) {
            uint256 rewards = _calculateRewards();
            _distributeRewards(rewards);
            _autoCompoud();
            recentReceived = 0;
        }
    }

    // Todo: Should we support directive withdrawing of IOTXs?
    /**
     * @dev This function handles manager revenue in this way:
     * 1. Mint uniIOTXs to the given recipient based on the given IOTX amount;
     * 2. Shift the corresponding amount of accountedManagerRevenue to totalPending.
     */
    function withdrawManagerFee(uint256 amount, address recipient) external nonReentrant onlyRole(MANAGER_ROLE)  {
        if (amount > accountedManagerRevenue) revert InsufficientManagerRevenue();

        toMint = _convertIotxToUniIotx(amount);
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

    // Todo: Consider whitelisting KYC users?
    function _mint(uint256 minToMint) internal notZeroMint returns (uint256 minted) {
        accountedBalance += msg.value;

        toMint = _convertIotxToUniIotx(msg.value);
        if (toMint < minToMint) revert ExchangeRatioMismatch(minToMint, toMint);
        uniIOTX.mint(msg.sender, toMint);
        minted = toMint;

        totalPending += msg.value;

        emit Minted(msg.sender, minted);
    }

    function _stake() internal returns (bool staked) {
        for (uint256 i = sequenceLength-1; i >= 0 && totalPending >= startAmount; i--) {
            // Determine stake amount
            uint256 amount = startAmount * (commonRatio**i);
            uint256 count = totalPending / amount;

            if (count == 0) continue

            uint256 totalAmount = amount * count;

            // Call system stake service
            firstTokenId = systemStakeContract.stake{value:totalAmount}(amount, stakeDuration, globalDelegate, count);

            // Record minted & staked tokens
            // Todo: Recheck the implement very carefully.
            if (i == sequenceLenth-1) {
                TopTokenQueue storage tq = topTokenQueue;
                for (uint256 j = 0; j < count; j++) {
                    uint256 nextTokenId = firstTokenId+j;
                    tq.tokenIds[tq.nextPushIndex] = nextTokenId;
                    tq.nextPushIndex++
                }
                tq.stakedCount += count;
            } else {
                SubTokenQueue storage tq = subTokenQueues[i];
                for (uint256 j = 0; j < count; j++) {
                    uint256 nextTokenId = firstTokenId+j;
                    tq.tokenIds[tq.nextPushIndex] = nextTokenId;
                    tq.nextPushIndex =（tq.nextPushIndex+1) % (commonRatio*2);
                }
                tq.stakedCount += count;
            }

            // Update fund status
            totalPending -= totalAmount;
            totalStaked  += totalAmount;
            accountedBalance -= totalAmount;

            staked = true;

            emit Staked(firstTokenId, amount, globalDelegate, count);
        }
    }

    function _merge() internal {
        for (uint256 i = 0; i < sequenceLength-1; i++) {
            // Check merge condition
            SubTokenQueue storage tq = subTokenQueues[i];
            if (tq.stakedCount < commonRatio) break;

            // Extract tokens to merge
            // Todo: Recheck the implement very carefully.
            uint256[] memory tokenIdsToMerge = new uint256[](commonRatio);
            uint256 counted;
            for (uint256 i = 0; i < commonRatio; ;) {
                tokenIdsToMerge[i] = tq.tokenIds[tq.nextMergeIndex];
                tq.nextMergeIndex =（tq.nextMergeIndex+1) % (commonRatio*2);
            }
            tq.stakedCount -= commonRatio;

            // Call system merge service
            // All tokens will be merged into the first token in tokenIdsToMerge
            // Reference: https://github.com/iotexproject/iip13-contracts/blob/main/src/SystemStaking.sol#L302
            systemStake.merge(tokenIdsToMerge, stakeDuration);

            // Record the merged token to upper queue
            SubTokenQueue storage tqUpper = tokenQueues[i+1];
            tqUpper.tokenIds[tqUpper.nextPushIndex] = tokenIdsToMerge[0];
            tqUpper.nextPushIndex =（tqUpper.nextPushIndex+1) % (commonRatio*2);
            tqUpper.stakedCount++;

            emit Merged(tokenIdsToMerge, targetAmount);
        }
    }

    // Todo: Recheck the implement very carefully.
    function _redeem(uint256 iotxsToRedeem, uint256 maxToBurn) internal returns(uint256 burned) {
        // Check redeem condition
        // Todo: This condition is very picky and maybe needs a more flexible policy.
        uint256 allowedAmount = startAmount * (commonRatio ** (sequenceLength-1)); // Todo: This value can be calculated at contract initialization?
        if (iotxsToRedeem < allowedAmount ||  iotxsToRedeem % allowedAmount != 0) revert InvalidRedeemAmount(iotxsToRedeem, allowedAmount);

        // Burn uniIOTXs
        toBurn = _convertIotxToUniIotx(msg.value);
        if (toBurn > maxToBurn) revert ExchangeRatioMismatch(minToMint, toMint);
        uniIOTX.safeTransferFrom(msg.sender, address(this), toBurn); // Todo: Why transfer and burn, but not just burn?
        uniIOTX.burn(toBurn);
        burned = toBurn;
        totalStaked -= iotxsToRedeem;

        // Extract tokens to unlock
        TopTokenQueue storage tq = topTokenQueue;
        uint256 count = iotxsToRedeem / allowedAmount;
        uint256[] tokenIdsToUnlock = new uint256[](count);
        for (uint256 i = 0; i < count; i++) {
            tokenIdsToUnlock[i] = tq.tokenIds[tq.nextRedeemIndex];
            tq.nextRedeemIndex++;
        }
        tq.stakedCount -= count;

        // Call system unlock service
        systemStake.unlock(tokenIdsToUnlock);

        // Transfer unlocked tokens to IOTXClear contract
        for (i = 0; i < count; i++) {
            systemStake.safeTransferFrom(address(this), address(iotxClear), tokenIds[i]);
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
    function _convertIotxToUniIotx(uint256 amountIOTX) internal pure returns (uint256 amountUniIOTX) {
        uint256 totalSupply = uniIOTX.totalSupply();
        uint256 currentReserve = currentReserve();
        amountUniIOTX = defaultExchangeRatio * amountIOTX

        if (currentReserve > 0) { // avert division overflow
            amountUniIOTX = totalSupply * amountIOTX / currentReserve; // TODO: further consideration on the fractional part
        }
    }

    function _syncBalance() internal returns (bool changed) {
        uint256 thisBalance = address(this).balance
        if (thisBalance > accountedBalance) {
            uint256 diff = thisBalance - accountedBalance;
            accountedBalance = thisBalance;
            recentReceived += diff;
            changed = true;
            // Todo: Trigger vectorClockTick ?

            emit BalanceSynced(diff);
        }
    }

    // Todo: Reconsider the rules for calculating rewards.
    function _calculateRewards() internal returns (uint256) {
        // Todo: Check whether to account in slash
        return recentReceived;
    }

    function _distributeRewards(uint256 rewards) internal {
        uint256 fee = rewards * managerFeeShares / 1000;
        accountedManagerRevenue += fee;
        accountedUserRevenue += rewards - fee;

        emit RevenueAccounted(rewards);
    }

    // Todo: Reconsider the amount for auto compound
    function _autoCompound() internal { // Todo: What if disabling this feature?
        uint256 amount = accountedUserRevenue - rewardDebts;
        totalPending += amount;
        rewardDebts += amount;
    }

}