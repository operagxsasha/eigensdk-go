// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIndexRegistry

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

// IIndexRegistryOperatorUpdate is an auto generated low-level Go binding around an user-defined struct.
type IIndexRegistryOperatorUpdate struct {
	FromBlockNumber uint32
	OperatorId      [32]byte
}

// IIndexRegistryQuorumUpdate is an auto generated low-level Go binding around an user-defined struct.
type IIndexRegistryQuorumUpdate struct {
	FromBlockNumber uint32
	NumOperators    uint32
}

// ContractIndexRegistryMetaData contains all meta data concerning the ContractIndexRegistry contract.
var ContractIndexRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"OPERATOR_DOES_NOT_EXIST_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentOperatorIndex\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLatestOperatorUpdate\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIndexRegistry.OperatorUpdate\",\"components\":[{\"name\":\"fromBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestQuorumUpdate\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIndexRegistry.QuorumUpdate\",\"components\":[{\"name\":\"fromBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"numOperators\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorListAtBlockNumber\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorUpdateAtIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"arrayIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIndexRegistry.OperatorUpdate\",\"components\":[{\"name\":\"fromBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumUpdateAtIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"quorumIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIndexRegistry.QuorumUpdate\",\"components\":[{\"name\":\"fromBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"numOperators\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializeQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalOperatorsForQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumIndexUpdate\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"newOperatorIndex\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false}]",
	Bin: "0x60a060405234801561000f575f5ffd5b506040516112b53803806112b583398101604081905261002e91610108565b6001600160a01b0381166080528061004461004b565b5050610135565b5f54610100900460ff16156100b65760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610106575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f60208284031215610118575f5ffd5b81516001600160a01b038116811461012e575f5ffd5b9392505050565b6080516111616101545f395f818161013e015261082001526111615ff3fe608060405234801561000f575f5ffd5b50600436106100b0575f3560e01c8063890262451161006e57806389026245146101af578063a48bb0ac146101cf578063bd29b8cd146101e2578063caa3cd76146101f5578063e2e685801461020a578063f34109221461024f575f5ffd5b8062bff04d146100b457806312d1d74d146100dd57806326d941f2146101115780632ed583e5146101265780636d14a987146101395780638121906f14610178575b5f5ffd5b6100c76100c2366004610e6a565b610262565b6040516100d49190610ee1565b60405180910390f35b6100f06100eb366004610f51565b610373565b60408051825163ffffffff16815260209283015192810192909252016100d4565b61012461011f366004610f82565b6103b8565b005b6100f0610134366004610f9b565b61049a565b6101607f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100d4565b61018b610186366004610f82565b61051d565b60408051825163ffffffff90811682526020938401511692810192909252016100d4565b6101c26101bd366004610f51565b610563565b6040516100d49190610fdb565b61018b6101dd366004610f51565b6106be565b6101246101f0366004610e6a565b610732565b6101fc5f81565b6040519081526020016100d4565b61023a610218366004611012565b600160209081525f928352604080842090915290825290205463ffffffff1681565b60405163ffffffff90911681526020016100d4565b61023a61025d366004610f82565b6107f7565b606061026c610815565b5f8267ffffffffffffffff8111156102865761028661103a565b6040519080825280602002602001820160405280156102af578160200160208202803683370190505b5090505f5b83811015610368575f8585838181106102cf576102cf61104e565b919091013560f81c5f818152600360205260408120549193509091508190036103135760405162461bcd60e51b815260040161030a90611062565b60405180910390fd5b5f61031d836108cd565b9050610334898461032f6001856110cb565b6109c4565b808585815181106103475761034761104e565b63ffffffff92909216602092830291909101909101525050506001016102b4565b5090505b9392505050565b604080518082019091525f80825260208201526103908383610a4c565b60408051808201909152815463ffffffff168152600190910154602082015290505b92915050565b6103c0610815565b60ff81165f90815260036020526040902054156104395760405162461bcd60e51b815260206004820152603160248201527f496e64657852656769737472792e63726561746551756f72756d3a2071756f72604482015270756d20616c72656164792065786973747360781b606482015260840161030a565b60ff165f908152600360209081526040808320815180830190925263ffffffff438116835282840185815282546001810184559286529390942091519101805492518416600160201b0267ffffffffffffffff199093169190931617179055565b604080518082019091525f808252602082015260ff84165f90815260026020908152604080832063ffffffff808816855292529091208054909184169081106104e5576104e561104e565b5f91825260209182902060408051808201909152600290920201805463ffffffff168252600101549181019190915290509392505050565b604080518082019091525f808252602082015261053982610aa1565b60408051808201909152905463ffffffff8082168352600160201b90910416602082015292915050565b60605f6105708484610ae0565b90505f8163ffffffff1667ffffffffffffffff8111156105925761059261103a565b6040519080825280602002602001820160405280156105bb578160200160208202803683370190505b5090505f5b8263ffffffff168110156106b5576105d9868287610c12565b8282815181106105eb576105eb61104e565b6020026020010181815250505f5f1b82828151811061060c5761060c61104e565b6020026020010151036106ad5760405162461bcd60e51b815260206004820152605d60248201527f496e64657852656769737472792e6765744f70657261746f724c69737441744260448201527f6c6f636b4e756d6265723a206f70657261746f7220646f6573206e6f7420657860648201527f6973742061742074686520676976656e20626c6f636b206e756d626572000000608482015260a40161030a565b6001016105c0565b50949350505050565b604080518082019091525f808252602082015260ff83165f908152600360205260409020805463ffffffff84169081106106fa576106fa61104e565b5f9182526020918290206040805180820190915291015463ffffffff8082168352600160201b90910416918101919091529392505050565b61073a610815565b5f5b818110156107f1575f8383838181106107575761075761104e565b919091013560f81c5f818152600360205260408120549193509091508190036107925760405162461bcd60e51b815260040161030a90611062565b60ff82165f90815260016020908152604080832089845290915281205463ffffffff16906107bf84610ce5565b90505f6107cc8583610d1d565b90508089146107e0576107e08186856109c4565b50506001909301925061073c915050565b50505050565b5f61080182610aa1565b54600160201b900463ffffffff1692915050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108cb5760405162461bcd60e51b815260206004820152604f60248201527f496e64657852656769737472792e5f636865636b5265676973747279436f6f7260448201527f64696e61746f723a2063616c6c6572206973206e6f742074686520726567697360648201526e3a393c9031b7b7b93234b730ba37b960891b608482015260a40161030a565b565b5f5f6108d883610aa1565b80549091505f906108f790600160201b900463ffffffff1660016110e7565b9050610904848383610d45565b60ff84165f908152600260205260408120906109216001846110cb565b63ffffffff16815260208101919091526040015f90812054900361036c5760ff84165f9081526002602052604081209061095c6001846110cb565b63ffffffff908116825260208083019390935260409182015f908120835180850190945243831684528385018281528154600180820184559284529590922093516002909502909301805463ffffffff19169490921693909317815591519101559392505050565b5f6109cf8383610a4c565b90506109dd83838387610de2565b60ff83165f818152600160209081526040808320888452825291829020805463ffffffff191663ffffffff871690811790915582519384529083015285917f6ee1e4f4075f3d067176140d34e87874244dd273294c05b2218133e49a2ba6f6910160405180910390a250505050565b60ff82165f90815260026020908152604080832063ffffffff851684529091528120805490610a7c600183611103565b81548110610a8c57610a8c61104e565b905f5260205f20906002020191505092915050565b60ff81165f908152600360205260408120805490610ac0600183611103565b81548110610ad057610ad061104e565b905f5260205f2001915050919050565b60ff82165f90815260036020526040812054805b8015610b855760ff85165f908152600360205260408120610b16600184611103565b81548110610b2657610b2661104e565b5f9182526020918290206040805180820190915291015463ffffffff808216808452600160201b90920481169383019390935290925090861610610b72576020015192506103b2915050565b5080610b7d81611116565b915050610af4565b5060405162461bcd60e51b815260206004820152605560248201527f496e64657852656769737472792e5f6f70657261746f72436f756e744174426c60448201527f6f636b4e756d6265723a2071756f72756d20646964206e6f742065786973742060648201527430ba1033b4bb32b710313637b1b590373ab6b132b960591b608482015260a40161030a565b60ff83165f90815260026020908152604080832063ffffffff86168452909152812054805b8015610cda5760ff86165f90815260026020908152604080832063ffffffff891684529091528120610c6a600184611103565b81548110610c7a57610c7a61104e565b5f91825260209182902060408051808201909152600290920201805463ffffffff9081168084526001909201549383019390935290925090861610610cc75760200151925061036c915050565b5080610cd281611116565b915050610c37565b505f95945050505050565b5f5f610cf083610aa1565b80549091505f90610d1090600190600160201b900463ffffffff166110cb565b905061036c848383610d45565b5f5f610d298484610a4c565b6001810154909150610d3d8585845f610de2565b949350505050565b815463ffffffff438116911603610d7a57815463ffffffff8216600160201b0267ffffffff0000000019909116178255505050565b60ff83165f908152600360209081526040808320815180830190925263ffffffff438116835285811683850190815282546001810184559286529390942091519101805492518416600160201b0267ffffffffffffffff199093169190931617179055505050565b815463ffffffff438116911603610dff57600182018190556107f1565b60ff939093165f90815260026020818152604080842063ffffffff968716855282528084208151808301909252438716825281830197885280546001808201835591865292909420905191909202909101805463ffffffff1916919094161783559251919092015550565b5f5f5f60408486031215610e7c575f5ffd5b83359250602084013567ffffffffffffffff811115610e99575f5ffd5b8401601f81018613610ea9575f5ffd5b803567ffffffffffffffff811115610ebf575f5ffd5b866020828401011115610ed0575f5ffd5b939660209190910195509293505050565b602080825282518282018190525f918401906040840190835b81811015610f1e57835163ffffffff16835260209384019390920191600101610efa565b509095945050505050565b803560ff81168114610f39575f5ffd5b919050565b803563ffffffff81168114610f39575f5ffd5b5f5f60408385031215610f62575f5ffd5b610f6b83610f29565b9150610f7960208401610f3e565b90509250929050565b5f60208284031215610f92575f5ffd5b61036c82610f29565b5f5f5f60608486031215610fad575f5ffd5b610fb684610f29565b9250610fc460208501610f3e565b9150610fd260408501610f3e565b90509250925092565b602080825282518282018190525f918401906040840190835b81811015610f1e578351835260209384019390920191600101610ff4565b5f5f60408385031215611023575f5ffd5b61102c83610f29565b946020939093013593505050565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b60208082526035908201527f496e64657852656769737472792e72656769737465724f70657261746f723a206040820152741c5d5bdc9d5b48191bd95cc81b9bdd08195e1a5cdd605a1b606082015260800190565b634e487b7160e01b5f52601160045260245ffd5b63ffffffff82811682821603908111156103b2576103b26110b7565b63ffffffff81811683821601908111156103b2576103b26110b7565b818103818111156103b2576103b26110b7565b5f81611124576111246110b7565b505f19019056fea2646970667358221220b150369cdcfc3f64114bf764591cf5d397b1ab8c0d1f490171b515e1dc50cbd664736f6c634300081b0033",
}

// ContractIndexRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIndexRegistryMetaData.ABI instead.
var ContractIndexRegistryABI = ContractIndexRegistryMetaData.ABI

// ContractIndexRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractIndexRegistryMetaData.Bin instead.
var ContractIndexRegistryBin = ContractIndexRegistryMetaData.Bin

// DeployContractIndexRegistry deploys a new Ethereum contract, binding an instance of ContractIndexRegistry to it.
func DeployContractIndexRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address) (common.Address, *types.Transaction, *ContractIndexRegistry, error) {
	parsed, err := ContractIndexRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractIndexRegistryBin), backend, _registryCoordinator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractIndexRegistry{ContractIndexRegistryCaller: ContractIndexRegistryCaller{contract: contract}, ContractIndexRegistryTransactor: ContractIndexRegistryTransactor{contract: contract}, ContractIndexRegistryFilterer: ContractIndexRegistryFilterer{contract: contract}}, nil
}

// ContractIndexRegistryMethods is an auto generated interface around an Ethereum contract.
type ContractIndexRegistryMethods interface {
	ContractIndexRegistryCalls
	ContractIndexRegistryTransacts
	ContractIndexRegistryFilters
}

// ContractIndexRegistryCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractIndexRegistryCalls interface {
	OPERATORDOESNOTEXISTID(opts *bind.CallOpts) ([32]byte, error)

	CurrentOperatorIndex(opts *bind.CallOpts, arg0 uint8, arg1 [32]byte) (uint32, error)

	GetLatestOperatorUpdate(opts *bind.CallOpts, quorumNumber uint8, operatorIndex uint32) (IIndexRegistryOperatorUpdate, error)

	GetLatestQuorumUpdate(opts *bind.CallOpts, quorumNumber uint8) (IIndexRegistryQuorumUpdate, error)

	GetOperatorListAtBlockNumber(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32) ([][32]byte, error)

	GetOperatorUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, operatorIndex uint32, arrayIndex uint32) (IIndexRegistryOperatorUpdate, error)

	GetQuorumUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, quorumIndex uint32) (IIndexRegistryQuorumUpdate, error)

	RegistryCoordinator(opts *bind.CallOpts) (common.Address, error)

	TotalOperatorsForQuorum(opts *bind.CallOpts, quorumNumber uint8) (uint32, error)
}

// ContractIndexRegistryTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractIndexRegistryTransacts interface {
	DeregisterOperator(opts *bind.TransactOpts, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error)

	InitializeQuorum(opts *bind.TransactOpts, quorumNumber uint8) (*types.Transaction, error)

	RegisterOperator(opts *bind.TransactOpts, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error)
}

// ContractIndexRegistryFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractIndexRegistryFilters interface {
	FilterInitialized(opts *bind.FilterOpts) (*ContractIndexRegistryInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractIndexRegistryInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractIndexRegistryInitialized, error)

	FilterQuorumIndexUpdate(opts *bind.FilterOpts, operatorId [][32]byte) (*ContractIndexRegistryQuorumIndexUpdateIterator, error)
	WatchQuorumIndexUpdate(opts *bind.WatchOpts, sink chan<- *ContractIndexRegistryQuorumIndexUpdate, operatorId [][32]byte) (event.Subscription, error)
	ParseQuorumIndexUpdate(log types.Log) (*ContractIndexRegistryQuorumIndexUpdate, error)
}

// ContractIndexRegistry is an auto generated Go binding around an Ethereum contract.
type ContractIndexRegistry struct {
	ContractIndexRegistryCaller     // Read-only binding to the contract
	ContractIndexRegistryTransactor // Write-only binding to the contract
	ContractIndexRegistryFilterer   // Log filterer for contract events
}

