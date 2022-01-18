// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
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

// HeightData is an auto generated low-level Go binding around an user-defined struct.

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CREATE_CLIENT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTER_RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_CLIENT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessManager\",\"outputs\":[{\"internalType\":\"contractIAccessManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"clients\",\"outputs\":[{\"internalType\":\"contractIClient\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"clientAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"clientState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusState\",\"type\":\"bytes\"}],\"name\":\"createClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"}],\"name\":\"getClient\",\"outputs\":[{\"internalType\":\"contractIClient\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"}],\"name\":\"getLatestHeight\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"accessManagerContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"registerRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"relayers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"updateClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"clientState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusState\",\"type\":\"bytes\"}],\"name\":\"upgradeClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// CREATECLIENTROLE is a free data retrieval call binding the contract method 0xafde32ba.
//
// Solidity: function CREATE_CLIENT_ROLE() view returns(bytes32)
func (_Contracts *ContractsCaller) CREATECLIENTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "CREATE_CLIENT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CREATECLIENTROLE is a free data retrieval call binding the contract method 0xafde32ba.
//
// Solidity: function CREATE_CLIENT_ROLE() view returns(bytes32)
func (_Contracts *ContractsSession) CREATECLIENTROLE() ([32]byte, error) {
	return _Contracts.Contract.CREATECLIENTROLE(&_Contracts.CallOpts)
}

// CREATECLIENTROLE is a free data retrieval call binding the contract method 0xafde32ba.
//
// Solidity: function CREATE_CLIENT_ROLE() view returns(bytes32)
func (_Contracts *ContractsCallerSession) CREATECLIENTROLE() ([32]byte, error) {
	return _Contracts.Contract.CREATECLIENTROLE(&_Contracts.CallOpts)
}

// REGISTERRELAYERROLE is a free data retrieval call binding the contract method 0x1a371a90.
//
// Solidity: function REGISTER_RELAYER_ROLE() view returns(bytes32)
func (_Contracts *ContractsCaller) REGISTERRELAYERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "REGISTER_RELAYER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGISTERRELAYERROLE is a free data retrieval call binding the contract method 0x1a371a90.
//
// Solidity: function REGISTER_RELAYER_ROLE() view returns(bytes32)
func (_Contracts *ContractsSession) REGISTERRELAYERROLE() ([32]byte, error) {
	return _Contracts.Contract.REGISTERRELAYERROLE(&_Contracts.CallOpts)
}

// REGISTERRELAYERROLE is a free data retrieval call binding the contract method 0x1a371a90.
//
// Solidity: function REGISTER_RELAYER_ROLE() view returns(bytes32)
func (_Contracts *ContractsCallerSession) REGISTERRELAYERROLE() ([32]byte, error) {
	return _Contracts.Contract.REGISTERRELAYERROLE(&_Contracts.CallOpts)
}

// UPGRADECLIENTROLE is a free data retrieval call binding the contract method 0x8752ccf5.
//
// Solidity: function UPGRADE_CLIENT_ROLE() view returns(bytes32)
func (_Contracts *ContractsCaller) UPGRADECLIENTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "UPGRADE_CLIENT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADECLIENTROLE is a free data retrieval call binding the contract method 0x8752ccf5.
//
// Solidity: function UPGRADE_CLIENT_ROLE() view returns(bytes32)
func (_Contracts *ContractsSession) UPGRADECLIENTROLE() ([32]byte, error) {
	return _Contracts.Contract.UPGRADECLIENTROLE(&_Contracts.CallOpts)
}

// UPGRADECLIENTROLE is a free data retrieval call binding the contract method 0x8752ccf5.
//
// Solidity: function UPGRADE_CLIENT_ROLE() view returns(bytes32)
func (_Contracts *ContractsCallerSession) UPGRADECLIENTROLE() ([32]byte, error) {
	return _Contracts.Contract.UPGRADECLIENTROLE(&_Contracts.CallOpts)
}

// AccessManager is a free data retrieval call binding the contract method 0xfdcb6068.
//
// Solidity: function accessManager() view returns(address)
func (_Contracts *ContractsCaller) AccessManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "accessManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccessManager is a free data retrieval call binding the contract method 0xfdcb6068.
//
// Solidity: function accessManager() view returns(address)
func (_Contracts *ContractsSession) AccessManager() (common.Address, error) {
	return _Contracts.Contract.AccessManager(&_Contracts.CallOpts)
}

// AccessManager is a free data retrieval call binding the contract method 0xfdcb6068.
//
// Solidity: function accessManager() view returns(address)
func (_Contracts *ContractsCallerSession) AccessManager() (common.Address, error) {
	return _Contracts.Contract.AccessManager(&_Contracts.CallOpts)
}

