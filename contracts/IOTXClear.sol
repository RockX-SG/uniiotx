// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "interfaces/IUniIOTX.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";

contract IOTXClear is IIOTXClear, Initializable, PausableUpgradeable, ReentrancyGuardUpgradeable, OwnableUpgradeable, IERC721Receiver {
    using SafeERC20 for IERC20;
    using Address for address payable;

    mapping(address=>uint256) private balances;
    uint256 private totalBalance;

    event Paied(address account, uint256 amount);
    event Claimed(address indexed _from, address indexed _to, uint256 _value);

    /**
     * @dev initialization
     */
    function initialize() initializer public {
        __Ownable_init();
        __Pausable_init();
        __ReentrancyGuard_init();
    }

    // some convenient method to help show their claimable in wallet
    function name() external pure returns (string memory) { return "RockX Claimable IOTX"; }
    function symbol() external pure returns (string memory) { return "redeemIOTX"; }
    function decimals() external pure returns (uint8) { return 18; }
    function totalSupply() external view returns (uint256) { return totalBalance; }
    function balanceOf(address account) external view returns(uint256) { return balances[account]; }

    function claim(address to, uint256 amount) public nonReentrant returns (bool success) {
        // check
        require(balances[msg.sender] >= amount, "INSUFFICIENT_BALANCE");

        // modify
        balances[msg.sender] -= amount;
        totalBalance -= amount;
        payable(to).sendValue(amount);

        // log
        emit Claimed(msg.sender, to, amount);
        return true;
    }

    /**
     * @dev pay debts from rockx staking contract
     */
    function payDebt(address account) external override payable {
        balances[account] += msg.value;
        totalBalance += msg.value;

        // log
        emit Paied(account, msg.value);
    }


    /**
     * @dev IERC721Receiver implement for receiving staking NFT
     */
    function onERC721Received(
        address, // operator
        address, // from
        uint256, // tokenId
        bytes calldata // data
    ) external pure override returns (bytes4) {
        return this.onERC721Received.selector;
    }
}
