pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

contract RockXStaking is Initializable, PausableUpgradeable, AccessControlUpgradeable {
    // errors
    error MintZero();

    address public iotexSystemStakingContract;  // IoTeX system staking contract
    address public xIOTEXAddress;               // xIOTEX token address
    uint256 private totalPending;               // total pending IOTEXs awaiting to be staked

    /**
     * ======================================================================================
     *
     * SYSTEM SETTINGS, OPERATED VIA OWNER(DAO/TIMELOCK)
     *
     * ======================================================================================
     */

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
    function initialize() initializer public {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(PAUSER_ROLE, msg.sender);
    }


    /**
     * @dev set eth deposit contract address
     */
    function setIOTEXSystemStakingContract(address _iotexSystemStakingContract) external onlyRole(DEFAULT_ADMIN_ROLE) {
        iotexSystemStakingContract = _iotexSystemStakingContract;

        emit SystemStakingContractSet(_iotexSystemStakingContract);
    }

    /**
     * @dev set xIOTEX token contract address
     */
    function setXIOTEXContractAddress(address _xIOTEXAddress) external onlyRole(DEFAULT_ADMIN_ROLE) {
        xIOTEXAddress = _xIOTEXAddress;

        emit XIOTEXContractSet(_xIOTEXAddress);
    }


    /**
     * ======================================================================================
     *
     * VIEW FUNCTIONS
     *
     * ======================================================================================
     */

    /**
     * @dev return pending iotexs
     */
    function getPendingIotexs() external view returns (uint256) { return totalPending; }


    /**
     * ======================================================================================
     *
     * EXTERNAL FUNCTIONS
     *
     * ======================================================================================
     */

    /**
     * @dev mint xIOTEX with IOTEX
     */
    function mint() external payable returns (uint256 minted) {
        amount = msg.value;
        if  (amount == 0) revert MintZero();

        // merge
        merge();

        // TODO: to be optimized

        // mint xIOTEX
        uint256 toMint = 1 * amount; // default exchange ratio 1:1
        IMintableContract(xIOTEXAddress).mint(msg.sender, toMint);
        totalPending += amount;
        emit DepositEvent(msg.sender, amount);

        // stake
        stake();

        return toMint;
    }


    function merge() private {
        // TODO: to be implemented
    }

    function stake() private {
        // TODO: to be implemented
    }

    /**
     * ======================================================================================
     *
     * ROCKX SYSTEM EVENTS
     *
     * ======================================================================================
     */
    event SystemStakingContractSet(address addr);
    event XIOTEXContractSet(address addr);
    event DepositEvent(address indexed from, uint256 amount);
}