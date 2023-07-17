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

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20BurnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20SnapshotUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

import "./Roles.sol";

contract UniIOTX is Initializable, ERC20Upgradeable, ERC20BurnableUpgradeable, ERC20SnapshotUpgradeable, PausableUpgradeable, AccessControlUpgradeable {
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