// ContractIndexRegistry implements the ContractIndexRegistryMethods interface.
var _ ContractIndexRegistryMethods = (*ContractIndexRegistry)(nil)

// ContractIndexRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIndexRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIndexRegistryCaller implements the ContractIndexRegistryCalls interface.
var _ ContractIndexRegistryCalls = (*ContractIndexRegistryCaller)(nil)

// ContractIndexRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIndexRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIndexRegistryTransactor implements the ContractIndexRegistryTransacts interface.
var _ ContractIndexRegistryTransacts = (*ContractIndexRegistryTransactor)(nil)

// ContractIndexRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIndexRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIndexRegistryFilterer implements the ContractIndexRegistryFilters interface.
var _ ContractIndexRegistryFilters = (*ContractIndexRegistryFilterer)(nil)

// ContractIndexRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIndexRegistrySession struct {
	Contract     *ContractIndexRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContractIndexRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIndexRegistryCallerSession struct {
	Contract *ContractIndexRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ContractIndexRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIndexRegistryTransactorSession struct {
	Contract     *ContractIndexRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ContractIndexRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIndexRegistryRaw struct {
	Contract *ContractIndexRegistry // Generic contract binding to access the raw methods on
}

// ContractIndexRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIndexRegistryCallerRaw struct {
	Contract *ContractIndexRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIndexRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIndexRegistryTransactorRaw struct {
	Contract *ContractIndexRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIndexRegistry creates a new instance of ContractIndexRegistry, bound to a specific deployed contract.
func NewContractIndexRegistry(address common.Address, backend bind.ContractBackend) (*ContractIndexRegistry, error) {
	contract, err := bindContractIndexRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIndexRegistry{ContractIndexRegistryCaller: ContractIndexRegistryCaller{contract: contract}, ContractIndexRegistryTransactor: ContractIndexRegistryTransactor{contract: contract}, ContractIndexRegistryFilterer: ContractIndexRegistryFilterer{contract: contract}}, nil
}

// NewContractIndexRegistryCaller creates a new read-only instance of ContractIndexRegistry, bound to a specific deployed contract.
func NewContractIndexRegistryCaller(address common.Address, caller bind.ContractCaller) (*ContractIndexRegistryCaller, error) {
	contract, err := bindContractIndexRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIndexRegistryCaller{contract: contract}, nil
}

// NewContractIndexRegistryTransactor creates a new write-only instance of ContractIndexRegistry, bound to a specific deployed contract.
func NewContractIndexRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIndexRegistryTransactor, error) {
	contract, err := bindContractIndexRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIndexRegistryTransactor{contract: contract}, nil
}

// NewContractIndexRegistryFilterer creates a new log filterer instance of ContractIndexRegistry, bound to a specific deployed contract.
func NewContractIndexRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIndexRegistryFilterer, error) {
	contract, err := bindContractIndexRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIndexRegistryFilterer{contract: contract}, nil
}

// bindContractIndexRegistry binds a generic wrapper to an already deployed contract.
func bindContractIndexRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIndexRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIndexRegistry *ContractIndexRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIndexRegistry.Contract.ContractIndexRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIndexRegistry *ContractIndexRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.ContractIndexRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIndexRegistry *ContractIndexRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.ContractIndexRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIndexRegistry *ContractIndexRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIndexRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIndexRegistry *ContractIndexRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIndexRegistry *ContractIndexRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.contract.Transact(opts, method, params...)
}

// OPERATORDOESNOTEXISTID is a free data retrieval call binding the contract method 0xcaa3cd76.
//
// Solidity: function OPERATOR_DOES_NOT_EXIST_ID() view returns(bytes32)
func (_ContractIndexRegistry *ContractIndexRegistryCaller) OPERATORDOESNOTEXISTID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "OPERATOR_DOES_NOT_EXIST_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORDOESNOTEXISTID is a free data retrieval call binding the contract method 0xcaa3cd76.
//
// Solidity: function OPERATOR_DOES_NOT_EXIST_ID() view returns(bytes32)
func (_ContractIndexRegistry *ContractIndexRegistrySession) OPERATORDOESNOTEXISTID() ([32]byte, error) {
	return _ContractIndexRegistry.Contract.OPERATORDOESNOTEXISTID(&_ContractIndexRegistry.CallOpts)
}

// OPERATORDOESNOTEXISTID is a free data retrieval call binding the contract method 0xcaa3cd76.
//
// Solidity: function OPERATOR_DOES_NOT_EXIST_ID() view returns(bytes32)
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) OPERATORDOESNOTEXISTID() ([32]byte, error) {
	return _ContractIndexRegistry.Contract.OPERATORDOESNOTEXISTID(&_ContractIndexRegistry.CallOpts)
}

