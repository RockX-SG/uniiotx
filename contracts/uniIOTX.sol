// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20BurnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20SnapshotUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

contract uniIOTX is Initializable, ERC20Upgradeable, ERC20BurnableUpgradeable, ERC20SnapshotUpgradeable, PausableUpgradeable, AccessControlUpgradeable {
    function initialize(
        address iotxStakeAddress
    ) public initializer {
        __ERC20_init("Universal IOTX", "uniIOTX");
        __ERC20Burnable_init();
        __ERC20Snapshot_init();
        __Pausable_init();

        _grantRole(ROLE_PAUSE, iotxStakeAddress);
        _grantRole(ROLE_MINT, iotxStakeAddress);
    }

    function burn(uint amount) public override onlyRole(ROLE_MINT) {
        _burn(_msgSender(), amount);
    }

    function mint(address to, uint amount) public onlyRole(ROLE_MINT) {
        _mint(to, amount);
    }

    function snapshot() public onlyRole(ROLE_MINT) {
        _snapshot();
    }

    function pause() public onlyRole(ROLE_PAUSE) {
        _pause();
    }

    function unpause() public onlyRole(ROLE_PAUSE) {
        _unpause();
    }

    function _beforeTokenTransfer(address from, address to, uint amount)
        internal
        whenNotPaused
        override(ERC20Upgradeable, ERC20SnapshotUpgradeable)
    {
        super._beforeTokenTransfer(from, to, amount);
    }
}