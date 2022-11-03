// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappers

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

// BinaryMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type BinaryMerkleProof struct {
	SideNodes [][32]byte
	Key       *big.Int
	NumLeaves *big.Int
}

// DataRootTuple is an auto generated low-level Go binding around an user-defined struct.
type DataRootTuple struct {
	Height   *big.Int
	DataRoot [32]byte
}

// NamespaceMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type NamespaceMerkleProof struct {
	SideNodes []NamespaceNode
	Key       *big.Int
	NumLeaves *big.Int
}

// NamespaceNode is an auto generated low-level Go binding around an user-defined struct.
type NamespaceNode struct {
	Min    [8]byte
	Max    [8]byte
	Digest [32]byte
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// Validator is an auto generated low-level Go binding around an user-defined struct.
type Validator struct {
	Addr  common.Address
	Power *big.Int
}

// BinaryMerkleTreeMetaData contains all meta data concerning the BinaryMerkleTree contract.
var BinaryMerkleTreeMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207eadb80a7580ae404794766e2bdef99e27e061ef3413cdabf850b8262fd6415964736f6c63430008090033",
}

// BinaryMerkleTreeABI is the input ABI used to generate the binding from.
// Deprecated: Use BinaryMerkleTreeMetaData.ABI instead.
var BinaryMerkleTreeABI = BinaryMerkleTreeMetaData.ABI

// BinaryMerkleTreeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BinaryMerkleTreeMetaData.Bin instead.
var BinaryMerkleTreeBin = BinaryMerkleTreeMetaData.Bin

// DeployBinaryMerkleTree deploys a new Ethereum contract, binding an instance of BinaryMerkleTree to it.
func DeployBinaryMerkleTree(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BinaryMerkleTree, error) {
	parsed, err := BinaryMerkleTreeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BinaryMerkleTreeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BinaryMerkleTree{BinaryMerkleTreeCaller: BinaryMerkleTreeCaller{contract: contract}, BinaryMerkleTreeTransactor: BinaryMerkleTreeTransactor{contract: contract}, BinaryMerkleTreeFilterer: BinaryMerkleTreeFilterer{contract: contract}}, nil
}

// BinaryMerkleTree is an auto generated Go binding around an Ethereum contract.
type BinaryMerkleTree struct {
	BinaryMerkleTreeCaller     // Read-only binding to the contract
	BinaryMerkleTreeTransactor // Write-only binding to the contract
	BinaryMerkleTreeFilterer   // Log filterer for contract events
}

// BinaryMerkleTreeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BinaryMerkleTreeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BinaryMerkleTreeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BinaryMerkleTreeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BinaryMerkleTreeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BinaryMerkleTreeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BinaryMerkleTreeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BinaryMerkleTreeSession struct {
	Contract     *BinaryMerkleTree // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BinaryMerkleTreeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BinaryMerkleTreeCallerSession struct {
	Contract *BinaryMerkleTreeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BinaryMerkleTreeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BinaryMerkleTreeTransactorSession struct {
	Contract     *BinaryMerkleTreeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BinaryMerkleTreeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BinaryMerkleTreeRaw struct {
	Contract *BinaryMerkleTree // Generic contract binding to access the raw methods on
}

// BinaryMerkleTreeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BinaryMerkleTreeCallerRaw struct {
	Contract *BinaryMerkleTreeCaller // Generic read-only contract binding to access the raw methods on
}

// BinaryMerkleTreeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BinaryMerkleTreeTransactorRaw struct {
	Contract *BinaryMerkleTreeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBinaryMerkleTree creates a new instance of BinaryMerkleTree, bound to a specific deployed contract.
