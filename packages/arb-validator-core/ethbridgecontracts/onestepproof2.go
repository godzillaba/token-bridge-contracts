// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgecontracts

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OneStepProof2ABI is the input ABI used to generate the binding from.
const OneStepProof2ABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"inboxAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"messagesAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"logsAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bproof\",\"type\":\"bytes\"}],\"name\":\"executeStep\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[5]\",\"name\":\"fields\",\"type\":\"bytes32[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OneStepProof2FuncSigs maps the 4-byte function signature to its string representation.
var OneStepProof2FuncSigs = map[string]string{
	"1041c884": "executeStep(bytes32,bytes32,bytes32,bytes,bytes)",
}

// OneStepProof2Bin is the compiled bytecode used for deploying new contracts.
var OneStepProof2Bin = "0x608060405234801561001057600080fd5b506131fd806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80631041c88414610030575b600080fd5b610105600480360360a081101561004657600080fd5b8135916020810135916040820135919081019060808101606082013564010000000081111561007457600080fd5b82018360208201111561008657600080fd5b803590602001918460018302840111640100000000831117156100a857600080fd5b9193909290916020810190356401000000008111156100c657600080fd5b8201836020820111156100d857600080fd5b803590602001918460018302840111640100000000831117156100fa57600080fd5b50909250905061014e565b60405167ffffffffffffffff83168152602081018260a080838360005b8381101561013a578181015183820152602001610122565b505050509050019250505060405180910390f35b6000610158612fe4565b610160613002565b6101d68a8a8a8a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8e018190048102820181019092528c815292508c91508b90819084018382808284376000920191909152506101fb92505050565b90506101e1816106c1565b6101ea81610a29565b925092505097509795505050505050565b610203613002565b60008360008151811061021257fe5b602001015160f81c60f81b60f81c905060008460018151811061023157fe5b602001015160f81c60f81b60f81c905060006002905060608360040160ff1660405190808252806020026020018201604052801561028957816020015b61027661309b565b81526020019060019003908161026e5790505b50905060608360040160ff166040519080825280602002602001820160405280156102ce57816020015b6102bb61309b565b8152602001906001900390816102b35790505b50905060005b8560ff1681101561030c576102e98985610a8c565b84518590849081106102f757fe5b602090810291909101015293506001016102d4565b5060005b8460ff16811015610348576103258985610a8c565b835184908490811061033357fe5b60209081029190910101529350600101610310565b506103516130d8565b61035b8985610c4e565b9094509050610368613002565b604051806101e0016040528083815260200161038384610cff565b81526020018e81526020018d81526020018c8152602001600067ffffffffffffffff1681526020016103b3610d74565b81526020016000801b815260200160405180604001604052808a60ff16815260200187815250815260200160405180604001604052808960ff1681526020018681525081526020018b878151811061040757fe5b602001015160f81c60f81b60f81c60ff16600114151581526020018b876001018151811061043157fe5b0160209081015160f81c825281018c90526002870160408201526060018a90528a519091506000908b908790811061046557fe5b602001015160f81c60f81b60f81c905060008b876001018151811061048657fe5b01602001516002979097019660f81c905060ff821615806104aa57508160ff166001145b6040518060400160405280600b81526020016a04241445f494d4d5f5459560ac1b815250906105575760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561051c578181015183820152602001610504565b50505050905090810190601f1680156105495780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5061056061309b565b60ff831661057d57835151610576908390610dbb565b905061061d565b6000875111604051806040016040528060068152602001654e4f5f494d4d60d01b815250906105ed5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561051c578181015183820152602001610504565b5061061a828560000151600001518960018e0360ff168151811061060d57fe5b6020026020010151610e1f565b90505b61062681610ea5565b84515260005b838b0360ff1681101561066b5761066388828151811061064857fe5b60200260200101518660000151610fe990919063ffffffff16565b60010161062c565b5060005b8960ff168110156106ac576106a487828151811061068957fe5b6020026020010151866000015161100390919063ffffffff16565b60010161066f565b50929f9e505050505050505050505050505050565b60008060006131436106da85610160015160ff1661101d565b93509350935093506106ec858361110b565b156106fa5750505050610a26565b610100850151518411156107af57610718610713610d74565b610ea5565b610729866020015160200151610ea5565b146040518060400160405280600d81526020016c535441434b5f4d495353494e4760981b8152509061079c5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561051c578181015183820152602001610504565b506107a685611177565b50505050610a26565b6101208501515183111561084a576107c8610713610d74565b6107d9866020015160400151610ea5565b146040518060400160405280600b81526020016a4155585f4d495353494e4760a81b8152509061079c5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561051c578181015183820152602001610504565b600084118061085c5750846101400151155b801561086d57506101008501515184145b8061089557508461014001518015610883575083155b80156108955750610100850151516001145b6040518060400160405280600a815260200169535441434b5f4d414e5960b01b815250906109045760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561051c578181015183820152602001610504565b50610120850151516040805180820190915260088152674155585f4d414e5960c01b60208201529084146109795760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561051c578181015183820152602001610504565b50610987858263ffffffff16565b60005b610100860151518110156109d3576109cb8661010001516020015182815181106109b057fe5b60200260200101518760200151610fe990919063ffffffff16565b60010161098a565b5060005b61012086015151811015610a2057610a188661012001516020015182815181106109fd57fe5b6020026020010151876020015161100390919063ffffffff16565b6001016109d7565b50505050505b50565b6000610a33612fe4565b8260a001516040518060a00160405280610a5086600001516111e0565b8152602001610a6286602001516111e0565b81526020018560400151815260200185606001518152602001856080015181525091509150915091565b6000610a9661309b565b83518310610adc576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b600080610ae986866112ba565b91509150610af56112e1565b60ff168160ff161415610b29576000610b0e87846112e6565b909350905082610b1d8261135a565b94509450505050610c47565b610b31611413565b60ff168160ff161415610b5357610b488683611418565b935093505050610c47565b610b5b6114ba565b60ff168160ff161415610b83576000610b7487846112e6565b909350905082610b1d826114bf565b610b8b61157c565b60ff168160ff161415610ba257610b488683611581565b610baa611615565b60ff168160ff1610158015610bcb5750610bc261161a565b60ff168160ff16105b15610c07576000610bda611615565b820390506060610beb82898661161f565b909450905083610bfa826116b8565b9550955050505050610c47565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b9250929050565b6000610c586130d8565b610c606130d8565b6000610100820181905280610c7587876112e6565b9096509150610c848787611581565b60208501529550610c958787611581565b60408501529550610ca68787610a8c565b60608501529550610cb78787610a8c565b60808501529550610cc887876112e6565b60a08501529550610cd987876112e6565b9096509050610ce88787610a8c565b60e085015291835260c08301529590945092505050565b610d076130d8565b60405180610120016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c0015181526020018360e00151815260200183610100015181525090505b919050565b610d7c61309b565b60408051600080825260208201909252610db691610db0565b610d9d61309b565b815260200190600190039081610d955790505b506116b8565b905090565b610dc361309b565b6040805160608101825260ff851681526020808201859052825160008082529181018452610e1693830191610e0e565b610dfb61309b565b815260200190600190039081610df35790505b5090526117d1565b90505b92915050565b610e2761309b565b604080516001808252818301909252606091816020015b610e4661309b565b815260200190600190039081610e3e5790505090508281600081518110610e6957fe5b6020026020010181905250610e9a60405180606001604052808760ff168152602001868152602001838152506117d1565b9150505b9392505050565b6000610eaf6112e1565b60ff16826080015160ff161415610ed2578151610ecb9061183f565b9050610d6f565b610eda611413565b60ff16826080015160ff161415610ef857610ecb8260200151611863565b610f0061157c565b60ff16826080015160ff161415610f2257815160a0830151610ecb9190611960565b610f2a611615565b60ff16826080015160ff161415610f6357610f4361309b565b610f5083604001516119b1565b9050610f5b81610ea5565b915050610d6f565b610f6b611b13565b60ff16826080015160ff161415610f8457508051610d6f565b610f8c6114ba565b60ff16826080015160ff161415610fa857506060810151610d6f565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b610ff7826020015182611b18565b82602001819052505050565b611011826040015182611b18565b82604001819052505050565b6000808061314360a085141561104157506001925060009150829050611b96611104565b60a185141561105f57506002925060009150600a9050611c0e611104565b60a285141561107d57506002925060009150600a9050611caa611104565b60a385141561109b57506002925060009150600a9050611d7e611104565b60a48514156110b957506003925060009150600a9050611e52611104565b60a58514156110d757506003925060009150600a9050611f1e611104565b60a68514156110f557506003925060009150600a9050611fc7611104565b50600092508291508190506120705b9193509193565b60a0808301805167ffffffffffffffff90840181169091526020840151909101516000918316111561115657602083015160001960a09091015261114e83611177565b506001610e19565b50602082015160a001805167ffffffffffffffff8316900390526000610e19565b60408051600160f81b6020808301919091526000602183018190526022808401919091528351808403909101815260429092019092528051908201209082015160c0015114156111d3576111ce8160200151612079565b610a26565b6020015160c08101519052565b6000600282610100015114156111f857506000610d6f565b6001826101000151141561120e57506001610d6f565b8151602083015161121e90610ea5565b61122b8460400151610ea5565b6112388560600151610ea5565b6112458660800151610ea5565b8660a001518760c0015161125c8960e00151610ea5565b6040516020018089815260200188815260200187815260200186815260200185815260200184815260200183815260200182815260200198505050505050505050604051602081830303815290604052805190602001209050610d6f565b600080826001018484815181106112cd57fe5b016020015190925060f81c90509250929050565b600090565b600080828451101580156112fe575060208385510310155b61133b576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b6020830161134f858563ffffffff61208416565b915091509250929050565b61136261309b565b6040805160c08101825283815281516060810183526000808252602082810182905284518281528082018652939490850193908301916113b8565b6113a561309b565b81526020019060019003908161139d5790505b509052815260408051600080825260208281019093529190920191906113f4565b6113e161309b565b8152602001906001900390816113d95790505b5081526000602082018190526040820152600160609091015292915050565b600190565b600061142261309b565b8260008061142e61309b565b600061143a89866112ba565b909550935061144989866112ba565b9095509250600160ff8516141561146a576114648986610a8c565b90955091505b61147489866120dd565b9095509050600160ff8516141561149f5784611491848385610e1f565b965096505050505050610c47565b846114aa8483610dbb565b9650965050505050509250929050565b600c90565b6114c761309b565b6040805160c0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190611520565b61150d61309b565b8152602001906001900390816115055790505b5090528152604080516000808252602082810190935291909201919061155c565b61154961309b565b8152602001906001900390816115415790505b50815260208101849052600c604082015260016060909101529050919050565b600290565b600061158b61309b565b828451101580156115a0575060408385510310155b6115dc576040805162461bcd60e51b81526020600482015260086024820152671d1bc81cda1bdc9d60c21b604482015290519081900360640190fd5b6000806115e986866120dd565b90945091506115f886856112e6565b90945090508361160883836120f4565b9350935050509250929050565b600390565b600d90565b60006060600083905060608660ff1660405190808252806020026020018201604052801561166757816020015b61165461309b565b81526020019060019003908161164c5790505b50905060005b8760ff168160ff1610156116ab576116858784610a8c565b8351849060ff851690811061169657fe5b6020908102919091010152925060010161166d565b5090969095509350505050565b6116c061309b565b6116ca82516121ac565b61171b576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b83518110156117525783818151811061173557fe5b602002602001015160a00151820191508080600101915050611720565b506040805160c08101825260008082528251606081018452818152602081810183905284518381528082018652939490850193919290830191906117ac565b61179961309b565b8152602001906001900390816117915790505b5090528152602081019490945260006040850152600360608501526080909301525090565b6117d961309b565b6040805160c0810182526000808252602080830186905283518281529081018452919283019190611820565b61180d61309b565b8152602001906001900390816118055790505b5081526000602082015260016040820181905260609091015292915050565b60408051602080820193909352815180820384018152908201909152805191012090565b600060028260400151511061187457fe5b6040820151516118d957611886611413565b8251602080850151604080516001600160f81b031960f896871b8116828601529490951b909316602185015260228085019190915282518085039091018152604290930190915281519101209050610d6f565b6118e1611413565b826000015161190784604001516000815181106118fa57fe5b6020026020010151610ea5565b8460200151604051602001808560ff1660ff1660f81b81526001018460ff1660ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b600061196a611615565b8383604051602001808460ff1660ff1660f81b8152600101838152602001828152602001935050505060405160208183030381529060405280519060200120905092915050565b6119b961309b565b600882511115611a07576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015611a34578160200160208202803883390190505b508051909150600160005b82811015611a9757611a568682815181106118fa57fe5b848281518110611a6257fe5b602002602001018181525050858181518110611a7a57fe5b602002602001015160a00151820191508080600101915050611a3f565b506000835184604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015611adc578181015183820152602001611ac4565b5050505090500192505050604051602081830303815290604052805190602001209050611b0981836120f4565b9695505050505050565b606490565b611b2061309b565b6040805160028082526060828101909352816020015b611b3e61309b565b815260200190600190039081611b365790505090508281600081518110611b6157fe5b60200260200101819052508381600181518110611b7a57fe5b6020026020010181905250611b8e816119b1565b949350505050565b611b9e61309b565b611bac8261010001516121b3565b9050611bb7816121f5565b611bca57611bc482612200565b50610a26565b611c0a826101000151611c056000801b60405160200180828152602001915050604051602081830303815290604052805190602001206114bf565b61221f565b5050565b611c1661309b565b611c248261010001516121b3565b9050611c2e61309b565b611c3c8361010001516121b3565b9050611c4781612249565b1580611c595750611c5782612267565b155b15611c6e57611c6783612200565b5050610a26565b6000611c9083606001518360000151611c8b876101c00151612274565b61236a565b9050611ca4846101000151611c058361135a565b50505050565b611cb261309b565b611cc08261010001516121b3565b9050611cca61309b565b611cd88361010001516121b3565b9050611ce381612249565b1580611cf55750611cf382612267565b155b15611d0357611c6783612200565b8151600160401b11611d5c576040805162461bcd60e51b815260206004820152601b60248201527f62756666657220696e646578206d7573742062652036342d6269740000000000604482015290519081900360640190fd5b6000611c9083606001518360000151611d79876101c00151612274565b61238c565b611d8661309b565b611d948261010001516121b3565b9050611d9e61309b565b611dac8361010001516121b3565b9050611db781612249565b1580611dc95750611dc782612267565b155b15611dd757611c6783612200565b8151600160401b11611e30576040805162461bcd60e51b815260206004820152601b60248201527f62756666657220696e646578206d7573742062652036342d6269740000000000604482015290519081900360640190fd5b6000611c9083606001518360000151611e4d876101c00151612274565b6124eb565b611e5a61309b565b611e688261010001516121b3565b9050611e7261309b565b611e808361010001516121b3565b9050611e8a61309b565b611e988461010001516121b3565b9050611ea382612249565b1580611eb55750611eb3816121f5565b155b80611ec65750611ec483612267565b155b15611edc57611ed484612200565b505050610a26565b6000611f03846060015184600001518460000151611efe896101c00151612274565b61261e565b9050611f17856101000151611c05836114bf565b5050505050565b611f2661309b565b611f348261010001516121b3565b9050611f3e61309b565b611f4c8361010001516121b3565b9050611f5661309b565b611f648461010001516121b3565b9050611f6f82612249565b1580611f815750611f7f816121f5565b155b80611f925750611f9083612267565b155b15611fa057611ed484612200565b6000611f03846060015184600001518460000151611fc2896101c00151612274565b612667565b611fcf61309b565b611fdd8261010001516121b3565b9050611fe761309b565b611ff58361010001516121b3565b9050611fff61309b565b61200d8461010001516121b3565b905061201882612249565b158061202a5750612028816121f5565b155b8061203b575061203983612267565b155b1561204957611ed484612200565b6000611f0384606001518460000151846000015161206b896101c00151612274565b6127af565b610a2681612200565b600161010090910152565b600081602001835110156120d4576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b6000806020830161134f858563ffffffff61208416565b6120fc61309b565b6040805160c0810182528481528151606081018352600080825260208281018290528451828152808201865293949085019390830191612152565b61213f61309b565b8152602001906001900390816121375790505b5090528152604080516000808252602082810190935291909201919061218e565b61217b61309b565b8152602001906001900390816121735790505b50815260006020820152600260408201526060019290925250919050565b6008101590565b6121bb61309b565b6121c361309b565b82602001516001846000015103815181106121da57fe5b60209081029190910101518351600019018452915050919050565b6080015160ff161590565b61220981611177565b6101008101516000908190526101209091015152565b80826020015183600001518151811061223457fe5b60209081029190910101525080516001019052565b608081015160009060ff16158015610e1957505051600160401b1190565b6080015160ff16600c1490565b61227c613145565b60606122bb838460008151811061228f57fe5b602001015160f81c60f81b856001815181106122a757fe5b01602001516001600160f81b031916612880565b905060606122e884856001815181106122d057fe5b602001015160f81c60f81b866002815181106122a757fe5b9050606061231585866002815181106122fd57fe5b602001015160f81c60f81b876003815181106122a757fe5b90506060612342868760038151811061232a57fe5b602001015160f81c60f81b886004815181106122a757fe5b6040805160808101825295865260208601949094529284019190915250606082015292915050565b6000611b8e612382856020865b048560000151612913565b6020855b06612a81565b6040805160088082528183019092526000916060919060208201818038833901905050905060006123c6866020875b048660000151612913565b90506020808606600801106124995760006123ed876020885b046001018760400151612913565b905060005b6018601f88166008030181101561244157612413838260208a5b0601612a81565b60f81b84828151811061242257fe5b60200101906001600160f81b031916908160001a9053506001016123f2565b506018601f8716600803015b600881101561249257612464826020898401612386565b60f81b84828151811061247357fe5b60200101906001600160f81b031916908160001a90535060010161244d565b50506124e2565b60005b60088110156124e0576124b2828260208961240c565b60f81b8382815181106124c157fe5b60200101906001600160f81b031916908160001a90535060010161249c565b505b611b0982612a8e565b60408051602080825281830190925260009160609190602082018180388339019050509050600061251e866020876123bb565b90506020808606602001106125d757600061253b876020886123df565b905060005b601f87166020038110156125895761255b838260208a61240c565b60f81b84828151811061256a57fe5b60200101906001600160f81b031916908160001a905350600101612540565b50601f86166008035b6020811015612492576125a9826020898401612386565b60f81b8482815181106125b857fe5b60200101906001600160f81b031916908160001a905350600101612592565b60005b60208110156124e0576125f0828260208961240c565b60f81b8382815181106125ff57fe5b60200101906001600160f81b031916908160001a9053506001016125da565b60008061262d86602087612377565b9050600061263f826020880687612acb565b9050600061265b88602089048488600001518960200151612b0a565b98975050505050505050565b6000606061267484612b9a565b90506000612684876020886123bb565b90506020808706600801106127655760005b6018601f8816600803018110156126de576126d4826020898401068584601801815181106126c057fe5b01602001516001600160f81b031916612c0c565b9150600101612696565b506126f8876020885b048387600001518860200151612b0a565b96506000612708886020896123df565b90506018601f8816600803015b6008811015612741576127378260208a8401068684601801815181106126c057fe5b9150600101612715565b5061275d88602089046001018388604001518960600151612b0a565b9750506127a4565b60005b60088110156127945761278a828260208a06018584601801815181106126c057fe5b9150600101612768565b506127a1876020886126e7565b96505b509495945050505050565b600060606127bc84612b9a565b905060006127cc876020886123bb565b905060208087066020011061285d5760005b601f871660200381101561280d57612803828260208a5b06018584815181106126c057fe5b91506001016127de565b5061281a876020886126e7565b9650600061282a886020896123df565b9050601f87166020035b6020811015612741576128538260208a8401068684815181106126c057fe5b9150600101612834565b60005b602081101561279457612876828260208a6127f5565b9150600101612860565b606060008360f81c8360f81c0360ff16905060008460f81c60ff1690506060826040519080825280602002602001820160405280156128c9578160200160208202803883390190505b50905060005b83811015612908576128e688828501602002612c28565b60001b8282815181106128f557fe5b60209081029190910101526001016128cf565b509695505050505050565b600081516000141561297c57612929600061183f565b8414612974576040805162461bcd60e51b815260206004820152601560248201527432bc3832b1ba32b21032b6b83a3c90313ab33332b960591b604482015290519081900360640190fd5b506000610e9e565b600061299b8360008151811061298e57fe5b602002602001015161183f565b905060015b8351811015612a055784600116600114156129d9576129d28482815181106129c457fe5b602002602001015183612c60565b91506129f9565b6129f6828583815181106129e957fe5b6020026020010151612c60565b91505b600194851c94016129a0565b50848114612a52576040805162461bcd60e51b8152602060048201526015602482015274195e1c1958dd19590818dbdc9c9958dd081c9bdbdd605a1b604482015290519081900360640190fd5b8315612a62575060009050610e9e565b82600081518110612a6f57fe5b60200260200101519150509392505050565b601f036008021c60ff1690565b600080805b8351811015612ac457600882901b9150838181518110612aaf57fe5b016020015160f81c9190911790600101612a93565b5092915050565b60006060612ad885612b9a565b90508260f81b818581518110612aea57fe5b60200101906001600160f81b031916908160001a905350610e9a81612a8e565b60008151600314612b4c5760405162461bcd60e51b81526004018080602001828103825260228152602001806131a76022913960400191505060405180910390fd5b611b098686868686600081518110612b6057fe5b602002602001015160001c87600181518110612b7857fe5b602002602001015188600281518110612b8d57fe5b6020026020010151612c8c565b6040805160208082528183019092526060918391839160208201818038833901905050905060005b6020811015612c04578260f81b8282601f0381518110612bde57fe5b60200101906001600160f81b031916908160001a90535060089290921c91600101612bc2565b509392505050565b60006060612c1985612b9a565b905082818581518110612aea57fe5b600080805b6020811015612c0457600882901b91508481850181518110612c4b57fe5b016020015160f81c9190911790600101612c2d565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600080612c988761183f565b9050612ca5898988612913565b506060612cb0612f1d565b905060018751036001901b8910612d825787612cd0578992505050612f12565b6000612cdb8a612fbe565b88519091505b60018203811015612d0957612cff8c8460018403815181106129e957fe5b9b50600101612ce1565b5060015b60018203811015612d6d578a60011660011415612d4b57612d44836001830381518110612d3657fe5b602002602001015185612c60565b9350612d61565b612d5e848460018403815181106129e957fe5b93505b60019a8b1c9a01612d0d565b50612d788b84612c60565b9350505050612f12565b60015b8751811015612e025760008a600116600114612da15783612db6565b888281518110612dad57fe5b60200260200101515b905060008b600116600114612dde57898381518110612dd157fe5b6020026020010151612de0565b845b9050612dec8282612c60565b60019c8d1c9c909550929092019150612d859050565b508715612e1157509050612f12565b808681518110612e1d57fe5b602002602001015184141580612e31575085155b612e82576040805162461bcd60e51b815260206004820152601c60248201527f726967687420737562747265652063616e6e6f74206265207a65726f00000000604482015290519081900360640190fd5b60008615612e9957612e948686612c60565b612e9b565b855b905080875b60018a5103811015612ec657612ebc828583815181106129e957fe5b9150600101612ea0565b50838114612f0c576040805162461bcd60e51b815260206004820152600e60248201526d0caf0e0cac6e8cac840dac2e8c6d60931b604482015290519081900360640190fd5b50925050505b979650505050505050565b60408051818152610820810182526060918291906020820161080080388339019050509050612f4c600061183f565b81600081518110612f5957fe5b602090810291909101015260015b6040811015612fb857612f99826001830381518110612f8257fe5b60200260200101518360018403815181106129e957fe5b828281518110612fa557fe5b6020908102919091010152600101612f67565b50905090565b600081612fcd57506001610d6f565b612fda600183901c612fbe565b6001019050610d6f565b6040518060a001604052806005906020820280388339509192915050565b604051806101e001604052806130166130d8565b81526020016130236130d8565b81526000602082018190526040820181905260608201819052608082015260a00161304c61309b565b81526000602082015260400161306061316d565b815260200161306d61316d565b8152602001600015158152602001600060ff1681526020016060815260200160008152602001606081525090565b6040518060c00160405280600081526020016130b5613187565b815260606020820181905260006040830181905290820181905260809091015290565b60408051610120810190915260008152602081016130f461309b565b815260200161310161309b565b815260200161310e61309b565b815260200161311b61309b565b8152600060208201819052604082015260600161313661309b565b8152602001600081525090565bfe5b6040518060800160405280606081526020016060815260200160608152602001606081525090565b604051806040016040528060008152602001606081525090565b604080516060808201835260008083526020830152918101919091529056fe6e6f726d616c697a6174696f6e2070726f6f66206861732077726f6e672073697a65a265627a7a723158200cc01bcb243ba1545ebe21e342c9a67c731b3e0ceed39b00a1386dce47d1cdee64736f6c63430005110032"

