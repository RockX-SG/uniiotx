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
    ISystemStake public systemStake;
    IUniIOTX public uniIOTX;
    IIOTXClear public iotxClear;

    // Constants
    uint public defaultExchangeRatio = 1;
    uint public constant MULTIPLIER = 1e18;

    // Type declarations

    // TopTokenQueue contains a dynamic-size queue, designed for managing tokens of the highest staking amount.
    // All tokens with less staking amount will be eventually merged into TopTokenQueue under proper conditions.
    // Besides, TopTokenQueue is the only container that provides claimable tokens.
    struct TopTokenQueue {
        // Designates where the next staked token should reside.
        // Its value increments continuously from 0.
        uint nextPushIndex;
        // This field stands for the start token queued that should be applied for the next redeem/unlock request.
        // Its value increments continuously from 0 and also equal to the total count of redeemed/unlocked tokens.
        // Tokens which are indexed less than this value can be applied for later unstake and withdraw requests.
        uint nextRedeemIndex;
        uint stakedCount;
        uint[] tokenIds;
    }

    // SubTokenQueue contains a fixed-size queue, designed for managing tokens with less staking amount that can be merged upward level by level.
    // Its underlying array is actually a ring queue and only holds less than commonRatio*2 tokens because of the merging operation.
   struct SubTokenQueue {
        // Designates where the next staked token should reside.
        // Its value increments continuously from 0 to (commonRatio*2 - 1) again and again.
        uint nextPushIndex;
        // This field stands for the start token queued that should be applied for the next merging operation.
        // Its value increments continuously from 0 to (commonRatio*2 - 1) again and again.
        uint nextMergeIndex;
        uint stakedCount;
        uint[] tokenIds;
    }

    // State variables

    address public globalDelegate;

    // The geometric sequence of the staking amount that can be staked onto the IoTeX network by our liquid staking protocol.
    uint public immutable startAmount;
    uint public immutable commonRatio;
    uint public immutable sequenceLength;

    uint public immutable stakeDuration;

    // Token containers
    TopTokenQueue public topTokenQueue; // implicitly sequenceIndex == sequenceLength-1
    mapping(uint => SubTokenQueue) public subTokenQueues; // sequenceIndex => SubTokenQueue

    uint public accountedBalance; // TODO: Double check whether its value can be negative.

    uint public recentReceived;


    // Stake variables

    // Exchange ratio related variables
    // Track user deposits & redeem (uniIOTX mint & burn)
    uint public totalPending;               // Total pending IOTXs awaiting to be staked
    uint public totalStaked;            // Track current staked iotxs for delegates

    uint public managerFeeShares; // Shares range: [0, 1000]

    // Track revenue from validators to form exchange ratio
    uint public accountedUserRevenue;           // Accounted shared user revenue
    uint public accountedManagerRevenue;        // Accounted manager's revenue
    uint public rewardDebts;

    // Events
    event DelegateStopped(uint stoppedCount);
    event ManagerFeeSharesSet(uint shares);
    event Minted(address user, uint minted);
    event Redeemed(address user, uint burned, uint[] tokenIds);
    event Staked(uint firstTokenId, uint amount, address delegate, uint count);
    event Merged(uint[] tokenIds, uint amount);
    event BalanceSynced(uint diff);
    event RevenueAccounted(uint amount);
    event ManagerFeeWithdrawed(uint amount, uint minted, address recipient);


