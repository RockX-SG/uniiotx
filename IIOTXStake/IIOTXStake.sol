pragma solidity ^0.8.9;

// @notice This is the IoTeX system contract interface. It issues an NFT token for each bucket creation.
// For more information see https://github.com/iotexproject/iip13-contracts/blob/main/src/SystemStaking.sol
interface IIOTXStake {
    function bucketOf(uint256 _tokenId) external view returns ( uint256 amount_, uint256 duration_, uint256 unlockedAt_, uint256 unstakedAt_, address delegate_);

    function stake(uint256 _duration, address _delegate) external payable returns (uint256);
//    function stake(uint256 _amount, uint256 _duration, address[] memory _delegates) external payable returns (uint256 firstTokenId_);
    function stake(uint256 _amount, uint256 _duration, address _delegate, uint256 _count) external payable returns (uint256 firstTokenId_);


    function unlock(uint256 _tokenId) external whenNotPaused onlyTokenOwner(_tokenId);
    function unlock(uint256[] calldata _tokenIds) external;

    function unstake(uint256 _tokenId) external whenNotPaused onlyTokenOwner(_tokenId);
    function unstake(uint256[] calldata _tokenIds) external;

    function withdraw( uint256 _tokenId, address payable _recipient) external whenNotPaused onlyTokenOwner(_tokenId);
    function withdraw(uint256[] calldata _tokenIds, address payable _recipient) external whenNotPaused;

    function merge(uint256[] calldata tokenIds, uint256 _newDuration) external payable;
}
