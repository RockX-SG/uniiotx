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

import "../interfaces/Roles.sol";

contract UniIOTX is Initializable, ERC20Upgradeable, ERC20BurnableUpgradeable, ERC20SnapshotUpgradeable, PausableUpgradeable, AccessControlUpgradeable {
    /**
     * @dev @custom:oz-upgrades-unsafe-allow constructor
     */
    constructor() {
        _disableInitializers();
    }

    /**
     * @dev This function initializes the contract
     */
    function initialize(
        address iotxStaker
    ) public initializer {
        // Init
        __ERC20_init("Universal IOTX", "uniIOTX");
        __ERC20Burnable_init();
        __ERC20Snapshot_init();
        __Pausable_init();
        __AccessControl_init();

        // Roles
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _setupRole(ROLE_PAUSER, iotxStaker);
        _setupRole(ROLE_MINTER, iotxStaker);
    }

    /**
     * @dev This function burns the specified quantity of uniIOTXs for the 'MINTER'
     */
    function burn(uint amount) public override onlyRole(ROLE_MINTER) {
        _burn(_msgSender(), amount);
    }

    /**
     * @dev This function burns the specified quantity of uniIOTXs for the specified account
     */
    function burnFrom(address account, uint256 amount) public override onlyRole(ROLE_MINTER) {
        _spendAllowance(account, _msgSender(), amount);
        _burn(account, amount);
    }

    /**
     * @dev This function mints the specified quantity of uniIOTXs for the 'MINTER'
     */
    function mint(address to, uint amount) public onlyRole(ROLE_MINTER) {
        _mint(to, amount);
    }

    /**
     * @dev This function captures a snapshot.
     */
    function snapshot() public onlyRole(ROLE_MINTER) {
        _snapshot();
    }

    /**
     * @dev This function pauses the contract
     */
    function pause() public onlyRole(ROLE_PAUSER) {
        _pause();
    }

    /**
     * @dev This function unpauses the contract
     */
    function unpause() public onlyRole(ROLE_PAUSER) {
        _unpause();
    }

    function _beforeTokenTransfer(address from, address to, uint amount)
        internal
        whenNotPaused
        override(ERC20Upgradeable, ERC20SnapshotUpgradeable)
    {
        super._beforeTokenTransfer(from, to, amount);
    }

    /**
     * @dev This function batch transfers amount to recipient
     * @notice that excessive gas consumption causes transaction revert
     */
    function batchTransfer(address[] memory recipients, uint256[] memory amounts) public {
        require(recipients.length > 0, "uniIOTX: least one recipient address");
        require(recipients.length == amounts.length, "uniIOTX: number of recipient addresses does not match the number of tokens");

        for(uint256 i = 0; i < recipients.length; ++i) {
            _transfer(_msgSender(), recipients[i], amounts[i]);
        }
    }
}