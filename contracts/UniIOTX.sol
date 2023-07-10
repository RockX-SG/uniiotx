// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20BurnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20SnapshotUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

contract UniIOTX is Initializable, ERC20Upgradeable, ERC20BurnableUpgradeable, ERC20SnapshotUpgradeable, PausableUpgradeable, AccessControlUpgradeable {
    function initialize(
        address iotxStakeAddress
    ) public initializer {
        __ERC20_init("Universal IOTX", "uniIOTX");
        __ERC20Burnable_init();
        __ERC20Snapshot_init();
        __Pausable_init();

        _grantRole(MINTER_ROLE, iotxStakeAddress);
        _grantRole(PAUSE_ROLE, iotxStakeAddress);
    }

    function burn(uint amount) public override onlyRole(MINTER_ROLE) {
        _burn(_msgSender(), amount);
    }

    function mint(address to, uint amount) public onlyRole(MINTER_ROLE) {
        _mint(to, amount);
    }

    function snapshot() public onlyRole(MINTER_ROLE) {
        _snapshot();
    }

    function pause() public onlyRole(PAUSE_ROLE) {
        _pause();
    }

    function unpause() public onlyRole(PAUSE_ROLE) {
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