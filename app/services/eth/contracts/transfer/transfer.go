// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transfer

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

// PacketTypesResult is an auto generated low-level Go binding around an user-defined struct.
type PacketTypesResult struct {
	Result  []byte
	Message string
}

// TransferDataTypesBaseTransferData is an auto generated low-level Go binding around an user-defined struct.
type TransferDataTypesBaseTransferData struct {
	Receiver   string
	DestChain  string
	RelayChain string
}

// TransferDataTypesBaseTransferDataMulti is an auto generated low-level Go binding around an user-defined struct.
type TransferDataTypesBaseTransferDataMulti struct {
	Sender    common.Address
	Receiver  string
	DestChain string
}

// TransferDataTypesERC20TransferData is an auto generated low-level Go binding around an user-defined struct.
type TransferDataTypesERC20TransferData struct {
	TokenAddress common.Address
	Receiver     string
	Amount       *big.Int
	DestChain    string
	RelayChain   string
}

// TransferDataTypesERC20TransferDataMulti is an auto generated low-level Go binding around an user-defined struct.
type TransferDataTypesERC20TransferDataMulti struct {
	TokenAddress common.Address
	Sender       common.Address
	Receiver     string
	Amount       *big.Int
	DestChain    string
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BIND_TOKEN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTISEND_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessManager\",\"outputs\":[{\"internalType\":\"contractIAccessManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddres\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"oriToken\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"oriChain\",\"type\":\"string\"}],\"name\":\"bindToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"bindingTraces\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bindings\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"oriChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"oriToken\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"bound\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"boundTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clientManager\",\"outputs\":[{\"internalType\":\"contractIClientManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"packetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"clientManagerContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"accessManagerContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"onAcknowledgementPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onRecvPacket\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"internalType\":\"structPacketTypes.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"outTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"packet\",\"outputs\":[{\"internalType\":\"contractIPacket\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"}],\"internalType\":\"structTransferDataTypes.BaseTransferData\",\"name\":\"transferData\",\"type\":\"tuple\"}],\"name\":\"sendTransferBase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"}],\"internalType\":\"structTransferDataTypes.ERC20TransferData\",\"name\":\"transferData\",\"type\":\"tuple\"}],\"name\":\"sendTransferERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"}],\"internalType\":\"structTransferDataTypes.BaseTransferDataMulti\",\"name\":\"transferData\",\"type\":\"tuple\"}],\"name\":\"transferBase\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"}],\"internalType\":\"structTransferDataTypes.ERC20TransferDataMulti\",\"name\":\"transferData\",\"type\":\"tuple\"}],\"name\":\"transferERC20\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// BINDTOKENROLE is a free data retrieval call binding the contract method 0x9a3e594f.
//
// Solidity: function BIND_TOKEN_ROLE() view returns(bytes32)
func (_Contracts *ContractsCaller) BINDTOKENROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "BIND_TOKEN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BINDTOKENROLE is a free data retrieval call binding the contract method 0x9a3e594f.
//
// Solidity: function BIND_TOKEN_ROLE() view returns(bytes32)
func (_Contracts *ContractsSession) BINDTOKENROLE() ([32]byte, error) {
	return _Contracts.Contract.BINDTOKENROLE(&_Contracts.CallOpts)
}

// BINDTOKENROLE is a free data retrieval call binding the contract method 0x9a3e594f.
//
// Solidity: function BIND_TOKEN_ROLE() view returns(bytes32)
func (_Contracts *ContractsCallerSession) BINDTOKENROLE() ([32]byte, error) {
	return _Contracts.Contract.BINDTOKENROLE(&_Contracts.CallOpts)
}

// MULTISENDROLE is a free data retrieval call binding the contract method 0xbc86b303.
//
// Solidity: function MULTISEND_ROLE() view returns(bytes32)
func (_Contracts *ContractsCaller) MULTISENDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "MULTISEND_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MULTISENDROLE is a free data retrieval call binding the contract method 0xbc86b303.
//
// Solidity: function MULTISEND_ROLE() view returns(bytes32)
func (_Contracts *ContractsSession) MULTISENDROLE() ([32]byte, error) {
	return _Contracts.Contract.MULTISENDROLE(&_Contracts.CallOpts)
}

