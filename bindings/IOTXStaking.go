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

// IOTXStakingMetaData contains all meta data concerning the IOTXStaking contract.
var IOTXStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"DelegatesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"GlobalDelegateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"ManagerFeeSharesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ManagerFeeWithdrawed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromLevel\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toLevel\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Merged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"Redeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"firstTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountedManagerReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountedUserReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commonRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalDelegate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getManagerFeeShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getManagerReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRedeemAmountBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRedeemedTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRedeemedTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"j\",\"type\":\"uint256\"}],\"name\":\"getRedeemedTokenIdsSlice\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakeAmounts\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakeDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenQueueIndex\",\"type\":\"uint256\"}],\"name\":\"getStakedTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenQueueIndex\",\"type\":\"uint256\"}],\"name\":\"getStakedTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenQueueIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"j\",\"type\":\"uint256\"}],\"name\":\"getStakedTokenIdsSlice\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenQueueIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenIndex\",\"type\":\"uint256\"}],\"name\":\"getTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenQueueIndex\",\"type\":\"uint256\"}],\"name\":\"getTokenQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalDelegate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_systemStaking\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_uniIOTX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_iotxClear\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_commonRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sequenceLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakeDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"iotxClear\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerFeeShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"iotxsToRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"burned\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemAmountBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemedTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequenceLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"setGlobalDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"setManagerFeeShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemStaking\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniIOTX\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"updateDelegates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawManagerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// IOTXStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use IOTXStakingMetaData.ABI instead.
var IOTXStakingABI = IOTXStakingMetaData.ABI

// IOTXStaking is an auto generated Go binding around an Ethereum contract.
type IOTXStaking struct {
	IOTXStakingCaller     // Read-only binding to the contract
	IOTXStakingTransactor // Write-only binding to the contract
	IOTXStakingFilterer   // Log filterer for contract events
}

// IOTXStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOTXStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOTXStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOTXStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOTXStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOTXStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOTXStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOTXStakingSession struct {
	Contract     *IOTXStaking      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOTXStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOTXStakingCallerSession struct {
	Contract *IOTXStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IOTXStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOTXStakingTransactorSession struct {
	Contract     *IOTXStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IOTXStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOTXStakingRaw struct {
	Contract *IOTXStaking // Generic contract binding to access the raw methods on
}

// IOTXStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOTXStakingCallerRaw struct {
	Contract *IOTXStakingCaller // Generic read-only contract binding to access the raw methods on
}

// IOTXStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOTXStakingTransactorRaw struct {
	Contract *IOTXStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOTXStaking creates a new instance of IOTXStaking, bound to a specific deployed contract.
func NewIOTXStaking(address common.Address, backend bind.ContractBackend) (*IOTXStaking, error) {
	contract, err := bindIOTXStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOTXStaking{IOTXStakingCaller: IOTXStakingCaller{contract: contract}, IOTXStakingTransactor: IOTXStakingTransactor{contract: contract}, IOTXStakingFilterer: IOTXStakingFilterer{contract: contract}}, nil
}

// NewIOTXStakingCaller creates a new read-only instance of IOTXStaking, bound to a specific deployed contract.
func NewIOTXStakingCaller(address common.Address, caller bind.ContractCaller) (*IOTXStakingCaller, error) {
	contract, err := bindIOTXStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOTXStakingCaller{contract: contract}, nil
}

// NewIOTXStakingTransactor creates a new write-only instance of IOTXStaking, bound to a specific deployed contract.
func NewIOTXStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*IOTXStakingTransactor, error) {
	contract, err := bindIOTXStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOTXStakingTransactor{contract: contract}, nil
}

// NewIOTXStakingFilterer creates a new log filterer instance of IOTXStaking, bound to a specific deployed contract.
func NewIOTXStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*IOTXStakingFilterer, error) {
	contract, err := bindIOTXStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOTXStakingFilterer{contract: contract}, nil
}

// bindIOTXStaking binds a generic wrapper to an already deployed contract.
func bindIOTXStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOTXStakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOTXStaking *IOTXStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOTXStaking.Contract.IOTXStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOTXStaking *IOTXStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXStaking.Contract.IOTXStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOTXStaking *IOTXStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOTXStaking.Contract.IOTXStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOTXStaking *IOTXStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOTXStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOTXStaking *IOTXStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOTXStaking *IOTXStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOTXStaking.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IOTXStaking *IOTXStakingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IOTXStaking *IOTXStakingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IOTXStaking.Contract.DEFAULTADMINROLE(&_IOTXStaking.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IOTXStaking *IOTXStakingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IOTXStaking.Contract.DEFAULTADMINROLE(&_IOTXStaking.CallOpts)
}

// AccountedBalance is a free data retrieval call binding the contract method 0x0937eb54.
//
// Solidity: function accountedBalance() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) AccountedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "accountedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountedBalance is a free data retrieval call binding the contract method 0x0937eb54.
//
// Solidity: function accountedBalance() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) AccountedBalance() (*big.Int, error) {
	return _IOTXStaking.Contract.AccountedBalance(&_IOTXStaking.CallOpts)
}

// AccountedBalance is a free data retrieval call binding the contract method 0x0937eb54.
//
// Solidity: function accountedBalance() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) AccountedBalance() (*big.Int, error) {
	return _IOTXStaking.Contract.AccountedBalance(&_IOTXStaking.CallOpts)
}

// AccountedManagerReward is a free data retrieval call binding the contract method 0x48c26fce.
//
// Solidity: function accountedManagerReward() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) AccountedManagerReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "accountedManagerReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountedManagerReward is a free data retrieval call binding the contract method 0x48c26fce.
//
// Solidity: function accountedManagerReward() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) AccountedManagerReward() (*big.Int, error) {
	return _IOTXStaking.Contract.AccountedManagerReward(&_IOTXStaking.CallOpts)
}

// AccountedManagerReward is a free data retrieval call binding the contract method 0x48c26fce.
//
// Solidity: function accountedManagerReward() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) AccountedManagerReward() (*big.Int, error) {
	return _IOTXStaking.Contract.AccountedManagerReward(&_IOTXStaking.CallOpts)
}

// AccountedUserReward is a free data retrieval call binding the contract method 0x5ff35b5c.
//
// Solidity: function accountedUserReward() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) AccountedUserReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "accountedUserReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountedUserReward is a free data retrieval call binding the contract method 0x5ff35b5c.
//
// Solidity: function accountedUserReward() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) AccountedUserReward() (*big.Int, error) {
	return _IOTXStaking.Contract.AccountedUserReward(&_IOTXStaking.CallOpts)
}

// AccountedUserReward is a free data retrieval call binding the contract method 0x5ff35b5c.
//
// Solidity: function accountedUserReward() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) AccountedUserReward() (*big.Int, error) {
	return _IOTXStaking.Contract.AccountedUserReward(&_IOTXStaking.CallOpts)
}

// CommonRatio is a free data retrieval call binding the contract method 0x3dd419d0.
//
// Solidity: function commonRatio() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) CommonRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "commonRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommonRatio is a free data retrieval call binding the contract method 0x3dd419d0.
//
// Solidity: function commonRatio() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) CommonRatio() (*big.Int, error) {
	return _IOTXStaking.Contract.CommonRatio(&_IOTXStaking.CallOpts)
}

// CommonRatio is a free data retrieval call binding the contract method 0x3dd419d0.
//
// Solidity: function commonRatio() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) CommonRatio() (*big.Int, error) {
	return _IOTXStaking.Contract.CommonRatio(&_IOTXStaking.CallOpts)
}

// CurrentReserve is a free data retrieval call binding the contract method 0x2e12007c.
//
// Solidity: function currentReserve() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) CurrentReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "currentReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentReserve is a free data retrieval call binding the contract method 0x2e12007c.
//
// Solidity: function currentReserve() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) CurrentReserve() (*big.Int, error) {
	return _IOTXStaking.Contract.CurrentReserve(&_IOTXStaking.CallOpts)
}

// CurrentReserve is a free data retrieval call binding the contract method 0x2e12007c.
//
// Solidity: function currentReserve() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) CurrentReserve() (*big.Int, error) {
	return _IOTXStaking.Contract.CurrentReserve(&_IOTXStaking.CallOpts)
}

// ExchangeRatio is a free data retrieval call binding the contract method 0x4006ccc5.
//
// Solidity: function exchangeRatio() view returns(uint256 ratio)
func (_IOTXStaking *IOTXStakingCaller) ExchangeRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "exchangeRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRatio is a free data retrieval call binding the contract method 0x4006ccc5.
//
// Solidity: function exchangeRatio() view returns(uint256 ratio)
func (_IOTXStaking *IOTXStakingSession) ExchangeRatio() (*big.Int, error) {
	return _IOTXStaking.Contract.ExchangeRatio(&_IOTXStaking.CallOpts)
}

// ExchangeRatio is a free data retrieval call binding the contract method 0x4006ccc5.
//
// Solidity: function exchangeRatio() view returns(uint256 ratio)
func (_IOTXStaking *IOTXStakingCallerSession) ExchangeRatio() (*big.Int, error) {
	return _IOTXStaking.Contract.ExchangeRatio(&_IOTXStaking.CallOpts)
}

// GetAccountedBalance is a free data retrieval call binding the contract method 0x33e5761f.
//
// Solidity: function getAccountedBalance() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) GetAccountedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getAccountedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountedBalance is a free data retrieval call binding the contract method 0x33e5761f.
//
// Solidity: function getAccountedBalance() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) GetAccountedBalance() (*big.Int, error) {
	return _IOTXStaking.Contract.GetAccountedBalance(&_IOTXStaking.CallOpts)
}

