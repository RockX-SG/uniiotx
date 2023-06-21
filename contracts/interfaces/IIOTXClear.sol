pragma solidity ^0.8.9;

interface IIOTXClear {
    function joinDebt(address claimAddr, uint256 amount) external;
    function unstake(uint256[] calldata tokenIds) external;
    function withdraw(uint256[] calldata tokenIds) external;
    function claimRewards(uint256 amount) external;
    function changeDelegate(uint256 tokenId, address delegate) external;
    function setDelegate(address delegate);
    function payDebt(address account) external payable;
}
