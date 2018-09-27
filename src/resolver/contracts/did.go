// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package did

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// DidABI is the input ABI used to generate the binding from.
const DidABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"revokeAttribute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"authenticationChanged\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"authenticationType\",\"type\":\"bytes32\"},{\"name\":\"authentication\",\"type\":\"bytes32\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"addAuthentication\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"authenticationType\",\"type\":\"bytes32\"},{\"name\":\"authentication\",\"type\":\"bytes32\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"addAuthenticationSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"setAttributeSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"publicKeyChanged\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"publicKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnerSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"authentications\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"publicKeyType\",\"type\":\"bytes32\"},{\"name\":\"publicKey\",\"type\":\"bytes32\"}],\"name\":\"revokePublicKeySigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"authenticationType\",\"type\":\"bytes32\"},{\"name\":\"authentication\",\"type\":\"bytes32\"}],\"name\":\"revokeAuthentication\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"publicKeyType\",\"type\":\"bytes32\"},{\"name\":\"publicKey\",\"type\":\"bytes32\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"addPublicKeySigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"publicKeyType\",\"type\":\"bytes32\"},{\"name\":\"publicKey\",\"type\":\"bytes32\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"addPublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"publicKeyType\",\"type\":\"bytes32\"},{\"name\":\"publicKey\",\"type\":\"bytes32\"}],\"name\":\"revokePublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"setAttribute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"identityOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"attributeChanged\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"revokeAttributeSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"publicKeyType\",\"type\":\"bytes32\"},{\"name\":\"publicKey\",\"type\":\"bytes32\"}],\"name\":\"validPublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"authenticationType\",\"type\":\"bytes32\"},{\"name\":\"authentication\",\"type\":\"bytes32\"}],\"name\":\"revokeAuthenticationSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"authenticationType\",\"type\":\"bytes32\"},{\"name\":\"authentication\",\"type\":\"bytes32\"}],\"name\":\"validAuthentication\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"authenticationType\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"authentication\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"validTo\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"previousChange\",\"type\":\"uint256\"}],\"name\":\"DIDPublicKeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"authenticationType\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"authentication\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"validTo\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"previousChange\",\"type\":\"uint256\"}],\"name\":\"DIDAuthenticationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"validTo\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"previousChange\",\"type\":\"uint256\"}],\"name\":\"DIDAttributeChanged\",\"type\":\"event\"}]"

// Did is an auto generated Go binding around an Ethereum contract.
type Did struct {
	DidCaller     // Read-only binding to the contract
	DidTransactor // Write-only binding to the contract
	DidFilterer   // Log filterer for contract events
}

// DidCaller is an auto generated read-only Go binding around an Ethereum contract.
type DidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DidSession struct {
	Contract     *Did              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DidCallerSession struct {
	Contract *DidCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DidTransactorSession struct {
	Contract     *DidTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DidRaw is an auto generated low-level Go binding around an Ethereum contract.
type DidRaw struct {
	Contract *Did // Generic contract binding to access the raw methods on
}

// DidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DidCallerRaw struct {
	Contract *DidCaller // Generic read-only contract binding to access the raw methods on
}

// DidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DidTransactorRaw struct {
	Contract *DidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDid creates a new instance of Did, bound to a specific deployed contract.
func NewDid(address common.Address, backend bind.ContractBackend) (*Did, error) {
	contract, err := bindDid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Did{DidCaller: DidCaller{contract: contract}, DidTransactor: DidTransactor{contract: contract}, DidFilterer: DidFilterer{contract: contract}}, nil
}

// NewDidCaller creates a new read-only instance of Did, bound to a specific deployed contract.
func NewDidCaller(address common.Address, caller bind.ContractCaller) (*DidCaller, error) {
	contract, err := bindDid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DidCaller{contract: contract}, nil
}

// NewDidTransactor creates a new write-only instance of Did, bound to a specific deployed contract.
func NewDidTransactor(address common.Address, transactor bind.ContractTransactor) (*DidTransactor, error) {
	contract, err := bindDid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DidTransactor{contract: contract}, nil
}

// NewDidFilterer creates a new log filterer instance of Did, bound to a specific deployed contract.
func NewDidFilterer(address common.Address, filterer bind.ContractFilterer) (*DidFilterer, error) {
	contract, err := bindDid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DidFilterer{contract: contract}, nil
}

// bindDid binds a generic wrapper to an already deployed contract.
func bindDid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DidABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Did *DidRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Did.Contract.DidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Did *DidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Did.Contract.DidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Did *DidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Did.Contract.DidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Did *DidCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Did.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Did *DidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Did.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Did *DidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Did.Contract.contract.Transact(opts, method, params...)
}

// AttributeChanged is a free data retrieval call binding the contract method 0xd3c530d6.
//
// Solidity: function attributeChanged( address) constant returns(uint256)
func (_Did *DidCaller) AttributeChanged(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "attributeChanged", arg0)
	return *ret0, err
}