// GetAccountedBalance is a free data retrieval call binding the contract method 0x33e5761f.
//
// Solidity: function getAccountedBalance() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) GetAccountedBalance() (*big.Int, error) {
	return _IOTXStaking.Contract.GetAccountedBalance(&_IOTXStaking.CallOpts)
}

// GetGlobalDelegate is a free data retrieval call binding the contract method 0x79df5f41.
//
// Solidity: function getGlobalDelegate() view returns(address)
func (_IOTXStaking *IOTXStakingCaller) GetGlobalDelegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getGlobalDelegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGlobalDelegate is a free data retrieval call binding the contract method 0x79df5f41.
//
// Solidity: function getGlobalDelegate() view returns(address)
func (_IOTXStaking *IOTXStakingSession) GetGlobalDelegate() (common.Address, error) {
	return _IOTXStaking.Contract.GetGlobalDelegate(&_IOTXStaking.CallOpts)
}

// GetGlobalDelegate is a free data retrieval call binding the contract method 0x79df5f41.
//
// Solidity: function getGlobalDelegate() view returns(address)
func (_IOTXStaking *IOTXStakingCallerSession) GetGlobalDelegate() (common.Address, error) {
	return _IOTXStaking.Contract.GetGlobalDelegate(&_IOTXStaking.CallOpts)
}

// GetManagerFeeShares is a free data retrieval call binding the contract method 0x2a67bbe5.
//
// Solidity: function getManagerFeeShares() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) GetManagerFeeShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getManagerFeeShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetManagerFeeShares is a free data retrieval call binding the contract method 0x2a67bbe5.
//
// Solidity: function getManagerFeeShares() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) GetManagerFeeShares() (*big.Int, error) {
	return _IOTXStaking.Contract.GetManagerFeeShares(&_IOTXStaking.CallOpts)
}

// GetManagerFeeShares is a free data retrieval call binding the contract method 0x2a67bbe5.
//
// Solidity: function getManagerFeeShares() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) GetManagerFeeShares() (*big.Int, error) {
	return _IOTXStaking.Contract.GetManagerFeeShares(&_IOTXStaking.CallOpts)
}

// GetManagerReward is a free data retrieval call binding the contract method 0x865112c1.
//
// Solidity: function getManagerReward() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) GetManagerReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getManagerReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetManagerReward is a free data retrieval call binding the contract method 0x865112c1.
//
// Solidity: function getManagerReward() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) GetManagerReward() (*big.Int, error) {
	return _IOTXStaking.Contract.GetManagerReward(&_IOTXStaking.CallOpts)
}

// GetManagerReward is a free data retrieval call binding the contract method 0x865112c1.
//
// Solidity: function getManagerReward() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) GetManagerReward() (*big.Int, error) {
	return _IOTXStaking.Contract.GetManagerReward(&_IOTXStaking.CallOpts)
}

// GetRedeemAmountBase is a free data retrieval call binding the contract method 0xf2be86b3.
//
// Solidity: function getRedeemAmountBase() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) GetRedeemAmountBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getRedeemAmountBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedeemAmountBase is a free data retrieval call binding the contract method 0xf2be86b3.
//
// Solidity: function getRedeemAmountBase() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) GetRedeemAmountBase() (*big.Int, error) {
	return _IOTXStaking.Contract.GetRedeemAmountBase(&_IOTXStaking.CallOpts)
}

// GetRedeemAmountBase is a free data retrieval call binding the contract method 0xf2be86b3.
//
// Solidity: function getRedeemAmountBase() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) GetRedeemAmountBase() (*big.Int, error) {
	return _IOTXStaking.Contract.GetRedeemAmountBase(&_IOTXStaking.CallOpts)
}

// GetRedeemedTokenCount is a free data retrieval call binding the contract method 0xefa9058b.
//
// Solidity: function getRedeemedTokenCount() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) GetRedeemedTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getRedeemedTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedeemedTokenCount is a free data retrieval call binding the contract method 0xefa9058b.
//
// Solidity: function getRedeemedTokenCount() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) GetRedeemedTokenCount() (*big.Int, error) {
	return _IOTXStaking.Contract.GetRedeemedTokenCount(&_IOTXStaking.CallOpts)
}

// GetRedeemedTokenCount is a free data retrieval call binding the contract method 0xefa9058b.
//
// Solidity: function getRedeemedTokenCount() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) GetRedeemedTokenCount() (*big.Int, error) {
	return _IOTXStaking.Contract.GetRedeemedTokenCount(&_IOTXStaking.CallOpts)
}

// GetRedeemedTokenIds is a free data retrieval call binding the contract method 0xc081c90b.
//
// Solidity: function getRedeemedTokenIds() view returns(uint256[])
func (_IOTXStaking *IOTXStakingCaller) GetRedeemedTokenIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getRedeemedTokenIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetRedeemedTokenIds is a free data retrieval call binding the contract method 0xc081c90b.
//
// Solidity: function getRedeemedTokenIds() view returns(uint256[])
func (_IOTXStaking *IOTXStakingSession) GetRedeemedTokenIds() ([]*big.Int, error) {
	return _IOTXStaking.Contract.GetRedeemedTokenIds(&_IOTXStaking.CallOpts)
}

// GetRedeemedTokenIds is a free data retrieval call binding the contract method 0xc081c90b.
//
// Solidity: function getRedeemedTokenIds() view returns(uint256[])
func (_IOTXStaking *IOTXStakingCallerSession) GetRedeemedTokenIds() ([]*big.Int, error) {
	return _IOTXStaking.Contract.GetRedeemedTokenIds(&_IOTXStaking.CallOpts)
}

