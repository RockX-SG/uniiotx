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

interface IIOTXClear {
    function onERC721Received(address, address, uint, bytes calldata) external pure returns (bytes4);
    function updateDelegates(uint[] calldata tokenIds, address delegate) external;
    function getReward(address acount) external returns (uint);
    function joinDebt(address claimAddr, uint amount) external;
    function unstake(uint[] calldata tokenIds) external;
    function payDebts(uint[] calldata tokenIds) external;
    function claimRewards(uint amount, address recipient) external;
}
