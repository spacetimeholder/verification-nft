// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// StructsVerificationV0 is an auto generated low-level Go binding around an user-defined struct.
type StructsVerificationV0 struct {
	TokenId   *big.Int
	Service   string
	Timestamp *big.Int
}

// VerificationV0MetaData contains all meta data concerning the VerificationV0 contract.
var VerificationV0MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASE_TOKEN_URI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"service\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"verifications\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"service\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.VerificationV0[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526001600b553480156200001657600080fd5b5060405162001e2038038062001e2083398101604081905262000039916200022b565b828260006200004983826200034d565b5060016200005882826200034d565b50600c91506200006b905082826200034d565b5062000079600033620000b0565b50620000a67f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a633620000b0565b5050505062000419565b6000828152600a602090815260408083206001600160a01b038516845290915281205460ff1662000159576000838152600a602090815260408083206001600160a01b03861684529091529020805460ff19166001179055620001103390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016200015d565b5060005b92915050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200018b57600080fd5b81516001600160401b0380821115620001a857620001a862000163565b604051601f8301601f19908116603f01168101908282118183101715620001d357620001d362000163565b8160405283815260209250866020858801011115620001f157600080fd5b600091505b83821015620002155785820183015181830184015290820190620001f6565b6000602085830101528094505050505092915050565b6000806000606084860312156200024157600080fd5b83516001600160401b03808211156200025957600080fd5b620002678783880162000179565b945060208601519150808211156200027e57600080fd5b6200028c8783880162000179565b93506040860151915080821115620002a357600080fd5b50620002b28682870162000179565b9150509250925092565b600181811c90821680620002d157607f821691505b602082108103620002f257634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111562000348576000816000526020600020601f850160051c81016020861015620003235750805b601f850160051c820191505b8181101562000344578281556001016200032f565b5050505b505050565b81516001600160401b0381111562000369576200036962000163565b62000381816200037a8454620002bc565b84620002f8565b602080601f831160018114620003b95760008415620003a05750858301515b600019600386901b1c1916600185901b17855562000344565b600085815260208120601f198616915b82811015620003ea57888601518255948401946001909101908401620003c9565b5085821015620004095787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6119f780620004296000396000f3fe608060405234801561001057600080fd5b50600436106101545760003560e01c806301ffc9a71461015957806306fdde0314610181578063080c98da14610196578063081812fc146101b6578063095ea7b3146101d657806318160ddd146101eb57806323b872dd146101fd578063248a9ca3146102105780632f2ff15d146102235780632f745c591461023657806336568abe1461024957806342842e0e1461025c57806342966c681461026f5780634f6ccce7146102825780636352211e146102955780636f458133146102a857806370a08231146102db57806391d14854146102ee57806395d89b4114610301578063a217fddf14610309578063a22cb46514610311578063b88d4fde14610324578063ba7aef4314610337578063c87b56dd1461034a578063d53913931461035d578063d547741f14610372578063db0a087c14610385578063e985e9c51461038d575b600080fd5b61016c61016736600461136c565b6103a0565b60405190151581526020015b60405180910390f35b6101896103b1565b60405161017891906113d6565b6101a96101a4366004611405565b610443565b6040516101789190611420565b6101c96101c43660046114a3565b61061c565b60405161017891906114bc565b6101e96101e43660046114d0565b610631565b005b6008545b604051908152602001610178565b6101e961020b3660046114fa565b610640565b6101ef61021e3660046114a3565b610683565b6101e9610231366004611536565b610698565b6101ef6102443660046114d0565b6106ba565b6101e9610257366004611536565b610711565b6101e961026a3660046114fa565b610749565b6101e961027d3660046114a3565b610764565b6101ef6102903660046114a3565b610785565b6101c96102a33660046114a3565b6107d9565b6101896040518060400160405280601481526020017368747470733a2f2f617277656176652e6e65742f60601b81525081565b6101ef6102e9366004611405565b6107e4565b61016c6102fc366004611536565b61082c565b610189610857565b6101ef600081565b6101e961031f366004611562565b610866565b6101e9610332366004611629565b610871565b6101e96103453660046116a4565b610888565b6101896103583660046114a3565b6108f5565b6101ef6000805160206119a283398151915281565b6101e9610380366004611536565b610920565b61018961093c565b61016c61039b36600461170e565b61094b565b60006103ab82610979565b92915050565b6060600080546103c090611738565b80601f01602080910402602001604051908101604052809291908181526020018280546103ec90611738565b80156104395780601f1061040e57610100808354040283529160200191610439565b820191906000526020600020905b81548152906001019060200180831161041c57829003601f168201915b5050505050905090565b60606000610450836107e4565b90506000811161049a5760405162461bcd60e51b815260206004820152601060248201526f6e6f20766572696669636174696f6e7360801b60448201526064015b60405180910390fd5b6000816001600160401b038111156104b4576104b461159e565b60405190808252806020026020018201604052801561050957816020015b6104f660405180606001604052806000815260200160608152602001600081525090565b8152602001906001900390816104d25790505b50905060005b8281101561061457600061052386836106ba565b90506040518060600160405280828152602001600d6000848152602001908152602001600020805461055490611738565b80601f016020809104026020016040519081016040528092919081815260200182805461058090611738565b80156105cd5780601f106105a2576101008083540402835291602001916105cd565b820191906000526020600020905b8154815290600101906020018083116105b057829003601f168201915b50505050508152602001600e60008481526020019081526020016000205481525083838151811061060057610600611772565b60209081029190910101525060010161050f565b509392505050565b60006106278261099e565b506103ab826109d6565b61063c8282336109f1565b5050565b60405162461bcd60e51b8152602060048201526018602482015277746f6b656e73206e6f6e2d7472616e736665727261626c6560401b6044820152606401610491565b6000908152600a602052604090206001015490565b6106a182610683565b6106aa816109fe565b6106b48383610a0b565b50505050565b60006106c5836107e4565b82106106e857828260405163295f44f760e21b8152600401610491929190611788565b506001600160a01b03919091166000908152600660209081526040808320938352929052205490565b6001600160a01b038116331461073a5760405163334bd91960e11b815260040160405180910390fd5b6107448282610a9f565b505050565b61074483838360405180602001604052806000815250610871565b6000805160206119a283398151915261077c816109fe565b61063c82610b0c565b600061079060085490565b82106107b45760008260405163295f44f760e21b8152600401610491929190611788565b600882815481106107c7576107c7611772565b90600052602060002001549050919050565b60006103ab8261099e565b60006001600160a01b0382166108105760006040516322718ad960e21b815260040161049191906114bc565b506001600160a01b031660009081526003602052604090205490565b6000918252600a602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6060600180546103c090611738565b61063c338383610b47565b61087c848484610640565b6106b484848484610bdd565b6000805160206119a28339815191526108a0816109fe565b6108ac84600b54610cf4565b600b546000908152600d602052604090206108c784826117f1565b50600b80546000908152600e60205260408120849055815491906108ea836118c6565b919050555050505050565b60606109008261099e565b506040518060600160405280603f8152602001611963603f913992915050565b61092982610683565b610932816109fe565b6106b48383610a9f565b6060600c80546103c090611738565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b60006001600160e01b03198216637965db0b60e01b14806103ab57506103ab82610d0e565b6000806109aa83610d33565b90506001600160a01b0381166103ab57604051637e27328960e01b815260048101849052602401610491565b6000908152600460205260409020546001600160a01b031690565b6107448383836001610d4e565b610a088133610e4b565b50565b6000610a17838361082c565b610a97576000838152600a602090815260408083206001600160a01b03861684529091529020805460ff19166001179055610a4f3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016103ab565b5060006103ab565b6000610aab838361082c565b15610a97576000838152600a602090815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016103ab565b6000610b1b6000836000610e76565b90506001600160a01b03811661063c57604051637e27328960e01b815260048101839052602401610491565b6001600160a01b038216610b705781604051630b61174360e31b815260040161049191906114bc565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b0383163b156106b457604051630a85bd0160e11b81526001600160a01b0384169063150b7a0290610c1f9033908890879087906004016118df565b6020604051808303816000875af1925050508015610c5a575060408051601f3d908101601f19168201909252610c579181019061191c565b60015b610cba573d808015610c88576040519150601f19603f3d011682016040523d82523d6000602084013e610c8d565b606091505b508051600003610cb25783604051633250574960e11b815260040161049191906114bc565b805181602001fd5b6001600160e01b03198116630a85bd0160e11b14610ced5783604051633250574960e11b815260040161049191906114bc565b5050505050565b61063c828260405180602001604052806000815250610f4b565b60006001600160e01b0319821663780e9d6360e01b14806103ab57506103ab82610f62565b6000908152600260205260409020546001600160a01b031690565b8080610d6257506001600160a01b03821615155b15610e1b576000610d728461099e565b90506001600160a01b03831615801590610d9e5750826001600160a01b0316816001600160a01b031614155b8015610db15750610daf818461094b565b155b15610dd1578260405163a9fbf51f60e01b815260040161049191906114bc565b8115610e195783856001600160a01b0316826001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b5050600090815260046020526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b610e55828261082c565b61063c57808260405163e2517d3f60e01b8152600401610491929190611788565b600080610e84858585610fb2565b90506001600160a01b038116610ee157610edc84600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b610f04565b846001600160a01b0316816001600160a01b031614610f0457610f0481856110a6565b6001600160a01b038516610f2057610f1b84611137565b610f43565b846001600160a01b0316816001600160a01b031614610f4357610f4385856111e6565b949350505050565b610f558383611236565b6107446000848484610bdd565b60006001600160e01b031982166380ac58cd60e01b1480610f9357506001600160e01b03198216635b5e139f60e01b145b806103ab57506301ffc9a760e01b6001600160e01b03198316146103ab565b600080610fbe84610d33565b90506001600160a01b03831615610fda57610fda81848661129b565b6001600160a01b0381161561101857610ff7600085600080610d4e565b6001600160a01b038116600090815260036020526040902080546000190190555b6001600160a01b03851615611047576001600160a01b0385166000908152600360205260409020805460010190555b60008481526002602052604080822080546001600160a01b0319166001600160a01b0389811691821790925591518793918516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4949350505050565b60006110b1836107e4565b600083815260076020526040902054909150808214611104576001600160a01b03841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b5060009182526007602090815260408084208490556001600160a01b039094168352600681528383209183525290812055565b60085460009061114990600190611939565b6000838152600960205260408120546008805493945090928490811061117157611171611772565b90600052602060002001549050806008838154811061119257611192611772565b60009182526020808320909101929092558281526009909152604080822084905585825281205560088054806111ca576111ca61194c565b6001900381819060005260206000200160009055905550505050565b600060016111f3846107e4565b6111fd9190611939565b6001600160a01b039093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b6001600160a01b038216611260576000604051633250574960e11b815260040161049191906114bc565b600061126e83836000610e76565b90506001600160a01b038116156107445760006040516339e3563760e11b815260040161049191906114bc565b6112a68383836112f1565b610744576001600160a01b0383166112d457604051637e27328960e01b815260048101829052602401610491565b818160405163177e802f60e01b8152600401610491929190611788565b60006001600160a01b03831615801590610f435750826001600160a01b0316846001600160a01b0316148061132b575061132b848461094b565b80610f435750826001600160a01b0316611344836109d6565b6001600160a01b031614949350505050565b6001600160e01b031981168114610a0857600080fd5b60006020828403121561137e57600080fd5b813561138981611356565b9392505050565b6000815180845260005b818110156113b65760208185018101518683018201520161139a565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006113896020830184611390565b80356001600160a01b038116811461140057600080fd5b919050565b60006020828403121561141757600080fd5b611389826113e9565b600060208083018184528085518083526040925060408601915060408160051b87010184880160005b8381101561149557603f19898403018552815160608151855288820151818a87015261147782870182611390565b92890151958901959095525094870194925090860190600101611449565b509098975050505050505050565b6000602082840312156114b557600080fd5b5035919050565b6001600160a01b0391909116815260200190565b600080604083850312156114e357600080fd5b6114ec836113e9565b946020939093013593505050565b60008060006060848603121561150f57600080fd5b611518846113e9565b9250611526602085016113e9565b9150604084013590509250925092565b6000806040838503121561154957600080fd5b82359150611559602084016113e9565b90509250929050565b6000806040838503121561157557600080fd5b61157e836113e9565b91506020830135801515811461159357600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b60006001600160401b03808411156115ce576115ce61159e565b604051601f8501601f19908116603f011681019082821181831017156115f6576115f661159e565b8160405280935085815286868601111561160f57600080fd5b858560208301376000602087830101525050509392505050565b6000806000806080858703121561163f57600080fd5b611648856113e9565b9350611656602086016113e9565b92506040850135915060608501356001600160401b0381111561167857600080fd5b8501601f8101871361168957600080fd5b611698878235602084016115b4565b91505092959194509250565b6000806000606084860312156116b957600080fd5b6116c2846113e9565b925060208401356001600160401b038111156116dd57600080fd5b8401601f810186136116ee57600080fd5b6116fd868235602084016115b4565b925050604084013590509250925092565b6000806040838503121561172157600080fd5b61172a836113e9565b9150611559602084016113e9565b600181811c9082168061174c57607f821691505b60208210810361176c57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603260045260246000fd5b6001600160a01b03929092168252602082015260400190565b601f821115610744576000816000526020600020601f850160051c810160208610156117ca5750805b601f850160051c820191505b818110156117e9578281556001016117d6565b505050505050565b81516001600160401b0381111561180a5761180a61159e565b61181e816118188454611738565b846117a1565b602080601f831160018114611853576000841561183b5750858301515b600019600386901b1c1916600185901b1785556117e9565b600085815260208120601f198616915b8281101561188257888601518255948401946001909101908401611863565b50858210156118a05787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052601160045260246000fd5b6000600182016118d8576118d86118b0565b5060010190565b6001600160a01b038581168252841660208201526040810183905260806060820181905260009061191290830184611390565b9695505050505050565b60006020828403121561192e57600080fd5b815161138981611356565b818103818111156103ab576103ab6118b0565b634e487b7160e01b600052603160045260246000fdfe68747470733a2f2f617277656176652e6e65742f51687343554c49537655703043325467564c61516c4250626e30306c37435a6f544e4773307a3539427a559f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6a2646970667358221220a0a3d2c2e5f1ec6d3c0d788a5c59ea4db3986950ac804a01663d856d0a3db22f64736f6c63430008180033",
}