// Errors
    error ZeroDelegates();
    error TransactionExpired(uint deadline, uint now);
    error InvalidRedeemAmount(uint expectedAmount, uint allowedAmount);
    error ExchangeRatioMismatch(uint expectedAmount, uint gotAmount);
    error ManagerFeeSharesOutOfRange();
    error InsufficientManagerRevenue(uint withdrawAmount, uint availableAmount);

    // Modifiers // TODO: code reuse across smart contracts?
    modifier onlyValidTransaction(uint deadline) {
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
        uint[] _stakeAmountSequence,
        uint _globalStakeDuration
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
    function setManagerFeeShares(uint shares) external onlyRole(DEFAULT_ADMIN_ROLE)  {
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

        ratio = currentReserve() * MULTIPLIER / uniIOTXAmount; // Todo: further consideration on the fractional part
    }

    /**
     * @dev Returns current reserve of IOTXs
     */
    function currentReserve() public view returns(uint) {
        return totalPending + totalStaked + accountedUserRevenue - rewardDebts; // Todo: Second thought of the correctness, including potential total debts.
    }

    /**
     * @dev Return [i, j) slice of already redeemed/unlocked token id.
     */
    function getRedeemedTokenIds(uint i, uint j) external view returns (uint[] memory tokenIds) {
        if (i < j && j <= topTokenQueue.nextRedeemIndex) {
            for (uint k := 0; k < j-i; k++) {
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

    function setGlobalDelegate(address delegate) external whenNotPaused onlyRole(ORACLE_ROLE) {
        globalDelegate = delegate;
    }

    function updateDelegates(uint[] tokenIds, address delegate) external pure whenNotPaused onlyRole(ORACLE_ROLE) {
        systemStake.changeDelegates(tokenIds, delegate);
    }

    // Todo: Give an explanation of param minToMint
    // Todo: Prove that
    /**
     * @notice This function keeps the exchange ratio invariant to avoid user arbitrage.
     */
    function deposit(uint minToMint, uint deadline) external payable nonReentrant whenNotPaused onlyValidTransaction(deadline) returns (uint minted) {
        minted = _mint(minToMint);
        uint fromAmountIndex = _stake();
        if (fromAmountIndex < sequenceLength-1) _merge(fromAmountIndex);
    }

    function stake() external whenNotPaused onlyRole(ORACLE_ROLE) {
        uint fromAmountIndex = _stake();
        if (fromAmountIndex < sequenceLength-1) _merge(fromAmountIndex);
    }

    // Todo: to be optimized
    // Todo: Give an explanation of param maxToBurn
    /**
     * @notice This function keeps the exchange ratio invariant to avoid user arbitrage.
     * @param iotxsToRedeem The number of IOTXs to redeem must be a multiple of the accepted amount of redeeming base.
     */
    function redeem(uint iotxsToRedeem, uint maxToBurn, uint deadline) external nonReentrant onlyValidTransaction(deadline) returns(uint burned) {
        burned = _redeem(iotxsToRedeem, maxToBurn);
    }

    // Todo: reconsideration on economic params.
    function updateReward() external onlyRole(ORACLE_ROLE) { // Todo: Disable onlyRole check?
        if (_syncBalance()) {
            uint rewards = _calculateRewards();
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
    function withdrawManagerFee(uint amount, address recipient) external nonReentrant onlyRole(MANAGER_ROLE)  {
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
    function _mint(uint minToMint) internal notZeroMint returns (uint minted) {
        accountedBalance += msg.value;

        toMint = _convertIotxToUniIotx(msg.value);
        if (toMint < minToMint) revert ExchangeRatioMismatch(minToMint, toMint);
        uniIOTX.mint(msg.sender, toMint);
        minted = toMint;

        totalPending += msg.value;

        emit Minted(msg.sender, minted);
    }

    function _stake() internal returns (uint fromAmountIndex) {
        for (fromAmountIndex = sequenceLength-1; fromAmountIndex >= 0 && totalPending >= startAmount; fromAmountIndex--) {
            // Determine stake amount
            uint amount = startAmount * (commonRatio**i);
            uint count = totalPending / amount;

            if (count == 0) continue

            uint totalAmount = amount * count;

            // Call system stake service
            firstTokenId = systemStake.stake{value:totalAmount}(amount, stakeDuration, globalDelegate, count);

            // Record minted & staked tokens
            // Todo: Recheck the implement very carefully.
            if (i == sequenceLenth-1) {
                TopTokenQueue storage tq = topTokenQueue;
                for (uint j = 0; j < count; j++) {
                    uint nextTokenId = firstTokenId+j;
                    tq.tokenIds[tq.nextPushIndex] = nextTokenId;
                    tq.nextPushIndex++
                }
                tq.stakedCount += count;
            } else {
                SubTokenQueue storage tq = subTokenQueues[i];
                for (uint j = 0; j < count; j++) {
                    uint nextTokenId = firstTokenId+j;
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

    function _merge(uint fromAmountIndex) internal {
        for (uint i = fromAmountIndex; i < sequenceLength-1; i++) {
            // Check merge condition
            SubTokenQueue storage tq = subTokenQueues[i];
            if (tq.stakedCount < commonRatio) continue;

            // Extract tokens to merge
            // Todo: Recheck the implement very carefully.
            uint[] memory tokenIdsToMerge = new uint[](commonRatio);
            uint counted;
            for (uint i = 0; i < commonRatio; ;) {
                tokenIdsToMerge[i] = tq.tokenIds[tq.nextMergeIndex];
                tq.nextMergeIndex =（tq.nextMergeIndex+1) % (commonRatio*2);
            }
            tq.stakedCount -= commonRatio;

            // Call system merge service
            // All tokens will be merged into the first token in tokenIdsToMerge
            // Reference: https://github.com/iotexproject/iip13-contracts/blob/main/src/SystemStaking.sol#L302
            systemStake.merge(tokenIdsToMerge, stakeDuration);

            // Record the merged token to upper queue
            if (i+1 == sequenceLengh-1) {
                TopTokenQueue storage tqUpper = topTokenQueue;
                tqUpper.tokenIds[tqUpper.nextPushIndex] = tokenIdsToMerge[0];
                tqUpper.nextPushIndex++;
                tqUpper.stakedCount++;
            } else {
                SubTokenQueue storage tqUpper = tokenQueues[i+1];
                tqUpper.tokenIds[tqUpper.nextPushIndex] = tokenIdsToMerge[0];
                tqUpper.nextPushIndex =（tqUpper.nextPushIndex+1) % (commonRatio*2);
                tqUpper.stakedCount++;
            }

            emit Merged(tokenIdsToMerge, targetAmount);
        }
    }

    // Todo: Recheck the implement very carefully.
    function _redeem(uint iotxsToRedeem, uint maxToBurn) internal returns(uint burned) {
        // Check redeem condition
        // Todo: This condition is very picky and maybe needs a more flexible policy.
        uint allowedAmount = startAmount * (commonRatio ** (sequenceLength-1)); // Todo: This value can be calculated at contract initialization?
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
        uint count = iotxsToRedeem / allowedAmount;
        uint[] tokenIdsToUnlock = new uint[](count);
        for (uint i = 0; i < count; i++) {
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
    function _convertIotxToUniIotx(uint amountIOTX) internal pure returns (uint amountUniIOTX) {
        uint totalSupply = uniIOTX.totalSupply();
        uint currentReserve = currentReserve();
        amountUniIOTX = defaultExchangeRatio * amountIOTX

        if (currentReserve > 0) { // avert division overflow
            amountUniIOTX = totalSupply * amountIOTX / currentReserve; // TODO: further consideration on the fractional part
        }
    }

    function _syncBalance() internal returns (bool changed) {
        uint thisBalance = address(this).balance
        if (thisBalance > accountedBalance) {
            uint diff = thisBalance - accountedBalance;
            accountedBalance = thisBalance;
            recentReceived += diff;
            changed = true;
            // Todo: Trigger vectorClockTick ?

            emit BalanceSynced(diff);
        }
    }

    // Todo: Reconsider the rules for calculating rewards.
    function _calculateRewards() internal returns (uint) {
        // Todo: Check whether to account in slash
        return recentReceived;
    }

    function _distributeRewards(uint rewards) internal {
        uint fee = rewards * managerFeeShares / 1000;
        accountedManagerRevenue += fee;
        accountedUserRevenue += rewards - fee;

        emit RevenueAccounted(rewards);
    }

    // Todo: Reconsider the amount for auto compound
    function _autoCompound() internal { // Todo: What if disabling this feature?
        uint amount = accountedUserRevenue - rewardDebts;
        totalPending += amount;
        rewardDebts += amount;
    }
}