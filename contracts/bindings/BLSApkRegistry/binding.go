// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractBLSApkRegistry

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSApkRegistryApkUpdate is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryApkUpdate struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}

// IBLSApkRegistryPubkeyRegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryPubkeyRegistrationParams struct {
	PubkeyRegistrationSignature BN254G1Point
	PubkeyG1                    BN254G1Point
	PubkeyG2                    BN254G2Point
}

// ContractBLSApkRegistryMetaData contains all meta data concerning the ContractBLSApkRegistry contract.
var ContractBLSApkRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"apkHistory\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"apkHash\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"},{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentApk\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getApk\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkHashAtBlockNumberAndIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkHistoryLength\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkIndicesAtBlockNumber\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkUpdateAtIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.ApkUpdate\",\"components\":[{\"name\":\"apkHash\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"},{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorFromPubkeyHash\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorId\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredPubkey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializeQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorToPubkey\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorToPubkeyHash\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pubkeyHashToOperator\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerBLSPublicKey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"pubkeyRegistrationMessageHash\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewPubkeyRegistration\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToQuorums\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromQuorums\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50604051611f06380380611f0683398101604081905261002e91610108565b6001600160a01b0381166080528061004461004b565b5050610135565b5f54610100900460ff16156100b65760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610106575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f60208284031215610118575f5ffd5b81516001600160a01b038116811461012e575f5ffd5b9392505050565b608051611db26101545f395f81816103030152610f9a0152611db25ff3fe608060405234801561000f575f5ffd5b5060043610610110575f3560e01c80636d14a9871161009e578063bf79ce581161006e578063bf79ce58146103bf578063d5254a8c146103d2578063de29fac0146103f2578063e8bb9ae614610411578063f4e24fe514610439575f5ffd5b80636d14a987146102fe5780637916cea6146103255780637ff81a8714610366578063a3db80e214610399575f5ffd5b80633fb27952116100e45780633fb27952146101d657806347b314e8146101e95780635f61a88414610229578063605747d51461028357806368bccaac146102d1575f5ffd5b8062a1f4cb1461011457806313542a4e1461015457806326d941f21461018a578063377ed99d1461019f575b5f5ffd5b61013a610122366004611852565b60036020525f90815260409020805460019091015482565b604080519283526020830191909152015b60405180910390f35b61017c610162366004611852565b6001600160a01b03165f9081526001602052604090205490565b60405190815260200161014b565b61019d610198366004611882565b61044c565b005b6101c16101ad366004611882565b60ff165f9081526004602052604090205490565b60405163ffffffff909116815260200161014b565b61019d6101e4366004611909565b61055f565b6102116101f73660046119ae565b5f908152600260205260409020546001600160a01b031690565b6040516001600160a01b03909116815260200161014b565b610276610237366004611882565b604080518082019091525f80825260208201525060ff165f90815260056020908152604091829020825180840190935280548352600101549082015290565b60405161014b91906119c5565b6102966102913660046119dc565b6105db565b60408051825167ffffffffffffffff1916815260208084015163ffffffff90811691830191909152928201519092169082015260600161014b565b6102e46102df366004611a04565b61066c565b60405167ffffffffffffffff19909116815260200161014b565b6102117f000000000000000000000000000000000000000000000000000000000000000081565b6103386103333660046119dc565b610809565b6040805167ffffffffffffffff19909416845263ffffffff928316602085015291169082015260600161014b565b610379610374366004611852565b610850565b60408051835181526020938401519381019390935282015260600161014b565b61013a6103a7366004611882565b60056020525f90815260409020805460019091015482565b61017c6103cd366004611a48565b61091b565b6103e56103e0366004611aa0565b610d21565b60405161014b9190611b12565b61017c610400366004611852565b60016020525f908152604090205481565b61021161041f3660046119ae565b60026020525f90815260409020546001600160a01b031681565b61019d610447366004611909565b610f28565b610454610f8f565b60ff81165f90815260046020526040902054156104d75760405162461bcd60e51b815260206004820152603660248201527f424c5341706b52656769737472792e696e697469616c697a6551756f72756d3a6044820152752071756f72756d20616c72656164792065786973747360501b60648201526084015b60405180910390fd5b60ff165f908152600460209081526040808320815160608101835284815263ffffffff4381168286019081528285018781528454600181018655948852959096209151919092018054955194518316600160e01b026001600160e01b0395909316600160c01b026001600160e01b03199096169190931c179390931791909116919091179055565b610567610f8f565b5f61057183610850565b50905061057e8282611048565b7f73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e836105be856001600160a01b03165f9081526001602052604090205490565b846040516105ce93929190611b5a565b60405180910390a1505050565b604080516060810182525f808252602080830182905282840182905260ff86168252600490529190912080548390811061061757610617611ba5565b5f91825260209182902060408051606081018252919092015467ffffffffffffffff1981841b16825263ffffffff600160c01b8204811694830194909452600160e01b90049092169082015290505b92915050565b60ff83165f90815260046020526040812080548291908490811061069257610692611ba5565b5f91825260209182902060408051606081018252919092015467ffffffffffffffff1981841b16825263ffffffff600160c01b82048116948301859052600160e01b90910481169282019290925292508516101561075a576040805162461bcd60e51b81526020600482015260248101919091527f424c5341706b52656769737472792e67657441706b486173684174426c6f636b60448201527f4e756d626572416e64496e6465783a20696e64657820746f6f20726563656e7460648201526084016104ce565b604081015163ffffffff1615806107805750806040015163ffffffff168463ffffffff16105b6108005760405162461bcd60e51b815260206004820152604560248201527f424c5341706b52656769737472792e67657441706b486173684174426c6f636b60448201527f4e756d626572416e64496e6465783a206e6f74206c61746573742061706b2075606482015264706461746560d81b608482015260a4016104ce565b51949350505050565b6004602052815f5260405f208181548110610822575f80fd5b5f91825260209091200154604081901b925063ffffffff600160c01b820481169250600160e01b9091041683565b604080518082019091525f80825260208201526001600160a01b0382165f818152600360209081526040808320815180830183528154815260019182015481850152948452909152812054909190806109115760405162461bcd60e51b815260206004820152603e60248201527f424c5341706b52656769737472792e676574526567697374657265645075626b60448201527f65793a206f70657261746f72206973206e6f742072656769737465726564000060648201526084016104ce565b9094909350915050565b5f610924610f8f565b5f61095061093a36869003860160408701611bb9565b80515f9081526020918201519091526040902090565b90507fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb581036109d6576040805162461bcd60e51b81526020600482015260248101919091525f516020611d5d5f395f51905f5260448201527f4b65793a2063616e6e6f74207265676973746572207a65726f207075626b657960648201526084016104ce565b6001600160a01b0385165f9081526001602052604090205415610a5e5760405162461bcd60e51b815260206004820152604760248201525f516020611d5d5f395f51905f5260448201527f4b65793a206f70657261746f7220616c72656164792072656769737465726564606482015266207075626b657960c81b608482015260a4016104ce565b5f818152600260205260409020546001600160a01b031615610ae05760405162461bcd60e51b815260206004820152604260248201525f516020611d5d5f395f51905f5260448201527f4b65793a207075626c6963206b657920616c7265616479207265676973746572606482015261195960f21b608482015260a4016104ce565b604080515f917f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000191610b38918835916020808b0135928b01359160608c01359160808d019160c08e01918d35918e8201359101611bea565b604051602081830303815290604052805190602001205f1c610b5a9190611c2c565b9050610bf3610b93610b7e83610b78368a90038a0160408b01611bb9565b90611284565b610b8d36899003890189611bb9565b90611313565b610b9b6113a6565b610bdc610bcd85610b786040805180820182525f80825260209182015281518083019092526001825260029082015290565b610b8d368a90038a018a611bb9565b610bee368a90038a0160808b01611c8d565b611466565b610c8d5760405162461bcd60e51b815260206004820152606c60248201525f516020611d5d5f395f51905f5260448201527f4b65793a2065697468657220746865204731207369676e61747572652069732060648201527f77726f6e672c206f7220473120616e642047322070726976617465206b65792060848201526b0c8de40dcdee840dac2e8c6d60a31b60a482015260c4016104ce565b6001600160a01b0386165f8181526003602090815260408083208982018035825560608b013560019283015590835281842087905586845260029092529182902080546001600160a01b0319168417905590517fe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba382804191610d109160808a0190611ccc565b60405180910390a250949350505050565b60605f8367ffffffffffffffff811115610d3d57610d3d61189b565b604051908082528060200260200182016040528015610d66578160200160208202803683370190505b5090505f5b84811015610f1f575f868683818110610d8657610d86611ba5565b919091013560f81c5f818152600460205260409020549092509050801580610de6575060ff82165f9081526004602052604081208054909190610dcb57610dcb611ba5565b5f91825260209091200154600160c01b900463ffffffff1686105b15610e735760405162461bcd60e51b815260206004820152605160248201527f424c5341706b52656769737472792e67657441706b496e64696365734174426c60448201527f6f636b4e756d6265723a20626c6f636b4e756d626572206973206265666f7265606482015270207468652066697273742075706461746560781b608482015260a4016104ce565b805b8015610f145760ff83165f9081526004602052604090208790610e99600184611d0a565b81548110610ea957610ea9611ba5565b5f91825260209091200154600160c01b900463ffffffff1611610f0257610ed1600182611d0a565b858581518110610ee357610ee3611ba5565b602002602001019063ffffffff16908163ffffffff1681525050610f14565b80610f0c81611d1d565b915050610e75565b505050600101610d6b565b50949350505050565b610f30610f8f565b5f610f3a83610850565b509050610f4f82610f4a836116c4565b611048565b7ff843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e836105be856001600160a01b03165f9081526001602052604090205490565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146110465760405162461bcd60e51b815260206004820152605060248201527f424c5341706b52656769737472792e5f636865636b5265676973747279436f6f60448201527f7264696e61746f723a2063616c6c6572206973206e6f7420746865207265676960648201526f39ba393c9031b7b7b93234b730ba37b960811b608482015260a4016104ce565b565b604080518082019091525f80825260208201525f5b835181101561127e575f84828151811061107957611079611ba5565b0160209081015160f81c5f818152600490925260408220549092509081900361110a5760405162461bcd60e51b815260206004820152603d60248201527f424c5341706b52656769737472792e5f70726f6365737351756f72756d41706b60448201527f5570646174653a2071756f72756d20646f6573206e6f7420657869737400000060648201526084016104ce565b60ff82165f90815260056020908152604091829020825180840190935280548352600101549082015261113d9086611313565b60ff83165f8181526005602090815260408083208551808255868401805160019384015590855251835281842094845260049092528220939750919290916111859085611d0a565b8154811061119557611195611ba5565b5f918252602090912001805490915063ffffffff438116600160c01b90920416036111d35780546001600160c01b031916604083901c17815561126e565b805463ffffffff438116600160e01b8181026001600160e01b0394851617855560ff88165f908152600460209081526040808320815160608101835267ffffffffffffffff198b16815280840196875280830185815282546001810184559286529390942093519301805495519251871690940291909516600160c01b026001600160e01b0319949094169190941c17919091179092161790555b50506001909201915061105d9050565b50505050565b604080518082019091525f808252602082015261129f611780565b835181526020808501519082015260408082018490525f908360608460076107d05a03fa905080806112cd57fe5b508061130b5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b60448201526064016104ce565b505092915050565b604080518082019091525f808252602082015261132e61179e565b835181526020808501518183015283516040808401919091529084015160608301525f908360808460066107d05a03fa9050808061136857fe5b508061130b5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b60448201526064016104ce565b6113ae6117bc565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b6040805180820182528581526020808201859052825180840190935285835282018390525f916114946117e1565b5f5b600281101561164b575f6114ab826006611d32565b90508482600281106114bf576114bf611ba5565b602002015151836114d0835f611d49565b600c81106114e0576114e0611ba5565b60200201528482600281106114f7576114f7611ba5565b6020020151602001518382600161150e9190611d49565b600c811061151e5761151e611ba5565b602002015283826002811061153557611535611ba5565b6020020151515183611548836002611d49565b600c811061155857611558611ba5565b602002015283826002811061156f5761156f611ba5565b6020020151516001602002015183611588836003611d49565b600c811061159857611598611ba5565b60200201528382600281106115af576115af611ba5565b6020020151602001515f600281106115c9576115c9611ba5565b6020020151836115da836004611d49565b600c81106115ea576115ea611ba5565b602002015283826002811061160157611601611ba5565b60200201516020015160016002811061161c5761161c611ba5565b60200201518361162d836005611d49565b600c811061163d5761163d611ba5565b602002015250600101611496565b50611654611800565b5f6020826101808560086107d05a03fa9050808061166e57fe5b50806116b45760405162461bcd60e51b81526020600482015260156024820152741c185a5c9a5b99cb5bdc18dbd9194b59985a5b1959605a1b60448201526064016104ce565b5051151598975050505050505050565b604080518082019091525f808252602082015281511580156116e857506020820151155b15611705575050604080518082019091525f808252602082015290565b6040518060400160405280835f015181526020017f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4784602001516117499190611c2c565b611773907f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47611d0a565b905292915050565b919050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806117cf61181e565b81526020016117dc61181e565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b80356001600160a01b038116811461177b575f5ffd5b5f60208284031215611862575f5ffd5b61186b8261183c565b9392505050565b803560ff8116811461177b575f5ffd5b5f60208284031215611892575f5ffd5b61186b82611872565b634e487b7160e01b5f52604160045260245ffd5b6040805190810167ffffffffffffffff811182821017156118d2576118d261189b565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156119015761190161189b565b604052919050565b5f5f6040838503121561191a575f5ffd5b6119238361183c565b9150602083013567ffffffffffffffff81111561193e575f5ffd5b8301601f8101851361194e575f5ffd5b803567ffffffffffffffff8111156119685761196861189b565b61197b601f8201601f19166020016118d8565b81815286602083850101111561198f575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f602082840312156119be575f5ffd5b5035919050565b815181526020808301519082015260408101610666565b5f5f604083850312156119ed575f5ffd5b6119f683611872565b946020939093013593505050565b5f5f5f60608486031215611a16575f5ffd5b611a1f84611872565b9250602084013563ffffffff81168114611a37575f5ffd5b929592945050506040919091013590565b5f5f5f838503610160811215611a5c575f5ffd5b611a658561183c565b9350610100601f1982011215611a79575f5ffd5b602085019250604061011f1982011215611a91575f5ffd5b50610120840190509250925092565b5f5f5f60408486031215611ab2575f5ffd5b833567ffffffffffffffff811115611ac8575f5ffd5b8401601f81018613611ad8575f5ffd5b803567ffffffffffffffff811115611aee575f5ffd5b866020828401011115611aff575f5ffd5b6020918201979096509401359392505050565b602080825282518282018190525f918401906040840190835b81811015611b4f57835163ffffffff16835260209384019390920191600101611b2b565b509095945050505050565b60018060a01b0384168152826020820152606060408201525f82518060608401528060208501608085015e5f608082850101526080601f19601f830116840101915050949350505050565b634e487b7160e01b5f52603260045260245ffd5b5f6040828403128015611bca575f5ffd5b50611bd36118af565b823581526020928301359281019290925250919050565b888152876020820152866040820152856060820152604085608083013760408460c0830137610100810192909252610120820152610140019695505050505050565b5f82611c4657634e487b7160e01b5f52601260045260245ffd5b500690565b5f82601f830112611c5a575f5ffd5b611c626118af565b806040840185811115611c73575f5ffd5b845b81811015611b4f578035845260209384019301611c75565b5f6080828403128015611c9e575f5ffd5b50611ca76118af565b611cb18484611c4b565b8152611cc08460408501611c4b565b60208201529392505050565b823581526020808401359082015260c0810160408381840137604080840160808401379392505050565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561066657610666611cf6565b5f81611d2b57611d2b611cf6565b505f190190565b808202811582820484141761066657610666611cf6565b8082018082111561066657610666611cf656fe424c5341706b52656769737472792e7265676973746572424c535075626c6963a26469706673582212200be37c9cf2e321a1cfe5210ea064163dddfd4022873736fd9937f4e974f04a0664736f6c634300081b0033",
}

// ContractBLSApkRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractBLSApkRegistryMetaData.ABI instead.
var ContractBLSApkRegistryABI = ContractBLSApkRegistryMetaData.ABI

// ContractBLSApkRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractBLSApkRegistryMetaData.Bin instead.
var ContractBLSApkRegistryBin = ContractBLSApkRegistryMetaData.Bin

// DeployContractBLSApkRegistry deploys a new Ethereum contract, binding an instance of ContractBLSApkRegistry to it.
func DeployContractBLSApkRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address) (common.Address, *types.Transaction, *ContractBLSApkRegistry, error) {
	parsed, err := ContractBLSApkRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBLSApkRegistryBin), backend, _registryCoordinator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractBLSApkRegistry{ContractBLSApkRegistryCaller: ContractBLSApkRegistryCaller{contract: contract}, ContractBLSApkRegistryTransactor: ContractBLSApkRegistryTransactor{contract: contract}, ContractBLSApkRegistryFilterer: ContractBLSApkRegistryFilterer{contract: contract}}, nil
}

// ContractBLSApkRegistryMethods is an auto generated interface around an Ethereum contract.
type ContractBLSApkRegistryMethods interface {
	ContractBLSApkRegistryCalls
	ContractBLSApkRegistryTransacts
	ContractBLSApkRegistryFilters
}

// ContractBLSApkRegistryCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractBLSApkRegistryCalls interface {
	ApkHistory(opts *bind.CallOpts, arg0 uint8, arg1 *big.Int) (struct {
		ApkHash               [24]byte
		UpdateBlockNumber     uint32
		NextUpdateBlockNumber uint32
	}, error)

	CurrentApk(opts *bind.CallOpts, arg0 uint8) (struct {
		X *big.Int
		Y *big.Int
	}, error)

	GetApk(opts *bind.CallOpts, quorumNumber uint8) (BN254G1Point, error)

	GetApkHashAtBlockNumberAndIndex(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error)

	GetApkHistoryLength(opts *bind.CallOpts, quorumNumber uint8) (uint32, error)

	GetApkIndicesAtBlockNumber(opts *bind.CallOpts, quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error)

	GetApkUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, index *big.Int) (IBLSApkRegistryApkUpdate, error)

	GetOperatorFromPubkeyHash(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error)

	GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error)

	GetRegisteredPubkey(opts *bind.CallOpts, operator common.Address) (BN254G1Point, [32]byte, error)

	OperatorToPubkey(opts *bind.CallOpts, arg0 common.Address) (struct {
		X *big.Int
		Y *big.Int
	}, error)

	OperatorToPubkeyHash(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error)

	PubkeyHashToOperator(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error)

	RegistryCoordinator(opts *bind.CallOpts) (common.Address, error)
}

// ContractBLSApkRegistryTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractBLSApkRegistryTransacts interface {
	DeregisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error)

	InitializeQuorum(opts *bind.TransactOpts, quorumNumber uint8) (*types.Transaction, error)

	RegisterBLSPublicKey(opts *bind.TransactOpts, operator common.Address, params IBLSApkRegistryPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error)

	RegisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error)
}

// ContractBLSApkRegistryFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractBLSApkRegistryFilters interface {
	FilterInitialized(opts *bind.FilterOpts) (*ContractBLSApkRegistryInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractBLSApkRegistryInitialized, error)

	FilterNewPubkeyRegistration(opts *bind.FilterOpts, operator []common.Address) (*ContractBLSApkRegistryNewPubkeyRegistrationIterator, error)
	WatchNewPubkeyRegistration(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryNewPubkeyRegistration, operator []common.Address) (event.Subscription, error)
	ParseNewPubkeyRegistration(log types.Log) (*ContractBLSApkRegistryNewPubkeyRegistration, error)

	FilterOperatorAddedToQuorums(opts *bind.FilterOpts) (*ContractBLSApkRegistryOperatorAddedToQuorumsIterator, error)
	WatchOperatorAddedToQuorums(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryOperatorAddedToQuorums) (event.Subscription, error)
	ParseOperatorAddedToQuorums(log types.Log) (*ContractBLSApkRegistryOperatorAddedToQuorums, error)

	FilterOperatorRemovedFromQuorums(opts *bind.FilterOpts) (*ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator, error)
	WatchOperatorRemovedFromQuorums(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryOperatorRemovedFromQuorums) (event.Subscription, error)
	ParseOperatorRemovedFromQuorums(log types.Log) (*ContractBLSApkRegistryOperatorRemovedFromQuorums, error)
}

// ContractBLSApkRegistry is an auto generated Go binding around an Ethereum contract.
type ContractBLSApkRegistry struct {
	ContractBLSApkRegistryCaller     // Read-only binding to the contract
	ContractBLSApkRegistryTransactor // Write-only binding to the contract
	ContractBLSApkRegistryFilterer   // Log filterer for contract events
}

// ContractBLSApkRegistry implements the ContractBLSApkRegistryMethods interface.
var _ ContractBLSApkRegistryMethods = (*ContractBLSApkRegistry)(nil)

// ContractBLSApkRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractBLSApkRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBLSApkRegistryCaller implements the ContractBLSApkRegistryCalls interface.
var _ ContractBLSApkRegistryCalls = (*ContractBLSApkRegistryCaller)(nil)

// ContractBLSApkRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractBLSApkRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBLSApkRegistryTransactor implements the ContractBLSApkRegistryTransacts interface.
var _ ContractBLSApkRegistryTransacts = (*ContractBLSApkRegistryTransactor)(nil)

// ContractBLSApkRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractBLSApkRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBLSApkRegistryFilterer implements the ContractBLSApkRegistryFilters interface.
var _ ContractBLSApkRegistryFilters = (*ContractBLSApkRegistryFilterer)(nil)

// ContractBLSApkRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractBLSApkRegistrySession struct {
	Contract     *ContractBLSApkRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ContractBLSApkRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractBLSApkRegistryCallerSession struct {
	Contract *ContractBLSApkRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ContractBLSApkRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractBLSApkRegistryTransactorSession struct {
	Contract     *ContractBLSApkRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ContractBLSApkRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractBLSApkRegistryRaw struct {
	Contract *ContractBLSApkRegistry // Generic contract binding to access the raw methods on
}

// ContractBLSApkRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractBLSApkRegistryCallerRaw struct {
	Contract *ContractBLSApkRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractBLSApkRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractBLSApkRegistryTransactorRaw struct {
	Contract *ContractBLSApkRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractBLSApkRegistry creates a new instance of ContractBLSApkRegistry, bound to a specific deployed contract.
func NewContractBLSApkRegistry(address common.Address, backend bind.ContractBackend) (*ContractBLSApkRegistry, error) {
	contract, err := bindContractBLSApkRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistry{ContractBLSApkRegistryCaller: ContractBLSApkRegistryCaller{contract: contract}, ContractBLSApkRegistryTransactor: ContractBLSApkRegistryTransactor{contract: contract}, ContractBLSApkRegistryFilterer: ContractBLSApkRegistryFilterer{contract: contract}}, nil
}

// NewContractBLSApkRegistryCaller creates a new read-only instance of ContractBLSApkRegistry, bound to a specific deployed contract.
func NewContractBLSApkRegistryCaller(address common.Address, caller bind.ContractCaller) (*ContractBLSApkRegistryCaller, error) {
	contract, err := bindContractBLSApkRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryCaller{contract: contract}, nil
}

// NewContractBLSApkRegistryTransactor creates a new write-only instance of ContractBLSApkRegistry, bound to a specific deployed contract.
func NewContractBLSApkRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractBLSApkRegistryTransactor, error) {
	contract, err := bindContractBLSApkRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryTransactor{contract: contract}, nil
}

// NewContractBLSApkRegistryFilterer creates a new log filterer instance of ContractBLSApkRegistry, bound to a specific deployed contract.
func NewContractBLSApkRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractBLSApkRegistryFilterer, error) {
	contract, err := bindContractBLSApkRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryFilterer{contract: contract}, nil
}

// bindContractBLSApkRegistry binds a generic wrapper to an already deployed contract.
func bindContractBLSApkRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractBLSApkRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractBLSApkRegistry *ContractBLSApkRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractBLSApkRegistry.Contract.ContractBLSApkRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractBLSApkRegistry *ContractBLSApkRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.ContractBLSApkRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractBLSApkRegistry *ContractBLSApkRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.ContractBLSApkRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractBLSApkRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.contract.Transact(opts, method, params...)
}

// ApkHistory is a free data retrieval call binding the contract method 0x7916cea6.
//
// Solidity: function apkHistory(uint8 , uint256 ) view returns(bytes24 apkHash, uint32 updateBlockNumber, uint32 nextUpdateBlockNumber)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) ApkHistory(opts *bind.CallOpts, arg0 uint8, arg1 *big.Int) (struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "apkHistory", arg0, arg1)

	outstruct := new(struct {
		ApkHash               [24]byte
		UpdateBlockNumber     uint32
		NextUpdateBlockNumber uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ApkHash = *abi.ConvertType(out[0], new([24]byte)).(*[24]byte)
	outstruct.UpdateBlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.NextUpdateBlockNumber = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// ApkHistory is a free data retrieval call binding the contract method 0x7916cea6.
//
// Solidity: function apkHistory(uint8 , uint256 ) view returns(bytes24 apkHash, uint32 updateBlockNumber, uint32 nextUpdateBlockNumber)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) ApkHistory(arg0 uint8, arg1 *big.Int) (struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}, error) {
	return _ContractBLSApkRegistry.Contract.ApkHistory(&_ContractBLSApkRegistry.CallOpts, arg0, arg1)
}

// ApkHistory is a free data retrieval call binding the contract method 0x7916cea6.
//
// Solidity: function apkHistory(uint8 , uint256 ) view returns(bytes24 apkHash, uint32 updateBlockNumber, uint32 nextUpdateBlockNumber)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) ApkHistory(arg0 uint8, arg1 *big.Int) (struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}, error) {
	return _ContractBLSApkRegistry.Contract.ApkHistory(&_ContractBLSApkRegistry.CallOpts, arg0, arg1)
}

// CurrentApk is a free data retrieval call binding the contract method 0xa3db80e2.
//
// Solidity: function currentApk(uint8 ) view returns(uint256 X, uint256 Y)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) CurrentApk(opts *bind.CallOpts, arg0 uint8) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "currentApk", arg0)

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentApk is a free data retrieval call binding the contract method 0xa3db80e2.
//
// Solidity: function currentApk(uint8 ) view returns(uint256 X, uint256 Y)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) CurrentApk(arg0 uint8) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _ContractBLSApkRegistry.Contract.CurrentApk(&_ContractBLSApkRegistry.CallOpts, arg0)
}

// CurrentApk is a free data retrieval call binding the contract method 0xa3db80e2.
//
// Solidity: function currentApk(uint8 ) view returns(uint256 X, uint256 Y)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) CurrentApk(arg0 uint8) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _ContractBLSApkRegistry.Contract.CurrentApk(&_ContractBLSApkRegistry.CallOpts, arg0)
}