func NewBinaryMerkleTree(address common.Address, backend bind.ContractBackend) (*BinaryMerkleTree, error) {
	contract, err := bindBinaryMerkleTree(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BinaryMerkleTree{BinaryMerkleTreeCaller: BinaryMerkleTreeCaller{contract: contract}, BinaryMerkleTreeTransactor: BinaryMerkleTreeTransactor{contract: contract}, BinaryMerkleTreeFilterer: BinaryMerkleTreeFilterer{contract: contract}}, nil
}

// NewBinaryMerkleTreeCaller creates a new read-only instance of BinaryMerkleTree, bound to a specific deployed contract.
func NewBinaryMerkleTreeCaller(address common.Address, caller bind.ContractCaller) (*BinaryMerkleTreeCaller, error) {
	contract, err := bindBinaryMerkleTree(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BinaryMerkleTreeCaller{contract: contract}, nil
}

// NewBinaryMerkleTreeTransactor creates a new write-only instance of BinaryMerkleTree, bound to a specific deployed contract.
func NewBinaryMerkleTreeTransactor(address common.Address, transactor bind.ContractTransactor) (*BinaryMerkleTreeTransactor, error) {
	contract, err := bindBinaryMerkleTree(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BinaryMerkleTreeTransactor{contract: contract}, nil
}

// NewBinaryMerkleTreeFilterer creates a new log filterer instance of BinaryMerkleTree, bound to a specific deployed contract.
func NewBinaryMerkleTreeFilterer(address common.Address, filterer bind.ContractFilterer) (*BinaryMerkleTreeFilterer, error) {
	contract, err := bindBinaryMerkleTree(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BinaryMerkleTreeFilterer{contract: contract}, nil
}

// bindBinaryMerkleTree binds a generic wrapper to an already deployed contract.
func bindBinaryMerkleTree(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BinaryMerkleTreeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BinaryMerkleTree *BinaryMerkleTreeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BinaryMerkleTree.Contract.BinaryMerkleTreeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BinaryMerkleTree *BinaryMerkleTreeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BinaryMerkleTree.Contract.BinaryMerkleTreeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BinaryMerkleTree *BinaryMerkleTreeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BinaryMerkleTree.Contract.BinaryMerkleTreeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BinaryMerkleTree *BinaryMerkleTreeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BinaryMerkleTree.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BinaryMerkleTree *BinaryMerkleTreeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BinaryMerkleTree.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BinaryMerkleTree *BinaryMerkleTreeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BinaryMerkleTree.Contract.contract.Transact(opts, method, params...)
}

// ConstantsMetaData contains all meta data concerning the Constants contract.
var ConstantsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208f868974afcbbf9b36869e2284f3eb5ebe16ddff65eb1593f2f10039702dce5664736f6c63430008090033",
}

// ConstantsABI is the input ABI used to generate the binding from.
// Deprecated: Use ConstantsMetaData.ABI instead.
var ConstantsABI = ConstantsMetaData.ABI

// ConstantsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConstantsMetaData.Bin instead.
var ConstantsBin = ConstantsMetaData.Bin

// DeployConstants deploys a new Ethereum contract, binding an instance of Constants to it.
func DeployConstants(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Constants, error) {
	parsed, err := ConstantsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConstantsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Constants{ConstantsCaller: ConstantsCaller{contract: contract}, ConstantsTransactor: ConstantsTransactor{contract: contract}, ConstantsFilterer: ConstantsFilterer{contract: contract}}, nil
}

// Constants is an auto generated Go binding around an Ethereum contract.
type Constants struct {
	ConstantsCaller     // Read-only binding to the contract
	ConstantsTransactor // Write-only binding to the contract
	ConstantsFilterer   // Log filterer for contract events
}

// ConstantsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConstantsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConstantsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConstantsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConstantsSession struct {
	Contract     *Constants        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConstantsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConstantsCallerSession struct {
	Contract *ConstantsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ConstantsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConstantsTransactorSession struct {
	Contract     *ConstantsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ConstantsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConstantsRaw struct {
	Contract *Constants // Generic contract binding to access the raw methods on
}

// ConstantsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConstantsCallerRaw struct {
	Contract *ConstantsCaller // Generic read-only contract binding to access the raw methods on
}

// ConstantsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConstantsTransactorRaw struct {
	Contract *ConstantsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConstants creates a new instance of Constants, bound to a specific deployed contract.
func NewConstants(address common.Address, backend bind.ContractBackend) (*Constants, error) {
	contract, err := bindConstants(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Constants{ConstantsCaller: ConstantsCaller{contract: contract}, ConstantsTransactor: ConstantsTransactor{contract: contract}, ConstantsFilterer: ConstantsFilterer{contract: contract}}, nil
}

// NewConstantsCaller creates a new read-only instance of Constants, bound to a specific deployed contract.
func NewConstantsCaller(address common.Address, caller bind.ContractCaller) (*ConstantsCaller, error) {
	contract, err := bindConstants(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConstantsCaller{contract: contract}, nil
}

// NewConstantsTransactor creates a new write-only instance of Constants, bound to a specific deployed contract.
func NewConstantsTransactor(address common.Address, transactor bind.ContractTransactor) (*ConstantsTransactor, error) {
	contract, err := bindConstants(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConstantsTransactor{contract: contract}, nil
}

// NewConstantsFilterer creates a new log filterer instance of Constants, bound to a specific deployed contract.
func NewConstantsFilterer(address common.Address, filterer bind.ContractFilterer) (*ConstantsFilterer, error) {
	contract, err := bindConstants(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConstantsFilterer{contract: contract}, nil
}

// bindConstants binds a generic wrapper to an already deployed contract.
func bindConstants(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConstantsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Constants *ConstantsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Constants.Contract.ConstantsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Constants *ConstantsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constants.Contract.ConstantsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Constants *ConstantsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Constants.Contract.ConstantsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Constants *ConstantsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Constants.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Constants *ConstantsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constants.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Constants *ConstantsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Constants.Contract.contract.Transact(opts, method, params...)
}

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220333ac9e98839ed5aa8c16d658fffd68e1cdbc5fb35ef9efab71b442b1f20217b64736f6c63430008090033",
}

// ECDSAABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSAMetaData.ABI instead.
var ECDSAABI = ECDSAMetaData.ABI

// ECDSABin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ECDSAMetaData.Bin instead.
var ECDSABin = ECDSAMetaData.Bin

// DeployECDSA deploys a new Ethereum contract, binding an instance of ECDSA to it.
func DeployECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSA, error) {
	parsed, err := ECDSAMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ECDSABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// ECDSA is an auto generated Go binding around an Ethereum contract.
type ECDSA struct {
	ECDSACaller     // Read-only binding to the contract
	ECDSATransactor // Write-only binding to the contract
	ECDSAFilterer   // Log filterer for contract events
}

// ECDSACaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSASession struct {
	Contract     *ECDSA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACallerSession struct {
	Contract *ECDSACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ECDSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATransactorSession struct {
	Contract     *ECDSATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSARaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSARaw struct {
	Contract *ECDSA // Generic contract binding to access the raw methods on
}

// ECDSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACallerRaw struct {
	Contract *ECDSACaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATransactorRaw struct {
	Contract *ECDSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSA creates a new instance of ECDSA, bound to a specific deployed contract.
func NewECDSA(address common.Address, backend bind.ContractBackend) (*ECDSA, error) {
	contract, err := bindECDSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// NewECDSACaller creates a new read-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSACaller(address common.Address, caller bind.ContractCaller) (*ECDSACaller, error) {
	contract, err := bindECDSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACaller{contract: contract}, nil
}

// NewECDSATransactor creates a new write-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSATransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATransactor, error) {
	contract, err := bindECDSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATransactor{contract: contract}, nil
}

// NewECDSAFilterer creates a new log filterer instance of ECDSA, bound to a specific deployed contract.
func NewECDSAFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAFilterer, error) {
	contract, err := bindECDSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAFilterer{contract: contract}, nil
}

// bindECDSA binds a generic wrapper to an already deployed contract.
func bindECDSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.ECDSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transact(opts, method, params...)
}

// IDAOracleMetaData contains all meta data concerning the IDAOracle contract.
var IDAOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes8\",\"name\":\"min\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"max\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"_namespacedRowRoots\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_sharesSets\",\"type\":\"bytes[]\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes8\",\"name\":\"min\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"max\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"sideNodes\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structNamespaceMerkleProof[]\",\"name\":\"_sharesSetsProofs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes8\",\"name\":\"_namespaceID\",\"type\":\"bytes8\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof[]\",\"name\":\"_rowsProofs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_rowRoots\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"_tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"_tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"_dataRootProof\",\"type\":\"tuple\"}],\"name\":\"verifySharesInclusion\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a93527b3": "verifySharesInclusion((bytes8,bytes8,bytes32)[],bytes[],((bytes8,bytes8,bytes32)[],uint256,uint256)[],bytes8,(bytes32[],uint256,uint256)[],bytes[],uint256,(uint256,bytes32),(bytes32[],uint256,uint256))",
	},
}

// IDAOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use IDAOracleMetaData.ABI instead.
var IDAOracleABI = IDAOracleMetaData.ABI

// Deprecated: Use IDAOracleMetaData.Sigs instead.
// IDAOracleFuncSigs maps the 4-byte function signature to its string representation.
var IDAOracleFuncSigs = IDAOracleMetaData.Sigs

// IDAOracle is an auto generated Go binding around an Ethereum contract.
type IDAOracle struct {
	IDAOracleCaller     // Read-only binding to the contract
	IDAOracleTransactor // Write-only binding to the contract
	IDAOracleFilterer   // Log filterer for contract events
}

// IDAOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDAOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDAOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDAOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDAOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDAOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDAOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDAOracleSession struct {
	Contract     *IDAOracle        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDAOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDAOracleCallerSession struct {
	Contract *IDAOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IDAOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDAOracleTransactorSession struct {
	Contract     *IDAOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IDAOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDAOracleRaw struct {
	Contract *IDAOracle // Generic contract binding to access the raw methods on
}

// IDAOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDAOracleCallerRaw struct {
	Contract *IDAOracleCaller // Generic read-only contract binding to access the raw methods on
}

// IDAOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDAOracleTransactorRaw struct {
	Contract *IDAOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDAOracle creates a new instance of IDAOracle, bound to a specific deployed contract.
func NewIDAOracle(address common.Address, backend bind.ContractBackend) (*IDAOracle, error) {
	contract, err := bindIDAOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDAOracle{IDAOracleCaller: IDAOracleCaller{contract: contract}, IDAOracleTransactor: IDAOracleTransactor{contract: contract}, IDAOracleFilterer: IDAOracleFilterer{contract: contract}}, nil
}

// NewIDAOracleCaller creates a new read-only instance of IDAOracle, bound to a specific deployed contract.
func NewIDAOracleCaller(address common.Address, caller bind.ContractCaller) (*IDAOracleCaller, error) {
	contract, err := bindIDAOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDAOracleCaller{contract: contract}, nil
}

// NewIDAOracleTransactor creates a new write-only instance of IDAOracle, bound to a specific deployed contract.
func NewIDAOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*IDAOracleTransactor, error) {
	contract, err := bindIDAOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDAOracleTransactor{contract: contract}, nil
}

// NewIDAOracleFilterer creates a new log filterer instance of IDAOracle, bound to a specific deployed contract.
func NewIDAOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*IDAOracleFilterer, error) {
	contract, err := bindIDAOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDAOracleFilterer{contract: contract}, nil
}

// bindIDAOracle binds a generic wrapper to an already deployed contract.
func bindIDAOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDAOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDAOracle *IDAOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDAOracle.Contract.IDAOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDAOracle *IDAOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDAOracle.Contract.IDAOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDAOracle *IDAOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDAOracle.Contract.IDAOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDAOracle *IDAOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDAOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDAOracle *IDAOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDAOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDAOracle *IDAOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDAOracle.Contract.contract.Transact(opts, method, params...)
}

// VerifySharesInclusion is a free data retrieval call binding the contract method 0xa93527b3.
//
// Solidity: function verifySharesInclusion((bytes8,bytes8,bytes32)[] _namespacedRowRoots, bytes[] _sharesSets, ((bytes8,bytes8,bytes32)[],uint256,uint256)[] _sharesSetsProofs, bytes8 _namespaceID, (bytes32[],uint256,uint256)[] _rowsProofs, bytes[] _rowRoots, uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _dataRootProof) view returns(bool)
func (_IDAOracle *IDAOracleCaller) VerifySharesInclusion(opts *bind.CallOpts, _namespacedRowRoots []NamespaceNode, _sharesSets [][]byte, _sharesSetsProofs []NamespaceMerkleProof, _namespaceID [8]byte, _rowsProofs []BinaryMerkleProof, _rowRoots [][]byte, _tupleRootNonce *big.Int, _tuple DataRootTuple, _dataRootProof BinaryMerkleProof) (bool, error) {
	var out []interface{}
	err := _IDAOracle.contract.Call(opts, &out, "verifySharesInclusion", _namespacedRowRoots, _sharesSets, _sharesSetsProofs, _namespaceID, _rowsProofs, _rowRoots, _tupleRootNonce, _tuple, _dataRootProof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySharesInclusion is a free data retrieval call binding the contract method 0xa93527b3.
//
// Solidity: function verifySharesInclusion((bytes8,bytes8,bytes32)[] _namespacedRowRoots, bytes[] _sharesSets, ((bytes8,bytes8,bytes32)[],uint256,uint256)[] _sharesSetsProofs, bytes8 _namespaceID, (bytes32[],uint256,uint256)[] _rowsProofs, bytes[] _rowRoots, uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _dataRootProof) view returns(bool)
func (_IDAOracle *IDAOracleSession) VerifySharesInclusion(_namespacedRowRoots []NamespaceNode, _sharesSets [][]byte, _sharesSetsProofs []NamespaceMerkleProof, _namespaceID [8]byte, _rowsProofs []BinaryMerkleProof, _rowRoots [][]byte, _tupleRootNonce *big.Int, _tuple DataRootTuple, _dataRootProof BinaryMerkleProof) (bool, error) {
	return _IDAOracle.Contract.VerifySharesInclusion(&_IDAOracle.CallOpts, _namespacedRowRoots, _sharesSets, _sharesSetsProofs, _namespaceID, _rowsProofs, _rowRoots, _tupleRootNonce, _tuple, _dataRootProof)
}

// VerifySharesInclusion is a free data retrieval call binding the contract method 0xa93527b3.
//
// Solidity: function verifySharesInclusion((bytes8,bytes8,bytes32)[] _namespacedRowRoots, bytes[] _sharesSets, ((bytes8,bytes8,bytes32)[],uint256,uint256)[] _sharesSetsProofs, bytes8 _namespaceID, (bytes32[],uint256,uint256)[] _rowsProofs, bytes[] _rowRoots, uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _dataRootProof) view returns(bool)
func (_IDAOracle *IDAOracleCallerSession) VerifySharesInclusion(_namespacedRowRoots []NamespaceNode, _sharesSets [][]byte, _sharesSetsProofs []NamespaceMerkleProof, _namespaceID [8]byte, _rowsProofs []BinaryMerkleProof, _rowRoots [][]byte, _tupleRootNonce *big.Int, _tuple DataRootTuple, _dataRootProof BinaryMerkleProof) (bool, error) {
	return _IDAOracle.Contract.VerifySharesInclusion(&_IDAOracle.CallOpts, _namespacedRowRoots, _sharesSets, _sharesSetsProofs, _namespaceID, _rowsProofs, _rowRoots, _tupleRootNonce, _tuple, _dataRootProof)
}

// NamespaceMerkleTreeMetaData contains all meta data concerning the NamespaceMerkleTree contract.
var NamespaceMerkleTreeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes8\",\"name\":\"min\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"max\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode\",\"name\":\"root\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes8\",\"name\":\"min\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"max\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"sideNodes\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structNamespaceMerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes8\",\"name\":\"minmaxNID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e8c25a15": "verify(NamespaceNode,NamespaceMerkleProof,bytes8,bytes)",
	},
	Bin: "0x610af761003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063e8c25a151461003a575b600080fd5b61004d610048366004610854565b610061565b604051901515815260200160405180910390f35b60008061006e8484610089565b905061007d868683600161014d565b9150505b949350505050565b604080516060810182526000808252602082018190529181019190915260006002600060f81b8586866040516020016100c594939291906109c2565b60408051601f19818403018152908290526100df91610a00565b602060405180830381855afa1580156100fc573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061011f9190610a0c565b604080516060810182526001600160c01b031987168082526020820152908101919091529150505b92915050565b6000600182101561016057506000610081565b600061016d600184610a3b565b90506001856040015111610192578451511561018d576000915050610081565b6101c1565b6101a485602001518660400151610433565b8551516101b2908390610a52565b146101c1576000915050610081565b84604001518560200151106101da576000915050610081565b84515161020a57846040015160011415610200576101f886856104ba565b915050610081565b6000915050610081565b602085015183905b60208701516000906001841b9061022a908290610a6a565b6102349190610a8c565b90506000600161024681861b84610a52565b6102509190610a3b565b905088604001518110610264575050610356565b915081610272600185610a3b565b895151610280908790610a52565b1161029357600095505050505050610081565b600161029f8686610a3b565b6102a99190610a3b565b6001901b828a602001516102bd9190610a3b565b10156103055788516102fe90899060016102d78989610a3b565b6102e19190610a3b565b815181106102f1576102f1610aab565b6020026020010151610511565b9750610342565b885161033f9060016103178888610a3b565b6103219190610a3b565b8151811061033157610331610aab565b602002602001015189610511565b97505b61034d600185610a52565b93505050610212565b600187604001516103679190610a3b565b81146103b057610378600183610a3b565b8751511161038c5760009350505050610081565b86516103a090879060016102d78787610a3b565b95506103ad600183610a52565b91505b86515160016103bf8585610a3b565b6103c99190610a3b565b101561041d5786516104099060016103e18686610a3b565b6103eb9190610a3b565b815181106103fb576103fb610aab565b602002602001015187610511565b9550610416600183610a52565b91506103b0565b61042788876104ba565b98975050505050505050565b600061043e8261068b565b61044a90610100610a3b565b90506000610459600183610a3b565b6001901b905060018161046c9190610a3b565b84116104785750610147565b806001141561048b576001915050610147565b6104a76104988286610a3b565b6104a28386610a3b565b610433565b6104b2906001610a52565b915050610147565b805182516000916001600160c01b031991821691161480156104f7575081602001516001600160c01b03191683602001516001600160c01b031916145b801561050a575081604001518360400151145b9392505050565b60408051606081018252600080825260208201819052918101919091528251825160009161053e916106b8565b84519091506000906001600160c01b0319908116141561056757506001600160c01b031961059b565b83516001600160c01b031990811614156105865750602084015161059b565b610598856020015185602001516106d7565b90505b84516020808701516040808901518851898501518a8401519351600160f81b968101969096526001600160c01b03199687166021870152938616602986015260318501919091528416605184015292166059820152606181019190915260009060029060810160408051601f198184030181529082905261061b91610a00565b602060405180830381855afa158015610638573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061065b9190610a0c565b604080516060810182526001600160c01b0319958616815293909416602084015292820192909252949350505050565b60005b81816001901b10156106ac576106a5600182610a52565b905061068e565b61014781610100610a3b565b600060c082811c9084901c10156106d0575081610147565b5080610147565b600060c082811c9084901c11156106d0575081610147565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715610728576107286106ef565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610757576107576106ef565b604052919050565b80356001600160c01b03198116811461077757600080fd5b919050565b60006060828403121561078e57600080fd5b6040516060810181811067ffffffffffffffff821117156107b1576107b16106ef565b6040529050806107c08361075f565b81526107ce6020840161075f565b6020820152604083013560408201525092915050565b600082601f8301126107f557600080fd5b813567ffffffffffffffff81111561080f5761080f6106ef565b610822601f8201601f191660200161072e565b81815284602083860101111561083757600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060c0858703121561086a57600080fd5b610874868661077c565b935060608086013567ffffffffffffffff8082111561089257600080fd5b818801915082828a0312156108a657600080fd5b6108ae610705565b8235828111156108bd57600080fd5b8301601f81018b136108ce57600080fd5b80356020848211156108e2576108e26106ef565b6108f0818360051b0161072e565b8281529187028301810191818101908e84111561090c57600080fd5b938201935b83851015610932576109238f8661077c565b82529388019390820190610911565b855250858101359084015250506040808401359082015295506109576080890161075f565b945060a088013592508083111561096d57600080fd5b505061097b878288016107e4565b91505092959194509250565b6000815160005b818110156109a8576020818501810151868301520161098e565b818111156109b7576000828601525b509290920192915050565b6001600160f81b0319851681526001600160c01b031984811660018301528316600982015260006109f66011830184610987565b9695505050505050565b600061050a8284610987565b600060208284031215610a1e57600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b600082821015610a4d57610a4d610a25565b500390565b60008219821115610a6557610a65610a25565b500190565b600082610a8757634e487b7160e01b600052601260045260246000fd5b500490565b6000816000190483118215151615610aa657610aa6610a25565b500290565b634e487b7160e01b600052603260045260246000fdfea2646970667358221220f499939139bfffb2edd7c3f064da78ec9f0d359d075100685e0ede502035b9bc64736f6c63430008090033",
}

// NamespaceMerkleTreeABI is the input ABI used to generate the binding from.
// Deprecated: Use NamespaceMerkleTreeMetaData.ABI instead.
var NamespaceMerkleTreeABI = NamespaceMerkleTreeMetaData.ABI

// Deprecated: Use NamespaceMerkleTreeMetaData.Sigs instead.
// NamespaceMerkleTreeFuncSigs maps the 4-byte function signature to its string representation.
var NamespaceMerkleTreeFuncSigs = NamespaceMerkleTreeMetaData.Sigs

// NamespaceMerkleTreeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NamespaceMerkleTreeMetaData.Bin instead.
var NamespaceMerkleTreeBin = NamespaceMerkleTreeMetaData.Bin

// DeployNamespaceMerkleTree deploys a new Ethereum contract, binding an instance of NamespaceMerkleTree to it.
func DeployNamespaceMerkleTree(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NamespaceMerkleTree, error) {
	parsed, err := NamespaceMerkleTreeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NamespaceMerkleTreeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NamespaceMerkleTree{NamespaceMerkleTreeCaller: NamespaceMerkleTreeCaller{contract: contract}, NamespaceMerkleTreeTransactor: NamespaceMerkleTreeTransactor{contract: contract}, NamespaceMerkleTreeFilterer: NamespaceMerkleTreeFilterer{contract: contract}}, nil
}

// NamespaceMerkleTree is an auto generated Go binding around an Ethereum contract.
type NamespaceMerkleTree struct {
	NamespaceMerkleTreeCaller     // Read-only binding to the contract
	NamespaceMerkleTreeTransactor // Write-only binding to the contract
	NamespaceMerkleTreeFilterer   // Log filterer for contract events
}

// NamespaceMerkleTreeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NamespaceMerkleTreeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NamespaceMerkleTreeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NamespaceMerkleTreeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NamespaceMerkleTreeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NamespaceMerkleTreeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NamespaceMerkleTreeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NamespaceMerkleTreeSession struct {
	Contract     *NamespaceMerkleTree // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NamespaceMerkleTreeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NamespaceMerkleTreeCallerSession struct {
	Contract *NamespaceMerkleTreeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// NamespaceMerkleTreeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NamespaceMerkleTreeTransactorSession struct {
	Contract     *NamespaceMerkleTreeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// NamespaceMerkleTreeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NamespaceMerkleTreeRaw struct {
	Contract *NamespaceMerkleTree // Generic contract binding to access the raw methods on
}

// NamespaceMerkleTreeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NamespaceMerkleTreeCallerRaw struct {
	Contract *NamespaceMerkleTreeCaller // Generic read-only contract binding to access the raw methods on
}

// NamespaceMerkleTreeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NamespaceMerkleTreeTransactorRaw struct {
	Contract *NamespaceMerkleTreeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNamespaceMerkleTree creates a new instance of NamespaceMerkleTree, bound to a specific deployed contract.
func NewNamespaceMerkleTree(address common.Address, backend bind.ContractBackend) (*NamespaceMerkleTree, error) {
	contract, err := bindNamespaceMerkleTree(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NamespaceMerkleTree{NamespaceMerkleTreeCaller: NamespaceMerkleTreeCaller{contract: contract}, NamespaceMerkleTreeTransactor: NamespaceMerkleTreeTransactor{contract: contract}, NamespaceMerkleTreeFilterer: NamespaceMerkleTreeFilterer{contract: contract}}, nil
}

// NewNamespaceMerkleTreeCaller creates a new read-only instance of NamespaceMerkleTree, bound to a specific deployed contract.
func NewNamespaceMerkleTreeCaller(address common.Address, caller bind.ContractCaller) (*NamespaceMerkleTreeCaller, error) {
	contract, err := bindNamespaceMerkleTree(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NamespaceMerkleTreeCaller{contract: contract}, nil
}

// NewNamespaceMerkleTreeTransactor creates a new write-only instance of NamespaceMerkleTree, bound to a specific deployed contract.
func NewNamespaceMerkleTreeTransactor(address common.Address, transactor bind.ContractTransactor) (*NamespaceMerkleTreeTransactor, error) {
	contract, err := bindNamespaceMerkleTree(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NamespaceMerkleTreeTransactor{contract: contract}, nil
}

// NewNamespaceMerkleTreeFilterer creates a new log filterer instance of NamespaceMerkleTree, bound to a specific deployed contract.
func NewNamespaceMerkleTreeFilterer(address common.Address, filterer bind.ContractFilterer) (*NamespaceMerkleTreeFilterer, error) {
	contract, err := bindNamespaceMerkleTree(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NamespaceMerkleTreeFilterer{contract: contract}, nil
}

// bindNamespaceMerkleTree binds a generic wrapper to an already deployed contract.
func bindNamespaceMerkleTree(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NamespaceMerkleTreeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NamespaceMerkleTree *NamespaceMerkleTreeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NamespaceMerkleTree.Contract.NamespaceMerkleTreeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NamespaceMerkleTree *NamespaceMerkleTreeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NamespaceMerkleTree.Contract.NamespaceMerkleTreeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NamespaceMerkleTree *NamespaceMerkleTreeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NamespaceMerkleTree.Contract.NamespaceMerkleTreeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NamespaceMerkleTree *NamespaceMerkleTreeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NamespaceMerkleTree.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NamespaceMerkleTree *NamespaceMerkleTreeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NamespaceMerkleTree.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NamespaceMerkleTree *NamespaceMerkleTreeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NamespaceMerkleTree.Contract.contract.Transact(opts, method, params...)
}

// Verify is a free data retrieval call binding the contract method 0xfba69670.
//
// Solidity: function verify((bytes8,bytes8,bytes32) root, ((bytes8,bytes8,bytes32)[],uint256,uint256) proof, bytes8 minmaxNID, bytes data) pure returns(bool)
func (_NamespaceMerkleTree *NamespaceMerkleTreeCaller) Verify(opts *bind.CallOpts, root NamespaceNode, proof NamespaceMerkleProof, minmaxNID [8]byte, data []byte) (bool, error) {
	var out []interface{}
	err := _NamespaceMerkleTree.contract.Call(opts, &out, "verify", root, proof, minmaxNID, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xfba69670.
//
// Solidity: function verify((bytes8,bytes8,bytes32) root, ((bytes8,bytes8,bytes32)[],uint256,uint256) proof, bytes8 minmaxNID, bytes data) pure returns(bool)
func (_NamespaceMerkleTree *NamespaceMerkleTreeSession) Verify(root NamespaceNode, proof NamespaceMerkleProof, minmaxNID [8]byte, data []byte) (bool, error) {
	return _NamespaceMerkleTree.Contract.Verify(&_NamespaceMerkleTree.CallOpts, root, proof, minmaxNID, data)
}

// Verify is a free data retrieval call binding the contract method 0xfba69670.
//
// Solidity: function verify((bytes8,bytes8,bytes32) root, ((bytes8,bytes8,bytes32)[],uint256,uint256) proof, bytes8 minmaxNID, bytes data) pure returns(bool)
func (_NamespaceMerkleTree *NamespaceMerkleTreeCallerSession) Verify(root NamespaceNode, proof NamespaceMerkleProof, minmaxNID [8]byte, data []byte) (bool, error) {
	return _NamespaceMerkleTree.Contract.Verify(&_NamespaceMerkleTree.CallOpts, root, proof, minmaxNID, data)
}

// QuantumGravityBridgeMetaData contains all meta data concerning the QuantumGravityBridge contract.
var QuantumGravityBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bridge_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_powerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_validatorSetHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientVotingPower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDataRootTupleRootNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValidatorSetNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MalformedCurrentValidatorSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SuppliedValidatorSetInvalid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataRootTupleRoot\",\"type\":\"bytes32\"}],\"name\":\"DataRootTupleRootEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"powerThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"validatorSetHash\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetUpdatedEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"state_dataRootTupleRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_eventNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValidatorSetCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_powerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorSetNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_dataRootTupleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"_currentValidatorSet\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_sigs\",\"type\":\"tuple[]\"}],\"name\":\"submitDataRootTupleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_oldNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newPowerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_newValidatorSetHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"_currentValidatorSet\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_sigs\",\"type\":\"tuple[]\"}],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"_tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes8\",\"name\":\"min\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"max\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"_namespacedRowRoots\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_sharesSets\",\"type\":\"bytes[]\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes8\",\"name\":\"min\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"max\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"sideNodes\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structNamespaceMerkleProof[]\",\"name\":\"_sharesSetsProofs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes8\",\"name\":\"_namespaceID\",\"type\":\"bytes8\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof[]\",\"name\":\"_rowsProofs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_rowRoots\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"_tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"_tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"_dataRootProof\",\"type\":\"tuple\"}],\"name\":\"verifySharesInclusion\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"af4482af": "BRIDGE_ID()",
		"817f985b": "state_dataRootTupleRoots(uint256)",
		"cdade866": "state_eventNonce()",
		"5433218c": "state_lastValidatorSetCheckpoint()",
		"e5a2b5d2": "state_powerThreshold()",
		"e23eb326": "submitDataRootTupleRoot(uint256,uint256,bytes32,(address,uint256)[],(uint8,bytes32,bytes32)[])",
		"05d85c13": "updateValidatorSet(uint256,uint256,uint256,bytes32,(address,uint256)[],(uint8,bytes32,bytes32)[])",
		"1f3302a9": "verifyAttestation(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))",
		"a93527b3": "verifySharesInclusion((bytes8,bytes8,bytes32)[],bytes[],((bytes8,bytes8,bytes32)[],uint256,uint256)[],bytes8,(bytes32[],uint256,uint256)[],bytes[],uint256,(uint256,bytes32),(bytes32[],uint256,uint256))",
	},
	Bin: "0x60a060405234801561001057600080fd5b50604051611c70380380611c7083398101604081905261002f916100d3565b60808481526040805160208082018890526918da1958dadc1bda5b9d60b21b828401526060820187905292810185905260a08082018590528251808303909101815260c08201808452815191909401206002879055600081905560018690559285905260e08101849052905185917fe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c91908190036101000190a25050505050610109565b600080600080608085870312156100e957600080fd5b505082516020840151604085015160609095015191969095509092509050565b608051611b3061014060003960008181610124015281816101d30152818161021f0152818161056f01526105ba0152611b306000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063a93527b311610066578063a93527b31461010c578063af4482af1461011f578063cdade86614610146578063e23eb3261461014f578063e5a2b5d21461016257600080fd5b806305d85c13146100985780631f3302a9146100ad5780635433218c146100d5578063817f985b146100ec575b600080fd5b6100ab6100a636600461106a565b61016b565b005b6100c06100bb366004611293565b6102ae565b60405190151581526020015b60405180910390f35b6100de60005481565b6040519081526020016100cc565b6100de6100fa3660046112ea565b60036020526000908152604090205481565b6100c061011a3660046115f9565b610310565b6100de7f000000000000000000000000000000000000000000000000000000000000000081565b6100de60025481565b6100ab61015d36600461171b565b610507565b6100de60015481565b600254600180549061017e9083906117bd565b8a1461019d576040516368a35ffd60e11b815260040160405180910390fd5b8483146101bd5760405163c6617b7b60e01b815260040160405180910390fd5b60006101c9878761068d565b90506000546101fa7f00000000000000000000000000000000000000000000000000000000000000008c85856106c1565b1461021857604051630bbdaec960e11b815260040160405180910390fd5b60006102467f00000000000000000000000000000000000000000000000000000000000000008d8c8c6106c1565b9050610256888888888588610712565b600081905560018a905560028c9055604080518b8152602081018b90528d917fe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c910160405180910390a2505050505050505050505050565b60006002548411156102c257506000610309565b60008481526003602090815260408083205481518751818501529287015183830152815180840383018152606090930190915291906103049083908690610820565b925050505b9392505050565b60008951865114610323575060006104fa565b8951885114610334575060006104fa565b8951855114610345575060006104fa565b6000805b89518163ffffffff1610156104665773__$549b49878917c9869da759c74e6164f91f$__63e8c25a158d8363ffffffff168151811061038a5761038a6117d5565b60200260200101518c8463ffffffff16815181106103aa576103aa6117d5565b60200260200101518c8f8663ffffffff16815181106103cb576103cb6117d5565b60200260200101516040518563ffffffff1660e01b81526004016103f29493929190611847565b60206040518083038186803b15801561040a57600080fd5b505af415801561041e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104429190611926565b915081610454576000925050506104fa565b8061045e81611948565b915050610349565b5060005b87518163ffffffff1610156104ea576104c68560200151898363ffffffff1681518110610499576104996117d5565b6020026020010151898463ffffffff16815181106104b9576104b96117d5565b6020026020010151610820565b9150816104d8576000925050506104fa565b806104e281611948565b91505061046a565b506104f68585856102ae565b9150505b9998505050505050505050565b600254600180549061051a9083906117bd565b89146105395760405163e869766d60e01b815260040160405180910390fd5b8483146105595760405163c6617b7b60e01b815260040160405180910390fd5b6000610565878761068d565b90506000546105967f00000000000000000000000000000000000000000000000000000000000000008b85856106c1565b146105b457604051630bbdaec960e11b815260040160405180910390fd5b604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091526f0e8e4c2dce6c2c6e8d2dedc84c2e8c6d60831b82840152606082018d905260808083018c90528351808403909101815260a0909201909252805191012061062e888888888588610712565b60028b905560008b815260036020526040908190208a9055518b907f6614d037bde4905e31ca5ff05de61964c267f28b0320ed49e59f7d99752e1c4f90610678908c815260200190565b60405180910390a25050505050505050505050565b600082826040516020016106a2929190611983565b6040516020818303038152906040528051906020012090505b92915050565b604080516020808201969096526918da1958dadc1bda5b9d60b21b818301526060810194909452608084019290925260a0808401919091528151808403909101815260c09092019052805191012090565b6000805b868110156107f55761073e868683818110610733576107336117d5565b905060600201610a8e565b15610748576107e3565b61079188888381811061075d5761075d6117d5565b61077392602060409092020190810191506119cd565b85888885818110610786576107866117d5565b905060600201610ac2565b6107ae57604051638baa579f60e01b815260040160405180910390fd5b8787828181106107c0576107c06117d5565b90506040020160200135826107d591906117bd565b91508282106107e3576107f5565b806107ed816119e8565b915050610716565b50818110156108175760405163cabeb65560e01b815260040160405180910390fd5b50505050505050565b60006001836040015111610843578251511561083e57506000610309565b610865565b61085583602001518460400151610b5c565b8351511461086557506000610309565b826040015183602001511061087c57506000610309565b600061088783610be3565b8451519091506108b1578360400151600114156108a75784149050610309565b6000915050610309565b60208401516001905b60208601516000906001841b906108d2908290611a03565b6108dc9190611a25565b9050600060016108ee81861b846117bd565b6108f89190611a44565b90508760400151811061090c5750506109d2565b91508161091a600185611a44565b8851511161093057600095505050505050610309565b61093b600185611a44565b6001901b82896020015161094f9190611a44565b101561098c578751610985908690610968600188611a44565b81518110610978576109786117d5565b6020026020010151610c58565b94506109be565b87516109bb9061099d600187611a44565b815181106109ad576109ad6117d5565b602002602001015186610c58565b94505b6109c96001856117bd565b935050506108ba565b600186604001516109e39190611a44565b8114610a2b576109f4600183611a44565b86515111610a085760009350505050610309565b8551610a1b908490610968600186611a44565b9250610a286001836117bd565b91505b855151610a39600184611a44565b1015610a82578551610a6e90610a50600185611a44565b81518110610a6057610a606117d5565b602002602001015184610c58565b9250610a7b6001836117bd565b9150610a2b565b50509093149392505050565b60006020820135158015610aa457506040820135155b80156106bb5750610ab86020830183611a5b565b60ff161592915050565b600080610b1c846040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b9050610b3e81610b2f6020860186611a5b565b85602001358660400135610cd6565b6001600160a01b0316856001600160a01b0316149150509392505050565b6000610b6782610cfe565b610b7390610100611a44565b90506000610b82600183611a44565b6001901b9050600181610b959190611a44565b8411610ba157506106bb565b8060011415610bb45760019150506106bb565b610bd0610bc18286611a44565b610bcb8386611a44565b610b5c565b610bdb9060016117bd565b9150506106bb565b60006002600060f81b83604051602001610bfe929190611a7e565b60408051601f1981840301815290829052610c1891611aaf565b602060405180830381855afa158015610c35573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906106bb9190611acb565b604051600160f81b6020820152602181018390526041810182905260009060029060610160408051601f1981840301815290829052610c9691611aaf565b602060405180830381855afa158015610cb3573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906103099190611acb565b6000806000610ce787878787610d2b565b91509150610cf481610e18565b5095945050505050565b60005b81816001901b1015610d1f57610d186001826117bd565b9050610d01565b6106bb81610100611a44565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610d625750600090506003610e0f565b8460ff16601b14158015610d7a57508460ff16601c14155b15610d8b5750600090506004610e0f565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610ddf573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610e0857600060019250925050610e0f565b9150600090505b94509492505050565b6000816004811115610e2c57610e2c611ae4565b1415610e355750565b6001816004811115610e4957610e49611ae4565b1415610e9c5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064015b60405180910390fd5b6002816004811115610eb057610eb0611ae4565b1415610efe5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610e93565b6003816004811115610f1257610f12611ae4565b1415610f6b5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610e93565b6004816004811115610f7f57610f7f611ae4565b1415610fd85760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610e93565b50565b60008083601f840112610fed57600080fd5b5081356001600160401b0381111561100457600080fd5b6020830191508360208260061b850101111561101f57600080fd5b9250929050565b60008083601f84011261103857600080fd5b5081356001600160401b0381111561104f57600080fd5b60208301915083602060608302850101111561101f57600080fd5b60008060008060008060008060c0898b03121561108657600080fd5b8835975060208901359650604089013595506060890135945060808901356001600160401b03808211156110b957600080fd5b6110c58c838d01610fdb565b909650945060a08b01359150808211156110de57600080fd5b506110eb8b828c01611026565b999c989b5096995094979396929594505050565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715611137576111376110ff565b60405290565b604051601f8201601f191681016001600160401b0381118282101715611165576111656110ff565b604052919050565b60006040828403121561117f57600080fd5b604051604081018181106001600160401b03821117156111a1576111a16110ff565b604052823581526020928301359281019290925250919050565b60006001600160401b038211156111d4576111d46110ff565b5060051b60200190565b6000606082840312156111f057600080fd5b6111f8611115565b905081356001600160401b0381111561121057600080fd5b8201601f8101841361122157600080fd5b80356020611236611231836111bb565b61113d565b82815260059290921b8301810191818101908784111561125557600080fd5b938201935b838510156112735784358252938201939082019061125a565b808652505080850135818501525050506040820135604082015292915050565b6000806000608084860312156112a857600080fd5b833592506112b9856020860161116d565b915060608401356001600160401b038111156112d457600080fd5b6112e0868287016111de565b9150509250925092565b6000602082840312156112fc57600080fd5b5035919050565b80356001600160c01b03198116811461131b57600080fd5b919050565b600082601f83011261133157600080fd5b81356020611341611231836111bb565b8281526060928302850182019282820191908785111561136057600080fd5b8387015b858110156113b55781818a03121561137c5760008081fd5b611384611115565b61138d82611303565b815261139a868301611303565b81870152604082810135908201528452928401928101611364565b5090979650505050505050565b6000601f83818401126113d457600080fd5b823560206113e4611231836111bb565b82815260059290921b8501810191818101908784111561140357600080fd5b8287015b848110156114995780356001600160401b03808211156114275760008081fd5b818a0191508a603f83011261143c5760008081fd5b85820135604082821115611452576114526110ff565b611463828b01601f1916890161113d565b92508183528c8183860101111561147a5760008081fd5b8181850189850137506000908201870152845250918301918301611407565b50979650505050505050565b600082601f8301126114b657600080fd5b813560206114c6611231836111bb565b82815260059290921b840181019181810190868411156114e557600080fd5b8286015b8481101561156f5780356001600160401b03808211156115095760008081fd5b908801906060828b03601f19018113156115235760008081fd5b61152b611115565b878401358381111561153d5760008081fd5b61154b8d8a83880101611320565b825250604084810135898301529190930135908301525083529183019183016114e9565b509695505050505050565b600082601f83011261158b57600080fd5b8135602061159b611231836111bb565b82815260059290921b840181019181810190868411156115ba57600080fd5b8286015b8481101561156f5780356001600160401b038111156115dd5760008081fd5b6115eb8986838b01016111de565b8452509183019183016115be565b60008060008060008060008060006101408a8c03121561161857600080fd5b89356001600160401b038082111561162f57600080fd5b61163b8d838e01611320565b9a5060208c013591508082111561165157600080fd5b61165d8d838e016113c2565b995060408c013591508082111561167357600080fd5b61167f8d838e016114a5565b985061168d60608d01611303565b975060808c01359150808211156116a357600080fd5b6116af8d838e0161157a565b965060a08c01359150808211156116c557600080fd5b6116d18d838e016113c2565b955060c08c013594506116e78d60e08e0161116d565b93506101208c01359150808211156116fe57600080fd5b5061170b8c828d016111de565b9150509295985092959850929598565b600080600080600080600060a0888a03121561173657600080fd5b87359650602088013595506040880135945060608801356001600160401b038082111561176257600080fd5b61176e8b838c01610fdb565b909650945060808a013591508082111561178757600080fd5b506117948a828b01611026565b989b979a50959850939692959293505050565b634e487b7160e01b600052601160045260246000fd5b600082198211156117d0576117d06117a7565b500190565b634e487b7160e01b600052603260045260246000fd5b60005b838110156118065781810151838201526020016117ee565b83811115611815576000848401525b50505050565b600081518084526118338160208601602086016117eb565b601f01601f19169290920160200192915050565b84516001600160c01b03199081168252602080870151909116818301526040808701519083015260c060608084018290528651918401819052815161012085018190526000939192820190610140860190855b818110156118e0576118d083855180516001600160c01b0319908116835260208083015190911690830152604090810151910152565b928401929185019160010161189a565b50508883015160e087015260408901516101008701526001600160c01b03198816608087015285810360a0870152611918818861181b565b9a9950505050505050505050565b60006020828403121561193857600080fd5b8151801515811461030957600080fd5b600063ffffffff80831681811415611962576119626117a7565b6001019392505050565b80356001600160a01b038116811461131b57600080fd5b6020808252818101839052600090604080840186845b878110156113b5576001600160a01b036119b28361196c565b16835281850135858401529183019190830190600101611999565b6000602082840312156119df57600080fd5b6103098261196c565b60006000198214156119fc576119fc6117a7565b5060010190565b600082611a2057634e487b7160e01b600052601260045260246000fd5b500490565b6000816000190483118215151615611a3f57611a3f6117a7565b500290565b600082821015611a5657611a566117a7565b500390565b600060208284031215611a6d57600080fd5b813560ff8116811461030957600080fd5b6001600160f81b0319831681528151600090611aa18160018501602087016117eb565b919091016001019392505050565b60008251611ac18184602087016117eb565b9190910192915050565b600060208284031215611add57600080fd5b5051919050565b634e487b7160e01b600052602160045260246000fdfea26469706673582212204175166c54e4999bf071dee2dcd812a81d5f2d194cc228e9d1b033e4084d5aae64736f6c63430008090033",
}

// QuantumGravityBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use QuantumGravityBridgeMetaData.ABI instead.
var QuantumGravityBridgeABI = QuantumGravityBridgeMetaData.ABI

// Deprecated: Use QuantumGravityBridgeMetaData.Sigs instead.
// QuantumGravityBridgeFuncSigs maps the 4-byte function signature to its string representation.
var QuantumGravityBridgeFuncSigs = QuantumGravityBridgeMetaData.Sigs

// QuantumGravityBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use QuantumGravityBridgeMetaData.Bin instead.
var QuantumGravityBridgeBin = QuantumGravityBridgeMetaData.Bin

// DeployQuantumGravityBridge deploys a new Ethereum contract, binding an instance of QuantumGravityBridge to it.
func DeployQuantumGravityBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge_id [32]byte, _nonce *big.Int, _powerThreshold *big.Int, _validatorSetHash [32]byte) (common.Address, *types.Transaction, *QuantumGravityBridge, error) {
	parsed, err := QuantumGravityBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	namespaceMerkleTreeAddr, _, _, _ := DeployNamespaceMerkleTree(auth, backend)
	QuantumGravityBridgeBin = strings.Replace(QuantumGravityBridgeBin, "__$549b49878917c9869da759c74e6164f91f$__", namespaceMerkleTreeAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(QuantumGravityBridgeBin), backend, _bridge_id, _nonce, _powerThreshold, _validatorSetHash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &QuantumGravityBridge{QuantumGravityBridgeCaller: QuantumGravityBridgeCaller{contract: contract}, QuantumGravityBridgeTransactor: QuantumGravityBridgeTransactor{contract: contract}, QuantumGravityBridgeFilterer: QuantumGravityBridgeFilterer{contract: contract}}, nil
}

// QuantumGravityBridge is an auto generated Go binding around an Ethereum contract.
type QuantumGravityBridge struct {
	QuantumGravityBridgeCaller     // Read-only binding to the contract
	QuantumGravityBridgeTransactor // Write-only binding to the contract
	QuantumGravityBridgeFilterer   // Log filterer for contract events
}

// QuantumGravityBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuantumGravityBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuantumGravityBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuantumGravityBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuantumGravityBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuantumGravityBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuantumGravityBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuantumGravityBridgeSession struct {
	Contract     *QuantumGravityBridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// QuantumGravityBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuantumGravityBridgeCallerSession struct {
	Contract *QuantumGravityBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// QuantumGravityBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuantumGravityBridgeTransactorSession struct {
	Contract     *QuantumGravityBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// QuantumGravityBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuantumGravityBridgeRaw struct {
	Contract *QuantumGravityBridge // Generic contract binding to access the raw methods on
}

// QuantumGravityBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuantumGravityBridgeCallerRaw struct {
	Contract *QuantumGravityBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// QuantumGravityBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuantumGravityBridgeTransactorRaw struct {
	Contract *QuantumGravityBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuantumGravityBridge creates a new instance of QuantumGravityBridge, bound to a specific deployed contract.
func NewQuantumGravityBridge(address common.Address, backend bind.ContractBackend) (*QuantumGravityBridge, error) {
	contract, err := bindQuantumGravityBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QuantumGravityBridge{QuantumGravityBridgeCaller: QuantumGravityBridgeCaller{contract: contract}, QuantumGravityBridgeTransactor: QuantumGravityBridgeTransactor{contract: contract}, QuantumGravityBridgeFilterer: QuantumGravityBridgeFilterer{contract: contract}}, nil
}

// NewQuantumGravityBridgeCaller creates a new read-only instance of QuantumGravityBridge, bound to a specific deployed contract.
func NewQuantumGravityBridgeCaller(address common.Address, caller bind.ContractCaller) (*QuantumGravityBridgeCaller, error) {
	contract, err := bindQuantumGravityBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuantumGravityBridgeCaller{contract: contract}, nil
}

// NewQuantumGravityBridgeTransactor creates a new write-only instance of QuantumGravityBridge, bound to a specific deployed contract.
func NewQuantumGravityBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*QuantumGravityBridgeTransactor, error) {
	contract, err := bindQuantumGravityBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuantumGravityBridgeTransactor{contract: contract}, nil
}

// NewQuantumGravityBridgeFilterer creates a new log filterer instance of QuantumGravityBridge, bound to a specific deployed contract.
func NewQuantumGravityBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*QuantumGravityBridgeFilterer, error) {
	contract, err := bindQuantumGravityBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuantumGravityBridgeFilterer{contract: contract}, nil
}

// bindQuantumGravityBridge binds a generic wrapper to an already deployed contract.
func bindQuantumGravityBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QuantumGravityBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuantumGravityBridge *QuantumGravityBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuantumGravityBridge.Contract.QuantumGravityBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuantumGravityBridge *QuantumGravityBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.QuantumGravityBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuantumGravityBridge *QuantumGravityBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.QuantumGravityBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuantumGravityBridge *QuantumGravityBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuantumGravityBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuantumGravityBridge *QuantumGravityBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuantumGravityBridge *QuantumGravityBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.contract.Transact(opts, method, params...)
}

// BRIDGEID is a free data retrieval call binding the contract method 0xaf4482af.
//
// Solidity: function BRIDGE_ID() view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) BRIDGEID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "BRIDGE_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BRIDGEID is a free data retrieval call binding the contract method 0xaf4482af.
//
// Solidity: function BRIDGE_ID() view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) BRIDGEID() ([32]byte, error) {
	return _QuantumGravityBridge.Contract.BRIDGEID(&_QuantumGravityBridge.CallOpts)
}

