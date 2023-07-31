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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"DelegatesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"GlobalDelegateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"ManagerFeeSharesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ManagerFeeWithdrawed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Merged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"Redeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"firstTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_EXCHANGE_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountedManagerReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountedUserReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commonRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minToMint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"j\",\"type\":\"uint256\"}],\"name\":\"getRedeemedTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenQueueIndex\",\"type\":\"uint256\"}],\"name\":\"getStakedTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalDelegate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_systemStake\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_uniIOTX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_iotxClear\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_commonRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sequenceLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakeDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"iotxClear\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerFeeShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"iotxsToRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxToBurn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"burned\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemAmountBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemedTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequenceLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"setGlobalDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"setManagerFeeShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemStake\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenQueues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniIOTX\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"updateDelegates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawManagerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// DEFAULTEXCHANGERATIO is a free data retrieval call binding the contract method 0x1bb42235.
//
// Solidity: function DEFAULT_EXCHANGE_RATIO() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) DEFAULTEXCHANGERATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "DEFAULT_EXCHANGE_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTEXCHANGERATIO is a free data retrieval call binding the contract method 0x1bb42235.
//
// Solidity: function DEFAULT_EXCHANGE_RATIO() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) DEFAULTEXCHANGERATIO() (*big.Int, error) {
	return _IOTXStake.Contract.DEFAULTEXCHANGERATIO(&_IOTXStake.CallOpts)
}

// DEFAULTEXCHANGERATIO is a free data retrieval call binding the contract method 0x1bb42235.
//
// Solidity: function DEFAULT_EXCHANGE_RATIO() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) DEFAULTEXCHANGERATIO() (*big.Int, error) {
	return _IOTXStake.Contract.DEFAULTEXCHANGERATIO(&_IOTXStake.CallOpts)
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

// AccountedManagerReward is a free data retrieval call binding the contract method 0x48c26fce.
//
// Solidity: function accountedManagerReward() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) AccountedManagerReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "accountedManagerReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountedManagerReward is a free data retrieval call binding the contract method 0x48c26fce.
//
// Solidity: function accountedManagerReward() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) AccountedManagerReward() (*big.Int, error) {
	return _IOTXStake.Contract.AccountedManagerReward(&_IOTXStake.CallOpts)
}

// AccountedManagerReward is a free data retrieval call binding the contract method 0x48c26fce.
//
// Solidity: function accountedManagerReward() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) AccountedManagerReward() (*big.Int, error) {
	return _IOTXStake.Contract.AccountedManagerReward(&_IOTXStake.CallOpts)
}

// AccountedUserReward is a free data retrieval call binding the contract method 0x5ff35b5c.
//
// Solidity: function accountedUserReward() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) AccountedUserReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "accountedUserReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountedUserReward is a free data retrieval call binding the contract method 0x5ff35b5c.
//
// Solidity: function accountedUserReward() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) AccountedUserReward() (*big.Int, error) {
	return _IOTXStake.Contract.AccountedUserReward(&_IOTXStake.CallOpts)
}

// AccountedUserReward is a free data retrieval call binding the contract method 0x5ff35b5c.
//
// Solidity: function accountedUserReward() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) AccountedUserReward() (*big.Int, error) {
	return _IOTXStake.Contract.AccountedUserReward(&_IOTXStake.CallOpts)
}

// CommonRatio is a free data retrieval call binding the contract method 0x3dd419d0.
//
// Solidity: function commonRatio() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) CommonRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "commonRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommonRatio is a free data retrieval call binding the contract method 0x3dd419d0.
//
// Solidity: function commonRatio() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) CommonRatio() (*big.Int, error) {
	return _IOTXStake.Contract.CommonRatio(&_IOTXStake.CallOpts)
}

// CommonRatio is a free data retrieval call binding the contract method 0x3dd419d0.
//
// Solidity: function commonRatio() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) CommonRatio() (*big.Int, error) {
	return _IOTXStake.Contract.CommonRatio(&_IOTXStake.CallOpts)
}

// CurrentReserve is a free data retrieval call binding the contract method 0x2e12007c.
//
// Solidity: function currentReserve() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) CurrentReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "currentReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentReserve is a free data retrieval call binding the contract method 0x2e12007c.
//
// Solidity: function currentReserve() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) CurrentReserve() (*big.Int, error) {
	return _IOTXStake.Contract.CurrentReserve(&_IOTXStake.CallOpts)
}

// CurrentReserve is a free data retrieval call binding the contract method 0x2e12007c.
//
// Solidity: function currentReserve() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) CurrentReserve() (*big.Int, error) {
	return _IOTXStake.Contract.CurrentReserve(&_IOTXStake.CallOpts)
}

// ExchangeRatio is a free data retrieval call binding the contract method 0x4006ccc5.
//
// Solidity: function exchangeRatio() view returns(uint256 ratio)
func (_IOTXStake *IOTXStakeCaller) ExchangeRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "exchangeRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRatio is a free data retrieval call binding the contract method 0x4006ccc5.
//
// Solidity: function exchangeRatio() view returns(uint256 ratio)
func (_IOTXStake *IOTXStakeSession) ExchangeRatio() (*big.Int, error) {
	return _IOTXStake.Contract.ExchangeRatio(&_IOTXStake.CallOpts)
}

// ExchangeRatio is a free data retrieval call binding the contract method 0x4006ccc5.
//
// Solidity: function exchangeRatio() view returns(uint256 ratio)
func (_IOTXStake *IOTXStakeCallerSession) ExchangeRatio() (*big.Int, error) {
	return _IOTXStake.Contract.ExchangeRatio(&_IOTXStake.CallOpts)
}

// GetRedeemedTokenIds is a free data retrieval call binding the contract method 0x69ced3c0.
//
// Solidity: function getRedeemedTokenIds(uint256 i, uint256 j) view returns(uint256[])
func (_IOTXStake *IOTXStakeCaller) GetRedeemedTokenIds(opts *bind.CallOpts, i *big.Int, j *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "getRedeemedTokenIds", i, j)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetRedeemedTokenIds is a free data retrieval call binding the contract method 0x69ced3c0.
//
// Solidity: function getRedeemedTokenIds(uint256 i, uint256 j) view returns(uint256[])
func (_IOTXStake *IOTXStakeSession) GetRedeemedTokenIds(i *big.Int, j *big.Int) ([]*big.Int, error) {
	return _IOTXStake.Contract.GetRedeemedTokenIds(&_IOTXStake.CallOpts, i, j)
}