// GetApk is a free data retrieval call binding the contract method 0x5f61a884.
//
// Solidity: function getApk(uint8 quorumNumber) view returns((uint256,uint256))
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetApk(opts *bind.CallOpts, quorumNumber uint8) (BN254G1Point, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getApk", quorumNumber)

	if err != nil {
		return *new(BN254G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)

	return out0, err

}

// GetApk is a free data retrieval call binding the contract method 0x5f61a884.
//
// Solidity: function getApk(uint8 quorumNumber) view returns((uint256,uint256))
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetApk(quorumNumber uint8) (BN254G1Point, error) {
	return _ContractBLSApkRegistry.Contract.GetApk(&_ContractBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApk is a free data retrieval call binding the contract method 0x5f61a884.
//
// Solidity: function getApk(uint8 quorumNumber) view returns((uint256,uint256))
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetApk(quorumNumber uint8) (BN254G1Point, error) {
	return _ContractBLSApkRegistry.Contract.GetApk(&_ContractBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApkHashAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0x68bccaac.
//
// Solidity: function getApkHashAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetApkHashAtBlockNumberAndIndex(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getApkHashAtBlockNumberAndIndex", quorumNumber, blockNumber, index)

	if err != nil {
		return *new([24]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([24]byte)).(*[24]byte)

	return out0, err

}

// GetApkHashAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0x68bccaac.
//
// Solidity: function getApkHashAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetApkHashAtBlockNumberAndIndex(quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	return _ContractBLSApkRegistry.Contract.GetApkHashAtBlockNumberAndIndex(&_ContractBLSApkRegistry.CallOpts, quorumNumber, blockNumber, index)
}

// GetApkHashAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0x68bccaac.
//
// Solidity: function getApkHashAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetApkHashAtBlockNumberAndIndex(quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	return _ContractBLSApkRegistry.Contract.GetApkHashAtBlockNumberAndIndex(&_ContractBLSApkRegistry.CallOpts, quorumNumber, blockNumber, index)
}

// GetApkHistoryLength is a free data retrieval call binding the contract method 0x377ed99d.
//
// Solidity: function getApkHistoryLength(uint8 quorumNumber) view returns(uint32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetApkHistoryLength(opts *bind.CallOpts, quorumNumber uint8) (uint32, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getApkHistoryLength", quorumNumber)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetApkHistoryLength is a free data retrieval call binding the contract method 0x377ed99d.
//
// Solidity: function getApkHistoryLength(uint8 quorumNumber) view returns(uint32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetApkHistoryLength(quorumNumber uint8) (uint32, error) {
	return _ContractBLSApkRegistry.Contract.GetApkHistoryLength(&_ContractBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApkHistoryLength is a free data retrieval call binding the contract method 0x377ed99d.
//
// Solidity: function getApkHistoryLength(uint8 quorumNumber) view returns(uint32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetApkHistoryLength(quorumNumber uint8) (uint32, error) {
	return _ContractBLSApkRegistry.Contract.GetApkHistoryLength(&_ContractBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApkIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xd5254a8c.
//
// Solidity: function getApkIndicesAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetApkIndicesAtBlockNumber(opts *bind.CallOpts, quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getApkIndicesAtBlockNumber", quorumNumbers, blockNumber)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetApkIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xd5254a8c.
//
// Solidity: function getApkIndicesAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetApkIndicesAtBlockNumber(quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	return _ContractBLSApkRegistry.Contract.GetApkIndicesAtBlockNumber(&_ContractBLSApkRegistry.CallOpts, quorumNumbers, blockNumber)
}

// GetApkIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xd5254a8c.
//
// Solidity: function getApkIndicesAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetApkIndicesAtBlockNumber(quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	return _ContractBLSApkRegistry.Contract.GetApkIndicesAtBlockNumber(&_ContractBLSApkRegistry.CallOpts, quorumNumbers, blockNumber)
}

// GetApkUpdateAtIndex is a free data retrieval call binding the contract method 0x605747d5.
//
// Solidity: function getApkUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetApkUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, index *big.Int) (IBLSApkRegistryApkUpdate, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getApkUpdateAtIndex", quorumNumber, index)

	if err != nil {
		return *new(IBLSApkRegistryApkUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSApkRegistryApkUpdate)).(*IBLSApkRegistryApkUpdate)

	return out0, err

}

// GetApkUpdateAtIndex is a free data retrieval call binding the contract method 0x605747d5.
//
// Solidity: function getApkUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetApkUpdateAtIndex(quorumNumber uint8, index *big.Int) (IBLSApkRegistryApkUpdate, error) {
	return _ContractBLSApkRegistry.Contract.GetApkUpdateAtIndex(&_ContractBLSApkRegistry.CallOpts, quorumNumber, index)
}

// GetApkUpdateAtIndex is a free data retrieval call binding the contract method 0x605747d5.
//
// Solidity: function getApkUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetApkUpdateAtIndex(quorumNumber uint8, index *big.Int) (IBLSApkRegistryApkUpdate, error) {
	return _ContractBLSApkRegistry.Contract.GetApkUpdateAtIndex(&_ContractBLSApkRegistry.CallOpts, quorumNumber, index)
}

// GetOperatorFromPubkeyHash is a free data retrieval call binding the contract method 0x47b314e8.
//
// Solidity: function getOperatorFromPubkeyHash(bytes32 pubkeyHash) view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetOperatorFromPubkeyHash(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getOperatorFromPubkeyHash", pubkeyHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorFromPubkeyHash is a free data retrieval call binding the contract method 0x47b314e8.
//
// Solidity: function getOperatorFromPubkeyHash(bytes32 pubkeyHash) view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetOperatorFromPubkeyHash(pubkeyHash [32]byte) (common.Address, error) {
	return _ContractBLSApkRegistry.Contract.GetOperatorFromPubkeyHash(&_ContractBLSApkRegistry.CallOpts, pubkeyHash)
}