// AttributeChanged is a free data retrieval call binding the contract method 0xd3c530d6.
//
// Solidity: function attributeChanged( address) constant returns(uint256)
func (_Did *DidSession) AttributeChanged(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.AttributeChanged(&_Did.CallOpts, arg0)
}

// AttributeChanged is a free data retrieval call binding the contract method 0xd3c530d6.
//
// Solidity: function attributeChanged( address) constant returns(uint256)
func (_Did *DidCallerSession) AttributeChanged(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.AttributeChanged(&_Did.CallOpts, arg0)
}

// AuthenticationChanged is a free data retrieval call binding the contract method 0x048ed937.
//
// Solidity: function authenticationChanged( address) constant returns(uint256)
func (_Did *DidCaller) AuthenticationChanged(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "authenticationChanged", arg0)
	return *ret0, err
}

// AuthenticationChanged is a free data retrieval call binding the contract method 0x048ed937.
//
// Solidity: function authenticationChanged( address) constant returns(uint256)
func (_Did *DidSession) AuthenticationChanged(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.AuthenticationChanged(&_Did.CallOpts, arg0)
}

// AuthenticationChanged is a free data retrieval call binding the contract method 0x048ed937.
//
// Solidity: function authenticationChanged( address) constant returns(uint256)
func (_Did *DidCallerSession) AuthenticationChanged(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.AuthenticationChanged(&_Did.CallOpts, arg0)
}

// Authentications is a free data retrieval call binding the contract method 0x39fdc728.
//
// Solidity: function authentications( address,  bytes32,  bytes32) constant returns(uint256)
func (_Did *DidCaller) Authentications(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte, arg2 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "authentications", arg0, arg1, arg2)
	return *ret0, err
}

// Authentications is a free data retrieval call binding the contract method 0x39fdc728.
//
// Solidity: function authentications( address,  bytes32,  bytes32) constant returns(uint256)
func (_Did *DidSession) Authentications(arg0 common.Address, arg1 [32]byte, arg2 [32]byte) (*big.Int, error) {
	return _Did.Contract.Authentications(&_Did.CallOpts, arg0, arg1, arg2)
}

// Authentications is a free data retrieval call binding the contract method 0x39fdc728.
//
// Solidity: function authentications( address,  bytes32,  bytes32) constant returns(uint256)
func (_Did *DidCallerSession) Authentications(arg0 common.Address, arg1 [32]byte, arg2 [32]byte) (*big.Int, error) {
	return _Did.Contract.Authentications(&_Did.CallOpts, arg0, arg1, arg2)
}

// IdentityOwner is a free data retrieval call binding the contract method 0x8733d4e8.
//
// Solidity: function identityOwner(identity address) constant returns(address)
func (_Did *DidCaller) IdentityOwner(opts *bind.CallOpts, identity common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "identityOwner", identity)
	return *ret0, err
}

// IdentityOwner is a free data retrieval call binding the contract method 0x8733d4e8.
//
// Solidity: function identityOwner(identity address) constant returns(address)
func (_Did *DidSession) IdentityOwner(identity common.Address) (common.Address, error) {
	return _Did.Contract.IdentityOwner(&_Did.CallOpts, identity)
}