// CurrentOperatorIndex is a free data retrieval call binding the contract method 0xe2e68580.
//
// Solidity: function currentOperatorIndex(uint8 , bytes32 ) view returns(uint32)
func (_ContractIndexRegistry *ContractIndexRegistryCaller) CurrentOperatorIndex(opts *bind.CallOpts, arg0 uint8, arg1 [32]byte) (uint32, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "currentOperatorIndex", arg0, arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CurrentOperatorIndex is a free data retrieval call binding the contract method 0xe2e68580.
//
// Solidity: function currentOperatorIndex(uint8 , bytes32 ) view returns(uint32)
func (_ContractIndexRegistry *ContractIndexRegistrySession) CurrentOperatorIndex(arg0 uint8, arg1 [32]byte) (uint32, error) {
	return _ContractIndexRegistry.Contract.CurrentOperatorIndex(&_ContractIndexRegistry.CallOpts, arg0, arg1)
}

// CurrentOperatorIndex is a free data retrieval call binding the contract method 0xe2e68580.
//
// Solidity: function currentOperatorIndex(uint8 , bytes32 ) view returns(uint32)
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) CurrentOperatorIndex(arg0 uint8, arg1 [32]byte) (uint32, error) {
	return _ContractIndexRegistry.Contract.CurrentOperatorIndex(&_ContractIndexRegistry.CallOpts, arg0, arg1)
}

// GetLatestOperatorUpdate is a free data retrieval call binding the contract method 0x12d1d74d.
//
// Solidity: function getLatestOperatorUpdate(uint8 quorumNumber, uint32 operatorIndex) view returns((uint32,bytes32))
func (_ContractIndexRegistry *ContractIndexRegistryCaller) GetLatestOperatorUpdate(opts *bind.CallOpts, quorumNumber uint8, operatorIndex uint32) (IIndexRegistryOperatorUpdate, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "getLatestOperatorUpdate", quorumNumber, operatorIndex)

	if err != nil {
		return *new(IIndexRegistryOperatorUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IIndexRegistryOperatorUpdate)).(*IIndexRegistryOperatorUpdate)

	return out0, err

}

// GetLatestOperatorUpdate is a free data retrieval call binding the contract method 0x12d1d74d.
//
// Solidity: function getLatestOperatorUpdate(uint8 quorumNumber, uint32 operatorIndex) view returns((uint32,bytes32))
func (_ContractIndexRegistry *ContractIndexRegistrySession) GetLatestOperatorUpdate(quorumNumber uint8, operatorIndex uint32) (IIndexRegistryOperatorUpdate, error) {
	return _ContractIndexRegistry.Contract.GetLatestOperatorUpdate(&_ContractIndexRegistry.CallOpts, quorumNumber, operatorIndex)
}

// GetLatestOperatorUpdate is a free data retrieval call binding the contract method 0x12d1d74d.
//
// Solidity: function getLatestOperatorUpdate(uint8 quorumNumber, uint32 operatorIndex) view returns((uint32,bytes32))
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) GetLatestOperatorUpdate(quorumNumber uint8, operatorIndex uint32) (IIndexRegistryOperatorUpdate, error) {
	return _ContractIndexRegistry.Contract.GetLatestOperatorUpdate(&_ContractIndexRegistry.CallOpts, quorumNumber, operatorIndex)
}

// GetLatestQuorumUpdate is a free data retrieval call binding the contract method 0x8121906f.
//
// Solidity: function getLatestQuorumUpdate(uint8 quorumNumber) view returns((uint32,uint32))
func (_ContractIndexRegistry *ContractIndexRegistryCaller) GetLatestQuorumUpdate(opts *bind.CallOpts, quorumNumber uint8) (IIndexRegistryQuorumUpdate, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "getLatestQuorumUpdate", quorumNumber)

	if err != nil {
		return *new(IIndexRegistryQuorumUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IIndexRegistryQuorumUpdate)).(*IIndexRegistryQuorumUpdate)

	return out0, err

}

// GetLatestQuorumUpdate is a free data retrieval call binding the contract method 0x8121906f.
//
// Solidity: function getLatestQuorumUpdate(uint8 quorumNumber) view returns((uint32,uint32))
func (_ContractIndexRegistry *ContractIndexRegistrySession) GetLatestQuorumUpdate(quorumNumber uint8) (IIndexRegistryQuorumUpdate, error) {
	return _ContractIndexRegistry.Contract.GetLatestQuorumUpdate(&_ContractIndexRegistry.CallOpts, quorumNumber)
}

