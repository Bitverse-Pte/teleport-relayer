// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// PacketTypesPacket is an auto generated low-level Go binding around an user-defined struct.
type PacketTypesPacket struct {
	Sequence    uint64
	SourceChain string
	DestChain   string
	RelayChain  string
	Ports       []string
	DataList    [][]byte
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"ports\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"dataList\",\"type\":\"bytes[]\"}],\"indexed\":false,\"internalType\":\"structPacketTypes.Packet\",\"name\":\"packet\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ack\",\"type\":\"bytes\"}],\"name\":\"AckPacket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"ports\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"dataList\",\"type\":\"bytes[]\"}],\"indexed\":false,\"internalType\":\"structPacketTypes.Packet\",\"name\":\"packet\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ack\",\"type\":\"bytes\"}],\"name\":\"AckWritten\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"ports\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"dataList\",\"type\":\"bytes[]\"}],\"indexed\":false,\"internalType\":\"structPacketTypes.Packet\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"PacketReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"ports\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"dataList\",\"type\":\"bytes[]\"}],\"indexed\":false,\"internalType\":\"structPacketTypes.Packet\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"PacketSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MULTISEND_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessManager\",\"outputs\":[{\"internalType\":\"contractIAccessManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"ports\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"dataList\",\"type\":\"bytes[]\"}],\"internalType\":\"structPacketTypes.Packet\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofAcked\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"name\":\"acknowledgePacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clientManager\",\"outputs\":[{\"internalType\":\"contractIClientManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"ports\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"dataList\",\"type\":\"bytes[]\"}],\"internalType\":\"structPacketTypes.Packet\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"executePacket\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"}],\"name\":\"getNextSequenceSend\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"clientManagerContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"routingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"accessManagerContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"receipts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"ports\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"dataList\",\"type\":\"bytes[]\"}],\"internalType\":\"structPacketTypes.Packet\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"name\":\"recvPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routing\",\"outputs\":[{\"internalType\":\"contractIRouting\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"ports\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"dataList\",\"type\":\"bytes[]\"}],\"internalType\":\"structPacketTypes.Packet\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"sendMultiPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"ports\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"dataList\",\"type\":\"bytes[]\"}],\"internalType\":\"structPacketTypes.Packet\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"sendPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sequences\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// MULTISENDROLE is a free data retrieval call binding the contract method 0xbc86b303.
//
// Solidity: function MULTISEND_ROLE() view returns(bytes32)
func (_Contract *ContractCaller) MULTISENDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "MULTISEND_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MULTISENDROLE is a free data retrieval call binding the contract method 0xbc86b303.
//
// Solidity: function MULTISEND_ROLE() view returns(bytes32)
func (_Contract *ContractSession) MULTISENDROLE() ([32]byte, error) {
	return _Contract.Contract.MULTISENDROLE(&_Contract.CallOpts)
}

// MULTISENDROLE is a free data retrieval call binding the contract method 0xbc86b303.
//
// Solidity: function MULTISEND_ROLE() view returns(bytes32)
func (_Contract *ContractCallerSession) MULTISENDROLE() ([32]byte, error) {
	return _Contract.Contract.MULTISENDROLE(&_Contract.CallOpts)
}

// AccessManager is a free data retrieval call binding the contract method 0xfdcb6068.
//
// Solidity: function accessManager() view returns(address)
func (_Contract *ContractCaller) AccessManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "accessManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccessManager is a free data retrieval call binding the contract method 0xfdcb6068.
//
// Solidity: function accessManager() view returns(address)
func (_Contract *ContractSession) AccessManager() (common.Address, error) {
	return _Contract.Contract.AccessManager(&_Contract.CallOpts)
}

// AccessManager is a free data retrieval call binding the contract method 0xfdcb6068.
//
// Solidity: function accessManager() view returns(address)
func (_Contract *ContractCallerSession) AccessManager() (common.Address, error) {
	return _Contract.Contract.AccessManager(&_Contract.CallOpts)
}

// ClientManager is a free data retrieval call binding the contract method 0x79e8be1d.
//
// Solidity: function clientManager() view returns(address)
func (_Contract *ContractCaller) ClientManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "clientManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClientManager is a free data retrieval call binding the contract method 0x79e8be1d.
//
// Solidity: function clientManager() view returns(address)
func (_Contract *ContractSession) ClientManager() (common.Address, error) {
	return _Contract.Contract.ClientManager(&_Contract.CallOpts)
}

// ClientManager is a free data retrieval call binding the contract method 0x79e8be1d.
//
// Solidity: function clientManager() view returns(address)
func (_Contract *ContractCallerSession) ClientManager() (common.Address, error) {
	return _Contract.Contract.ClientManager(&_Contract.CallOpts)
}

// Commitments is a free data retrieval call binding the contract method 0x7912b8e6.
//
// Solidity: function commitments(bytes ) view returns(bytes32)
func (_Contract *ContractCaller) Commitments(opts *bind.CallOpts, arg0 []byte) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "commitments", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0x7912b8e6.
//
// Solidity: function commitments(bytes ) view returns(bytes32)
func (_Contract *ContractSession) Commitments(arg0 []byte) ([32]byte, error) {
	return _Contract.Contract.Commitments(&_Contract.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0x7912b8e6.
//
// Solidity: function commitments(bytes ) view returns(bytes32)
func (_Contract *ContractCallerSession) Commitments(arg0 []byte) ([32]byte, error) {
	return _Contract.Contract.Commitments(&_Contract.CallOpts, arg0)
}

// GetNextSequenceSend is a free data retrieval call binding the contract method 0x582418b6.
//
// Solidity: function getNextSequenceSend(string sourceChain, string destChain) view returns(uint64)
func (_Contract *ContractCaller) GetNextSequenceSend(opts *bind.CallOpts, sourceChain string, destChain string) (uint64, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getNextSequenceSend", sourceChain, destChain)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNextSequenceSend is a free data retrieval call binding the contract method 0x582418b6.
//
// Solidity: function getNextSequenceSend(string sourceChain, string destChain) view returns(uint64)
func (_Contract *ContractSession) GetNextSequenceSend(sourceChain string, destChain string) (uint64, error) {
	return _Contract.Contract.GetNextSequenceSend(&_Contract.CallOpts, sourceChain, destChain)
}

// GetNextSequenceSend is a free data retrieval call binding the contract method 0x582418b6.
//
// Solidity: function getNextSequenceSend(string sourceChain, string destChain) view returns(uint64)
func (_Contract *ContractCallerSession) GetNextSequenceSend(sourceChain string, destChain string) (uint64, error) {
	return _Contract.Contract.GetNextSequenceSend(&_Contract.CallOpts, sourceChain, destChain)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Receipts is a free data retrieval call binding the contract method 0xa6992b83.
//
// Solidity: function receipts(bytes ) view returns(bool)
func (_Contract *ContractCaller) Receipts(opts *bind.CallOpts, arg0 []byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "receipts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Receipts is a free data retrieval call binding the contract method 0xa6992b83.
//
// Solidity: function receipts(bytes ) view returns(bool)
func (_Contract *ContractSession) Receipts(arg0 []byte) (bool, error) {
	return _Contract.Contract.Receipts(&_Contract.CallOpts, arg0)
}

// Receipts is a free data retrieval call binding the contract method 0xa6992b83.
//
// Solidity: function receipts(bytes ) view returns(bool)
func (_Contract *ContractCallerSession) Receipts(arg0 []byte) (bool, error) {
	return _Contract.Contract.Receipts(&_Contract.CallOpts, arg0)
}

// Routing is a free data retrieval call binding the contract method 0x1b77f489.
//
// Solidity: function routing() view returns(address)
func (_Contract *ContractCaller) Routing(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "routing")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Routing is a free data retrieval call binding the contract method 0x1b77f489.
//
// Solidity: function routing() view returns(address)
func (_Contract *ContractSession) Routing() (common.Address, error) {
	return _Contract.Contract.Routing(&_Contract.CallOpts)
}

// Routing is a free data retrieval call binding the contract method 0x1b77f489.
//
// Solidity: function routing() view returns(address)
func (_Contract *ContractCallerSession) Routing() (common.Address, error) {
	return _Contract.Contract.Routing(&_Contract.CallOpts)
}

// Sequences is a free data retrieval call binding the contract method 0xeeebb020.
//
// Solidity: function sequences(bytes ) view returns(uint64)
func (_Contract *ContractCaller) Sequences(opts *bind.CallOpts, arg0 []byte) (uint64, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "sequences", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Sequences is a free data retrieval call binding the contract method 0xeeebb020.
//
// Solidity: function sequences(bytes ) view returns(uint64)
func (_Contract *ContractSession) Sequences(arg0 []byte) (uint64, error) {
	return _Contract.Contract.Sequences(&_Contract.CallOpts, arg0)
}

// Sequences is a free data retrieval call binding the contract method 0xeeebb020.
//
// Solidity: function sequences(bytes ) view returns(uint64)
func (_Contract *ContractCallerSession) Sequences(arg0 []byte) (uint64, error) {
	return _Contract.Contract.Sequences(&_Contract.CallOpts, arg0)
}

// AcknowledgePacket is a paid mutator transaction binding the contract method 0xc5404e07.
//
// Solidity: function acknowledgePacket((uint64,string,string,string,string[],bytes[]) packet, bytes acknowledgement, bytes proofAcked, (uint64,uint64) height) returns()
func (_Contract *ContractTransactor) AcknowledgePacket(opts *bind.TransactOpts, packet PacketTypesPacket, acknowledgement []byte, proofAcked []byte, height HeightData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "acknowledgePacket", packet, acknowledgement, proofAcked, height)
}

// AcknowledgePacket is a paid mutator transaction binding the contract method 0xc5404e07.
//
// Solidity: function acknowledgePacket((uint64,string,string,string,string[],bytes[]) packet, bytes acknowledgement, bytes proofAcked, (uint64,uint64) height) returns()
func (_Contract *ContractSession) AcknowledgePacket(packet PacketTypesPacket, acknowledgement []byte, proofAcked []byte, height HeightData) (*types.Transaction, error) {
	return _Contract.Contract.AcknowledgePacket(&_Contract.TransactOpts, packet, acknowledgement, proofAcked, height)
}

// AcknowledgePacket is a paid mutator transaction binding the contract method 0xc5404e07.
//
// Solidity: function acknowledgePacket((uint64,string,string,string,string[],bytes[]) packet, bytes acknowledgement, bytes proofAcked, (uint64,uint64) height) returns()
func (_Contract *ContractTransactorSession) AcknowledgePacket(packet PacketTypesPacket, acknowledgement []byte, proofAcked []byte, height HeightData) (*types.Transaction, error) {
	return _Contract.Contract.AcknowledgePacket(&_Contract.TransactOpts, packet, acknowledgement, proofAcked, height)
}

// ExecutePacket is a paid mutator transaction binding the contract method 0xf2c77bf3.
//
// Solidity: function executePacket((uint64,string,string,string,string[],bytes[]) packet) returns(bytes[])
func (_Contract *ContractTransactor) ExecutePacket(opts *bind.TransactOpts, packet PacketTypesPacket) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "executePacket", packet)
}

// ExecutePacket is a paid mutator transaction binding the contract method 0xf2c77bf3.
//
// Solidity: function executePacket((uint64,string,string,string,string[],bytes[]) packet) returns(bytes[])
func (_Contract *ContractSession) ExecutePacket(packet PacketTypesPacket) (*types.Transaction, error) {
	return _Contract.Contract.ExecutePacket(&_Contract.TransactOpts, packet)
}

// ExecutePacket is a paid mutator transaction binding the contract method 0xf2c77bf3.
//
// Solidity: function executePacket((uint64,string,string,string,string[],bytes[]) packet) returns(bytes[])
func (_Contract *ContractTransactorSession) ExecutePacket(packet PacketTypesPacket) (*types.Transaction, error) {
	return _Contract.Contract.ExecutePacket(&_Contract.TransactOpts, packet)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address clientManagerContract, address routingContract, address accessManagerContract) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, clientManagerContract common.Address, routingContract common.Address, accessManagerContract common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", clientManagerContract, routingContract, accessManagerContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address clientManagerContract, address routingContract, address accessManagerContract) returns()
func (_Contract *ContractSession) Initialize(clientManagerContract common.Address, routingContract common.Address, accessManagerContract common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, clientManagerContract, routingContract, accessManagerContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address clientManagerContract, address routingContract, address accessManagerContract) returns()
func (_Contract *ContractTransactorSession) Initialize(clientManagerContract common.Address, routingContract common.Address, accessManagerContract common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, clientManagerContract, routingContract, accessManagerContract)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x55ba244d.
//
// Solidity: function recvPacket((uint64,string,string,string,string[],bytes[]) packet, bytes proof, (uint64,uint64) height) returns()
func (_Contract *ContractTransactor) RecvPacket(opts *bind.TransactOpts, packet PacketTypesPacket, proof []byte, height HeightData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "recvPacket", packet, proof, height)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x55ba244d.
//
// Solidity: function recvPacket((uint64,string,string,string,string[],bytes[]) packet, bytes proof, (uint64,uint64) height) returns()
func (_Contract *ContractSession) RecvPacket(packet PacketTypesPacket, proof []byte, height HeightData) (*types.Transaction, error) {
	return _Contract.Contract.RecvPacket(&_Contract.TransactOpts, packet, proof, height)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x55ba244d.
//
// Solidity: function recvPacket((uint64,string,string,string,string[],bytes[]) packet, bytes proof, (uint64,uint64) height) returns()
func (_Contract *ContractTransactorSession) RecvPacket(packet PacketTypesPacket, proof []byte, height HeightData) (*types.Transaction, error) {
	return _Contract.Contract.RecvPacket(&_Contract.TransactOpts, packet, proof, height)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// SendMultiPacket is a paid mutator transaction binding the contract method 0x0b98ae77.
//
// Solidity: function sendMultiPacket((uint64,string,string,string,string[],bytes[]) packet) returns()
func (_Contract *ContractTransactor) SendMultiPacket(opts *bind.TransactOpts, packet PacketTypesPacket) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sendMultiPacket", packet)
}

// SendMultiPacket is a paid mutator transaction binding the contract method 0x0b98ae77.
//
// Solidity: function sendMultiPacket((uint64,string,string,string,string[],bytes[]) packet) returns()
func (_Contract *ContractSession) SendMultiPacket(packet PacketTypesPacket) (*types.Transaction, error) {
	return _Contract.Contract.SendMultiPacket(&_Contract.TransactOpts, packet)
}

// SendMultiPacket is a paid mutator transaction binding the contract method 0x0b98ae77.
//
// Solidity: function sendMultiPacket((uint64,string,string,string,string[],bytes[]) packet) returns()
func (_Contract *ContractTransactorSession) SendMultiPacket(packet PacketTypesPacket) (*types.Transaction, error) {
	return _Contract.Contract.SendMultiPacket(&_Contract.TransactOpts, packet)
}

// SendPacket is a paid mutator transaction binding the contract method 0xfca8e45e.
//
// Solidity: function sendPacket((uint64,string,string,string,string[],bytes[]) packet) returns()
func (_Contract *ContractTransactor) SendPacket(opts *bind.TransactOpts, packet PacketTypesPacket) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sendPacket", packet)
}

// SendPacket is a paid mutator transaction binding the contract method 0xfca8e45e.
//
// Solidity: function sendPacket((uint64,string,string,string,string[],bytes[]) packet) returns()
func (_Contract *ContractSession) SendPacket(packet PacketTypesPacket) (*types.Transaction, error) {
	return _Contract.Contract.SendPacket(&_Contract.TransactOpts, packet)
}

// SendPacket is a paid mutator transaction binding the contract method 0xfca8e45e.
//
// Solidity: function sendPacket((uint64,string,string,string,string[],bytes[]) packet) returns()
func (_Contract *ContractTransactorSession) SendPacket(packet PacketTypesPacket) (*types.Transaction, error) {
	return _Contract.Contract.SendPacket(&_Contract.TransactOpts, packet)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// ContractAckPacketIterator is returned from FilterAckPacket and is used to iterate over the raw logs and unpacked data for AckPacket events raised by the Contract contract.
type ContractAckPacketIterator struct {
	Event *ContractAckPacket // Event containing the contract specifics and raw log

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
func (it *ContractAckPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAckPacket)
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
		it.Event = new(ContractAckPacket)
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
func (it *ContractAckPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAckPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAckPacket represents a AckPacket event raised by the Contract contract.
type ContractAckPacket struct {
	Packet PacketTypesPacket
	Ack    []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAckPacket is a free log retrieval operation binding the contract event 0xd9ee0a534aed479e69eedde038db68d32784e2f38ac5ca0a54fc5c9d64dfc46a.
//
// Solidity: event AckPacket((uint64,string,string,string,string[],bytes[]) packet, bytes ack)
func (_Contract *ContractFilterer) FilterAckPacket(opts *bind.FilterOpts) (*ContractAckPacketIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AckPacket")
	if err != nil {
		return nil, err
	}
	return &ContractAckPacketIterator{contract: _Contract.contract, event: "AckPacket", logs: logs, sub: sub}, nil
}

// WatchAckPacket is a free log subscription operation binding the contract event 0xd9ee0a534aed479e69eedde038db68d32784e2f38ac5ca0a54fc5c9d64dfc46a.
//
// Solidity: event AckPacket((uint64,string,string,string,string[],bytes[]) packet, bytes ack)
func (_Contract *ContractFilterer) WatchAckPacket(opts *bind.WatchOpts, sink chan<- *ContractAckPacket) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AckPacket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAckPacket)
				if err := _Contract.contract.UnpackLog(event, "AckPacket", log); err != nil {
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

// ParseAckPacket is a log parse operation binding the contract event 0xd9ee0a534aed479e69eedde038db68d32784e2f38ac5ca0a54fc5c9d64dfc46a.
//
// Solidity: event AckPacket((uint64,string,string,string,string[],bytes[]) packet, bytes ack)
func (_Contract *ContractFilterer) ParseAckPacket(log types.Log) (*ContractAckPacket, error) {
	event := new(ContractAckPacket)
	if err := _Contract.contract.UnpackLog(event, "AckPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAckWrittenIterator is returned from FilterAckWritten and is used to iterate over the raw logs and unpacked data for AckWritten events raised by the Contract contract.
type ContractAckWrittenIterator struct {
	Event *ContractAckWritten // Event containing the contract specifics and raw log

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
func (it *ContractAckWrittenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAckWritten)
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
		it.Event = new(ContractAckWritten)
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
func (it *ContractAckWrittenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAckWrittenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAckWritten represents a AckWritten event raised by the Contract contract.
type ContractAckWritten struct {
	Packet PacketTypesPacket
	Ack    []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAckWritten is a free log retrieval operation binding the contract event 0x9d1208da43d68810a4fb4015fd8115ce2efcfbe4143da9469995772e0c457c5f.
//
// Solidity: event AckWritten((uint64,string,string,string,string[],bytes[]) packet, bytes ack)
func (_Contract *ContractFilterer) FilterAckWritten(opts *bind.FilterOpts) (*ContractAckWrittenIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AckWritten")
	if err != nil {
		return nil, err
	}
	return &ContractAckWrittenIterator{contract: _Contract.contract, event: "AckWritten", logs: logs, sub: sub}, nil
}

// WatchAckWritten is a free log subscription operation binding the contract event 0x9d1208da43d68810a4fb4015fd8115ce2efcfbe4143da9469995772e0c457c5f.
//
// Solidity: event AckWritten((uint64,string,string,string,string[],bytes[]) packet, bytes ack)
func (_Contract *ContractFilterer) WatchAckWritten(opts *bind.WatchOpts, sink chan<- *ContractAckWritten) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AckWritten")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAckWritten)
				if err := _Contract.contract.UnpackLog(event, "AckWritten", log); err != nil {
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

// ParseAckWritten is a log parse operation binding the contract event 0x9d1208da43d68810a4fb4015fd8115ce2efcfbe4143da9469995772e0c457c5f.
//
// Solidity: event AckWritten((uint64,string,string,string,string[],bytes[]) packet, bytes ack)
func (_Contract *ContractFilterer) ParseAckWritten(log types.Log) (*ContractAckWritten, error) {
	event := new(ContractAckWritten)
	if err := _Contract.contract.UnpackLog(event, "AckWritten", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPacketReceivedIterator is returned from FilterPacketReceived and is used to iterate over the raw logs and unpacked data for PacketReceived events raised by the Contract contract.
type ContractPacketReceivedIterator struct {
	Event *ContractPacketReceived // Event containing the contract specifics and raw log

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
func (it *ContractPacketReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPacketReceived)
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
		it.Event = new(ContractPacketReceived)
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
func (it *ContractPacketReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPacketReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPacketReceived represents a PacketReceived event raised by the Contract contract.
type ContractPacketReceived struct {
	Packet PacketTypesPacket
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPacketReceived is a free log retrieval operation binding the contract event 0x36b6683568ac4e60d9b9a2bbefc6514aff71cc21c909409f810e4fc0c5961da8.
//
// Solidity: event PacketReceived((uint64,string,string,string,string[],bytes[]) packet)
func (_Contract *ContractFilterer) FilterPacketReceived(opts *bind.FilterOpts) (*ContractPacketReceivedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PacketReceived")
	if err != nil {
		return nil, err
	}
	return &ContractPacketReceivedIterator{contract: _Contract.contract, event: "PacketReceived", logs: logs, sub: sub}, nil
}

// WatchPacketReceived is a free log subscription operation binding the contract event 0x36b6683568ac4e60d9b9a2bbefc6514aff71cc21c909409f810e4fc0c5961da8.
//
// Solidity: event PacketReceived((uint64,string,string,string,string[],bytes[]) packet)
func (_Contract *ContractFilterer) WatchPacketReceived(opts *bind.WatchOpts, sink chan<- *ContractPacketReceived) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PacketReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPacketReceived)
				if err := _Contract.contract.UnpackLog(event, "PacketReceived", log); err != nil {
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

// ParsePacketReceived is a log parse operation binding the contract event 0x36b6683568ac4e60d9b9a2bbefc6514aff71cc21c909409f810e4fc0c5961da8.
//
// Solidity: event PacketReceived((uint64,string,string,string,string[],bytes[]) packet)
func (_Contract *ContractFilterer) ParsePacketReceived(log types.Log) (*ContractPacketReceived, error) {
	event := new(ContractPacketReceived)
	if err := _Contract.contract.UnpackLog(event, "PacketReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPacketSentIterator is returned from FilterPacketSent and is used to iterate over the raw logs and unpacked data for PacketSent events raised by the Contract contract.
type ContractPacketSentIterator struct {
	Event *ContractPacketSent // Event containing the contract specifics and raw log

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
func (it *ContractPacketSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPacketSent)
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
		it.Event = new(ContractPacketSent)
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
func (it *ContractPacketSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPacketSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPacketSent represents a PacketSent event raised by the Contract contract.
type ContractPacketSent struct {
	Packet PacketTypesPacket
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPacketSent is a free log retrieval operation binding the contract event 0x0de2903114df93fdb4d35577325a50b43b67f8cb0f7f8edcbe2d36dd0799cebf.
//
// Solidity: event PacketSent((uint64,string,string,string,string[],bytes[]) packet)
func (_Contract *ContractFilterer) FilterPacketSent(opts *bind.FilterOpts) (*ContractPacketSentIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PacketSent")
	if err != nil {
		return nil, err
	}
	return &ContractPacketSentIterator{contract: _Contract.contract, event: "PacketSent", logs: logs, sub: sub}, nil
}

// WatchPacketSent is a free log subscription operation binding the contract event 0x0de2903114df93fdb4d35577325a50b43b67f8cb0f7f8edcbe2d36dd0799cebf.
//
// Solidity: event PacketSent((uint64,string,string,string,string[],bytes[]) packet)
func (_Contract *ContractFilterer) WatchPacketSent(opts *bind.WatchOpts, sink chan<- *ContractPacketSent) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PacketSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPacketSent)
				if err := _Contract.contract.UnpackLog(event, "PacketSent", log); err != nil {
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

// ParsePacketSent is a log parse operation binding the contract event 0x0de2903114df93fdb4d35577325a50b43b67f8cb0f7f8edcbe2d36dd0799cebf.
//
// Solidity: event PacketSent((uint64,string,string,string,string[],bytes[]) packet)
func (_Contract *ContractFilterer) ParsePacketSent(log types.Log) (*ContractPacketSent, error) {
	event := new(ContractPacketSent)
	if err := _Contract.contract.UnpackLog(event, "PacketSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
