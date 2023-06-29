pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";

import "../interfaces/ISystemStake.sol";

contract DelegateManager is Initializable, AccessControlUpgradeable, PausableUpgradeable {
    ISystemStake public systemStakeContract;

    address public globalDelegate;

    function initialize(address _systemStakeContract) public initializer {
        _grantRole(ORACLE_ROLE, msg.sender);
        _grantRole(PAUSER_ROLE, msg.sender);

        systemStakeContract = _systemStakeContract;
    }

    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    function setGlobalDelegate(address delegate) external whenNotPaused onlyRole(ORACLE_ROLE) {
        globalDelegate = delegate;
    }

    function updateDelegates(uint256[] tokenIds, address delegate) external pure whenNotPaused onlyRole(ORACLE_ROLE) {
        systemStakeContract.changeDelegates(tokenIds, delegate);
    }

    function nextDelegate() external view whenNotPaused returns (address) {
        return globalDelegate;
    }
}
