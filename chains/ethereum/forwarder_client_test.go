// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"strings"
	"testing"

	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	ethtest "github.com/ChainSafe/ChainBridge/shared/ethereum/testing"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

const GsnForwarderBytecode = "60806040523480156200001157600080fd5b5060006040518060800160405280605d81526020016200130a605d9139604051602001620000409190620000c8565b60408051601f1981840301815291905290506200005d8162000064565b5062000174565b8051602080830191909120600081815291829052604091829020805460ff19166001179055905181907f64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb20290620000bc9085906200010c565b60405180910390a25050565b60006e08cdee4eec2e4c8a4cae2eacae6e85608b1b82528251620000f481600f85016020870162000141565b602960f81b600f939091019283015250601001919050565b60006020825282518060208401526200012d81604085016020870162000141565b601f01601f19169190910160400192915050565b60005b838110156200015e57818101518382015260200162000144565b838111156200016e576000848401525b50505050565b61118680620001846000396000f3fe6080604052600436106100955760003560e01c8063c3f28abd11610059578063c3f28abd14610168578063c722f1771461017d578063d9210be51461019d578063e024dc7f146101bd578063e2b62f2d146101de5761009c565b8063066a310c146100a157806321fe98df146100cc5780632d0335ab146100f95780639c7b459214610126578063ad9f99c7146101485761009c565b3661009c57005b600080fd5b3480156100ad57600080fd5b506100b66101fe565b6040516100c39190610dc0565b60405180910390f35b3480156100d857600080fd5b506100ec6100e7366004610a9a565b61021a565b6040516100c39190610d48565b34801561010557600080fd5b50610119610114366004610a6c565b61022f565b6040516100c39190611023565b34801561013257600080fd5b50610146610141366004610ab2565b61024a565b005b34801561015457600080fd5b50610146610163366004610b1b565b61032c565b34801561017457600080fd5b506100b661034d565b34801561018957600080fd5b506100ec610198366004610a9a565b610369565b3480156101a957600080fd5b506101466101b8366004610ab2565b61037e565b6101d06101cb366004610b1b565b61044d565b6040516100c3929190610d53565b3480156101ea57600080fd5b506100b66101f9366004610bbf565b610603565b6040518060800160405280605d81526020016110a2605d913981565b60006020819052908152604090205460ff1681565b6001600160a01b031660009081526002602052604090205490565b600046905060006040518060800160405280605281526020016110ff60529139805190602001208686604051610281929190610c8c565b60405180910390208585604051610299929190610c8c565b6040519081900381206102b493929186903090602001610d76565b60408051601f198184030181528282528051602080830191909120600081815260019283905293909320805460ff1916909117905592509081907f4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d89061031b908590610dc0565b60405180910390a250505050505050565b6103358761069d565b610344878787878787876106f4565b50505050505050565b6040518060800160405280605281526020016110ff6052913981565b60016020526000908152604090205460ff1681565b60005b838110156103f857600085858381811061039757fe5b909101356001600160f81b031916915050600560fb1b81148015906103ca5750602960f81b6001600160f81b0319821614155b6103ef5760405162461bcd60e51b81526004016103e690610fbd565b60405180910390fd5b50600101610381565b50600084846040518060800160405280605d81526020016110a2605d9139858560405160200161042c959493929190610ce1565b604051602081830303815290604052905061044681610816565b5050505050565b60006060610460898989898989896106f4565b61046989610878565b60c0890135158061047d5750438960c00135115b6104995760405162461bcd60e51b81526004016103e690610f58565b600060408a0135156104aa5750619c405b60006104b960a08c018c61102c565b6104c660208e018e610a6c565b6040516020016104d893929190610c9c565b6040516020818303038152906040529050818b606001350160405a603f02816104fd57fe5b04101561051c5760405162461bcd60e51b81526004016103e690610f29565b61052c60408c0160208d01610a6c565b6001600160a01b03168b606001358c604001358360405161054d9190610cc5565b600060405180830381858888f193505050503d806000811461058b576040519150601f19603f3d011682016040523d82523d6000602084013e610590565b606091505b50909450925060408b0135158015906105a95750600047115b156105f5576105bb60208c018c610a6c565b6001600160a01b03166108fc479081150290604051600060405180830381858888f193505050501580156105f3573d6000803e3d6000fd5b505b505097509795505050505050565b6060836106136020870187610a6c565b6001600160a01b031661062c6040880160208901610a6c565b6001600160a01b03166040880135606089013560808a013561065160a08c018c61102c565b60405161065f929190610c8c565b6040519081900381206106849796959493929160c08e0135908c908c90602001610c3f565b6040516020818303038152906040529050949350505050565b6080810135600260006106b36020850185610a6c565b6001600160a01b03166001600160a01b0316815260200190815260200160002054146106f15760405162461bcd60e51b81526004016103e690610efc565b50565b60008681526001602052604090205460ff166107225760405162461bcd60e51b81526004016103e690610f86565b60008581526020819052604090205460ff166107505760405162461bcd60e51b81526004016103e690610fec565b60008661075f89888888610603565b8051602091820120604051610775939201610d2d565b60408051601f198184030181529190528051602091820120915061079b90890189610a6c565b6001600160a01b03166107e684848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525086939250506108cd9050565b6001600160a01b03161461080c5760405162461bcd60e51b81526004016103e690610ec5565b5050505050505050565b8051602080830191909120600081815291829052604091829020805460ff19166001179055905181907f64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb2029061086c908590610dc0565b60405180910390a25050565b60808101356002600061088e6020850185610a6c565b6001600160a01b031681526020810191909152604001600020805460018101909155146106f15760405162461bcd60e51b81526004016103e690610efc565b600081516041146108f05760405162461bcd60e51b81526004016103e690610e0a565b60208201516040830151606084015160001a61090e86828585610918565b9695505050505050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561095a5760405162461bcd60e51b81526004016103e690610e41565b8360ff16601b148061096f57508360ff16601c145b61098b5760405162461bcd60e51b81526004016103e690610e83565b6000600186868686604051600081526020016040526040516109b09493929190610da2565b6020604051602081039080840390855afa1580156109d2573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610a055760405162461bcd60e51b81526004016103e690610dd3565b95945050505050565b60008083601f840112610a1f578182fd5b50813567ffffffffffffffff811115610a36578182fd5b602083019150836020828501011115610a4e57600080fd5b9250929050565b600060e08284031215610a66578081fd5b50919050565b600060208284031215610a7d578081fd5b81356001600160a01b0381168114610a93578182fd5b9392505050565b600060208284031215610aab578081fd5b5035919050565b60008060008060408587031215610ac7578283fd5b843567ffffffffffffffff80821115610ade578485fd5b610aea88838901610a0e565b90965094506020870135915080821115610b02578384fd5b50610b0f87828801610a0e565b95989497509550505050565b600080600080600080600060a0888a031215610b35578283fd5b873567ffffffffffffffff80821115610b4c578485fd5b610b588b838c01610a55565b985060208a0135975060408a0135965060608a0135915080821115610b7b578485fd5b610b878b838c01610a0e565b909650945060808a0135915080821115610b9f578384fd5b50610bac8a828b01610a0e565b989b979a50959850939692959293505050565b60008060008060608587031215610bd4578384fd5b843567ffffffffffffffff80821115610beb578586fd5b610bf788838901610a55565b9550602087013594506040870135915080821115610b02578384fd5b60008151808452610c2b816020860160208601611071565b601f01601f19169290920160200192915050565b60008b82528a60208301528960408301528860608301528760808301528660a08301528560c08301528460e083015261010083858285013791909201019081529998505050505050505050565b6000828483379101908152919050565b6000838583375060609190911b6bffffffffffffffffffffffff19169101908152601401919050565b60008251610cd7818460208701611071565b9190910192915050565b600085878337600560fb1b8287019081528551610d05816001840160208a01611071565b600b60fa1b600192909101918201528385600283013790920160020191825250949350505050565b61190160f01b81526002810192909252602282015260420190565b901515815260200190565b6000831515825260406020830152610d6e6040830184610c13565b949350505050565b9485526020850193909352604084019190915260608301526001600160a01b0316608082015260a00190565b93845260ff9290921660208401526040830152606082015260800190565b600060208252610a936020830184610c13565b60208082526018908201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604082015260600190565b6020808252601f908201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604082015260600190565b60208082526022908201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604082015261756560f01b606082015260800190565b60208082526022908201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604082015261756560f01b606082015260800190565b60208082526017908201527f4657443a207369676e6174757265206d69736d61746368000000000000000000604082015260600190565b60208082526013908201527208cae887440dcdedcc6ca40dad2e6dac2e8c6d606b1b604082015260600190565b6020808252601590820152744657443a20696e73756666696369656e742067617360581b604082015260600190565b6020808252601490820152731195d10e881c995c5d595cdd08195e1c1a5c995960621b604082015260600190565b6020808252601d908201527f4657443a20756e7265676973746572656420646f6d61696e207365702e000000604082015260600190565b6020808252601590820152744657443a20696e76616c696420747970656e616d6560581b604082015260600190565b6020808252601a908201527f4657443a20756e72656769737465726564207479706568617368000000000000604082015260600190565b90815260200190565b6000808335601e19843603018112611042578283fd5b83018035915067ffffffffffffffff82111561105c578283fd5b602001915036819003821315610a4e57600080fd5b60005b8381101561108c578181015183820152602001611074565b8381111561109b576000848401525b5050505056fe616464726573732066726f6d2c6164647265737320746f2c75696e743235362076616c75652c75696e74323536206761732c75696e74323536206e6f6e63652c627974657320646174612c75696e743235362076616c6964556e74696c454950373132446f6d61696e28737472696e67206e616d652c737472696e672076657273696f6e2c75696e7432353620636861696e49642c6164647265737320766572696679696e67436f6e747261637429a2646970667358221220a3c3167f48999c4f15c9c609504b0cac88185f26c2db9b5cb211592d60724f5264736f6c63430007060033616464726573732066726f6d2c6164647265737320746f2c75696e743235362076616c75652c75696e74323536206761732c75696e74323536206e6f6e63652c627974657320646174612c75696e743235362076616c6964556e74696c"

