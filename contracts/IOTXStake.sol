pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";

import "interfaces/IIOTXClear.sol";
import "interfaces/IUniIOTX.sol";
import "../interfaces/ISystemStake.sol";
import "./DelegateManager.sol"

contract IOTXStake is Initializable, PausableUpgradeable, AccessControlUpgradeable, IERC721Receiver, DelegateManager {
    // External dependencies
    ISystemStake public systemStakeContract;
    IUniIOTX public uniIOTXContract;
    IIOTXClear public iotxClearContract;

// Todo: code cleaning
//    enum stakeState (
//        LOCKED,
//        UNLOCKED,
//    )
//
//    struct tokenIdContainer {
//
//        stakeState
//    }

//    struct tokenIds {
//        uint256 size;
//        mapping(uint256 => uint256) data;   // startId => count
//    }

    // Constants
    uint256 public defaultExchangeRatio = 1;
    uint256 private constant MULTIPLIER = 1e18;

    // State variables

    uint256 public accountedBalance; // TODO: Double check whether its value can be negative.


    // Stake variables

    // exchange ratio related variables
    // track user deposits & redeem (uniIOTX mint & burn)
    uint256 private totalPending;               // total pending IOTXs awaiting to be staked
    uint256 private totalStaked;            // track current staked ethers for delegates
    uint256 private totalDebts;             // track current unpaid debts // TODO: ignore it?

    // track revenue from validators to form exchange ratio
    uint256 private accountedUserRevenue;           // accounted shared user revenue
    uint256 private accountedManagerRevenue;        // accounted manager's revenue
    uint256 private rewardDebts;                    // check validatorStopped function // TODO: modify this comment

    uint256[] public immutable stakeAmountBases; // In sorted ascending order, i.e. [10000, 100000, 1000000]
    uint256 public immutable stakeDuration;

    // Token ids: stake amount index => tokenIds
    mapping(uint256 => uint256[]) public stakedTokenIds;
    uint256[] public redeemedTokenIds;
//    mapping(BucketType => tokenIds) public redeemedTokenIds; // Todo: Code cleaning
//    mapping(BucketType => tokenIds) public withdrawedTokenIds;

    // Events
    event DelegateStopped(uint256 stoppedCount);
    event SystemStakingContractSet(address addr);
    event RedeemContractSet(address addr);
    event XIOTXContractSet(address addr);
    event Minted(address user, uint256 minToMint, uint256 minted);
    event Staked(uint256[] tokenIds, uint256 amount, address delegate);
    event Merged(uint256[] tokenIds, uint256 amount);

    // Errors
    error ZeroDelegates();
    error TransactionExpired(uint256 deadline, uint256 now);
    error InvalidRedeemAmount(uint256 redeemAmount, uint256 redeemBase);
    error ExchangeRatioMismatch(uint256 expectedAmount, uint256 gotAmount);

    // Modifiers // TODO: code reuse across smart contracts
    modifier onlyValidTransaction(uint256 deadline) {
        if (deadline <= block.timestamp) revert TransactionExpired(deadline, block.timestamp);
        _;
    }

    modifer notZeroMint() {
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
     * @dev pause the contract
     */
    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /**
     * @dev unpause the contract
     */
    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /**
     * @dev initialization address
     */
    function initialize(
        uint256[] _stakeAmountSequence,
        uint256 _globalStakeDuration
    ) initializer public {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(PAUSER_ROLE, msg.sender);

        stakeAmountSequence = _stakeAmountSequence
        globalStakeDuration = _globalStakeDuration

//        // init default values
//        firstDebt = 1;
//        lastDebt = 0;
    }

    // TODO: do it in initialization
    /**
     * @dev set eth deposit contract address
     */
    function setSystemStakingContract(address _systemStake) external onlyRole(DEFAULT_ADMIN_ROLE) {
        systemStake = _systemStake;

        emit SystemStakingContractSet(_SystemStakingContract);
    }

    /**
     * @dev set xIOTX token contract address
     */
    function setXIOTXContractAddress(address _xIOTXAddress) external onlyRole(DEFAULT_ADMIN_ROLE) {
        uniIOTX = _xIOTXAddress;

        emit XIOTXContractSet(_xIOTXAddress);
    }

    /**
     * @dev set redeem contract
     */
    function setRedeemContract(address _redeemContract) external onlyRole(DEFAULT_ADMIN_ROLE) {
        iotxClear = _redeemContract;

        emit RedeemContractSet(_redeemContract);
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

        ratio = currentReserve() * MULTIPLIER / uniIOTXAmount; // TODO: further consideration on the fractional part
    }

    /**
     * @dev returns current reserve of ethers
         */
    function currentReserve() public view returns(uint256) {
        return totalPending + totalStaked + accountedUserRevenue - totalDebts - rewardDebts;  // TODO: ignore total debts?
    }

    /**
     * @dev return pending IOTXs
     */
    function getPendingIOTXs() external view returns (uint256) { return totalPending; }

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

    function deposit(uint256 minToMint, uint256 deadline) external payable nonReentrant whenNotPaused onlyValidTransaction(deadline) returns (uint256 minted) {
        minted = _mint(minToMint);
        if (_stake()) _merge();
    }

    function stake() external whenNotPaused onlyRole(ORACLE_ROLE) {
        if (_stake()) _merge();
    }

    // TODO: to be modified
    function redeem(uint256 amount, uint256 maxToBurn, uint256 deadline) external onlyValidTransaction(deadline) nonReentrant returns(uint256 burned) {
        if (amount % stakeAmount03 != 0) revert InvalidRedeemAmount(amount, stakeAmount03);
        // TODO: verify the rights/preconditiions of msg.sender to redeem ? including:
        // TODO: check amount ceiling according to clients deposit amount.

        uint256 totalXIOTX = uniIOTX.totalSupply();
        uint256 xIOTXToBurn = totalXIOTX * amount / currentReserve(); // TODO:
        require(xIOTXToBurn <= maxToBurn, "USR004");

        // NOTE: the following procdure must keep exchangeRatio invariant:
        // transfer uniIOTX from sender & burn
        // TODO: why transfer and burn, but not only burn?
        uniIOTX.safeTransferFrom(msg.sender, address(this), xIOTXToBurn);
        uniIOTX.burn(xIOTXToBurn);

        // Unlock NFT(s)
        count = amount / stakeAmount03;
        tokenIds = _removestakedTokenIds(BucketType.BucketType3, count);
        systemStake.unlock(tokenIds);
        _addredeemedTokenIds(tokenIds);

        // Transfer NFT owner
        for (i = 0; i < tokenIds.length; i++) {
            systemStake.transferFrom(address(this), address(iotxClear), tokenIds[i]);
        }

        // Join debt
        iotxClear.joinDebt(msg.sender, amount);

        // return burned
        return xIOTXToBurn;
    }

//    function updateDelegates(uint256[] tokenIds, address delegate) external whenNotPaused onlyRole(ORACLE_ROLE) {
//        systemStake.changeDelegates(tokenIds, delegate);
//    }

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

    // TODO: Consider whitelisting KYC users?
    // TODO: Give an explanation of param minToMint
    function _mint(uint256 minToMint) private notZeroMint returns (uint256 minted) {
        accountedBalance += msg.value;

        toMint = _convertIotxToUniIotx(msg.value);
        if (toMint < minToMint) revert ExchangeRatioMismatch(minToMint, toMint);
        uniIOTX.mint(msg.sender, toMint);
        minted = toMint;

        totalPending += msg.value;

        emit Minted(msg.sender, minToMint, minted);
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
            for (k = 0; k < accumulatedLen; k++) {
                if (k <= j) {
                    tokenIdsToMerge[k] = stakedTokenIds[k];
                } else {
                    tokenIdsUnmerge[k] = stakedTokenIds[k];
                }
            }

            // Reference: https://github.com/iotexproject/iip13-contracts/blob/main/src/SystemStaking.sol#L302
            systemStake.merge(tokenIdsToMerge, stakeDuration);
            stakedTokenIds[i] = tokenIdsUnmerge;
            stakedTokenIds[i+1].push(tokenIdsToMerge[0]);

            emit Merged(tokenIdsToMerge, targetAmount);
        }
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

}