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

// BucketType is an auto generated low-level Go binding around an user-defined struct.
type BucketType struct {
	Amount      *big.Int
	Duration    *big.Int
	ActivatedAt *big.Int
}

// SystemStakingMetaData contains all meta data concerning the SystemStaking contract.
var SystemStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"BucketExpanded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"BucketTypeActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"BucketTypeDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"Merged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Unlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UINT256_MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNSTAKE_FREEZE_BLOCKS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"activateBucketType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"addBucketType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"blocksToUnstake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"blocksToWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"bucketOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockedAt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakedAt_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegate_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"name\":\"bucketTypes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activatedAt\",\"type\":\"uint256\"}],\"internalType\":\"structBucketType[]\",\"name\":\"types_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"changeDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"changeDelegates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"deactivateBucketType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newDuration\",\"type\":\"uint256\"}],\"name\":\"expandBucket\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"isActiveBucketType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_delegates\",\"type\":\"address[]\"}],\"name\":\"lockedVotesTo\",\"outputs\":[{\"internalType\":\"uint256[][]\",\"name\":\"counts_\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_newDuration\",\"type\":\"uint256\"}],\"name\":\"merge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfBucketTypes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"firstTokenId_\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_delegates\",\"type\":\"address[]\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"firstTokenId_\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_delegates\",\"type\":\"address[]\"}],\"name\":\"unlockedVotesTo\",\"outputs\":[{\"internalType\":\"uint256[][]\",\"name\":\"counts_\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"addresspayable\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SystemStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemStakingMetaData.ABI instead.
var SystemStakingABI = SystemStakingMetaData.ABI

// SystemStaking is an auto generated Go binding around an Ethereum contract.
type SystemStaking struct {
	SystemStakingCaller     // Read-only binding to the contract
	SystemStakingTransactor // Write-only binding to the contract
	SystemStakingFilterer   // Log filterer for contract events
}

// SystemStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemStakingSession struct {
	Contract     *SystemStaking    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemStakingCallerSession struct {
	Contract *SystemStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SystemStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemStakingTransactorSession struct {
	Contract     *SystemStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SystemStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemStakingRaw struct {
	Contract *SystemStaking // Generic contract binding to access the raw methods on
}

// SystemStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemStakingCallerRaw struct {
	Contract *SystemStakingCaller // Generic read-only contract binding to access the raw methods on
}

// SystemStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemStakingTransactorRaw struct {
	Contract *SystemStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemStaking creates a new instance of SystemStaking, bound to a specific deployed contract.
func NewSystemStaking(address common.Address, backend bind.ContractBackend) (*SystemStaking, error) {
	contract, err := bindSystemStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemStaking{SystemStakingCaller: SystemStakingCaller{contract: contract}, SystemStakingTransactor: SystemStakingTransactor{contract: contract}, SystemStakingFilterer: SystemStakingFilterer{contract: contract}}, nil
}

// NewSystemStakingCaller creates a new read-only instance of SystemStaking, bound to a specific deployed contract.
func NewSystemStakingCaller(address common.Address, caller bind.ContractCaller) (*SystemStakingCaller, error) {
	contract, err := bindSystemStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemStakingCaller{contract: contract}, nil
}

// NewSystemStakingTransactor creates a new write-only instance of SystemStaking, bound to a specific deployed contract.
func NewSystemStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemStakingTransactor, error) {
	contract, err := bindSystemStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemStakingTransactor{contract: contract}, nil
}

// NewSystemStakingFilterer creates a new log filterer instance of SystemStaking, bound to a specific deployed contract.
func NewSystemStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemStakingFilterer, error) {
	contract, err := bindSystemStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemStakingFilterer{contract: contract}, nil
}

// bindSystemStaking binds a generic wrapper to an already deployed contract.
func bindSystemStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemStakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemStaking *SystemStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemStaking.Contract.SystemStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemStaking *SystemStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemStaking.Contract.SystemStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemStaking *SystemStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemStaking.Contract.SystemStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemStaking *SystemStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemStaking *SystemStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemStaking *SystemStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemStaking.Contract.contract.Transact(opts, method, params...)
}

// UINT256MAX is a free data retrieval call binding the contract method 0xd0949f99.
//
// Solidity: function UINT256_MAX() view returns(uint256)
func (_SystemStaking *SystemStakingCaller) UINT256MAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "UINT256_MAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UINT256MAX is a free data retrieval call binding the contract method 0xd0949f99.
//
// Solidity: function UINT256_MAX() view returns(uint256)
func (_SystemStaking *SystemStakingSession) UINT256MAX() (*big.Int, error) {
	return _SystemStaking.Contract.UINT256MAX(&_SystemStaking.CallOpts)
}

// UINT256MAX is a free data retrieval call binding the contract method 0xd0949f99.
//
// Solidity: function UINT256_MAX() view returns(uint256)
func (_SystemStaking *SystemStakingCallerSession) UINT256MAX() (*big.Int, error) {
	return _SystemStaking.Contract.UINT256MAX(&_SystemStaking.CallOpts)
}

// UNSTAKEFREEZEBLOCKS is a free data retrieval call binding the contract method 0xf0b56b5d.
//
// Solidity: function UNSTAKE_FREEZE_BLOCKS() view returns(uint256)
func (_SystemStaking *SystemStakingCaller) UNSTAKEFREEZEBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "UNSTAKE_FREEZE_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNSTAKEFREEZEBLOCKS is a free data retrieval call binding the contract method 0xf0b56b5d.
//
// Solidity: function UNSTAKE_FREEZE_BLOCKS() view returns(uint256)
func (_SystemStaking *SystemStakingSession) UNSTAKEFREEZEBLOCKS() (*big.Int, error) {
	return _SystemStaking.Contract.UNSTAKEFREEZEBLOCKS(&_SystemStaking.CallOpts)
}