// IdentityOwner is a free data retrieval call binding the contract method 0x8733d4e8.
//
// Solidity: function identityOwner(identity address) constant returns(address)
func (_Did *DidCallerSession) IdentityOwner(identity common.Address) (common.Address, error) {
	return _Did.Contract.IdentityOwner(&_Did.CallOpts, identity)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce( address) constant returns(uint256)
func (_Did *DidCaller) Nonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "nonce", arg0)
	return *ret0, err
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce( address) constant returns(uint256)
func (_Did *DidSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.Nonce(&_Did.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce( address) constant returns(uint256)
func (_Did *DidCallerSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.Nonce(&_Did.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(address)
func (_Did *DidCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(address)
func (_Did *DidSession) Owners(arg0 common.Address) (common.Address, error) {
	return _Did.Contract.Owners(&_Did.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(address)
func (_Did *DidCallerSession) Owners(arg0 common.Address) (common.Address, error) {
	return _Did.Contract.Owners(&_Did.CallOpts, arg0)
}

// PublicKeyChanged is a free data retrieval call binding the contract method 0x1a2bcf4c.
//
// Solidity: function publicKeyChanged( address) constant returns(uint256)
func (_Did *DidCaller) PublicKeyChanged(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "publicKeyChanged", arg0)
	return *ret0, err
}

// PublicKeyChanged is a free data retrieval call binding the contract method 0x1a2bcf4c.
//
// Solidity: function publicKeyChanged( address) constant returns(uint256)
func (_Did *DidSession) PublicKeyChanged(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.PublicKeyChanged(&_Did.CallOpts, arg0)
}

// PublicKeyChanged is a free data retrieval call binding the contract method 0x1a2bcf4c.
//
// Solidity: function publicKeyChanged( address) constant returns(uint256)
func (_Did *DidCallerSession) PublicKeyChanged(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.PublicKeyChanged(&_Did.CallOpts, arg0)
}

// PublicKeys is a free data retrieval call binding the contract method 0x20cacc32.
//
// Solidity: function publicKeys( address,  bytes32,  bytes32) constant returns(uint256)
func (_Did *DidCaller) PublicKeys(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte, arg2 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "publicKeys", arg0, arg1, arg2)
	return *ret0, err
}

// PublicKeys is a free data retrieval call binding the contract method 0x20cacc32.
//
// Solidity: function publicKeys( address,  bytes32,  bytes32) constant returns(uint256)
func (_Did *DidSession) PublicKeys(arg0 common.Address, arg1 [32]byte, arg2 [32]byte) (*big.Int, error) {
	return _Did.Contract.PublicKeys(&_Did.CallOpts, arg0, arg1, arg2)
}

// PublicKeys is a free data retrieval call binding the contract method 0x20cacc32.
//
// Solidity: function publicKeys( address,  bytes32,  bytes32) constant returns(uint256)
func (_Did *DidCallerSession) PublicKeys(arg0 common.Address, arg1 [32]byte, arg2 [32]byte) (*big.Int, error) {
	return _Did.Contract.PublicKeys(&_Did.CallOpts, arg0, arg1, arg2)
}

// ValidAuthentication is a free data retrieval call binding the contract method 0xffb0254d.
//
// Solidity: function validAuthentication(identity address, authenticationType bytes32, authentication bytes32) constant returns(bool)
func (_Did *DidCaller) ValidAuthentication(opts *bind.CallOpts, identity common.Address, authenticationType [32]byte, authentication [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "validAuthentication", identity, authenticationType, authentication)
	return *ret0, err
}

// ValidAuthentication is a free data retrieval call binding the contract method 0xffb0254d.
//
// Solidity: function validAuthentication(identity address, authenticationType bytes32, authentication bytes32) constant returns(bool)
func (_Did *DidSession) ValidAuthentication(identity common.Address, authenticationType [32]byte, authentication [32]byte) (bool, error) {
	return _Did.Contract.ValidAuthentication(&_Did.CallOpts, identity, authenticationType, authentication)
}

// ValidAuthentication is a free data retrieval call binding the contract method 0xffb0254d.
//
// Solidity: function validAuthentication(identity address, authenticationType bytes32, authentication bytes32) constant returns(bool)
func (_Did *DidCallerSession) ValidAuthentication(identity common.Address, authenticationType [32]byte, authentication [32]byte) (bool, error) {
	return _Did.Contract.ValidAuthentication(&_Did.CallOpts, identity, authenticationType, authentication)
}

// ValidPublicKey is a free data retrieval call binding the contract method 0xe66bc6ba.
//
// Solidity: function validPublicKey(identity address, publicKeyType bytes32, publicKey bytes32) constant returns(bool)
func (_Did *DidCaller) ValidPublicKey(opts *bind.CallOpts, identity common.Address, publicKeyType [32]byte, publicKey [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "validPublicKey", identity, publicKeyType, publicKey)
	return *ret0, err
}

// ValidPublicKey is a free data retrieval call binding the contract method 0xe66bc6ba.
//
// Solidity: function validPublicKey(identity address, publicKeyType bytes32, publicKey bytes32) constant returns(bool)
func (_Did *DidSession) ValidPublicKey(identity common.Address, publicKeyType [32]byte, publicKey [32]byte) (bool, error) {
	return _Did.Contract.ValidPublicKey(&_Did.CallOpts, identity, publicKeyType, publicKey)
}

// ValidPublicKey is a free data retrieval call binding the contract method 0xe66bc6ba.
//
// Solidity: function validPublicKey(identity address, publicKeyType bytes32, publicKey bytes32) constant returns(bool)
func (_Did *DidCallerSession) ValidPublicKey(identity common.Address, publicKeyType [32]byte, publicKey [32]byte) (bool, error) {
	return _Did.Contract.ValidPublicKey(&_Did.CallOpts, identity, publicKeyType, publicKey)
}

// AddAuthentication is a paid mutator transaction binding the contract method 0x0824b3fa.
//
// Solidity: function addAuthentication(identity address, authenticationType bytes32, authentication bytes32, validity uint256) returns()
func (_Did *DidTransactor) AddAuthentication(opts *bind.TransactOpts, identity common.Address, authenticationType [32]byte, authentication [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "addAuthentication", identity, authenticationType, authentication, validity)
}

// AddAuthentication is a paid mutator transaction binding the contract method 0x0824b3fa.
//
// Solidity: function addAuthentication(identity address, authenticationType bytes32, authentication bytes32, validity uint256) returns()
func (_Did *DidSession) AddAuthentication(identity common.Address, authenticationType [32]byte, authentication [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddAuthentication(&_Did.TransactOpts, identity, authenticationType, authentication, validity)
}

// AddAuthentication is a paid mutator transaction binding the contract method 0x0824b3fa.
//
// Solidity: function addAuthentication(identity address, authenticationType bytes32, authentication bytes32, validity uint256) returns()
func (_Did *DidTransactorSession) AddAuthentication(identity common.Address, authenticationType [32]byte, authentication [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddAuthentication(&_Did.TransactOpts, identity, authenticationType, authentication, validity)
}

// AddAuthenticationSigned is a paid mutator transaction binding the contract method 0x0e110989.
//
// Solidity: function addAuthenticationSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, authenticationType bytes32, authentication bytes32, validity uint256) returns()
func (_Did *DidTransactor) AddAuthenticationSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, authenticationType [32]byte, authentication [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "addAuthenticationSigned", identity, sigV, sigR, sigS, authenticationType, authentication, validity)
}

// AddAuthenticationSigned is a paid mutator transaction binding the contract method 0x0e110989.
//
// Solidity: function addAuthenticationSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, authenticationType bytes32, authentication bytes32, validity uint256) returns()
func (_Did *DidSession) AddAuthenticationSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, authenticationType [32]byte, authentication [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddAuthenticationSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, authenticationType, authentication, validity)
}

// AddAuthenticationSigned is a paid mutator transaction binding the contract method 0x0e110989.
//
// Solidity: function addAuthenticationSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, authenticationType bytes32, authentication bytes32, validity uint256) returns()
func (_Did *DidTransactorSession) AddAuthenticationSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, authenticationType [32]byte, authentication [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddAuthenticationSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, authenticationType, authentication, validity)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0x6866a606.
//
// Solidity: function addPublicKey(identity address, publicKeyType bytes32, publicKey bytes32, validity uint256) returns()
func (_Did *DidTransactor) AddPublicKey(opts *bind.TransactOpts, identity common.Address, publicKeyType [32]byte, publicKey [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "addPublicKey", identity, publicKeyType, publicKey, validity)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0x6866a606.
//
// Solidity: function addPublicKey(identity address, publicKeyType bytes32, publicKey bytes32, validity uint256) returns()
func (_Did *DidSession) AddPublicKey(identity common.Address, publicKeyType [32]byte, publicKey [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddPublicKey(&_Did.TransactOpts, identity, publicKeyType, publicKey, validity)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0x6866a606.
//
// Solidity: function addPublicKey(identity address, publicKeyType bytes32, publicKey bytes32, validity uint256) returns()
func (_Did *DidTransactorSession) AddPublicKey(identity common.Address, publicKeyType [32]byte, publicKey [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddPublicKey(&_Did.TransactOpts, identity, publicKeyType, publicKey, validity)
}

// AddPublicKeySigned is a paid mutator transaction binding the contract method 0x456fcfeb.
//
// Solidity: function addPublicKeySigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, publicKeyType bytes32, publicKey bytes32, validity uint256) returns()
func (_Did *DidTransactor) AddPublicKeySigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, publicKeyType [32]byte, publicKey [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "addPublicKeySigned", identity, sigV, sigR, sigS, publicKeyType, publicKey, validity)
}

// AddPublicKeySigned is a paid mutator transaction binding the contract method 0x456fcfeb.
//
// Solidity: function addPublicKeySigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, publicKeyType bytes32, publicKey bytes32, validity uint256) returns()
func (_Did *DidSession) AddPublicKeySigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, publicKeyType [32]byte, publicKey [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddPublicKeySigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, publicKeyType, publicKey, validity)
}

// AddPublicKeySigned is a paid mutator transaction binding the contract method 0x456fcfeb.
//
// Solidity: function addPublicKeySigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, publicKeyType bytes32, publicKey bytes32, validity uint256) returns()
func (_Did *DidTransactorSession) AddPublicKeySigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, publicKeyType [32]byte, publicKey [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddPublicKeySigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, publicKeyType, publicKey, validity)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(identity address, newOwner address) returns()
func (_Did *DidTransactor) ChangeOwner(opts *bind.TransactOpts, identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "changeOwner", identity, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(identity address, newOwner address) returns()
func (_Did *DidSession) ChangeOwner(identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Did.Contract.ChangeOwner(&_Did.TransactOpts, identity, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(identity address, newOwner address) returns()
func (_Did *DidTransactorSession) ChangeOwner(identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Did.Contract.ChangeOwner(&_Did.TransactOpts, identity, newOwner)
}

// ChangeOwnerSigned is a paid mutator transaction binding the contract method 0x240cf1fa.
//
// Solidity: function changeOwnerSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, newOwner address) returns()
func (_Did *DidTransactor) ChangeOwnerSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "changeOwnerSigned", identity, sigV, sigR, sigS, newOwner)
}

// ChangeOwnerSigned is a paid mutator transaction binding the contract method 0x240cf1fa.
//
// Solidity: function changeOwnerSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, newOwner address) returns()
func (_Did *DidSession) ChangeOwnerSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _Did.Contract.ChangeOwnerSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, newOwner)
}

// ChangeOwnerSigned is a paid mutator transaction binding the contract method 0x240cf1fa.
//
// Solidity: function changeOwnerSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, newOwner address) returns()
func (_Did *DidTransactorSession) ChangeOwnerSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _Did.Contract.ChangeOwnerSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, newOwner)
}

// RevokeAttribute is a paid mutator transaction binding the contract method 0x00c023da.
//
// Solidity: function revokeAttribute(identity address, name bytes32, value bytes) returns()
func (_Did *DidTransactor) RevokeAttribute(opts *bind.TransactOpts, identity common.Address, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "revokeAttribute", identity, name, value)
}

// RevokeAttribute is a paid mutator transaction binding the contract method 0x00c023da.
//
// Solidity: function revokeAttribute(identity address, name bytes32, value bytes) returns()
func (_Did *DidSession) RevokeAttribute(identity common.Address, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAttribute(&_Did.TransactOpts, identity, name, value)
}

// RevokeAttribute is a paid mutator transaction binding the contract method 0x00c023da.
//
// Solidity: function revokeAttribute(identity address, name bytes32, value bytes) returns()
func (_Did *DidTransactorSession) RevokeAttribute(identity common.Address, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAttribute(&_Did.TransactOpts, identity, name, value)
}

// RevokeAttributeSigned is a paid mutator transaction binding the contract method 0xe476af5c.
//
// Solidity: function revokeAttributeSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, name bytes32, value bytes) returns()
func (_Did *DidTransactor) RevokeAttributeSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "revokeAttributeSigned", identity, sigV, sigR, sigS, name, value)
}

// RevokeAttributeSigned is a paid mutator transaction binding the contract method 0xe476af5c.
//
// Solidity: function revokeAttributeSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, name bytes32, value bytes) returns()
func (_Did *DidSession) RevokeAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAttributeSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, name, value)
}