// Clients is a free data retrieval call binding the contract method 0x20ba1e9f.
//
// Solidity: function clients(string ) view returns(address)
func (_Contracts *ContractsCaller) Clients(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "clients", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Clients is a free data retrieval call binding the contract method 0x20ba1e9f.
//
// Solidity: function clients(string ) view returns(address)
func (_Contracts *ContractsSession) Clients(arg0 string) (common.Address, error) {
	return _Contracts.Contract.Clients(&_Contracts.CallOpts, arg0)
}

// Clients is a free data retrieval call binding the contract method 0x20ba1e9f.
//
// Solidity: function clients(string ) view returns(address)
func (_Contracts *ContractsCallerSession) Clients(arg0 string) (common.Address, error) {
	return _Contracts.Contract.Clients(&_Contracts.CallOpts, arg0)
}

// GetChainName is a free data retrieval call binding the contract method 0xd722b0bc.
//
// Solidity: function getChainName() view returns(string)
func (_Contracts *ContractsCaller) GetChainName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getChainName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetChainName is a free data retrieval call binding the contract method 0xd722b0bc.
//
// Solidity: function getChainName() view returns(string)
func (_Contracts *ContractsSession) GetChainName() (string, error) {
	return _Contracts.Contract.GetChainName(&_Contracts.CallOpts)
}

// GetChainName is a free data retrieval call binding the contract method 0xd722b0bc.
//
// Solidity: function getChainName() view returns(string)
func (_Contracts *ContractsCallerSession) GetChainName() (string, error) {
	return _Contracts.Contract.GetChainName(&_Contracts.CallOpts)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string chainName) view returns((uint64,uint64))
func (_Contracts *ContractsCaller) GetLatestHeight(opts *bind.CallOpts, chainName string) (HeightData, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLatestHeight", chainName)

	if err != nil {
		return *new(HeightData), err
	}

	out0 := *abi.ConvertType(out[0], new(HeightData)).(*HeightData)

	return out0, err

}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string chainName) view returns((uint64,uint64))
func (_Contracts *ContractsSession) GetLatestHeight(chainName string) (HeightData, error) {
	return _Contracts.Contract.GetLatestHeight(&_Contracts.CallOpts, chainName)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string chainName) view returns((uint64,uint64))
