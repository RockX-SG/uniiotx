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

    // State variables

    uint256 public accountedBalance; // TODO: Double check whether its value can be negative.

    uint256 public recentReceived;


    // Stake variables

    // Exchange ratio related variables
    // Track user deposits & redeem (uniIOTX mint & burn)
    uint256 public totalPending;               // Total pending IOTXs awaiting to be staked
    uint256 public totalStaked;            // Track current staked ethers for delegates

    uint256 public managerFeeShares; // Shares range: [0, 1000]

    // Track revenue from validators to form exchange ratio
    uint256 public accountedUserRevenue;           // Accounted shared user revenue
    uint256 public accountedManagerRevenue;        // Accounted manager's revenue
    uint256 public rewardDebts;

    uint256[] public immutable stakeAmountBases; // In sorted ascending order, i.e. [10000, 100000, 1000000] // Todo: optimize with less data provided
    uint256 public immutable stakeDuration;
    // Todo: Add mergeMultiple // Rate of increase

    // Token ids: Stake amount index => tokenIds
    mapping(uint256 => uint256[]) public stakedTokenIds;
    uint256[] public redeemedTokenIds;

    // Events
    event DelegateStopped(uint256 stoppedCount);
    event SystemStakingContractSet(address addr);
    event RedeemContractSet(address addr);
    event XIOTXContractSet(address addr);
    event ManagerFeeSharesSet(uint256 shares);
    event Minted(address user, uint256 minted);
    event Redeemed(address user, uint256 burned, uint256[] tokenIds);
    event Staked(uint256[] tokenIds, uint256 amount, address delegate);
    event Merged(uint256[] tokenIds, uint256 amount);
    event BalanceSynced(uint256 diff);
    event RevenueAccounted(uint256 amount);
    event ManagerFeeWithdrawed(uint256 amount, uint256 minted, address recipient);


