pragma solidity ^0.8.9;

interface IIOTXClear {
    function onERC721Received(address, address, uint256, bytes calldata) external pure override returns (bytes4);
    function updateDelegates(uint256[] tokenIds, address delegate) external;
    function getReward(address acount) external returns (uint256);
    function joinDebt(address claimAddr, uint256 amount) public;
    function unstake(uint256[] calldata tokenIds) external;
    function withdraw(uint256[] calldata tokenIds) external;
    function updateReward(address acount) external returns (uint256);
    function claimRewards(uint256 amount) external;
}