func (_Contracts *ContractsCallerSession) GetLatestHeight(chainName string) (HeightData, error) {
	return _Contracts.Contract.GetLatestHeight(&_Contracts.CallOpts, chainName)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Relayers is a free data retrieval call binding the contract method 0xee1ceb62.
//
// Solidity: function relayers(string , address ) view returns(bool)
func (_Contracts *ContractsCaller) Relayers(opts *bind.CallOpts, arg0 string, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "relayers", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Relayers is a free data retrieval call binding the contract method 0xee1ceb62.
//
// Solidity: function relayers(string , address ) view returns(bool)
func (_Contracts *ContractsSession) Relayers(arg0 string, arg1 common.Address) (bool, error) {
	return _Contracts.Contract.Relayers(&_Contracts.CallOpts, arg0, arg1)
}

// Relayers is a free data retrieval call binding the contract method 0xee1ceb62.
//
// Solidity: function relayers(string , address ) view returns(bool)
func (_Contracts *ContractsCallerSession) Relayers(arg0 string, arg1 common.Address) (bool, error) {
	return _Contracts.Contract.Relayers(&_Contracts.CallOpts, arg0, arg1)
}

// CreateClient is a paid mutator transaction binding the contract method 0x76262a47.
//
// Solidity: function createClient(string chainName, address clientAddress, bytes clientState, bytes consensusState) returns()
func (_Contracts *ContractsTransactor) CreateClient(opts *bind.TransactOpts, chainName string, clientAddress common.Address, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createClient", chainName, clientAddress, clientState, consensusState)
}

// CreateClient is a paid mutator transaction binding the contract method 0x76262a47.
//
// Solidity: function createClient(string chainName, address clientAddress, bytes clientState, bytes consensusState) returns()
func (_Contracts *ContractsSession) CreateClient(chainName string, clientAddress common.Address, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Contracts.Contract.CreateClient(&_Contracts.TransactOpts, chainName, clientAddress, clientState, consensusState)
}

// CreateClient is a paid mutator transaction binding the contract method 0x76262a47.
//
// Solidity: function createClient(string chainName, address clientAddress, bytes clientState, bytes consensusState) returns()
func (_Contracts *ContractsTransactorSession) CreateClient(chainName string, clientAddress common.Address, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Contracts.Contract.CreateClient(&_Contracts.TransactOpts, chainName, clientAddress, clientState, consensusState)
}

// GetClient is a paid mutator transaction binding the contract method 0x7eb78932.
//
// Solidity: function getClient(string chainName) returns(address)
func (_Contracts *ContractsTransactor) GetClient(opts *bind.TransactOpts, chainName string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "getClient", chainName)
}

// GetClient is a paid mutator transaction binding the contract method 0x7eb78932.
//
// Solidity: function getClient(string chainName) returns(address)
func (_Contracts *ContractsSession) GetClient(chainName string) (*types.Transaction, error) {
	return _Contracts.Contract.GetClient(&_Contracts.TransactOpts, chainName)
}

// GetClient is a paid mutator transaction binding the contract method 0x7eb78932.
//
// Solidity: function getClient(string chainName) returns(address)
func (_Contracts *ContractsTransactorSession) GetClient(chainName string) (*types.Transaction, error) {
	return _Contracts.Contract.GetClient(&_Contracts.TransactOpts, chainName)
}

// Initialize is a paid mutator transaction binding the contract method 0x7ab4339d.
//
// Solidity: function initialize(string name, address accessManagerContract) returns()
func (_Contracts *ContractsTransactor) Initialize(opts *bind.TransactOpts, name string, accessManagerContract common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "initialize", name, accessManagerContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x7ab4339d.
//
// Solidity: function initialize(string name, address accessManagerContract) returns()
func (_Contracts *ContractsSession) Initialize(name string, accessManagerContract common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, name, accessManagerContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x7ab4339d.
//
// Solidity: function initialize(string name, address accessManagerContract) returns()
func (_Contracts *ContractsTransactorSession) Initialize(name string, accessManagerContract common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, name, accessManagerContract)
}

// RegisterRelayer is a paid mutator transaction binding the contract method 0x5330a758.
//
// Solidity: function registerRelayer(string chainName, address relayer) returns()
func (_Contracts *ContractsTransactor) RegisterRelayer(opts *bind.TransactOpts, chainName string, relayer common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerRelayer", chainName, relayer)
}

// RegisterRelayer is a paid mutator transaction binding the contract method 0x5330a758.
//
// Solidity: function registerRelayer(string chainName, address relayer) returns()
func (_Contracts *ContractsSession) RegisterRelayer(chainName string, relayer common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterRelayer(&_Contracts.TransactOpts, chainName, relayer)
}

// RegisterRelayer is a paid mutator transaction binding the contract method 0x5330a758.
//
// Solidity: function registerRelayer(string chainName, address relayer) returns()
func (_Contracts *ContractsTransactorSession) RegisterRelayer(chainName string, relayer common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterRelayer(&_Contracts.TransactOpts, chainName, relayer)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x6fbf8079.
//
// Solidity: function updateClient(string chainName, bytes header) returns()
func (_Contracts *ContractsTransactor) UpdateClient(opts *bind.TransactOpts, chainName string, header []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateClient", chainName, header)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x6fbf8079.
//
// Solidity: function updateClient(string chainName, bytes header) returns()
func (_Contracts *ContractsSession) UpdateClient(chainName string, header []byte) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateClient(&_Contracts.TransactOpts, chainName, header)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x6fbf8079.
//
// Solidity: function updateClient(string chainName, bytes header) returns()
func (_Contracts *ContractsTransactorSession) UpdateClient(chainName string, header []byte) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateClient(&_Contracts.TransactOpts, chainName, header)
}

// UpgradeClient is a paid mutator transaction binding the contract method 0x935aee64.
//
// Solidity: function upgradeClient(string chainName, bytes clientState, bytes consensusState) returns()
func (_Contracts *ContractsTransactor) UpgradeClient(opts *bind.TransactOpts, chainName string, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "upgradeClient", chainName, clientState, consensusState)
}

// UpgradeClient is a paid mutator transaction binding the contract method 0x935aee64.
//
// Solidity: function upgradeClient(string chainName, bytes clientState, bytes consensusState) returns()
func (_Contracts *ContractsSession) UpgradeClient(chainName string, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Contracts.Contract.UpgradeClient(&_Contracts.TransactOpts, chainName, clientState, consensusState)
}

// UpgradeClient is a paid mutator transaction binding the contract method 0x935aee64.
//
// Solidity: function upgradeClient(string chainName, bytes clientState, bytes consensusState) returns()
func (_Contracts *ContractsTransactorSession) UpgradeClient(chainName string, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Contracts.Contract.UpgradeClient(&_Contracts.TransactOpts, chainName, clientState, consensusState)
}

// ContractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contracts contract.
type ContractsOwnershipTransferredIterator struct {
	Event *ContractsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnershipTransferred)
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
		it.Event = new(ContractsOwnershipTransferred)
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
func (it *ContractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnershipTransferred represents a OwnershipTransferred event raised by the Contracts contract.
type ContractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOwnershipTransferredIterator{contract: _Contracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnershipTransferred)
				if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseOwnershipTransferred(log types.Log) (*ContractsOwnershipTransferred, error) {
	event := new(ContractsOwnershipTransferred)
	if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
