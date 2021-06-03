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

const ForwarderBytecode = "6101206040523480156200001257600080fd5b506040518060400160405280601081526020017f4d696e696d616c466f72776172646572000000000000000000000000000000008152506040518060400160405280600581526020017f302e302e3100000000000000000000000000000000000000000000000000000081525060008280519060200120905060008280519060200120905060007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f90508260c081815250508160e081815250504660a08181525050620000e78184846200010260201b60201c565b60808181525050806101008181525050505050505062000216565b600083838346306040516020016200011f95949392919062000171565b6040516020818303038152906040528051906020012090509392505050565b6200014981620001ce565b82525050565b6200015a81620001e2565b82525050565b6200016b816200020c565b82525050565b600060a0820190506200018860008301886200014f565b6200019760208301876200014f565b620001a660408301866200014f565b620001b5606083018562000160565b620001c460808301846200013e565b9695505050505050565b6000620001db82620001ec565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60805160a05160c05160e051610100516112046200025b60003960006105f40152600061063601526000610615015260006105a1015260006105c901526112046000f3fe6080604052600436106100345760003560e01c80632d0335ab1461003957806347153f8214610076578063bf5d3bdb146100a7575b600080fd5b34801561004557600080fd5b50610060600480360381019061005b91906108d2565b6100e4565b60405161006d9190610dc8565b60405180910390f35b610090600480360381019061008b91906108fb565b61012c565b60405161009e929190610bf1565b60405180910390f35b3480156100b357600080fd5b506100ce60048036038101906100c991906108fb565b610304565b6040516100db9190610bd6565b60405180910390f35b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000606061013b858585610304565b61017a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017190610d48565b60405180910390fd5b6001856080013561018b9190610e7d565b6000808760000160208101906101a191906108d2565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000808660200160208101906101f591906108d2565b73ffffffffffffffffffffffffffffffffffffffff1687606001358860400135898060a001906102259190610de3565b8b600001602081019061023891906108d2565b60405160200161024a93929190610b5e565b6040516020818303038152906040526040516102669190610b88565b600060405180830381858888f193505050503d80600081146102a4576040519150601f19603f3d011682016040523d82523d6000602084013e6102a9565b606091505b5091509150603f87606001356102bf9190610ed3565b5a116102f4577f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b8181935093505050935093915050565b60008061040d84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506103ff7fdd8f4b70b0f4393e889bd39128a30628a78b61816a9eb8199759e7a349657e4888600001602081019061038691906108d2565b89602001602081019061039991906108d2565b8a604001358b606001358c608001358d8060a001906103b89190610de3565b6040516103c6929190610b45565b60405180910390206040516020016103e49796959493929190610c21565b604051602081830303815290604052805190602001206104b9565b6104d390919063ffffffff16565b9050846080013560008087600001602081019061042a91906108d2565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541480156104af575084600001602081019061048091906108d2565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b9150509392505050565b60006104cc6104c661059d565b83610660565b9050919050565b600080600080604185511415610500576020850151925060408501519150606085015160001a9050610586565b60408551141561054a576040850151602086015193507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81169250601b8160ff1c01915050610585565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161057c90610d68565b60405180910390fd5b5b61059286828585610693565b935050505092915050565b60007f00000000000000000000000000000000000000000000000000000000000000004614156105ef577f0000000000000000000000000000000000000000000000000000000000000000905061065d565b61065a7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000061081e565b90505b90565b60008282604051602001610675929190610b9f565b60405160208183030381529060405280519060200120905092915050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c11156106fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f290610d88565b60405180910390fd5b601b8460ff1614806107105750601c8460ff16145b61074f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074690610da8565b60405180910390fd5b6000600186868686604051600081526020016040526040516107749493929190610ce3565b6020604051602081039080840390855afa158015610796573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610812576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161080990610d28565b60405180910390fd5b80915050949350505050565b60008383834630604051602001610839959493929190610c90565b6040516020818303038152906040528051906020012090509392505050565b600081359050610867816111b7565b92915050565b60008083601f84011261087f57600080fd5b8235905067ffffffffffffffff81111561089857600080fd5b6020830191508360018202830111156108b057600080fd5b9250929050565b600060c082840312156108c957600080fd5b81905092915050565b6000602082840312156108e457600080fd5b60006108f284828501610858565b91505092915050565b60008060006040848603121561091057600080fd5b600084013567ffffffffffffffff81111561092a57600080fd5b610936868287016108b7565b935050602084013567ffffffffffffffff81111561095357600080fd5b61095f8682870161086d565b92509250509250925092565b61097481610f04565b82525050565b61098b61098682610f04565b610fa5565b82525050565b61099a81610f16565b82525050565b6109a981610f22565b82525050565b6109c06109bb82610f22565b610fb7565b82525050565b60006109d28385610e56565b93506109df838584610f63565b82840190509392505050565b60006109f682610e3a565b610a008185610e45565b9350610a10818560208601610f72565b610a1981611031565b840191505092915050565b6000610a2f82610e3a565b610a398185610e56565b9350610a49818560208601610f72565b80840191505092915050565b6000610a62601883610e61565b9150610a6d8261104f565b602082019050919050565b6000610a85603283610e61565b9150610a9082611078565b604082019050919050565b6000610aa8601f83610e61565b9150610ab3826110c7565b602082019050919050565b6000610acb600283610e72565b9150610ad6826110f0565b600282019050919050565b6000610aee602283610e61565b9150610af982611119565b604082019050919050565b6000610b11602283610e61565b9150610b1c82611168565b604082019050919050565b610b3081610f4c565b82525050565b610b3f81610f56565b82525050565b6000610b528284866109c6565b91508190509392505050565b6000610b6b8285876109c6565b9150610b77828461097a565b601482019150819050949350505050565b6000610b948284610a24565b915081905092915050565b6000610baa82610abe565b9150610bb682856109af565b602082019150610bc682846109af565b6020820191508190509392505050565b6000602082019050610beb6000830184610991565b92915050565b6000604082019050610c066000830185610991565b8181036020830152610c1881846109eb565b90509392505050565b600060e082019050610c36600083018a6109a0565b610c43602083018961096b565b610c50604083018861096b565b610c5d6060830187610b27565b610c6a6080830186610b27565b610c7760a0830185610b27565b610c8460c08301846109a0565b98975050505050505050565b600060a082019050610ca560008301886109a0565b610cb260208301876109a0565b610cbf60408301866109a0565b610ccc6060830185610b27565b610cd9608083018461096b565b9695505050505050565b6000608082019050610cf860008301876109a0565b610d056020830186610b36565b610d1260408301856109a0565b610d1f60608301846109a0565b95945050505050565b60006020820190508181036000830152610d4181610a55565b9050919050565b60006020820190508181036000830152610d6181610a78565b9050919050565b60006020820190508181036000830152610d8181610a9b565b9050919050565b60006020820190508181036000830152610da181610ae1565b9050919050565b60006020820190508181036000830152610dc181610b04565b9050919050565b6000602082019050610ddd6000830184610b27565b92915050565b60008083356001602003843603038112610dfc57600080fd5b80840192508235915067ffffffffffffffff821115610e1a57600080fd5b602083019250600182023603831315610e3257600080fd5b509250929050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600081905092915050565b6000610e8882610f4c565b9150610e9383610f4c565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610ec857610ec7610fd3565b5b828201905092915050565b6000610ede82610f4c565b9150610ee983610f4c565b925082610ef957610ef8611002565b5b828204905092915050565b6000610f0f82610f2c565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015610f90578082015181840152602081019050610f75565b83811115610f9f576000848401525b50505050565b6000610fb082610fc1565b9050919050565b6000819050919050565b6000610fcc82611042565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f45434453413a20696e76616c6964207369676e61747572650000000000000000600082015250565b7f4d696e696d616c466f727761726465723a207369676e617475726520646f657360008201527f206e6f74206d6174636820726571756573740000000000000000000000000000602082015250565b7f45434453413a20696e76616c6964207369676e6174757265206c656e67746800600082015250565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b7f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b7f45434453413a20696e76616c6964207369676e6174757265202776272076616c60008201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b6111c081610f04565b81146111cb57600080fd5b5056fea2646970667358221220e33af66dcffb6010c71e0ff484c771afe2a79a1d6f26ecbadc28c45c0378356064736f6c63430008010033"

func TestCreateAndExecuteForwarder(t *testing.T) {
	pl := AliceKp
	client := ethtest.NewClient(t, TestEndpoint, pl)
	client.LockNonceAndUpdate()

	forwarderAbi, err := abi.JSON(strings.NewReader(ForwarderAbi))
	if err != nil {
		t.Fatal(err.Error())
	}
	forwarderAddress, tx, forwarderContract, err := bind.DeployContract(client.Opts, forwarderAbi, common.FromHex(ForwarderBytecode), client.Client)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = utils.WaitForTx(client, tx)
	if err != nil {
		t.Fatal(err.Error())
	}
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
		Gas:   gas.Uint64(),
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
