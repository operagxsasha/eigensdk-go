// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractOperatorStateRetriever

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

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractOperatorStateRetrieverMetaData contains all meta data concerning the ContractOperatorStateRetriever contract.
var ContractOperatorStateRetrieverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50611d178061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c806331b36bd9146100645780633563b0d11461008d5780634d2b57fe146100ad5780634f739f74146100cd5780635c155662146100ed578063cefdc1d41461010d575b5f5ffd5b61007761007236600461135c565b61012e565b6040516100849190611449565b60405180910390f35b6100a061009b366004611483565b61023f565b60405161008491906115e4565b6100c06100bb366004611651565b6106a7565b604051610084919061169d565b6100e06100db36600461172f565b6107b1565b6040516100849190611825565b6101006100fb3660046118db565b610eab565b604051610084919061193a565b61012061011b366004611971565b611060565b6040516100849291906119a5565b606081516001600160401b03811115610149576101496112f6565b604051908082528060200260200182016040528015610172578160200160208202803683370190505b5090505f5b825181101561023857836001600160a01b03166313542a4e8483815181106101a1576101a16119c5565b60200260200101516040518263ffffffff1660e01b81526004016101d491906001600160a01b0391909116815260200190565b602060405180830381865afa1580156101ef573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061021391906119d9565b828281518110610225576102256119c5565b6020908102919091010152600101610177565b5092915050565b60605f846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801561027e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102a291906119f0565b90505f856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102e1573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061030591906119f0565b90505f866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610344573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061036891906119f0565b90505f86516001600160401b03811115610384576103846112f6565b6040519080825280602002602001820160405280156103b757816020015b60608152602001906001900390816103a25790505b5090505f5b875181101561069b575f8882815181106103d8576103d86119c5565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291505f906001600160a01b038716906389026245906044015f60405180830381865afa158015610435573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261045c9190810190611a0b565b905080516001600160401b03811115610477576104776112f6565b6040519080825280602002602001820160405280156104c057816020015b604080516060810182525f80825260208083018290529282015282525f199092019101816104955790505b508484815181106104d3576104d36119c5565b60209081029190910101525f5b8151811015610690576040518060600160405280876001600160a01b03166347b314e8858581518110610515576105156119c5565b60200260200101516040518263ffffffff1660e01b815260040161053b91815260200190565b602060405180830381865afa158015610556573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061057a91906119f0565b6001600160a01b0316815260200183838151811061059a5761059a6119c5565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106105c8576105c86119c5565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610622573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106469190611a9b565b6001600160601b0316815250858581518110610664576106646119c5565b6020026020010151828151811061067d5761067d6119c5565b60209081029190910101526001016104e0565b5050506001016103bc565b50979650505050505050565b606081516001600160401b038111156106c2576106c26112f6565b6040519080825280602002602001820160405280156106eb578160200160208202803683370190505b5090505f5b825181101561023857836001600160a01b031663296bb06484838151811061071a5761071a6119c5565b60200260200101516040518263ffffffff1660e01b815260040161074091815260200190565b602060405180830381865afa15801561075b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061077f91906119f0565b828281518110610791576107916119c5565b6001600160a01b03909216602092830291909101909101526001016106f0565b6107dc6040518060800160405280606081526020016060815260200160608152602001606081525090565b5f876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610819573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061083d91906119f0565b905061086a6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e9061089a908b9089908990600401611ac1565b5f60405180830381865afa1580156108b4573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526108db9190810190611b06565b81526040516340e03a8160e11b81526001600160a01b038316906381c075029061090d908b908b908b90600401611bbd565b5f60405180830381865afa158015610927573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261094e9190810190611b06565b6040820152856001600160401b0381111561096b5761096b6112f6565b60405190808252806020026020018201604052801561099e57816020015b60608152602001906001900390816109895790505b5060608201525f5b60ff8116871115610dc3575f856001600160401b038111156109ca576109ca6112f6565b6040519080825280602002602001820160405280156109f3578160200160208202803683370190505b5083606001518360ff1681518110610a0d57610a0d6119c5565b60209081029190910101525f5b86811015610ccf575f8c6001600160a01b03166304ec63518a8a85818110610a4457610a446119c5565b905060200201358e885f01518681518110610a6157610a616119c5565b60200260200101516040518463ffffffff1660e01b8152600401610a9e9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015610ab9573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610add9190611be5565b9050806001600160c01b03165f03610b875760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a40160405180910390fd5b8a8a8560ff16818110610b9c57610b9c6119c5565b60016001600160c01b038516919093013560f81c1c82169091039050610cc657856001600160a01b031663dd9846b98a8a85818110610bdd57610bdd6119c5565b905060200201358d8d8860ff16818110610bf957610bf96119c5565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015610c4d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c719190611c0b565b85606001518560ff1681518110610c8a57610c8a6119c5565b60200260200101518481518110610ca357610ca36119c5565b63ffffffff9092166020928302919091019091015282610cc281611c3a565b9350505b50600101610a1a565b505f816001600160401b03811115610ce957610ce96112f6565b604051908082528060200260200182016040528015610d12578160200160208202803683370190505b5090505f5b82811015610d885784606001518460ff1681518110610d3857610d386119c5565b60200260200101518181518110610d5157610d516119c5565b6020026020010151828281518110610d6b57610d6b6119c5565b63ffffffff90921660209283029190910190910152600101610d17565b508084606001518460ff1681518110610da357610da36119c5565b602002602001018190525050508080610dbb90611c52565b9150506109a6565b505f896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e01573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e2591906119f0565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90610e58908b908b908e90600401611c70565b5f60405180830381865afa158015610e72573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610e999190810190611b06565b60208301525098975050505050505050565b60605f846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401610edc929190611c99565b5f60405180830381865afa158015610ef6573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610f1d9190810190611b06565b90505f84516001600160401b03811115610f3957610f396112f6565b604051908082528060200260200182016040528015610f62578160200160208202803683370190505b5090505f5b855181101561105657866001600160a01b03166304ec6351878381518110610f9157610f916119c5565b602002602001015187868581518110610fac57610fac6119c5565b60200260200101516040518463ffffffff1660e01b8152600401610fe99392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611004573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110289190611be5565b6001600160c01b0316828281518110611043576110436119c5565b6020908102919091010152600101610f67565b5095945050505050565b6040805160018082528183019092525f9160609183916020808301908036833701905050905084815f81518110611099576110996119c5565b60209081029190910101526040516361c8a12f60e11b81525f906001600160a01b0388169063c391425e906110d49088908690600401611c99565b5f60405180830381865afa1580156110ee573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526111159190810190611b06565b5f81518110611126576111266119c5565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291505f906001600160a01b038916906304ec635190606401602060405180830381865afa15801561118f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111b39190611be5565b6001600160c01b031690505f6111c8826111e6565b9050816111d68a838a61023f565b9550955050505050935093915050565b60605f5f6111f3846112af565b61ffff166001600160401b0381111561120e5761120e6112f6565b6040519080825280601f01601f191660200182016040528015611238576020820181803683370190505b5090505f805b82518210801561124f575061010081105b156112a5576001811b935085841615611295578060f81b838381518110611278576112786119c5565b60200101906001600160f81b03191690815f1a9053508160010191505b61129e81611c3a565b905061123e565b5090949350505050565b5f805b82156112d9576112c3600184611cb7565b90921691806112d181611cca565b9150506112b2565b92915050565b6001600160a01b03811681146112f3575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715611332576113326112f6565b604052919050565b5f6001600160401b03821115611352576113526112f6565b5060051b60200190565b5f5f6040838503121561136d575f5ffd5b8235611378816112df565b915060208301356001600160401b03811115611392575f5ffd5b8301601f810185136113a2575f5ffd5b80356113b56113b08261133a565b61130a565b8082825260208201915060208360051b8501019250878311156113d6575f5ffd5b6020840193505b828410156114015783356113f0816112df565b8252602093840193909101906113dd565b809450505050509250929050565b5f8151808452602084019350602083015f5b8281101561143f578151865260209586019590910190600101611421565b5093949350505050565b602081525f61145b602083018461140f565b9392505050565b63ffffffff811681146112f3575f5ffd5b803561147e81611462565b919050565b5f5f5f60608486031215611495575f5ffd5b83356114a0816112df565b925060208401356001600160401b038111156114ba575f5ffd5b8401601f810186136114ca575f5ffd5b80356001600160401b038111156114e3576114e36112f6565b6114f6601f8201601f191660200161130a565b81815287602083850101111561150a575f5ffd5b816020840160208301375f6020838301015280945050505061152e60408501611473565b90509250925092565b5f82825180855260208501945060208160051b830101602085015f5b838110156115d857848303601f19018852815180518085526020918201918501905f5b818110156115bf57835180516001600160a01b03168452602080820151818601526040918201516001600160601b03169185019190915290930192606090920191600101611576565b50506020998a0199909450929092019150600101611553565b50909695505050505050565b602081525f61145b6020830184611537565b5f82601f830112611605575f5ffd5b81356116136113b08261133a565b8082825260208201915060208360051b860101925085831115611634575f5ffd5b602085015b83811015611056578035835260209283019201611639565b5f5f60408385031215611662575f5ffd5b823561166d816112df565b915060208301356001600160401b03811115611687575f5ffd5b611693858286016115f6565b9150509250929050565b602080825282518282018190525f918401906040840190835b818110156116dd5783516001600160a01b03168352602093840193909201916001016116b6565b509095945050505050565b5f5f83601f8401126116f8575f5ffd5b5081356001600160401b0381111561170e575f5ffd5b6020830191508360208260051b8501011115611728575f5ffd5b9250929050565b5f5f5f5f5f5f60808789031215611744575f5ffd5b863561174f816112df565b9550602087013561175f81611462565b945060408701356001600160401b03811115611779575f5ffd5b8701601f81018913611789575f5ffd5b80356001600160401b0381111561179e575f5ffd5b8960208284010111156117af575f5ffd5b6020919091019450925060608701356001600160401b038111156117d1575f5ffd5b6117dd89828a016116e8565b979a9699509497509295939492505050565b5f8151808452602084019350602083015f5b8281101561143f57815163ffffffff16865260209586019590910190600101611801565b602081525f82516080602084015261184060a08401826117ef565b90506020840151601f1984830301604085015261185d82826117ef565b9150506040840151601f1984830301606085015261187b82826117ef565b6060860151858203601f190160808701528051808352919350602090810192508084019190600582901b8501015f5b8281101561069b57601f198683030184526118c68286516117ef565b602095860195949094019391506001016118aa565b5f5f5f606084860312156118ed575f5ffd5b83356118f8816112df565b925060208401356001600160401b03811115611912575f5ffd5b61191e868287016115f6565b925050604084013561192f81611462565b809150509250925092565b602080825282518282018190525f918401906040840190835b818110156116dd578351835260209384019390920191600101611953565b5f5f5f60608486031215611983575f5ffd5b833561198e816112df565b925060208401359150604084013561192f81611462565b828152604060208201525f6119bd6040830184611537565b949350505050565b634e487b7160e01b5f52603260045260245ffd5b5f602082840312156119e9575f5ffd5b5051919050565b5f60208284031215611a00575f5ffd5b815161145b816112df565b5f60208284031215611a1b575f5ffd5b81516001600160401b03811115611a30575f5ffd5b8201601f81018413611a40575f5ffd5b8051611a4e6113b08261133a565b8082825260208201915060208360051b850101925086831115611a6f575f5ffd5b6020840193505b82841015611a91578351825260209384019390910190611a76565b9695505050505050565b5f60208284031215611aab575f5ffd5b81516001600160601b038116811461145b575f5ffd5b63ffffffff8416815260406020820181905281018290525f6001600160fb1b03831115611aec575f5ffd5b8260051b8085606085013791909101606001949350505050565b5f60208284031215611b16575f5ffd5b81516001600160401b03811115611b2b575f5ffd5b8201601f81018413611b3b575f5ffd5b8051611b496113b08261133a565b8082825260208201915060208360051b850101925086831115611b6a575f5ffd5b6020840193505b82841015611a91578351611b8481611462565b825260209384019390910190611b71565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201525f611bdc604083018486611b95565b95945050505050565b5f60208284031215611bf5575f5ffd5b81516001600160c01b038116811461145b575f5ffd5b5f60208284031215611c1b575f5ffd5b815161145b81611462565b634e487b7160e01b5f52601160045260245ffd5b5f60018201611c4b57611c4b611c26565b5060010190565b5f60ff821660ff8103611c6757611c67611c26565b60010192915050565b604081525f611c83604083018587611b95565b905063ffffffff83166020830152949350505050565b63ffffffff83168152604060208201525f6119bd604083018461140f565b818103818111156112d9576112d9611c26565b5f61ffff821661ffff8103611c6757611c67611c2656fea26469706673582212201706ed58f23264aa591e6d38cc031666c60d696a4a550002a8b6177be300d7f864736f6c634300081b0033",
}

// ContractOperatorStateRetrieverABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractOperatorStateRetrieverMetaData.ABI instead.
var ContractOperatorStateRetrieverABI = ContractOperatorStateRetrieverMetaData.ABI

// ContractOperatorStateRetrieverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractOperatorStateRetrieverMetaData.Bin instead.
var ContractOperatorStateRetrieverBin = ContractOperatorStateRetrieverMetaData.Bin

// DeployContractOperatorStateRetriever deploys a new Ethereum contract, binding an instance of ContractOperatorStateRetriever to it.
func DeployContractOperatorStateRetriever(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractOperatorStateRetriever, error) {
	parsed, err := ContractOperatorStateRetrieverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractOperatorStateRetrieverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractOperatorStateRetriever{ContractOperatorStateRetrieverCaller: ContractOperatorStateRetrieverCaller{contract: contract}, ContractOperatorStateRetrieverTransactor: ContractOperatorStateRetrieverTransactor{contract: contract}, ContractOperatorStateRetrieverFilterer: ContractOperatorStateRetrieverFilterer{contract: contract}}, nil
}

// ContractOperatorStateRetrieverMethods is an auto generated interface around an Ethereum contract.
type ContractOperatorStateRetrieverMethods interface {
	ContractOperatorStateRetrieverCalls
	ContractOperatorStateRetrieverTransacts
	ContractOperatorStateRetrieverFilters
}

// ContractOperatorStateRetrieverCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractOperatorStateRetrieverCalls interface {
	GetBatchOperatorFromId(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error)

	GetBatchOperatorId(opts *bind.CallOpts, registryCoordinator common.Address, operators []common.Address) ([][32]byte, error)

	GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error)

	GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error)

	GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error)

	GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error)
}

// ContractOperatorStateRetrieverTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractOperatorStateRetrieverTransacts interface {
}

// ContractOperatorStateRetrieverFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractOperatorStateRetrieverFilters interface {
}

// ContractOperatorStateRetriever is an auto generated Go binding around an Ethereum contract.
type ContractOperatorStateRetriever struct {
	ContractOperatorStateRetrieverCaller     // Read-only binding to the contract
	ContractOperatorStateRetrieverTransactor // Write-only binding to the contract
	ContractOperatorStateRetrieverFilterer   // Log filterer for contract events
}

// ContractOperatorStateRetriever implements the ContractOperatorStateRetrieverMethods interface.
var _ ContractOperatorStateRetrieverMethods = (*ContractOperatorStateRetriever)(nil)

// ContractOperatorStateRetrieverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOperatorStateRetrieverCaller implements the ContractOperatorStateRetrieverCalls interface.
var _ ContractOperatorStateRetrieverCalls = (*ContractOperatorStateRetrieverCaller)(nil)

// ContractOperatorStateRetrieverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOperatorStateRetrieverTransactor implements the ContractOperatorStateRetrieverTransacts interface.
var _ ContractOperatorStateRetrieverTransacts = (*ContractOperatorStateRetrieverTransactor)(nil)

// ContractOperatorStateRetrieverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractOperatorStateRetrieverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOperatorStateRetrieverFilterer implements the ContractOperatorStateRetrieverFilters interface.
var _ ContractOperatorStateRetrieverFilters = (*ContractOperatorStateRetrieverFilterer)(nil)

// ContractOperatorStateRetrieverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractOperatorStateRetrieverSession struct {
	Contract     *ContractOperatorStateRetriever // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ContractOperatorStateRetrieverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractOperatorStateRetrieverCallerSession struct {
	Contract *ContractOperatorStateRetrieverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// ContractOperatorStateRetrieverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractOperatorStateRetrieverTransactorSession struct {
	Contract     *ContractOperatorStateRetrieverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// ContractOperatorStateRetrieverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverRaw struct {
	Contract *ContractOperatorStateRetriever // Generic contract binding to access the raw methods on
}

// ContractOperatorStateRetrieverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverCallerRaw struct {
	Contract *ContractOperatorStateRetrieverCaller // Generic read-only contract binding to access the raw methods on
}

// ContractOperatorStateRetrieverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverTransactorRaw struct {
	Contract *ContractOperatorStateRetrieverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractOperatorStateRetriever creates a new instance of ContractOperatorStateRetriever, bound to a specific deployed contract.
func NewContractOperatorStateRetriever(address common.Address, backend bind.ContractBackend) (*ContractOperatorStateRetriever, error) {
	contract, err := bindContractOperatorStateRetriever(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorStateRetriever{ContractOperatorStateRetrieverCaller: ContractOperatorStateRetrieverCaller{contract: contract}, ContractOperatorStateRetrieverTransactor: ContractOperatorStateRetrieverTransactor{contract: contract}, ContractOperatorStateRetrieverFilterer: ContractOperatorStateRetrieverFilterer{contract: contract}}, nil
}

// NewContractOperatorStateRetrieverCaller creates a new read-only instance of ContractOperatorStateRetriever, bound to a specific deployed contract.
func NewContractOperatorStateRetrieverCaller(address common.Address, caller bind.ContractCaller) (*ContractOperatorStateRetrieverCaller, error) {
	contract, err := bindContractOperatorStateRetriever(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorStateRetrieverCaller{contract: contract}, nil
}

// NewContractOperatorStateRetrieverTransactor creates a new write-only instance of ContractOperatorStateRetriever, bound to a specific deployed contract.
func NewContractOperatorStateRetrieverTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractOperatorStateRetrieverTransactor, error) {
	contract, err := bindContractOperatorStateRetriever(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorStateRetrieverTransactor{contract: contract}, nil
}

// NewContractOperatorStateRetrieverFilterer creates a new log filterer instance of ContractOperatorStateRetriever, bound to a specific deployed contract.
func NewContractOperatorStateRetrieverFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractOperatorStateRetrieverFilterer, error) {
	contract, err := bindContractOperatorStateRetriever(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorStateRetrieverFilterer{contract: contract}, nil
}

// bindContractOperatorStateRetriever binds a generic wrapper to an already deployed contract.
func bindContractOperatorStateRetriever(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractOperatorStateRetrieverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOperatorStateRetriever.Contract.ContractOperatorStateRetrieverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOperatorStateRetriever.Contract.ContractOperatorStateRetrieverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOperatorStateRetriever.Contract.ContractOperatorStateRetrieverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOperatorStateRetriever.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOperatorStateRetriever.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOperatorStateRetriever.Contract.contract.Transact(opts, method, params...)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetBatchOperatorFromId(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getBatchOperatorFromId", registryCoordinator, operatorIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractOperatorStateRetriever.Contract.GetBatchOperatorFromId(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractOperatorStateRetriever.Contract.GetBatchOperatorFromId(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetBatchOperatorId(opts *bind.CallOpts, registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getBatchOperatorId", registryCoordinator, operators)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractOperatorStateRetriever.Contract.GetBatchOperatorId(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operators)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractOperatorStateRetriever.Contract.GetBatchOperatorId(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operators)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractOperatorStateRetriever.Contract.GetCheckSignaturesIndices(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractOperatorStateRetriever.Contract.GetCheckSignaturesIndices(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractOperatorStateRetriever.Contract.GetOperatorState(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractOperatorStateRetriever.Contract.GetOperatorState(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractOperatorStateRetriever.Contract.GetOperatorState0(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractOperatorStateRetriever.Contract.GetOperatorState0(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractOperatorStateRetriever.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractOperatorStateRetriever.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorIds, blockNumber)
}