// RevokeAttributeSigned is a paid mutator transaction binding the contract method 0xe476af5c.
//
// Solidity: function revokeAttributeSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, name bytes32, value bytes) returns()
func (_Did *DidTransactorSession) RevokeAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAttributeSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, name, value)
}

// RevokeAuthentication is a paid mutator transaction binding the contract method 0x3ea1a651.
//
// Solidity: function revokeAuthentication(identity address, authenticationType bytes32, authentication bytes32) returns()
func (_Did *DidTransactor) RevokeAuthentication(opts *bind.TransactOpts, identity common.Address, authenticationType [32]byte, authentication [32]byte) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "revokeAuthentication", identity, authenticationType, authentication)
}

// RevokeAuthentication is a paid mutator transaction binding the contract method 0x3ea1a651.
//
// Solidity: function revokeAuthentication(identity address, authenticationType bytes32, authentication bytes32) returns()
func (_Did *DidSession) RevokeAuthentication(identity common.Address, authenticationType [32]byte, authentication [32]byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAuthentication(&_Did.TransactOpts, identity, authenticationType, authentication)
}

// RevokeAuthentication is a paid mutator transaction binding the contract method 0x3ea1a651.
//
// Solidity: function revokeAuthentication(identity address, authenticationType bytes32, authentication bytes32) returns()
func (_Did *DidTransactorSession) RevokeAuthentication(identity common.Address, authenticationType [32]byte, authentication [32]byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAuthentication(&_Did.TransactOpts, identity, authenticationType, authentication)
}

