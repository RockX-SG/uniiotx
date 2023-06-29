
interface IDelegateManager {
    function setGlobalDelegate(address delegate) external;
    function updateDelegates(uint256[] tokenIds, address delegate) external;
    function nextDelegate() external returns (address);
}