// BRIDGEID is a free data retrieval call binding the contract method 0xaf4482af.
//
// Solidity: function BRIDGE_ID() view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) BRIDGEID() ([32]byte, error) {
	return _QuantumGravityBridge.Contract.BRIDGEID(&_QuantumGravityBridge.CallOpts)
}

// StateDataRootTupleRoots is a free data retrieval call binding the contract method 0x817f985b.
//
// Solidity: function state_dataRootTupleRoots(uint256 ) view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) StateDataRootTupleRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "state_dataRootTupleRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateDataRootTupleRoots is a free data retrieval call binding the contract method 0x817f985b.
//
// Solidity: function state_dataRootTupleRoots(uint256 ) view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) StateDataRootTupleRoots(arg0 *big.Int) ([32]byte, error) {
	return _QuantumGravityBridge.Contract.StateDataRootTupleRoots(&_QuantumGravityBridge.CallOpts, arg0)
}

// StateDataRootTupleRoots is a free data retrieval call binding the contract method 0x817f985b.
//
// Solidity: function state_dataRootTupleRoots(uint256 ) view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) StateDataRootTupleRoots(arg0 *big.Int) ([32]byte, error) {
	return _QuantumGravityBridge.Contract.StateDataRootTupleRoots(&_QuantumGravityBridge.CallOpts, arg0)
}