// DeployOneStepProof2 deploys a new Ethereum contract, binding an instance of OneStepProof2 to it.
func DeployOneStepProof2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProof2, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProof2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OneStepProof2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProof2{OneStepProof2Caller: OneStepProof2Caller{contract: contract}, OneStepProof2Transactor: OneStepProof2Transactor{contract: contract}, OneStepProof2Filterer: OneStepProof2Filterer{contract: contract}}, nil
}

// OneStepProof2 is an auto generated Go binding around an Ethereum contract.
type OneStepProof2 struct {
	OneStepProof2Caller     // Read-only binding to the contract
	OneStepProof2Transactor // Write-only binding to the contract
	OneStepProof2Filterer   // Log filterer for contract events
}

// OneStepProof2Caller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProof2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProof2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProof2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProof2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProof2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProof2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProof2Session struct {
	Contract     *OneStepProof2    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneStepProof2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProof2CallerSession struct {
	Contract *OneStepProof2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OneStepProof2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProof2TransactorSession struct {
	Contract     *OneStepProof2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OneStepProof2Raw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProof2Raw struct {
	Contract *OneStepProof2 // Generic contract binding to access the raw methods on
}

// OneStepProof2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProof2CallerRaw struct {
	Contract *OneStepProof2Caller // Generic read-only contract binding to access the raw methods on
}

// OneStepProof2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProof2TransactorRaw struct {
	Contract *OneStepProof2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProof2 creates a new instance of OneStepProof2, bound to a specific deployed contract.
func NewOneStepProof2(address common.Address, backend bind.ContractBackend) (*OneStepProof2, error) {
	contract, err := bindOneStepProof2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProof2{OneStepProof2Caller: OneStepProof2Caller{contract: contract}, OneStepProof2Transactor: OneStepProof2Transactor{contract: contract}, OneStepProof2Filterer: OneStepProof2Filterer{contract: contract}}, nil
}

// NewOneStepProof2Caller creates a new read-only instance of OneStepProof2, bound to a specific deployed contract.
func NewOneStepProof2Caller(address common.Address, caller bind.ContractCaller) (*OneStepProof2Caller, error) {
	contract, err := bindOneStepProof2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProof2Caller{contract: contract}, nil
}

// NewOneStepProof2Transactor creates a new write-only instance of OneStepProof2, bound to a specific deployed contract.
func NewOneStepProof2Transactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProof2Transactor, error) {
	contract, err := bindOneStepProof2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProof2Transactor{contract: contract}, nil
}