// RevokeAuthenticationSigned is a paid mutator transaction binding the contract method 0xf8156eb8.
//
// Solidity: function revokeAuthenticationSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, authenticationType bytes32, authentication bytes32) returns()
func (_Did *DidTransactor) RevokeAuthenticationSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, authenticationType [32]byte, authentication [32]byte) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "revokeAuthenticationSigned", identity, sigV, sigR, sigS, authenticationType, authentication)
}

// RevokeAuthenticationSigned is a paid mutator transaction binding the contract method 0xf8156eb8.
//
// Solidity: function revokeAuthenticationSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, authenticationType bytes32, authentication bytes32) returns()
func (_Did *DidSession) RevokeAuthenticationSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, authenticationType [32]byte, authentication [32]byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAuthenticationSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, authenticationType, authentication)
}

// RevokeAuthenticationSigned is a paid mutator transaction binding the contract method 0xf8156eb8.
//
// Solidity: function revokeAuthenticationSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, authenticationType bytes32, authentication bytes32) returns()
func (_Did *DidTransactorSession) RevokeAuthenticationSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, authenticationType [32]byte, authentication [32]byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAuthenticationSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, authenticationType, authentication)
}

// RevokePublicKey is a paid mutator transaction binding the contract method 0x787768ef.
//
// Solidity: function revokePublicKey(identity address, publicKeyType bytes32, publicKey bytes32) returns()
func (_Did *DidTransactor) RevokePublicKey(opts *bind.TransactOpts, identity common.Address, publicKeyType [32]byte, publicKey [32]byte) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "revokePublicKey", identity, publicKeyType, publicKey)
}