// StateEventNonce is a free data retrieval call binding the contract method 0xcdade866.
//
// Solidity: function state_eventNonce() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) StateEventNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "state_eventNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateEventNonce is a free data retrieval call binding the contract method 0xcdade866.
//
// Solidity: function state_eventNonce() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) StateEventNonce() (*big.Int, error) {
	return _QuantumGravityBridge.Contract.StateEventNonce(&_QuantumGravityBridge.CallOpts)
}

// StateEventNonce is a free data retrieval call binding the contract method 0xcdade866.
//
// Solidity: function state_eventNonce() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) StateEventNonce() (*big.Int, error) {
	return _QuantumGravityBridge.Contract.StateEventNonce(&_QuantumGravityBridge.CallOpts)
}

// StateLastValidatorSetCheckpoint is a free data retrieval call binding the contract method 0x5433218c.
//
// Solidity: function state_lastValidatorSetCheckpoint() view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) StateLastValidatorSetCheckpoint(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "state_lastValidatorSetCheckpoint")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateLastValidatorSetCheckpoint is a free data retrieval call binding the contract method 0x5433218c.
//
// Solidity: function state_lastValidatorSetCheckpoint() view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) StateLastValidatorSetCheckpoint() ([32]byte, error) {
	return _QuantumGravityBridge.Contract.StateLastValidatorSetCheckpoint(&_QuantumGravityBridge.CallOpts)
}