// VerificationV0ABI is the input ABI used to generate the binding from.
// Deprecated: Use VerificationV0MetaData.ABI instead.
var VerificationV0ABI = VerificationV0MetaData.ABI

// VerificationV0Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerificationV0MetaData.Bin instead.
var VerificationV0Bin = VerificationV0MetaData.Bin

// DeployVerificationV0 deploys a new Ethereum contract, binding an instance of VerificationV0 to it.
func DeployVerificationV0(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, typeName string) (common.Address, *types.Transaction, *VerificationV0, error) {
	parsed, err := VerificationV0MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerificationV0Bin), backend, name, symbol, typeName)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerificationV0{VerificationV0Caller: VerificationV0Caller{contract: contract}, VerificationV0Transactor: VerificationV0Transactor{contract: contract}, VerificationV0Filterer: VerificationV0Filterer{contract: contract}}, nil
}

// VerificationV0 is an auto generated Go binding around an Ethereum contract.
type VerificationV0 struct {
	VerificationV0Caller     // Read-only binding to the contract
	VerificationV0Transactor // Write-only binding to the contract
	VerificationV0Filterer   // Log filterer for contract events
}

// VerificationV0Caller is an auto generated read-only Go binding around an Ethereum contract.
type VerificationV0Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationV0Transactor is an auto generated write-only Go binding around an Ethereum contract.
type VerificationV0Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationV0Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerificationV0Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationV0Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerificationV0Session struct {
	Contract     *VerificationV0   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerificationV0CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerificationV0CallerSession struct {
	Contract *VerificationV0Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VerificationV0TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerificationV0TransactorSession struct {
	Contract     *VerificationV0Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VerificationV0Raw is an auto generated low-level Go binding around an Ethereum contract.
type VerificationV0Raw struct {
	Contract *VerificationV0 // Generic contract binding to access the raw methods on
}

// VerificationV0CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerificationV0CallerRaw struct {
	Contract *VerificationV0Caller // Generic read-only contract binding to access the raw methods on
}

// VerificationV0TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerificationV0TransactorRaw struct {
	Contract *VerificationV0Transactor // Generic write-only contract binding to access the raw methods on
}

// NewVerificationV0 creates a new instance of VerificationV0, bound to a specific deployed contract.
func NewVerificationV0(address common.Address, backend bind.ContractBackend) (*VerificationV0, error) {
	contract, err := bindVerificationV0(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerificationV0{VerificationV0Caller: VerificationV0Caller{contract: contract}, VerificationV0Transactor: VerificationV0Transactor{contract: contract}, VerificationV0Filterer: VerificationV0Filterer{contract: contract}}, nil
}

// NewVerificationV0Caller creates a new read-only instance of VerificationV0, bound to a specific deployed contract.
func NewVerificationV0Caller(address common.Address, caller bind.ContractCaller) (*VerificationV0Caller, error) {
	contract, err := bindVerificationV0(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerificationV0Caller{contract: contract}, nil
}

// NewVerificationV0Transactor creates a new write-only instance of VerificationV0, bound to a specific deployed contract.
func NewVerificationV0Transactor(address common.Address, transactor bind.ContractTransactor) (*VerificationV0Transactor, error) {
	contract, err := bindVerificationV0(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerificationV0Transactor{contract: contract}, nil
}

// NewVerificationV0Filterer creates a new log filterer instance of VerificationV0, bound to a specific deployed contract.
func NewVerificationV0Filterer(address common.Address, filterer bind.ContractFilterer) (*VerificationV0Filterer, error) {
	contract, err := bindVerificationV0(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerificationV0Filterer{contract: contract}, nil
}

// bindVerificationV0 binds a generic wrapper to an already deployed contract.
func bindVerificationV0(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerificationV0MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerificationV0 *VerificationV0Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerificationV0.Contract.VerificationV0Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerificationV0 *VerificationV0Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerificationV0.Contract.VerificationV0Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerificationV0 *VerificationV0Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerificationV0.Contract.VerificationV0Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerificationV0 *VerificationV0CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerificationV0.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerificationV0 *VerificationV0TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerificationV0.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerificationV0 *VerificationV0TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerificationV0.Contract.contract.Transact(opts, method, params...)
}

// BASETOKENURI is a free data retrieval call binding the contract method 0x6f458133.
//
// Solidity: function BASE_TOKEN_URI() view returns(string)
func (_VerificationV0 *VerificationV0Caller) BASETOKENURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "BASE_TOKEN_URI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BASETOKENURI is a free data retrieval call binding the contract method 0x6f458133.
//
// Solidity: function BASE_TOKEN_URI() view returns(string)
func (_VerificationV0 *VerificationV0Session) BASETOKENURI() (string, error) {
	return _VerificationV0.Contract.BASETOKENURI(&_VerificationV0.CallOpts)
}

// BASETOKENURI is a free data retrieval call binding the contract method 0x6f458133.
//
// Solidity: function BASE_TOKEN_URI() view returns(string)
func (_VerificationV0 *VerificationV0CallerSession) BASETOKENURI() (string, error) {
	return _VerificationV0.Contract.BASETOKENURI(&_VerificationV0.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VerificationV0 *VerificationV0Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VerificationV0 *VerificationV0Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _VerificationV0.Contract.DEFAULTADMINROLE(&_VerificationV0.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VerificationV0 *VerificationV0CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VerificationV0.Contract.DEFAULTADMINROLE(&_VerificationV0.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_VerificationV0 *VerificationV0Caller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_VerificationV0 *VerificationV0Session) MINTERROLE() ([32]byte, error) {
	return _VerificationV0.Contract.MINTERROLE(&_VerificationV0.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_VerificationV0 *VerificationV0CallerSession) MINTERROLE() ([32]byte, error) {
	return _VerificationV0.Contract.MINTERROLE(&_VerificationV0.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_VerificationV0 *VerificationV0Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_VerificationV0 *VerificationV0Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _VerificationV0.Contract.BalanceOf(&_VerificationV0.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_VerificationV0 *VerificationV0CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _VerificationV0.Contract.BalanceOf(&_VerificationV0.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_VerificationV0 *VerificationV0Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_VerificationV0 *VerificationV0Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _VerificationV0.Contract.GetApproved(&_VerificationV0.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_VerificationV0 *VerificationV0CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _VerificationV0.Contract.GetApproved(&_VerificationV0.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VerificationV0 *VerificationV0Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VerificationV0 *VerificationV0Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VerificationV0.Contract.GetRoleAdmin(&_VerificationV0.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VerificationV0 *VerificationV0CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VerificationV0.Contract.GetRoleAdmin(&_VerificationV0.CallOpts, role)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() view returns(string)
func (_VerificationV0 *VerificationV0Caller) GetTypeName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "getTypeName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() view returns(string)
func (_VerificationV0 *VerificationV0Session) GetTypeName() (string, error) {
	return _VerificationV0.Contract.GetTypeName(&_VerificationV0.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() view returns(string)
func (_VerificationV0 *VerificationV0CallerSession) GetTypeName() (string, error) {
	return _VerificationV0.Contract.GetTypeName(&_VerificationV0.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VerificationV0 *VerificationV0Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VerificationV0 *VerificationV0Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VerificationV0.Contract.HasRole(&_VerificationV0.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VerificationV0 *VerificationV0CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VerificationV0.Contract.HasRole(&_VerificationV0.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_VerificationV0 *VerificationV0Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_VerificationV0 *VerificationV0Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _VerificationV0.Contract.IsApprovedForAll(&_VerificationV0.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_VerificationV0 *VerificationV0CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _VerificationV0.Contract.IsApprovedForAll(&_VerificationV0.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VerificationV0 *VerificationV0Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VerificationV0 *VerificationV0Session) Name() (string, error) {
	return _VerificationV0.Contract.Name(&_VerificationV0.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VerificationV0 *VerificationV0CallerSession) Name() (string, error) {
	return _VerificationV0.Contract.Name(&_VerificationV0.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_VerificationV0 *VerificationV0Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_VerificationV0 *VerificationV0Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _VerificationV0.Contract.OwnerOf(&_VerificationV0.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_VerificationV0 *VerificationV0CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _VerificationV0.Contract.OwnerOf(&_VerificationV0.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VerificationV0 *VerificationV0Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VerificationV0 *VerificationV0Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VerificationV0.Contract.SupportsInterface(&_VerificationV0.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VerificationV0 *VerificationV0CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VerificationV0.Contract.SupportsInterface(&_VerificationV0.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VerificationV0 *VerificationV0Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VerificationV0 *VerificationV0Session) Symbol() (string, error) {
	return _VerificationV0.Contract.Symbol(&_VerificationV0.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VerificationV0 *VerificationV0CallerSession) Symbol() (string, error) {
	return _VerificationV0.Contract.Symbol(&_VerificationV0.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_VerificationV0 *VerificationV0Caller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_VerificationV0 *VerificationV0Session) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _VerificationV0.Contract.TokenByIndex(&_VerificationV0.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_VerificationV0 *VerificationV0CallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _VerificationV0.Contract.TokenByIndex(&_VerificationV0.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_VerificationV0 *VerificationV0Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_VerificationV0 *VerificationV0Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _VerificationV0.Contract.TokenOfOwnerByIndex(&_VerificationV0.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_VerificationV0 *VerificationV0CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _VerificationV0.Contract.TokenOfOwnerByIndex(&_VerificationV0.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_VerificationV0 *VerificationV0Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_VerificationV0 *VerificationV0Session) TokenURI(tokenId *big.Int) (string, error) {
	return _VerificationV0.Contract.TokenURI(&_VerificationV0.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_VerificationV0 *VerificationV0CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _VerificationV0.Contract.TokenURI(&_VerificationV0.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VerificationV0 *VerificationV0Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VerificationV0 *VerificationV0Session) TotalSupply() (*big.Int, error) {
	return _VerificationV0.Contract.TotalSupply(&_VerificationV0.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VerificationV0 *VerificationV0CallerSession) TotalSupply() (*big.Int, error) {
	return _VerificationV0.Contract.TotalSupply(&_VerificationV0.CallOpts)
}

// TransferFrom is a free data retrieval call binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) pure returns()
func (_VerificationV0 *VerificationV0Caller) TransferFrom(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) error {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "transferFrom", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// TransferFrom is a free data retrieval call binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) pure returns()
func (_VerificationV0 *VerificationV0Session) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) error {
	return _VerificationV0.Contract.TransferFrom(&_VerificationV0.CallOpts, arg0, arg1, arg2)
}

// TransferFrom is a free data retrieval call binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) pure returns()
func (_VerificationV0 *VerificationV0CallerSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) error {
	return _VerificationV0.Contract.TransferFrom(&_VerificationV0.CallOpts, arg0, arg1, arg2)
}

// Verifications is a free data retrieval call binding the contract method 0x080c98da.
//
// Solidity: function verifications(address owner) view returns((uint256,string,uint256)[])
func (_VerificationV0 *VerificationV0Caller) Verifications(opts *bind.CallOpts, owner common.Address) ([]StructsVerificationV0, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "verifications", owner)

	if err != nil {
		return *new([]StructsVerificationV0), err
	}

	out0 := *abi.ConvertType(out[0], new([]StructsVerificationV0)).(*[]StructsVerificationV0)

	return out0, err

}

// Verifications is a free data retrieval call binding the contract method 0x080c98da.
//
// Solidity: function verifications(address owner) view returns((uint256,string,uint256)[])
func (_VerificationV0 *VerificationV0Session) Verifications(owner common.Address) ([]StructsVerificationV0, error) {
	return _VerificationV0.Contract.Verifications(&_VerificationV0.CallOpts, owner)
}

// Verifications is a free data retrieval call binding the contract method 0x080c98da.
//
// Solidity: function verifications(address owner) view returns((uint256,string,uint256)[])
func (_VerificationV0 *VerificationV0CallerSession) Verifications(owner common.Address) ([]StructsVerificationV0, error) {
	return _VerificationV0.Contract.Verifications(&_VerificationV0.CallOpts, owner)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_VerificationV0 *VerificationV0Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerificationV0.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_VerificationV0 *VerificationV0Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerificationV0.Contract.Approve(&_VerificationV0.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_VerificationV0 *VerificationV0TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerificationV0.Contract.Approve(&_VerificationV0.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_VerificationV0 *VerificationV0Transactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _VerificationV0.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_VerificationV0 *VerificationV0Session) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _VerificationV0.Contract.Burn(&_VerificationV0.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_VerificationV0 *VerificationV0TransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _VerificationV0.Contract.Burn(&_VerificationV0.TransactOpts, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VerificationV0 *VerificationV0Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VerificationV0.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VerificationV0 *VerificationV0Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VerificationV0.Contract.GrantRole(&_VerificationV0.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VerificationV0 *VerificationV0TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VerificationV0.Contract.GrantRole(&_VerificationV0.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0xba7aef43.
//
// Solidity: function mint(address to, string service, uint256 timestamp) returns()
func (_VerificationV0 *VerificationV0Transactor) Mint(opts *bind.TransactOpts, to common.Address, service string, timestamp *big.Int) (*types.Transaction, error) {
	return _VerificationV0.contract.Transact(opts, "mint", to, service, timestamp)
}

// Mint is a paid mutator transaction binding the contract method 0xba7aef43.
//
// Solidity: function mint(address to, string service, uint256 timestamp) returns()
func (_VerificationV0 *VerificationV0Session) Mint(to common.Address, service string, timestamp *big.Int) (*types.Transaction, error) {
	return _VerificationV0.Contract.Mint(&_VerificationV0.TransactOpts, to, service, timestamp)
}

// Mint is a paid mutator transaction binding the contract method 0xba7aef43.
//
// Solidity: function mint(address to, string service, uint256 timestamp) returns()
func (_VerificationV0 *VerificationV0TransactorSession) Mint(to common.Address, service string, timestamp *big.Int) (*types.Transaction, error) {
	return _VerificationV0.Contract.Mint(&_VerificationV0.TransactOpts, to, service, timestamp)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_VerificationV0 *VerificationV0Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _VerificationV0.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_VerificationV0 *VerificationV0Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _VerificationV0.Contract.RenounceRole(&_VerificationV0.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_VerificationV0 *VerificationV0TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _VerificationV0.Contract.RenounceRole(&_VerificationV0.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VerificationV0 *VerificationV0Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VerificationV0.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VerificationV0 *VerificationV0Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VerificationV0.Contract.RevokeRole(&_VerificationV0.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VerificationV0 *VerificationV0TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VerificationV0.Contract.RevokeRole(&_VerificationV0.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_VerificationV0 *VerificationV0Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerificationV0.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_VerificationV0 *VerificationV0Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerificationV0.Contract.SafeTransferFrom(&_VerificationV0.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_VerificationV0 *VerificationV0TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerificationV0.Contract.SafeTransferFrom(&_VerificationV0.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_VerificationV0 *VerificationV0Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _VerificationV0.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_VerificationV0 *VerificationV0Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _VerificationV0.Contract.SafeTransferFrom0(&_VerificationV0.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_VerificationV0 *VerificationV0TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _VerificationV0.Contract.SafeTransferFrom0(&_VerificationV0.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_VerificationV0 *VerificationV0Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _VerificationV0.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_VerificationV0 *VerificationV0Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _VerificationV0.Contract.SetApprovalForAll(&_VerificationV0.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_VerificationV0 *VerificationV0TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _VerificationV0.Contract.SetApprovalForAll(&_VerificationV0.TransactOpts, operator, approved)
}

// VerificationV0ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the VerificationV0 contract.
type VerificationV0ApprovalIterator struct {
	Event *VerificationV0Approval // Event containing the contract specifics and raw log

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
func (it *VerificationV0ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerificationV0Approval)
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
		it.Event = new(VerificationV0Approval)
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
func (it *VerificationV0ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerificationV0ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerificationV0Approval represents a Approval event raised by the VerificationV0 contract.
type VerificationV0Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_VerificationV0 *VerificationV0Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*VerificationV0ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VerificationV0.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VerificationV0ApprovalIterator{contract: _VerificationV0.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_VerificationV0 *VerificationV0Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VerificationV0Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VerificationV0.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerificationV0Approval)
				if err := _VerificationV0.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_VerificationV0 *VerificationV0Filterer) ParseApproval(log types.Log) (*VerificationV0Approval, error) {
	event := new(VerificationV0Approval)
	if err := _VerificationV0.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerificationV0ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the VerificationV0 contract.
type VerificationV0ApprovalForAllIterator struct {
	Event *VerificationV0ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *VerificationV0ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerificationV0ApprovalForAll)
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
		it.Event = new(VerificationV0ApprovalForAll)
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
func (it *VerificationV0ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerificationV0ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerificationV0ApprovalForAll represents a ApprovalForAll event raised by the VerificationV0 contract.
type VerificationV0ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_VerificationV0 *VerificationV0Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*VerificationV0ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _VerificationV0.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &VerificationV0ApprovalForAllIterator{contract: _VerificationV0.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_VerificationV0 *VerificationV0Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *VerificationV0ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _VerificationV0.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerificationV0ApprovalForAll)
				if err := _VerificationV0.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_VerificationV0 *VerificationV0Filterer) ParseApprovalForAll(log types.Log) (*VerificationV0ApprovalForAll, error) {
	event := new(VerificationV0ApprovalForAll)
	if err := _VerificationV0.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerificationV0RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the VerificationV0 contract.
type VerificationV0RoleAdminChangedIterator struct {
	Event *VerificationV0RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *VerificationV0RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerificationV0RoleAdminChanged)
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
		it.Event = new(VerificationV0RoleAdminChanged)
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
func (it *VerificationV0RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerificationV0RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerificationV0RoleAdminChanged represents a RoleAdminChanged event raised by the VerificationV0 contract.
type VerificationV0RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VerificationV0 *VerificationV0Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VerificationV0RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _VerificationV0.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VerificationV0RoleAdminChangedIterator{contract: _VerificationV0.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VerificationV0 *VerificationV0Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VerificationV0RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _VerificationV0.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerificationV0RoleAdminChanged)
				if err := _VerificationV0.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VerificationV0 *VerificationV0Filterer) ParseRoleAdminChanged(log types.Log) (*VerificationV0RoleAdminChanged, error) {
	event := new(VerificationV0RoleAdminChanged)
	if err := _VerificationV0.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerificationV0RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the VerificationV0 contract.
type VerificationV0RoleGrantedIterator struct {
	Event *VerificationV0RoleGranted // Event containing the contract specifics and raw log

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
func (it *VerificationV0RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerificationV0RoleGranted)
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
		it.Event = new(VerificationV0RoleGranted)
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
func (it *VerificationV0RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerificationV0RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerificationV0RoleGranted represents a RoleGranted event raised by the VerificationV0 contract.
type VerificationV0RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VerificationV0 *VerificationV0Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VerificationV0RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VerificationV0.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VerificationV0RoleGrantedIterator{contract: _VerificationV0.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VerificationV0 *VerificationV0Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VerificationV0RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VerificationV0.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerificationV0RoleGranted)
				if err := _VerificationV0.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VerificationV0 *VerificationV0Filterer) ParseRoleGranted(log types.Log) (*VerificationV0RoleGranted, error) {
	event := new(VerificationV0RoleGranted)
	if err := _VerificationV0.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerificationV0RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the VerificationV0 contract.
type VerificationV0RoleRevokedIterator struct {
	Event *VerificationV0RoleRevoked // Event containing the contract specifics and raw log

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
func (it *VerificationV0RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerificationV0RoleRevoked)
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
		it.Event = new(VerificationV0RoleRevoked)
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
func (it *VerificationV0RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerificationV0RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerificationV0RoleRevoked represents a RoleRevoked event raised by the VerificationV0 contract.
type VerificationV0RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VerificationV0 *VerificationV0Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VerificationV0RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VerificationV0.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VerificationV0RoleRevokedIterator{contract: _VerificationV0.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VerificationV0 *VerificationV0Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VerificationV0RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VerificationV0.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerificationV0RoleRevoked)
				if err := _VerificationV0.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VerificationV0 *VerificationV0Filterer) ParseRoleRevoked(log types.Log) (*VerificationV0RoleRevoked, error) {
	event := new(VerificationV0RoleRevoked)
	if err := _VerificationV0.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerificationV0TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the VerificationV0 contract.
type VerificationV0TransferIterator struct {
	Event *VerificationV0Transfer // Event containing the contract specifics and raw log

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
func (it *VerificationV0TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerificationV0Transfer)
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
		it.Event = new(VerificationV0Transfer)
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
func (it *VerificationV0TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerificationV0TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerificationV0Transfer represents a Transfer event raised by the VerificationV0 contract.
type VerificationV0Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_VerificationV0 *VerificationV0Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*VerificationV0TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VerificationV0.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VerificationV0TransferIterator{contract: _VerificationV0.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_VerificationV0 *VerificationV0Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VerificationV0Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VerificationV0.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerificationV0Transfer)
				if err := _VerificationV0.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_VerificationV0 *VerificationV0Filterer) ParseTransfer(log types.Log) (*VerificationV0Transfer, error) {
	event := new(VerificationV0Transfer)
	if err := _VerificationV0.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
