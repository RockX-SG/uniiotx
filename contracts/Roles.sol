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

// Roles
bytes32 constant ROLE_PAUSE = keccak256("ROLE_PAUSE");
bytes32 constant ROLE_MINT = keccak256("ROLE_MINT");
bytes32 constant ROLE_STAKE = keccak256("ROLE_STAKE");
bytes32 constant ROLE_FEE_MANAGER = keccak256("ROLE_FEE_MANAGER");
bytes32 constant ROLE_ORACLE = keccak256("ROLE_ORACLE");