// MULTISENDROLE is a free data retrieval call binding the contract method 0xbc86b303.
//
// Solidity: function MULTISEND_ROLE() view returns(bytes32)
func (_Contracts *ContractsCallerSession) MULTISENDROLE() ([32]byte, error) {
	return _Contracts.Contract.MULTISENDROLE(&_Contracts.CallOpts)
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

// BindingTraces is a free data retrieval call binding the contract method 0xb319271d.
//
// Solidity: function bindingTraces(string ) view returns(address)
func (_Contracts *ContractsCaller) BindingTraces(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "bindingTraces", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BindingTraces is a free data retrieval call binding the contract method 0xb319271d.
//
// Solidity: function bindingTraces(string ) view returns(address)
func (_Contracts *ContractsSession) BindingTraces(arg0 string) (common.Address, error) {
	return _Contracts.Contract.BindingTraces(&_Contracts.CallOpts, arg0)
}

// BindingTraces is a free data retrieval call binding the contract method 0xb319271d.
//
// Solidity: function bindingTraces(string ) view returns(address)
func (_Contracts *ContractsCallerSession) BindingTraces(arg0 string) (common.Address, error) {
	return _Contracts.Contract.BindingTraces(&_Contracts.CallOpts, arg0)
}

// Bindings is a free data retrieval call binding the contract method 0x78022813.
//
// Solidity: function bindings(address ) view returns(string oriChain, string oriToken, uint256 amount, bool bound)
func (_Contracts *ContractsCaller) Bindings(opts *bind.CallOpts, arg0 common.Address) (struct {
	OriChain string
	OriToken string
	Amount   *big.Int
	Bound    bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "bindings", arg0)

	outstruct := new(struct {
		OriChain string
		OriToken string
		Amount   *big.Int
		Bound    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OriChain = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.OriToken = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Bound = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Bindings is a free data retrieval call binding the contract method 0x78022813.
//
// Solidity: function bindings(address ) view returns(string oriChain, string oriToken, uint256 amount, bool bound)
func (_Contracts *ContractsSession) Bindings(arg0 common.Address) (struct {
	OriChain string
	OriToken string
	Amount   *big.Int
	Bound    bool
}, error) {
	return _Contracts.Contract.Bindings(&_Contracts.CallOpts, arg0)
}

// Bindings is a free data retrieval call binding the contract method 0x78022813.
//
// Solidity: function bindings(address ) view returns(string oriChain, string oriToken, uint256 amount, bool bound)
func (_Contracts *ContractsCallerSession) Bindings(arg0 common.Address) (struct {
	OriChain string
	OriToken string
	Amount   *big.Int
	Bound    bool
}, error) {
	return _Contracts.Contract.Bindings(&_Contracts.CallOpts, arg0)
}

// BoundTokens is a free data retrieval call binding the contract method 0xe8e8da7c.
//
// Solidity: function boundTokens(uint256 ) view returns(address)
func (_Contracts *ContractsCaller) BoundTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "boundTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BoundTokens is a free data retrieval call binding the contract method 0xe8e8da7c.
//
// Solidity: function boundTokens(uint256 ) view returns(address)
func (_Contracts *ContractsSession) BoundTokens(arg0 *big.Int) (common.Address, error) {
	return _Contracts.Contract.BoundTokens(&_Contracts.CallOpts, arg0)
}

// BoundTokens is a free data retrieval call binding the contract method 0xe8e8da7c.
//
// Solidity: function boundTokens(uint256 ) view returns(address)
func (_Contracts *ContractsCallerSession) BoundTokens(arg0 *big.Int) (common.Address, error) {
	return _Contracts.Contract.BoundTokens(&_Contracts.CallOpts, arg0)
}

// ClientManager is a free data retrieval call binding the contract method 0x79e8be1d.
//
// Solidity: function clientManager() view returns(address)
func (_Contracts *ContractsCaller) ClientManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "clientManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClientManager is a free data retrieval call binding the contract method 0x79e8be1d.
//
// Solidity: function clientManager() view returns(address)
func (_Contracts *ContractsSession) ClientManager() (common.Address, error) {
	return _Contracts.Contract.ClientManager(&_Contracts.CallOpts)
}

// ClientManager is a free data retrieval call binding the contract method 0x79e8be1d.
//
// Solidity: function clientManager() view returns(address)
func (_Contracts *ContractsCallerSession) ClientManager() (common.Address, error) {
	return _Contracts.Contract.ClientManager(&_Contracts.CallOpts)
}

// OutTokens is a free data retrieval call binding the contract method 0xae22ea80.
//
// Solidity: function outTokens(address , string ) view returns(uint256)
func (_Contracts *ContractsCaller) OutTokens(opts *bind.CallOpts, arg0 common.Address, arg1 string) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "outTokens", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutTokens is a free data retrieval call binding the contract method 0xae22ea80.
//
// Solidity: function outTokens(address , string ) view returns(uint256)
func (_Contracts *ContractsSession) OutTokens(arg0 common.Address, arg1 string) (*big.Int, error) {
	return _Contracts.Contract.OutTokens(&_Contracts.CallOpts, arg0, arg1)
}

// OutTokens is a free data retrieval call binding the contract method 0xae22ea80.
//
// Solidity: function outTokens(address , string ) view returns(uint256)
func (_Contracts *ContractsCallerSession) OutTokens(arg0 common.Address, arg1 string) (*big.Int, error) {
	return _Contracts.Contract.OutTokens(&_Contracts.CallOpts, arg0, arg1)
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

// Packet is a free data retrieval call binding the contract method 0x1c3083ef.
//
// Solidity: function packet() view returns(address)
func (_Contracts *ContractsCaller) Packet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "packet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Packet is a free data retrieval call binding the contract method 0x1c3083ef.
//
// Solidity: function packet() view returns(address)
func (_Contracts *ContractsSession) Packet() (common.Address, error) {
	return _Contracts.Contract.Packet(&_Contracts.CallOpts)
}

// Packet is a free data retrieval call binding the contract method 0x1c3083ef.
//
// Solidity: function packet() view returns(address)
func (_Contracts *ContractsCallerSession) Packet() (common.Address, error) {
	return _Contracts.Contract.Packet(&_Contracts.CallOpts)
}

// BindToken is a paid mutator transaction binding the contract method 0x8155f9a2.
//
// Solidity: function bindToken(address tokenAddres, string oriToken, string oriChain) returns()
func (_Contracts *ContractsTransactor) BindToken(opts *bind.TransactOpts, tokenAddres common.Address, oriToken string, oriChain string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "bindToken", tokenAddres, oriToken, oriChain)
}

// BindToken is a paid mutator transaction binding the contract method 0x8155f9a2.
//
// Solidity: function bindToken(address tokenAddres, string oriToken, string oriChain) returns()
func (_Contracts *ContractsSession) BindToken(tokenAddres common.Address, oriToken string, oriChain string) (*types.Transaction, error) {
	return _Contracts.Contract.BindToken(&_Contracts.TransactOpts, tokenAddres, oriToken, oriChain)
}

// BindToken is a paid mutator transaction binding the contract method 0x8155f9a2.
//
// Solidity: function bindToken(address tokenAddres, string oriToken, string oriChain) returns()
func (_Contracts *ContractsTransactorSession) BindToken(tokenAddres common.Address, oriToken string, oriChain string) (*types.Transaction, error) {
	return _Contracts.Contract.BindToken(&_Contracts.TransactOpts, tokenAddres, oriToken, oriChain)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address packetContract, address clientManagerContract, address accessManagerContract) returns()
func (_Contracts *ContractsTransactor) Initialize(opts *bind.TransactOpts, packetContract common.Address, clientManagerContract common.Address, accessManagerContract common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "initialize", packetContract, clientManagerContract, accessManagerContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address packetContract, address clientManagerContract, address accessManagerContract) returns()
func (_Contracts *ContractsSession) Initialize(packetContract common.Address, clientManagerContract common.Address, accessManagerContract common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, packetContract, clientManagerContract, accessManagerContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address packetContract, address clientManagerContract, address accessManagerContract) returns()
func (_Contracts *ContractsTransactorSession) Initialize(packetContract common.Address, clientManagerContract common.Address, accessManagerContract common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, packetContract, clientManagerContract, accessManagerContract)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x6f43f8b0.
//
// Solidity: function onAcknowledgementPacket(bytes data, bytes result) returns()
func (_Contracts *ContractsTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, data []byte, result []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "onAcknowledgementPacket", data, result)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x6f43f8b0.
//
// Solidity: function onAcknowledgementPacket(bytes data, bytes result) returns()
func (_Contracts *ContractsSession) OnAcknowledgementPacket(data []byte, result []byte) (*types.Transaction, error) {
	return _Contracts.Contract.OnAcknowledgementPacket(&_Contracts.TransactOpts, data, result)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x6f43f8b0.
//
// Solidity: function onAcknowledgementPacket(bytes data, bytes result) returns()
func (_Contracts *ContractsTransactorSession) OnAcknowledgementPacket(data []byte, result []byte) (*types.Transaction, error) {
	return _Contracts.Contract.OnAcknowledgementPacket(&_Contracts.TransactOpts, data, result)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x2b3fa679.
//
// Solidity: function onRecvPacket(bytes data) returns((bytes,string))
func (_Contracts *ContractsTransactor) OnRecvPacket(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "onRecvPacket", data)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x2b3fa679.
//
// Solidity: function onRecvPacket(bytes data) returns((bytes,string))
func (_Contracts *ContractsSession) OnRecvPacket(data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.OnRecvPacket(&_Contracts.TransactOpts, data)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x2b3fa679.
//
// Solidity: function onRecvPacket(bytes data) returns((bytes,string))
func (_Contracts *ContractsTransactorSession) OnRecvPacket(data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.OnRecvPacket(&_Contracts.TransactOpts, data)
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

// SendTransferBase is a paid mutator transaction binding the contract method 0xf9bde822.
//
// Solidity: function sendTransferBase((string,string,string) transferData) payable returns()
func (_Contracts *ContractsTransactor) SendTransferBase(opts *bind.TransactOpts, transferData TransferDataTypesBaseTransferData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "sendTransferBase", transferData)
}

// SendTransferBase is a paid mutator transaction binding the contract method 0xf9bde822.
//
// Solidity: function sendTransferBase((string,string,string) transferData) payable returns()
func (_Contracts *ContractsSession) SendTransferBase(transferData TransferDataTypesBaseTransferData) (*types.Transaction, error) {
	return _Contracts.Contract.SendTransferBase(&_Contracts.TransactOpts, transferData)
}

// SendTransferBase is a paid mutator transaction binding the contract method 0xf9bde822.
//
// Solidity: function sendTransferBase((string,string,string) transferData) payable returns()
func (_Contracts *ContractsTransactorSession) SendTransferBase(transferData TransferDataTypesBaseTransferData) (*types.Transaction, error) {
	return _Contracts.Contract.SendTransferBase(&_Contracts.TransactOpts, transferData)
}

// SendTransferERC20 is a paid mutator transaction binding the contract method 0xe0c51f15.
//
// Solidity: function sendTransferERC20((address,string,uint256,string,string) transferData) returns()
func (_Contracts *ContractsTransactor) SendTransferERC20(opts *bind.TransactOpts, transferData TransferDataTypesERC20TransferData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "sendTransferERC20", transferData)
}

// SendTransferERC20 is a paid mutator transaction binding the contract method 0xe0c51f15.
//
// Solidity: function sendTransferERC20((address,string,uint256,string,string) transferData) returns()
func (_Contracts *ContractsSession) SendTransferERC20(transferData TransferDataTypesERC20TransferData) (*types.Transaction, error) {
	return _Contracts.Contract.SendTransferERC20(&_Contracts.TransactOpts, transferData)
}

// SendTransferERC20 is a paid mutator transaction binding the contract method 0xe0c51f15.
//
// Solidity: function sendTransferERC20((address,string,uint256,string,string) transferData) returns()
func (_Contracts *ContractsTransactorSession) SendTransferERC20(transferData TransferDataTypesERC20TransferData) (*types.Transaction, error) {
	return _Contracts.Contract.SendTransferERC20(&_Contracts.TransactOpts, transferData)
}

// TransferBase is a paid mutator transaction binding the contract method 0x1387da63.
//
// Solidity: function transferBase((address,string,string) transferData) payable returns(bytes)
func (_Contracts *ContractsTransactor) TransferBase(opts *bind.TransactOpts, transferData TransferDataTypesBaseTransferDataMulti) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferBase", transferData)
}

// TransferBase is a paid mutator transaction binding the contract method 0x1387da63.
//
// Solidity: function transferBase((address,string,string) transferData) payable returns(bytes)
func (_Contracts *ContractsSession) TransferBase(transferData TransferDataTypesBaseTransferDataMulti) (*types.Transaction, error) {
	return _Contracts.Contract.TransferBase(&_Contracts.TransactOpts, transferData)
}

// TransferBase is a paid mutator transaction binding the contract method 0x1387da63.
//
// Solidity: function transferBase((address,string,string) transferData) payable returns(bytes)
func (_Contracts *ContractsTransactorSession) TransferBase(transferData TransferDataTypesBaseTransferDataMulti) (*types.Transaction, error) {
	return _Contracts.Contract.TransferBase(&_Contracts.TransactOpts, transferData)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0xe568c141.
//
// Solidity: function transferERC20((address,address,string,uint256,string) transferData) returns(bytes)
func (_Contracts *ContractsTransactor) TransferERC20(opts *bind.TransactOpts, transferData TransferDataTypesERC20TransferDataMulti) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferERC20", transferData)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0xe568c141.
//
// Solidity: function transferERC20((address,address,string,uint256,string) transferData) returns(bytes)
func (_Contracts *ContractsSession) TransferERC20(transferData TransferDataTypesERC20TransferDataMulti) (*types.Transaction, error) {
	return _Contracts.Contract.TransferERC20(&_Contracts.TransactOpts, transferData)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0xe568c141.
//
// Solidity: function transferERC20((address,address,string,uint256,string) transferData) returns(bytes)
func (_Contracts *ContractsTransactorSession) TransferERC20(transferData TransferDataTypesERC20TransferDataMulti) (*types.Transaction, error) {
	return _Contracts.Contract.TransferERC20(&_Contracts.TransactOpts, transferData)
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