// GetRedeemedTokenIds is a free data retrieval call binding the contract method 0x69ced3c0.
//
// Solidity: function getRedeemedTokenIds(uint256 i, uint256 j) view returns(uint256[])
func (_IOTXStake *IOTXStakeCallerSession) GetRedeemedTokenIds(i *big.Int, j *big.Int) ([]*big.Int, error) {
	return _IOTXStake.Contract.GetRedeemedTokenIds(&_IOTXStake.CallOpts, i, j)
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

// GetStakedTokenCount is a free data retrieval call binding the contract method 0x2291f70a.
//
// Solidity: function getStakedTokenCount(uint256 tokenQueueIndex) view returns(uint256 count)
func (_IOTXStake *IOTXStakeCaller) GetStakedTokenCount(opts *bind.CallOpts, tokenQueueIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "getStakedTokenCount", tokenQueueIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakedTokenCount is a free data retrieval call binding the contract method 0x2291f70a.
//
// Solidity: function getStakedTokenCount(uint256 tokenQueueIndex) view returns(uint256 count)
func (_IOTXStake *IOTXStakeSession) GetStakedTokenCount(tokenQueueIndex *big.Int) (*big.Int, error) {
	return _IOTXStake.Contract.GetStakedTokenCount(&_IOTXStake.CallOpts, tokenQueueIndex)
}

// GetStakedTokenCount is a free data retrieval call binding the contract method 0x2291f70a.
//
// Solidity: function getStakedTokenCount(uint256 tokenQueueIndex) view returns(uint256 count)
func (_IOTXStake *IOTXStakeCallerSession) GetStakedTokenCount(tokenQueueIndex *big.Int) (*big.Int, error) {
	return _IOTXStake.Contract.GetStakedTokenCount(&_IOTXStake.CallOpts, tokenQueueIndex)
}

// GlobalDelegate is a free data retrieval call binding the contract method 0x11e4bc2d.
//
// Solidity: function globalDelegate() view returns(address)
func (_IOTXStake *IOTXStakeCaller) GlobalDelegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "globalDelegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalDelegate is a free data retrieval call binding the contract method 0x11e4bc2d.
//
// Solidity: function globalDelegate() view returns(address)
func (_IOTXStake *IOTXStakeSession) GlobalDelegate() (common.Address, error) {
	return _IOTXStake.Contract.GlobalDelegate(&_IOTXStake.CallOpts)
}

// GlobalDelegate is a free data retrieval call binding the contract method 0x11e4bc2d.
//
// Solidity: function globalDelegate() view returns(address)
func (_IOTXStake *IOTXStakeCallerSession) GlobalDelegate() (common.Address, error) {
	return _IOTXStake.Contract.GlobalDelegate(&_IOTXStake.CallOpts)
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

// IotxClear is a free data retrieval call binding the contract method 0x5b3f86b4.
//
// Solidity: function iotxClear() view returns(address)
func (_IOTXStake *IOTXStakeCaller) IotxClear(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "iotxClear")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IotxClear is a free data retrieval call binding the contract method 0x5b3f86b4.
//
// Solidity: function iotxClear() view returns(address)
func (_IOTXStake *IOTXStakeSession) IotxClear() (common.Address, error) {
	return _IOTXStake.Contract.IotxClear(&_IOTXStake.CallOpts)
}

// IotxClear is a free data retrieval call binding the contract method 0x5b3f86b4.
//
// Solidity: function iotxClear() view returns(address)
func (_IOTXStake *IOTXStakeCallerSession) IotxClear() (common.Address, error) {
	return _IOTXStake.Contract.IotxClear(&_IOTXStake.CallOpts)
}

// ManagerFeeShares is a free data retrieval call binding the contract method 0xf6af7527.
//
// Solidity: function managerFeeShares() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) ManagerFeeShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "managerFeeShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerFeeShares is a free data retrieval call binding the contract method 0xf6af7527.
//
// Solidity: function managerFeeShares() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) ManagerFeeShares() (*big.Int, error) {
	return _IOTXStake.Contract.ManagerFeeShares(&_IOTXStake.CallOpts)
}

// ManagerFeeShares is a free data retrieval call binding the contract method 0xf6af7527.
//
// Solidity: function managerFeeShares() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) ManagerFeeShares() (*big.Int, error) {
	return _IOTXStake.Contract.ManagerFeeShares(&_IOTXStake.CallOpts)
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

// RedeemAmountBase is a free data retrieval call binding the contract method 0x37f52345.
//
// Solidity: function redeemAmountBase() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) RedeemAmountBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "redeemAmountBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemAmountBase is a free data retrieval call binding the contract method 0x37f52345.
//
// Solidity: function redeemAmountBase() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) RedeemAmountBase() (*big.Int, error) {
	return _IOTXStake.Contract.RedeemAmountBase(&_IOTXStake.CallOpts)
}

// RedeemAmountBase is a free data retrieval call binding the contract method 0x37f52345.
//
// Solidity: function redeemAmountBase() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) RedeemAmountBase() (*big.Int, error) {
	return _IOTXStake.Contract.RedeemAmountBase(&_IOTXStake.CallOpts)
}

// RedeemedTokenCount is a free data retrieval call binding the contract method 0xb9bafb72.
//
// Solidity: function redeemedTokenCount() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) RedeemedTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "redeemedTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemedTokenCount is a free data retrieval call binding the contract method 0xb9bafb72.
//
// Solidity: function redeemedTokenCount() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) RedeemedTokenCount() (*big.Int, error) {
	return _IOTXStake.Contract.RedeemedTokenCount(&_IOTXStake.CallOpts)
}

// RedeemedTokenCount is a free data retrieval call binding the contract method 0xb9bafb72.
//
// Solidity: function redeemedTokenCount() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) RedeemedTokenCount() (*big.Int, error) {
	return _IOTXStake.Contract.RedeemedTokenCount(&_IOTXStake.CallOpts)
}