// RevokePublicKey is a paid mutator transaction binding the contract method 0x787768ef.
//
// Solidity: function revokePublicKey(identity address, publicKeyType bytes32, publicKey bytes32) returns()
func (_Did *DidSession) RevokePublicKey(identity common.Address, publicKeyType [32]byte, publicKey [32]byte) (*types.Transaction, error) {
	return _Did.Contract.RevokePublicKey(&_Did.TransactOpts, identity, publicKeyType, publicKey)
}

// RevokePublicKey is a paid mutator transaction binding the contract method 0x787768ef.
//
// Solidity: function revokePublicKey(identity address, publicKeyType bytes32, publicKey bytes32) returns()
func (_Did *DidTransactorSession) RevokePublicKey(identity common.Address, publicKeyType [32]byte, publicKey [32]byte) (*types.Transaction, error) {
	return _Did.Contract.RevokePublicKey(&_Did.TransactOpts, identity, publicKeyType, publicKey)
}

// RevokePublicKeySigned is a paid mutator transaction binding the contract method 0x3be52425.
//
// Solidity: function revokePublicKeySigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, publicKeyType bytes32, publicKey bytes32) returns()
func (_Did *DidTransactor) RevokePublicKeySigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, publicKeyType [32]byte, publicKey [32]byte) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "revokePublicKeySigned", identity, sigV, sigR, sigS, publicKeyType, publicKey)
}

// RevokePublicKeySigned is a paid mutator transaction binding the contract method 0x3be52425.
//
// Solidity: function revokePublicKeySigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, publicKeyType bytes32, publicKey bytes32) returns()
func (_Did *DidSession) RevokePublicKeySigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, publicKeyType [32]byte, publicKey [32]byte) (*types.Transaction, error) {
	return _Did.Contract.RevokePublicKeySigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, publicKeyType, publicKey)
}

