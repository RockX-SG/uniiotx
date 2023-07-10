// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

interface IIOTXClear {
    function onERC721Received(address, address, uint, bytes calldata) external pure override returns (bytes4);
    function updateDelegates(uint[] tokenIds, address delegate) external;
    function getReward(address acount) external returns (uint);
    function joinDebt(address claimAddr, uint amount) public;
    function unstake(uint[] calldata tokenIds) external;
    function withdraw(uint[] calldata tokenIds) external;
    function updateReward(address acount) external returns (uint);
    function claimRewards(uint amount) external;
}