// StateLastValidatorSetCheckpoint is a free data retrieval call binding the contract method 0x5433218c.
//
// Solidity: function state_lastValidatorSetCheckpoint() view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) StateLastValidatorSetCheckpoint() ([32]byte, error) {
	return _QuantumGravityBridge.Contract.StateLastValidatorSetCheckpoint(&_QuantumGravityBridge.CallOpts)
}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) StatePowerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "state_powerThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) StatePowerThreshold() (*big.Int, error) {
	return _QuantumGravityBridge.Contract.StatePowerThreshold(&_QuantumGravityBridge.CallOpts)
}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) StatePowerThreshold() (*big.Int, error) {
	return _QuantumGravityBridge.Contract.StatePowerThreshold(&_QuantumGravityBridge.CallOpts)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) VerifyAttestation(opts *bind.CallOpts, _tupleRootNonce *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "verifyAttestation", _tupleRootNonce, _tuple, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) VerifyAttestation(_tupleRootNonce *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	return _QuantumGravityBridge.Contract.VerifyAttestation(&_QuantumGravityBridge.CallOpts, _tupleRootNonce, _tuple, _proof)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) VerifyAttestation(_tupleRootNonce *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	return _QuantumGravityBridge.Contract.VerifyAttestation(&_QuantumGravityBridge.CallOpts, _tupleRootNonce, _tuple, _proof)
}