// SequenceLength is a free data retrieval call binding the contract method 0xc030a6b9.
//
// Solidity: function sequenceLength() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) SequenceLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "sequenceLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequenceLength is a free data retrieval call binding the contract method 0xc030a6b9.
//
// Solidity: function sequenceLength() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) SequenceLength() (*big.Int, error) {
	return _IOTXStake.Contract.SequenceLength(&_IOTXStake.CallOpts)
}

// SequenceLength is a free data retrieval call binding the contract method 0xc030a6b9.
//
// Solidity: function sequenceLength() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) SequenceLength() (*big.Int, error) {
	return _IOTXStake.Contract.SequenceLength(&_IOTXStake.CallOpts)
}

// StakeDuration is a free data retrieval call binding the contract method 0x40f02ab6.
//
// Solidity: function stakeDuration() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) StakeDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "stakeDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeDuration is a free data retrieval call binding the contract method 0x40f02ab6.
//
// Solidity: function stakeDuration() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) StakeDuration() (*big.Int, error) {
	return _IOTXStake.Contract.StakeDuration(&_IOTXStake.CallOpts)
}

// StakeDuration is a free data retrieval call binding the contract method 0x40f02ab6.
//
// Solidity: function stakeDuration() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) StakeDuration() (*big.Int, error) {
	return _IOTXStake.Contract.StakeDuration(&_IOTXStake.CallOpts)
}

// StartAmount is a free data retrieval call binding the contract method 0x0f767872.
//
// Solidity: function startAmount() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) StartAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "startAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartAmount is a free data retrieval call binding the contract method 0x0f767872.
//
// Solidity: function startAmount() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) StartAmount() (*big.Int, error) {
	return _IOTXStake.Contract.StartAmount(&_IOTXStake.CallOpts)
}

// StartAmount is a free data retrieval call binding the contract method 0x0f767872.
//
// Solidity: function startAmount() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) StartAmount() (*big.Int, error) {
	return _IOTXStake.Contract.StartAmount(&_IOTXStake.CallOpts)
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

// TokenQueues is a free data retrieval call binding the contract method 0x4777aa71.
//
// Solidity: function tokenQueues(uint256 , uint256 ) view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) TokenQueues(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "tokenQueues", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenQueues is a free data retrieval call binding the contract method 0x4777aa71.
//
// Solidity: function tokenQueues(uint256 , uint256 ) view returns(uint256)
func (_IOTXStake *IOTXStakeSession) TokenQueues(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _IOTXStake.Contract.TokenQueues(&_IOTXStake.CallOpts, arg0, arg1)
}

// TokenQueues is a free data retrieval call binding the contract method 0x4777aa71.
//
// Solidity: function tokenQueues(uint256 , uint256 ) view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) TokenQueues(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _IOTXStake.Contract.TokenQueues(&_IOTXStake.CallOpts, arg0, arg1)
}

// TotalPending is a free data retrieval call binding the contract method 0x3f90916a.
//
// Solidity: function totalPending() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) TotalPending(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "totalPending")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPending is a free data retrieval call binding the contract method 0x3f90916a.
//
// Solidity: function totalPending() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) TotalPending() (*big.Int, error) {
	return _IOTXStake.Contract.TotalPending(&_IOTXStake.CallOpts)
}

