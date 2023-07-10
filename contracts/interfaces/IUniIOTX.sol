pragma solidity ^0.8.9;

interface IUniIOTX is IERC20 {
    function mint(address account, uint amount) external;
    function burn(uint amount) external;
}