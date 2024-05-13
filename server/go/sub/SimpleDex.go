// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sub

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
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tradeToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"buyOrderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sellOrderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"OrderExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBuy\",\"type\":\"bool\"}],\"name\":\"OrderPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ethBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"buyOrderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellOrderId\",\"type\":\"uint256\"}],\"name\":\"executeOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balanceETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOrderId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBuy\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBuy\",\"type\":\"bool\"}],\"name\":\"placeOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// EthBalances is a free data retrieval call binding the contract method 0x3cfba0e3.
//
// Solidity: function ethBalances(address ) view returns(uint256)
func (_Contract *ContractCaller) EthBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ethBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthBalances is a free data retrieval call binding the contract method 0x3cfba0e3.
//
// Solidity: function ethBalances(address ) view returns(uint256)
func (_Contract *ContractSession) EthBalances(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.EthBalances(&_Contract.CallOpts, arg0)
}

// EthBalances is a free data retrieval call binding the contract method 0x3cfba0e3.
//
// Solidity: function ethBalances(address ) view returns(uint256)
func (_Contract *ContractCallerSession) EthBalances(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.EthBalances(&_Contract.CallOpts, arg0)
}

// GetBalances is a free data retrieval call binding the contract method 0xc84aae17.
//
// Solidity: function getBalances(address _user) view returns(uint256 balanceETH, uint256 balanceToken)
func (_Contract *ContractCaller) GetBalances(opts *bind.CallOpts, _user common.Address) (struct {
	BalanceETH   *big.Int
	BalanceToken *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getBalances", _user)

	outstruct := new(struct {
		BalanceETH   *big.Int
		BalanceToken *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BalanceETH = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BalanceToken = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBalances is a free data retrieval call binding the contract method 0xc84aae17.
//
// Solidity: function getBalances(address _user) view returns(uint256 balanceETH, uint256 balanceToken)
func (_Contract *ContractSession) GetBalances(_user common.Address) (struct {
	BalanceETH   *big.Int
	BalanceToken *big.Int
}, error) {
	return _Contract.Contract.GetBalances(&_Contract.CallOpts, _user)
}

// GetBalances is a free data retrieval call binding the contract method 0xc84aae17.
//
// Solidity: function getBalances(address _user) view returns(uint256 balanceETH, uint256 balanceToken)
func (_Contract *ContractCallerSession) GetBalances(_user common.Address) (struct {
	BalanceETH   *big.Int
	BalanceToken *big.Int
}, error) {
	return _Contract.Contract.GetBalances(&_Contract.CallOpts, _user)
}

// NextOrderId is a free data retrieval call binding the contract method 0x2a58b330.
//
// Solidity: function nextOrderId() view returns(uint256)
func (_Contract *ContractCaller) NextOrderId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nextOrderId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextOrderId is a free data retrieval call binding the contract method 0x2a58b330.
//
// Solidity: function nextOrderId() view returns(uint256)
func (_Contract *ContractSession) NextOrderId() (*big.Int, error) {
	return _Contract.Contract.NextOrderId(&_Contract.CallOpts)
}

// NextOrderId is a free data retrieval call binding the contract method 0x2a58b330.
//
// Solidity: function nextOrderId() view returns(uint256)
func (_Contract *ContractCallerSession) NextOrderId() (*big.Int, error) {
	return _Contract.Contract.NextOrderId(&_Contract.CallOpts)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(uint256 id, address user, uint256 amount, uint256 price, bool isBuy)
func (_Contract *ContractCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id     *big.Int
	User   common.Address
	Amount *big.Int
	Price  *big.Int
	IsBuy  bool
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		Id     *big.Int
		User   common.Address
		Amount *big.Int
		Price  *big.Int
		IsBuy  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.User = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsBuy = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(uint256 id, address user, uint256 amount, uint256 price, bool isBuy)
func (_Contract *ContractSession) Orders(arg0 *big.Int) (struct {
	Id     *big.Int
	User   common.Address
	Amount *big.Int
	Price  *big.Int
	IsBuy  bool
}, error) {
	return _Contract.Contract.Orders(&_Contract.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(uint256 id, address user, uint256 amount, uint256 price, bool isBuy)
func (_Contract *ContractCallerSession) Orders(arg0 *big.Int) (struct {
	Id     *big.Int
	User   common.Address
	Amount *big.Int
	Price  *big.Int
	IsBuy  bool
}, error) {
	return _Contract.Contract.Orders(&_Contract.CallOpts, arg0)
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

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_Contract *ContractCaller) TokenBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tokenBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_Contract *ContractSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.TokenBalances(&_Contract.CallOpts, arg0)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_Contract *ContractCallerSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.TokenBalances(&_Contract.CallOpts, arg0)
}

// TradeToken is a free data retrieval call binding the contract method 0xd83678ac.
//
// Solidity: function tradeToken() view returns(address)
func (_Contract *ContractCaller) TradeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tradeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TradeToken is a free data retrieval call binding the contract method 0xd83678ac.
//
// Solidity: function tradeToken() view returns(address)
func (_Contract *ContractSession) TradeToken() (common.Address, error) {
	return _Contract.Contract.TradeToken(&_Contract.CallOpts)
}

// TradeToken is a free data retrieval call binding the contract method 0xd83678ac.
//
// Solidity: function tradeToken() view returns(address)
func (_Contract *ContractCallerSession) TradeToken() (common.Address, error) {
	return _Contract.Contract.TradeToken(&_Contract.CallOpts)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 orderId) returns()
func (_Contract *ContractTransactor) CancelOrder(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "cancelOrder", orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 orderId) returns()
func (_Contract *ContractSession) CancelOrder(orderId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CancelOrder(&_Contract.TransactOpts, orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 orderId) returns()
func (_Contract *ContractTransactorSession) CancelOrder(orderId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CancelOrder(&_Contract.TransactOpts, orderId)
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns()
func (_Contract *ContractTransactor) DepositETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "depositETH")
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns()
func (_Contract *ContractSession) DepositETH() (*types.Transaction, error) {
	return _Contract.Contract.DepositETH(&_Contract.TransactOpts)
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns()
func (_Contract *ContractTransactorSession) DepositETH() (*types.Transaction, error) {
	return _Contract.Contract.DepositETH(&_Contract.TransactOpts)
}

// DepositToken is a paid mutator transaction binding the contract method 0x6215be77.
//
// Solidity: function depositToken(uint256 amount) returns()
func (_Contract *ContractTransactor) DepositToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "depositToken", amount)
}

// DepositToken is a paid mutator transaction binding the contract method 0x6215be77.
//
// Solidity: function depositToken(uint256 amount) returns()
func (_Contract *ContractSession) DepositToken(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DepositToken(&_Contract.TransactOpts, amount)
}

// DepositToken is a paid mutator transaction binding the contract method 0x6215be77.
//
// Solidity: function depositToken(uint256 amount) returns()
func (_Contract *ContractTransactorSession) DepositToken(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DepositToken(&_Contract.TransactOpts, amount)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0xef46e0ca.
//
// Solidity: function executeOrder(uint256 buyOrderId, uint256 sellOrderId) returns()
func (_Contract *ContractTransactor) ExecuteOrder(opts *bind.TransactOpts, buyOrderId *big.Int, sellOrderId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "executeOrder", buyOrderId, sellOrderId)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0xef46e0ca.
//
// Solidity: function executeOrder(uint256 buyOrderId, uint256 sellOrderId) returns()
func (_Contract *ContractSession) ExecuteOrder(buyOrderId *big.Int, sellOrderId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ExecuteOrder(&_Contract.TransactOpts, buyOrderId, sellOrderId)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0xef46e0ca.
//
// Solidity: function executeOrder(uint256 buyOrderId, uint256 sellOrderId) returns()
func (_Contract *ContractTransactorSession) ExecuteOrder(buyOrderId *big.Int, sellOrderId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ExecuteOrder(&_Contract.TransactOpts, buyOrderId, sellOrderId)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0xbd2d447d.
//
// Solidity: function placeOrder(uint256 amount, uint256 price, bool isBuy) returns()
func (_Contract *ContractTransactor) PlaceOrder(opts *bind.TransactOpts, amount *big.Int, price *big.Int, isBuy bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "placeOrder", amount, price, isBuy)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0xbd2d447d.
//
// Solidity: function placeOrder(uint256 amount, uint256 price, bool isBuy) returns()
func (_Contract *ContractSession) PlaceOrder(amount *big.Int, price *big.Int, isBuy bool) (*types.Transaction, error) {
	return _Contract.Contract.PlaceOrder(&_Contract.TransactOpts, amount, price, isBuy)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0xbd2d447d.
//
// Solidity: function placeOrder(uint256 amount, uint256 price, bool isBuy) returns()
func (_Contract *ContractTransactorSession) PlaceOrder(amount *big.Int, price *big.Int, isBuy bool) (*types.Transaction, error) {
	return _Contract.Contract.PlaceOrder(&_Contract.TransactOpts, amount, price, isBuy)
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

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_Contract *ContractTransactor) WithdrawETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawETH", amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_Contract *ContractSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawETH(&_Contract.TransactOpts, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_Contract *ContractTransactorSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawETH(&_Contract.TransactOpts, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x50baa622.
//
// Solidity: function withdrawToken(uint256 amount) returns()
func (_Contract *ContractTransactor) WithdrawToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawToken", amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x50baa622.
//
// Solidity: function withdrawToken(uint256 amount) returns()
func (_Contract *ContractSession) WithdrawToken(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawToken(&_Contract.TransactOpts, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x50baa622.
//
// Solidity: function withdrawToken(uint256 amount) returns()
func (_Contract *ContractTransactorSession) WithdrawToken(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawToken(&_Contract.TransactOpts, amount)
}

// ContractOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Contract contract.
type ContractOrderCancelledIterator struct {
	Event *ContractOrderCancelled // Event containing the contract specifics and raw log

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
func (it *ContractOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOrderCancelled)
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
		it.Event = new(ContractOrderCancelled)
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
func (it *ContractOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOrderCancelled represents a OrderCancelled event raised by the Contract contract.
type ContractOrderCancelled struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x61b9399f2f0f32ca39ce8d7be32caed5ec22fe07a6daba3a467ed479ec606582.
//
// Solidity: event OrderCancelled(uint256 indexed orderId)
func (_Contract *ContractFilterer) FilterOrderCancelled(opts *bind.FilterOpts, orderId []*big.Int) (*ContractOrderCancelledIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OrderCancelled", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractOrderCancelledIterator{contract: _Contract.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x61b9399f2f0f32ca39ce8d7be32caed5ec22fe07a6daba3a467ed479ec606582.
//
// Solidity: event OrderCancelled(uint256 indexed orderId)
func (_Contract *ContractFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *ContractOrderCancelled, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OrderCancelled", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOrderCancelled)
				if err := _Contract.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x61b9399f2f0f32ca39ce8d7be32caed5ec22fe07a6daba3a467ed479ec606582.
//
// Solidity: event OrderCancelled(uint256 indexed orderId)
func (_Contract *ContractFilterer) ParseOrderCancelled(log types.Log) (*ContractOrderCancelled, error) {
	event := new(ContractOrderCancelled)
	if err := _Contract.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOrderExecutedIterator is returned from FilterOrderExecuted and is used to iterate over the raw logs and unpacked data for OrderExecuted events raised by the Contract contract.
type ContractOrderExecutedIterator struct {
	Event *ContractOrderExecuted // Event containing the contract specifics and raw log

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
func (it *ContractOrderExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOrderExecuted)
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
		it.Event = new(ContractOrderExecuted)
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
func (it *ContractOrderExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOrderExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOrderExecuted represents a OrderExecuted event raised by the Contract contract.
type ContractOrderExecuted struct {
	BuyOrderId  *big.Int
	SellOrderId *big.Int
	Amount      *big.Int
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrderExecuted is a free log retrieval operation binding the contract event 0xd03c0b3fc99c13b1deab0d6c6efb77b59dbe857f327840eca401774933e31120.
//
// Solidity: event OrderExecuted(uint256 indexed buyOrderId, uint256 indexed sellOrderId, uint256 amount, uint256 price)
func (_Contract *ContractFilterer) FilterOrderExecuted(opts *bind.FilterOpts, buyOrderId []*big.Int, sellOrderId []*big.Int) (*ContractOrderExecutedIterator, error) {

	var buyOrderIdRule []interface{}
	for _, buyOrderIdItem := range buyOrderId {
		buyOrderIdRule = append(buyOrderIdRule, buyOrderIdItem)
	}
	var sellOrderIdRule []interface{}
	for _, sellOrderIdItem := range sellOrderId {
		sellOrderIdRule = append(sellOrderIdRule, sellOrderIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OrderExecuted", buyOrderIdRule, sellOrderIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractOrderExecutedIterator{contract: _Contract.contract, event: "OrderExecuted", logs: logs, sub: sub}, nil
}

// WatchOrderExecuted is a free log subscription operation binding the contract event 0xd03c0b3fc99c13b1deab0d6c6efb77b59dbe857f327840eca401774933e31120.
//
// Solidity: event OrderExecuted(uint256 indexed buyOrderId, uint256 indexed sellOrderId, uint256 amount, uint256 price)
func (_Contract *ContractFilterer) WatchOrderExecuted(opts *bind.WatchOpts, sink chan<- *ContractOrderExecuted, buyOrderId []*big.Int, sellOrderId []*big.Int) (event.Subscription, error) {

	var buyOrderIdRule []interface{}
	for _, buyOrderIdItem := range buyOrderId {
		buyOrderIdRule = append(buyOrderIdRule, buyOrderIdItem)
	}
	var sellOrderIdRule []interface{}
	for _, sellOrderIdItem := range sellOrderId {
		sellOrderIdRule = append(sellOrderIdRule, sellOrderIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OrderExecuted", buyOrderIdRule, sellOrderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOrderExecuted)
				if err := _Contract.contract.UnpackLog(event, "OrderExecuted", log); err != nil {
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

// ParseOrderExecuted is a log parse operation binding the contract event 0xd03c0b3fc99c13b1deab0d6c6efb77b59dbe857f327840eca401774933e31120.
//
// Solidity: event OrderExecuted(uint256 indexed buyOrderId, uint256 indexed sellOrderId, uint256 amount, uint256 price)
func (_Contract *ContractFilterer) ParseOrderExecuted(log types.Log) (*ContractOrderExecuted, error) {
	event := new(ContractOrderExecuted)
	if err := _Contract.contract.UnpackLog(event, "OrderExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOrderPlacedIterator is returned from FilterOrderPlaced and is used to iterate over the raw logs and unpacked data for OrderPlaced events raised by the Contract contract.
type ContractOrderPlacedIterator struct {
	Event *ContractOrderPlaced // Event containing the contract specifics and raw log

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
func (it *ContractOrderPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOrderPlaced)
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
		it.Event = new(ContractOrderPlaced)
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
func (it *ContractOrderPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOrderPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOrderPlaced represents a OrderPlaced event raised by the Contract contract.
type ContractOrderPlaced struct {
	OrderId *big.Int
	User    common.Address
	Amount  *big.Int
	Price   *big.Int
	IsBuy   bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderPlaced is a free log retrieval operation binding the contract event 0x1ad1a5ae4ffda7ac4728a3f999d312ee8c3bf93f8f9bec6ea55609ec2980c0d7.
//
// Solidity: event OrderPlaced(uint256 indexed orderId, address indexed user, uint256 amount, uint256 price, bool isBuy)
func (_Contract *ContractFilterer) FilterOrderPlaced(opts *bind.FilterOpts, orderId []*big.Int, user []common.Address) (*ContractOrderPlacedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OrderPlaced", orderIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ContractOrderPlacedIterator{contract: _Contract.contract, event: "OrderPlaced", logs: logs, sub: sub}, nil
}

// WatchOrderPlaced is a free log subscription operation binding the contract event 0x1ad1a5ae4ffda7ac4728a3f999d312ee8c3bf93f8f9bec6ea55609ec2980c0d7.
//
// Solidity: event OrderPlaced(uint256 indexed orderId, address indexed user, uint256 amount, uint256 price, bool isBuy)
func (_Contract *ContractFilterer) WatchOrderPlaced(opts *bind.WatchOpts, sink chan<- *ContractOrderPlaced, orderId []*big.Int, user []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OrderPlaced", orderIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOrderPlaced)
				if err := _Contract.contract.UnpackLog(event, "OrderPlaced", log); err != nil {
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

// ParseOrderPlaced is a log parse operation binding the contract event 0x1ad1a5ae4ffda7ac4728a3f999d312ee8c3bf93f8f9bec6ea55609ec2980c0d7.
//
// Solidity: event OrderPlaced(uint256 indexed orderId, address indexed user, uint256 amount, uint256 price, bool isBuy)
func (_Contract *ContractFilterer) ParseOrderPlaced(log types.Log) (*ContractOrderPlaced, error) {
	event := new(ContractOrderPlaced)
	if err := _Contract.contract.UnpackLog(event, "OrderPlaced", log); err != nil {
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