// TotalPending is a free data retrieval call binding the contract method 0x3f90916a.
//
// Solidity: function totalPending() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) TotalPending() (*big.Int, error) {
	return _IOTXStake.Contract.TotalPending(&_IOTXStake.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_IOTXStake *IOTXStakeCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_IOTXStake *IOTXStakeSession) TotalStaked() (*big.Int, error) {
	return _IOTXStake.Contract.TotalStaked(&_IOTXStake.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_IOTXStake *IOTXStakeCallerSession) TotalStaked() (*big.Int, error) {
	return _IOTXStake.Contract.TotalStaked(&_IOTXStake.CallOpts)
}

// UniIOTX is a free data retrieval call binding the contract method 0x35b83d28.
//
// Solidity: function uniIOTX() view returns(address)
func (_IOTXStake *IOTXStakeCaller) UniIOTX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IOTXStake.contract.Call(opts, &out, "uniIOTX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniIOTX is a free data retrieval call binding the contract method 0x35b83d28.
//
// Solidity: function uniIOTX() view returns(address)
func (_IOTXStake *IOTXStakeSession) UniIOTX() (common.Address, error) {
	return _IOTXStake.Contract.UniIOTX(&_IOTXStake.CallOpts)
}

// UniIOTX is a free data retrieval call binding the contract method 0x35b83d28.
//
// Solidity: function uniIOTX() view returns(address)
func (_IOTXStake *IOTXStakeCallerSession) UniIOTX() (common.Address, error) {
	return _IOTXStake.Contract.UniIOTX(&_IOTXStake.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 minToMint, uint256 deadline) payable returns(uint256 minted)
func (_IOTXStake *IOTXStakeTransactor) Deposit(opts *bind.TransactOpts, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "deposit", minToMint, deadline)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 minToMint, uint256 deadline) payable returns(uint256 minted)
func (_IOTXStake *IOTXStakeSession) Deposit(minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IOTXStake.Contract.Deposit(&_IOTXStake.TransactOpts, minToMint, deadline)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 minToMint, uint256 deadline) payable returns(uint256 minted)
func (_IOTXStake *IOTXStakeTransactorSession) Deposit(minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IOTXStake.Contract.Deposit(&_IOTXStake.TransactOpts, minToMint, deadline)
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

// Initialize is a paid mutator transaction binding the contract method 0xc843e81c.
//
// Solidity: function initialize(address _systemStake, address _uniIOTX, address _iotxClear, address _oracle, uint256 _startAmount, uint256 _commonRatio, uint256 _sequenceLength, uint256 _stakeDuration) returns()
func (_IOTXStake *IOTXStakeTransactor) Initialize(opts *bind.TransactOpts, _systemStake common.Address, _uniIOTX common.Address, _iotxClear common.Address, _oracle common.Address, _startAmount *big.Int, _commonRatio *big.Int, _sequenceLength *big.Int, _stakeDuration *big.Int) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "initialize", _systemStake, _uniIOTX, _iotxClear, _oracle, _startAmount, _commonRatio, _sequenceLength, _stakeDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xc843e81c.
//
// Solidity: function initialize(address _systemStake, address _uniIOTX, address _iotxClear, address _oracle, uint256 _startAmount, uint256 _commonRatio, uint256 _sequenceLength, uint256 _stakeDuration) returns()
func (_IOTXStake *IOTXStakeSession) Initialize(_systemStake common.Address, _uniIOTX common.Address, _iotxClear common.Address, _oracle common.Address, _startAmount *big.Int, _commonRatio *big.Int, _sequenceLength *big.Int, _stakeDuration *big.Int) (*types.Transaction, error) {
	return _IOTXStake.Contract.Initialize(&_IOTXStake.TransactOpts, _systemStake, _uniIOTX, _iotxClear, _oracle, _startAmount, _commonRatio, _sequenceLength, _stakeDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xc843e81c.
//
// Solidity: function initialize(address _systemStake, address _uniIOTX, address _iotxClear, address _oracle, uint256 _startAmount, uint256 _commonRatio, uint256 _sequenceLength, uint256 _stakeDuration) returns()
func (_IOTXStake *IOTXStakeTransactorSession) Initialize(_systemStake common.Address, _uniIOTX common.Address, _iotxClear common.Address, _oracle common.Address, _startAmount *big.Int, _commonRatio *big.Int, _sequenceLength *big.Int, _stakeDuration *big.Int) (*types.Transaction, error) {
	return _IOTXStake.Contract.Initialize(&_IOTXStake.TransactOpts, _systemStake, _uniIOTX, _iotxClear, _oracle, _startAmount, _commonRatio, _sequenceLength, _stakeDuration)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IOTXStake *IOTXStakeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IOTXStake *IOTXStakeSession) Pause() (*types.Transaction, error) {
	return _IOTXStake.Contract.Pause(&_IOTXStake.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IOTXStake *IOTXStakeTransactorSession) Pause() (*types.Transaction, error) {
	return _IOTXStake.Contract.Pause(&_IOTXStake.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xb8192205.
//
// Solidity: function redeem(uint256 iotxsToRedeem, uint256 maxToBurn, uint256 deadline) returns(uint256 burned)
func (_IOTXStake *IOTXStakeTransactor) Redeem(opts *bind.TransactOpts, iotxsToRedeem *big.Int, maxToBurn *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "redeem", iotxsToRedeem, maxToBurn, deadline)
}

// Redeem is a paid mutator transaction binding the contract method 0xb8192205.
//
// Solidity: function redeem(uint256 iotxsToRedeem, uint256 maxToBurn, uint256 deadline) returns(uint256 burned)
func (_IOTXStake *IOTXStakeSession) Redeem(iotxsToRedeem *big.Int, maxToBurn *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IOTXStake.Contract.Redeem(&_IOTXStake.TransactOpts, iotxsToRedeem, maxToBurn, deadline)
}

// Redeem is a paid mutator transaction binding the contract method 0xb8192205.
//
// Solidity: function redeem(uint256 iotxsToRedeem, uint256 maxToBurn, uint256 deadline) returns(uint256 burned)
func (_IOTXStake *IOTXStakeTransactorSession) Redeem(iotxsToRedeem *big.Int, maxToBurn *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IOTXStake.Contract.Redeem(&_IOTXStake.TransactOpts, iotxsToRedeem, maxToBurn, deadline)
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

// SetGlobalDelegate is a paid mutator transaction binding the contract method 0x229bf1b6.
//
// Solidity: function setGlobalDelegate(address delegate) returns()
func (_IOTXStake *IOTXStakeTransactor) SetGlobalDelegate(opts *bind.TransactOpts, delegate common.Address) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "setGlobalDelegate", delegate)
}

// SetGlobalDelegate is a paid mutator transaction binding the contract method 0x229bf1b6.
//
// Solidity: function setGlobalDelegate(address delegate) returns()
func (_IOTXStake *IOTXStakeSession) SetGlobalDelegate(delegate common.Address) (*types.Transaction, error) {
	return _IOTXStake.Contract.SetGlobalDelegate(&_IOTXStake.TransactOpts, delegate)
}

// SetGlobalDelegate is a paid mutator transaction binding the contract method 0x229bf1b6.
//
// Solidity: function setGlobalDelegate(address delegate) returns()
func (_IOTXStake *IOTXStakeTransactorSession) SetGlobalDelegate(delegate common.Address) (*types.Transaction, error) {
	return _IOTXStake.Contract.SetGlobalDelegate(&_IOTXStake.TransactOpts, delegate)
}

// SetManagerFeeShares is a paid mutator transaction binding the contract method 0xbecd0858.
//
// Solidity: function setManagerFeeShares(uint256 shares) returns()
func (_IOTXStake *IOTXStakeTransactor) SetManagerFeeShares(opts *bind.TransactOpts, shares *big.Int) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "setManagerFeeShares", shares)
}

// SetManagerFeeShares is a paid mutator transaction binding the contract method 0xbecd0858.
//
// Solidity: function setManagerFeeShares(uint256 shares) returns()
func (_IOTXStake *IOTXStakeSession) SetManagerFeeShares(shares *big.Int) (*types.Transaction, error) {
	return _IOTXStake.Contract.SetManagerFeeShares(&_IOTXStake.TransactOpts, shares)
}

// SetManagerFeeShares is a paid mutator transaction binding the contract method 0xbecd0858.
//
// Solidity: function setManagerFeeShares(uint256 shares) returns()
func (_IOTXStake *IOTXStakeTransactorSession) SetManagerFeeShares(shares *big.Int) (*types.Transaction, error) {
	return _IOTXStake.Contract.SetManagerFeeShares(&_IOTXStake.TransactOpts, shares)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_IOTXStake *IOTXStakeTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_IOTXStake *IOTXStakeSession) Stake() (*types.Transaction, error) {
	return _IOTXStake.Contract.Stake(&_IOTXStake.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_IOTXStake *IOTXStakeTransactorSession) Stake() (*types.Transaction, error) {
	return _IOTXStake.Contract.Stake(&_IOTXStake.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IOTXStake *IOTXStakeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IOTXStake *IOTXStakeSession) Unpause() (*types.Transaction, error) {
	return _IOTXStake.Contract.Unpause(&_IOTXStake.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IOTXStake *IOTXStakeTransactorSession) Unpause() (*types.Transaction, error) {
	return _IOTXStake.Contract.Unpause(&_IOTXStake.TransactOpts)
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

// UpdateReward is a paid mutator transaction binding the contract method 0xf36c0a72.
//
// Solidity: function updateReward() returns()
func (_IOTXStake *IOTXStakeTransactor) UpdateReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "updateReward")
}

// UpdateReward is a paid mutator transaction binding the contract method 0xf36c0a72.
//
// Solidity: function updateReward() returns()
func (_IOTXStake *IOTXStakeSession) UpdateReward() (*types.Transaction, error) {
	return _IOTXStake.Contract.UpdateReward(&_IOTXStake.TransactOpts)
}

// UpdateReward is a paid mutator transaction binding the contract method 0xf36c0a72.
//
// Solidity: function updateReward() returns()
func (_IOTXStake *IOTXStakeTransactorSession) UpdateReward() (*types.Transaction, error) {
	return _IOTXStake.Contract.UpdateReward(&_IOTXStake.TransactOpts)
}

// WithdrawManagerFee is a paid mutator transaction binding the contract method 0xfa0ec2ce.
//
// Solidity: function withdrawManagerFee(uint256 amount, address recipient) returns()
func (_IOTXStake *IOTXStakeTransactor) WithdrawManagerFee(opts *bind.TransactOpts, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IOTXStake.contract.Transact(opts, "withdrawManagerFee", amount, recipient)
}

// WithdrawManagerFee is a paid mutator transaction binding the contract method 0xfa0ec2ce.
//
// Solidity: function withdrawManagerFee(uint256 amount, address recipient) returns()
func (_IOTXStake *IOTXStakeSession) WithdrawManagerFee(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IOTXStake.Contract.WithdrawManagerFee(&_IOTXStake.TransactOpts, amount, recipient)
}

// WithdrawManagerFee is a paid mutator transaction binding the contract method 0xfa0ec2ce.
//
// Solidity: function withdrawManagerFee(uint256 amount, address recipient) returns()
func (_IOTXStake *IOTXStakeTransactorSession) WithdrawManagerFee(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IOTXStake.Contract.WithdrawManagerFee(&_IOTXStake.TransactOpts, amount, recipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IOTXStake *IOTXStakeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXStake.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IOTXStake *IOTXStakeSession) Receive() (*types.Transaction, error) {
	return _IOTXStake.Contract.Receive(&_IOTXStake.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IOTXStake *IOTXStakeTransactorSession) Receive() (*types.Transaction, error) {
	return _IOTXStake.Contract.Receive(&_IOTXStake.TransactOpts)
}

// IOTXStakeDelegatesUpdatedIterator is returned from FilterDelegatesUpdated and is used to iterate over the raw logs and unpacked data for DelegatesUpdated events raised by the IOTXStake contract.
type IOTXStakeDelegatesUpdatedIterator struct {
	Event *IOTXStakeDelegatesUpdated // Event containing the contract specifics and raw log

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
func (it *IOTXStakeDelegatesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakeDelegatesUpdated)
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
		it.Event = new(IOTXStakeDelegatesUpdated)
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
func (it *IOTXStakeDelegatesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakeDelegatesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakeDelegatesUpdated represents a DelegatesUpdated event raised by the IOTXStake contract.
type IOTXStakeDelegatesUpdated struct {
	TokenIds []*big.Int
	Delegate common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegatesUpdated is a free log retrieval operation binding the contract event 0x055e970cad70d45557223e4ef35ac45bb162f8fc6f8a5986159f49d7d7fc742b.
//
// Solidity: event DelegatesUpdated(uint256[] tokenIds, address delegate)
func (_IOTXStake *IOTXStakeFilterer) FilterDelegatesUpdated(opts *bind.FilterOpts) (*IOTXStakeDelegatesUpdatedIterator, error) {

	logs, sub, err := _IOTXStake.contract.FilterLogs(opts, "DelegatesUpdated")
	if err != nil {
		return nil, err
	}
	return &IOTXStakeDelegatesUpdatedIterator{contract: _IOTXStake.contract, event: "DelegatesUpdated", logs: logs, sub: sub}, nil
}

// WatchDelegatesUpdated is a free log subscription operation binding the contract event 0x055e970cad70d45557223e4ef35ac45bb162f8fc6f8a5986159f49d7d7fc742b.
//
// Solidity: event DelegatesUpdated(uint256[] tokenIds, address delegate)
func (_IOTXStake *IOTXStakeFilterer) WatchDelegatesUpdated(opts *bind.WatchOpts, sink chan<- *IOTXStakeDelegatesUpdated) (event.Subscription, error) {

	logs, sub, err := _IOTXStake.contract.WatchLogs(opts, "DelegatesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakeDelegatesUpdated)
				if err := _IOTXStake.contract.UnpackLog(event, "DelegatesUpdated", log); err != nil {
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
func (_IOTXStake *IOTXStakeFilterer) ParseDelegatesUpdated(log types.Log) (*IOTXStakeDelegatesUpdated, error) {
	event := new(IOTXStakeDelegatesUpdated)
	if err := _IOTXStake.contract.UnpackLog(event, "DelegatesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakeGlobalDelegateSetIterator is returned from FilterGlobalDelegateSet and is used to iterate over the raw logs and unpacked data for GlobalDelegateSet events raised by the IOTXStake contract.
type IOTXStakeGlobalDelegateSetIterator struct {
	Event *IOTXStakeGlobalDelegateSet // Event containing the contract specifics and raw log

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
func (it *IOTXStakeGlobalDelegateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakeGlobalDelegateSet)
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
		it.Event = new(IOTXStakeGlobalDelegateSet)
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
func (it *IOTXStakeGlobalDelegateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakeGlobalDelegateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakeGlobalDelegateSet represents a GlobalDelegateSet event raised by the IOTXStake contract.
type IOTXStakeGlobalDelegateSet struct {
	Delegate common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGlobalDelegateSet is a free log retrieval operation binding the contract event 0x19320da70b0c544e76e5dbc3c85129ac6d5d0d5486b4aa42982640788a157d71.
//
// Solidity: event GlobalDelegateSet(address delegate)
func (_IOTXStake *IOTXStakeFilterer) FilterGlobalDelegateSet(opts *bind.FilterOpts) (*IOTXStakeGlobalDelegateSetIterator, error) {

	logs, sub, err := _IOTXStake.contract.FilterLogs(opts, "GlobalDelegateSet")
	if err != nil {
		return nil, err
	}
	return &IOTXStakeGlobalDelegateSetIterator{contract: _IOTXStake.contract, event: "GlobalDelegateSet", logs: logs, sub: sub}, nil
}

// WatchGlobalDelegateSet is a free log subscription operation binding the contract event 0x19320da70b0c544e76e5dbc3c85129ac6d5d0d5486b4aa42982640788a157d71.
//
// Solidity: event GlobalDelegateSet(address delegate)
func (_IOTXStake *IOTXStakeFilterer) WatchGlobalDelegateSet(opts *bind.WatchOpts, sink chan<- *IOTXStakeGlobalDelegateSet) (event.Subscription, error) {

	logs, sub, err := _IOTXStake.contract.WatchLogs(opts, "GlobalDelegateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakeGlobalDelegateSet)
				if err := _IOTXStake.contract.UnpackLog(event, "GlobalDelegateSet", log); err != nil {
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

// ParseGlobalDelegateSet is a log parse operation binding the contract event 0x19320da70b0c544e76e5dbc3c85129ac6d5d0d5486b4aa42982640788a157d71.
//
// Solidity: event GlobalDelegateSet(address delegate)
func (_IOTXStake *IOTXStakeFilterer) ParseGlobalDelegateSet(log types.Log) (*IOTXStakeGlobalDelegateSet, error) {
	event := new(IOTXStakeGlobalDelegateSet)
	if err := _IOTXStake.contract.UnpackLog(event, "GlobalDelegateSet", log); err != nil {
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

// IOTXStakeManagerFeeSharesSetIterator is returned from FilterManagerFeeSharesSet and is used to iterate over the raw logs and unpacked data for ManagerFeeSharesSet events raised by the IOTXStake contract.
type IOTXStakeManagerFeeSharesSetIterator struct {
	Event *IOTXStakeManagerFeeSharesSet // Event containing the contract specifics and raw log

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
func (it *IOTXStakeManagerFeeSharesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakeManagerFeeSharesSet)
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
		it.Event = new(IOTXStakeManagerFeeSharesSet)
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
func (it *IOTXStakeManagerFeeSharesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakeManagerFeeSharesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakeManagerFeeSharesSet represents a ManagerFeeSharesSet event raised by the IOTXStake contract.
type IOTXStakeManagerFeeSharesSet struct {
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterManagerFeeSharesSet is a free log retrieval operation binding the contract event 0xb6efe094bbb68f7f840c92b10663dc9fdee2f9c7b4b7ba75d88b13e0b685b228.
//
// Solidity: event ManagerFeeSharesSet(uint256 shares)
func (_IOTXStake *IOTXStakeFilterer) FilterManagerFeeSharesSet(opts *bind.FilterOpts) (*IOTXStakeManagerFeeSharesSetIterator, error) {

	logs, sub, err := _IOTXStake.contract.FilterLogs(opts, "ManagerFeeSharesSet")
	if err != nil {
		return nil, err
	}
	return &IOTXStakeManagerFeeSharesSetIterator{contract: _IOTXStake.contract, event: "ManagerFeeSharesSet", logs: logs, sub: sub}, nil
}

// WatchManagerFeeSharesSet is a free log subscription operation binding the contract event 0xb6efe094bbb68f7f840c92b10663dc9fdee2f9c7b4b7ba75d88b13e0b685b228.
//
// Solidity: event ManagerFeeSharesSet(uint256 shares)
func (_IOTXStake *IOTXStakeFilterer) WatchManagerFeeSharesSet(opts *bind.WatchOpts, sink chan<- *IOTXStakeManagerFeeSharesSet) (event.Subscription, error) {

	logs, sub, err := _IOTXStake.contract.WatchLogs(opts, "ManagerFeeSharesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakeManagerFeeSharesSet)
				if err := _IOTXStake.contract.UnpackLog(event, "ManagerFeeSharesSet", log); err != nil {
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

// ParseManagerFeeSharesSet is a log parse operation binding the contract event 0xb6efe094bbb68f7f840c92b10663dc9fdee2f9c7b4b7ba75d88b13e0b685b228.
//
// Solidity: event ManagerFeeSharesSet(uint256 shares)
func (_IOTXStake *IOTXStakeFilterer) ParseManagerFeeSharesSet(log types.Log) (*IOTXStakeManagerFeeSharesSet, error) {
	event := new(IOTXStakeManagerFeeSharesSet)
	if err := _IOTXStake.contract.UnpackLog(event, "ManagerFeeSharesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakeManagerFeeWithdrawedIterator is returned from FilterManagerFeeWithdrawed and is used to iterate over the raw logs and unpacked data for ManagerFeeWithdrawed events raised by the IOTXStake contract.
type IOTXStakeManagerFeeWithdrawedIterator struct {
	Event *IOTXStakeManagerFeeWithdrawed // Event containing the contract specifics and raw log

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
func (it *IOTXStakeManagerFeeWithdrawedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakeManagerFeeWithdrawed)
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
		it.Event = new(IOTXStakeManagerFeeWithdrawed)
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
func (it *IOTXStakeManagerFeeWithdrawedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakeManagerFeeWithdrawedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakeManagerFeeWithdrawed represents a ManagerFeeWithdrawed event raised by the IOTXStake contract.
type IOTXStakeManagerFeeWithdrawed struct {
	Amount    *big.Int
	Minted    *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterManagerFeeWithdrawed is a free log retrieval operation binding the contract event 0xce3d910b07bcccb2bc98f93c5292e13a264d588f34ef3740940f872cea14a7eb.
//
// Solidity: event ManagerFeeWithdrawed(uint256 amount, uint256 minted, address recipient)
func (_IOTXStake *IOTXStakeFilterer) FilterManagerFeeWithdrawed(opts *bind.FilterOpts) (*IOTXStakeManagerFeeWithdrawedIterator, error) {

	logs, sub, err := _IOTXStake.contract.FilterLogs(opts, "ManagerFeeWithdrawed")
	if err != nil {
		return nil, err
	}
	return &IOTXStakeManagerFeeWithdrawedIterator{contract: _IOTXStake.contract, event: "ManagerFeeWithdrawed", logs: logs, sub: sub}, nil
}

// WatchManagerFeeWithdrawed is a free log subscription operation binding the contract event 0xce3d910b07bcccb2bc98f93c5292e13a264d588f34ef3740940f872cea14a7eb.
//
// Solidity: event ManagerFeeWithdrawed(uint256 amount, uint256 minted, address recipient)
func (_IOTXStake *IOTXStakeFilterer) WatchManagerFeeWithdrawed(opts *bind.WatchOpts, sink chan<- *IOTXStakeManagerFeeWithdrawed) (event.Subscription, error) {

	logs, sub, err := _IOTXStake.contract.WatchLogs(opts, "ManagerFeeWithdrawed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakeManagerFeeWithdrawed)
				if err := _IOTXStake.contract.UnpackLog(event, "ManagerFeeWithdrawed", log); err != nil {
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

// ParseManagerFeeWithdrawed is a log parse operation binding the contract event 0xce3d910b07bcccb2bc98f93c5292e13a264d588f34ef3740940f872cea14a7eb.
//
// Solidity: event ManagerFeeWithdrawed(uint256 amount, uint256 minted, address recipient)
func (_IOTXStake *IOTXStakeFilterer) ParseManagerFeeWithdrawed(log types.Log) (*IOTXStakeManagerFeeWithdrawed, error) {
	event := new(IOTXStakeManagerFeeWithdrawed)
	if err := _IOTXStake.contract.UnpackLog(event, "ManagerFeeWithdrawed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakeMergedIterator is returned from FilterMerged and is used to iterate over the raw logs and unpacked data for Merged events raised by the IOTXStake contract.
type IOTXStakeMergedIterator struct {
	Event *IOTXStakeMerged // Event containing the contract specifics and raw log

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
func (it *IOTXStakeMergedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakeMerged)
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
		it.Event = new(IOTXStakeMerged)
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
func (it *IOTXStakeMergedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakeMergedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakeMerged represents a Merged event raised by the IOTXStake contract.
type IOTXStakeMerged struct {
	TokenIds []*big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMerged is a free log retrieval operation binding the contract event 0x0c834fb4dca660c7c0016cf0716f341c5bd26cd91f18223279e15942b2fe0cda.
//
// Solidity: event Merged(uint256[] tokenIds, uint256 amount)
func (_IOTXStake *IOTXStakeFilterer) FilterMerged(opts *bind.FilterOpts) (*IOTXStakeMergedIterator, error) {

	logs, sub, err := _IOTXStake.contract.FilterLogs(opts, "Merged")
	if err != nil {
		return nil, err
	}
	return &IOTXStakeMergedIterator{contract: _IOTXStake.contract, event: "Merged", logs: logs, sub: sub}, nil
}

// WatchMerged is a free log subscription operation binding the contract event 0x0c834fb4dca660c7c0016cf0716f341c5bd26cd91f18223279e15942b2fe0cda.
//
// Solidity: event Merged(uint256[] tokenIds, uint256 amount)
func (_IOTXStake *IOTXStakeFilterer) WatchMerged(opts *bind.WatchOpts, sink chan<- *IOTXStakeMerged) (event.Subscription, error) {

	logs, sub, err := _IOTXStake.contract.WatchLogs(opts, "Merged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakeMerged)
				if err := _IOTXStake.contract.UnpackLog(event, "Merged", log); err != nil {
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

// ParseMerged is a log parse operation binding the contract event 0x0c834fb4dca660c7c0016cf0716f341c5bd26cd91f18223279e15942b2fe0cda.
//
// Solidity: event Merged(uint256[] tokenIds, uint256 amount)
func (_IOTXStake *IOTXStakeFilterer) ParseMerged(log types.Log) (*IOTXStakeMerged, error) {
	event := new(IOTXStakeMerged)
	if err := _IOTXStake.contract.UnpackLog(event, "Merged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakeMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the IOTXStake contract.
type IOTXStakeMintedIterator struct {
	Event *IOTXStakeMinted // Event containing the contract specifics and raw log

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
func (it *IOTXStakeMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakeMinted)
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
		it.Event = new(IOTXStakeMinted)
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
func (it *IOTXStakeMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakeMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakeMinted represents a Minted event raised by the IOTXStake contract.
type IOTXStakeMinted struct {
	User   common.Address
	Minted *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address user, uint256 minted)
func (_IOTXStake *IOTXStakeFilterer) FilterMinted(opts *bind.FilterOpts) (*IOTXStakeMintedIterator, error) {

	logs, sub, err := _IOTXStake.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &IOTXStakeMintedIterator{contract: _IOTXStake.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address user, uint256 minted)
func (_IOTXStake *IOTXStakeFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *IOTXStakeMinted) (event.Subscription, error) {

	logs, sub, err := _IOTXStake.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakeMinted)
				if err := _IOTXStake.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address user, uint256 minted)
func (_IOTXStake *IOTXStakeFilterer) ParseMinted(log types.Log) (*IOTXStakeMinted, error) {
	event := new(IOTXStakeMinted)
	if err := _IOTXStake.contract.UnpackLog(event, "Minted", log); err != nil {
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

// IOTXStakeRedeemedIterator is returned from FilterRedeemed and is used to iterate over the raw logs and unpacked data for Redeemed events raised by the IOTXStake contract.
type IOTXStakeRedeemedIterator struct {
	Event *IOTXStakeRedeemed // Event containing the contract specifics and raw log

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
func (it *IOTXStakeRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakeRedeemed)
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
		it.Event = new(IOTXStakeRedeemed)
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
func (it *IOTXStakeRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakeRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakeRedeemed represents a Redeemed event raised by the IOTXStake contract.
type IOTXStakeRedeemed struct {
	User     common.Address
	Burned   *big.Int
	TokenIds []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRedeemed is a free log retrieval operation binding the contract event 0x6b921527f65949c80fab346e1e330b442e353b32e7b29f428d3d0c7766dd9152.
//
// Solidity: event Redeemed(address user, uint256 burned, uint256[] tokenIds)
func (_IOTXStake *IOTXStakeFilterer) FilterRedeemed(opts *bind.FilterOpts) (*IOTXStakeRedeemedIterator, error) {

	logs, sub, err := _IOTXStake.contract.FilterLogs(opts, "Redeemed")
	if err != nil {
		return nil, err
	}
	return &IOTXStakeRedeemedIterator{contract: _IOTXStake.contract, event: "Redeemed", logs: logs, sub: sub}, nil
}

// WatchRedeemed is a free log subscription operation binding the contract event 0x6b921527f65949c80fab346e1e330b442e353b32e7b29f428d3d0c7766dd9152.
//
// Solidity: event Redeemed(address user, uint256 burned, uint256[] tokenIds)
func (_IOTXStake *IOTXStakeFilterer) WatchRedeemed(opts *bind.WatchOpts, sink chan<- *IOTXStakeRedeemed) (event.Subscription, error) {

	logs, sub, err := _IOTXStake.contract.WatchLogs(opts, "Redeemed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakeRedeemed)
				if err := _IOTXStake.contract.UnpackLog(event, "Redeemed", log); err != nil {
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

// ParseRedeemed is a log parse operation binding the contract event 0x6b921527f65949c80fab346e1e330b442e353b32e7b29f428d3d0c7766dd9152.
//
// Solidity: event Redeemed(address user, uint256 burned, uint256[] tokenIds)
func (_IOTXStake *IOTXStakeFilterer) ParseRedeemed(log types.Log) (*IOTXStakeRedeemed, error) {
	event := new(IOTXStakeRedeemed)
	if err := _IOTXStake.contract.UnpackLog(event, "Redeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakeRewardUpdatedIterator is returned from FilterRewardUpdated and is used to iterate over the raw logs and unpacked data for RewardUpdated events raised by the IOTXStake contract.
type IOTXStakeRewardUpdatedIterator struct {
	Event *IOTXStakeRewardUpdated // Event containing the contract specifics and raw log

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
func (it *IOTXStakeRewardUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakeRewardUpdated)
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
		it.Event = new(IOTXStakeRewardUpdated)
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
func (it *IOTXStakeRewardUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakeRewardUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakeRewardUpdated represents a RewardUpdated event raised by the IOTXStake contract.
type IOTXStakeRewardUpdated struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardUpdated is a free log retrieval operation binding the contract event 0xcb94909754d27c309adf4167150f1f7aa04de40b6a0e6bb98b2ae80a2bf438f6.
//
// Solidity: event RewardUpdated(uint256 amount)
func (_IOTXStake *IOTXStakeFilterer) FilterRewardUpdated(opts *bind.FilterOpts) (*IOTXStakeRewardUpdatedIterator, error) {

	logs, sub, err := _IOTXStake.contract.FilterLogs(opts, "RewardUpdated")
	if err != nil {
		return nil, err
	}
	return &IOTXStakeRewardUpdatedIterator{contract: _IOTXStake.contract, event: "RewardUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardUpdated is a free log subscription operation binding the contract event 0xcb94909754d27c309adf4167150f1f7aa04de40b6a0e6bb98b2ae80a2bf438f6.
//
// Solidity: event RewardUpdated(uint256 amount)
func (_IOTXStake *IOTXStakeFilterer) WatchRewardUpdated(opts *bind.WatchOpts, sink chan<- *IOTXStakeRewardUpdated) (event.Subscription, error) {

	logs, sub, err := _IOTXStake.contract.WatchLogs(opts, "RewardUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakeRewardUpdated)
				if err := _IOTXStake.contract.UnpackLog(event, "RewardUpdated", log); err != nil {
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

// ParseRewardUpdated is a log parse operation binding the contract event 0xcb94909754d27c309adf4167150f1f7aa04de40b6a0e6bb98b2ae80a2bf438f6.
//
// Solidity: event RewardUpdated(uint256 amount)
func (_IOTXStake *IOTXStakeFilterer) ParseRewardUpdated(log types.Log) (*IOTXStakeRewardUpdated, error) {
	event := new(IOTXStakeRewardUpdated)
	if err := _IOTXStake.contract.UnpackLog(event, "RewardUpdated", log); err != nil {
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

// IOTXStakeStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the IOTXStake contract.
type IOTXStakeStakedIterator struct {
	Event *IOTXStakeStaked // Event containing the contract specifics and raw log

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
func (it *IOTXStakeStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakeStaked)
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
		it.Event = new(IOTXStakeStaked)
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
func (it *IOTXStakeStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakeStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakeStaked represents a Staked event raised by the IOTXStake contract.
type IOTXStakeStaked struct {
	FirstTokenId *big.Int
	Amount       *big.Int
	Delegate     common.Address
	Count        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0xc49cabe9d2bbfea8d9f51b0961c30a1081ff6fd3c4d6a9182cd65de8cea2df00.
//
// Solidity: event Staked(uint256 firstTokenId, uint256 amount, address delegate, uint256 count)
func (_IOTXStake *IOTXStakeFilterer) FilterStaked(opts *bind.FilterOpts) (*IOTXStakeStakedIterator, error) {

	logs, sub, err := _IOTXStake.contract.FilterLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return &IOTXStakeStakedIterator{contract: _IOTXStake.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0xc49cabe9d2bbfea8d9f51b0961c30a1081ff6fd3c4d6a9182cd65de8cea2df00.
//
// Solidity: event Staked(uint256 firstTokenId, uint256 amount, address delegate, uint256 count)
func (_IOTXStake *IOTXStakeFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *IOTXStakeStaked) (event.Subscription, error) {

	logs, sub, err := _IOTXStake.contract.WatchLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakeStaked)
				if err := _IOTXStake.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0xc49cabe9d2bbfea8d9f51b0961c30a1081ff6fd3c4d6a9182cd65de8cea2df00.
//
// Solidity: event Staked(uint256 firstTokenId, uint256 amount, address delegate, uint256 count)
func (_IOTXStake *IOTXStakeFilterer) ParseStaked(log types.Log) (*IOTXStakeStaked, error) {
	event := new(IOTXStakeStaked)
	if err := _IOTXStake.contract.UnpackLog(event, "Staked", log); err != nil {
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