// VerifySharesInclusion is a free data retrieval call binding the contract method 0xa93527b3.
//
// Solidity: function verifySharesInclusion((bytes8,bytes8,bytes32)[] _namespacedRowRoots, bytes[] _sharesSets, ((bytes8,bytes8,bytes32)[],uint256,uint256)[] _sharesSetsProofs, bytes8 _namespaceID, (bytes32[],uint256,uint256)[] _rowsProofs, bytes[] _rowRoots, uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _dataRootProof) view returns(bool)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) VerifySharesInclusion(opts *bind.CallOpts, _namespacedRowRoots []NamespaceNode, _sharesSets [][]byte, _sharesSetsProofs []NamespaceMerkleProof, _namespaceID [8]byte, _rowsProofs []BinaryMerkleProof, _rowRoots [][]byte, _tupleRootNonce *big.Int, _tuple DataRootTuple, _dataRootProof BinaryMerkleProof) (bool, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "verifySharesInclusion", _namespacedRowRoots, _sharesSets, _sharesSetsProofs, _namespaceID, _rowsProofs, _rowRoots, _tupleRootNonce, _tuple, _dataRootProof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySharesInclusion is a free data retrieval call binding the contract method 0xa93527b3.
//
// Solidity: function verifySharesInclusion((bytes8,bytes8,bytes32)[] _namespacedRowRoots, bytes[] _sharesSets, ((bytes8,bytes8,bytes32)[],uint256,uint256)[] _sharesSetsProofs, bytes8 _namespaceID, (bytes32[],uint256,uint256)[] _rowsProofs, bytes[] _rowRoots, uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _dataRootProof) view returns(bool)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) VerifySharesInclusion(_namespacedRowRoots []NamespaceNode, _sharesSets [][]byte, _sharesSetsProofs []NamespaceMerkleProof, _namespaceID [8]byte, _rowsProofs []BinaryMerkleProof, _rowRoots [][]byte, _tupleRootNonce *big.Int, _tuple DataRootTuple, _dataRootProof BinaryMerkleProof) (bool, error) {
	return _QuantumGravityBridge.Contract.VerifySharesInclusion(&_QuantumGravityBridge.CallOpts, _namespacedRowRoots, _sharesSets, _sharesSetsProofs, _namespaceID, _rowsProofs, _rowRoots, _tupleRootNonce, _tuple, _dataRootProof)
}

