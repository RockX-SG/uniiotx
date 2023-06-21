
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20BurnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20SnapshotUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";

contract uniIOTX is Initializable, ERC20Upgradeable, ERC20BurnableUpgradeable, ERC20SnapshotUpgradeable, OwnableUpgradeable, PausableUpgradeable {
    /**
      * @dev Emitted when an account is set mintable
     */
    event Mintable(address account);
    /**
     * @dev Emitted when an account is set unmintable
     */
    event Unmintable(address account);

    // @dev mintable group
    mapping(address => bool) public mintableGroup;

    modifier onlyMintableGroup() {
        require(mintableGroup[msg.sender], "uniIOTX: not in mintable group");
        _;
    }

    function initialize() initializer public {
        __ERC20_init("Universal IOTX", "uniIOTX");
        __ERC20Burnable_init();
        __ERC20Snapshot_init();
        __Ownable_init();
        __Pausable_init();

        setMintable(owner(), true); // default mintable at constructor
    }

    /**
     * @dev set or remove address to mintable group
     */
    function setMintable(address account, bool allow) public onlyOwner {
        require(mintableGroup[account] != allow, "already set");
        mintableGroup[account] = allow;

        if (allow) {
            emit Mintable(account);
        }  else {
            emit Unmintable(account);
        }
    }

    /**
     * @dev Destroys `amount` tokens from the user.
     *
     * See {ERC20-_burn}.
     */
    function burn(uint256 amount) public override {
        _burn(_msgSender(), amount);
    }

    /**
     * @dev Destroys `amount` tokens from `account`, deducting from the caller's
     * allowance.
     *
     * See {ERC20-_burn} and {ERC20-allowance}.
     *
     * Requirements:
     *
     * - the caller must have allowance for ``accounts``'s tokens of at least
     * `amount`.
     */
    function burnFrom(address account, uint256 amount) public override onlyMintableGroup {
        _spendAllowance(account, _msgSender(), amount);
        _burn(account, amount);
    }

    function mint(address to, uint256 amount) public {
        _mint(to, amount);
    }

    function snapshot() public onlyOwner {
        _snapshot();
    }

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

    function _beforeTokenTransfer(address from, address to, uint256 amount)
        internal
        whenNotPaused
        override(ERC20Upgradeable, ERC20SnapshotUpgradeable)
    {
        super._beforeTokenTransfer(from, to, amount);
    }

    /**
     * @dev Batch transfer amount to recipient
     * @notice that excessive gas consumption causes transaction revert
     */
    function batchTransfer(address[] memory recipients, uint256[] memory amounts) public {
        require(recipients.length > 0, "uniIOTX: least one recipient address");
        require(recipients.length == amounts.length, "uniIOTX: number of recipient addresses does not match the number of tokens");

        for(uint256 i = 0; i < recipients.length; ++i) {
            _transfer(_msgSender(), recipients[i], amounts[i]);
        }
    }
}