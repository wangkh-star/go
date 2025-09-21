// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v

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

// VotingMetaData contains all meta data concerning the Voting contract.
var VotingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voteTo\",\"type\":\"address\"}],\"name\":\"getHasVotes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voteTo\",\"type\":\"address\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasVotes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voteTo\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061083b8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061007b575f3560e01c80639ab24eb0116100595780639ab24eb0146100fb578063b9830ff11461012b578063d8bff5a514610135578063da58c7d9146101655761007b565b80633aed16cd1461007f5780636dd7d8ea146100af5780636f5012f7146100cb575b5f5ffd5b610099600480360381019061009491906105f4565b610195565b6040516100a69190610639565b60405180910390f35b6100c960048036038101906100c491906105f4565b6101e7565b005b6100e560048036038101906100e091906105f4565b6103b2565b6040516100f29190610639565b60405180910390f35b610115600480360381019061011091906105f4565b6103cf565b604051610122919061066a565b60405180910390f35b610133610414565b005b61014f600480360381019061014a91906105f4565b61050e565b60405161015c919061066a565b60405180910390f35b61017f600480360381019061017a91906106ad565b610522565b60405161018c91906106e7565b60405180910390f35b5f60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610271576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102689061075a565b60405180910390fd5b60015f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546102ba91906107a5565b5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550600133908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555050565b6002602052805f5260405f205f915054906101000a900460ff1681565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f5f90505b6001805490508110156104fe575f6001828154811061043b5761043a6107d8565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f5f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f905560025f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81549060ff0219169055508080600101915050610419565b5060015f61050c919061055d565b565b5f602052805f5260405f205f915090505481565b60018181548110610531575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5080545f8255905f5260205f2090810190610578919061057b565b50565b5b80821115610592575f815f90555060010161057c565b5090565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6105c38261059a565b9050919050565b6105d3816105b9565b81146105dd575f5ffd5b50565b5f813590506105ee816105ca565b92915050565b5f6020828403121561060957610608610596565b5b5f610616848285016105e0565b91505092915050565b5f8115159050919050565b6106338161061f565b82525050565b5f60208201905061064c5f83018461062a565b92915050565b5f819050919050565b61066481610652565b82525050565b5f60208201905061067d5f83018461065b565b92915050565b61068c81610652565b8114610696575f5ffd5b50565b5f813590506106a781610683565b92915050565b5f602082840312156106c2576106c1610596565b5b5f6106cf84828501610699565b91505092915050565b6106e1816105b9565b82525050565b5f6020820190506106fa5f8301846106d8565b92915050565b5f82825260208201905092915050565b7f20796f752061726520616c726561647920766f746500000000000000000000005f82015250565b5f610744601583610700565b915061074f82610710565b602082019050919050565b5f6020820190508181035f83015261077181610738565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6107af82610652565b91506107ba83610652565b92508282019050808211156107d2576107d1610778565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea2646970667358221220f77b3976be23b72c49862a07efab9618f97b7f2eb19c122c78def4506a53ffb864736f6c634300081e0033",
}

// VotingABI is the input ABI used to generate the binding from.
// Deprecated: Use VotingMetaData.ABI instead.
var VotingABI = VotingMetaData.ABI

// VotingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VotingMetaData.Bin instead.
var VotingBin = VotingMetaData.Bin

// DeployVoting deploys a new Ethereum contract, binding an instance of Voting to it.
func DeployVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Voting, error) {
	parsed, err := VotingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// Voting is an auto generated Go binding around an Ethereum contract.
type Voting struct {
	VotingCaller     // Read-only binding to the contract
	VotingTransactor // Write-only binding to the contract
	VotingFilterer   // Log filterer for contract events
}

// VotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingSession struct {
	Contract     *Voting           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingCallerSession struct {
	Contract *VotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingTransactorSession struct {
	Contract     *VotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingRaw struct {
	Contract *Voting // Generic contract binding to access the raw methods on
}

// VotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingCallerRaw struct {
	Contract *VotingCaller // Generic read-only contract binding to access the raw methods on
}

// VotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingTransactorRaw struct {
	Contract *VotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoting creates a new instance of Voting, bound to a specific deployed contract.
func NewVoting(address common.Address, backend bind.ContractBackend) (*Voting, error) {
	contract, err := bindVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// NewVotingCaller creates a new read-only instance of Voting, bound to a specific deployed contract.
func NewVotingCaller(address common.Address, caller bind.ContractCaller) (*VotingCaller, error) {
	contract, err := bindVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingCaller{contract: contract}, nil
}

// NewVotingTransactor creates a new write-only instance of Voting, bound to a specific deployed contract.
func NewVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingTransactor, error) {
	contract, err := bindVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingTransactor{contract: contract}, nil
}

// NewVotingFilterer creates a new log filterer instance of Voting, bound to a specific deployed contract.
func NewVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingFilterer, error) {
	contract, err := bindVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingFilterer{contract: contract}, nil
}

// bindVoting binds a generic wrapper to an already deployed contract.
func bindVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VotingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.VotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transact(opts, method, params...)
}