// GetRedeemedTokenIdsSlice is a free data retrieval call binding the contract method 0x3d11446c.
//
// Solidity: function getRedeemedTokenIdsSlice(uint256 i, uint256 j) view returns(uint256[])
func (_IOTXStaking *IOTXStakingCaller) GetRedeemedTokenIdsSlice(opts *bind.CallOpts, i *big.Int, j *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getRedeemedTokenIdsSlice", i, j)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetRedeemedTokenIdsSlice is a free data retrieval call binding the contract method 0x3d11446c.
//
// Solidity: function getRedeemedTokenIdsSlice(uint256 i, uint256 j) view returns(uint256[])
func (_IOTXStaking *IOTXStakingSession) GetRedeemedTokenIdsSlice(i *big.Int, j *big.Int) ([]*big.Int, error) {
	return _IOTXStaking.Contract.GetRedeemedTokenIdsSlice(&_IOTXStaking.CallOpts, i, j)
}

// GetRedeemedTokenIdsSlice is a free data retrieval call binding the contract method 0x3d11446c.
//
// Solidity: function getRedeemedTokenIdsSlice(uint256 i, uint256 j) view returns(uint256[])
func (_IOTXStaking *IOTXStakingCallerSession) GetRedeemedTokenIdsSlice(i *big.Int, j *big.Int) ([]*big.Int, error) {
	return _IOTXStaking.Contract.GetRedeemedTokenIdsSlice(&_IOTXStaking.CallOpts, i, j)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IOTXStaking *IOTXStakingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IOTXStaking *IOTXStakingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IOTXStaking.Contract.GetRoleAdmin(&_IOTXStaking.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IOTXStaking *IOTXStakingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IOTXStaking.Contract.GetRoleAdmin(&_IOTXStaking.CallOpts, role)
}

// GetStakeAmounts is a free data retrieval call binding the contract method 0x64d5bae4.
//
// Solidity: function getStakeAmounts() view returns(uint256[])
func (_IOTXStaking *IOTXStakingCaller) GetStakeAmounts(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getStakeAmounts")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetStakeAmounts is a free data retrieval call binding the contract method 0x64d5bae4.
//
// Solidity: function getStakeAmounts() view returns(uint256[])
func (_IOTXStaking *IOTXStakingSession) GetStakeAmounts() ([]*big.Int, error) {
	return _IOTXStaking.Contract.GetStakeAmounts(&_IOTXStaking.CallOpts)
}

// GetStakeAmounts is a free data retrieval call binding the contract method 0x64d5bae4.
//
// Solidity: function getStakeAmounts() view returns(uint256[])
func (_IOTXStaking *IOTXStakingCallerSession) GetStakeAmounts() ([]*big.Int, error) {
	return _IOTXStaking.Contract.GetStakeAmounts(&_IOTXStaking.CallOpts)
}

// GetStakeDuration is a free data retrieval call binding the contract method 0x9f1bfc6c.
//
// Solidity: function getStakeDuration() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) GetStakeDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getStakeDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeDuration is a free data retrieval call binding the contract method 0x9f1bfc6c.
//
// Solidity: function getStakeDuration() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) GetStakeDuration() (*big.Int, error) {
	return _IOTXStaking.Contract.GetStakeDuration(&_IOTXStaking.CallOpts)
}

// GetStakeDuration is a free data retrieval call binding the contract method 0x9f1bfc6c.
//
// Solidity: function getStakeDuration() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) GetStakeDuration() (*big.Int, error) {
	return _IOTXStaking.Contract.GetStakeDuration(&_IOTXStaking.CallOpts)
}

// GetStakedTokenCount is a free data retrieval call binding the contract method 0x2291f70a.
//
// Solidity: function getStakedTokenCount(uint256 tokenQueueIndex) view returns(uint256 count)
func (_IOTXStaking *IOTXStakingCaller) GetStakedTokenCount(opts *bind.CallOpts, tokenQueueIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getStakedTokenCount", tokenQueueIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakedTokenCount is a free data retrieval call binding the contract method 0x2291f70a.
//
// Solidity: function getStakedTokenCount(uint256 tokenQueueIndex) view returns(uint256 count)
func (_IOTXStaking *IOTXStakingSession) GetStakedTokenCount(tokenQueueIndex *big.Int) (*big.Int, error) {
	return _IOTXStaking.Contract.GetStakedTokenCount(&_IOTXStaking.CallOpts, tokenQueueIndex)
}

// GetStakedTokenCount is a free data retrieval call binding the contract method 0x2291f70a.
//
// Solidity: function getStakedTokenCount(uint256 tokenQueueIndex) view returns(uint256 count)
func (_IOTXStaking *IOTXStakingCallerSession) GetStakedTokenCount(tokenQueueIndex *big.Int) (*big.Int, error) {
	return _IOTXStaking.Contract.GetStakedTokenCount(&_IOTXStaking.CallOpts, tokenQueueIndex)
}

// GetStakedTokenIds is a free data retrieval call binding the contract method 0xa59ec997.
//
// Solidity: function getStakedTokenIds(uint256 tokenQueueIndex) view returns(uint256[])
func (_IOTXStaking *IOTXStakingCaller) GetStakedTokenIds(opts *bind.CallOpts, tokenQueueIndex *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getStakedTokenIds", tokenQueueIndex)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetStakedTokenIds is a free data retrieval call binding the contract method 0xa59ec997.
//
// Solidity: function getStakedTokenIds(uint256 tokenQueueIndex) view returns(uint256[])
func (_IOTXStaking *IOTXStakingSession) GetStakedTokenIds(tokenQueueIndex *big.Int) ([]*big.Int, error) {
	return _IOTXStaking.Contract.GetStakedTokenIds(&_IOTXStaking.CallOpts, tokenQueueIndex)
}

// GetStakedTokenIds is a free data retrieval call binding the contract method 0xa59ec997.
//
// Solidity: function getStakedTokenIds(uint256 tokenQueueIndex) view returns(uint256[])
func (_IOTXStaking *IOTXStakingCallerSession) GetStakedTokenIds(tokenQueueIndex *big.Int) ([]*big.Int, error) {
	return _IOTXStaking.Contract.GetStakedTokenIds(&_IOTXStaking.CallOpts, tokenQueueIndex)
}

// GetStakedTokenIdsSlice is a free data retrieval call binding the contract method 0x654869d4.
//
// Solidity: function getStakedTokenIdsSlice(uint256 tokenQueueIndex, uint256 i, uint256 j) view returns(uint256[])
func (_IOTXStaking *IOTXStakingCaller) GetStakedTokenIdsSlice(opts *bind.CallOpts, tokenQueueIndex *big.Int, i *big.Int, j *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getStakedTokenIdsSlice", tokenQueueIndex, i, j)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetStakedTokenIdsSlice is a free data retrieval call binding the contract method 0x654869d4.
//
// Solidity: function getStakedTokenIdsSlice(uint256 tokenQueueIndex, uint256 i, uint256 j) view returns(uint256[])
func (_IOTXStaking *IOTXStakingSession) GetStakedTokenIdsSlice(tokenQueueIndex *big.Int, i *big.Int, j *big.Int) ([]*big.Int, error) {
	return _IOTXStaking.Contract.GetStakedTokenIdsSlice(&_IOTXStaking.CallOpts, tokenQueueIndex, i, j)
}

// GetStakedTokenIdsSlice is a free data retrieval call binding the contract method 0x654869d4.
//
// Solidity: function getStakedTokenIdsSlice(uint256 tokenQueueIndex, uint256 i, uint256 j) view returns(uint256[])
func (_IOTXStaking *IOTXStakingCallerSession) GetStakedTokenIdsSlice(tokenQueueIndex *big.Int, i *big.Int, j *big.Int) ([]*big.Int, error) {
	return _IOTXStaking.Contract.GetStakedTokenIdsSlice(&_IOTXStaking.CallOpts, tokenQueueIndex, i, j)
}

// GetTokenId is a free data retrieval call binding the contract method 0x3112de9a.
//
// Solidity: function getTokenId(uint256 tokenQueueIndex, uint256 tokenIndex) view returns(uint256 tokenId)
func (_IOTXStaking *IOTXStakingCaller) GetTokenId(opts *bind.CallOpts, tokenQueueIndex *big.Int, tokenIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getTokenId", tokenQueueIndex, tokenIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenId is a free data retrieval call binding the contract method 0x3112de9a.
//
// Solidity: function getTokenId(uint256 tokenQueueIndex, uint256 tokenIndex) view returns(uint256 tokenId)
func (_IOTXStaking *IOTXStakingSession) GetTokenId(tokenQueueIndex *big.Int, tokenIndex *big.Int) (*big.Int, error) {
	return _IOTXStaking.Contract.GetTokenId(&_IOTXStaking.CallOpts, tokenQueueIndex, tokenIndex)
}

// GetTokenId is a free data retrieval call binding the contract method 0x3112de9a.
//
// Solidity: function getTokenId(uint256 tokenQueueIndex, uint256 tokenIndex) view returns(uint256 tokenId)
func (_IOTXStaking *IOTXStakingCallerSession) GetTokenId(tokenQueueIndex *big.Int, tokenIndex *big.Int) (*big.Int, error) {
	return _IOTXStaking.Contract.GetTokenId(&_IOTXStaking.CallOpts, tokenQueueIndex, tokenIndex)
}

// GetTokenQueueLength is a free data retrieval call binding the contract method 0x797fa5c2.
//
// Solidity: function getTokenQueueLength(uint256 tokenQueueIndex) view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) GetTokenQueueLength(opts *bind.CallOpts, tokenQueueIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getTokenQueueLength", tokenQueueIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenQueueLength is a free data retrieval call binding the contract method 0x797fa5c2.
//
// Solidity: function getTokenQueueLength(uint256 tokenQueueIndex) view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) GetTokenQueueLength(tokenQueueIndex *big.Int) (*big.Int, error) {
	return _IOTXStaking.Contract.GetTokenQueueLength(&_IOTXStaking.CallOpts, tokenQueueIndex)
}

// GetTokenQueueLength is a free data retrieval call binding the contract method 0x797fa5c2.
//
// Solidity: function getTokenQueueLength(uint256 tokenQueueIndex) view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) GetTokenQueueLength(tokenQueueIndex *big.Int) (*big.Int, error) {
	return _IOTXStaking.Contract.GetTokenQueueLength(&_IOTXStaking.CallOpts, tokenQueueIndex)
}

// GetTotalPending is a free data retrieval call binding the contract method 0x3bd05400.
//
// Solidity: function getTotalPending() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) GetTotalPending(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getTotalPending")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPending is a free data retrieval call binding the contract method 0x3bd05400.
//
// Solidity: function getTotalPending() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) GetTotalPending() (*big.Int, error) {
	return _IOTXStaking.Contract.GetTotalPending(&_IOTXStaking.CallOpts)
}

// GetTotalPending is a free data retrieval call binding the contract method 0x3bd05400.
//
// Solidity: function getTotalPending() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) GetTotalPending() (*big.Int, error) {
	return _IOTXStaking.Contract.GetTotalPending(&_IOTXStaking.CallOpts)
}

// GetTotalStaked is a free data retrieval call binding the contract method 0x0917e776.
//
// Solidity: function getTotalStaked() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) GetTotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getTotalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStaked is a free data retrieval call binding the contract method 0x0917e776.
//
// Solidity: function getTotalStaked() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) GetTotalStaked() (*big.Int, error) {
	return _IOTXStaking.Contract.GetTotalStaked(&_IOTXStaking.CallOpts)
}

// GetTotalStaked is a free data retrieval call binding the contract method 0x0917e776.
//
// Solidity: function getTotalStaked() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) GetTotalStaked() (*big.Int, error) {
	return _IOTXStaking.Contract.GetTotalStaked(&_IOTXStaking.CallOpts)
}

// GetUserReward is a free data retrieval call binding the contract method 0xd793d83e.
//
// Solidity: function getUserReward() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) GetUserReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "getUserReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserReward is a free data retrieval call binding the contract method 0xd793d83e.
//
// Solidity: function getUserReward() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) GetUserReward() (*big.Int, error) {
	return _IOTXStaking.Contract.GetUserReward(&_IOTXStaking.CallOpts)
}

// GetUserReward is a free data retrieval call binding the contract method 0xd793d83e.
//
// Solidity: function getUserReward() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) GetUserReward() (*big.Int, error) {
	return _IOTXStaking.Contract.GetUserReward(&_IOTXStaking.CallOpts)
}

// GlobalDelegate is a free data retrieval call binding the contract method 0x11e4bc2d.
//
// Solidity: function globalDelegate() view returns(address)
func (_IOTXStaking *IOTXStakingCaller) GlobalDelegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "globalDelegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalDelegate is a free data retrieval call binding the contract method 0x11e4bc2d.
//
// Solidity: function globalDelegate() view returns(address)
func (_IOTXStaking *IOTXStakingSession) GlobalDelegate() (common.Address, error) {
	return _IOTXStaking.Contract.GlobalDelegate(&_IOTXStaking.CallOpts)
}

