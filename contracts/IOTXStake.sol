// SPDX-License-Identifier: UNLICENSED
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
    // External dependencies
    ISystemStake public systemStake;
    IUniIOTX public uniIOTX;
    IIOTXClear public iotxClear;

    // Constants
    uint public defaultExchangeRatio = 1;
    uint public constant MULTIPLIER = 1e18;

    // Type declarations

    // State variables

    address public globalDelegate;

    // The geometric sequence of the staking amount that can be staked onto the IoTeX network by our liquid staking protocol.
    // Every value populated in this sequence must be a valid bucket amount predefined by the IoTeX network.
    // Every value in this sequence corresponds to a unique indexed level in the range of [0, sequenceLength-1]
    uint public startAmount;
    uint public commonRatio;
    uint public sequenceLength;

    uint public stakeDuration;

    // Token queue map: The KEY corresponds to the bucket amount level defined as above;  the VALUE is a dynamic array of token IDs.
    mapping(uint => uint[]) public tokenQueues;

    // Redeemed token count. It can change only when redeeming operation performed.
    // Note: The staked token count at top staking level is: tokenQueues[sequenceLength-1].length - redeemedTokenCount
    uint public redeemedTokenCount;

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
    function pause() public onlyRole(ROLE_PAUSE) {
        _pause();
    }

    /**
     * @dev Unpause the contract
     */
    function unpause() public onlyRole(ROLE_PAUSE) {
        _unpause();
    }

    // Todo: Tune it properly, including initialize token container
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
        uint _sequenceLength
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
    }

    /**
     * @dev Set Manager's fee in range [0, 1000]
     */
    function setManagerFeeShares(uint shares) external onlyRole(ROLE_FEE_MANAGER)  {
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
        if (i < j && j <= redeemedTokenCount) {
            uint[] memory tq = tokenQueues[sequenceLength-1];
            for (uint k = 0; k < j-i; k++) {
                tokenIds[k] = tq[i+k];
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
    }

    function updateDelegates(uint[] calldata tokenIds, address delegate) external whenNotPaused onlyRole(ROLE_ORACLE) {
        systemStake.changeDelegates(tokenIds, delegate);
    }

    // Todo: Give an explanation of param minToMint
    // Todo: Prove that
    /**
     * @notice This function keeps the exchange ratio invariant to avoid user arbitrage.
     */
    function deposit(uint minToMint, uint deadline) external payable nonReentrant whenNotPaused onlyValidTransaction(deadline) returns (uint minted) {
        minted = _mint(minToMint);
        _stakeAtTopLevel();
        _stakeAndMergeAtSubLevel();
    }

    function stake() external whenNotPaused onlyRole(ROLE_ORACLE) {
        _stakeAtTopLevel();
        _stakeAndMergeAtSubLevel();
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
    function updateReward() external onlyRole(ROLE_ORACLE) { // Todo: Disable onlyRole check?
        if (_syncBalance()) {
            uint rewards = _calculateRewards();
            _distributeRewards(rewards);
            _autoCompound();
            recentReceived = 0;
        }
    }

    // Todo: Should we support directive withdrawing of IOTXs?
    /**
     * @dev This function handles manager revenue in this way:
     * 1. Mint uniIOTXs to the given recipient based on the given IOTX amount;
     * 2. Shift the corresponding amount of accountedManagerRevenue to totalPending.
     */
    function withdrawManagerFee(uint amount, address recipient) external nonReentrant onlyRole(ROLE_FEE_MANAGER)  {
        if (amount > accountedManagerRevenue) revert InsufficientManagerRevenue(amount, accountedManagerRevenue);

        uint toMint = _convertIotxTouniIOTX(amount);
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

        uint toMint = _convertIotxTouniIOTX(msg.value);
        if (toMint < minToMint) revert ExchangeRatioMismatch(minToMint, toMint);
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
        for (uint level = sequenceLength-2; level >= 0;  ) {
            // Determine values of stake params
            (uint amount, uint count) = _getStakeAmountAndCount(level);
            if (count == 0) continue;

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
            if (count == 0) level--;
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
            delete tokenQueues[i];
            uint[] storage tqUpper = tokenQueues[i+1];
            tqUpper.push(tq[0]);

            emit Merged(tq, startAmount * (commonRatio**(i+1)));
        }
    }

    // Todo: Recheck the implement very carefully.
    function _redeem(uint iotxsToRedeem, uint maxToBurn) internal returns(uint burned) {
        // Check redeem condition
        // Todo: This condition is very picky and maybe needs a more flexible policy.
        uint allowedAmount = startAmount * (commonRatio ** (sequenceLength-1)); // Todo: This value can be calculated at contract initialization?
        if (iotxsToRedeem < allowedAmount ||  iotxsToRedeem % allowedAmount != 0) revert InvalidRedeemAmount(iotxsToRedeem, allowedAmount);

        // Burn uniIOTXs
        uint toBurn = _convertIotxTouniIOTX(msg.value);
        if (toBurn > maxToBurn) revert ExchangeRatioMismatch(maxToBurn, toBurn);
        uniIOTX.burn(toBurn);
        burned = toBurn;
        totalStaked -= iotxsToRedeem;

        // Extract tokens to unlock
        uint [] memory tq = tokenQueues[sequenceLength-1];
        uint count = iotxsToRedeem / allowedAmount;
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
    function _convertIotxTouniIOTX(uint amountIOTX) internal view returns (uint amountuniIOTX) {
        uint totalSupply = uniIOTX.totalSupply();
        uint _currentReserve = currentReserve();
        amountuniIOTX = defaultExchangeRatio * amountIOTX;

        if (_currentReserve > 0) { // avert division overflow
            amountuniIOTX = totalSupply * amountIOTX / _currentReserve; // TODO: further consideration on the fractional part
        }
    }

    function _syncBalance() internal returns (bool changed) {
        uint thisBalance = address(this).balance;
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
    function _calculateRewards() internal view returns (uint) {
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