// VerifySharesInclusion is a free data retrieval call binding the contract method 0xa93527b3.
//
// Solidity: function verifySharesInclusion((bytes8,bytes8,bytes32)[] _namespacedRowRoots, bytes[] _sharesSets, ((bytes8,bytes8,bytes32)[],uint256,uint256)[] _sharesSetsProofs, bytes8 _namespaceID, (bytes32[],uint256,uint256)[] _rowsProofs, bytes[] _rowRoots, uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _dataRootProof) view returns(bool)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) VerifySharesInclusion(_namespacedRowRoots []NamespaceNode, _sharesSets [][]byte, _sharesSetsProofs []NamespaceMerkleProof, _namespaceID [8]byte, _rowsProofs []BinaryMerkleProof, _rowRoots [][]byte, _tupleRootNonce *big.Int, _tuple DataRootTuple, _dataRootProof BinaryMerkleProof) (bool, error) {
	return _QuantumGravityBridge.Contract.VerifySharesInclusion(&_QuantumGravityBridge.CallOpts, _namespacedRowRoots, _sharesSets, _sharesSetsProofs, _namespaceID, _rowsProofs, _rowRoots, _tupleRootNonce, _tuple, _dataRootProof)
}

// SubmitDataRootTupleRoot is a paid mutator transaction binding the contract method 0xe23eb326.
//
// Solidity: function submitDataRootTupleRoot(uint256 _newNonce, uint256 _validatorSetNonce, bytes32 _dataRootTupleRoot, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactor) SubmitDataRootTupleRoot(opts *bind.TransactOpts, _newNonce *big.Int, _validatorSetNonce *big.Int, _dataRootTupleRoot [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.contract.Transact(opts, "submitDataRootTupleRoot", _newNonce, _validatorSetNonce, _dataRootTupleRoot, _currentValidatorSet, _sigs)
}

// SubmitDataRootTupleRoot is a paid mutator transaction binding the contract method 0xe23eb326.
//
// Solidity: function submitDataRootTupleRoot(uint256 _newNonce, uint256 _validatorSetNonce, bytes32 _dataRootTupleRoot, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeSession) SubmitDataRootTupleRoot(_newNonce *big.Int, _validatorSetNonce *big.Int, _dataRootTupleRoot [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.SubmitDataRootTupleRoot(&_QuantumGravityBridge.TransactOpts, _newNonce, _validatorSetNonce, _dataRootTupleRoot, _currentValidatorSet, _sigs)
}