// GetOperatorFromPubkeyHash is a free data retrieval call binding the contract method 0x47b314e8.
//
// Solidity: function getOperatorFromPubkeyHash(bytes32 pubkeyHash) view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetOperatorFromPubkeyHash(pubkeyHash [32]byte) (common.Address, error) {
	return _ContractBLSApkRegistry.Contract.GetOperatorFromPubkeyHash(&_ContractBLSApkRegistry.CallOpts, pubkeyHash)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getOperatorId", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractBLSApkRegistry.Contract.GetOperatorId(&_ContractBLSApkRegistry.CallOpts, operator)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractBLSApkRegistry.Contract.GetOperatorId(&_ContractBLSApkRegistry.CallOpts, operator)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetRegisteredPubkey(opts *bind.CallOpts, operator common.Address) (BN254G1Point, [32]byte, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getRegisteredPubkey", operator)

	if err != nil {
		return *new(BN254G1Point), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetRegisteredPubkey(operator common.Address) (BN254G1Point, [32]byte, error) {
	return _ContractBLSApkRegistry.Contract.GetRegisteredPubkey(&_ContractBLSApkRegistry.CallOpts, operator)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetRegisteredPubkey(operator common.Address) (BN254G1Point, [32]byte, error) {
	return _ContractBLSApkRegistry.Contract.GetRegisteredPubkey(&_ContractBLSApkRegistry.CallOpts, operator)
}

// OperatorToPubkey is a free data retrieval call binding the contract method 0x00a1f4cb.
//
// Solidity: function operatorToPubkey(address ) view returns(uint256 X, uint256 Y)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) OperatorToPubkey(opts *bind.CallOpts, arg0 common.Address) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "operatorToPubkey", arg0)

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OperatorToPubkey is a free data retrieval call binding the contract method 0x00a1f4cb.
//
// Solidity: function operatorToPubkey(address ) view returns(uint256 X, uint256 Y)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) OperatorToPubkey(arg0 common.Address) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _ContractBLSApkRegistry.Contract.OperatorToPubkey(&_ContractBLSApkRegistry.CallOpts, arg0)
}

// OperatorToPubkey is a free data retrieval call binding the contract method 0x00a1f4cb.
//
// Solidity: function operatorToPubkey(address ) view returns(uint256 X, uint256 Y)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) OperatorToPubkey(arg0 common.Address) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _ContractBLSApkRegistry.Contract.OperatorToPubkey(&_ContractBLSApkRegistry.CallOpts, arg0)
}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address ) view returns(bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) OperatorToPubkeyHash(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "operatorToPubkeyHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address ) view returns(bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) OperatorToPubkeyHash(arg0 common.Address) ([32]byte, error) {
	return _ContractBLSApkRegistry.Contract.OperatorToPubkeyHash(&_ContractBLSApkRegistry.CallOpts, arg0)
}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address ) view returns(bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) OperatorToPubkeyHash(arg0 common.Address) ([32]byte, error) {
	return _ContractBLSApkRegistry.Contract.OperatorToPubkeyHash(&_ContractBLSApkRegistry.CallOpts, arg0)
}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 ) view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) PubkeyHashToOperator(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "pubkeyHashToOperator", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 ) view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) PubkeyHashToOperator(arg0 [32]byte) (common.Address, error) {
	return _ContractBLSApkRegistry.Contract.PubkeyHashToOperator(&_ContractBLSApkRegistry.CallOpts, arg0)
}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 ) view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) PubkeyHashToOperator(arg0 [32]byte) (common.Address, error) {
	return _ContractBLSApkRegistry.Contract.PubkeyHashToOperator(&_ContractBLSApkRegistry.CallOpts, arg0)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) RegistryCoordinator() (common.Address, error) {
	return _ContractBLSApkRegistry.Contract.RegistryCoordinator(&_ContractBLSApkRegistry.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractBLSApkRegistry.Contract.RegistryCoordinator(&_ContractBLSApkRegistry.CallOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.contract.Transact(opts, "deregisterOperator", operator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) DeregisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.DeregisterOperator(&_ContractBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorSession) DeregisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.DeregisterOperator(&_ContractBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactor) InitializeQuorum(opts *bind.TransactOpts, quorumNumber uint8) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.contract.Transact(opts, "initializeQuorum", quorumNumber)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) InitializeQuorum(quorumNumber uint8) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.InitializeQuorum(&_ContractBLSApkRegistry.TransactOpts, quorumNumber)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorSession) InitializeQuorum(quorumNumber uint8) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.InitializeQuorum(&_ContractBLSApkRegistry.TransactOpts, quorumNumber)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactor) RegisterBLSPublicKey(opts *bind.TransactOpts, operator common.Address, params IBLSApkRegistryPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.contract.Transact(opts, "registerBLSPublicKey", operator, params, pubkeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) RegisterBLSPublicKey(operator common.Address, params IBLSApkRegistryPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.RegisterBLSPublicKey(&_ContractBLSApkRegistry.TransactOpts, operator, params, pubkeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorSession) RegisterBLSPublicKey(operator common.Address, params IBLSApkRegistryPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.RegisterBLSPublicKey(&_ContractBLSApkRegistry.TransactOpts, operator, params, pubkeyRegistrationMessageHash)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3fb27952.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.contract.Transact(opts, "registerOperator", operator, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3fb27952.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) RegisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.RegisterOperator(&_ContractBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3fb27952.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorSession) RegisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.RegisterOperator(&_ContractBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// ContractBLSApkRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryInitializedIterator struct {
	Event *ContractBLSApkRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ContractBLSApkRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBLSApkRegistryInitialized)
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
		it.Event = new(ContractBLSApkRegistryInitialized)
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
func (it *ContractBLSApkRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBLSApkRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBLSApkRegistryInitialized represents a Initialized event raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractBLSApkRegistryInitializedIterator, error) {

	logs, sub, err := _ContractBLSApkRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryInitializedIterator{contract: _ContractBLSApkRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractBLSApkRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBLSApkRegistryInitialized)
				if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) ParseInitialized(log types.Log) (*ContractBLSApkRegistryInitialized, error) {
	event := new(ContractBLSApkRegistryInitialized)
	if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBLSApkRegistryNewPubkeyRegistrationIterator is returned from FilterNewPubkeyRegistration and is used to iterate over the raw logs and unpacked data for NewPubkeyRegistration events raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryNewPubkeyRegistrationIterator struct {
	Event *ContractBLSApkRegistryNewPubkeyRegistration // Event containing the contract specifics and raw log

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
func (it *ContractBLSApkRegistryNewPubkeyRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBLSApkRegistryNewPubkeyRegistration)
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
		it.Event = new(ContractBLSApkRegistryNewPubkeyRegistration)
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
func (it *ContractBLSApkRegistryNewPubkeyRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBLSApkRegistryNewPubkeyRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBLSApkRegistryNewPubkeyRegistration represents a NewPubkeyRegistration event raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryNewPubkeyRegistration struct {
	Operator common.Address
	PubkeyG1 BN254G1Point
	PubkeyG2 BN254G2Point
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewPubkeyRegistration is a free log retrieval operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) FilterNewPubkeyRegistration(opts *bind.FilterOpts, operator []common.Address) (*ContractBLSApkRegistryNewPubkeyRegistrationIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractBLSApkRegistry.contract.FilterLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryNewPubkeyRegistrationIterator{contract: _ContractBLSApkRegistry.contract, event: "NewPubkeyRegistration", logs: logs, sub: sub}, nil
}

// WatchNewPubkeyRegistration is a free log subscription operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) WatchNewPubkeyRegistration(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryNewPubkeyRegistration, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractBLSApkRegistry.contract.WatchLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBLSApkRegistryNewPubkeyRegistration)
				if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
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

// ParseNewPubkeyRegistration is a log parse operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) ParseNewPubkeyRegistration(log types.Log) (*ContractBLSApkRegistryNewPubkeyRegistration, error) {
	event := new(ContractBLSApkRegistryNewPubkeyRegistration)
	if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBLSApkRegistryOperatorAddedToQuorumsIterator is returned from FilterOperatorAddedToQuorums and is used to iterate over the raw logs and unpacked data for OperatorAddedToQuorums events raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryOperatorAddedToQuorumsIterator struct {
	Event *ContractBLSApkRegistryOperatorAddedToQuorums // Event containing the contract specifics and raw log

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
func (it *ContractBLSApkRegistryOperatorAddedToQuorumsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBLSApkRegistryOperatorAddedToQuorums)
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
		it.Event = new(ContractBLSApkRegistryOperatorAddedToQuorums)
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
func (it *ContractBLSApkRegistryOperatorAddedToQuorumsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBLSApkRegistryOperatorAddedToQuorumsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBLSApkRegistryOperatorAddedToQuorums represents a OperatorAddedToQuorums event raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryOperatorAddedToQuorums struct {
	Operator      common.Address
	OperatorId    [32]byte
	QuorumNumbers []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToQuorums is a free log retrieval operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) FilterOperatorAddedToQuorums(opts *bind.FilterOpts) (*ContractBLSApkRegistryOperatorAddedToQuorumsIterator, error) {

	logs, sub, err := _ContractBLSApkRegistry.contract.FilterLogs(opts, "OperatorAddedToQuorums")
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryOperatorAddedToQuorumsIterator{contract: _ContractBLSApkRegistry.contract, event: "OperatorAddedToQuorums", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToQuorums is a free log subscription operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) WatchOperatorAddedToQuorums(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryOperatorAddedToQuorums) (event.Subscription, error) {

	logs, sub, err := _ContractBLSApkRegistry.contract.WatchLogs(opts, "OperatorAddedToQuorums")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBLSApkRegistryOperatorAddedToQuorums)
				if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "OperatorAddedToQuorums", log); err != nil {
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

// ParseOperatorAddedToQuorums is a log parse operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) ParseOperatorAddedToQuorums(log types.Log) (*ContractBLSApkRegistryOperatorAddedToQuorums, error) {
	event := new(ContractBLSApkRegistryOperatorAddedToQuorums)
	if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "OperatorAddedToQuorums", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator is returned from FilterOperatorRemovedFromQuorums and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromQuorums events raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator struct {
	Event *ContractBLSApkRegistryOperatorRemovedFromQuorums // Event containing the contract specifics and raw log

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
func (it *ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBLSApkRegistryOperatorRemovedFromQuorums)
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
		it.Event = new(ContractBLSApkRegistryOperatorRemovedFromQuorums)
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
func (it *ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBLSApkRegistryOperatorRemovedFromQuorums represents a OperatorRemovedFromQuorums event raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryOperatorRemovedFromQuorums struct {
	Operator      common.Address
	OperatorId    [32]byte
	QuorumNumbers []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromQuorums is a free log retrieval operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) FilterOperatorRemovedFromQuorums(opts *bind.FilterOpts) (*ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator, error) {

	logs, sub, err := _ContractBLSApkRegistry.contract.FilterLogs(opts, "OperatorRemovedFromQuorums")
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator{contract: _ContractBLSApkRegistry.contract, event: "OperatorRemovedFromQuorums", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromQuorums is a free log subscription operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) WatchOperatorRemovedFromQuorums(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryOperatorRemovedFromQuorums) (event.Subscription, error) {

	logs, sub, err := _ContractBLSApkRegistry.contract.WatchLogs(opts, "OperatorRemovedFromQuorums")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBLSApkRegistryOperatorRemovedFromQuorums)
				if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "OperatorRemovedFromQuorums", log); err != nil {
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

// ParseOperatorRemovedFromQuorums is a log parse operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) ParseOperatorRemovedFromQuorums(log types.Log) (*ContractBLSApkRegistryOperatorRemovedFromQuorums, error) {
	event := new(ContractBLSApkRegistryOperatorRemovedFromQuorums)
	if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "OperatorRemovedFromQuorums", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