// NewOneStepProof2Filterer creates a new log filterer instance of OneStepProof2, bound to a specific deployed contract.
func NewOneStepProof2Filterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProof2Filterer, error) {
	contract, err := bindOneStepProof2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProof2Filterer{contract: contract}, nil
}

// bindOneStepProof2 binds a generic wrapper to an already deployed contract.
func bindOneStepProof2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProof2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof2 *OneStepProof2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProof2.Contract.OneStepProof2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof2 *OneStepProof2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof2.Contract.OneStepProof2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof2 *OneStepProof2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof2.Contract.OneStepProof2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof2 *OneStepProof2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProof2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof2 *OneStepProof2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof2 *OneStepProof2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof2.Contract.contract.Transact(opts, method, params...)
}

// ExecuteStep is a free data retrieval call binding the contract method 0x1041c884.
//
// Solidity: function executeStep(bytes32 inboxAcc, bytes32 messagesAcc, bytes32 logsAcc, bytes proof, bytes bproof) view returns(uint64 gas, bytes32[5] fields)
func (_OneStepProof2 *OneStepProof2Caller) ExecuteStep(opts *bind.CallOpts, inboxAcc [32]byte, messagesAcc [32]byte, logsAcc [32]byte, proof []byte, bproof []byte) (struct {
	Gas    uint64
	Fields [5][32]byte
}, error) {
	var out []interface{}
	err := _OneStepProof2.contract.Call(opts, &out, "executeStep", inboxAcc, messagesAcc, logsAcc, proof, bproof)

	outstruct := new(struct {
		Gas    uint64
		Fields [5][32]byte
	})

	outstruct.Gas = out[0].(uint64)
	outstruct.Fields = out[1].([5][32]byte)

	return *outstruct, err

}

