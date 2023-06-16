pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";

contract RockXIOTEX is Initializable, ERC20Upgradeable, PausableUpgradeable {
    function initialize() initializer public {
        __ERC20_init("Universal IOTEX", "uniIOTEX");
        __Pausable_init();
    }

    function mint(address to, uint256 amount) public {
        _mint(to, amount);
    }

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }
}