// GetHasVotes is a free data retrieval call binding the contract method 0x3aed16cd.
//
// Solidity: function getHasVotes(address voteTo) view returns(bool)
func (_Voting *VotingCaller) GetHasVotes(opts *bind.CallOpts, voteTo common.Address) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getHasVotes", voteTo)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetHasVotes is a free data retrieval call binding the contract method 0x3aed16cd.
//
// Solidity: function getHasVotes(address voteTo) view returns(bool)
func (_Voting *VotingSession) GetHasVotes(voteTo common.Address) (bool, error) {
	return _Voting.Contract.GetHasVotes(&_Voting.CallOpts, voteTo)
}

// GetHasVotes is a free data retrieval call binding the contract method 0x3aed16cd.
//
// Solidity: function getHasVotes(address voteTo) view returns(bool)
func (_Voting *VotingCallerSession) GetHasVotes(voteTo common.Address) (bool, error) {
	return _Voting.Contract.GetHasVotes(&_Voting.CallOpts, voteTo)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address voteTo) view returns(uint256)
func (_Voting *VotingCaller) GetVotes(opts *bind.CallOpts, voteTo common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getVotes", voteTo)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address voteTo) view returns(uint256)
func (_Voting *VotingSession) GetVotes(voteTo common.Address) (*big.Int, error) {
	return _Voting.Contract.GetVotes(&_Voting.CallOpts, voteTo)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address voteTo) view returns(uint256)
func (_Voting *VotingCallerSession) GetVotes(voteTo common.Address) (*big.Int, error) {
	return _Voting.Contract.GetVotes(&_Voting.CallOpts, voteTo)
}

// HasVotes is a free data retrieval call binding the contract method 0x6f5012f7.
//
// Solidity: function hasVotes(address ) view returns(bool)
func (_Voting *VotingCaller) HasVotes(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "hasVotes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVotes is a free data retrieval call binding the contract method 0x6f5012f7.
//
// Solidity: function hasVotes(address ) view returns(bool)
func (_Voting *VotingSession) HasVotes(arg0 common.Address) (bool, error) {
	return _Voting.Contract.HasVotes(&_Voting.CallOpts, arg0)
}

// HasVotes is a free data retrieval call binding the contract method 0x6f5012f7.
//
// Solidity: function hasVotes(address ) view returns(bool)
func (_Voting *VotingCallerSession) HasVotes(arg0 common.Address) (bool, error) {
	return _Voting.Contract.HasVotes(&_Voting.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters(uint256 ) view returns(address)
func (_Voting *VotingCaller) Voters(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "voters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters(uint256 ) view returns(address)
func (_Voting *VotingSession) Voters(arg0 *big.Int) (common.Address, error) {
	return _Voting.Contract.Voters(&_Voting.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters(uint256 ) view returns(address)
func (_Voting *VotingCallerSession) Voters(arg0 *big.Int) (common.Address, error) {
	return _Voting.Contract.Voters(&_Voting.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(uint256)
func (_Voting *VotingCaller) Votes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "votes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(uint256)
func (_Voting *VotingSession) Votes(arg0 common.Address) (*big.Int, error) {
	return _Voting.Contract.Votes(&_Voting.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(uint256)
func (_Voting *VotingCallerSession) Votes(arg0 common.Address) (*big.Int, error) {
	return _Voting.Contract.Votes(&_Voting.CallOpts, arg0)
}

// ResetVotes is a paid mutator transaction binding the contract method 0xb9830ff1.
//
// Solidity: function resetVotes() returns()
func (_Voting *VotingTransactor) ResetVotes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "resetVotes")
}

// ResetVotes is a paid mutator transaction binding the contract method 0xb9830ff1.
//
// Solidity: function resetVotes() returns()
func (_Voting *VotingSession) ResetVotes() (*types.Transaction, error) {
	return _Voting.Contract.ResetVotes(&_Voting.TransactOpts)
}

// ResetVotes is a paid mutator transaction binding the contract method 0xb9830ff1.
//
// Solidity: function resetVotes() returns()
func (_Voting *VotingTransactorSession) ResetVotes() (*types.Transaction, error) {
	return _Voting.Contract.ResetVotes(&_Voting.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address voteTo) returns()
func (_Voting *VotingTransactor) Vote(opts *bind.TransactOpts, voteTo common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "vote", voteTo)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address voteTo) returns()
func (_Voting *VotingSession) Vote(voteTo common.Address) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, voteTo)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address voteTo) returns()
func (_Voting *VotingTransactorSession) Vote(voteTo common.Address) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, voteTo)
}