// GetLatestQuorumUpdate is a free data retrieval call binding the contract method 0x8121906f.
//
// Solidity: function getLatestQuorumUpdate(uint8 quorumNumber) view returns((uint32,uint32))
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) GetLatestQuorumUpdate(quorumNumber uint8) (IIndexRegistryQuorumUpdate, error) {
	return _ContractIndexRegistry.Contract.GetLatestQuorumUpdate(&_ContractIndexRegistry.CallOpts, quorumNumber)
}

// GetOperatorListAtBlockNumber is a free data retrieval call binding the contract method 0x89026245.
//
// Solidity: function getOperatorListAtBlockNumber(uint8 quorumNumber, uint32 blockNumber) view returns(bytes32[])
func (_ContractIndexRegistry *ContractIndexRegistryCaller) GetOperatorListAtBlockNumber(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32) ([][32]byte, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "getOperatorListAtBlockNumber", quorumNumber, blockNumber)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetOperatorListAtBlockNumber is a free data retrieval call binding the contract method 0x89026245.
//
// Solidity: function getOperatorListAtBlockNumber(uint8 quorumNumber, uint32 blockNumber) view returns(bytes32[])
func (_ContractIndexRegistry *ContractIndexRegistrySession) GetOperatorListAtBlockNumber(quorumNumber uint8, blockNumber uint32) ([][32]byte, error) {
	return _ContractIndexRegistry.Contract.GetOperatorListAtBlockNumber(&_ContractIndexRegistry.CallOpts, quorumNumber, blockNumber)
}

// GetOperatorListAtBlockNumber is a free data retrieval call binding the contract method 0x89026245.
//
// Solidity: function getOperatorListAtBlockNumber(uint8 quorumNumber, uint32 blockNumber) view returns(bytes32[])
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) GetOperatorListAtBlockNumber(quorumNumber uint8, blockNumber uint32) ([][32]byte, error) {
	return _ContractIndexRegistry.Contract.GetOperatorListAtBlockNumber(&_ContractIndexRegistry.CallOpts, quorumNumber, blockNumber)
}

// GetOperatorUpdateAtIndex is a free data retrieval call binding the contract method 0x2ed583e5.
//
// Solidity: function getOperatorUpdateAtIndex(uint8 quorumNumber, uint32 operatorIndex, uint32 arrayIndex) view returns((uint32,bytes32))
func (_ContractIndexRegistry *ContractIndexRegistryCaller) GetOperatorUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, operatorIndex uint32, arrayIndex uint32) (IIndexRegistryOperatorUpdate, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "getOperatorUpdateAtIndex", quorumNumber, operatorIndex, arrayIndex)

	if err != nil {
		return *new(IIndexRegistryOperatorUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IIndexRegistryOperatorUpdate)).(*IIndexRegistryOperatorUpdate)

	return out0, err

}

// GetOperatorUpdateAtIndex is a free data retrieval call binding the contract method 0x2ed583e5.
//
// Solidity: function getOperatorUpdateAtIndex(uint8 quorumNumber, uint32 operatorIndex, uint32 arrayIndex) view returns((uint32,bytes32))
func (_ContractIndexRegistry *ContractIndexRegistrySession) GetOperatorUpdateAtIndex(quorumNumber uint8, operatorIndex uint32, arrayIndex uint32) (IIndexRegistryOperatorUpdate, error) {
	return _ContractIndexRegistry.Contract.GetOperatorUpdateAtIndex(&_ContractIndexRegistry.CallOpts, quorumNumber, operatorIndex, arrayIndex)
}

// GetOperatorUpdateAtIndex is a free data retrieval call binding the contract method 0x2ed583e5.
//
// Solidity: function getOperatorUpdateAtIndex(uint8 quorumNumber, uint32 operatorIndex, uint32 arrayIndex) view returns((uint32,bytes32))
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) GetOperatorUpdateAtIndex(quorumNumber uint8, operatorIndex uint32, arrayIndex uint32) (IIndexRegistryOperatorUpdate, error) {
	return _ContractIndexRegistry.Contract.GetOperatorUpdateAtIndex(&_ContractIndexRegistry.CallOpts, quorumNumber, operatorIndex, arrayIndex)
}

// GetQuorumUpdateAtIndex is a free data retrieval call binding the contract method 0xa48bb0ac.
//
// Solidity: function getQuorumUpdateAtIndex(uint8 quorumNumber, uint32 quorumIndex) view returns((uint32,uint32))
func (_ContractIndexRegistry *ContractIndexRegistryCaller) GetQuorumUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, quorumIndex uint32) (IIndexRegistryQuorumUpdate, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "getQuorumUpdateAtIndex", quorumNumber, quorumIndex)

	if err != nil {
		return *new(IIndexRegistryQuorumUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IIndexRegistryQuorumUpdate)).(*IIndexRegistryQuorumUpdate)

	return out0, err

}