// ExecuteStep is a free data retrieval call binding the contract method 0x1041c884.
//
// Solidity: function executeStep(bytes32 inboxAcc, bytes32 messagesAcc, bytes32 logsAcc, bytes proof, bytes bproof) view returns(uint64 gas, bytes32[5] fields)
func (_OneStepProof2 *OneStepProof2Session) ExecuteStep(inboxAcc [32]byte, messagesAcc [32]byte, logsAcc [32]byte, proof []byte, bproof []byte) (struct {
	Gas    uint64
	Fields [5][32]byte
}, error) {
	return _OneStepProof2.Contract.ExecuteStep(&_OneStepProof2.CallOpts, inboxAcc, messagesAcc, logsAcc, proof, bproof)
}

// ExecuteStep is a free data retrieval call binding the contract method 0x1041c884.
//
// Solidity: function executeStep(bytes32 inboxAcc, bytes32 messagesAcc, bytes32 logsAcc, bytes proof, bytes bproof) view returns(uint64 gas, bytes32[5] fields)
func (_OneStepProof2 *OneStepProof2CallerSession) ExecuteStep(inboxAcc [32]byte, messagesAcc [32]byte, logsAcc [32]byte, proof []byte, bproof []byte) (struct {
	Gas    uint64
	Fields [5][32]byte
}, error) {
	return _OneStepProof2.Contract.ExecuteStep(&_OneStepProof2.CallOpts, inboxAcc, messagesAcc, logsAcc, proof, bproof)
}
