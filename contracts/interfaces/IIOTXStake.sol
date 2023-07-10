// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

interface IIOTXStake {
    function onERC721Received(address, address, uint, bytes calldata) external pure override returns (bytes4);
    function exchangeRatio() external returns (uint ratio);
    function currentReserve() public view returns(uint);
    function getRedeemedTokenIds(uint i, uint j) external view returns (uint[] memory tokenIds);
    function setGlobalDelegate(address delegate) external;
    function updateDelegates(uint[] tokenIds, address delegate) external;
    function deposit(uint minToMint, uint deadline) external payable returns (uint minted);
    function stake() external;
    function redeem(uint iotxsToRedeem, uint maxToBurn, uint deadline) external returns (uint burned);
    function updateReward() external;
    function withdrawManagerFee(uint amount, address recipient) external;
}