// SubmitDataRootTupleRoot is a paid mutator transaction binding the contract method 0xe23eb326.
//
// Solidity: function submitDataRootTupleRoot(uint256 _newNonce, uint256 _validatorSetNonce, bytes32 _dataRootTupleRoot, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactorSession) SubmitDataRootTupleRoot(_newNonce *big.Int, _validatorSetNonce *big.Int, _dataRootTupleRoot [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.SubmitDataRootTupleRoot(&_QuantumGravityBridge.TransactOpts, _newNonce, _validatorSetNonce, _dataRootTupleRoot, _currentValidatorSet, _sigs)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x05d85c13.
//
// Solidity: function updateValidatorSet(uint256 _newNonce, uint256 _oldNonce, uint256 _newPowerThreshold, bytes32 _newValidatorSetHash, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactor) UpdateValidatorSet(opts *bind.TransactOpts, _newNonce *big.Int, _oldNonce *big.Int, _newPowerThreshold *big.Int, _newValidatorSetHash [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.contract.Transact(opts, "updateValidatorSet", _newNonce, _oldNonce, _newPowerThreshold, _newValidatorSetHash, _currentValidatorSet, _sigs)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x05d85c13.
//
// Solidity: function updateValidatorSet(uint256 _newNonce, uint256 _oldNonce, uint256 _newPowerThreshold, bytes32 _newValidatorSetHash, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeSession) UpdateValidatorSet(_newNonce *big.Int, _oldNonce *big.Int, _newPowerThreshold *big.Int, _newValidatorSetHash [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.UpdateValidatorSet(&_QuantumGravityBridge.TransactOpts, _newNonce, _oldNonce, _newPowerThreshold, _newValidatorSetHash, _currentValidatorSet, _sigs)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x05d85c13.
//
// Solidity: function updateValidatorSet(uint256 _newNonce, uint256 _oldNonce, uint256 _newPowerThreshold, bytes32 _newValidatorSetHash, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactorSession) UpdateValidatorSet(_newNonce *big.Int, _oldNonce *big.Int, _newPowerThreshold *big.Int, _newValidatorSetHash [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.UpdateValidatorSet(&_QuantumGravityBridge.TransactOpts, _newNonce, _oldNonce, _newPowerThreshold, _newValidatorSetHash, _currentValidatorSet, _sigs)
}

// QuantumGravityBridgeDataRootTupleRootEventIterator is returned from FilterDataRootTupleRootEvent and is used to iterate over the raw logs and unpacked data for DataRootTupleRootEvent events raised by the QuantumGravityBridge contract.
type QuantumGravityBridgeDataRootTupleRootEventIterator struct {
	Event *QuantumGravityBridgeDataRootTupleRootEvent // Event containing the contract specifics and raw log

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
func (it *QuantumGravityBridgeDataRootTupleRootEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuantumGravityBridgeDataRootTupleRootEvent)
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
		it.Event = new(QuantumGravityBridgeDataRootTupleRootEvent)
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
func (it *QuantumGravityBridgeDataRootTupleRootEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuantumGravityBridgeDataRootTupleRootEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuantumGravityBridgeDataRootTupleRootEvent represents a DataRootTupleRootEvent event raised by the QuantumGravityBridge contract.
type QuantumGravityBridgeDataRootTupleRootEvent struct {
	Nonce             *big.Int
	DataRootTupleRoot [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDataRootTupleRootEvent is a free log retrieval operation binding the contract event 0x6614d037bde4905e31ca5ff05de61964c267f28b0320ed49e59f7d99752e1c4f.
//
// Solidity: event DataRootTupleRootEvent(uint256 indexed nonce, bytes32 dataRootTupleRoot)
func (_QuantumGravityBridge *QuantumGravityBridgeFilterer) FilterDataRootTupleRootEvent(opts *bind.FilterOpts, nonce []*big.Int) (*QuantumGravityBridgeDataRootTupleRootEventIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _QuantumGravityBridge.contract.FilterLogs(opts, "DataRootTupleRootEvent", nonceRule)
	if err != nil {
		return nil, err
	}
	return &QuantumGravityBridgeDataRootTupleRootEventIterator{contract: _QuantumGravityBridge.contract, event: "DataRootTupleRootEvent", logs: logs, sub: sub}, nil
}

// WatchDataRootTupleRootEvent is a free log subscription operation binding the contract event 0x6614d037bde4905e31ca5ff05de61964c267f28b0320ed49e59f7d99752e1c4f.
//
// Solidity: event DataRootTupleRootEvent(uint256 indexed nonce, bytes32 dataRootTupleRoot)
func (_QuantumGravityBridge *QuantumGravityBridgeFilterer) WatchDataRootTupleRootEvent(opts *bind.WatchOpts, sink chan<- *QuantumGravityBridgeDataRootTupleRootEvent, nonce []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _QuantumGravityBridge.contract.WatchLogs(opts, "DataRootTupleRootEvent", nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuantumGravityBridgeDataRootTupleRootEvent)
				if err := _QuantumGravityBridge.contract.UnpackLog(event, "DataRootTupleRootEvent", log); err != nil {
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

// ParseDataRootTupleRootEvent is a log parse operation binding the contract event 0x6614d037bde4905e31ca5ff05de61964c267f28b0320ed49e59f7d99752e1c4f.
//
// Solidity: event DataRootTupleRootEvent(uint256 indexed nonce, bytes32 dataRootTupleRoot)
func (_QuantumGravityBridge *QuantumGravityBridgeFilterer) ParseDataRootTupleRootEvent(log types.Log) (*QuantumGravityBridgeDataRootTupleRootEvent, error) {
	event := new(QuantumGravityBridgeDataRootTupleRootEvent)
	if err := _QuantumGravityBridge.contract.UnpackLog(event, "DataRootTupleRootEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuantumGravityBridgeValidatorSetUpdatedEventIterator is returned from FilterValidatorSetUpdatedEvent and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdatedEvent events raised by the QuantumGravityBridge contract.
type QuantumGravityBridgeValidatorSetUpdatedEventIterator struct {
	Event *QuantumGravityBridgeValidatorSetUpdatedEvent // Event containing the contract specifics and raw log

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
func (it *QuantumGravityBridgeValidatorSetUpdatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuantumGravityBridgeValidatorSetUpdatedEvent)
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
		it.Event = new(QuantumGravityBridgeValidatorSetUpdatedEvent)
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
func (it *QuantumGravityBridgeValidatorSetUpdatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuantumGravityBridgeValidatorSetUpdatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuantumGravityBridgeValidatorSetUpdatedEvent represents a ValidatorSetUpdatedEvent event raised by the QuantumGravityBridge contract.
type QuantumGravityBridgeValidatorSetUpdatedEvent struct {
	Nonce            *big.Int
	PowerThreshold   *big.Int
	ValidatorSetHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdatedEvent is a free log retrieval operation binding the contract event 0xe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c.
//
// Solidity: event ValidatorSetUpdatedEvent(uint256 indexed nonce, uint256 powerThreshold, bytes32 validatorSetHash)
func (_QuantumGravityBridge *QuantumGravityBridgeFilterer) FilterValidatorSetUpdatedEvent(opts *bind.FilterOpts, nonce []*big.Int) (*QuantumGravityBridgeValidatorSetUpdatedEventIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _QuantumGravityBridge.contract.FilterLogs(opts, "ValidatorSetUpdatedEvent", nonceRule)
	if err != nil {
		return nil, err
	}
	return &QuantumGravityBridgeValidatorSetUpdatedEventIterator{contract: _QuantumGravityBridge.contract, event: "ValidatorSetUpdatedEvent", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdatedEvent is a free log subscription operation binding the contract event 0xe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c.
//
// Solidity: event ValidatorSetUpdatedEvent(uint256 indexed nonce, uint256 powerThreshold, bytes32 validatorSetHash)
func (_QuantumGravityBridge *QuantumGravityBridgeFilterer) WatchValidatorSetUpdatedEvent(opts *bind.WatchOpts, sink chan<- *QuantumGravityBridgeValidatorSetUpdatedEvent, nonce []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _QuantumGravityBridge.contract.WatchLogs(opts, "ValidatorSetUpdatedEvent", nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuantumGravityBridgeValidatorSetUpdatedEvent)
				if err := _QuantumGravityBridge.contract.UnpackLog(event, "ValidatorSetUpdatedEvent", log); err != nil {
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

// ParseValidatorSetUpdatedEvent is a log parse operation binding the contract event 0xe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c.
//
// Solidity: event ValidatorSetUpdatedEvent(uint256 indexed nonce, uint256 powerThreshold, bytes32 validatorSetHash)
func (_QuantumGravityBridge *QuantumGravityBridgeFilterer) ParseValidatorSetUpdatedEvent(log types.Log) (*QuantumGravityBridgeValidatorSetUpdatedEvent, error) {
	event := new(QuantumGravityBridgeValidatorSetUpdatedEvent)
	if err := _QuantumGravityBridge.contract.UnpackLog(event, "ValidatorSetUpdatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