// GlobalDelegate is a free data retrieval call binding the contract method 0x11e4bc2d.
//
// Solidity: function globalDelegate() view returns(address)
func (_IOTXStaking *IOTXStakingCallerSession) GlobalDelegate() (common.Address, error) {
	return _IOTXStaking.Contract.GlobalDelegate(&_IOTXStaking.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IOTXStaking *IOTXStakingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IOTXStaking *IOTXStakingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IOTXStaking.Contract.HasRole(&_IOTXStaking.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IOTXStaking *IOTXStakingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IOTXStaking.Contract.HasRole(&_IOTXStaking.CallOpts, role, account)
}

// IotxClear is a free data retrieval call binding the contract method 0x5b3f86b4.
//
// Solidity: function iotxClear() view returns(address)
func (_IOTXStaking *IOTXStakingCaller) IotxClear(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "iotxClear")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IotxClear is a free data retrieval call binding the contract method 0x5b3f86b4.
//
// Solidity: function iotxClear() view returns(address)
func (_IOTXStaking *IOTXStakingSession) IotxClear() (common.Address, error) {
	return _IOTXStaking.Contract.IotxClear(&_IOTXStaking.CallOpts)
}

// IotxClear is a free data retrieval call binding the contract method 0x5b3f86b4.
//
// Solidity: function iotxClear() view returns(address)
func (_IOTXStaking *IOTXStakingCallerSession) IotxClear() (common.Address, error) {
	return _IOTXStaking.Contract.IotxClear(&_IOTXStaking.CallOpts)
}

// ManagerFeeShares is a free data retrieval call binding the contract method 0xf6af7527.
//
// Solidity: function managerFeeShares() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) ManagerFeeShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "managerFeeShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerFeeShares is a free data retrieval call binding the contract method 0xf6af7527.
//
// Solidity: function managerFeeShares() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) ManagerFeeShares() (*big.Int, error) {
	return _IOTXStaking.Contract.ManagerFeeShares(&_IOTXStaking.CallOpts)
}

// ManagerFeeShares is a free data retrieval call binding the contract method 0xf6af7527.
//
// Solidity: function managerFeeShares() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) ManagerFeeShares() (*big.Int, error) {
	return _IOTXStaking.Contract.ManagerFeeShares(&_IOTXStaking.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_IOTXStaking *IOTXStakingCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_IOTXStaking *IOTXStakingSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _IOTXStaking.Contract.OnERC721Received(&_IOTXStaking.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_IOTXStaking *IOTXStakingCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _IOTXStaking.Contract.OnERC721Received(&_IOTXStaking.CallOpts, arg0, arg1, arg2, arg3)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IOTXStaking *IOTXStakingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IOTXStaking *IOTXStakingSession) Paused() (bool, error) {
	return _IOTXStaking.Contract.Paused(&_IOTXStaking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IOTXStaking *IOTXStakingCallerSession) Paused() (bool, error) {
	return _IOTXStaking.Contract.Paused(&_IOTXStaking.CallOpts)
}

// RedeemAmountBase is a free data retrieval call binding the contract method 0x37f52345.
//
// Solidity: function redeemAmountBase() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) RedeemAmountBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "redeemAmountBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemAmountBase is a free data retrieval call binding the contract method 0x37f52345.
//
// Solidity: function redeemAmountBase() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) RedeemAmountBase() (*big.Int, error) {
	return _IOTXStaking.Contract.RedeemAmountBase(&_IOTXStaking.CallOpts)
}

// RedeemAmountBase is a free data retrieval call binding the contract method 0x37f52345.
//
// Solidity: function redeemAmountBase() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) RedeemAmountBase() (*big.Int, error) {
	return _IOTXStaking.Contract.RedeemAmountBase(&_IOTXStaking.CallOpts)
}

// RedeemedTokenCount is a free data retrieval call binding the contract method 0xb9bafb72.
//
// Solidity: function redeemedTokenCount() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) RedeemedTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "redeemedTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemedTokenCount is a free data retrieval call binding the contract method 0xb9bafb72.
//
// Solidity: function redeemedTokenCount() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) RedeemedTokenCount() (*big.Int, error) {
	return _IOTXStaking.Contract.RedeemedTokenCount(&_IOTXStaking.CallOpts)
}

// RedeemedTokenCount is a free data retrieval call binding the contract method 0xb9bafb72.
//
// Solidity: function redeemedTokenCount() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) RedeemedTokenCount() (*big.Int, error) {
	return _IOTXStaking.Contract.RedeemedTokenCount(&_IOTXStaking.CallOpts)
}

// SequenceLength is a free data retrieval call binding the contract method 0xc030a6b9.
//
// Solidity: function sequenceLength() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) SequenceLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "sequenceLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequenceLength is a free data retrieval call binding the contract method 0xc030a6b9.
//
// Solidity: function sequenceLength() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) SequenceLength() (*big.Int, error) {
	return _IOTXStaking.Contract.SequenceLength(&_IOTXStaking.CallOpts)
}

// SequenceLength is a free data retrieval call binding the contract method 0xc030a6b9.
//
// Solidity: function sequenceLength() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) SequenceLength() (*big.Int, error) {
	return _IOTXStaking.Contract.SequenceLength(&_IOTXStaking.CallOpts)
}

// StakeDuration is a free data retrieval call binding the contract method 0x40f02ab6.
//
// Solidity: function stakeDuration() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) StakeDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "stakeDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeDuration is a free data retrieval call binding the contract method 0x40f02ab6.
//
// Solidity: function stakeDuration() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) StakeDuration() (*big.Int, error) {
	return _IOTXStaking.Contract.StakeDuration(&_IOTXStaking.CallOpts)
}

// StakeDuration is a free data retrieval call binding the contract method 0x40f02ab6.
//
// Solidity: function stakeDuration() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) StakeDuration() (*big.Int, error) {
	return _IOTXStaking.Contract.StakeDuration(&_IOTXStaking.CallOpts)
}

// StartAmount is a free data retrieval call binding the contract method 0x0f767872.
//
// Solidity: function startAmount() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) StartAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "startAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartAmount is a free data retrieval call binding the contract method 0x0f767872.
//
// Solidity: function startAmount() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) StartAmount() (*big.Int, error) {
	return _IOTXStaking.Contract.StartAmount(&_IOTXStaking.CallOpts)
}

// StartAmount is a free data retrieval call binding the contract method 0x0f767872.
//
// Solidity: function startAmount() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) StartAmount() (*big.Int, error) {
	return _IOTXStaking.Contract.StartAmount(&_IOTXStaking.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IOTXStaking *IOTXStakingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IOTXStaking *IOTXStakingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IOTXStaking.Contract.SupportsInterface(&_IOTXStaking.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IOTXStaking *IOTXStakingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IOTXStaking.Contract.SupportsInterface(&_IOTXStaking.CallOpts, interfaceId)
}

// SystemStaking is a free data retrieval call binding the contract method 0xa5e35ef6.
//
// Solidity: function systemStaking() view returns(address)
func (_IOTXStaking *IOTXStakingCaller) SystemStaking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "systemStaking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemStaking is a free data retrieval call binding the contract method 0xa5e35ef6.
//
// Solidity: function systemStaking() view returns(address)
func (_IOTXStaking *IOTXStakingSession) SystemStaking() (common.Address, error) {
	return _IOTXStaking.Contract.SystemStaking(&_IOTXStaking.CallOpts)
}

// SystemStaking is a free data retrieval call binding the contract method 0xa5e35ef6.
//
// Solidity: function systemStaking() view returns(address)
func (_IOTXStaking *IOTXStakingCallerSession) SystemStaking() (common.Address, error) {
	return _IOTXStaking.Contract.SystemStaking(&_IOTXStaking.CallOpts)
}

// TotalPending is a free data retrieval call binding the contract method 0x3f90916a.
//
// Solidity: function totalPending() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) TotalPending(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "totalPending")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPending is a free data retrieval call binding the contract method 0x3f90916a.
//
// Solidity: function totalPending() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) TotalPending() (*big.Int, error) {
	return _IOTXStaking.Contract.TotalPending(&_IOTXStaking.CallOpts)
}