// Errors
    error ZeroDelegates();
    error TransactionExpired(uint256 deadline, uint256 now);
    error InvalidRedeemAmount(uint256 redeemAmount, uint256 redeemBase);
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

    /**
     * @dev Initialization address
     */
    function initialize(
        uint256[] _stakeAmountSequence,
        uint256 _globalStakeDuration
    ) initializer public {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(PAUSER_ROLE, msg.sender);

        stakeAmountSequence = _stakeAmountSequence // Todo: Validate amount
        globalStakeDuration = _globalStakeDuration

//        // init default values
//        firstDebt = 1;
//        lastDebt = 0;
    }

    // TODO: do it in initialization?
    /**
     * @dev Set eth deposit contract address
     */
    function setSystemStakingContract(address _systemStake) external onlyRole(DEFAULT_ADMIN_ROLE) {
        systemStake = _systemStake;

        emit SystemStakingContractSet(_SystemStakingContract);
    }

    /**
     * @dev Set xIOTX token contract address
     */
    function setXIOTXContractAddress(address _xIOTXAddress) external onlyRole(DEFAULT_ADMIN_ROLE) {
        uniIOTX = _xIOTXAddress;

        emit XIOTXContractSet(_xIOTXAddress);
    }

    /**
     * @dev Set redeem contract
     */
    function setRedeemContract(address _redeemContract) external onlyRole(DEFAULT_ADMIN_ROLE) {
        iotxClear = _redeemContract;

        emit RedeemContractSet(_redeemContract);
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
     * @dev Returns current reserve of ethers
         */
    function currentReserve() public view returns(uint256) {
        return totalPending + totalStaked + accountedUserRevenue - rewardDebts; // Todo: Second thought of the correctness, including potential total debts.
    }

    /**
     * @dev Return the length of redeemedTokenIds
     */
    function getRedeemedTokensCount() external view returns (uint256) {
        return redeemedTokenIds.length;
    }

    /**
     * @dev Return the [i, j) slice of redeemedTokenIds
     */
    function getRedeemedTokenIds(uint256 i, uint256 j) external view returns (uint256[] memory tokenIds) {
        for (uint256 k := i; k < j; k++) {
            tokenIds.push(redeemedTokenIds[k])
        }
    }

    /**
     * ======================================================================================
     *
     * EXTERNAL FUNCTIONS
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
    functon updateReward() external onlyRole(ORACLE_ROLE) {
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


    /**
     * ======================================================================================
     *
     * PRIVATE FUNCTIONS
     *
     * ======================================================================================
     */

    // Todo: Consider whitelisting KYC users?
    function _mint(uint256 minToMint) private notZeroMint returns (uint256 minted) {
        accountedBalance += msg.value;

        toMint = _convertIotxToUniIotx(msg.value);
        if (toMint < minToMint) revert ExchangeRatioMismatch(minToMint, toMint);
        uniIOTX.mint(msg.sender, toMint);
        minted = toMint;

        totalPending += msg.value;

        emit Minted(msg.sender, minted);
    }

    function _stake() private returns (bool staked) {
        for (uint256 i = stakeAmountBases.length-1; i >= 0; i--) {
            base = stakeAmountBases[i];
            count = totalPending / base;
            amount = base * count;

            if (amount == 0) continue;

            startTokenId = systemStakeContract.stake{value:amount}(amount, stakeDuration, globalDelegate, count);
            uint256[] tokenIds = new uint256[](count);
            for (uint256 j = 0; j < count; j++) {
                uint256 tokenId = startTokenId+j;
                stakedTokenIds[i].push(tokenId)
                tokenIds[j] = tokenId;
            }
            staked = true;
            totalPending -= amount;
            totalStaked  += amount;
            accountedBalance -= amount;

            emit Staked(tokenIds, amount, globalDelegate);
        }
    }

    // Todo: Optimize performance?
    function _merge() private {
        for (uint256 i = 0; i < stakeAmountBases.length-1; i++) {
            targetAmount = stakeAmountBases[i+1];
            baseAmount = stakeAmountBases[i];
            accumulatedLen = stakedTokenIds[i].length

            uint256 accumulatedAmount;
            uint256 j;
            for (; j < accumulatedLen; j++) {
                accumulatedAmount += baseAmount;
                if (accumulatedAmount == targetAmount) break;
            }

            if (accumulatedAmount != targetAmount) continue;

            uint256[] memory tokenIdsToMerge = new uint256[](j+1);
            uint256[] memory tokenIdsUnmerge = new uint256[](accumulatedLen-j-1);
            for (k = 0; k < accumulatedLen; k++) { // Todo: Code reuse and remove loop.
                if (k <= j) {
                    tokenIdsToMerge[k] = stakedTokenIds[k];
                } else {
                    tokenIdsUnmerge[k] = stakedTokenIds[k]; // Todo: Assignment with incorrect index?
                }
            }

            // Reference: https://github.com/iotexproject/iip13-contracts/blob/main/src/SystemStaking.sol#L302
            systemStake.merge(tokenIdsToMerge, stakeDuration);
            stakedTokenIds[i] = tokenIdsUnmerge;
            stakedTokenIds[i+1].push(tokenIdsToMerge[0]);

            emit Merged(tokenIdsToMerge, targetAmount);
        }
    }

    // Todo: Optimization is needed
    function _redeem(uint256 iotxsToRedeem, uint256 maxToBurn) private returns(uint256 burned) {
        uint256 baseIndex = stakeAmountBases.legnth-1;
        uint256 base = stakeAmountBases[baseIndex];
        if (iotxsToRedeem % base != 0) revert InvalidRedeemAmount(iotxsToRedeem, base);

        // Burn uniIOTXs
        toBurn = _convertIotxToUniIotx(msg.value);
        if (toBurn > maxToBurn) revert ExchangeRatioMismatch(minToMint, toMint);
        uniIOTX.safeTransferFrom(msg.sender, address(this), toBurn); // Todo: Why transfer and burn, but not just burn?
        uniIOTX.burn(toBurn);
        burned = toBurn;
        totalStaked -= iotxsToRedeem; // Todo: double check whether we should decrease it here.

        // Unlock NFT(s)
        count = iotxsToRedeem / base;
        uint256[] stakedTokenIds = stakedTokenIds[baseIndex];
        uint256[] tokenIdsToUnlock = new uint256[](count);
        uint256[] tokenIdsReserve = new uint256[](stakedTokenIds.length-count);
        for (i = 0; i < stakedTokenIds.length; i++) { // Todo: Code reuse and remove loop.
        if (i <= j) {
            tokenIdsToUnlock[i] = stakedTokenIds[i];
            redeemedTokenIds.push(stakedTokenIds[i]);
        } else {
            tokenIdsReserve[i] = stakedTokenIds[i]; // Todo: Assignment with incorrect index
        }
        }
        systemStake.unlock(tokenIdsToUnlock);
        stakedTokenIds[baseIndex] = tokenIdsReserve;

        // Transfer NFT(s)
        for (i = 0; i < count; i++) {
            systemStake.safeTransferFrom(address(this), address(iotxClear), tokenIds[i]);
        }

        // Join debt
        iotxClear.joinDebt(msg.sender, iotxsToRedeem);

        emit Redeemed(msg.sender, burned, tokenIdsToUnlock);
    }

    /**
     * @dev Calculates uniIOTX amount based on IOTX amount for mint and burn operation,
     * aiming to keep the exchange ratio invariant to avoid user arbitrage.
     * Reference: https://github.com/RockX-SG/stake/blob/main/doc/uniETH_ETH2_0_Liquid_Staking_Explained.pdf
     */
    function _convertIotxToUniIotx(uint256 amountIOTX) private pure returns (uint256 amountUniIOTX) {
        uint256 totalSupply = uniIOTX.totalSupply();
        uint256 currentReserve = currentReserve();
        amountUniIOTX = defaultExchangeRatio * amountIOTX

        if (currentReserve > 0) { // avert division overflow
            amountUniIOTX = totalSupply * amountIOTX / currentReserve; // TODO: further consideration on the fractional part
        }
    }

    function _syncBalance() private returns (bool changed) {
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
    function _calculateRewards() private returns (uint256) {
        // Todo: Check whether to account in slash
        return recentReceived;
    }

    function _distributeRewards(uint256 rewards) private {
        uint256 fee = rewards * managerFeeShares / 1000;
        accountedManagerRevenue += fee;
        accountedUserRevenue += rewards - fee;

        emit RevenueAccounted(rewards);
    }

    // Todo: Reconsider the amount for auto compound
    function _autoCompound() private {
        uint256 amount = accountedUserRevenue - rewardDebts;
        totalPending += amount;
        rewardDebts += amount;
    }

}