pragma solidity ^0.8.9;

interface IMintableContract is IERC20 {
    function mint(address account, uint256 amount) external;
}

// @notice This is the IoTeX system contract interface. It issues an NFT token for each bucket creation.
// For more information see https://github.com/iotexproject/iip13-contracts/blob/main/src/SystemStaking.sol
interface SystemStakingContract {
    function stake(uint256 _duration, address _delegate) external payable returns (uint256);
    function merge(uint256[] calldata tokenIds, uint256 _newDuration) external payable;
}

interface IRockXRedeem {
    function pay(address account) external payable;
}