func TestCreateAndExecuteGsnForwarder(t *testing.T) {
	pl := AliceKp
	client := ethtest.NewClient(t, TestEndpoint, pl)
	client.LockNonceAndUpdate()

	forwarderAbi, err := abi.JSON(strings.NewReader(GsnForwarderAbi))
	if err != nil {
		t.Fatal(err.Error())
	}
	forwarderAddress, tx, forwarderContract, err := bind.DeployContract(client.Opts, forwarderAbi, common.FromHex(GsnForwarderBytecode), client.Client)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = utils.WaitForTx(client, tx)
	if err != nil {
		t.Fatal(err.Error())
	}
	client.UnlockNonce()
	client.LockNonceAndUpdate()
	domainRegistrationPacked, err := forwarderAbi.Pack("registerDomainSeparator", "GSN Relayed Transaction", "2")
	if err != nil {
		t.Fatal(err.Error())
	}
	sendDomainTx, err := forwarderContract.RawTransact(client.Opts, domainRegistrationPacked)
	if err != nil {
		t.Fatal(err.Error())
	}
	utils.WaitForTx(client, sendDomainTx)
	client.UnlockNonce()

	forwarderClient := NewForwarderClient(client.Client, forwarderAddress, pl.CommonAddress())
	nonce, err := forwarderClient.LockAndNextNonce()
	if err != nil {
		t.Fatal(err.Error())
	}

	value := big.NewInt(0)
	gas := big.NewInt(100000)
	chainId, err := client.Client.ChainID(client.Opts.Context)
	if err != nil {
		t.Fatal(err.Error())
	}

	packed, err := forwarderClient.PackAndSignForwarderArg(
		pl.CommonAddress(),
		pl.CommonAddress(),
		common.Hex2Bytes("0x"),
		nonce,
		value,
		gas,
		uint(chainId.Uint64()),
		*pl)

	if err != nil {
		t.Fatal(err.Error())
	}

	// now make a call with the packed data
	callMsg := eth.CallMsg{
		To:    &forwarderAddress,
		Data:  packed,
		From:  pl.CommonAddress(),
		Gas:   gas.Uint64() + 100000,
		Value: value,
	}
	res, err := client.Client.CallContract(client.Opts.Context, callMsg, nil)

	if err != nil {
		t.Fatal(err.Error())
	} else {
		p, unErr := forwarderClient.forwarderAbi.Unpack("execute", res)
		if unErr != nil {
			t.Fatal(unErr.Error())
		}
		if p[0] == false {
			t.Fatal("Inner call failed")
		}
	}

	err = client.LockNonceAndUpdate()
	if err != nil {
		t.Fatal(err.Error())
	}

	onChainNonce, err := forwarderClient.GetOnChainNonce()
	if err != nil {
		t.Fatal(err.Error())
	}

	if onChainNonce.Cmp(nonce) != 0 || onChainNonce.Cmp(big.NewInt(0)) != 0 {
		t.Fatal("Invalid start nonce", onChainNonce, nonce)
	}

	// now send a transaction and check that the nonce was updated
	sendTx, err := forwarderContract.RawTransact(client.Opts, packed)
	if err != nil {
		t.Fatal(err.Error())
	}
	utils.WaitForTx(client, sendTx)

	onChainNonce, err = forwarderClient.GetOnChainNonce()
	if err != nil {
		t.Fatal(err.Error())
	}
	if onChainNonce.Cmp(nonce.Add(nonce, big.NewInt(1))) != 0 || onChainNonce.Cmp(big.NewInt(1)) != 0 {
		t.Fatal("Invalid end nonce", onChainNonce, nonce)
	}

	client.UnlockNonce()
	forwarderClient.UnlockAndSetNonce(nonce)
}