// UNSTAKEFREEZEBLOCKS is a free data retrieval call binding the contract method 0xf0b56b5d.
//
// Solidity: function UNSTAKE_FREEZE_BLOCKS() view returns(uint256)
func (_SystemStaking *SystemStakingCallerSession) UNSTAKEFREEZEBLOCKS() (*big.Int, error) {
	return _SystemStaking.Contract.UNSTAKEFREEZEBLOCKS(&_SystemStaking.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SystemStaking *SystemStakingCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SystemStaking *SystemStakingSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SystemStaking.Contract.BalanceOf(&_SystemStaking.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SystemStaking *SystemStakingCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SystemStaking.Contract.BalanceOf(&_SystemStaking.CallOpts, owner)
}

// BlocksToUnstake is a free data retrieval call binding the contract method 0x93b6ef59.
//
// Solidity: function blocksToUnstake(uint256 _tokenId) view returns(uint256)
func (_SystemStaking *SystemStakingCaller) BlocksToUnstake(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "blocksToUnstake", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlocksToUnstake is a free data retrieval call binding the contract method 0x93b6ef59.
//
// Solidity: function blocksToUnstake(uint256 _tokenId) view returns(uint256)
func (_SystemStaking *SystemStakingSession) BlocksToUnstake(_tokenId *big.Int) (*big.Int, error) {
	return _SystemStaking.Contract.BlocksToUnstake(&_SystemStaking.CallOpts, _tokenId)
}

// BlocksToUnstake is a free data retrieval call binding the contract method 0x93b6ef59.
//
// Solidity: function blocksToUnstake(uint256 _tokenId) view returns(uint256)
func (_SystemStaking *SystemStakingCallerSession) BlocksToUnstake(_tokenId *big.Int) (*big.Int, error) {
	return _SystemStaking.Contract.BlocksToUnstake(&_SystemStaking.CallOpts, _tokenId)
}

// BlocksToWithdraw is a free data retrieval call binding the contract method 0x03459b16.
//
// Solidity: function blocksToWithdraw(uint256 _tokenId) view returns(uint256)
func (_SystemStaking *SystemStakingCaller) BlocksToWithdraw(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "blocksToWithdraw", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlocksToWithdraw is a free data retrieval call binding the contract method 0x03459b16.
//
// Solidity: function blocksToWithdraw(uint256 _tokenId) view returns(uint256)
func (_SystemStaking *SystemStakingSession) BlocksToWithdraw(_tokenId *big.Int) (*big.Int, error) {
	return _SystemStaking.Contract.BlocksToWithdraw(&_SystemStaking.CallOpts, _tokenId)
}

// BlocksToWithdraw is a free data retrieval call binding the contract method 0x03459b16.
//
// Solidity: function blocksToWithdraw(uint256 _tokenId) view returns(uint256)
func (_SystemStaking *SystemStakingCallerSession) BlocksToWithdraw(_tokenId *big.Int) (*big.Int, error) {
	return _SystemStaking.Contract.BlocksToWithdraw(&_SystemStaking.CallOpts, _tokenId)
}

// BucketOf is a free data retrieval call binding the contract method 0x431cd92a.
//
// Solidity: function bucketOf(uint256 _tokenId) view returns(uint256 amount_, uint256 duration_, uint256 unlockedAt_, uint256 unstakedAt_, address delegate_)
func (_SystemStaking *SystemStakingCaller) BucketOf(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Amount     *big.Int
	Duration   *big.Int
	UnlockedAt *big.Int
	UnstakedAt *big.Int
	Delegate   common.Address
}, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "bucketOf", _tokenId)

	outstruct := new(struct {
		Amount     *big.Int
		Duration   *big.Int
		UnlockedAt *big.Int
		UnstakedAt *big.Int
		Delegate   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UnlockedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnstakedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Delegate = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// BucketOf is a free data retrieval call binding the contract method 0x431cd92a.
//
// Solidity: function bucketOf(uint256 _tokenId) view returns(uint256 amount_, uint256 duration_, uint256 unlockedAt_, uint256 unstakedAt_, address delegate_)
func (_SystemStaking *SystemStakingSession) BucketOf(_tokenId *big.Int) (struct {
	Amount     *big.Int
	Duration   *big.Int
	UnlockedAt *big.Int
	UnstakedAt *big.Int
	Delegate   common.Address
}, error) {
	return _SystemStaking.Contract.BucketOf(&_SystemStaking.CallOpts, _tokenId)
}

// BucketOf is a free data retrieval call binding the contract method 0x431cd92a.
//
// Solidity: function bucketOf(uint256 _tokenId) view returns(uint256 amount_, uint256 duration_, uint256 unlockedAt_, uint256 unstakedAt_, address delegate_)
func (_SystemStaking *SystemStakingCallerSession) BucketOf(_tokenId *big.Int) (struct {
	Amount     *big.Int
	Duration   *big.Int
	UnlockedAt *big.Int
	UnstakedAt *big.Int
	Delegate   common.Address
}, error) {
	return _SystemStaking.Contract.BucketOf(&_SystemStaking.CallOpts, _tokenId)
}

// BucketTypes is a free data retrieval call binding the contract method 0x78bfca10.
//
// Solidity: function bucketTypes(uint256 _offset, uint256 _size) view returns((uint256,uint256,uint256)[] types_)
func (_SystemStaking *SystemStakingCaller) BucketTypes(opts *bind.CallOpts, _offset *big.Int, _size *big.Int) ([]BucketType, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "bucketTypes", _offset, _size)

	if err != nil {
		return *new([]BucketType), err
	}

	out0 := *abi.ConvertType(out[0], new([]BucketType)).(*[]BucketType)

	return out0, err

}

// BucketTypes is a free data retrieval call binding the contract method 0x78bfca10.
//
// Solidity: function bucketTypes(uint256 _offset, uint256 _size) view returns((uint256,uint256,uint256)[] types_)
func (_SystemStaking *SystemStakingSession) BucketTypes(_offset *big.Int, _size *big.Int) ([]BucketType, error) {
	return _SystemStaking.Contract.BucketTypes(&_SystemStaking.CallOpts, _offset, _size)
}

// BucketTypes is a free data retrieval call binding the contract method 0x78bfca10.
//
// Solidity: function bucketTypes(uint256 _offset, uint256 _size) view returns((uint256,uint256,uint256)[] types_)
func (_SystemStaking *SystemStakingCallerSession) BucketTypes(_offset *big.Int, _size *big.Int) ([]BucketType, error) {
	return _SystemStaking.Contract.BucketTypes(&_SystemStaking.CallOpts, _offset, _size)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SystemStaking *SystemStakingCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SystemStaking *SystemStakingSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SystemStaking.Contract.GetApproved(&_SystemStaking.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SystemStaking *SystemStakingCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SystemStaking.Contract.GetApproved(&_SystemStaking.CallOpts, tokenId)
}

// IsActiveBucketType is a free data retrieval call binding the contract method 0x43e06c59.
//
// Solidity: function isActiveBucketType(uint256 _amount, uint256 _duration) view returns(bool)
func (_SystemStaking *SystemStakingCaller) IsActiveBucketType(opts *bind.CallOpts, _amount *big.Int, _duration *big.Int) (bool, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "isActiveBucketType", _amount, _duration)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveBucketType is a free data retrieval call binding the contract method 0x43e06c59.
//
// Solidity: function isActiveBucketType(uint256 _amount, uint256 _duration) view returns(bool)
func (_SystemStaking *SystemStakingSession) IsActiveBucketType(_amount *big.Int, _duration *big.Int) (bool, error) {
	return _SystemStaking.Contract.IsActiveBucketType(&_SystemStaking.CallOpts, _amount, _duration)
}

// IsActiveBucketType is a free data retrieval call binding the contract method 0x43e06c59.
//
// Solidity: function isActiveBucketType(uint256 _amount, uint256 _duration) view returns(bool)
func (_SystemStaking *SystemStakingCallerSession) IsActiveBucketType(_amount *big.Int, _duration *big.Int) (bool, error) {
	return _SystemStaking.Contract.IsActiveBucketType(&_SystemStaking.CallOpts, _amount, _duration)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SystemStaking *SystemStakingCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SystemStaking *SystemStakingSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SystemStaking.Contract.IsApprovedForAll(&_SystemStaking.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SystemStaking *SystemStakingCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SystemStaking.Contract.IsApprovedForAll(&_SystemStaking.CallOpts, owner, operator)
}

// LockedVotesTo is a free data retrieval call binding the contract method 0x3fd140df.
//
// Solidity: function lockedVotesTo(address[] _delegates) view returns(uint256[][] counts_)
func (_SystemStaking *SystemStakingCaller) LockedVotesTo(opts *bind.CallOpts, _delegates []common.Address) ([][]*big.Int, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "lockedVotesTo", _delegates)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// LockedVotesTo is a free data retrieval call binding the contract method 0x3fd140df.
//
// Solidity: function lockedVotesTo(address[] _delegates) view returns(uint256[][] counts_)
func (_SystemStaking *SystemStakingSession) LockedVotesTo(_delegates []common.Address) ([][]*big.Int, error) {
	return _SystemStaking.Contract.LockedVotesTo(&_SystemStaking.CallOpts, _delegates)
}

// LockedVotesTo is a free data retrieval call binding the contract method 0x3fd140df.
//
// Solidity: function lockedVotesTo(address[] _delegates) view returns(uint256[][] counts_)
func (_SystemStaking *SystemStakingCallerSession) LockedVotesTo(_delegates []common.Address) ([][]*big.Int, error) {
	return _SystemStaking.Contract.LockedVotesTo(&_SystemStaking.CallOpts, _delegates)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SystemStaking *SystemStakingCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SystemStaking *SystemStakingSession) Name() (string, error) {
	return _SystemStaking.Contract.Name(&_SystemStaking.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SystemStaking *SystemStakingCallerSession) Name() (string, error) {
	return _SystemStaking.Contract.Name(&_SystemStaking.CallOpts)
}

// NumOfBucketTypes is a free data retrieval call binding the contract method 0x9f7d5b00.
//
// Solidity: function numOfBucketTypes() view returns(uint256)
func (_SystemStaking *SystemStakingCaller) NumOfBucketTypes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "numOfBucketTypes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfBucketTypes is a free data retrieval call binding the contract method 0x9f7d5b00.
//
// Solidity: function numOfBucketTypes() view returns(uint256)
func (_SystemStaking *SystemStakingSession) NumOfBucketTypes() (*big.Int, error) {
	return _SystemStaking.Contract.NumOfBucketTypes(&_SystemStaking.CallOpts)
}

// NumOfBucketTypes is a free data retrieval call binding the contract method 0x9f7d5b00.
//
// Solidity: function numOfBucketTypes() view returns(uint256)
func (_SystemStaking *SystemStakingCallerSession) NumOfBucketTypes() (*big.Int, error) {
	return _SystemStaking.Contract.NumOfBucketTypes(&_SystemStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemStaking *SystemStakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemStaking *SystemStakingSession) Owner() (common.Address, error) {
	return _SystemStaking.Contract.Owner(&_SystemStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemStaking *SystemStakingCallerSession) Owner() (common.Address, error) {
	return _SystemStaking.Contract.Owner(&_SystemStaking.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SystemStaking *SystemStakingCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SystemStaking *SystemStakingSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SystemStaking.Contract.OwnerOf(&_SystemStaking.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SystemStaking *SystemStakingCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SystemStaking.Contract.OwnerOf(&_SystemStaking.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SystemStaking *SystemStakingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SystemStaking *SystemStakingSession) Paused() (bool, error) {
	return _SystemStaking.Contract.Paused(&_SystemStaking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SystemStaking *SystemStakingCallerSession) Paused() (bool, error) {
	return _SystemStaking.Contract.Paused(&_SystemStaking.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SystemStaking *SystemStakingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SystemStaking *SystemStakingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SystemStaking.Contract.SupportsInterface(&_SystemStaking.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SystemStaking *SystemStakingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SystemStaking.Contract.SupportsInterface(&_SystemStaking.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SystemStaking *SystemStakingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SystemStaking *SystemStakingSession) Symbol() (string, error) {
	return _SystemStaking.Contract.Symbol(&_SystemStaking.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SystemStaking *SystemStakingCallerSession) Symbol() (string, error) {
	return _SystemStaking.Contract.Symbol(&_SystemStaking.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SystemStaking *SystemStakingCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SystemStaking *SystemStakingSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SystemStaking.Contract.TokenURI(&_SystemStaking.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SystemStaking *SystemStakingCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SystemStaking.Contract.TokenURI(&_SystemStaking.CallOpts, tokenId)
}

// UnlockedVotesTo is a free data retrieval call binding the contract method 0x960014bd.
//
// Solidity: function unlockedVotesTo(address[] _delegates) view returns(uint256[][] counts_)
func (_SystemStaking *SystemStakingCaller) UnlockedVotesTo(opts *bind.CallOpts, _delegates []common.Address) ([][]*big.Int, error) {
	var out []interface{}
	err := _SystemStaking.contract.Call(opts, &out, "unlockedVotesTo", _delegates)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// UnlockedVotesTo is a free data retrieval call binding the contract method 0x960014bd.
//
// Solidity: function unlockedVotesTo(address[] _delegates) view returns(uint256[][] counts_)
func (_SystemStaking *SystemStakingSession) UnlockedVotesTo(_delegates []common.Address) ([][]*big.Int, error) {
	return _SystemStaking.Contract.UnlockedVotesTo(&_SystemStaking.CallOpts, _delegates)
}

// UnlockedVotesTo is a free data retrieval call binding the contract method 0x960014bd.
//
// Solidity: function unlockedVotesTo(address[] _delegates) view returns(uint256[][] counts_)
func (_SystemStaking *SystemStakingCallerSession) UnlockedVotesTo(_delegates []common.Address) ([][]*big.Int, error) {
	return _SystemStaking.Contract.UnlockedVotesTo(&_SystemStaking.CallOpts, _delegates)
}

// ActivateBucketType is a paid mutator transaction binding the contract method 0xe0028ecf.
//
// Solidity: function activateBucketType(uint256 _amount, uint256 _duration) returns()
func (_SystemStaking *SystemStakingTransactor) ActivateBucketType(opts *bind.TransactOpts, _amount *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "activateBucketType", _amount, _duration)
}

// ActivateBucketType is a paid mutator transaction binding the contract method 0xe0028ecf.
//
// Solidity: function activateBucketType(uint256 _amount, uint256 _duration) returns()
func (_SystemStaking *SystemStakingSession) ActivateBucketType(_amount *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.ActivateBucketType(&_SystemStaking.TransactOpts, _amount, _duration)
}

// ActivateBucketType is a paid mutator transaction binding the contract method 0xe0028ecf.
//
// Solidity: function activateBucketType(uint256 _amount, uint256 _duration) returns()
func (_SystemStaking *SystemStakingTransactorSession) ActivateBucketType(_amount *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.ActivateBucketType(&_SystemStaking.TransactOpts, _amount, _duration)
}

// AddBucketType is a paid mutator transaction binding the contract method 0xc8e77923.
//
// Solidity: function addBucketType(uint256 _amount, uint256 _duration) returns()
func (_SystemStaking *SystemStakingTransactor) AddBucketType(opts *bind.TransactOpts, _amount *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "addBucketType", _amount, _duration)
}

// AddBucketType is a paid mutator transaction binding the contract method 0xc8e77923.
//
// Solidity: function addBucketType(uint256 _amount, uint256 _duration) returns()
func (_SystemStaking *SystemStakingSession) AddBucketType(_amount *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.AddBucketType(&_SystemStaking.TransactOpts, _amount, _duration)
}

// AddBucketType is a paid mutator transaction binding the contract method 0xc8e77923.
//
// Solidity: function addBucketType(uint256 _amount, uint256 _duration) returns()
func (_SystemStaking *SystemStakingTransactorSession) AddBucketType(_amount *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.AddBucketType(&_SystemStaking.TransactOpts, _amount, _duration)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SystemStaking *SystemStakingTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SystemStaking *SystemStakingSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.Approve(&_SystemStaking.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SystemStaking *SystemStakingTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.Approve(&_SystemStaking.TransactOpts, to, tokenId)
}

// ChangeDelegate is a paid mutator transaction binding the contract method 0x0f5b2ca5.
//
// Solidity: function changeDelegate(uint256 _tokenId, address _delegate) returns()
func (_SystemStaking *SystemStakingTransactor) ChangeDelegate(opts *bind.TransactOpts, _tokenId *big.Int, _delegate common.Address) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "changeDelegate", _tokenId, _delegate)
}

// ChangeDelegate is a paid mutator transaction binding the contract method 0x0f5b2ca5.
//
// Solidity: function changeDelegate(uint256 _tokenId, address _delegate) returns()
func (_SystemStaking *SystemStakingSession) ChangeDelegate(_tokenId *big.Int, _delegate common.Address) (*types.Transaction, error) {
	return _SystemStaking.Contract.ChangeDelegate(&_SystemStaking.TransactOpts, _tokenId, _delegate)
}

// ChangeDelegate is a paid mutator transaction binding the contract method 0x0f5b2ca5.
//
// Solidity: function changeDelegate(uint256 _tokenId, address _delegate) returns()
func (_SystemStaking *SystemStakingTransactorSession) ChangeDelegate(_tokenId *big.Int, _delegate common.Address) (*types.Transaction, error) {
	return _SystemStaking.Contract.ChangeDelegate(&_SystemStaking.TransactOpts, _tokenId, _delegate)
}

// ChangeDelegates is a paid mutator transaction binding the contract method 0x98ca3b76.
//
// Solidity: function changeDelegates(uint256[] _tokenIds, address _delegate) returns()
func (_SystemStaking *SystemStakingTransactor) ChangeDelegates(opts *bind.TransactOpts, _tokenIds []*big.Int, _delegate common.Address) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "changeDelegates", _tokenIds, _delegate)
}

// ChangeDelegates is a paid mutator transaction binding the contract method 0x98ca3b76.
//
// Solidity: function changeDelegates(uint256[] _tokenIds, address _delegate) returns()
func (_SystemStaking *SystemStakingSession) ChangeDelegates(_tokenIds []*big.Int, _delegate common.Address) (*types.Transaction, error) {
	return _SystemStaking.Contract.ChangeDelegates(&_SystemStaking.TransactOpts, _tokenIds, _delegate)
}

// ChangeDelegates is a paid mutator transaction binding the contract method 0x98ca3b76.
//
// Solidity: function changeDelegates(uint256[] _tokenIds, address _delegate) returns()
func (_SystemStaking *SystemStakingTransactorSession) ChangeDelegates(_tokenIds []*big.Int, _delegate common.Address) (*types.Transaction, error) {
	return _SystemStaking.Contract.ChangeDelegates(&_SystemStaking.TransactOpts, _tokenIds, _delegate)
}

// DeactivateBucketType is a paid mutator transaction binding the contract method 0xeb0ffb2e.
//
// Solidity: function deactivateBucketType(uint256 _amount, uint256 _duration) returns()
func (_SystemStaking *SystemStakingTransactor) DeactivateBucketType(opts *bind.TransactOpts, _amount *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "deactivateBucketType", _amount, _duration)
}

// DeactivateBucketType is a paid mutator transaction binding the contract method 0xeb0ffb2e.
//
// Solidity: function deactivateBucketType(uint256 _amount, uint256 _duration) returns()
func (_SystemStaking *SystemStakingSession) DeactivateBucketType(_amount *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.DeactivateBucketType(&_SystemStaking.TransactOpts, _amount, _duration)
}

// DeactivateBucketType is a paid mutator transaction binding the contract method 0xeb0ffb2e.
//
// Solidity: function deactivateBucketType(uint256 _amount, uint256 _duration) returns()
func (_SystemStaking *SystemStakingTransactorSession) DeactivateBucketType(_amount *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.DeactivateBucketType(&_SystemStaking.TransactOpts, _amount, _duration)
}

// ExpandBucket is a paid mutator transaction binding the contract method 0x025008ed.
//
// Solidity: function expandBucket(uint256 _tokenId, uint256 _newAmount, uint256 _newDuration) payable returns()
func (_SystemStaking *SystemStakingTransactor) ExpandBucket(opts *bind.TransactOpts, _tokenId *big.Int, _newAmount *big.Int, _newDuration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "expandBucket", _tokenId, _newAmount, _newDuration)
}

// ExpandBucket is a paid mutator transaction binding the contract method 0x025008ed.
//
// Solidity: function expandBucket(uint256 _tokenId, uint256 _newAmount, uint256 _newDuration) payable returns()
func (_SystemStaking *SystemStakingSession) ExpandBucket(_tokenId *big.Int, _newAmount *big.Int, _newDuration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.ExpandBucket(&_SystemStaking.TransactOpts, _tokenId, _newAmount, _newDuration)
}

// ExpandBucket is a paid mutator transaction binding the contract method 0x025008ed.
//
// Solidity: function expandBucket(uint256 _tokenId, uint256 _newAmount, uint256 _newDuration) payable returns()
func (_SystemStaking *SystemStakingTransactorSession) ExpandBucket(_tokenId *big.Int, _newAmount *big.Int, _newDuration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.ExpandBucket(&_SystemStaking.TransactOpts, _tokenId, _newAmount, _newDuration)
}

// Lock is a paid mutator transaction binding the contract method 0x1338736f.
//
// Solidity: function lock(uint256 _tokenId, uint256 _duration) returns()
func (_SystemStaking *SystemStakingTransactor) Lock(opts *bind.TransactOpts, _tokenId *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "lock", _tokenId, _duration)
}

// Lock is a paid mutator transaction binding the contract method 0x1338736f.
//
// Solidity: function lock(uint256 _tokenId, uint256 _duration) returns()
func (_SystemStaking *SystemStakingSession) Lock(_tokenId *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.Lock(&_SystemStaking.TransactOpts, _tokenId, _duration)
}

// Lock is a paid mutator transaction binding the contract method 0x1338736f.
//
// Solidity: function lock(uint256 _tokenId, uint256 _duration) returns()
func (_SystemStaking *SystemStakingTransactorSession) Lock(_tokenId *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.Lock(&_SystemStaking.TransactOpts, _tokenId, _duration)
}

// Lock0 is a paid mutator transaction binding the contract method 0x5ceb8b5b.
//
// Solidity: function lock(uint256[] _tokenIds, uint256 _duration) returns()
func (_SystemStaking *SystemStakingTransactor) Lock0(opts *bind.TransactOpts, _tokenIds []*big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "lock0", _tokenIds, _duration)
}

// Lock0 is a paid mutator transaction binding the contract method 0x5ceb8b5b.
//
// Solidity: function lock(uint256[] _tokenIds, uint256 _duration) returns()
func (_SystemStaking *SystemStakingSession) Lock0(_tokenIds []*big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.Lock0(&_SystemStaking.TransactOpts, _tokenIds, _duration)
}

// Lock0 is a paid mutator transaction binding the contract method 0x5ceb8b5b.
//
// Solidity: function lock(uint256[] _tokenIds, uint256 _duration) returns()
func (_SystemStaking *SystemStakingTransactorSession) Lock0(_tokenIds []*big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.Lock0(&_SystemStaking.TransactOpts, _tokenIds, _duration)
}

// Merge is a paid mutator transaction binding the contract method 0xbbe33ea5.
//
// Solidity: function merge(uint256[] tokenIds, uint256 _newDuration) payable returns()
func (_SystemStaking *SystemStakingTransactor) Merge(opts *bind.TransactOpts, tokenIds []*big.Int, _newDuration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "merge", tokenIds, _newDuration)
}

// Merge is a paid mutator transaction binding the contract method 0xbbe33ea5.
//
// Solidity: function merge(uint256[] tokenIds, uint256 _newDuration) payable returns()
func (_SystemStaking *SystemStakingSession) Merge(tokenIds []*big.Int, _newDuration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.Merge(&_SystemStaking.TransactOpts, tokenIds, _newDuration)
}

// Merge is a paid mutator transaction binding the contract method 0xbbe33ea5.
//
// Solidity: function merge(uint256[] tokenIds, uint256 _newDuration) payable returns()
func (_SystemStaking *SystemStakingTransactorSession) Merge(tokenIds []*big.Int, _newDuration *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.Merge(&_SystemStaking.TransactOpts, tokenIds, _newDuration)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SystemStaking *SystemStakingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SystemStaking *SystemStakingSession) Pause() (*types.Transaction, error) {
	return _SystemStaking.Contract.Pause(&_SystemStaking.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SystemStaking *SystemStakingTransactorSession) Pause() (*types.Transaction, error) {
	return _SystemStaking.Contract.Pause(&_SystemStaking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemStaking *SystemStakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemStaking *SystemStakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _SystemStaking.Contract.RenounceOwnership(&_SystemStaking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemStaking *SystemStakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SystemStaking.Contract.RenounceOwnership(&_SystemStaking.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SystemStaking *SystemStakingTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SystemStaking *SystemStakingSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.SafeTransferFrom(&_SystemStaking.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SystemStaking *SystemStakingTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.SafeTransferFrom(&_SystemStaking.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SystemStaking *SystemStakingTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SystemStaking *SystemStakingSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SystemStaking.Contract.SafeTransferFrom0(&_SystemStaking.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SystemStaking *SystemStakingTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SystemStaking.Contract.SafeTransferFrom0(&_SystemStaking.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SystemStaking *SystemStakingTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SystemStaking *SystemStakingSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SystemStaking.Contract.SetApprovalForAll(&_SystemStaking.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SystemStaking *SystemStakingTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SystemStaking.Contract.SetApprovalForAll(&_SystemStaking.TransactOpts, operator, approved)
}

// Stake is a paid mutator transaction binding the contract method 0x711563d4.
//
// Solidity: function stake(uint256 _amount, uint256 _duration, address _delegate, uint256 _count) payable returns(uint256 firstTokenId_)
func (_SystemStaking *SystemStakingTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int, _duration *big.Int, _delegate common.Address, _count *big.Int) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "stake", _amount, _duration, _delegate, _count)
}

// Stake is a paid mutator transaction binding the contract method 0x711563d4.
//
// Solidity: function stake(uint256 _amount, uint256 _duration, address _delegate, uint256 _count) payable returns(uint256 firstTokenId_)
func (_SystemStaking *SystemStakingSession) Stake(_amount *big.Int, _duration *big.Int, _delegate common.Address, _count *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.Stake(&_SystemStaking.TransactOpts, _amount, _duration, _delegate, _count)
}

// Stake is a paid mutator transaction binding the contract method 0x711563d4.
//
// Solidity: function stake(uint256 _amount, uint256 _duration, address _delegate, uint256 _count) payable returns(uint256 firstTokenId_)
func (_SystemStaking *SystemStakingTransactorSession) Stake(_amount *big.Int, _duration *big.Int, _delegate common.Address, _count *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.Stake(&_SystemStaking.TransactOpts, _amount, _duration, _delegate, _count)
}

// Stake0 is a paid mutator transaction binding the contract method 0x7acb7757.
//
// Solidity: function stake(uint256 _duration, address _delegate) payable returns(uint256)
func (_SystemStaking *SystemStakingTransactor) Stake0(opts *bind.TransactOpts, _duration *big.Int, _delegate common.Address) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "stake0", _duration, _delegate)
}

// Stake0 is a paid mutator transaction binding the contract method 0x7acb7757.
//
// Solidity: function stake(uint256 _duration, address _delegate) payable returns(uint256)
func (_SystemStaking *SystemStakingSession) Stake0(_duration *big.Int, _delegate common.Address) (*types.Transaction, error) {
	return _SystemStaking.Contract.Stake0(&_SystemStaking.TransactOpts, _duration, _delegate)
}

// Stake0 is a paid mutator transaction binding the contract method 0x7acb7757.
//
// Solidity: function stake(uint256 _duration, address _delegate) payable returns(uint256)
func (_SystemStaking *SystemStakingTransactorSession) Stake0(_duration *big.Int, _delegate common.Address) (*types.Transaction, error) {
	return _SystemStaking.Contract.Stake0(&_SystemStaking.TransactOpts, _duration, _delegate)
}

// Stake1 is a paid mutator transaction binding the contract method 0xeec7ee73.
//
// Solidity: function stake(uint256 _amount, uint256 _duration, address[] _delegates) payable returns(uint256 firstTokenId_)
func (_SystemStaking *SystemStakingTransactor) Stake1(opts *bind.TransactOpts, _amount *big.Int, _duration *big.Int, _delegates []common.Address) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "stake1", _amount, _duration, _delegates)
}

// Stake1 is a paid mutator transaction binding the contract method 0xeec7ee73.
//
// Solidity: function stake(uint256 _amount, uint256 _duration, address[] _delegates) payable returns(uint256 firstTokenId_)
func (_SystemStaking *SystemStakingSession) Stake1(_amount *big.Int, _duration *big.Int, _delegates []common.Address) (*types.Transaction, error) {
	return _SystemStaking.Contract.Stake1(&_SystemStaking.TransactOpts, _amount, _duration, _delegates)
}

// Stake1 is a paid mutator transaction binding the contract method 0xeec7ee73.
//
// Solidity: function stake(uint256 _amount, uint256 _duration, address[] _delegates) payable returns(uint256 firstTokenId_)
func (_SystemStaking *SystemStakingTransactorSession) Stake1(_amount *big.Int, _duration *big.Int, _delegates []common.Address) (*types.Transaction, error) {
	return _SystemStaking.Contract.Stake1(&_SystemStaking.TransactOpts, _amount, _duration, _delegates)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SystemStaking *SystemStakingTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SystemStaking *SystemStakingSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.TransferFrom(&_SystemStaking.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SystemStaking *SystemStakingTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.TransferFrom(&_SystemStaking.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemStaking *SystemStakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemStaking *SystemStakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SystemStaking.Contract.TransferOwnership(&_SystemStaking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemStaking *SystemStakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SystemStaking.Contract.TransferOwnership(&_SystemStaking.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0x5d36598f.
//
// Solidity: function unlock(uint256[] _tokenIds) returns()
func (_SystemStaking *SystemStakingTransactor) Unlock(opts *bind.TransactOpts, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "unlock", _tokenIds)
}

// Unlock is a paid mutator transaction binding the contract method 0x5d36598f.
//
// Solidity: function unlock(uint256[] _tokenIds) returns()
func (_SystemStaking *SystemStakingSession) Unlock(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.Unlock(&_SystemStaking.TransactOpts, _tokenIds)
}

// Unlock is a paid mutator transaction binding the contract method 0x5d36598f.
//
// Solidity: function unlock(uint256[] _tokenIds) returns()
func (_SystemStaking *SystemStakingTransactorSession) Unlock(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.Unlock(&_SystemStaking.TransactOpts, _tokenIds)
}

// Unlock0 is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 _tokenId) returns()
func (_SystemStaking *SystemStakingTransactor) Unlock0(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "unlock0", _tokenId)
}

// Unlock0 is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 _tokenId) returns()
func (_SystemStaking *SystemStakingSession) Unlock0(_tokenId *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.Unlock0(&_SystemStaking.TransactOpts, _tokenId)
}

// Unlock0 is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 _tokenId) returns()
func (_SystemStaking *SystemStakingTransactorSession) Unlock0(_tokenId *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.Unlock0(&_SystemStaking.TransactOpts, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SystemStaking *SystemStakingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SystemStaking *SystemStakingSession) Unpause() (*types.Transaction, error) {
	return _SystemStaking.Contract.Unpause(&_SystemStaking.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SystemStaking *SystemStakingTransactorSession) Unpause() (*types.Transaction, error) {
	return _SystemStaking.Contract.Unpause(&_SystemStaking.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _tokenId) returns()
func (_SystemStaking *SystemStakingTransactor) Unstake(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "unstake", _tokenId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _tokenId) returns()
func (_SystemStaking *SystemStakingSession) Unstake(_tokenId *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.Unstake(&_SystemStaking.TransactOpts, _tokenId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _tokenId) returns()
func (_SystemStaking *SystemStakingTransactorSession) Unstake(_tokenId *big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.Unstake(&_SystemStaking.TransactOpts, _tokenId)
}

// Unstake0 is a paid mutator transaction binding the contract method 0xe449f341.
//
// Solidity: function unstake(uint256[] _tokenIds) returns()
func (_SystemStaking *SystemStakingTransactor) Unstake0(opts *bind.TransactOpts, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "unstake0", _tokenIds)
}

// Unstake0 is a paid mutator transaction binding the contract method 0xe449f341.
//
// Solidity: function unstake(uint256[] _tokenIds) returns()
func (_SystemStaking *SystemStakingSession) Unstake0(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.Unstake0(&_SystemStaking.TransactOpts, _tokenIds)
}

// Unstake0 is a paid mutator transaction binding the contract method 0xe449f341.
//
// Solidity: function unstake(uint256[] _tokenIds) returns()
func (_SystemStaking *SystemStakingTransactorSession) Unstake0(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _SystemStaking.Contract.Unstake0(&_SystemStaking.TransactOpts, _tokenIds)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 _tokenId, address _recipient) returns()
func (_SystemStaking *SystemStakingTransactor) Withdraw(opts *bind.TransactOpts, _tokenId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "withdraw", _tokenId, _recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 _tokenId, address _recipient) returns()
func (_SystemStaking *SystemStakingSession) Withdraw(_tokenId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _SystemStaking.Contract.Withdraw(&_SystemStaking.TransactOpts, _tokenId, _recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 _tokenId, address _recipient) returns()
func (_SystemStaking *SystemStakingTransactorSession) Withdraw(_tokenId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _SystemStaking.Contract.Withdraw(&_SystemStaking.TransactOpts, _tokenId, _recipient)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xb8f4bd7b.
//
// Solidity: function withdraw(uint256[] _tokenIds, address _recipient) returns()
func (_SystemStaking *SystemStakingTransactor) Withdraw0(opts *bind.TransactOpts, _tokenIds []*big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _SystemStaking.contract.Transact(opts, "withdraw0", _tokenIds, _recipient)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xb8f4bd7b.
//
// Solidity: function withdraw(uint256[] _tokenIds, address _recipient) returns()
func (_SystemStaking *SystemStakingSession) Withdraw0(_tokenIds []*big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _SystemStaking.Contract.Withdraw0(&_SystemStaking.TransactOpts, _tokenIds, _recipient)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xb8f4bd7b.
//
// Solidity: function withdraw(uint256[] _tokenIds, address _recipient) returns()
func (_SystemStaking *SystemStakingTransactorSession) Withdraw0(_tokenIds []*big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _SystemStaking.Contract.Withdraw0(&_SystemStaking.TransactOpts, _tokenIds, _recipient)
}

// SystemStakingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SystemStaking contract.
type SystemStakingApprovalIterator struct {
	Event *SystemStakingApproval // Event containing the contract specifics and raw log

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
func (it *SystemStakingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemStakingApproval)
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
		it.Event = new(SystemStakingApproval)
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
func (it *SystemStakingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemStakingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemStakingApproval represents a Approval event raised by the SystemStaking contract.
type SystemStakingApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SystemStaking *SystemStakingFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SystemStakingApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SystemStaking.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SystemStakingApprovalIterator{contract: _SystemStaking.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SystemStaking *SystemStakingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SystemStakingApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SystemStaking.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemStakingApproval)
				if err := _SystemStaking.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SystemStaking *SystemStakingFilterer) ParseApproval(log types.Log) (*SystemStakingApproval, error) {
	event := new(SystemStakingApproval)
	if err := _SystemStaking.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemStakingApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the SystemStaking contract.
type SystemStakingApprovalForAllIterator struct {
	Event *SystemStakingApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SystemStakingApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemStakingApprovalForAll)
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
		it.Event = new(SystemStakingApprovalForAll)
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
func (it *SystemStakingApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemStakingApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemStakingApprovalForAll represents a ApprovalForAll event raised by the SystemStaking contract.
type SystemStakingApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SystemStaking *SystemStakingFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SystemStakingApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SystemStaking.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SystemStakingApprovalForAllIterator{contract: _SystemStaking.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SystemStaking *SystemStakingFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SystemStakingApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SystemStaking.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemStakingApprovalForAll)
				if err := _SystemStaking.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SystemStaking *SystemStakingFilterer) ParseApprovalForAll(log types.Log) (*SystemStakingApprovalForAll, error) {
	event := new(SystemStakingApprovalForAll)
	if err := _SystemStaking.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemStakingBucketExpandedIterator is returned from FilterBucketExpanded and is used to iterate over the raw logs and unpacked data for BucketExpanded events raised by the SystemStaking contract.
type SystemStakingBucketExpandedIterator struct {
	Event *SystemStakingBucketExpanded // Event containing the contract specifics and raw log

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
func (it *SystemStakingBucketExpandedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemStakingBucketExpanded)
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
		it.Event = new(SystemStakingBucketExpanded)
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
func (it *SystemStakingBucketExpandedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemStakingBucketExpandedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemStakingBucketExpanded represents a BucketExpanded event raised by the SystemStaking contract.
type SystemStakingBucketExpanded struct {
	TokenId  *big.Int
	Amount   *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBucketExpanded is a free log retrieval operation binding the contract event 0xd29e04160a74f0dbab5e7b82ef0392d86d11ac2939e5883eb3353be4cfedb83e.
//
// Solidity: event BucketExpanded(uint256 indexed tokenId, uint256 amount, uint256 duration)
func (_SystemStaking *SystemStakingFilterer) FilterBucketExpanded(opts *bind.FilterOpts, tokenId []*big.Int) (*SystemStakingBucketExpandedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SystemStaking.contract.FilterLogs(opts, "BucketExpanded", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SystemStakingBucketExpandedIterator{contract: _SystemStaking.contract, event: "BucketExpanded", logs: logs, sub: sub}, nil
}

// WatchBucketExpanded is a free log subscription operation binding the contract event 0xd29e04160a74f0dbab5e7b82ef0392d86d11ac2939e5883eb3353be4cfedb83e.
//
// Solidity: event BucketExpanded(uint256 indexed tokenId, uint256 amount, uint256 duration)
func (_SystemStaking *SystemStakingFilterer) WatchBucketExpanded(opts *bind.WatchOpts, sink chan<- *SystemStakingBucketExpanded, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SystemStaking.contract.WatchLogs(opts, "BucketExpanded", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemStakingBucketExpanded)
				if err := _SystemStaking.contract.UnpackLog(event, "BucketExpanded", log); err != nil {
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

// ParseBucketExpanded is a log parse operation binding the contract event 0xd29e04160a74f0dbab5e7b82ef0392d86d11ac2939e5883eb3353be4cfedb83e.
//
// Solidity: event BucketExpanded(uint256 indexed tokenId, uint256 amount, uint256 duration)
func (_SystemStaking *SystemStakingFilterer) ParseBucketExpanded(log types.Log) (*SystemStakingBucketExpanded, error) {
	event := new(SystemStakingBucketExpanded)
	if err := _SystemStaking.contract.UnpackLog(event, "BucketExpanded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemStakingBucketTypeActivatedIterator is returned from FilterBucketTypeActivated and is used to iterate over the raw logs and unpacked data for BucketTypeActivated events raised by the SystemStaking contract.
type SystemStakingBucketTypeActivatedIterator struct {
	Event *SystemStakingBucketTypeActivated // Event containing the contract specifics and raw log

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
func (it *SystemStakingBucketTypeActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemStakingBucketTypeActivated)
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
		it.Event = new(SystemStakingBucketTypeActivated)
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
func (it *SystemStakingBucketTypeActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemStakingBucketTypeActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemStakingBucketTypeActivated represents a BucketTypeActivated event raised by the SystemStaking contract.
type SystemStakingBucketTypeActivated struct {
	Amount   *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBucketTypeActivated is a free log retrieval operation binding the contract event 0x6b39e3267efcd6611c8d7d2534c4715dcb4824322b90d85540a3a82967b6e7b7.
//
// Solidity: event BucketTypeActivated(uint256 amount, uint256 duration)
func (_SystemStaking *SystemStakingFilterer) FilterBucketTypeActivated(opts *bind.FilterOpts) (*SystemStakingBucketTypeActivatedIterator, error) {

	logs, sub, err := _SystemStaking.contract.FilterLogs(opts, "BucketTypeActivated")
	if err != nil {
		return nil, err
	}
	return &SystemStakingBucketTypeActivatedIterator{contract: _SystemStaking.contract, event: "BucketTypeActivated", logs: logs, sub: sub}, nil
}

// WatchBucketTypeActivated is a free log subscription operation binding the contract event 0x6b39e3267efcd6611c8d7d2534c4715dcb4824322b90d85540a3a82967b6e7b7.
//
// Solidity: event BucketTypeActivated(uint256 amount, uint256 duration)
func (_SystemStaking *SystemStakingFilterer) WatchBucketTypeActivated(opts *bind.WatchOpts, sink chan<- *SystemStakingBucketTypeActivated) (event.Subscription, error) {

	logs, sub, err := _SystemStaking.contract.WatchLogs(opts, "BucketTypeActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemStakingBucketTypeActivated)
				if err := _SystemStaking.contract.UnpackLog(event, "BucketTypeActivated", log); err != nil {
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

// ParseBucketTypeActivated is a log parse operation binding the contract event 0x6b39e3267efcd6611c8d7d2534c4715dcb4824322b90d85540a3a82967b6e7b7.
//
// Solidity: event BucketTypeActivated(uint256 amount, uint256 duration)
func (_SystemStaking *SystemStakingFilterer) ParseBucketTypeActivated(log types.Log) (*SystemStakingBucketTypeActivated, error) {
	event := new(SystemStakingBucketTypeActivated)
	if err := _SystemStaking.contract.UnpackLog(event, "BucketTypeActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemStakingBucketTypeDeactivatedIterator is returned from FilterBucketTypeDeactivated and is used to iterate over the raw logs and unpacked data for BucketTypeDeactivated events raised by the SystemStaking contract.
type SystemStakingBucketTypeDeactivatedIterator struct {
	Event *SystemStakingBucketTypeDeactivated // Event containing the contract specifics and raw log

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
func (it *SystemStakingBucketTypeDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemStakingBucketTypeDeactivated)
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
		it.Event = new(SystemStakingBucketTypeDeactivated)
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
func (it *SystemStakingBucketTypeDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemStakingBucketTypeDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemStakingBucketTypeDeactivated represents a BucketTypeDeactivated event raised by the SystemStaking contract.
type SystemStakingBucketTypeDeactivated struct {
	Amount   *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBucketTypeDeactivated is a free log retrieval operation binding the contract event 0x099df2bf9247b43481cf1b791a4dd5fa1220c40c62940da539082fbcb30241d6.
//
// Solidity: event BucketTypeDeactivated(uint256 amount, uint256 duration)
func (_SystemStaking *SystemStakingFilterer) FilterBucketTypeDeactivated(opts *bind.FilterOpts) (*SystemStakingBucketTypeDeactivatedIterator, error) {

	logs, sub, err := _SystemStaking.contract.FilterLogs(opts, "BucketTypeDeactivated")
	if err != nil {
		return nil, err
	}
	return &SystemStakingBucketTypeDeactivatedIterator{contract: _SystemStaking.contract, event: "BucketTypeDeactivated", logs: logs, sub: sub}, nil
}

// WatchBucketTypeDeactivated is a free log subscription operation binding the contract event 0x099df2bf9247b43481cf1b791a4dd5fa1220c40c62940da539082fbcb30241d6.
//
// Solidity: event BucketTypeDeactivated(uint256 amount, uint256 duration)
func (_SystemStaking *SystemStakingFilterer) WatchBucketTypeDeactivated(opts *bind.WatchOpts, sink chan<- *SystemStakingBucketTypeDeactivated) (event.Subscription, error) {

	logs, sub, err := _SystemStaking.contract.WatchLogs(opts, "BucketTypeDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemStakingBucketTypeDeactivated)
				if err := _SystemStaking.contract.UnpackLog(event, "BucketTypeDeactivated", log); err != nil {
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

// ParseBucketTypeDeactivated is a log parse operation binding the contract event 0x099df2bf9247b43481cf1b791a4dd5fa1220c40c62940da539082fbcb30241d6.
//
// Solidity: event BucketTypeDeactivated(uint256 amount, uint256 duration)
func (_SystemStaking *SystemStakingFilterer) ParseBucketTypeDeactivated(log types.Log) (*SystemStakingBucketTypeDeactivated, error) {
	event := new(SystemStakingBucketTypeDeactivated)
	if err := _SystemStaking.contract.UnpackLog(event, "BucketTypeDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemStakingDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the SystemStaking contract.
type SystemStakingDelegateChangedIterator struct {
	Event *SystemStakingDelegateChanged // Event containing the contract specifics and raw log

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
func (it *SystemStakingDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemStakingDelegateChanged)
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
		it.Event = new(SystemStakingDelegateChanged)
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
func (it *SystemStakingDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemStakingDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemStakingDelegateChanged represents a DelegateChanged event raised by the SystemStaking contract.
type SystemStakingDelegateChanged struct {
	TokenId     *big.Int
	NewDelegate common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x6f08c7e76d830d5f3d0a18fd27f4d8c0049b24a8689ddb39625e0864d894a9c1.
//
// Solidity: event DelegateChanged(uint256 indexed tokenId, address newDelegate)
func (_SystemStaking *SystemStakingFilterer) FilterDelegateChanged(opts *bind.FilterOpts, tokenId []*big.Int) (*SystemStakingDelegateChangedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SystemStaking.contract.FilterLogs(opts, "DelegateChanged", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SystemStakingDelegateChangedIterator{contract: _SystemStaking.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x6f08c7e76d830d5f3d0a18fd27f4d8c0049b24a8689ddb39625e0864d894a9c1.
//
// Solidity: event DelegateChanged(uint256 indexed tokenId, address newDelegate)
func (_SystemStaking *SystemStakingFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *SystemStakingDelegateChanged, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SystemStaking.contract.WatchLogs(opts, "DelegateChanged", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemStakingDelegateChanged)
				if err := _SystemStaking.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0x6f08c7e76d830d5f3d0a18fd27f4d8c0049b24a8689ddb39625e0864d894a9c1.
//
// Solidity: event DelegateChanged(uint256 indexed tokenId, address newDelegate)
func (_SystemStaking *SystemStakingFilterer) ParseDelegateChanged(log types.Log) (*SystemStakingDelegateChanged, error) {
	event := new(SystemStakingDelegateChanged)
	if err := _SystemStaking.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemStakingLockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the SystemStaking contract.
type SystemStakingLockedIterator struct {
	Event *SystemStakingLocked // Event containing the contract specifics and raw log

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
func (it *SystemStakingLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemStakingLocked)
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
		it.Event = new(SystemStakingLocked)
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
func (it *SystemStakingLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemStakingLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemStakingLocked represents a Locked event raised by the SystemStaking contract.
type SystemStakingLocked struct {
	TokenId  *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x907fece23ce39fbcbceb71e515043fe29408353fbb393b25b35eb8a70a4bad0b.
//
// Solidity: event Locked(uint256 indexed tokenId, uint256 duration)
func (_SystemStaking *SystemStakingFilterer) FilterLocked(opts *bind.FilterOpts, tokenId []*big.Int) (*SystemStakingLockedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SystemStaking.contract.FilterLogs(opts, "Locked", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SystemStakingLockedIterator{contract: _SystemStaking.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x907fece23ce39fbcbceb71e515043fe29408353fbb393b25b35eb8a70a4bad0b.
//
// Solidity: event Locked(uint256 indexed tokenId, uint256 duration)
func (_SystemStaking *SystemStakingFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *SystemStakingLocked, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SystemStaking.contract.WatchLogs(opts, "Locked", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemStakingLocked)
				if err := _SystemStaking.contract.UnpackLog(event, "Locked", log); err != nil {
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

// ParseLocked is a log parse operation binding the contract event 0x907fece23ce39fbcbceb71e515043fe29408353fbb393b25b35eb8a70a4bad0b.
//
// Solidity: event Locked(uint256 indexed tokenId, uint256 duration)
func (_SystemStaking *SystemStakingFilterer) ParseLocked(log types.Log) (*SystemStakingLocked, error) {
	event := new(SystemStakingLocked)
	if err := _SystemStaking.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemStakingMergedIterator is returned from FilterMerged and is used to iterate over the raw logs and unpacked data for Merged events raised by the SystemStaking contract.
type SystemStakingMergedIterator struct {
	Event *SystemStakingMerged // Event containing the contract specifics and raw log

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
func (it *SystemStakingMergedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemStakingMerged)
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
		it.Event = new(SystemStakingMerged)
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
func (it *SystemStakingMergedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemStakingMergedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemStakingMerged represents a Merged event raised by the SystemStaking contract.
type SystemStakingMerged struct {
	TokenIds []*big.Int
	Amount   *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMerged is a free log retrieval operation binding the contract event 0xb3f4c8ca702dbbd32d9a25ce17b1942a5060284d9d69fc4fcac8fb0397891b12.
//
// Solidity: event Merged(uint256[] tokenIds, uint256 amount, uint256 duration)
func (_SystemStaking *SystemStakingFilterer) FilterMerged(opts *bind.FilterOpts) (*SystemStakingMergedIterator, error) {

	logs, sub, err := _SystemStaking.contract.FilterLogs(opts, "Merged")
	if err != nil {
		return nil, err
	}
	return &SystemStakingMergedIterator{contract: _SystemStaking.contract, event: "Merged", logs: logs, sub: sub}, nil
}

// WatchMerged is a free log subscription operation binding the contract event 0xb3f4c8ca702dbbd32d9a25ce17b1942a5060284d9d69fc4fcac8fb0397891b12.
//
// Solidity: event Merged(uint256[] tokenIds, uint256 amount, uint256 duration)
func (_SystemStaking *SystemStakingFilterer) WatchMerged(opts *bind.WatchOpts, sink chan<- *SystemStakingMerged) (event.Subscription, error) {

	logs, sub, err := _SystemStaking.contract.WatchLogs(opts, "Merged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemStakingMerged)
				if err := _SystemStaking.contract.UnpackLog(event, "Merged", log); err != nil {
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

// ParseMerged is a log parse operation binding the contract event 0xb3f4c8ca702dbbd32d9a25ce17b1942a5060284d9d69fc4fcac8fb0397891b12.
//
// Solidity: event Merged(uint256[] tokenIds, uint256 amount, uint256 duration)
func (_SystemStaking *SystemStakingFilterer) ParseMerged(log types.Log) (*SystemStakingMerged, error) {
	event := new(SystemStakingMerged)
	if err := _SystemStaking.contract.UnpackLog(event, "Merged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemStakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SystemStaking contract.
type SystemStakingOwnershipTransferredIterator struct {
	Event *SystemStakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SystemStakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemStakingOwnershipTransferred)
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
		it.Event = new(SystemStakingOwnershipTransferred)
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
func (it *SystemStakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemStakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemStakingOwnershipTransferred represents a OwnershipTransferred event raised by the SystemStaking contract.
type SystemStakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemStaking *SystemStakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SystemStakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SystemStaking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SystemStakingOwnershipTransferredIterator{contract: _SystemStaking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemStaking *SystemStakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SystemStakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SystemStaking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemStakingOwnershipTransferred)
				if err := _SystemStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemStaking *SystemStakingFilterer) ParseOwnershipTransferred(log types.Log) (*SystemStakingOwnershipTransferred, error) {
	event := new(SystemStakingOwnershipTransferred)
	if err := _SystemStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemStakingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SystemStaking contract.
type SystemStakingPausedIterator struct {
	Event *SystemStakingPaused // Event containing the contract specifics and raw log

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
func (it *SystemStakingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemStakingPaused)
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
		it.Event = new(SystemStakingPaused)
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
func (it *SystemStakingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemStakingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemStakingPaused represents a Paused event raised by the SystemStaking contract.
type SystemStakingPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SystemStaking *SystemStakingFilterer) FilterPaused(opts *bind.FilterOpts) (*SystemStakingPausedIterator, error) {

	logs, sub, err := _SystemStaking.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SystemStakingPausedIterator{contract: _SystemStaking.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SystemStaking *SystemStakingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SystemStakingPaused) (event.Subscription, error) {

	logs, sub, err := _SystemStaking.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemStakingPaused)
				if err := _SystemStaking.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_SystemStaking *SystemStakingFilterer) ParsePaused(log types.Log) (*SystemStakingPaused, error) {
	event := new(SystemStakingPaused)
	if err := _SystemStaking.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemStakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the SystemStaking contract.
type SystemStakingStakedIterator struct {
	Event *SystemStakingStaked // Event containing the contract specifics and raw log

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
func (it *SystemStakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemStakingStaked)
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
		it.Event = new(SystemStakingStaked)
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
func (it *SystemStakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemStakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemStakingStaked represents a Staked event raised by the SystemStaking contract.
type SystemStakingStaked struct {
	TokenId  *big.Int
	Delegate common.Address
	Amount   *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x17700ceb1658b18206f427c1578048e87504106b14ec69e9b4586d9a95174a32.
//
// Solidity: event Staked(uint256 indexed tokenId, address delegate, uint256 amount, uint256 duration)
func (_SystemStaking *SystemStakingFilterer) FilterStaked(opts *bind.FilterOpts, tokenId []*big.Int) (*SystemStakingStakedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SystemStaking.contract.FilterLogs(opts, "Staked", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SystemStakingStakedIterator{contract: _SystemStaking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x17700ceb1658b18206f427c1578048e87504106b14ec69e9b4586d9a95174a32.
//
// Solidity: event Staked(uint256 indexed tokenId, address delegate, uint256 amount, uint256 duration)
func (_SystemStaking *SystemStakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *SystemStakingStaked, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SystemStaking.contract.WatchLogs(opts, "Staked", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemStakingStaked)
				if err := _SystemStaking.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x17700ceb1658b18206f427c1578048e87504106b14ec69e9b4586d9a95174a32.
//
// Solidity: event Staked(uint256 indexed tokenId, address delegate, uint256 amount, uint256 duration)
func (_SystemStaking *SystemStakingFilterer) ParseStaked(log types.Log) (*SystemStakingStaked, error) {
	event := new(SystemStakingStaked)
	if err := _SystemStaking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemStakingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SystemStaking contract.
type SystemStakingTransferIterator struct {
	Event *SystemStakingTransfer // Event containing the contract specifics and raw log

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
func (it *SystemStakingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemStakingTransfer)
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
		it.Event = new(SystemStakingTransfer)
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
func (it *SystemStakingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemStakingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemStakingTransfer represents a Transfer event raised by the SystemStaking contract.
type SystemStakingTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SystemStaking *SystemStakingFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SystemStakingTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SystemStaking.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SystemStakingTransferIterator{contract: _SystemStaking.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SystemStaking *SystemStakingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SystemStakingTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SystemStaking.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemStakingTransfer)
				if err := _SystemStaking.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SystemStaking *SystemStakingFilterer) ParseTransfer(log types.Log) (*SystemStakingTransfer, error) {
	event := new(SystemStakingTransfer)
	if err := _SystemStaking.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemStakingUnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the SystemStaking contract.
type SystemStakingUnlockedIterator struct {
	Event *SystemStakingUnlocked // Event containing the contract specifics and raw log

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
func (it *SystemStakingUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemStakingUnlocked)
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
		it.Event = new(SystemStakingUnlocked)
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
func (it *SystemStakingUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemStakingUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemStakingUnlocked represents a Unlocked event raised by the SystemStaking contract.
type SystemStakingUnlocked struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0xf27b6ce5b2f5e68ddb2fd95a8a909d4ecf1daaac270935fff052feacb24f1842.
//
// Solidity: event Unlocked(uint256 indexed tokenId)
func (_SystemStaking *SystemStakingFilterer) FilterUnlocked(opts *bind.FilterOpts, tokenId []*big.Int) (*SystemStakingUnlockedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SystemStaking.contract.FilterLogs(opts, "Unlocked", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SystemStakingUnlockedIterator{contract: _SystemStaking.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0xf27b6ce5b2f5e68ddb2fd95a8a909d4ecf1daaac270935fff052feacb24f1842.
//
// Solidity: event Unlocked(uint256 indexed tokenId)
func (_SystemStaking *SystemStakingFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *SystemStakingUnlocked, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SystemStaking.contract.WatchLogs(opts, "Unlocked", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemStakingUnlocked)
				if err := _SystemStaking.contract.UnpackLog(event, "Unlocked", log); err != nil {
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

// ParseUnlocked is a log parse operation binding the contract event 0xf27b6ce5b2f5e68ddb2fd95a8a909d4ecf1daaac270935fff052feacb24f1842.
//
// Solidity: event Unlocked(uint256 indexed tokenId)
func (_SystemStaking *SystemStakingFilterer) ParseUnlocked(log types.Log) (*SystemStakingUnlocked, error) {
	event := new(SystemStakingUnlocked)
	if err := _SystemStaking.contract.UnpackLog(event, "Unlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemStakingUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SystemStaking contract.
type SystemStakingUnpausedIterator struct {
	Event *SystemStakingUnpaused // Event containing the contract specifics and raw log

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
func (it *SystemStakingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemStakingUnpaused)
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
		it.Event = new(SystemStakingUnpaused)
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
func (it *SystemStakingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemStakingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemStakingUnpaused represents a Unpaused event raised by the SystemStaking contract.
type SystemStakingUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SystemStaking *SystemStakingFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SystemStakingUnpausedIterator, error) {

	logs, sub, err := _SystemStaking.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SystemStakingUnpausedIterator{contract: _SystemStaking.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SystemStaking *SystemStakingFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SystemStakingUnpaused) (event.Subscription, error) {

	logs, sub, err := _SystemStaking.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemStakingUnpaused)
				if err := _SystemStaking.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_SystemStaking *SystemStakingFilterer) ParseUnpaused(log types.Log) (*SystemStakingUnpaused, error) {
	event := new(SystemStakingUnpaused)
	if err := _SystemStaking.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemStakingUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the SystemStaking contract.
type SystemStakingUnstakedIterator struct {
	Event *SystemStakingUnstaked // Event containing the contract specifics and raw log

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
func (it *SystemStakingUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemStakingUnstaked)
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
		it.Event = new(SystemStakingUnstaked)
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
func (it *SystemStakingUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemStakingUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemStakingUnstaked represents a Unstaked event raised by the SystemStaking contract.
type SystemStakingUnstaked struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x11725367022c3ff288940f4b5473aa61c2da6a24af7363a1128ee2401e8983b2.
//
// Solidity: event Unstaked(uint256 indexed tokenId)
func (_SystemStaking *SystemStakingFilterer) FilterUnstaked(opts *bind.FilterOpts, tokenId []*big.Int) (*SystemStakingUnstakedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SystemStaking.contract.FilterLogs(opts, "Unstaked", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SystemStakingUnstakedIterator{contract: _SystemStaking.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x11725367022c3ff288940f4b5473aa61c2da6a24af7363a1128ee2401e8983b2.
//
// Solidity: event Unstaked(uint256 indexed tokenId)
func (_SystemStaking *SystemStakingFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *SystemStakingUnstaked, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SystemStaking.contract.WatchLogs(opts, "Unstaked", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemStakingUnstaked)
				if err := _SystemStaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x11725367022c3ff288940f4b5473aa61c2da6a24af7363a1128ee2401e8983b2.
//
// Solidity: event Unstaked(uint256 indexed tokenId)
func (_SystemStaking *SystemStakingFilterer) ParseUnstaked(log types.Log) (*SystemStakingUnstaked, error) {
	event := new(SystemStakingUnstaked)
	if err := _SystemStaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemStakingWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the SystemStaking contract.
type SystemStakingWithdrawalIterator struct {
	Event *SystemStakingWithdrawal // Event containing the contract specifics and raw log

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
func (it *SystemStakingWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemStakingWithdrawal)
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
		it.Event = new(SystemStakingWithdrawal)
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
func (it *SystemStakingWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemStakingWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemStakingWithdrawal represents a Withdrawal event raised by the SystemStaking contract.
type SystemStakingWithdrawal struct {
	TokenId   *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xd964a27d45f595739c13d8b1160b57491050cacf3a2e5602207277d6228f64ee.
//
// Solidity: event Withdrawal(uint256 indexed tokenId, address indexed recipient)
func (_SystemStaking *SystemStakingFilterer) FilterWithdrawal(opts *bind.FilterOpts, tokenId []*big.Int, recipient []common.Address) (*SystemStakingWithdrawalIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SystemStaking.contract.FilterLogs(opts, "Withdrawal", tokenIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &SystemStakingWithdrawalIterator{contract: _SystemStaking.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xd964a27d45f595739c13d8b1160b57491050cacf3a2e5602207277d6228f64ee.
//
// Solidity: event Withdrawal(uint256 indexed tokenId, address indexed recipient)
func (_SystemStaking *SystemStakingFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *SystemStakingWithdrawal, tokenId []*big.Int, recipient []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SystemStaking.contract.WatchLogs(opts, "Withdrawal", tokenIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemStakingWithdrawal)
				if err := _SystemStaking.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xd964a27d45f595739c13d8b1160b57491050cacf3a2e5602207277d6228f64ee.
//
// Solidity: event Withdrawal(uint256 indexed tokenId, address indexed recipient)
func (_SystemStaking *SystemStakingFilterer) ParseWithdrawal(log types.Log) (*SystemStakingWithdrawal, error) {
	event := new(SystemStakingWithdrawal)
	if err := _SystemStaking.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
