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

// IOTXStakeMetaData contains all meta data concerning the IOTXStake contract.
var IOTXStakeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toPayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"queuedAmount\",\"type\":\"uint256\"}],\"name\":\"DebtAmountMismatched\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"InsufficientReward\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DebtAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DebtPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_systemStakeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_iotxStakeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracleAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"iotxDebts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"joinDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemStake\",\"outputs\":[{\"internalType\":\"contractISystemStake\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDebts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"updateDelegates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"updateReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IOTXStakeABI is the input ABI used to generate the binding from.
// Deprecated: Use IOTXStakeMetaData.ABI instead.
var IOTXStakeABI = IOTXStakeMetaData.ABI

// IOTXStake is an auto generated Go binding around an Ethereum contract.
type IOTXStake struct {
	IOTXStakeCaller     // Read-only binding to the contract
	IOTXStakeTransactor // Write-only binding to the contract
	IOTXStakeFilterer   // Log filterer for contract events
}

// IOTXStakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOTXStakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOTXStakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOTXStakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOTXStakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOTXStakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOTXStakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOTXStakeSession struct {
	Contract     *IOTXStake        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOTXStakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOTXStakeCallerSession struct {
	Contract *IOTXStakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IOTXStakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOTXStakeTransactorSession struct {
	Contract     *IOTXStakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IOTXStakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOTXStakeRaw struct {
	Contract *IOTXStake // Generic contract binding to access the raw methods on
}

// IOTXStakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOTXStakeCallerRaw struct {
	Contract *IOTXStakeCaller // Generic read-only contract binding to access the raw methods on
}

// IOTXStakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOTXStakeTransactorRaw struct {
	Contract *IOTXStakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOTXStake creates a new instance of IOTXStake, bound to a specific deployed contract.
func NewIOTXStake(address common.Address, backend bind.ContractBackend) (*IOTXStake, error) {
	contract, err := bindIOTXStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOTXStake{IOTXStakeCaller: IOTXStakeCaller{contract: contract}, IOTXStakeTransactor: IOTXStakeTransactor{contract: contract}, IOTXStakeFilterer: IOTXStakeFilterer{contract: contract}}, nil
}

// NewIOTXStakeCaller creates a new read-only instance of IOTXStake, bound to a specific deployed contract.
func NewIOTXStakeCaller(address common.Address, caller bind.ContractCaller) (*IOTXStakeCaller, error) {
	contract, err := bindIOTXStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOTXStakeCaller{contract: contract}, nil
}

// NewIOTXStakeTransactor creates a new write-only instance of IOTXStake, bound to a specific deployed contract.
func NewIOTXStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*IOTXStakeTransactor, error) {
	contract, err := bindIOTXStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOTXStakeTransactor{contract: contract}, nil
}

// NewIOTXStakeFilterer creates a new log filterer instance of IOTXStake, bound to a specific deployed contract.
func NewIOTXStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*IOTXStakeFilterer, error) {
	contract, err := bindIOTXStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOTXStakeFilterer{contract: contract}, nil
}

// bindIOTXStake binds a generic wrapper to an already deployed contract.
func bindIOTXStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOTXStakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOTXStake *IOTXStakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOTXStake.Contract.IOTXStakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOTXStake *IOTXStakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXStake.Contract.IOTXStakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOTXStake *IOTXStakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOTXStake.Contract.IOTXStakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOTXStake *IOTXStakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOTXStake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOTXStake *IOTXStakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXStake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOTXStake *IOTXStakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOTXStake.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IOTXStake *IOTXStakeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IOTXStake *IOTXStakeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IOTXStake.Contract.DEFAULTADMINROLE(&_IOTXStake.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IOTXStake *IOTXStakeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IOTXStake.Contract.DEFAULTADMINROLE(&_IOTXStake.CallOpts)
}

// MULTIPLIER is a free data retrieval call binding the contract method 0x059f8b16.
//
// Solidity: function MULTIPLIER() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) MULTIPLIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "MULTIPLIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIPLIER is a free data retrieval call binding the contract method 0x059f8b16.
//
// Solidity: function MULTIPLIER() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) MULTIPLIER() (*big.Int, error) {
	return _IOTXStake.Contract.MULTIPLIER(&_IOTXStake.CallOpts)
}