// GetQuorumUpdateAtIndex is a free data retrieval call binding the contract method 0xa48bb0ac.
//
// Solidity: function getQuorumUpdateAtIndex(uint8 quorumNumber, uint32 quorumIndex) view returns((uint32,uint32))
func (_ContractIndexRegistry *ContractIndexRegistrySession) GetQuorumUpdateAtIndex(quorumNumber uint8, quorumIndex uint32) (IIndexRegistryQuorumUpdate, error) {
	return _ContractIndexRegistry.Contract.GetQuorumUpdateAtIndex(&_ContractIndexRegistry.CallOpts, quorumNumber, quorumIndex)
}

// GetQuorumUpdateAtIndex is a free data retrieval call binding the contract method 0xa48bb0ac.
//
// Solidity: function getQuorumUpdateAtIndex(uint8 quorumNumber, uint32 quorumIndex) view returns((uint32,uint32))
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) GetQuorumUpdateAtIndex(quorumNumber uint8, quorumIndex uint32) (IIndexRegistryQuorumUpdate, error) {
	return _ContractIndexRegistry.Contract.GetQuorumUpdateAtIndex(&_ContractIndexRegistry.CallOpts, quorumNumber, quorumIndex)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIndexRegistry *ContractIndexRegistryCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIndexRegistry *ContractIndexRegistrySession) RegistryCoordinator() (common.Address, error) {
	return _ContractIndexRegistry.Contract.RegistryCoordinator(&_ContractIndexRegistry.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIndexRegistry.Contract.RegistryCoordinator(&_ContractIndexRegistry.CallOpts)
}

// TotalOperatorsForQuorum is a free data retrieval call binding the contract method 0xf3410922.
//
// Solidity: function totalOperatorsForQuorum(uint8 quorumNumber) view returns(uint32)
func (_ContractIndexRegistry *ContractIndexRegistryCaller) TotalOperatorsForQuorum(opts *bind.CallOpts, quorumNumber uint8) (uint32, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "totalOperatorsForQuorum", quorumNumber)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TotalOperatorsForQuorum is a free data retrieval call binding the contract method 0xf3410922.
//
// Solidity: function totalOperatorsForQuorum(uint8 quorumNumber) view returns(uint32)
func (_ContractIndexRegistry *ContractIndexRegistrySession) TotalOperatorsForQuorum(quorumNumber uint8) (uint32, error) {
	return _ContractIndexRegistry.Contract.TotalOperatorsForQuorum(&_ContractIndexRegistry.CallOpts, quorumNumber)
}

// TotalOperatorsForQuorum is a free data retrieval call binding the contract method 0xf3410922.
//
// Solidity: function totalOperatorsForQuorum(uint8 quorumNumber) view returns(uint32)
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) TotalOperatorsForQuorum(quorumNumber uint8) (uint32, error) {
	return _ContractIndexRegistry.Contract.TotalOperatorsForQuorum(&_ContractIndexRegistry.CallOpts, quorumNumber)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xbd29b8cd.
//
// Solidity: function deregisterOperator(bytes32 operatorId, bytes quorumNumbers) returns()
func (_ContractIndexRegistry *ContractIndexRegistryTransactor) DeregisterOperator(opts *bind.TransactOpts, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIndexRegistry.contract.Transact(opts, "deregisterOperator", operatorId, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xbd29b8cd.
//
// Solidity: function deregisterOperator(bytes32 operatorId, bytes quorumNumbers) returns()
func (_ContractIndexRegistry *ContractIndexRegistrySession) DeregisterOperator(operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.DeregisterOperator(&_ContractIndexRegistry.TransactOpts, operatorId, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xbd29b8cd.
//
// Solidity: function deregisterOperator(bytes32 operatorId, bytes quorumNumbers) returns()
func (_ContractIndexRegistry *ContractIndexRegistryTransactorSession) DeregisterOperator(operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.DeregisterOperator(&_ContractIndexRegistry.TransactOpts, operatorId, quorumNumbers)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_ContractIndexRegistry *ContractIndexRegistryTransactor) InitializeQuorum(opts *bind.TransactOpts, quorumNumber uint8) (*types.Transaction, error) {
	return _ContractIndexRegistry.contract.Transact(opts, "initializeQuorum", quorumNumber)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_ContractIndexRegistry *ContractIndexRegistrySession) InitializeQuorum(quorumNumber uint8) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.InitializeQuorum(&_ContractIndexRegistry.TransactOpts, quorumNumber)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_ContractIndexRegistry *ContractIndexRegistryTransactorSession) InitializeQuorum(quorumNumber uint8) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.InitializeQuorum(&_ContractIndexRegistry.TransactOpts, quorumNumber)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x00bff04d.
//
// Solidity: function registerOperator(bytes32 operatorId, bytes quorumNumbers) returns(uint32[])
func (_ContractIndexRegistry *ContractIndexRegistryTransactor) RegisterOperator(opts *bind.TransactOpts, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIndexRegistry.contract.Transact(opts, "registerOperator", operatorId, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x00bff04d.
//
// Solidity: function registerOperator(bytes32 operatorId, bytes quorumNumbers) returns(uint32[])
func (_ContractIndexRegistry *ContractIndexRegistrySession) RegisterOperator(operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.RegisterOperator(&_ContractIndexRegistry.TransactOpts, operatorId, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x00bff04d.
//
// Solidity: function registerOperator(bytes32 operatorId, bytes quorumNumbers) returns(uint32[])
func (_ContractIndexRegistry *ContractIndexRegistryTransactorSession) RegisterOperator(operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.RegisterOperator(&_ContractIndexRegistry.TransactOpts, operatorId, quorumNumbers)
}

// ContractIndexRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractIndexRegistry contract.
type ContractIndexRegistryInitializedIterator struct {
	Event *ContractIndexRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ContractIndexRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIndexRegistryInitialized)
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
		it.Event = new(ContractIndexRegistryInitialized)
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
func (it *ContractIndexRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIndexRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIndexRegistryInitialized represents a Initialized event raised by the ContractIndexRegistry contract.
type ContractIndexRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIndexRegistry *ContractIndexRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractIndexRegistryInitializedIterator, error) {

	logs, sub, err := _ContractIndexRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractIndexRegistryInitializedIterator{contract: _ContractIndexRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIndexRegistry *ContractIndexRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractIndexRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractIndexRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIndexRegistryInitialized)
				if err := _ContractIndexRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractIndexRegistry *ContractIndexRegistryFilterer) ParseInitialized(log types.Log) (*ContractIndexRegistryInitialized, error) {
	event := new(ContractIndexRegistryInitialized)
	if err := _ContractIndexRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIndexRegistryQuorumIndexUpdateIterator is returned from FilterQuorumIndexUpdate and is used to iterate over the raw logs and unpacked data for QuorumIndexUpdate events raised by the ContractIndexRegistry contract.
type ContractIndexRegistryQuorumIndexUpdateIterator struct {
	Event *ContractIndexRegistryQuorumIndexUpdate // Event containing the contract specifics and raw log

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
func (it *ContractIndexRegistryQuorumIndexUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIndexRegistryQuorumIndexUpdate)
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
		it.Event = new(ContractIndexRegistryQuorumIndexUpdate)
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
func (it *ContractIndexRegistryQuorumIndexUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIndexRegistryQuorumIndexUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIndexRegistryQuorumIndexUpdate represents a QuorumIndexUpdate event raised by the ContractIndexRegistry contract.
type ContractIndexRegistryQuorumIndexUpdate struct {
	OperatorId       [32]byte
	QuorumNumber     uint8
	NewOperatorIndex uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterQuorumIndexUpdate is a free log retrieval operation binding the contract event 0x6ee1e4f4075f3d067176140d34e87874244dd273294c05b2218133e49a2ba6f6.
//
// Solidity: event QuorumIndexUpdate(bytes32 indexed operatorId, uint8 quorumNumber, uint32 newOperatorIndex)
func (_ContractIndexRegistry *ContractIndexRegistryFilterer) FilterQuorumIndexUpdate(opts *bind.FilterOpts, operatorId [][32]byte) (*ContractIndexRegistryQuorumIndexUpdateIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractIndexRegistry.contract.FilterLogs(opts, "QuorumIndexUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractIndexRegistryQuorumIndexUpdateIterator{contract: _ContractIndexRegistry.contract, event: "QuorumIndexUpdate", logs: logs, sub: sub}, nil
}

// WatchQuorumIndexUpdate is a free log subscription operation binding the contract event 0x6ee1e4f4075f3d067176140d34e87874244dd273294c05b2218133e49a2ba6f6.
//
// Solidity: event QuorumIndexUpdate(bytes32 indexed operatorId, uint8 quorumNumber, uint32 newOperatorIndex)
func (_ContractIndexRegistry *ContractIndexRegistryFilterer) WatchQuorumIndexUpdate(opts *bind.WatchOpts, sink chan<- *ContractIndexRegistryQuorumIndexUpdate, operatorId [][32]byte) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractIndexRegistry.contract.WatchLogs(opts, "QuorumIndexUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIndexRegistryQuorumIndexUpdate)
				if err := _ContractIndexRegistry.contract.UnpackLog(event, "QuorumIndexUpdate", log); err != nil {
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

// ParseQuorumIndexUpdate is a log parse operation binding the contract event 0x6ee1e4f4075f3d067176140d34e87874244dd273294c05b2218133e49a2ba6f6.
//
// Solidity: event QuorumIndexUpdate(bytes32 indexed operatorId, uint8 quorumNumber, uint32 newOperatorIndex)
func (_ContractIndexRegistry *ContractIndexRegistryFilterer) ParseQuorumIndexUpdate(log types.Log) (*ContractIndexRegistryQuorumIndexUpdate, error) {
	event := new(ContractIndexRegistryQuorumIndexUpdate)
	if err := _ContractIndexRegistry.contract.UnpackLog(event, "QuorumIndexUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
