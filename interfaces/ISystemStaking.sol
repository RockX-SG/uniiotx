// SPDX-License-Identifier: MIT
/*
 * ==================================================================
 * Copyright (C) 2023 Altstake Technology Pte. Ltd. (RockX)
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

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

// @notice This is the IoTeX system contract interface. It issues an NFT token for each bucket creation.
// For more information see https://github.com/iotexproject/iip13-contracts/blob/main/src/SystemStaking.sol
interface ISystemStaking is IERC721 {
    function isActiveBucketType(uint256 _amount, uint256 _duration) external view returns (bool);
    function bucketOf(uint _tokenId) external view returns ( uint amount_, uint duration_, uint unlockedAt_, uint unstakedAt_, address delegate_);

    function stake(uint _amount, uint _duration, address _delegate, uint _count) external payable returns (uint firstTokenId_);
    function unlock(uint[] calldata _tokenIds) external;
    function unstake(uint[] calldata _tokenIds) external;
    function withdraw( uint[] calldata _tokenIds, address payable _recipient) external;

    function merge(uint[] calldata tokenIds, uint _newDuration) external payable;

    function changeDelegates(uint[] calldata _tokenIds, address _delegate) external;
}
