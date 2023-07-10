// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

// Roles
bytes32 constant ROLE_PAUSE = keccak256("ROLE_PAUSE");
bytes32 constant ROLE_MINT = keccak256("ROLE_MINT"); // Todo: Maybe remove it
bytes32 constant ROLE_STAKE = keccak256("ROLE_STAKE"); // Todo: Maybe remove it
bytes32 constant ROLE_FEE_MANAGER = keccak256("ROLE_FEE_MANAGER");
bytes32 constant ROLE_ORACLE = keccak256("ROLE_ORACLE");

bytes32 constant ROLE_PROTOCOL_MANAGER = keccak256("ROLE_PROTOCOL_MANAGER"); // Todo: Maybe remove it