// RevokePublicKeySigned is a paid mutator transaction binding the contract method 0x3be52425.
//
// Solidity: function revokePublicKeySigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, publicKeyType bytes32, publicKey bytes32) returns()
func (_Did *DidTransactorSession) RevokePublicKeySigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, publicKeyType [32]byte, publicKey [32]byte) (*types.Transaction, error) {
	return _Did.Contract.RevokePublicKeySigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, publicKeyType, publicKey)
}

// SetAttribute is a paid mutator transaction binding the contract method 0x7ad4b0a4.
//
// Solidity: function setAttribute(identity address, name bytes32, value bytes, validity uint256) returns()
func (_Did *DidTransactor) SetAttribute(opts *bind.TransactOpts, identity common.Address, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "setAttribute", identity, name, value, validity)
}

// SetAttribute is a paid mutator transaction binding the contract method 0x7ad4b0a4.
//
// Solidity: function setAttribute(identity address, name bytes32, value bytes, validity uint256) returns()
func (_Did *DidSession) SetAttribute(identity common.Address, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.SetAttribute(&_Did.TransactOpts, identity, name, value, validity)
}

// SetAttribute is a paid mutator transaction binding the contract method 0x7ad4b0a4.
//
// Solidity: function setAttribute(identity address, name bytes32, value bytes, validity uint256) returns()
func (_Did *DidTransactorSession) SetAttribute(identity common.Address, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.SetAttribute(&_Did.TransactOpts, identity, name, value, validity)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0x123b5e98.
//
// Solidity: function setAttributeSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, name bytes32, value bytes, validity uint256) returns()
func (_Did *DidTransactor) SetAttributeSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "setAttributeSigned", identity, sigV, sigR, sigS, name, value, validity)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0x123b5e98.
//
// Solidity: function setAttributeSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, name bytes32, value bytes, validity uint256) returns()
func (_Did *DidSession) SetAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.SetAttributeSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, name, value, validity)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0x123b5e98.
//
// Solidity: function setAttributeSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, name bytes32, value bytes, validity uint256) returns()
func (_Did *DidTransactorSession) SetAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.SetAttributeSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, name, value, validity)
}

