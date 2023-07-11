// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IUniIOTX is IERC20 {
    function mint(address account, uint amount) external;
    function burn(uint amount) external;
}