// MULTIPLIER is a free data retrieval call binding the contract method 0x059f8b16.
//
// Solidity: function MULTIPLIER() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) MULTIPLIER() (*big.Int, error) {
	return _IOTXStake.Contract.MULTIPLIER(&_IOTXStake.CallOpts)
}

// AccountedBalance is a free data retrieval call binding the contract method 0x0937eb54.
//
// Solidity: function accountedBalance() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) AccountedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "accountedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountedBalance is a free data retrieval call binding the contract method 0x0937eb54.
//
// Solidity: function accountedBalance() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) AccountedBalance() (*big.Int, error) {
	return _IOTXStake.Contract.AccountedBalance(&_IOTXStake.CallOpts)
}

// AccountedBalance is a free data retrieval call binding the contract method 0x0937eb54.
//
// Solidity: function accountedBalance() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) AccountedBalance() (*big.Int, error) {
	return _IOTXStake.Contract.AccountedBalance(&_IOTXStake.CallOpts)
}

// FirstIndex is a free data retrieval call binding the contract method 0x6b4f2c0f.
//
// Solidity: function firstIndex() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) FirstIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "firstIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstIndex is a free data retrieval call binding the contract method 0x6b4f2c0f.
//
// Solidity: function firstIndex() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) FirstIndex() (*big.Int, error) {
	return _IOTXStake.Contract.FirstIndex(&_IOTXStake.CallOpts)
}

// FirstIndex is a free data retrieval call binding the contract method 0x6b4f2c0f.
//
// Solidity: function firstIndex() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) FirstIndex() (*big.Int, error) {
	return _IOTXStake.Contract.FirstIndex(&_IOTXStake.CallOpts)
}

// GetReward is a free data retrieval call binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address account) view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) GetReward(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "getReward", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReward is a free data retrieval call binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address account) view returns(uint256)
func (_IOTXStake *IOTXStakeSession) GetReward(account common.Address) (*big.Int, error) {
	return _IOTXStake.Contract.GetReward(&_IOTXStake.CallOpts, account)
}