// DidDIDAttributeChangedIterator is returned from FilterDIDAttributeChanged and is used to iterate over the raw logs and unpacked data for DIDAttributeChanged events raised by the Did contract.
type DidDIDAttributeChangedIterator struct {
	Event *DidDIDAttributeChanged // Event containing the contract specifics and raw log

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
func (it *DidDIDAttributeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DidDIDAttributeChanged)
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
		it.Event = new(DidDIDAttributeChanged)
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
func (it *DidDIDAttributeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DidDIDAttributeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DidDIDAttributeChanged represents a DIDAttributeChanged event raised by the Did contract.
type DidDIDAttributeChanged struct {
	Identity       common.Address
	Name           [32]byte
	Value          []byte
	ValidTo        *big.Int
	PreviousChange *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDIDAttributeChanged is a free log retrieval operation binding the contract event 0x18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e4.
//
// Solidity: e DIDAttributeChanged(identity indexed address, name bytes32, value bytes, validTo uint256, previousChange uint256)
func (_Did *DidFilterer) FilterDIDAttributeChanged(opts *bind.FilterOpts, identity []common.Address) (*DidDIDAttributeChangedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.FilterLogs(opts, "DIDAttributeChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return &DidDIDAttributeChangedIterator{contract: _Did.contract, event: "DIDAttributeChanged", logs: logs, sub: sub}, nil
}

// WatchDIDAttributeChanged is a free log subscription operation binding the contract event 0x18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e4.
//
// Solidity: e DIDAttributeChanged(identity indexed address, name bytes32, value bytes, validTo uint256, previousChange uint256)
func (_Did *DidFilterer) WatchDIDAttributeChanged(opts *bind.WatchOpts, sink chan<- *DidDIDAttributeChanged, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.WatchLogs(opts, "DIDAttributeChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DidDIDAttributeChanged)
				if err := _Did.contract.UnpackLog(event, "DIDAttributeChanged", log); err != nil {
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

// DidDIDAuthenticationChangedIterator is returned from FilterDIDAuthenticationChanged and is used to iterate over the raw logs and unpacked data for DIDAuthenticationChanged events raised by the Did contract.
type DidDIDAuthenticationChangedIterator struct {
	Event *DidDIDAuthenticationChanged // Event containing the contract specifics and raw log

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
func (it *DidDIDAuthenticationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DidDIDAuthenticationChanged)
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
		it.Event = new(DidDIDAuthenticationChanged)
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
func (it *DidDIDAuthenticationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DidDIDAuthenticationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DidDIDAuthenticationChanged represents a DIDAuthenticationChanged event raised by the Did contract.
type DidDIDAuthenticationChanged struct {
	Identity           common.Address
	AuthenticationType [32]byte
	Authentication     [32]byte
	ValidTo            *big.Int
	PreviousChange     *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDIDAuthenticationChanged is a free log retrieval operation binding the contract event 0x6c1587d1ee8f4bdfb125ab048123336ebb2fb5f901a4561a8218e340ceffd2e0.
//
// Solidity: e DIDAuthenticationChanged(identity indexed address, authenticationType bytes32, authentication bytes32, validTo uint256, previousChange uint256)
func (_Did *DidFilterer) FilterDIDAuthenticationChanged(opts *bind.FilterOpts, identity []common.Address) (*DidDIDAuthenticationChangedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.FilterLogs(opts, "DIDAuthenticationChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return &DidDIDAuthenticationChangedIterator{contract: _Did.contract, event: "DIDAuthenticationChanged", logs: logs, sub: sub}, nil
}

// WatchDIDAuthenticationChanged is a free log subscription operation binding the contract event 0x6c1587d1ee8f4bdfb125ab048123336ebb2fb5f901a4561a8218e340ceffd2e0.
//
// Solidity: e DIDAuthenticationChanged(identity indexed address, authenticationType bytes32, authentication bytes32, validTo uint256, previousChange uint256)
func (_Did *DidFilterer) WatchDIDAuthenticationChanged(opts *bind.WatchOpts, sink chan<- *DidDIDAuthenticationChanged, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.WatchLogs(opts, "DIDAuthenticationChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DidDIDAuthenticationChanged)
				if err := _Did.contract.UnpackLog(event, "DIDAuthenticationChanged", log); err != nil {
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

// DidDIDPublicKeyChangedIterator is returned from FilterDIDPublicKeyChanged and is used to iterate over the raw logs and unpacked data for DIDPublicKeyChanged events raised by the Did contract.
type DidDIDPublicKeyChangedIterator struct {
	Event *DidDIDPublicKeyChanged // Event containing the contract specifics and raw log

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
func (it *DidDIDPublicKeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DidDIDPublicKeyChanged)
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
		it.Event = new(DidDIDPublicKeyChanged)
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
func (it *DidDIDPublicKeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DidDIDPublicKeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DidDIDPublicKeyChanged represents a DIDPublicKeyChanged event raised by the Did contract.
type DidDIDPublicKeyChanged struct {
	Identity           common.Address
	AuthenticationType [32]byte
	Authentication     [32]byte
	ValidTo            *big.Int
	PreviousChange     *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDIDPublicKeyChanged is a free log retrieval operation binding the contract event 0x213edff8255aa608dd36ca4cf7ed92d5480ddcd178865af1053f9d0bf5c77338.
//
// Solidity: e DIDPublicKeyChanged(identity indexed address, authenticationType bytes32, authentication bytes32, validTo uint256, previousChange uint256)
func (_Did *DidFilterer) FilterDIDPublicKeyChanged(opts *bind.FilterOpts, identity []common.Address) (*DidDIDPublicKeyChangedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.FilterLogs(opts, "DIDPublicKeyChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return &DidDIDPublicKeyChangedIterator{contract: _Did.contract, event: "DIDPublicKeyChanged", logs: logs, sub: sub}, nil
}

// WatchDIDPublicKeyChanged is a free log subscription operation binding the contract event 0x213edff8255aa608dd36ca4cf7ed92d5480ddcd178865af1053f9d0bf5c77338.
//
// Solidity: e DIDPublicKeyChanged(identity indexed address, authenticationType bytes32, authentication bytes32, validTo uint256, previousChange uint256)
func (_Did *DidFilterer) WatchDIDPublicKeyChanged(opts *bind.WatchOpts, sink chan<- *DidDIDPublicKeyChanged, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.WatchLogs(opts, "DIDPublicKeyChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DidDIDPublicKeyChanged)
				if err := _Did.contract.UnpackLog(event, "DIDPublicKeyChanged", log); err != nil {
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
