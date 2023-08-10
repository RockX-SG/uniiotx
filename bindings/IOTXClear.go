// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IOTXClearDebt is an auto generated low-level Go binding around an user-defined struct.
type IOTXClearDebt struct {
	Account common.Address
	Amount  *big.Int
}

// IOTXClearUserInfo is an auto generated low-level Go binding around an user-defined struct.
type IOTXClearUserInfo struct {
	Debt       *big.Int
	Principal  *big.Int
	Reward     *big.Int
	RewardRate *big.Int
}

// IOTXClearMetaData contains all meta data concerning the IOTXClear contract.
var IOTXClearMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtIndex\",\"type\":\"uint256\"}],\"name\":\"DebtPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtIndex\",\"type\":\"uint256\"}],\"name\":\"DebtQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"DelegatesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PrincipalClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"claimPrincipals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"debtAmountBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDebtAmountBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPaidDebtItemCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPrincipal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalDebtItemCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalDebts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"unpaidDebtIndex\",\"type\":\"uint256\"}],\"name\":\"getUnpaidDebtItem\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIOTXClear.Debt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnpaidDebtItemCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnpaidDebtItems\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIOTXClear.Debt[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getUserInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"}],\"internalType\":\"structIOTXClear.UserInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_systemStaking\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_iotxStaking\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"joinDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"payDebts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemStaking\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDebts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"updateDelegates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// IOTXClearABI is the input ABI used to generate the binding from.
// Deprecated: Use IOTXClearMetaData.ABI instead.
var IOTXClearABI = IOTXClearMetaData.ABI

// IOTXClear is an auto generated Go binding around an Ethereum contract.
type IOTXClear struct {
	IOTXClearCaller     // Read-only binding to the contract
	IOTXClearTransactor // Write-only binding to the contract
	IOTXClearFilterer   // Log filterer for contract events
}

// IOTXClearCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOTXClearCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOTXClearTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOTXClearTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOTXClearFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOTXClearFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOTXClearSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOTXClearSession struct {
	Contract     *IOTXClear        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOTXClearCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOTXClearCallerSession struct {
	Contract *IOTXClearCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IOTXClearTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOTXClearTransactorSession struct {
	Contract     *IOTXClearTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IOTXClearRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOTXClearRaw struct {
	Contract *IOTXClear // Generic contract binding to access the raw methods on
}

// IOTXClearCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOTXClearCallerRaw struct {
	Contract *IOTXClearCaller // Generic read-only contract binding to access the raw methods on
}

// IOTXClearTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOTXClearTransactorRaw struct {
	Contract *IOTXClearTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOTXClear creates a new instance of IOTXClear, bound to a specific deployed contract.
func NewIOTXClear(address common.Address, backend bind.ContractBackend) (*IOTXClear, error) {
	contract, err := bindIOTXClear(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOTXClear{IOTXClearCaller: IOTXClearCaller{contract: contract}, IOTXClearTransactor: IOTXClearTransactor{contract: contract}, IOTXClearFilterer: IOTXClearFilterer{contract: contract}}, nil
}

// NewIOTXClearCaller creates a new read-only instance of IOTXClear, bound to a specific deployed contract.
func NewIOTXClearCaller(address common.Address, caller bind.ContractCaller) (*IOTXClearCaller, error) {
	contract, err := bindIOTXClear(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOTXClearCaller{contract: contract}, nil
}

// NewIOTXClearTransactor creates a new write-only instance of IOTXClear, bound to a specific deployed contract.
func NewIOTXClearTransactor(address common.Address, transactor bind.ContractTransactor) (*IOTXClearTransactor, error) {
	contract, err := bindIOTXClear(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOTXClearTransactor{contract: contract}, nil
}

// NewIOTXClearFilterer creates a new log filterer instance of IOTXClear, bound to a specific deployed contract.
func NewIOTXClearFilterer(address common.Address, filterer bind.ContractFilterer) (*IOTXClearFilterer, error) {
	contract, err := bindIOTXClear(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOTXClearFilterer{contract: contract}, nil
}

// bindIOTXClear binds a generic wrapper to an already deployed contract.
func bindIOTXClear(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOTXClearABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOTXClear *IOTXClearRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOTXClear.Contract.IOTXClearCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOTXClear *IOTXClearRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXClear.Contract.IOTXClearTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOTXClear *IOTXClearRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOTXClear.Contract.IOTXClearTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOTXClear *IOTXClearCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOTXClear.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOTXClear *IOTXClearTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXClear.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOTXClear *IOTXClearTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOTXClear.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IOTXClear *IOTXClearCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IOTXClear *IOTXClearSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IOTXClear.Contract.DEFAULTADMINROLE(&_IOTXClear.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IOTXClear *IOTXClearCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IOTXClear.Contract.DEFAULTADMINROLE(&_IOTXClear.CallOpts)
}

// AccountedBalance is a free data retrieval call binding the contract method 0x0937eb54.
//
// Solidity: function accountedBalance() view returns(uint256)
func (_IOTXClear *IOTXClearCaller) AccountedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "accountedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountedBalance is a free data retrieval call binding the contract method 0x0937eb54.
//
// Solidity: function accountedBalance() view returns(uint256)
func (_IOTXClear *IOTXClearSession) AccountedBalance() (*big.Int, error) {
	return _IOTXClear.Contract.AccountedBalance(&_IOTXClear.CallOpts)
}

// AccountedBalance is a free data retrieval call binding the contract method 0x0937eb54.
//
// Solidity: function accountedBalance() view returns(uint256)
func (_IOTXClear *IOTXClearCallerSession) AccountedBalance() (*big.Int, error) {
	return _IOTXClear.Contract.AccountedBalance(&_IOTXClear.CallOpts)
}

// DebtAmountBase is a free data retrieval call binding the contract method 0x257f9133.
//
// Solidity: function debtAmountBase() view returns(uint256)
func (_IOTXClear *IOTXClearCaller) DebtAmountBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "debtAmountBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtAmountBase is a free data retrieval call binding the contract method 0x257f9133.
//
// Solidity: function debtAmountBase() view returns(uint256)
func (_IOTXClear *IOTXClearSession) DebtAmountBase() (*big.Int, error) {
	return _IOTXClear.Contract.DebtAmountBase(&_IOTXClear.CallOpts)
}

// DebtAmountBase is a free data retrieval call binding the contract method 0x257f9133.
//
// Solidity: function debtAmountBase() view returns(uint256)
func (_IOTXClear *IOTXClearCallerSession) DebtAmountBase() (*big.Int, error) {
	return _IOTXClear.Contract.DebtAmountBase(&_IOTXClear.CallOpts)
}

// GetAccountedBalance is a free data retrieval call binding the contract method 0x33e5761f.
//
// Solidity: function getAccountedBalance() view returns(uint256)
func (_IOTXClear *IOTXClearCaller) GetAccountedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "getAccountedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountedBalance is a free data retrieval call binding the contract method 0x33e5761f.
//
// Solidity: function getAccountedBalance() view returns(uint256)
func (_IOTXClear *IOTXClearSession) GetAccountedBalance() (*big.Int, error) {
	return _IOTXClear.Contract.GetAccountedBalance(&_IOTXClear.CallOpts)
}

// GetAccountedBalance is a free data retrieval call binding the contract method 0x33e5761f.
//
// Solidity: function getAccountedBalance() view returns(uint256)
func (_IOTXClear *IOTXClearCallerSession) GetAccountedBalance() (*big.Int, error) {
	return _IOTXClear.Contract.GetAccountedBalance(&_IOTXClear.CallOpts)
}

// GetDebt is a free data retrieval call binding the contract method 0x9a78e72e.
//
// Solidity: function getDebt(address account) view returns(uint256)
func (_IOTXClear *IOTXClearCaller) GetDebt(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "getDebt", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDebt is a free data retrieval call binding the contract method 0x9a78e72e.
//
// Solidity: function getDebt(address account) view returns(uint256)
func (_IOTXClear *IOTXClearSession) GetDebt(account common.Address) (*big.Int, error) {
	return _IOTXClear.Contract.GetDebt(&_IOTXClear.CallOpts, account)
}

// GetDebt is a free data retrieval call binding the contract method 0x9a78e72e.
//
// Solidity: function getDebt(address account) view returns(uint256)
func (_IOTXClear *IOTXClearCallerSession) GetDebt(account common.Address) (*big.Int, error) {
	return _IOTXClear.Contract.GetDebt(&_IOTXClear.CallOpts, account)
}

// GetDebtAmountBase is a free data retrieval call binding the contract method 0x8dc43dd3.
//
// Solidity: function getDebtAmountBase() view returns(uint256)
func (_IOTXClear *IOTXClearCaller) GetDebtAmountBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "getDebtAmountBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDebtAmountBase is a free data retrieval call binding the contract method 0x8dc43dd3.
//
// Solidity: function getDebtAmountBase() view returns(uint256)
func (_IOTXClear *IOTXClearSession) GetDebtAmountBase() (*big.Int, error) {
	return _IOTXClear.Contract.GetDebtAmountBase(&_IOTXClear.CallOpts)
}

// GetDebtAmountBase is a free data retrieval call binding the contract method 0x8dc43dd3.
//
// Solidity: function getDebtAmountBase() view returns(uint256)
func (_IOTXClear *IOTXClearCallerSession) GetDebtAmountBase() (*big.Int, error) {
	return _IOTXClear.Contract.GetDebtAmountBase(&_IOTXClear.CallOpts)
}

// GetPaidDebtItemCount is a free data retrieval call binding the contract method 0xf3ec464e.
//
// Solidity: function getPaidDebtItemCount() view returns(uint256)
func (_IOTXClear *IOTXClearCaller) GetPaidDebtItemCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "getPaidDebtItemCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPaidDebtItemCount is a free data retrieval call binding the contract method 0xf3ec464e.
//
// Solidity: function getPaidDebtItemCount() view returns(uint256)
func (_IOTXClear *IOTXClearSession) GetPaidDebtItemCount() (*big.Int, error) {
	return _IOTXClear.Contract.GetPaidDebtItemCount(&_IOTXClear.CallOpts)
}

// GetPaidDebtItemCount is a free data retrieval call binding the contract method 0xf3ec464e.
//
// Solidity: function getPaidDebtItemCount() view returns(uint256)
func (_IOTXClear *IOTXClearCallerSession) GetPaidDebtItemCount() (*big.Int, error) {
	return _IOTXClear.Contract.GetPaidDebtItemCount(&_IOTXClear.CallOpts)
}

// GetPrincipal is a free data retrieval call binding the contract method 0x7abdd2d1.
//
// Solidity: function getPrincipal(address account) view returns(uint256)
func (_IOTXClear *IOTXClearCaller) GetPrincipal(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "getPrincipal", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrincipal is a free data retrieval call binding the contract method 0x7abdd2d1.
//
// Solidity: function getPrincipal(address account) view returns(uint256)
func (_IOTXClear *IOTXClearSession) GetPrincipal(account common.Address) (*big.Int, error) {
	return _IOTXClear.Contract.GetPrincipal(&_IOTXClear.CallOpts, account)
}

// GetPrincipal is a free data retrieval call binding the contract method 0x7abdd2d1.
//
// Solidity: function getPrincipal(address account) view returns(uint256)
func (_IOTXClear *IOTXClearCallerSession) GetPrincipal(account common.Address) (*big.Int, error) {
	return _IOTXClear.Contract.GetPrincipal(&_IOTXClear.CallOpts, account)
}

// GetReward is a free data retrieval call binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address account) view returns(uint256)
func (_IOTXClear *IOTXClearCaller) GetReward(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "getReward", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReward is a free data retrieval call binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address account) view returns(uint256)
func (_IOTXClear *IOTXClearSession) GetReward(account common.Address) (*big.Int, error) {
	return _IOTXClear.Contract.GetReward(&_IOTXClear.CallOpts, account)
}

// GetReward is a free data retrieval call binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address account) view returns(uint256)
func (_IOTXClear *IOTXClearCallerSession) GetReward(account common.Address) (*big.Int, error) {
	return _IOTXClear.Contract.GetReward(&_IOTXClear.CallOpts, account)
}

// GetRewardRate is a free data retrieval call binding the contract method 0x7e1a3786.
//
// Solidity: function getRewardRate() view returns(uint256)
func (_IOTXClear *IOTXClearCaller) GetRewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "getRewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardRate is a free data retrieval call binding the contract method 0x7e1a3786.
//
// Solidity: function getRewardRate() view returns(uint256)
func (_IOTXClear *IOTXClearSession) GetRewardRate() (*big.Int, error) {
	return _IOTXClear.Contract.GetRewardRate(&_IOTXClear.CallOpts)
}

// GetRewardRate is a free data retrieval call binding the contract method 0x7e1a3786.
//
// Solidity: function getRewardRate() view returns(uint256)
func (_IOTXClear *IOTXClearCallerSession) GetRewardRate() (*big.Int, error) {
	return _IOTXClear.Contract.GetRewardRate(&_IOTXClear.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IOTXClear *IOTXClearCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IOTXClear *IOTXClearSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IOTXClear.Contract.GetRoleAdmin(&_IOTXClear.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IOTXClear *IOTXClearCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IOTXClear.Contract.GetRoleAdmin(&_IOTXClear.CallOpts, role)
}

// GetTotalDebtItemCount is a free data retrieval call binding the contract method 0x8ffe1db4.
//
// Solidity: function getTotalDebtItemCount() view returns(uint256)
func (_IOTXClear *IOTXClearCaller) GetTotalDebtItemCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "getTotalDebtItemCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDebtItemCount is a free data retrieval call binding the contract method 0x8ffe1db4.
//
// Solidity: function getTotalDebtItemCount() view returns(uint256)
func (_IOTXClear *IOTXClearSession) GetTotalDebtItemCount() (*big.Int, error) {
	return _IOTXClear.Contract.GetTotalDebtItemCount(&_IOTXClear.CallOpts)
}

// GetTotalDebtItemCount is a free data retrieval call binding the contract method 0x8ffe1db4.
//
// Solidity: function getTotalDebtItemCount() view returns(uint256)
func (_IOTXClear *IOTXClearCallerSession) GetTotalDebtItemCount() (*big.Int, error) {
	return _IOTXClear.Contract.GetTotalDebtItemCount(&_IOTXClear.CallOpts)
}

// GetTotalDebts is a free data retrieval call binding the contract method 0x477562a7.
//
// Solidity: function getTotalDebts() view returns(uint256)
func (_IOTXClear *IOTXClearCaller) GetTotalDebts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "getTotalDebts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDebts is a free data retrieval call binding the contract method 0x477562a7.
//
// Solidity: function getTotalDebts() view returns(uint256)
func (_IOTXClear *IOTXClearSession) GetTotalDebts() (*big.Int, error) {
	return _IOTXClear.Contract.GetTotalDebts(&_IOTXClear.CallOpts)
}

// GetTotalDebts is a free data retrieval call binding the contract method 0x477562a7.
//
// Solidity: function getTotalDebts() view returns(uint256)
func (_IOTXClear *IOTXClearCallerSession) GetTotalDebts() (*big.Int, error) {
	return _IOTXClear.Contract.GetTotalDebts(&_IOTXClear.CallOpts)
}

// GetUnpaidDebtItem is a free data retrieval call binding the contract method 0x36d0f761.
//
// Solidity: function getUnpaidDebtItem(uint256 unpaidDebtIndex) view returns((address,uint256))
func (_IOTXClear *IOTXClearCaller) GetUnpaidDebtItem(opts *bind.CallOpts, unpaidDebtIndex *big.Int) (IOTXClearDebt, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "getUnpaidDebtItem", unpaidDebtIndex)

	if err != nil {
		return *new(IOTXClearDebt), err
	}

	out0 := *abi.ConvertType(out[0], new(IOTXClearDebt)).(*IOTXClearDebt)

	return out0, err

}

// GetUnpaidDebtItem is a free data retrieval call binding the contract method 0x36d0f761.
//
// Solidity: function getUnpaidDebtItem(uint256 unpaidDebtIndex) view returns((address,uint256))
func (_IOTXClear *IOTXClearSession) GetUnpaidDebtItem(unpaidDebtIndex *big.Int) (IOTXClearDebt, error) {
	return _IOTXClear.Contract.GetUnpaidDebtItem(&_IOTXClear.CallOpts, unpaidDebtIndex)
}

// GetUnpaidDebtItem is a free data retrieval call binding the contract method 0x36d0f761.
//
// Solidity: function getUnpaidDebtItem(uint256 unpaidDebtIndex) view returns((address,uint256))
func (_IOTXClear *IOTXClearCallerSession) GetUnpaidDebtItem(unpaidDebtIndex *big.Int) (IOTXClearDebt, error) {
	return _IOTXClear.Contract.GetUnpaidDebtItem(&_IOTXClear.CallOpts, unpaidDebtIndex)
}

// GetUnpaidDebtItemCount is a free data retrieval call binding the contract method 0xd2e349ae.
//
// Solidity: function getUnpaidDebtItemCount() view returns(uint256)
func (_IOTXClear *IOTXClearCaller) GetUnpaidDebtItemCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "getUnpaidDebtItemCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnpaidDebtItemCount is a free data retrieval call binding the contract method 0xd2e349ae.
//
// Solidity: function getUnpaidDebtItemCount() view returns(uint256)
func (_IOTXClear *IOTXClearSession) GetUnpaidDebtItemCount() (*big.Int, error) {
	return _IOTXClear.Contract.GetUnpaidDebtItemCount(&_IOTXClear.CallOpts)
}

// GetUnpaidDebtItemCount is a free data retrieval call binding the contract method 0xd2e349ae.
//
// Solidity: function getUnpaidDebtItemCount() view returns(uint256)
func (_IOTXClear *IOTXClearCallerSession) GetUnpaidDebtItemCount() (*big.Int, error) {
	return _IOTXClear.Contract.GetUnpaidDebtItemCount(&_IOTXClear.CallOpts)
}

// GetUnpaidDebtItems is a free data retrieval call binding the contract method 0x97fe6417.
//
// Solidity: function getUnpaidDebtItems() view returns((address,uint256)[])
func (_IOTXClear *IOTXClearCaller) GetUnpaidDebtItems(opts *bind.CallOpts) ([]IOTXClearDebt, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "getUnpaidDebtItems")

	if err != nil {
		return *new([]IOTXClearDebt), err
	}

	out0 := *abi.ConvertType(out[0], new([]IOTXClearDebt)).(*[]IOTXClearDebt)

	return out0, err

}

// GetUnpaidDebtItems is a free data retrieval call binding the contract method 0x97fe6417.
//
// Solidity: function getUnpaidDebtItems() view returns((address,uint256)[])
func (_IOTXClear *IOTXClearSession) GetUnpaidDebtItems() ([]IOTXClearDebt, error) {
	return _IOTXClear.Contract.GetUnpaidDebtItems(&_IOTXClear.CallOpts)
}

// GetUnpaidDebtItems is a free data retrieval call binding the contract method 0x97fe6417.
//
// Solidity: function getUnpaidDebtItems() view returns((address,uint256)[])
func (_IOTXClear *IOTXClearCallerSession) GetUnpaidDebtItems() ([]IOTXClearDebt, error) {
	return _IOTXClear.Contract.GetUnpaidDebtItems(&_IOTXClear.CallOpts)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address account) view returns((uint256,uint256,uint256,uint256))
func (_IOTXClear *IOTXClearCaller) GetUserInfo(opts *bind.CallOpts, account common.Address) (IOTXClearUserInfo, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "getUserInfo", account)

	if err != nil {
		return *new(IOTXClearUserInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IOTXClearUserInfo)).(*IOTXClearUserInfo)

	return out0, err

}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address account) view returns((uint256,uint256,uint256,uint256))
func (_IOTXClear *IOTXClearSession) GetUserInfo(account common.Address) (IOTXClearUserInfo, error) {
	return _IOTXClear.Contract.GetUserInfo(&_IOTXClear.CallOpts, account)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address account) view returns((uint256,uint256,uint256,uint256))
func (_IOTXClear *IOTXClearCallerSession) GetUserInfo(account common.Address) (IOTXClearUserInfo, error) {
	return _IOTXClear.Contract.GetUserInfo(&_IOTXClear.CallOpts, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IOTXClear *IOTXClearCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IOTXClear *IOTXClearSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IOTXClear.Contract.HasRole(&_IOTXClear.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IOTXClear *IOTXClearCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IOTXClear.Contract.HasRole(&_IOTXClear.CallOpts, role, account)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_IOTXClear *IOTXClearCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_IOTXClear *IOTXClearSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _IOTXClear.Contract.OnERC721Received(&_IOTXClear.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_IOTXClear *IOTXClearCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _IOTXClear.Contract.OnERC721Received(&_IOTXClear.CallOpts, arg0, arg1, arg2, arg3)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IOTXClear *IOTXClearCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IOTXClear *IOTXClearSession) Paused() (bool, error) {
	return _IOTXClear.Contract.Paused(&_IOTXClear.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IOTXClear *IOTXClearCallerSession) Paused() (bool, error) {
	return _IOTXClear.Contract.Paused(&_IOTXClear.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_IOTXClear *IOTXClearCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_IOTXClear *IOTXClearSession) RewardRate() (*big.Int, error) {
	return _IOTXClear.Contract.RewardRate(&_IOTXClear.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_IOTXClear *IOTXClearCallerSession) RewardRate() (*big.Int, error) {
	return _IOTXClear.Contract.RewardRate(&_IOTXClear.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IOTXClear *IOTXClearCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IOTXClear *IOTXClearSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IOTXClear.Contract.SupportsInterface(&_IOTXClear.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IOTXClear *IOTXClearCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IOTXClear.Contract.SupportsInterface(&_IOTXClear.CallOpts, interfaceId)
}

// SystemStaking is a free data retrieval call binding the contract method 0xa5e35ef6.
//
// Solidity: function systemStaking() view returns(address)
func (_IOTXClear *IOTXClearCaller) SystemStaking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "systemStaking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemStaking is a free data retrieval call binding the contract method 0xa5e35ef6.
//
// Solidity: function systemStaking() view returns(address)
func (_IOTXClear *IOTXClearSession) SystemStaking() (common.Address, error) {
	return _IOTXClear.Contract.SystemStaking(&_IOTXClear.CallOpts)
}

// SystemStaking is a free data retrieval call binding the contract method 0xa5e35ef6.
//
// Solidity: function systemStaking() view returns(address)
func (_IOTXClear *IOTXClearCallerSession) SystemStaking() (common.Address, error) {
	return _IOTXClear.Contract.SystemStaking(&_IOTXClear.CallOpts)
}

// TotalDebts is a free data retrieval call binding the contract method 0x14a1c32d.
//
// Solidity: function totalDebts() view returns(uint256)
func (_IOTXClear *IOTXClearCaller) TotalDebts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXClear.contract.Call(opts, &out, "totalDebts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebts is a free data retrieval call binding the contract method 0x14a1c32d.
//
// Solidity: function totalDebts() view returns(uint256)
func (_IOTXClear *IOTXClearSession) TotalDebts() (*big.Int, error) {
	return _IOTXClear.Contract.TotalDebts(&_IOTXClear.CallOpts)
}

// TotalDebts is a free data retrieval call binding the contract method 0x14a1c32d.
//
// Solidity: function totalDebts() view returns(uint256)
func (_IOTXClear *IOTXClearCallerSession) TotalDebts() (*big.Int, error) {
	return _IOTXClear.Contract.TotalDebts(&_IOTXClear.CallOpts)
}

// ClaimPrincipals is a paid mutator transaction binding the contract method 0xdc8e238b.
//
// Solidity: function claimPrincipals(uint256 amount, address recipient) returns()
func (_IOTXClear *IOTXClearTransactor) ClaimPrincipals(opts *bind.TransactOpts, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IOTXClear.contract.Transact(opts, "claimPrincipals", amount, recipient)
}

// ClaimPrincipals is a paid mutator transaction binding the contract method 0xdc8e238b.
//
// Solidity: function claimPrincipals(uint256 amount, address recipient) returns()
func (_IOTXClear *IOTXClearSession) ClaimPrincipals(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IOTXClear.Contract.ClaimPrincipals(&_IOTXClear.TransactOpts, amount, recipient)
}

// ClaimPrincipals is a paid mutator transaction binding the contract method 0xdc8e238b.
//
// Solidity: function claimPrincipals(uint256 amount, address recipient) returns()
func (_IOTXClear *IOTXClearTransactorSession) ClaimPrincipals(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IOTXClear.Contract.ClaimPrincipals(&_IOTXClear.TransactOpts, amount, recipient)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6c7b69cb.
//
// Solidity: function claimRewards(uint256 amount, address recipient) returns()
func (_IOTXClear *IOTXClearTransactor) ClaimRewards(opts *bind.TransactOpts, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IOTXClear.contract.Transact(opts, "claimRewards", amount, recipient)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6c7b69cb.
//
// Solidity: function claimRewards(uint256 amount, address recipient) returns()
func (_IOTXClear *IOTXClearSession) ClaimRewards(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IOTXClear.Contract.ClaimRewards(&_IOTXClear.TransactOpts, amount, recipient)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6c7b69cb.
//
// Solidity: function claimRewards(uint256 amount, address recipient) returns()
func (_IOTXClear *IOTXClearTransactorSession) ClaimRewards(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IOTXClear.Contract.ClaimRewards(&_IOTXClear.TransactOpts, amount, recipient)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IOTXClear *IOTXClearTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXClear.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IOTXClear *IOTXClearSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXClear.Contract.GrantRole(&_IOTXClear.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IOTXClear *IOTXClearTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXClear.Contract.GrantRole(&_IOTXClear.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _systemStaking, address _iotxStaking, address _oracle) returns()
func (_IOTXClear *IOTXClearTransactor) Initialize(opts *bind.TransactOpts, _systemStaking common.Address, _iotxStaking common.Address, _oracle common.Address) (*types.Transaction, error) {
	return _IOTXClear.contract.Transact(opts, "initialize", _systemStaking, _iotxStaking, _oracle)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _systemStaking, address _iotxStaking, address _oracle) returns()
func (_IOTXClear *IOTXClearSession) Initialize(_systemStaking common.Address, _iotxStaking common.Address, _oracle common.Address) (*types.Transaction, error) {
	return _IOTXClear.Contract.Initialize(&_IOTXClear.TransactOpts, _systemStaking, _iotxStaking, _oracle)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _systemStaking, address _iotxStaking, address _oracle) returns()
func (_IOTXClear *IOTXClearTransactorSession) Initialize(_systemStaking common.Address, _iotxStaking common.Address, _oracle common.Address) (*types.Transaction, error) {
	return _IOTXClear.Contract.Initialize(&_IOTXClear.TransactOpts, _systemStaking, _iotxStaking, _oracle)
}

// JoinDebt is a paid mutator transaction binding the contract method 0x7180f85b.
//
// Solidity: function joinDebt(address account, uint256 amount) returns()
func (_IOTXClear *IOTXClearTransactor) JoinDebt(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IOTXClear.contract.Transact(opts, "joinDebt", account, amount)
}

// JoinDebt is a paid mutator transaction binding the contract method 0x7180f85b.
//
// Solidity: function joinDebt(address account, uint256 amount) returns()
func (_IOTXClear *IOTXClearSession) JoinDebt(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IOTXClear.Contract.JoinDebt(&_IOTXClear.TransactOpts, account, amount)
}

// JoinDebt is a paid mutator transaction binding the contract method 0x7180f85b.
//
// Solidity: function joinDebt(address account, uint256 amount) returns()
func (_IOTXClear *IOTXClearTransactorSession) JoinDebt(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IOTXClear.Contract.JoinDebt(&_IOTXClear.TransactOpts, account, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IOTXClear *IOTXClearTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXClear.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IOTXClear *IOTXClearSession) Pause() (*types.Transaction, error) {
	return _IOTXClear.Contract.Pause(&_IOTXClear.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IOTXClear *IOTXClearTransactorSession) Pause() (*types.Transaction, error) {
	return _IOTXClear.Contract.Pause(&_IOTXClear.TransactOpts)
}

// PayDebts is a paid mutator transaction binding the contract method 0xff50bf5e.
//
// Solidity: function payDebts(uint256[] tokenIds) returns()
func (_IOTXClear *IOTXClearTransactor) PayDebts(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _IOTXClear.contract.Transact(opts, "payDebts", tokenIds)
}

// PayDebts is a paid mutator transaction binding the contract method 0xff50bf5e.
//
// Solidity: function payDebts(uint256[] tokenIds) returns()
func (_IOTXClear *IOTXClearSession) PayDebts(tokenIds []*big.Int) (*types.Transaction, error) {
	return _IOTXClear.Contract.PayDebts(&_IOTXClear.TransactOpts, tokenIds)
}

// PayDebts is a paid mutator transaction binding the contract method 0xff50bf5e.
//
// Solidity: function payDebts(uint256[] tokenIds) returns()
func (_IOTXClear *IOTXClearTransactorSession) PayDebts(tokenIds []*big.Int) (*types.Transaction, error) {
	return _IOTXClear.Contract.PayDebts(&_IOTXClear.TransactOpts, tokenIds)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IOTXClear *IOTXClearTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXClear.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IOTXClear *IOTXClearSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXClear.Contract.RenounceRole(&_IOTXClear.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IOTXClear *IOTXClearTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXClear.Contract.RenounceRole(&_IOTXClear.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IOTXClear *IOTXClearTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXClear.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IOTXClear *IOTXClearSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXClear.Contract.RevokeRole(&_IOTXClear.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IOTXClear *IOTXClearTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXClear.Contract.RevokeRole(&_IOTXClear.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IOTXClear *IOTXClearTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXClear.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IOTXClear *IOTXClearSession) Unpause() (*types.Transaction, error) {
	return _IOTXClear.Contract.Unpause(&_IOTXClear.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IOTXClear *IOTXClearTransactorSession) Unpause() (*types.Transaction, error) {
	return _IOTXClear.Contract.Unpause(&_IOTXClear.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0xe449f341.
//
// Solidity: function unstake(uint256[] tokenIds) returns()
func (_IOTXClear *IOTXClearTransactor) Unstake(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _IOTXClear.contract.Transact(opts, "unstake", tokenIds)
}

// Unstake is a paid mutator transaction binding the contract method 0xe449f341.
//
// Solidity: function unstake(uint256[] tokenIds) returns()
func (_IOTXClear *IOTXClearSession) Unstake(tokenIds []*big.Int) (*types.Transaction, error) {
	return _IOTXClear.Contract.Unstake(&_IOTXClear.TransactOpts, tokenIds)
}

// Unstake is a paid mutator transaction binding the contract method 0xe449f341.
//
// Solidity: function unstake(uint256[] tokenIds) returns()
func (_IOTXClear *IOTXClearTransactorSession) Unstake(tokenIds []*big.Int) (*types.Transaction, error) {
	return _IOTXClear.Contract.Unstake(&_IOTXClear.TransactOpts, tokenIds)
}

// UpdateDelegates is a paid mutator transaction binding the contract method 0x2328782f.
//
// Solidity: function updateDelegates(uint256[] tokenIds, address delegate) returns()
func (_IOTXClear *IOTXClearTransactor) UpdateDelegates(opts *bind.TransactOpts, tokenIds []*big.Int, delegate common.Address) (*types.Transaction, error) {
	return _IOTXClear.contract.Transact(opts, "updateDelegates", tokenIds, delegate)
}

// UpdateDelegates is a paid mutator transaction binding the contract method 0x2328782f.
//
// Solidity: function updateDelegates(uint256[] tokenIds, address delegate) returns()
func (_IOTXClear *IOTXClearSession) UpdateDelegates(tokenIds []*big.Int, delegate common.Address) (*types.Transaction, error) {
	return _IOTXClear.Contract.UpdateDelegates(&_IOTXClear.TransactOpts, tokenIds, delegate)
}

// UpdateDelegates is a paid mutator transaction binding the contract method 0x2328782f.
//
// Solidity: function updateDelegates(uint256[] tokenIds, address delegate) returns()
func (_IOTXClear *IOTXClearTransactorSession) UpdateDelegates(tokenIds []*big.Int, delegate common.Address) (*types.Transaction, error) {
	return _IOTXClear.Contract.UpdateDelegates(&_IOTXClear.TransactOpts, tokenIds, delegate)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IOTXClear *IOTXClearTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXClear.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IOTXClear *IOTXClearSession) Receive() (*types.Transaction, error) {
	return _IOTXClear.Contract.Receive(&_IOTXClear.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IOTXClear *IOTXClearTransactorSession) Receive() (*types.Transaction, error) {
	return _IOTXClear.Contract.Receive(&_IOTXClear.TransactOpts)
}

// IOTXClearDebtPaidIterator is returned from FilterDebtPaid and is used to iterate over the raw logs and unpacked data for DebtPaid events raised by the IOTXClear contract.
type IOTXClearDebtPaidIterator struct {
	Event *IOTXClearDebtPaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IOTXClearDebtPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXClearDebtPaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IOTXClearDebtPaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IOTXClearDebtPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXClearDebtPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXClearDebtPaid represents a DebtPaid event raised by the IOTXClear contract.
type IOTXClearDebtPaid struct {
	Account   common.Address
	Amount    *big.Int
	DebtIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDebtPaid is a free log retrieval operation binding the contract event 0xc167477d6269a6698dabd8aabaff469b2fc5301c9ba41f9d187eeeb30e12488e.
//
// Solidity: event DebtPaid(address account, uint256 amount, uint256 debtIndex)
func (_IOTXClear *IOTXClearFilterer) FilterDebtPaid(opts *bind.FilterOpts) (*IOTXClearDebtPaidIterator, error) {

	logs, sub, err := _IOTXClear.contract.FilterLogs(opts, "DebtPaid")
	if err != nil {
		return nil, err
	}
	return &IOTXClearDebtPaidIterator{contract: _IOTXClear.contract, event: "DebtPaid", logs: logs, sub: sub}, nil
}

// WatchDebtPaid is a free log subscription operation binding the contract event 0xc167477d6269a6698dabd8aabaff469b2fc5301c9ba41f9d187eeeb30e12488e.
//
// Solidity: event DebtPaid(address account, uint256 amount, uint256 debtIndex)
func (_IOTXClear *IOTXClearFilterer) WatchDebtPaid(opts *bind.WatchOpts, sink chan<- *IOTXClearDebtPaid) (event.Subscription, error) {

	logs, sub, err := _IOTXClear.contract.WatchLogs(opts, "DebtPaid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXClearDebtPaid)
				if err := _IOTXClear.contract.UnpackLog(event, "DebtPaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDebtPaid is a log parse operation binding the contract event 0xc167477d6269a6698dabd8aabaff469b2fc5301c9ba41f9d187eeeb30e12488e.
//
// Solidity: event DebtPaid(address account, uint256 amount, uint256 debtIndex)
func (_IOTXClear *IOTXClearFilterer) ParseDebtPaid(log types.Log) (*IOTXClearDebtPaid, error) {
	event := new(IOTXClearDebtPaid)
	if err := _IOTXClear.contract.UnpackLog(event, "DebtPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXClearDebtQueuedIterator is returned from FilterDebtQueued and is used to iterate over the raw logs and unpacked data for DebtQueued events raised by the IOTXClear contract.
type IOTXClearDebtQueuedIterator struct {
	Event *IOTXClearDebtQueued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IOTXClearDebtQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXClearDebtQueued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IOTXClearDebtQueued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IOTXClearDebtQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXClearDebtQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXClearDebtQueued represents a DebtQueued event raised by the IOTXClear contract.
type IOTXClearDebtQueued struct {
	Account   common.Address
	Amount    *big.Int
	DebtIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDebtQueued is a free log retrieval operation binding the contract event 0xb5de2a99082d929f0850ba77b6c1ff97cd4238074ff4be97da5ae22aa0564228.
//
// Solidity: event DebtQueued(address account, uint256 amount, uint256 debtIndex)
func (_IOTXClear *IOTXClearFilterer) FilterDebtQueued(opts *bind.FilterOpts) (*IOTXClearDebtQueuedIterator, error) {

	logs, sub, err := _IOTXClear.contract.FilterLogs(opts, "DebtQueued")
	if err != nil {
		return nil, err
	}
	return &IOTXClearDebtQueuedIterator{contract: _IOTXClear.contract, event: "DebtQueued", logs: logs, sub: sub}, nil
}

// WatchDebtQueued is a free log subscription operation binding the contract event 0xb5de2a99082d929f0850ba77b6c1ff97cd4238074ff4be97da5ae22aa0564228.
//
// Solidity: event DebtQueued(address account, uint256 amount, uint256 debtIndex)
func (_IOTXClear *IOTXClearFilterer) WatchDebtQueued(opts *bind.WatchOpts, sink chan<- *IOTXClearDebtQueued) (event.Subscription, error) {

	logs, sub, err := _IOTXClear.contract.WatchLogs(opts, "DebtQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXClearDebtQueued)
				if err := _IOTXClear.contract.UnpackLog(event, "DebtQueued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDebtQueued is a log parse operation binding the contract event 0xb5de2a99082d929f0850ba77b6c1ff97cd4238074ff4be97da5ae22aa0564228.
//
// Solidity: event DebtQueued(address account, uint256 amount, uint256 debtIndex)
func (_IOTXClear *IOTXClearFilterer) ParseDebtQueued(log types.Log) (*IOTXClearDebtQueued, error) {
	event := new(IOTXClearDebtQueued)
	if err := _IOTXClear.contract.UnpackLog(event, "DebtQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXClearDelegatesUpdatedIterator is returned from FilterDelegatesUpdated and is used to iterate over the raw logs and unpacked data for DelegatesUpdated events raised by the IOTXClear contract.
type IOTXClearDelegatesUpdatedIterator struct {
	Event *IOTXClearDelegatesUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IOTXClearDelegatesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXClearDelegatesUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IOTXClearDelegatesUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IOTXClearDelegatesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXClearDelegatesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXClearDelegatesUpdated represents a DelegatesUpdated event raised by the IOTXClear contract.
type IOTXClearDelegatesUpdated struct {
	TokenIds []*big.Int
	Delegate common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegatesUpdated is a free log retrieval operation binding the contract event 0x055e970cad70d45557223e4ef35ac45bb162f8fc6f8a5986159f49d7d7fc742b.
//
// Solidity: event DelegatesUpdated(uint256[] tokenIds, address delegate)
func (_IOTXClear *IOTXClearFilterer) FilterDelegatesUpdated(opts *bind.FilterOpts) (*IOTXClearDelegatesUpdatedIterator, error) {

	logs, sub, err := _IOTXClear.contract.FilterLogs(opts, "DelegatesUpdated")
	if err != nil {
		return nil, err
	}
	return &IOTXClearDelegatesUpdatedIterator{contract: _IOTXClear.contract, event: "DelegatesUpdated", logs: logs, sub: sub}, nil
}

// WatchDelegatesUpdated is a free log subscription operation binding the contract event 0x055e970cad70d45557223e4ef35ac45bb162f8fc6f8a5986159f49d7d7fc742b.
//
// Solidity: event DelegatesUpdated(uint256[] tokenIds, address delegate)
func (_IOTXClear *IOTXClearFilterer) WatchDelegatesUpdated(opts *bind.WatchOpts, sink chan<- *IOTXClearDelegatesUpdated) (event.Subscription, error) {

	logs, sub, err := _IOTXClear.contract.WatchLogs(opts, "DelegatesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXClearDelegatesUpdated)
				if err := _IOTXClear.contract.UnpackLog(event, "DelegatesUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegatesUpdated is a log parse operation binding the contract event 0x055e970cad70d45557223e4ef35ac45bb162f8fc6f8a5986159f49d7d7fc742b.
//
// Solidity: event DelegatesUpdated(uint256[] tokenIds, address delegate)
func (_IOTXClear *IOTXClearFilterer) ParseDelegatesUpdated(log types.Log) (*IOTXClearDelegatesUpdated, error) {
	event := new(IOTXClearDelegatesUpdated)
	if err := _IOTXClear.contract.UnpackLog(event, "DelegatesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXClearInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IOTXClear contract.
type IOTXClearInitializedIterator struct {
	Event *IOTXClearInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IOTXClearInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXClearInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IOTXClearInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IOTXClearInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXClearInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXClearInitialized represents a Initialized event raised by the IOTXClear contract.
type IOTXClearInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IOTXClear *IOTXClearFilterer) FilterInitialized(opts *bind.FilterOpts) (*IOTXClearInitializedIterator, error) {

	logs, sub, err := _IOTXClear.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IOTXClearInitializedIterator{contract: _IOTXClear.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IOTXClear *IOTXClearFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IOTXClearInitialized) (event.Subscription, error) {

	logs, sub, err := _IOTXClear.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXClearInitialized)
				if err := _IOTXClear.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IOTXClear *IOTXClearFilterer) ParseInitialized(log types.Log) (*IOTXClearInitialized, error) {
	event := new(IOTXClearInitialized)
	if err := _IOTXClear.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXClearPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IOTXClear contract.
type IOTXClearPausedIterator struct {
	Event *IOTXClearPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IOTXClearPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXClearPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IOTXClearPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IOTXClearPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXClearPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXClearPaused represents a Paused event raised by the IOTXClear contract.
type IOTXClearPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IOTXClear *IOTXClearFilterer) FilterPaused(opts *bind.FilterOpts) (*IOTXClearPausedIterator, error) {

	logs, sub, err := _IOTXClear.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IOTXClearPausedIterator{contract: _IOTXClear.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IOTXClear *IOTXClearFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IOTXClearPaused) (event.Subscription, error) {

	logs, sub, err := _IOTXClear.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXClearPaused)
				if err := _IOTXClear.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IOTXClear *IOTXClearFilterer) ParsePaused(log types.Log) (*IOTXClearPaused, error) {
	event := new(IOTXClearPaused)
	if err := _IOTXClear.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXClearPrincipalClaimedIterator is returned from FilterPrincipalClaimed and is used to iterate over the raw logs and unpacked data for PrincipalClaimed events raised by the IOTXClear contract.
type IOTXClearPrincipalClaimedIterator struct {
	Event *IOTXClearPrincipalClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IOTXClearPrincipalClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXClearPrincipalClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IOTXClearPrincipalClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IOTXClearPrincipalClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXClearPrincipalClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXClearPrincipalClaimed represents a PrincipalClaimed event raised by the IOTXClear contract.
type IOTXClearPrincipalClaimed struct {
	Claimer   common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPrincipalClaimed is a free log retrieval operation binding the contract event 0x7f8a3859c11803a8872da707dae770c23ab5b3eb3ea899d87a00b1c2bfa7749b.
//
// Solidity: event PrincipalClaimed(address claimer, address recipient, uint256 amount)
func (_IOTXClear *IOTXClearFilterer) FilterPrincipalClaimed(opts *bind.FilterOpts) (*IOTXClearPrincipalClaimedIterator, error) {

	logs, sub, err := _IOTXClear.contract.FilterLogs(opts, "PrincipalClaimed")
	if err != nil {
		return nil, err
	}
	return &IOTXClearPrincipalClaimedIterator{contract: _IOTXClear.contract, event: "PrincipalClaimed", logs: logs, sub: sub}, nil
}

// WatchPrincipalClaimed is a free log subscription operation binding the contract event 0x7f8a3859c11803a8872da707dae770c23ab5b3eb3ea899d87a00b1c2bfa7749b.
//
// Solidity: event PrincipalClaimed(address claimer, address recipient, uint256 amount)
func (_IOTXClear *IOTXClearFilterer) WatchPrincipalClaimed(opts *bind.WatchOpts, sink chan<- *IOTXClearPrincipalClaimed) (event.Subscription, error) {

	logs, sub, err := _IOTXClear.contract.WatchLogs(opts, "PrincipalClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXClearPrincipalClaimed)
				if err := _IOTXClear.contract.UnpackLog(event, "PrincipalClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePrincipalClaimed is a log parse operation binding the contract event 0x7f8a3859c11803a8872da707dae770c23ab5b3eb3ea899d87a00b1c2bfa7749b.
//
// Solidity: event PrincipalClaimed(address claimer, address recipient, uint256 amount)
func (_IOTXClear *IOTXClearFilterer) ParsePrincipalClaimed(log types.Log) (*IOTXClearPrincipalClaimed, error) {
	event := new(IOTXClearPrincipalClaimed)
	if err := _IOTXClear.contract.UnpackLog(event, "PrincipalClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXClearRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the IOTXClear contract.
type IOTXClearRewardClaimedIterator struct {
	Event *IOTXClearRewardClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IOTXClearRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXClearRewardClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IOTXClearRewardClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IOTXClearRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXClearRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXClearRewardClaimed represents a RewardClaimed event raised by the IOTXClear contract.
type IOTXClearRewardClaimed struct {
	Claimer   common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address claimer, address recipient, uint256 amount)
func (_IOTXClear *IOTXClearFilterer) FilterRewardClaimed(opts *bind.FilterOpts) (*IOTXClearRewardClaimedIterator, error) {

	logs, sub, err := _IOTXClear.contract.FilterLogs(opts, "RewardClaimed")
	if err != nil {
		return nil, err
	}
	return &IOTXClearRewardClaimedIterator{contract: _IOTXClear.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address claimer, address recipient, uint256 amount)
func (_IOTXClear *IOTXClearFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *IOTXClearRewardClaimed) (event.Subscription, error) {

	logs, sub, err := _IOTXClear.contract.WatchLogs(opts, "RewardClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXClearRewardClaimed)
				if err := _IOTXClear.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardClaimed is a log parse operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address claimer, address recipient, uint256 amount)
func (_IOTXClear *IOTXClearFilterer) ParseRewardClaimed(log types.Log) (*IOTXClearRewardClaimed, error) {
	event := new(IOTXClearRewardClaimed)
	if err := _IOTXClear.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXClearRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the IOTXClear contract.
type IOTXClearRoleAdminChangedIterator struct {
	Event *IOTXClearRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IOTXClearRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXClearRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IOTXClearRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IOTXClearRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXClearRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXClearRoleAdminChanged represents a RoleAdminChanged event raised by the IOTXClear contract.
type IOTXClearRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IOTXClear *IOTXClearFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*IOTXClearRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IOTXClear.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &IOTXClearRoleAdminChangedIterator{contract: _IOTXClear.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IOTXClear *IOTXClearFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *IOTXClearRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IOTXClear.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXClearRoleAdminChanged)
				if err := _IOTXClear.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IOTXClear *IOTXClearFilterer) ParseRoleAdminChanged(log types.Log) (*IOTXClearRoleAdminChanged, error) {
	event := new(IOTXClearRoleAdminChanged)
	if err := _IOTXClear.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXClearRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IOTXClear contract.
type IOTXClearRoleGrantedIterator struct {
	Event *IOTXClearRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IOTXClearRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXClearRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IOTXClearRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IOTXClearRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXClearRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXClearRoleGranted represents a RoleGranted event raised by the IOTXClear contract.
type IOTXClearRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IOTXClear *IOTXClearFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IOTXClearRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IOTXClear.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IOTXClearRoleGrantedIterator{contract: _IOTXClear.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IOTXClear *IOTXClearFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IOTXClearRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IOTXClear.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXClearRoleGranted)
				if err := _IOTXClear.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IOTXClear *IOTXClearFilterer) ParseRoleGranted(log types.Log) (*IOTXClearRoleGranted, error) {
	event := new(IOTXClearRoleGranted)
	if err := _IOTXClear.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXClearRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IOTXClear contract.
type IOTXClearRoleRevokedIterator struct {
	Event *IOTXClearRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IOTXClearRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXClearRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IOTXClearRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IOTXClearRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXClearRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXClearRoleRevoked represents a RoleRevoked event raised by the IOTXClear contract.
type IOTXClearRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IOTXClear *IOTXClearFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IOTXClearRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IOTXClear.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IOTXClearRoleRevokedIterator{contract: _IOTXClear.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IOTXClear *IOTXClearFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IOTXClearRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IOTXClear.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXClearRoleRevoked)
				if err := _IOTXClear.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IOTXClear *IOTXClearFilterer) ParseRoleRevoked(log types.Log) (*IOTXClearRoleRevoked, error) {
	event := new(IOTXClearRoleRevoked)
	if err := _IOTXClear.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXClearUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IOTXClear contract.
type IOTXClearUnpausedIterator struct {
	Event *IOTXClearUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IOTXClearUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXClearUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IOTXClearUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IOTXClearUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXClearUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXClearUnpaused represents a Unpaused event raised by the IOTXClear contract.
type IOTXClearUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IOTXClear *IOTXClearFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IOTXClearUnpausedIterator, error) {

	logs, sub, err := _IOTXClear.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IOTXClearUnpausedIterator{contract: _IOTXClear.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IOTXClear *IOTXClearFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IOTXClearUnpaused) (event.Subscription, error) {

	logs, sub, err := _IOTXClear.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXClearUnpaused)
				if err := _IOTXClear.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IOTXClear *IOTXClearFilterer) ParseUnpaused(log types.Log) (*IOTXClearUnpaused, error) {
	event := new(IOTXClearUnpaused)
	if err := _IOTXClear.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
