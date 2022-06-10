// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package client

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

// HeightData is an auto generated low-level Go binding around an user-defined struct.
type HeightData struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// ClientMetaData contains all meta data concerning the Client contract.
var ClientMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CREATE_CLIENT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_CLIENT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessManager\",\"outputs\":[{\"internalType\":\"contractIAccessManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"client\",\"outputs\":[{\"internalType\":\"contractIClient\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"clientAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"clientState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusState\",\"type\":\"bytes\"}],\"name\":\"createClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClientType\",\"outputs\":[{\"internalType\":\"enumIClient.Type\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestHeight\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accessManagerContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"clientAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"clientState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusState\",\"type\":\"bytes\"}],\"name\":\"toggleClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"updateClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"clientState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusState\",\"type\":\"bytes\"}],\"name\":\"upgradeClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ClientABI is the input ABI used to generate the binding from.
// Deprecated: Use ClientMetaData.ABI instead.
var ClientABI = ClientMetaData.ABI

// Client is an auto generated Go binding around an Ethereum contract.
type Client struct {
	ClientCaller     // Read-only binding to the contract
	ClientTransactor // Write-only binding to the contract
	ClientFilterer   // Log filterer for contract events
}

// ClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClientSession struct {
	Contract     *Client           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClientCallerSession struct {
	Contract *ClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClientTransactorSession struct {
	Contract     *ClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClientRaw struct {
	Contract *Client // Generic contract binding to access the raw methods on
}

// ClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClientCallerRaw struct {
	Contract *ClientCaller // Generic read-only contract binding to access the raw methods on
}

// ClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClientTransactorRaw struct {
	Contract *ClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClient creates a new instance of Client, bound to a specific deployed contract.
func NewClient(address common.Address, backend bind.ContractBackend) (*Client, error) {
	contract, err := bindClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Client{ClientCaller: ClientCaller{contract: contract}, ClientTransactor: ClientTransactor{contract: contract}, ClientFilterer: ClientFilterer{contract: contract}}, nil
}

// NewClientCaller creates a new read-only instance of Client, bound to a specific deployed contract.
func NewClientCaller(address common.Address, caller bind.ContractCaller) (*ClientCaller, error) {
	contract, err := bindClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClientCaller{contract: contract}, nil
}

// NewClientTransactor creates a new write-only instance of Client, bound to a specific deployed contract.
func NewClientTransactor(address common.Address, transactor bind.ContractTransactor) (*ClientTransactor, error) {
	contract, err := bindClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClientTransactor{contract: contract}, nil
}

// NewClientFilterer creates a new log filterer instance of Client, bound to a specific deployed contract.
func NewClientFilterer(address common.Address, filterer bind.ContractFilterer) (*ClientFilterer, error) {
	contract, err := bindClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClientFilterer{contract: contract}, nil
}

// bindClient binds a generic wrapper to an already deployed contract.
func bindClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Client *ClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Client.Contract.ClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Client *ClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Client.Contract.ClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Client *ClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Client.Contract.ClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Client *ClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Client.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Client *ClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Client.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Client *ClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Client.Contract.contract.Transact(opts, method, params...)
}