// GetReward is a free data retrieval call binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address account) view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) GetReward(account common.Address) (*big.Int, error) {
	return _IOTXStake.Contract.GetReward(&_IOTXStake.CallOpts, account)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IOTXStake *IOTXStakeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IOTXStake *IOTXStakeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IOTXStake.Contract.GetRoleAdmin(&_IOTXStake.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IOTXStake *IOTXStakeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IOTXStake.Contract.GetRoleAdmin(&_IOTXStake.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IOTXStake *IOTXStakeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IOTXStake *IOTXStakeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IOTXStake.Contract.HasRole(&_IOTXStake.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IOTXStake *IOTXStakeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IOTXStake.Contract.HasRole(&_IOTXStake.CallOpts, role, account)
}

// IotxDebts is a free data retrieval call binding the contract method 0x5c14675e.
//
// Solidity: function iotxDebts(uint256 ) view returns(address account, uint256 amount)
func (_IOTXStake *IOTXStakeCaller) IotxDebts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Account common.Address
	Amount  *big.Int
}, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "iotxDebts", arg0)

	outstruct := new(struct {
		Account common.Address
		Amount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Account = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// IotxDebts is a free data retrieval call binding the contract method 0x5c14675e.
//
// Solidity: function iotxDebts(uint256 ) view returns(address account, uint256 amount)
func (_IOTXStake *IOTXStakeSession) IotxDebts(arg0 *big.Int) (struct {
	Account common.Address
	Amount  *big.Int
}, error) {
	return _IOTXStake.Contract.IotxDebts(&_IOTXStake.CallOpts, arg0)
}

// IotxDebts is a free data retrieval call binding the contract method 0x5c14675e.
//
// Solidity: function iotxDebts(uint256 ) view returns(address account, uint256 amount)
func (_IOTXStake *IOTXStakeCallerSession) IotxDebts(arg0 *big.Int) (struct {
	Account common.Address
	Amount  *big.Int
}, error) {
	return _IOTXStake.Contract.IotxDebts(&_IOTXStake.CallOpts, arg0)
}

// LastIndex is a free data retrieval call binding the contract method 0xf3f6f0d7.
//
// Solidity: function lastIndex() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) LastIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "lastIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastIndex is a free data retrieval call binding the contract method 0xf3f6f0d7.
//
// Solidity: function lastIndex() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) LastIndex() (*big.Int, error) {
	return _IOTXStake.Contract.LastIndex(&_IOTXStake.CallOpts)
}

// LastIndex is a free data retrieval call binding the contract method 0xf3f6f0d7.
//
// Solidity: function lastIndex() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) LastIndex() (*big.Int, error) {
	return _IOTXStake.Contract.LastIndex(&_IOTXStake.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_IOTXStake *IOTXStakeCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_IOTXStake *IOTXStakeSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _IOTXStake.Contract.OnERC721Received(&_IOTXStake.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_IOTXStake *IOTXStakeCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _IOTXStake.Contract.OnERC721Received(&_IOTXStake.CallOpts, arg0, arg1, arg2, arg3)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IOTXStake *IOTXStakeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IOTXStake *IOTXStakeSession) Paused() (bool, error) {
	return _IOTXStake.Contract.Paused(&_IOTXStake.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IOTXStake *IOTXStakeCallerSession) Paused() (bool, error) {
	return _IOTXStake.Contract.Paused(&_IOTXStake.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) RewardRate() (*big.Int, error) {
	return _IOTXStake.Contract.RewardRate(&_IOTXStake.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) RewardRate() (*big.Int, error) {
	return _IOTXStake.Contract.RewardRate(&_IOTXStake.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IOTXStake *IOTXStakeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IOTXStake *IOTXStakeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IOTXStake.Contract.SupportsInterface(&_IOTXStake.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IOTXStake *IOTXStakeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IOTXStake.Contract.SupportsInterface(&_IOTXStake.CallOpts, interfaceId)
}

// SystemStake is a free data retrieval call binding the contract method 0xcae8406b.
//
// Solidity: function systemStake() view returns(address)
func (_IOTXStake *IOTXStakeCaller) SystemStake(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "systemStake")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemStake is a free data retrieval call binding the contract method 0xcae8406b.
//
// Solidity: function systemStake() view returns(address)
func (_IOTXStake *IOTXStakeSession) SystemStake() (common.Address, error) {
	return _IOTXStake.Contract.SystemStake(&_IOTXStake.CallOpts)
}

// SystemStake is a free data retrieval call binding the contract method 0xcae8406b.
//
// Solidity: function systemStake() view returns(address)
func (_IOTXStake *IOTXStakeCallerSession) SystemStake() (common.Address, error) {
	return _IOTXStake.Contract.SystemStake(&_IOTXStake.CallOpts)
}

// TotalDebts is a free data retrieval call binding the contract method 0x14a1c32d.
//
// Solidity: function totalDebts() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) TotalDebts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "totalDebts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebts is a free data retrieval call binding the contract method 0x14a1c32d.
//
// Solidity: function totalDebts() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) TotalDebts() (*big.Int, error) {
	return _IOTXStake.Contract.TotalDebts(&_IOTXStake.CallOpts)
}

// TotalDebts is a free data retrieval call binding the contract method 0x14a1c32d.
//
// Solidity: function totalDebts() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) TotalDebts() (*big.Int, error) {
	return _IOTXStake.Contract.TotalDebts(&_IOTXStake.CallOpts)
}

// UserInfos is a free data retrieval call binding the contract method 0x43b0215f.
//
// Solidity: function userInfos(address ) view returns(uint256 debt, uint256 reward, uint256 rewardRate)
func (_IOTXStake *IOTXStakeCaller) UserInfos(opts *bind.CallOpts, arg0 common.Address) (struct {
	Debt       *big.Int
	Reward     *big.Int
	RewardRate *big.Int
}, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "userInfos", arg0)

	outstruct := new(struct {
		Debt       *big.Int
		Reward     *big.Int
		RewardRate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Debt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RewardRate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfos is a free data retrieval call binding the contract method 0x43b0215f.
//
// Solidity: function userInfos(address ) view returns(uint256 debt, uint256 reward, uint256 rewardRate)
func (_IOTXStake *IOTXStakeSession) UserInfos(arg0 common.Address) (struct {
	Debt       *big.Int
	Reward     *big.Int
	RewardRate *big.Int
}, error) {
	return _IOTXStake.Contract.UserInfos(&_IOTXStake.CallOpts, arg0)
}

// UserInfos is a free data retrieval call binding the contract method 0x43b0215f.
//
// Solidity: function userInfos(address ) view returns(uint256 debt, uint256 reward, uint256 rewardRate)
func (_IOTXStake *IOTXStakeCallerSession) UserInfos(arg0 common.Address) (struct {
	Debt       *big.Int
	Reward     *big.Int
	RewardRate *big.Int
}, error) {
	return _IOTXStake.Contract.UserInfos(&_IOTXStake.CallOpts, arg0)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6c7b69cb.
//
// Solidity: function claimRewards(uint256 amount, address recipient) returns()
func (_IOTXStake *IOTXStakeTransactor) ClaimRewards(opts *bind.TransactOpts, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "claimRewards", amount, recipient)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6c7b69cb.
//
// Solidity: function claimRewards(uint256 amount, address recipient) returns()
func (_IOTXStake *IOTXStakeSession) ClaimRewards(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IOTXStake.Contract.ClaimRewards(&_IOTXStake.TransactOpts, amount, recipient)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6c7b69cb.
//
// Solidity: function claimRewards(uint256 amount, address recipient) returns()
func (_IOTXStake *IOTXStakeTransactorSession) ClaimRewards(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IOTXStake.Contract.ClaimRewards(&_IOTXStake.TransactOpts, amount, recipient)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IOTXStake *IOTXStakeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IOTXStake *IOTXStakeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXStake.Contract.GrantRole(&_IOTXStake.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IOTXStake *IOTXStakeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXStake.Contract.GrantRole(&_IOTXStake.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _systemStakeAddress, address _iotxStakeAddress, address _oracleAddress) returns()
func (_IOTXStake *IOTXStakeTransactor) Initialize(opts *bind.TransactOpts, _systemStakeAddress common.Address, _iotxStakeAddress common.Address, _oracleAddress common.Address) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "initialize", _systemStakeAddress, _iotxStakeAddress, _oracleAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _systemStakeAddress, address _iotxStakeAddress, address _oracleAddress) returns()
func (_IOTXStake *IOTXStakeSession) Initialize(_systemStakeAddress common.Address, _iotxStakeAddress common.Address, _oracleAddress common.Address) (*types.Transaction, error) {
	return _IOTXStake.Contract.Initialize(&_IOTXStake.TransactOpts, _systemStakeAddress, _iotxStakeAddress, _oracleAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _systemStakeAddress, address _iotxStakeAddress, address _oracleAddress) returns()
func (_IOTXStake *IOTXStakeTransactorSession) Initialize(_systemStakeAddress common.Address, _iotxStakeAddress common.Address, _oracleAddress common.Address) (*types.Transaction, error) {
	return _IOTXStake.Contract.Initialize(&_IOTXStake.TransactOpts, _systemStakeAddress, _iotxStakeAddress, _oracleAddress)
}

// JoinDebt is a paid mutator transaction binding the contract method 0x7180f85b.
//
// Solidity: function joinDebt(address account, uint256 amount) returns()
func (_IOTXStake *IOTXStakeTransactor) JoinDebt(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "joinDebt", account, amount)
}

// JoinDebt is a paid mutator transaction binding the contract method 0x7180f85b.
//
// Solidity: function joinDebt(address account, uint256 amount) returns()
func (_IOTXStake *IOTXStakeSession) JoinDebt(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IOTXStake.Contract.JoinDebt(&_IOTXStake.TransactOpts, account, amount)
}

// JoinDebt is a paid mutator transaction binding the contract method 0x7180f85b.
//
// Solidity: function joinDebt(address account, uint256 amount) returns()
func (_IOTXStake *IOTXStakeTransactorSession) JoinDebt(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IOTXStake.Contract.JoinDebt(&_IOTXStake.TransactOpts, account, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IOTXStake *IOTXStakeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IOTXStake *IOTXStakeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXStake.Contract.RenounceRole(&_IOTXStake.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IOTXStake *IOTXStakeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXStake.Contract.RenounceRole(&_IOTXStake.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IOTXStake *IOTXStakeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IOTXStake *IOTXStakeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXStake.Contract.RevokeRole(&_IOTXStake.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IOTXStake *IOTXStakeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXStake.Contract.RevokeRole(&_IOTXStake.TransactOpts, role, account)
}

// Unstake is a paid mutator transaction binding the contract method 0xe449f341.
//
// Solidity: function unstake(uint256[] tokenIds) returns()
func (_IOTXStake *IOTXStakeTransactor) Unstake(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "unstake", tokenIds)
}

// Unstake is a paid mutator transaction binding the contract method 0xe449f341.
//
// Solidity: function unstake(uint256[] tokenIds) returns()
func (_IOTXStake *IOTXStakeSession) Unstake(tokenIds []*big.Int) (*types.Transaction, error) {
	return _IOTXStake.Contract.Unstake(&_IOTXStake.TransactOpts, tokenIds)
}

// Unstake is a paid mutator transaction binding the contract method 0xe449f341.
//
// Solidity: function unstake(uint256[] tokenIds) returns()
func (_IOTXStake *IOTXStakeTransactorSession) Unstake(tokenIds []*big.Int) (*types.Transaction, error) {
	return _IOTXStake.Contract.Unstake(&_IOTXStake.TransactOpts, tokenIds)
}

// UpdateDelegates is a paid mutator transaction binding the contract method 0x2328782f.
//
// Solidity: function updateDelegates(uint256[] tokenIds, address delegate) returns()
func (_IOTXStake *IOTXStakeTransactor) UpdateDelegates(opts *bind.TransactOpts, tokenIds []*big.Int, delegate common.Address) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "updateDelegates", tokenIds, delegate)
}

// UpdateDelegates is a paid mutator transaction binding the contract method 0x2328782f.
//
// Solidity: function updateDelegates(uint256[] tokenIds, address delegate) returns()
func (_IOTXStake *IOTXStakeSession) UpdateDelegates(tokenIds []*big.Int, delegate common.Address) (*types.Transaction, error) {
	return _IOTXStake.Contract.UpdateDelegates(&_IOTXStake.TransactOpts, tokenIds, delegate)
}

// UpdateDelegates is a paid mutator transaction binding the contract method 0x2328782f.
//
// Solidity: function updateDelegates(uint256[] tokenIds, address delegate) returns()
func (_IOTXStake *IOTXStakeTransactorSession) UpdateDelegates(tokenIds []*big.Int, delegate common.Address) (*types.Transaction, error) {
	return _IOTXStake.Contract.UpdateDelegates(&_IOTXStake.TransactOpts, tokenIds, delegate)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x632447c9.
//
// Solidity: function updateReward(address account) returns(uint256)
func (_IOTXStake *IOTXStakeTransactor) UpdateReward(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "updateReward", account)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x632447c9.
//
// Solidity: function updateReward(address account) returns(uint256)
func (_IOTXStake *IOTXStakeSession) UpdateReward(account common.Address) (*types.Transaction, error) {
	return _IOTXStake.Contract.UpdateReward(&_IOTXStake.TransactOpts, account)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x632447c9.
//
// Solidity: function updateReward(address account) returns(uint256)
func (_IOTXStake *IOTXStakeTransactorSession) UpdateReward(account common.Address) (*types.Transaction, error) {
	return _IOTXStake.Contract.UpdateReward(&_IOTXStake.TransactOpts, account)
}

// Withdraw is a paid mutator transaction binding the contract method 0x983d95ce.
//
// Solidity: function withdraw(uint256[] tokenIds) returns()
func (_IOTXStake *IOTXStakeTransactor) Withdraw(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "withdraw", tokenIds)
}

// Withdraw is a paid mutator transaction binding the contract method 0x983d95ce.
//
// Solidity: function withdraw(uint256[] tokenIds) returns()
func (_IOTXStake *IOTXStakeSession) Withdraw(tokenIds []*big.Int) (*types.Transaction, error) {
	return _IOTXStake.Contract.Withdraw(&_IOTXStake.TransactOpts, tokenIds)
}

// Withdraw is a paid mutator transaction binding the contract method 0x983d95ce.
//
// Solidity: function withdraw(uint256[] tokenIds) returns()
func (_IOTXStake *IOTXStakeTransactorSession) Withdraw(tokenIds []*big.Int) (*types.Transaction, error) {
	return _IOTXStake.Contract.Withdraw(&_IOTXStake.TransactOpts, tokenIds)
}

// IOTXStakeDebtAddedIterator is returned from FilterDebtAdded and is used to iterate over the raw logs and unpacked data for DebtAdded events raised by the IOTXStake contract.
type IOTXStakeDebtAddedIterator struct {
	Event *IOTXStakeDebtAdded // Event containing the contract specifics and raw log

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
func (it *IOTXStakeDebtAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakeDebtAdded)
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
		it.Event = new(IOTXStakeDebtAdded)
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
func (it *IOTXStakeDebtAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakeDebtAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakeDebtAdded represents a DebtAdded event raised by the IOTXStake contract.
type IOTXStakeDebtAdded struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDebtAdded is a free log retrieval operation binding the contract event 0x183076a1e61b553fbb03b0fb4641f565bba910da3e423f6931f9d08a16539692.
//
// Solidity: event DebtAdded(address account, uint256 amount)
func (_IOTXStake *IOTXStakeFilterer) FilterDebtAdded(opts *bind.FilterOpts) (*IOTXStakeDebtAddedIterator, error) {

	logs, sub, err := _IOTXStake.contract.FilterLogs(opts, "DebtAdded")
	if err != nil {
		return nil, err
	}
	return &IOTXStakeDebtAddedIterator{contract: _IOTXStake.contract, event: "DebtAdded", logs: logs, sub: sub}, nil
}

// WatchDebtAdded is a free log subscription operation binding the contract event 0x183076a1e61b553fbb03b0fb4641f565bba910da3e423f6931f9d08a16539692.
//
// Solidity: event DebtAdded(address account, uint256 amount)
func (_IOTXStake *IOTXStakeFilterer) WatchDebtAdded(opts *bind.WatchOpts, sink chan<- *IOTXStakeDebtAdded) (event.Subscription, error) {

	logs, sub, err := _IOTXStake.contract.WatchLogs(opts, "DebtAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakeDebtAdded)
				if err := _IOTXStake.contract.UnpackLog(event, "DebtAdded", log); err != nil {
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

// ParseDebtAdded is a log parse operation binding the contract event 0x183076a1e61b553fbb03b0fb4641f565bba910da3e423f6931f9d08a16539692.
//
// Solidity: event DebtAdded(address account, uint256 amount)
func (_IOTXStake *IOTXStakeFilterer) ParseDebtAdded(log types.Log) (*IOTXStakeDebtAdded, error) {
	event := new(IOTXStakeDebtAdded)
	if err := _IOTXStake.contract.UnpackLog(event, "DebtAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakeDebtPaidIterator is returned from FilterDebtPaid and is used to iterate over the raw logs and unpacked data for DebtPaid events raised by the IOTXStake contract.
type IOTXStakeDebtPaidIterator struct {
	Event *IOTXStakeDebtPaid // Event containing the contract specifics and raw log

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
func (it *IOTXStakeDebtPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakeDebtPaid)
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
		it.Event = new(IOTXStakeDebtPaid)
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
func (it *IOTXStakeDebtPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakeDebtPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakeDebtPaid represents a DebtPaid event raised by the IOTXStake contract.
type IOTXStakeDebtPaid struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDebtPaid is a free log retrieval operation binding the contract event 0xc75e95ca1d8e8b97e6452ce2acc93a4f8f9f5779fbced06e43809b0d4e4e3ce4.
//
// Solidity: event DebtPaid(address account, uint256 amount)
func (_IOTXStake *IOTXStakeFilterer) FilterDebtPaid(opts *bind.FilterOpts) (*IOTXStakeDebtPaidIterator, error) {

	logs, sub, err := _IOTXStake.contract.FilterLogs(opts, "DebtPaid")
	if err != nil {
		return nil, err
	}
	return &IOTXStakeDebtPaidIterator{contract: _IOTXStake.contract, event: "DebtPaid", logs: logs, sub: sub}, nil
}

// WatchDebtPaid is a free log subscription operation binding the contract event 0xc75e95ca1d8e8b97e6452ce2acc93a4f8f9f5779fbced06e43809b0d4e4e3ce4.
//
// Solidity: event DebtPaid(address account, uint256 amount)
func (_IOTXStake *IOTXStakeFilterer) WatchDebtPaid(opts *bind.WatchOpts, sink chan<- *IOTXStakeDebtPaid) (event.Subscription, error) {

	logs, sub, err := _IOTXStake.contract.WatchLogs(opts, "DebtPaid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakeDebtPaid)
				if err := _IOTXStake.contract.UnpackLog(event, "DebtPaid", log); err != nil {
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

// ParseDebtPaid is a log parse operation binding the contract event 0xc75e95ca1d8e8b97e6452ce2acc93a4f8f9f5779fbced06e43809b0d4e4e3ce4.
//
// Solidity: event DebtPaid(address account, uint256 amount)
func (_IOTXStake *IOTXStakeFilterer) ParseDebtPaid(log types.Log) (*IOTXStakeDebtPaid, error) {
	event := new(IOTXStakeDebtPaid)
	if err := _IOTXStake.contract.UnpackLog(event, "DebtPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IOTXStake contract.
type IOTXStakeInitializedIterator struct {
	Event *IOTXStakeInitialized // Event containing the contract specifics and raw log

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
func (it *IOTXStakeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakeInitialized)
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
		it.Event = new(IOTXStakeInitialized)
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
func (it *IOTXStakeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakeInitialized represents a Initialized event raised by the IOTXStake contract.
type IOTXStakeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IOTXStake *IOTXStakeFilterer) FilterInitialized(opts *bind.FilterOpts) (*IOTXStakeInitializedIterator, error) {

	logs, sub, err := _IOTXStake.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IOTXStakeInitializedIterator{contract: _IOTXStake.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IOTXStake *IOTXStakeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IOTXStakeInitialized) (event.Subscription, error) {

	logs, sub, err := _IOTXStake.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakeInitialized)
				if err := _IOTXStake.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_IOTXStake *IOTXStakeFilterer) ParseInitialized(log types.Log) (*IOTXStakeInitialized, error) {
	event := new(IOTXStakeInitialized)
	if err := _IOTXStake.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IOTXStake contract.
type IOTXStakePausedIterator struct {
	Event *IOTXStakePaused // Event containing the contract specifics and raw log

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
func (it *IOTXStakePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakePaused)
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
		it.Event = new(IOTXStakePaused)
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
func (it *IOTXStakePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakePaused represents a Paused event raised by the IOTXStake contract.
type IOTXStakePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IOTXStake *IOTXStakeFilterer) FilterPaused(opts *bind.FilterOpts) (*IOTXStakePausedIterator, error) {

	logs, sub, err := _IOTXStake.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IOTXStakePausedIterator{contract: _IOTXStake.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IOTXStake *IOTXStakeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IOTXStakePaused) (event.Subscription, error) {

	logs, sub, err := _IOTXStake.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakePaused)
				if err := _IOTXStake.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_IOTXStake *IOTXStakeFilterer) ParsePaused(log types.Log) (*IOTXStakePaused, error) {
	event := new(IOTXStakePaused)
	if err := _IOTXStake.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakeRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the IOTXStake contract.
type IOTXStakeRewardClaimedIterator struct {
	Event *IOTXStakeRewardClaimed // Event containing the contract specifics and raw log

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
func (it *IOTXStakeRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakeRewardClaimed)
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
		it.Event = new(IOTXStakeRewardClaimed)
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
func (it *IOTXStakeRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakeRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakeRewardClaimed represents a RewardClaimed event raised by the IOTXStake contract.
type IOTXStakeRewardClaimed struct {
	Claimer   common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address claimer, address recipient, uint256 amount)
func (_IOTXStake *IOTXStakeFilterer) FilterRewardClaimed(opts *bind.FilterOpts) (*IOTXStakeRewardClaimedIterator, error) {

	logs, sub, err := _IOTXStake.contract.FilterLogs(opts, "RewardClaimed")
	if err != nil {
		return nil, err
	}
	return &IOTXStakeRewardClaimedIterator{contract: _IOTXStake.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address claimer, address recipient, uint256 amount)
func (_IOTXStake *IOTXStakeFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *IOTXStakeRewardClaimed) (event.Subscription, error) {

	logs, sub, err := _IOTXStake.contract.WatchLogs(opts, "RewardClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakeRewardClaimed)
				if err := _IOTXStake.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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
func (_IOTXStake *IOTXStakeFilterer) ParseRewardClaimed(log types.Log) (*IOTXStakeRewardClaimed, error) {
	event := new(IOTXStakeRewardClaimed)
	if err := _IOTXStake.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the IOTXStake contract.
type IOTXStakeRoleAdminChangedIterator struct {
	Event *IOTXStakeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *IOTXStakeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakeRoleAdminChanged)
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
		it.Event = new(IOTXStakeRoleAdminChanged)
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
func (it *IOTXStakeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakeRoleAdminChanged represents a RoleAdminChanged event raised by the IOTXStake contract.
type IOTXStakeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IOTXStake *IOTXStakeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*IOTXStakeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _IOTXStake.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &IOTXStakeRoleAdminChangedIterator{contract: _IOTXStake.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IOTXStake *IOTXStakeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *IOTXStakeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _IOTXStake.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakeRoleAdminChanged)
				if err := _IOTXStake.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_IOTXStake *IOTXStakeFilterer) ParseRoleAdminChanged(log types.Log) (*IOTXStakeRoleAdminChanged, error) {
	event := new(IOTXStakeRoleAdminChanged)
	if err := _IOTXStake.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IOTXStake contract.
type IOTXStakeRoleGrantedIterator struct {
	Event *IOTXStakeRoleGranted // Event containing the contract specifics and raw log

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
func (it *IOTXStakeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakeRoleGranted)
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
		it.Event = new(IOTXStakeRoleGranted)
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
func (it *IOTXStakeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakeRoleGranted represents a RoleGranted event raised by the IOTXStake contract.
type IOTXStakeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IOTXStake *IOTXStakeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IOTXStakeRoleGrantedIterator, error) {

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

	logs, sub, err := _IOTXStake.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IOTXStakeRoleGrantedIterator{contract: _IOTXStake.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IOTXStake *IOTXStakeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IOTXStakeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IOTXStake.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakeRoleGranted)
				if err := _IOTXStake.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_IOTXStake *IOTXStakeFilterer) ParseRoleGranted(log types.Log) (*IOTXStakeRoleGranted, error) {
	event := new(IOTXStakeRoleGranted)
	if err := _IOTXStake.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IOTXStake contract.
type IOTXStakeRoleRevokedIterator struct {
	Event *IOTXStakeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *IOTXStakeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakeRoleRevoked)
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
		it.Event = new(IOTXStakeRoleRevoked)
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
func (it *IOTXStakeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakeRoleRevoked represents a RoleRevoked event raised by the IOTXStake contract.
type IOTXStakeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IOTXStake *IOTXStakeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IOTXStakeRoleRevokedIterator, error) {

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

	logs, sub, err := _IOTXStake.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IOTXStakeRoleRevokedIterator{contract: _IOTXStake.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IOTXStake *IOTXStakeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IOTXStakeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IOTXStake.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakeRoleRevoked)
				if err := _IOTXStake.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_IOTXStake *IOTXStakeFilterer) ParseRoleRevoked(log types.Log) (*IOTXStakeRoleRevoked, error) {
	event := new(IOTXStakeRoleRevoked)
	if err := _IOTXStake.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IOTXStake contract.
type IOTXStakeUnpausedIterator struct {
	Event *IOTXStakeUnpaused // Event containing the contract specifics and raw log

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
func (it *IOTXStakeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakeUnpaused)
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
		it.Event = new(IOTXStakeUnpaused)
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
func (it *IOTXStakeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakeUnpaused represents a Unpaused event raised by the IOTXStake contract.
type IOTXStakeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IOTXStake *IOTXStakeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IOTXStakeUnpausedIterator, error) {

	logs, sub, err := _IOTXStake.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IOTXStakeUnpausedIterator{contract: _IOTXStake.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IOTXStake *IOTXStakeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IOTXStakeUnpaused) (event.Subscription, error) {

	logs, sub, err := _IOTXStake.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakeUnpaused)
				if err := _IOTXStake.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_IOTXStake *IOTXStakeFilterer) ParseUnpaused(log types.Log) (*IOTXStakeUnpaused, error) {
	event := new(IOTXStakeUnpaused)
	if err := _IOTXStake.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
