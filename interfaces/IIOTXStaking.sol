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

import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";

interface IIOTXStaking is IERC721Receiver {
    function exchangeRatio() external returns (uint ratio);
    function currentReserve() external view returns(uint);
    function redeemAmountBase() external returns(uint);
    function getRedeemedTokenIds(uint i, uint j) external view returns (uint[] memory tokenIds);
    function getStakedTokenCount(uint tokenQueueIndex) external view returns (uint count);
    function setGlobalDelegate(address delegate) external;
    function updateDelegates(uint[] calldata tokenIds, address delegate) external;
    function deposit(uint deadline) external payable returns (uint minted);
    function stake() external;
    function redeem(uint iotxsToRedeem, uint deadline) external returns (uint burned);
    function updateReward() external;
    function withdrawManagerFee(uint amount, address recipient) external;
}