// TotalPending is a free data retrieval call binding the contract method 0x3f90916a.
//
// Solidity: function totalPending() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) TotalPending() (*big.Int, error) {
	return _IOTXStaking.Contract.TotalPending(&_IOTXStaking.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_IOTXStaking *IOTXStakingCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_IOTXStaking *IOTXStakingSession) TotalStaked() (*big.Int, error) {
	return _IOTXStaking.Contract.TotalStaked(&_IOTXStaking.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_IOTXStaking *IOTXStakingCallerSession) TotalStaked() (*big.Int, error) {
	return _IOTXStaking.Contract.TotalStaked(&_IOTXStaking.CallOpts)
}

// UniIOTX is a free data retrieval call binding the contract method 0x35b83d28.
//
// Solidity: function uniIOTX() view returns(address)
func (_IOTXStaking *IOTXStakingCaller) UniIOTX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IOTXStaking.contract.Call(opts, &out, "uniIOTX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniIOTX is a free data retrieval call binding the contract method 0x35b83d28.
//
// Solidity: function uniIOTX() view returns(address)
func (_IOTXStaking *IOTXStakingSession) UniIOTX() (common.Address, error) {
	return _IOTXStaking.Contract.UniIOTX(&_IOTXStaking.CallOpts)
}

// UniIOTX is a free data retrieval call binding the contract method 0x35b83d28.
//
// Solidity: function uniIOTX() view returns(address)
func (_IOTXStaking *IOTXStakingCallerSession) UniIOTX() (common.Address, error) {
	return _IOTXStaking.Contract.UniIOTX(&_IOTXStaking.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 deadline) payable returns(uint256 minted)
func (_IOTXStaking *IOTXStakingTransactor) Deposit(opts *bind.TransactOpts, deadline *big.Int) (*types.Transaction, error) {
	return _IOTXStaking.contract.Transact(opts, "deposit", deadline)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 deadline) payable returns(uint256 minted)
func (_IOTXStaking *IOTXStakingSession) Deposit(deadline *big.Int) (*types.Transaction, error) {
	return _IOTXStaking.Contract.Deposit(&_IOTXStaking.TransactOpts, deadline)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 deadline) payable returns(uint256 minted)
func (_IOTXStaking *IOTXStakingTransactorSession) Deposit(deadline *big.Int) (*types.Transaction, error) {
	return _IOTXStaking.Contract.Deposit(&_IOTXStaking.TransactOpts, deadline)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IOTXStaking *IOTXStakingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXStaking.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IOTXStaking *IOTXStakingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXStaking.Contract.GrantRole(&_IOTXStaking.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IOTXStaking *IOTXStakingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXStaking.Contract.GrantRole(&_IOTXStaking.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc843e81c.
//
// Solidity: function initialize(address _systemStaking, address _uniIOTX, address _iotxClear, address _oracle, uint256 _startAmount, uint256 _commonRatio, uint256 _sequenceLength, uint256 _stakeDuration) returns()
func (_IOTXStaking *IOTXStakingTransactor) Initialize(opts *bind.TransactOpts, _systemStaking common.Address, _uniIOTX common.Address, _iotxClear common.Address, _oracle common.Address, _startAmount *big.Int, _commonRatio *big.Int, _sequenceLength *big.Int, _stakeDuration *big.Int) (*types.Transaction, error) {
	return _IOTXStaking.contract.Transact(opts, "initialize", _systemStaking, _uniIOTX, _iotxClear, _oracle, _startAmount, _commonRatio, _sequenceLength, _stakeDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xc843e81c.
//
// Solidity: function initialize(address _systemStaking, address _uniIOTX, address _iotxClear, address _oracle, uint256 _startAmount, uint256 _commonRatio, uint256 _sequenceLength, uint256 _stakeDuration) returns()
func (_IOTXStaking *IOTXStakingSession) Initialize(_systemStaking common.Address, _uniIOTX common.Address, _iotxClear common.Address, _oracle common.Address, _startAmount *big.Int, _commonRatio *big.Int, _sequenceLength *big.Int, _stakeDuration *big.Int) (*types.Transaction, error) {
	return _IOTXStaking.Contract.Initialize(&_IOTXStaking.TransactOpts, _systemStaking, _uniIOTX, _iotxClear, _oracle, _startAmount, _commonRatio, _sequenceLength, _stakeDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xc843e81c.
//
// Solidity: function initialize(address _systemStaking, address _uniIOTX, address _iotxClear, address _oracle, uint256 _startAmount, uint256 _commonRatio, uint256 _sequenceLength, uint256 _stakeDuration) returns()
func (_IOTXStaking *IOTXStakingTransactorSession) Initialize(_systemStaking common.Address, _uniIOTX common.Address, _iotxClear common.Address, _oracle common.Address, _startAmount *big.Int, _commonRatio *big.Int, _sequenceLength *big.Int, _stakeDuration *big.Int) (*types.Transaction, error) {
	return _IOTXStaking.Contract.Initialize(&_IOTXStaking.TransactOpts, _systemStaking, _uniIOTX, _iotxClear, _oracle, _startAmount, _commonRatio, _sequenceLength, _stakeDuration)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IOTXStaking *IOTXStakingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXStaking.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IOTXStaking *IOTXStakingSession) Pause() (*types.Transaction, error) {
	return _IOTXStaking.Contract.Pause(&_IOTXStaking.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IOTXStaking *IOTXStakingTransactorSession) Pause() (*types.Transaction, error) {
	return _IOTXStaking.Contract.Pause(&_IOTXStaking.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x7cbc2373.
//
// Solidity: function redeem(uint256 iotxsToRedeem, uint256 deadline) returns(uint256 burned)
func (_IOTXStaking *IOTXStakingTransactor) Redeem(opts *bind.TransactOpts, iotxsToRedeem *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IOTXStaking.contract.Transact(opts, "redeem", iotxsToRedeem, deadline)
}

// Redeem is a paid mutator transaction binding the contract method 0x7cbc2373.
//
// Solidity: function redeem(uint256 iotxsToRedeem, uint256 deadline) returns(uint256 burned)
func (_IOTXStaking *IOTXStakingSession) Redeem(iotxsToRedeem *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IOTXStaking.Contract.Redeem(&_IOTXStaking.TransactOpts, iotxsToRedeem, deadline)
}

// Redeem is a paid mutator transaction binding the contract method 0x7cbc2373.
//
// Solidity: function redeem(uint256 iotxsToRedeem, uint256 deadline) returns(uint256 burned)
func (_IOTXStaking *IOTXStakingTransactorSession) Redeem(iotxsToRedeem *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IOTXStaking.Contract.Redeem(&_IOTXStaking.TransactOpts, iotxsToRedeem, deadline)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IOTXStaking *IOTXStakingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXStaking.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IOTXStaking *IOTXStakingSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXStaking.Contract.RenounceRole(&_IOTXStaking.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IOTXStaking *IOTXStakingTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXStaking.Contract.RenounceRole(&_IOTXStaking.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IOTXStaking *IOTXStakingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXStaking.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IOTXStaking *IOTXStakingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXStaking.Contract.RevokeRole(&_IOTXStaking.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IOTXStaking *IOTXStakingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IOTXStaking.Contract.RevokeRole(&_IOTXStaking.TransactOpts, role, account)
}

// SetGlobalDelegate is a paid mutator transaction binding the contract method 0x229bf1b6.
//
// Solidity: function setGlobalDelegate(address delegate) returns()
func (_IOTXStaking *IOTXStakingTransactor) SetGlobalDelegate(opts *bind.TransactOpts, delegate common.Address) (*types.Transaction, error) {
	return _IOTXStaking.contract.Transact(opts, "setGlobalDelegate", delegate)
}

// SetGlobalDelegate is a paid mutator transaction binding the contract method 0x229bf1b6.
//
// Solidity: function setGlobalDelegate(address delegate) returns()
func (_IOTXStaking *IOTXStakingSession) SetGlobalDelegate(delegate common.Address) (*types.Transaction, error) {
	return _IOTXStaking.Contract.SetGlobalDelegate(&_IOTXStaking.TransactOpts, delegate)
}

// SetGlobalDelegate is a paid mutator transaction binding the contract method 0x229bf1b6.
//
// Solidity: function setGlobalDelegate(address delegate) returns()
func (_IOTXStaking *IOTXStakingTransactorSession) SetGlobalDelegate(delegate common.Address) (*types.Transaction, error) {
	return _IOTXStaking.Contract.SetGlobalDelegate(&_IOTXStaking.TransactOpts, delegate)
}

// SetManagerFeeShares is a paid mutator transaction binding the contract method 0xbecd0858.
//
// Solidity: function setManagerFeeShares(uint256 shares) returns()
func (_IOTXStaking *IOTXStakingTransactor) SetManagerFeeShares(opts *bind.TransactOpts, shares *big.Int) (*types.Transaction, error) {
	return _IOTXStaking.contract.Transact(opts, "setManagerFeeShares", shares)
}

// SetManagerFeeShares is a paid mutator transaction binding the contract method 0xbecd0858.
//
// Solidity: function setManagerFeeShares(uint256 shares) returns()
func (_IOTXStaking *IOTXStakingSession) SetManagerFeeShares(shares *big.Int) (*types.Transaction, error) {
	return _IOTXStaking.Contract.SetManagerFeeShares(&_IOTXStaking.TransactOpts, shares)
}

// SetManagerFeeShares is a paid mutator transaction binding the contract method 0xbecd0858.
//
// Solidity: function setManagerFeeShares(uint256 shares) returns()
func (_IOTXStaking *IOTXStakingTransactorSession) SetManagerFeeShares(shares *big.Int) (*types.Transaction, error) {
	return _IOTXStaking.Contract.SetManagerFeeShares(&_IOTXStaking.TransactOpts, shares)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_IOTXStaking *IOTXStakingTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXStaking.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_IOTXStaking *IOTXStakingSession) Stake() (*types.Transaction, error) {
	return _IOTXStaking.Contract.Stake(&_IOTXStaking.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_IOTXStaking *IOTXStakingTransactorSession) Stake() (*types.Transaction, error) {
	return _IOTXStaking.Contract.Stake(&_IOTXStaking.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IOTXStaking *IOTXStakingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXStaking.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IOTXStaking *IOTXStakingSession) Unpause() (*types.Transaction, error) {
	return _IOTXStaking.Contract.Unpause(&_IOTXStaking.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IOTXStaking *IOTXStakingTransactorSession) Unpause() (*types.Transaction, error) {
	return _IOTXStaking.Contract.Unpause(&_IOTXStaking.TransactOpts)
}

// UpdateDelegates is a paid mutator transaction binding the contract method 0x2328782f.
//
// Solidity: function updateDelegates(uint256[] tokenIds, address delegate) returns()
func (_IOTXStaking *IOTXStakingTransactor) UpdateDelegates(opts *bind.TransactOpts, tokenIds []*big.Int, delegate common.Address) (*types.Transaction, error) {
	return _IOTXStaking.contract.Transact(opts, "updateDelegates", tokenIds, delegate)
}

// UpdateDelegates is a paid mutator transaction binding the contract method 0x2328782f.
//
// Solidity: function updateDelegates(uint256[] tokenIds, address delegate) returns()
func (_IOTXStaking *IOTXStakingSession) UpdateDelegates(tokenIds []*big.Int, delegate common.Address) (*types.Transaction, error) {
	return _IOTXStaking.Contract.UpdateDelegates(&_IOTXStaking.TransactOpts, tokenIds, delegate)
}

// UpdateDelegates is a paid mutator transaction binding the contract method 0x2328782f.
//
// Solidity: function updateDelegates(uint256[] tokenIds, address delegate) returns()
func (_IOTXStaking *IOTXStakingTransactorSession) UpdateDelegates(tokenIds []*big.Int, delegate common.Address) (*types.Transaction, error) {
	return _IOTXStaking.Contract.UpdateDelegates(&_IOTXStaking.TransactOpts, tokenIds, delegate)
}

// UpdateReward is a paid mutator transaction binding the contract method 0xf36c0a72.
//
// Solidity: function updateReward() returns()
func (_IOTXStaking *IOTXStakingTransactor) UpdateReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXStaking.contract.Transact(opts, "updateReward")
}

// UpdateReward is a paid mutator transaction binding the contract method 0xf36c0a72.
//
// Solidity: function updateReward() returns()
func (_IOTXStaking *IOTXStakingSession) UpdateReward() (*types.Transaction, error) {
	return _IOTXStaking.Contract.UpdateReward(&_IOTXStaking.TransactOpts)
}

// UpdateReward is a paid mutator transaction binding the contract method 0xf36c0a72.
//
// Solidity: function updateReward() returns()
func (_IOTXStaking *IOTXStakingTransactorSession) UpdateReward() (*types.Transaction, error) {
	return _IOTXStaking.Contract.UpdateReward(&_IOTXStaking.TransactOpts)
}

// WithdrawManagerFee is a paid mutator transaction binding the contract method 0xfa0ec2ce.
//
// Solidity: function withdrawManagerFee(uint256 amount, address recipient) returns()
func (_IOTXStaking *IOTXStakingTransactor) WithdrawManagerFee(opts *bind.TransactOpts, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IOTXStaking.contract.Transact(opts, "withdrawManagerFee", amount, recipient)
}

// WithdrawManagerFee is a paid mutator transaction binding the contract method 0xfa0ec2ce.
//
// Solidity: function withdrawManagerFee(uint256 amount, address recipient) returns()
func (_IOTXStaking *IOTXStakingSession) WithdrawManagerFee(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IOTXStaking.Contract.WithdrawManagerFee(&_IOTXStaking.TransactOpts, amount, recipient)
}

// WithdrawManagerFee is a paid mutator transaction binding the contract method 0xfa0ec2ce.
//
// Solidity: function withdrawManagerFee(uint256 amount, address recipient) returns()
func (_IOTXStaking *IOTXStakingTransactorSession) WithdrawManagerFee(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IOTXStaking.Contract.WithdrawManagerFee(&_IOTXStaking.TransactOpts, amount, recipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IOTXStaking *IOTXStakingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOTXStaking.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IOTXStaking *IOTXStakingSession) Receive() (*types.Transaction, error) {
	return _IOTXStaking.Contract.Receive(&_IOTXStaking.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IOTXStaking *IOTXStakingTransactorSession) Receive() (*types.Transaction, error) {
	return _IOTXStaking.Contract.Receive(&_IOTXStaking.TransactOpts)
}

// IOTXStakingDelegatesUpdatedIterator is returned from FilterDelegatesUpdated and is used to iterate over the raw logs and unpacked data for DelegatesUpdated events raised by the IOTXStaking contract.
type IOTXStakingDelegatesUpdatedIterator struct {
	Event *IOTXStakingDelegatesUpdated // Event containing the contract specifics and raw log

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
func (it *IOTXStakingDelegatesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakingDelegatesUpdated)
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
		it.Event = new(IOTXStakingDelegatesUpdated)
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
func (it *IOTXStakingDelegatesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakingDelegatesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakingDelegatesUpdated represents a DelegatesUpdated event raised by the IOTXStaking contract.
type IOTXStakingDelegatesUpdated struct {
	TokenIds []*big.Int
	Delegate common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegatesUpdated is a free log retrieval operation binding the contract event 0x055e970cad70d45557223e4ef35ac45bb162f8fc6f8a5986159f49d7d7fc742b.
//
// Solidity: event DelegatesUpdated(uint256[] tokenIds, address delegate)
func (_IOTXStaking *IOTXStakingFilterer) FilterDelegatesUpdated(opts *bind.FilterOpts) (*IOTXStakingDelegatesUpdatedIterator, error) {

	logs, sub, err := _IOTXStaking.contract.FilterLogs(opts, "DelegatesUpdated")
	if err != nil {
		return nil, err
	}
	return &IOTXStakingDelegatesUpdatedIterator{contract: _IOTXStaking.contract, event: "DelegatesUpdated", logs: logs, sub: sub}, nil
}

// WatchDelegatesUpdated is a free log subscription operation binding the contract event 0x055e970cad70d45557223e4ef35ac45bb162f8fc6f8a5986159f49d7d7fc742b.
//
// Solidity: event DelegatesUpdated(uint256[] tokenIds, address delegate)
func (_IOTXStaking *IOTXStakingFilterer) WatchDelegatesUpdated(opts *bind.WatchOpts, sink chan<- *IOTXStakingDelegatesUpdated) (event.Subscription, error) {

	logs, sub, err := _IOTXStaking.contract.WatchLogs(opts, "DelegatesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakingDelegatesUpdated)
				if err := _IOTXStaking.contract.UnpackLog(event, "DelegatesUpdated", log); err != nil {
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
func (_IOTXStaking *IOTXStakingFilterer) ParseDelegatesUpdated(log types.Log) (*IOTXStakingDelegatesUpdated, error) {
	event := new(IOTXStakingDelegatesUpdated)
	if err := _IOTXStaking.contract.UnpackLog(event, "DelegatesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakingGlobalDelegateSetIterator is returned from FilterGlobalDelegateSet and is used to iterate over the raw logs and unpacked data for GlobalDelegateSet events raised by the IOTXStaking contract.
type IOTXStakingGlobalDelegateSetIterator struct {
	Event *IOTXStakingGlobalDelegateSet // Event containing the contract specifics and raw log

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
func (it *IOTXStakingGlobalDelegateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakingGlobalDelegateSet)
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
		it.Event = new(IOTXStakingGlobalDelegateSet)
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
func (it *IOTXStakingGlobalDelegateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakingGlobalDelegateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakingGlobalDelegateSet represents a GlobalDelegateSet event raised by the IOTXStaking contract.
type IOTXStakingGlobalDelegateSet struct {
	Delegate common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGlobalDelegateSet is a free log retrieval operation binding the contract event 0x19320da70b0c544e76e5dbc3c85129ac6d5d0d5486b4aa42982640788a157d71.
//
// Solidity: event GlobalDelegateSet(address delegate)
func (_IOTXStaking *IOTXStakingFilterer) FilterGlobalDelegateSet(opts *bind.FilterOpts) (*IOTXStakingGlobalDelegateSetIterator, error) {

	logs, sub, err := _IOTXStaking.contract.FilterLogs(opts, "GlobalDelegateSet")
	if err != nil {
		return nil, err
	}
	return &IOTXStakingGlobalDelegateSetIterator{contract: _IOTXStaking.contract, event: "GlobalDelegateSet", logs: logs, sub: sub}, nil
}

// WatchGlobalDelegateSet is a free log subscription operation binding the contract event 0x19320da70b0c544e76e5dbc3c85129ac6d5d0d5486b4aa42982640788a157d71.
//
// Solidity: event GlobalDelegateSet(address delegate)
func (_IOTXStaking *IOTXStakingFilterer) WatchGlobalDelegateSet(opts *bind.WatchOpts, sink chan<- *IOTXStakingGlobalDelegateSet) (event.Subscription, error) {

	logs, sub, err := _IOTXStaking.contract.WatchLogs(opts, "GlobalDelegateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakingGlobalDelegateSet)
				if err := _IOTXStaking.contract.UnpackLog(event, "GlobalDelegateSet", log); err != nil {
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
func (_IOTXStaking *IOTXStakingFilterer) ParseGlobalDelegateSet(log types.Log) (*IOTXStakingGlobalDelegateSet, error) {
	event := new(IOTXStakingGlobalDelegateSet)
	if err := _IOTXStaking.contract.UnpackLog(event, "GlobalDelegateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IOTXStaking contract.
type IOTXStakingInitializedIterator struct {
	Event *IOTXStakingInitialized // Event containing the contract specifics and raw log

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
func (it *IOTXStakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakingInitialized)
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
		it.Event = new(IOTXStakingInitialized)
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
func (it *IOTXStakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakingInitialized represents a Initialized event raised by the IOTXStaking contract.
type IOTXStakingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IOTXStaking *IOTXStakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*IOTXStakingInitializedIterator, error) {

	logs, sub, err := _IOTXStaking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IOTXStakingInitializedIterator{contract: _IOTXStaking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IOTXStaking *IOTXStakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IOTXStakingInitialized) (event.Subscription, error) {

	logs, sub, err := _IOTXStaking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakingInitialized)
				if err := _IOTXStaking.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_IOTXStaking *IOTXStakingFilterer) ParseInitialized(log types.Log) (*IOTXStakingInitialized, error) {
	event := new(IOTXStakingInitialized)
	if err := _IOTXStaking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakingManagerFeeSharesSetIterator is returned from FilterManagerFeeSharesSet and is used to iterate over the raw logs and unpacked data for ManagerFeeSharesSet events raised by the IOTXStaking contract.
type IOTXStakingManagerFeeSharesSetIterator struct {
	Event *IOTXStakingManagerFeeSharesSet // Event containing the contract specifics and raw log

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
func (it *IOTXStakingManagerFeeSharesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakingManagerFeeSharesSet)
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
		it.Event = new(IOTXStakingManagerFeeSharesSet)
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
func (it *IOTXStakingManagerFeeSharesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakingManagerFeeSharesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakingManagerFeeSharesSet represents a ManagerFeeSharesSet event raised by the IOTXStaking contract.
type IOTXStakingManagerFeeSharesSet struct {
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterManagerFeeSharesSet is a free log retrieval operation binding the contract event 0xb6efe094bbb68f7f840c92b10663dc9fdee2f9c7b4b7ba75d88b13e0b685b228.
//
// Solidity: event ManagerFeeSharesSet(uint256 shares)
func (_IOTXStaking *IOTXStakingFilterer) FilterManagerFeeSharesSet(opts *bind.FilterOpts) (*IOTXStakingManagerFeeSharesSetIterator, error) {

	logs, sub, err := _IOTXStaking.contract.FilterLogs(opts, "ManagerFeeSharesSet")
	if err != nil {
		return nil, err
	}
	return &IOTXStakingManagerFeeSharesSetIterator{contract: _IOTXStaking.contract, event: "ManagerFeeSharesSet", logs: logs, sub: sub}, nil
}

// WatchManagerFeeSharesSet is a free log subscription operation binding the contract event 0xb6efe094bbb68f7f840c92b10663dc9fdee2f9c7b4b7ba75d88b13e0b685b228.
//
// Solidity: event ManagerFeeSharesSet(uint256 shares)
func (_IOTXStaking *IOTXStakingFilterer) WatchManagerFeeSharesSet(opts *bind.WatchOpts, sink chan<- *IOTXStakingManagerFeeSharesSet) (event.Subscription, error) {

	logs, sub, err := _IOTXStaking.contract.WatchLogs(opts, "ManagerFeeSharesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakingManagerFeeSharesSet)
				if err := _IOTXStaking.contract.UnpackLog(event, "ManagerFeeSharesSet", log); err != nil {
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
func (_IOTXStaking *IOTXStakingFilterer) ParseManagerFeeSharesSet(log types.Log) (*IOTXStakingManagerFeeSharesSet, error) {
	event := new(IOTXStakingManagerFeeSharesSet)
	if err := _IOTXStaking.contract.UnpackLog(event, "ManagerFeeSharesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakingManagerFeeWithdrawedIterator is returned from FilterManagerFeeWithdrawed and is used to iterate over the raw logs and unpacked data for ManagerFeeWithdrawed events raised by the IOTXStaking contract.
type IOTXStakingManagerFeeWithdrawedIterator struct {
	Event *IOTXStakingManagerFeeWithdrawed // Event containing the contract specifics and raw log

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
func (it *IOTXStakingManagerFeeWithdrawedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakingManagerFeeWithdrawed)
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
		it.Event = new(IOTXStakingManagerFeeWithdrawed)
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
func (it *IOTXStakingManagerFeeWithdrawedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakingManagerFeeWithdrawedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakingManagerFeeWithdrawed represents a ManagerFeeWithdrawed event raised by the IOTXStaking contract.
type IOTXStakingManagerFeeWithdrawed struct {
	Amount    *big.Int
	Minted    *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterManagerFeeWithdrawed is a free log retrieval operation binding the contract event 0xce3d910b07bcccb2bc98f93c5292e13a264d588f34ef3740940f872cea14a7eb.
//
// Solidity: event ManagerFeeWithdrawed(uint256 amount, uint256 minted, address recipient)
func (_IOTXStaking *IOTXStakingFilterer) FilterManagerFeeWithdrawed(opts *bind.FilterOpts) (*IOTXStakingManagerFeeWithdrawedIterator, error) {

	logs, sub, err := _IOTXStaking.contract.FilterLogs(opts, "ManagerFeeWithdrawed")
	if err != nil {
		return nil, err
	}
	return &IOTXStakingManagerFeeWithdrawedIterator{contract: _IOTXStaking.contract, event: "ManagerFeeWithdrawed", logs: logs, sub: sub}, nil
}

// WatchManagerFeeWithdrawed is a free log subscription operation binding the contract event 0xce3d910b07bcccb2bc98f93c5292e13a264d588f34ef3740940f872cea14a7eb.
//
// Solidity: event ManagerFeeWithdrawed(uint256 amount, uint256 minted, address recipient)
func (_IOTXStaking *IOTXStakingFilterer) WatchManagerFeeWithdrawed(opts *bind.WatchOpts, sink chan<- *IOTXStakingManagerFeeWithdrawed) (event.Subscription, error) {

	logs, sub, err := _IOTXStaking.contract.WatchLogs(opts, "ManagerFeeWithdrawed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakingManagerFeeWithdrawed)
				if err := _IOTXStaking.contract.UnpackLog(event, "ManagerFeeWithdrawed", log); err != nil {
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
func (_IOTXStaking *IOTXStakingFilterer) ParseManagerFeeWithdrawed(log types.Log) (*IOTXStakingManagerFeeWithdrawed, error) {
	event := new(IOTXStakingManagerFeeWithdrawed)
	if err := _IOTXStaking.contract.UnpackLog(event, "ManagerFeeWithdrawed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakingMergedIterator is returned from FilterMerged and is used to iterate over the raw logs and unpacked data for Merged events raised by the IOTXStaking contract.
type IOTXStakingMergedIterator struct {
	Event *IOTXStakingMerged // Event containing the contract specifics and raw log

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
func (it *IOTXStakingMergedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakingMerged)
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
		it.Event = new(IOTXStakingMerged)
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
func (it *IOTXStakingMergedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakingMergedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakingMerged represents a Merged event raised by the IOTXStaking contract.
type IOTXStakingMerged struct {
	FromLevel *big.Int
	ToLevel   *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMerged is a free log retrieval operation binding the contract event 0x7032ce2394fdfd8d373af7fe97df414c4002860109755d42748fd91b6a2a0082.
//
// Solidity: event Merged(uint256 fromLevel, uint256 toLevel, uint256 amount)
func (_IOTXStaking *IOTXStakingFilterer) FilterMerged(opts *bind.FilterOpts) (*IOTXStakingMergedIterator, error) {

	logs, sub, err := _IOTXStaking.contract.FilterLogs(opts, "Merged")
	if err != nil {
		return nil, err
	}
	return &IOTXStakingMergedIterator{contract: _IOTXStaking.contract, event: "Merged", logs: logs, sub: sub}, nil
}

// WatchMerged is a free log subscription operation binding the contract event 0x7032ce2394fdfd8d373af7fe97df414c4002860109755d42748fd91b6a2a0082.
//
// Solidity: event Merged(uint256 fromLevel, uint256 toLevel, uint256 amount)
func (_IOTXStaking *IOTXStakingFilterer) WatchMerged(opts *bind.WatchOpts, sink chan<- *IOTXStakingMerged) (event.Subscription, error) {

	logs, sub, err := _IOTXStaking.contract.WatchLogs(opts, "Merged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakingMerged)
				if err := _IOTXStaking.contract.UnpackLog(event, "Merged", log); err != nil {
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

// ParseMerged is a log parse operation binding the contract event 0x7032ce2394fdfd8d373af7fe97df414c4002860109755d42748fd91b6a2a0082.
//
// Solidity: event Merged(uint256 fromLevel, uint256 toLevel, uint256 amount)
func (_IOTXStaking *IOTXStakingFilterer) ParseMerged(log types.Log) (*IOTXStakingMerged, error) {
	event := new(IOTXStakingMerged)
	if err := _IOTXStaking.contract.UnpackLog(event, "Merged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakingMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the IOTXStaking contract.
type IOTXStakingMintedIterator struct {
	Event *IOTXStakingMinted // Event containing the contract specifics and raw log

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
func (it *IOTXStakingMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakingMinted)
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
		it.Event = new(IOTXStakingMinted)
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
func (it *IOTXStakingMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakingMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakingMinted represents a Minted event raised by the IOTXStaking contract.
type IOTXStakingMinted struct {
	User   common.Address
	Minted *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address user, uint256 minted)
func (_IOTXStaking *IOTXStakingFilterer) FilterMinted(opts *bind.FilterOpts) (*IOTXStakingMintedIterator, error) {

	logs, sub, err := _IOTXStaking.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &IOTXStakingMintedIterator{contract: _IOTXStaking.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address user, uint256 minted)
func (_IOTXStaking *IOTXStakingFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *IOTXStakingMinted) (event.Subscription, error) {

	logs, sub, err := _IOTXStaking.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakingMinted)
				if err := _IOTXStaking.contract.UnpackLog(event, "Minted", log); err != nil {
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
func (_IOTXStaking *IOTXStakingFilterer) ParseMinted(log types.Log) (*IOTXStakingMinted, error) {
	event := new(IOTXStakingMinted)
	if err := _IOTXStaking.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IOTXStaking contract.
type IOTXStakingPausedIterator struct {
	Event *IOTXStakingPaused // Event containing the contract specifics and raw log

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
func (it *IOTXStakingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakingPaused)
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
		it.Event = new(IOTXStakingPaused)
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
func (it *IOTXStakingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakingPaused represents a Paused event raised by the IOTXStaking contract.
type IOTXStakingPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IOTXStaking *IOTXStakingFilterer) FilterPaused(opts *bind.FilterOpts) (*IOTXStakingPausedIterator, error) {

	logs, sub, err := _IOTXStaking.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IOTXStakingPausedIterator{contract: _IOTXStaking.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IOTXStaking *IOTXStakingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IOTXStakingPaused) (event.Subscription, error) {

	logs, sub, err := _IOTXStaking.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakingPaused)
				if err := _IOTXStaking.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_IOTXStaking *IOTXStakingFilterer) ParsePaused(log types.Log) (*IOTXStakingPaused, error) {
	event := new(IOTXStakingPaused)
	if err := _IOTXStaking.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakingRedeemedIterator is returned from FilterRedeemed and is used to iterate over the raw logs and unpacked data for Redeemed events raised by the IOTXStaking contract.
type IOTXStakingRedeemedIterator struct {
	Event *IOTXStakingRedeemed // Event containing the contract specifics and raw log

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
func (it *IOTXStakingRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakingRedeemed)
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
		it.Event = new(IOTXStakingRedeemed)
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
func (it *IOTXStakingRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakingRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakingRedeemed represents a Redeemed event raised by the IOTXStaking contract.
type IOTXStakingRedeemed struct {
	User     common.Address
	Burned   *big.Int
	TokenIds []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRedeemed is a free log retrieval operation binding the contract event 0x6b921527f65949c80fab346e1e330b442e353b32e7b29f428d3d0c7766dd9152.
//
// Solidity: event Redeemed(address user, uint256 burned, uint256[] tokenIds)
func (_IOTXStaking *IOTXStakingFilterer) FilterRedeemed(opts *bind.FilterOpts) (*IOTXStakingRedeemedIterator, error) {

	logs, sub, err := _IOTXStaking.contract.FilterLogs(opts, "Redeemed")
	if err != nil {
		return nil, err
	}
	return &IOTXStakingRedeemedIterator{contract: _IOTXStaking.contract, event: "Redeemed", logs: logs, sub: sub}, nil
}

// WatchRedeemed is a free log subscription operation binding the contract event 0x6b921527f65949c80fab346e1e330b442e353b32e7b29f428d3d0c7766dd9152.
//
// Solidity: event Redeemed(address user, uint256 burned, uint256[] tokenIds)
func (_IOTXStaking *IOTXStakingFilterer) WatchRedeemed(opts *bind.WatchOpts, sink chan<- *IOTXStakingRedeemed) (event.Subscription, error) {

	logs, sub, err := _IOTXStaking.contract.WatchLogs(opts, "Redeemed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakingRedeemed)
				if err := _IOTXStaking.contract.UnpackLog(event, "Redeemed", log); err != nil {
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
func (_IOTXStaking *IOTXStakingFilterer) ParseRedeemed(log types.Log) (*IOTXStakingRedeemed, error) {
	event := new(IOTXStakingRedeemed)
	if err := _IOTXStaking.contract.UnpackLog(event, "Redeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakingRewardUpdatedIterator is returned from FilterRewardUpdated and is used to iterate over the raw logs and unpacked data for RewardUpdated events raised by the IOTXStaking contract.
type IOTXStakingRewardUpdatedIterator struct {
	Event *IOTXStakingRewardUpdated // Event containing the contract specifics and raw log

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
func (it *IOTXStakingRewardUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakingRewardUpdated)
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
		it.Event = new(IOTXStakingRewardUpdated)
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
func (it *IOTXStakingRewardUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakingRewardUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakingRewardUpdated represents a RewardUpdated event raised by the IOTXStaking contract.
type IOTXStakingRewardUpdated struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardUpdated is a free log retrieval operation binding the contract event 0xcb94909754d27c309adf4167150f1f7aa04de40b6a0e6bb98b2ae80a2bf438f6.
//
// Solidity: event RewardUpdated(uint256 amount)
func (_IOTXStaking *IOTXStakingFilterer) FilterRewardUpdated(opts *bind.FilterOpts) (*IOTXStakingRewardUpdatedIterator, error) {

	logs, sub, err := _IOTXStaking.contract.FilterLogs(opts, "RewardUpdated")
	if err != nil {
		return nil, err
	}
	return &IOTXStakingRewardUpdatedIterator{contract: _IOTXStaking.contract, event: "RewardUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardUpdated is a free log subscription operation binding the contract event 0xcb94909754d27c309adf4167150f1f7aa04de40b6a0e6bb98b2ae80a2bf438f6.
//
// Solidity: event RewardUpdated(uint256 amount)
func (_IOTXStaking *IOTXStakingFilterer) WatchRewardUpdated(opts *bind.WatchOpts, sink chan<- *IOTXStakingRewardUpdated) (event.Subscription, error) {

	logs, sub, err := _IOTXStaking.contract.WatchLogs(opts, "RewardUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakingRewardUpdated)
				if err := _IOTXStaking.contract.UnpackLog(event, "RewardUpdated", log); err != nil {
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
func (_IOTXStaking *IOTXStakingFilterer) ParseRewardUpdated(log types.Log) (*IOTXStakingRewardUpdated, error) {
	event := new(IOTXStakingRewardUpdated)
	if err := _IOTXStaking.contract.UnpackLog(event, "RewardUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the IOTXStaking contract.
type IOTXStakingRoleAdminChangedIterator struct {
	Event *IOTXStakingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *IOTXStakingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakingRoleAdminChanged)
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
		it.Event = new(IOTXStakingRoleAdminChanged)
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
func (it *IOTXStakingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakingRoleAdminChanged represents a RoleAdminChanged event raised by the IOTXStaking contract.
type IOTXStakingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IOTXStaking *IOTXStakingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*IOTXStakingRoleAdminChangedIterator, error) {

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

	logs, sub, err := _IOTXStaking.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &IOTXStakingRoleAdminChangedIterator{contract: _IOTXStaking.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IOTXStaking *IOTXStakingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *IOTXStakingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _IOTXStaking.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakingRoleAdminChanged)
				if err := _IOTXStaking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_IOTXStaking *IOTXStakingFilterer) ParseRoleAdminChanged(log types.Log) (*IOTXStakingRoleAdminChanged, error) {
	event := new(IOTXStakingRoleAdminChanged)
	if err := _IOTXStaking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IOTXStaking contract.
type IOTXStakingRoleGrantedIterator struct {
	Event *IOTXStakingRoleGranted // Event containing the contract specifics and raw log

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
func (it *IOTXStakingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakingRoleGranted)
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
		it.Event = new(IOTXStakingRoleGranted)
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
func (it *IOTXStakingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakingRoleGranted represents a RoleGranted event raised by the IOTXStaking contract.
type IOTXStakingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IOTXStaking *IOTXStakingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IOTXStakingRoleGrantedIterator, error) {

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

	logs, sub, err := _IOTXStaking.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IOTXStakingRoleGrantedIterator{contract: _IOTXStaking.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IOTXStaking *IOTXStakingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IOTXStakingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IOTXStaking.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakingRoleGranted)
				if err := _IOTXStaking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_IOTXStaking *IOTXStakingFilterer) ParseRoleGranted(log types.Log) (*IOTXStakingRoleGranted, error) {
	event := new(IOTXStakingRoleGranted)
	if err := _IOTXStaking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IOTXStaking contract.
type IOTXStakingRoleRevokedIterator struct {
	Event *IOTXStakingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *IOTXStakingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakingRoleRevoked)
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
		it.Event = new(IOTXStakingRoleRevoked)
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
func (it *IOTXStakingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakingRoleRevoked represents a RoleRevoked event raised by the IOTXStaking contract.
type IOTXStakingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IOTXStaking *IOTXStakingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IOTXStakingRoleRevokedIterator, error) {

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

	logs, sub, err := _IOTXStaking.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IOTXStakingRoleRevokedIterator{contract: _IOTXStaking.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IOTXStaking *IOTXStakingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IOTXStakingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IOTXStaking.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakingRoleRevoked)
				if err := _IOTXStaking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_IOTXStaking *IOTXStakingFilterer) ParseRoleRevoked(log types.Log) (*IOTXStakingRoleRevoked, error) {
	event := new(IOTXStakingRoleRevoked)
	if err := _IOTXStaking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the IOTXStaking contract.
type IOTXStakingStakedIterator struct {
	Event *IOTXStakingStaked // Event containing the contract specifics and raw log

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
func (it *IOTXStakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakingStaked)
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
		it.Event = new(IOTXStakingStaked)
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
func (it *IOTXStakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakingStaked represents a Staked event raised by the IOTXStaking contract.
type IOTXStakingStaked struct {
	FirstTokenId *big.Int
	Amount       *big.Int
	Delegate     common.Address
	Count        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0xc49cabe9d2bbfea8d9f51b0961c30a1081ff6fd3c4d6a9182cd65de8cea2df00.
//
// Solidity: event Staked(uint256 firstTokenId, uint256 amount, address delegate, uint256 count)
func (_IOTXStaking *IOTXStakingFilterer) FilterStaked(opts *bind.FilterOpts) (*IOTXStakingStakedIterator, error) {

	logs, sub, err := _IOTXStaking.contract.FilterLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return &IOTXStakingStakedIterator{contract: _IOTXStaking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0xc49cabe9d2bbfea8d9f51b0961c30a1081ff6fd3c4d6a9182cd65de8cea2df00.
//
// Solidity: event Staked(uint256 firstTokenId, uint256 amount, address delegate, uint256 count)
func (_IOTXStaking *IOTXStakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *IOTXStakingStaked) (event.Subscription, error) {

	logs, sub, err := _IOTXStaking.contract.WatchLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakingStaked)
				if err := _IOTXStaking.contract.UnpackLog(event, "Staked", log); err != nil {
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
func (_IOTXStaking *IOTXStakingFilterer) ParseStaked(log types.Log) (*IOTXStakingStaked, error) {
	event := new(IOTXStakingStaked)
	if err := _IOTXStaking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOTXStakingUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IOTXStaking contract.
type IOTXStakingUnpausedIterator struct {
	Event *IOTXStakingUnpaused // Event containing the contract specifics and raw log

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
func (it *IOTXStakingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOTXStakingUnpaused)
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
		it.Event = new(IOTXStakingUnpaused)
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
func (it *IOTXStakingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOTXStakingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOTXStakingUnpaused represents a Unpaused event raised by the IOTXStaking contract.
type IOTXStakingUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IOTXStaking *IOTXStakingFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IOTXStakingUnpausedIterator, error) {

	logs, sub, err := _IOTXStaking.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IOTXStakingUnpausedIterator{contract: _IOTXStaking.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IOTXStaking *IOTXStakingFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IOTXStakingUnpaused) (event.Subscription, error) {

	logs, sub, err := _IOTXStaking.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOTXStakingUnpaused)
				if err := _IOTXStaking.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_IOTXStaking *IOTXStakingFilterer) ParseUnpaused(log types.Log) (*IOTXStakingUnpaused, error) {
	event := new(IOTXStakingUnpaused)
	if err := _IOTXStaking.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
