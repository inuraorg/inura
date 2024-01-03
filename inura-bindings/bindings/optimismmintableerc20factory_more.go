// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/inuraorg/inura/inura-bindings/solc"
)

const OptimismMintableERC20FactoryStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/universal/OptimismMintableERC20Factory.sol:OptimismMintableERC20Factory\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/universal/OptimismMintableERC20Factory.sol:OptimismMintableERC20Factory\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/universal/OptimismMintableERC20Factory.sol:OptimismMintableERC20Factory\",\"label\":\"bridge\",\"offset\":2,\"slot\":\"0\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var OptimismMintableERC20FactoryStorageLayout = new(solc.StorageLayout)

var OptimismMintableERC20FactoryDeployedBin = "0x60806040523480156200001157600080fd5b5060043610620000875760003560e01c8063c4d66de81162000062578063c4d66de81462000135578063ce5ac90f146200014e578063e78cea921462000165578063ee9a31a2146200018c57600080fd5b806354fd4d50146200008c578063896f93d114620000e15780638cf0629c146200011e575b600080fd5b620000c96040518060400160405280600581526020017f312e372e3000000000000000000000000000000000000000000000000000000081525081565b604051620000d89190620005d1565b60405180910390f35b620000f8620000f2366004620006f9565b620001b1565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001620000d8565b620000f86200012f36600462000776565b620001c8565b6200014c620001463660046200080d565b620003c6565b005b620000f86200015f366004620006f9565b62000544565b600054620000f89062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff16620000f8565b6000620001c084848462000544565b949350505050565b600073ffffffffffffffffffffffffffffffffffffffff851662000273576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4f7074696d69736d4d696e7461626c654552433230466163746f72793a206d7560448201527f73742070726f766964652072656d6f746520746f6b656e20616464726573730060648201526084015b60405180910390fd5b6000858585856040516020016200028e94939291906200082b565b604051602081830303815290604052805190602001209050600081600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1688888888604051620002de9062000555565b620002ee95949392919062000885565b8190604051809103906000f59050801580156200030f573d6000803e3d6000fd5b5090508073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fceeb8e7d520d7f3b65fc11a262b91066940193b05d4f93df07cfdced0eb551cf60405160405180910390a360405133815273ffffffffffffffffffffffffffffffffffffffff80891691908316907f52fe89dd5930f343d25650b62fd367bae47088bcddffd2a88350a6ecdd620cdb9060200160405180910390a39695505050505050565b600054600390610100900460ff16158015620003e9575060005460ff8083169116105b62000477576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016200026a565b6000805461010060ff84167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00009092168217177fffffffffffffffffffff000000000000000000000000000000000000000000ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff6201000073ffffffffffffffffffffffffffffffffffffffff87160216179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b6000620001c08484846012620001c8565b61178a80620008eb83390190565b6000815180845260005b818110156200058b576020818501810151868301820152016200056d565b818111156200059e576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000620005e6602083018462000563565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146200061257600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126200065857600080fd5b813567ffffffffffffffff8082111562000676576200067662000617565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715620006bf57620006bf62000617565b81604052838152866020858801011115620006d957600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156200070f57600080fd5b6200071a84620005ed565b9250602084013567ffffffffffffffff808211156200073857600080fd5b620007468783880162000646565b935060408601359150808211156200075d57600080fd5b506200076c8682870162000646565b9150509250925092565b600080600080608085870312156200078d57600080fd5b6200079885620005ed565b9350602085013567ffffffffffffffff80821115620007b657600080fd5b620007c48883890162000646565b94506040870135915080821115620007db57600080fd5b50620007ea8782880162000646565b925050606085013560ff811681146200080257600080fd5b939692955090935050565b6000602082840312156200082057600080fd5b620005e682620005ed565b73ffffffffffffffffffffffffffffffffffffffff851681526080602082015260006200085c608083018662000563565b828103604084015262000870818662000563565b91505060ff8316606083015295945050505050565b600073ffffffffffffffffffffffffffffffffffffffff808816835280871660208401525060a06040830152620008c060a083018662000563565b8281036060840152620008d4818662000563565b91505060ff83166080830152969550505050505056fe60e06040523480156200001157600080fd5b506040516200178a3803806200178a833981016040819052620000349162000163565b828260036200004483826200029e565b5060046200005382826200029e565b5050506001600160a01b039384166080529390921660a052505060ff1660c0526200036a565b80516001600160a01b03811681146200009157600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620000be57600080fd5b81516001600160401b0380821115620000db57620000db62000096565b604051601f8301601f19908116603f0116810190828211818310171562000106576200010662000096565b816040528381526020925086838588010111156200012357600080fd5b600091505b8382101562000147578582018301518183018401529082019062000128565b83821115620001595760008385830101525b9695505050505050565b600080600080600060a086880312156200017c57600080fd5b620001878662000079565b9450620001976020870162000079565b60408701519094506001600160401b0380821115620001b557600080fd5b620001c389838a01620000ac565b94506060880151915080821115620001da57600080fd5b50620001e988828901620000ac565b925050608086015160ff811681146200020157600080fd5b809150509295509295909350565b600181811c908216806200022457607f821691505b6020821081036200024557634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200029957600081815260208120601f850160051c81016020861015620002745750805b601f850160051c820191505b81811015620002955782815560010162000280565b5050505b505050565b81516001600160401b03811115620002ba57620002ba62000096565b620002d281620002cb84546200020f565b846200024b565b602080601f8311600181146200030a5760008415620002f15750858301515b600019600386901b1c1916600185901b17855562000295565b600085815260208120601f198616915b828110156200033b578886015182559484019460019091019084016200031a565b50858210156200035a5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a05160c0516113d4620003b6600039600061024401526000818161034b015281816103e001528181610625015261075c0152600081816101a9015261037101526113d46000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063ae1f6aaf1161008c578063dd62ed3e11610066578063dd62ed3e14610395578063e78cea9214610349578063ee9a31a2146103db57600080fd5b8063ae1f6aaf14610349578063c01e1bd61461036f578063d6c0b2c41461036f57600080fd5b80639dc29fac116100bd5780639dc29fac14610310578063a457c2d714610323578063a9059cbb1461033657600080fd5b806370a08231146102d257806395d89b411461030857600080fd5b806323b872dd1161012f5780633950935111610114578063395093511461026e57806340c10f191461028157806354fd4d501461029657600080fd5b806323b872dd1461022a578063313ce5671461023d57600080fd5b806306fdde031161016057806306fdde03146101f0578063095ea7b31461020557806318160ddd1461021857600080fd5b806301ffc9a71461017c578063033964be146101a4575b600080fd5b61018f61018a36600461117d565b610402565b60405190151581526020015b60405180910390f35b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f86104f3565b60405161019b91906111c6565b61018f610213366004611262565b610585565b6002545b60405190815260200161019b565b61018f61023836600461128c565b61059d565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016815260200161019b565b61018f61027c366004611262565b6105c1565b61029461028f366004611262565b61060d565b005b6101f86040518060400160405280600581526020017f312e332e3000000000000000000000000000000000000000000000000000000081525081565b61021c6102e03660046112c8565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6101f8610735565b61029461031e366004611262565b610744565b61018f610331366004611262565b61085b565b61018f610344366004611262565b61092c565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b61021c6103a33660046112e3565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f1d1d8b63000000000000000000000000000000000000000000000000000000007fec4fc8e3000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000085168314806104bb57507fffffffff00000000000000000000000000000000000000000000000000000000858116908316145b806104ea57507fffffffff00000000000000000000000000000000000000000000000000000000858116908216145b95945050505050565b60606003805461050290611316565b80601f016020809104026020016040519081016040528092919081815260200182805461052e90611316565b801561057b5780601f106105505761010080835404028352916020019161057b565b820191906000526020600020905b81548152906001019060200180831161055e57829003601f168201915b5050505050905090565b60003361059381858561093a565b5060019392505050565b6000336105ab858285610aee565b6105b6858585610bc5565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906105939082908690610608908790611398565b61093a565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146106d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4f7074696d69736d4d696e7461626c6545524332303a206f6e6c79206272696460448201527f67652063616e206d696e7420616e64206275726e00000000000000000000000060648201526084015b60405180910390fd5b6106e18282610e78565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405161072991815260200190565b60405180910390a25050565b60606004805461050290611316565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610809576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4f7074696d69736d4d696e7461626c6545524332303a206f6e6c79206272696460448201527f67652063616e206d696e7420616e64206275726e00000000000000000000000060648201526084016106ce565b6108138282610f98565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca58260405161072991815260200190565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091908381101561091f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084016106ce565b6105b6828686840361093a565b600033610593818585610bc5565b73ffffffffffffffffffffffffffffffffffffffff83166109dc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f726573730000000000000000000000000000000000000000000000000000000060648201526084016106ce565b73ffffffffffffffffffffffffffffffffffffffff8216610a7f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f737300000000000000000000000000000000000000000000000000000000000060648201526084016106ce565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610bbf5781811015610bb2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016106ce565b610bbf848484840361093a565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c68576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f647265737300000000000000000000000000000000000000000000000000000060648201526084016106ce565b73ffffffffffffffffffffffffffffffffffffffff8216610d0b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f657373000000000000000000000000000000000000000000000000000000000060648201526084016106ce565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610dc1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e6365000000000000000000000000000000000000000000000000000060648201526084016106ce565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260208190526040808220858503905591851681529081208054849290610e05908490611398565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610e6b91815260200190565b60405180910390a3610bbf565b73ffffffffffffffffffffffffffffffffffffffff8216610ef5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016106ce565b8060026000828254610f079190611398565b909155505073ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604081208054839290610f41908490611398565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff821661103b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f730000000000000000000000000000000000000000000000000000000000000060648201526084016106ce565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260208190526040902054818110156110f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f636500000000000000000000000000000000000000000000000000000000000060648201526084016106ce565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040812083830390556002805484929061112d9084906113b0565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610ae1565b60006020828403121561118f57600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146111bf57600080fd5b9392505050565b600060208083528351808285015260005b818110156111f3578581018301518582016040015282016111d7565b81811115611205576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461125d57600080fd5b919050565b6000806040838503121561127557600080fd5b61127e83611239565b946020939093013593505050565b6000806000606084860312156112a157600080fd5b6112aa84611239565b92506112b860208501611239565b9150604084013590509250925092565b6000602082840312156112da57600080fd5b6111bf82611239565b600080604083850312156112f657600080fd5b6112ff83611239565b915061130d60208401611239565b90509250929050565b600181811c9082168061132a57607f821691505b602082108103611363577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156113ab576113ab611369565b500190565b6000828210156113c2576113c2611369565b50039056fea164736f6c634300080f000aa164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(OptimismMintableERC20FactoryStorageLayoutJSON), OptimismMintableERC20FactoryStorageLayout); err != nil {
		panic(err)
	}

	layouts["OptimismMintableERC20Factory"] = OptimismMintableERC20FactoryStorageLayout
	deployedBytecodes["OptimismMintableERC20Factory"] = OptimismMintableERC20FactoryDeployedBin
}
