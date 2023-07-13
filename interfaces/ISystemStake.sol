// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

// @notice This is the IoTeX system contract interface. It issues an NFT token for each bucket creation.
// For more information see https://github.com/iotexproject/iip13-contracts/blob/main/src/SystemStaking.sol
interface ISystemStake is IERC721 {
    function bucketOf(uint _tokenId) external view returns ( uint amount_, uint duration_, uint unlockedAt_, uint unstakedAt_, address delegate_);

    function stake(uint _amount, uint _duration, address _delegate, uint _count) external payable returns (uint firstTokenId_);
    function unlock(uint[] calldata _tokenIds) external;
    function unstake(uint[] calldata _tokenIds) external;
    function withdraw( uint _tokenId, address payable _recipient) external;

    function merge(uint[] calldata tokenIds, uint _newDuration) external payable;

    function changeDelegates(uint[] calldata _tokenIds, address _delegate) external;
}
