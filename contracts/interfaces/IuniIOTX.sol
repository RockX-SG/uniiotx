// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";

interface IuniIOTX is ERC20Upgradeable {
    function mint(address account, uint amount) external;
    function burn(uint amount) external;
}