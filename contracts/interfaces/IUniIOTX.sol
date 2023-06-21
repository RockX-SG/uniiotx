pragma solidity ^0.8.9;

interface IUniIOTX is IERC20 {
    function mint(address account, uint256 amount) external;
    function burn(uint256 amount) external;
}