// CREATECLIENTROLE is a free data retrieval call binding the contract method 0xafde32ba.
//
// Solidity: function CREATE_CLIENT_ROLE() view returns(bytes32)
func (_Client *ClientCaller) CREATECLIENTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Client.contract.Call(opts, &out, "CREATE_CLIENT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CREATECLIENTROLE is a free data retrieval call binding the contract method 0xafde32ba.
//
// Solidity: function CREATE_CLIENT_ROLE() view returns(bytes32)
func (_Client *ClientSession) CREATECLIENTROLE() ([32]byte, error) {
	return _Client.Contract.CREATECLIENTROLE(&_Client.CallOpts)
}

// CREATECLIENTROLE is a free data retrieval call binding the contract method 0xafde32ba.
//
// Solidity: function CREATE_CLIENT_ROLE() view returns(bytes32)
func (_Client *ClientCallerSession) CREATECLIENTROLE() ([32]byte, error) {
	return _Client.Contract.CREATECLIENTROLE(&_Client.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Client *ClientCaller) RELAYERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Client.contract.Call(opts, &out, "RELAYER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Client *ClientSession) RELAYERROLE() ([32]byte, error) {
	return _Client.Contract.RELAYERROLE(&_Client.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Client *ClientCallerSession) RELAYERROLE() ([32]byte, error) {
	return _Client.Contract.RELAYERROLE(&_Client.CallOpts)
}

// UPGRADECLIENTROLE is a free data retrieval call binding the contract method 0x8752ccf5.
//
// Solidity: function UPGRADE_CLIENT_ROLE() view returns(bytes32)
func (_Client *ClientCaller) UPGRADECLIENTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Client.contract.Call(opts, &out, "UPGRADE_CLIENT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADECLIENTROLE is a free data retrieval call binding the contract method 0x8752ccf5.
//
// Solidity: function UPGRADE_CLIENT_ROLE() view returns(bytes32)
func (_Client *ClientSession) UPGRADECLIENTROLE() ([32]byte, error) {
	return _Client.Contract.UPGRADECLIENTROLE(&_Client.CallOpts)
}

// UPGRADECLIENTROLE is a free data retrieval call binding the contract method 0x8752ccf5.
//
// Solidity: function UPGRADE_CLIENT_ROLE() view returns(bytes32)
func (_Client *ClientCallerSession) UPGRADECLIENTROLE() ([32]byte, error) {
	return _Client.Contract.UPGRADECLIENTROLE(&_Client.CallOpts)
}

// AccessManager is a free data retrieval call binding the contract method 0xfdcb6068.
//
// Solidity: function accessManager() view returns(address)
func (_Client *ClientCaller) AccessManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Client.contract.Call(opts, &out, "accessManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccessManager is a free data retrieval call binding the contract method 0xfdcb6068.
//
// Solidity: function accessManager() view returns(address)
func (_Client *ClientSession) AccessManager() (common.Address, error) {
	return _Client.Contract.AccessManager(&_Client.CallOpts)
}

// AccessManager is a free data retrieval call binding the contract method 0xfdcb6068.
//
// Solidity: function accessManager() view returns(address)
func (_Client *ClientCallerSession) AccessManager() (common.Address, error) {
	return _Client.Contract.AccessManager(&_Client.CallOpts)
}

// Client is a free data retrieval call binding the contract method 0x109e94cf.
//
// Solidity: function client() view returns(address)
func (_Client *ClientCaller) Client(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Client.contract.Call(opts, &out, "client")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Client is a free data retrieval call binding the contract method 0x109e94cf.
//
// Solidity: function client() view returns(address)
func (_Client *ClientSession) Client() (common.Address, error) {
	return _Client.Contract.Client(&_Client.CallOpts)
}

// Client is a free data retrieval call binding the contract method 0x109e94cf.
//
// Solidity: function client() view returns(address)
func (_Client *ClientCallerSession) Client() (common.Address, error) {
	return _Client.Contract.Client(&_Client.CallOpts)
}

// GetClientType is a free data retrieval call binding the contract method 0xaac909b8.
//
// Solidity: function getClientType() view returns(uint8)
func (_Client *ClientCaller) GetClientType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Client.contract.Call(opts, &out, "getClientType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetClientType is a free data retrieval call binding the contract method 0xaac909b8.
//
// Solidity: function getClientType() view returns(uint8)
func (_Client *ClientSession) GetClientType() (uint8, error) {
	return _Client.Contract.GetClientType(&_Client.CallOpts)
}

// GetClientType is a free data retrieval call binding the contract method 0xaac909b8.
//
// Solidity: function getClientType() view returns(uint8)
func (_Client *ClientCallerSession) GetClientType() (uint8, error) {
	return _Client.Contract.GetClientType(&_Client.CallOpts)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x4ed1d8cc.
//
// Solidity: function getLatestHeight() view returns((uint64,uint64))
func (_Client *ClientCaller) GetLatestHeight(opts *bind.CallOpts) (HeightData, error) {
	var out []interface{}
	err := _Client.contract.Call(opts, &out, "getLatestHeight")

	if err != nil {
		return *new(HeightData), err
	}

	out0 := *abi.ConvertType(out[0], new(HeightData)).(*HeightData)

	return out0, err

}

// GetLatestHeight is a free data retrieval call binding the contract method 0x4ed1d8cc.
//
// Solidity: function getLatestHeight() view returns((uint64,uint64))
func (_Client *ClientSession) GetLatestHeight() (HeightData, error) {
	return _Client.Contract.GetLatestHeight(&_Client.CallOpts)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x4ed1d8cc.
//
// Solidity: function getLatestHeight() view returns((uint64,uint64))
func (_Client *ClientCallerSession) GetLatestHeight() (HeightData, error) {
	return _Client.Contract.GetLatestHeight(&_Client.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Client *ClientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Client.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Client *ClientSession) Owner() (common.Address, error) {
	return _Client.Contract.Owner(&_Client.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Client *ClientCallerSession) Owner() (common.Address, error) {
	return _Client.Contract.Owner(&_Client.CallOpts)
}

// CreateClient is a paid mutator transaction binding the contract method 0x83a3d97e.
//
// Solidity: function createClient(address clientAddress, bytes clientState, bytes consensusState) returns()
func (_Client *ClientTransactor) CreateClient(opts *bind.TransactOpts, clientAddress common.Address, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "createClient", clientAddress, clientState, consensusState)
}

// CreateClient is a paid mutator transaction binding the contract method 0x83a3d97e.
//
// Solidity: function createClient(address clientAddress, bytes clientState, bytes consensusState) returns()
func (_Client *ClientSession) CreateClient(clientAddress common.Address, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Client.Contract.CreateClient(&_Client.TransactOpts, clientAddress, clientState, consensusState)
}

// CreateClient is a paid mutator transaction binding the contract method 0x83a3d97e.
//
// Solidity: function createClient(address clientAddress, bytes clientState, bytes consensusState) returns()
func (_Client *ClientTransactorSession) CreateClient(clientAddress common.Address, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Client.Contract.CreateClient(&_Client.TransactOpts, clientAddress, clientState, consensusState)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address accessManagerContract) returns()
func (_Client *ClientTransactor) Initialize(opts *bind.TransactOpts, accessManagerContract common.Address) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "initialize", accessManagerContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address accessManagerContract) returns()
func (_Client *ClientSession) Initialize(accessManagerContract common.Address) (*types.Transaction, error) {
	return _Client.Contract.Initialize(&_Client.TransactOpts, accessManagerContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address accessManagerContract) returns()
func (_Client *ClientTransactorSession) Initialize(accessManagerContract common.Address) (*types.Transaction, error) {
	return _Client.Contract.Initialize(&_Client.TransactOpts, accessManagerContract)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Client *ClientTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Client *ClientSession) RenounceOwnership() (*types.Transaction, error) {
	return _Client.Contract.RenounceOwnership(&_Client.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Client *ClientTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Client.Contract.RenounceOwnership(&_Client.TransactOpts)
}

// ToggleClient is a paid mutator transaction binding the contract method 0x4bd1bf5a.
//
// Solidity: function toggleClient(address clientAddress, bytes clientState, bytes consensusState) returns()
func (_Client *ClientTransactor) ToggleClient(opts *bind.TransactOpts, clientAddress common.Address, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "toggleClient", clientAddress, clientState, consensusState)
}

// ToggleClient is a paid mutator transaction binding the contract method 0x4bd1bf5a.
//
// Solidity: function toggleClient(address clientAddress, bytes clientState, bytes consensusState) returns()
func (_Client *ClientSession) ToggleClient(clientAddress common.Address, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Client.Contract.ToggleClient(&_Client.TransactOpts, clientAddress, clientState, consensusState)
}

// ToggleClient is a paid mutator transaction binding the contract method 0x4bd1bf5a.
//
// Solidity: function toggleClient(address clientAddress, bytes clientState, bytes consensusState) returns()
func (_Client *ClientTransactorSession) ToggleClient(clientAddress common.Address, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Client.Contract.ToggleClient(&_Client.TransactOpts, clientAddress, clientState, consensusState)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Client *ClientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Client *ClientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Client.Contract.TransferOwnership(&_Client.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Client *ClientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Client.Contract.TransferOwnership(&_Client.TransactOpts, newOwner)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x0bece356.
//
// Solidity: function updateClient(bytes header) returns()
func (_Client *ClientTransactor) UpdateClient(opts *bind.TransactOpts, header []byte) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "updateClient", header)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x0bece356.
//
// Solidity: function updateClient(bytes header) returns()
func (_Client *ClientSession) UpdateClient(header []byte) (*types.Transaction, error) {
	return _Client.Contract.UpdateClient(&_Client.TransactOpts, header)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x0bece356.
//
// Solidity: function updateClient(bytes header) returns()
func (_Client *ClientTransactorSession) UpdateClient(header []byte) (*types.Transaction, error) {
	return _Client.Contract.UpdateClient(&_Client.TransactOpts, header)
}

// UpgradeClient is a paid mutator transaction binding the contract method 0xf59e5c29.
//
// Solidity: function upgradeClient(bytes clientState, bytes consensusState) returns()
func (_Client *ClientTransactor) UpgradeClient(opts *bind.TransactOpts, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "upgradeClient", clientState, consensusState)
}

// UpgradeClient is a paid mutator transaction binding the contract method 0xf59e5c29.
//
// Solidity: function upgradeClient(bytes clientState, bytes consensusState) returns()
func (_Client *ClientSession) UpgradeClient(clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Client.Contract.UpgradeClient(&_Client.TransactOpts, clientState, consensusState)
}

// UpgradeClient is a paid mutator transaction binding the contract method 0xf59e5c29.
//
// Solidity: function upgradeClient(bytes clientState, bytes consensusState) returns()
func (_Client *ClientTransactorSession) UpgradeClient(clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Client.Contract.UpgradeClient(&_Client.TransactOpts, clientState, consensusState)
}

// ClientInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Client contract.
type ClientInitializedIterator struct {
	Event *ClientInitialized // Event containing the contract specifics and raw log

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
func (it *ClientInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClientInitialized)
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
		it.Event = new(ClientInitialized)
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
func (it *ClientInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClientInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClientInitialized represents a Initialized event raised by the Client contract.
type ClientInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Client *ClientFilterer) FilterInitialized(opts *bind.FilterOpts) (*ClientInitializedIterator, error) {

	logs, sub, err := _Client.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ClientInitializedIterator{contract: _Client.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Client *ClientFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ClientInitialized) (event.Subscription, error) {

	logs, sub, err := _Client.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClientInitialized)
				if err := _Client.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Client *ClientFilterer) ParseInitialized(log types.Log) (*ClientInitialized, error) {
	event := new(ClientInitialized)
	if err := _Client.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClientOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Client contract.
type ClientOwnershipTransferredIterator struct {
	Event *ClientOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ClientOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClientOwnershipTransferred)
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
		it.Event = new(ClientOwnershipTransferred)
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
func (it *ClientOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClientOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClientOwnershipTransferred represents a OwnershipTransferred event raised by the Client contract.
type ClientOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Client *ClientFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ClientOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Client.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ClientOwnershipTransferredIterator{contract: _Client.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Client *ClientFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ClientOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Client.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClientOwnershipTransferred)
				if err := _Client.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Client *ClientFilterer) ParseOwnershipTransferred(log types.Log) (*ClientOwnershipTransferred, error) {
	event := new(ClientOwnershipTransferred)
	if err := _Client.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
