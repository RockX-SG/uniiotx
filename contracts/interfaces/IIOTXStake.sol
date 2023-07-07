pragma solidity ^0.8.9;

interface IIOTXStake {
    function onERC721Received(address, address, uint256, bytes calldata) external pure override returns (bytes4);
    function exchangeRatio() external returns (uint256 ratio);
    function currentReserve() public view returns(uint256);
    function getRedeemedTokenIds(uint256 i, uint256 j) external view returns (uint256[] memory tokenIds);
    function setGlobalDelegate(address delegate) external;
    function updateDelegates(uint256[] tokenIds, address delegate) external;
    function deposit(uint256 minToMint, uint256 deadline) external payable returns (uint256 minted);
    function stake() external;
    function redeem(uint256 iotxsToRedeem, uint256 maxToBurn, uint256 deadline) external returns (uint256 burned);
    function updateReward() external;
    function withdrawManagerFee(uint256 amount, address recipient) external;
}