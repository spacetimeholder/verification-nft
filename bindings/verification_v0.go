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

// StructsVerificationData is an auto generated low-level Go binding around an user-defined struct.
type StructsVerificationData struct {
	TokenId   *big.Int
	Service   string
	Timestamp *big.Int
}

// VerificationV0MetaData contains all meta data concerning the VerificationV0 contract.
var VerificationV0MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASE_TOKEN_URI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"service\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"verifications\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"service\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.VerificationData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526001600b553480156200001657600080fd5b5060405162001f3e38038062001f3e83398101604081905262000039916200022b565b828260006200004983826200034d565b5060016200005882826200034d565b50600c91506200006b905082826200034d565b5062000079600033620000b0565b50620000a67f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a633620000b0565b5050505062000419565b6000828152600a602090815260408083206001600160a01b038516845290915281205460ff1662000159576000838152600a602090815260408083206001600160a01b03861684529091529020805460ff19166001179055620001103390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016200015d565b5060005b92915050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200018b57600080fd5b81516001600160401b0380821115620001a857620001a862000163565b604051601f8301601f19908116603f01168101908282118183101715620001d357620001d362000163565b8160405283815260209250866020858801011115620001f157600080fd5b600091505b83821015620002155785820183015181830184015290820190620001f6565b6000602085830101528094505050505092915050565b6000806000606084860312156200024157600080fd5b83516001600160401b03808211156200025957600080fd5b620002678783880162000179565b945060208601519150808211156200027e57600080fd5b6200028c8783880162000179565b93506040860151915080821115620002a357600080fd5b50620002b28682870162000179565b9150509250925092565b600181811c90821680620002d157607f821691505b602082108103620002f257634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111562000348576000816000526020600020601f850160051c81016020861015620003235750805b601f850160051c820191505b8181101562000344578281556001016200032f565b5050505b505050565b81516001600160401b0381111562000369576200036962000163565b62000381816200037a8454620002bc565b84620002f8565b602080601f831160018114620003b95760008415620003a05750858301515b600019600386901b1c1916600185901b17855562000344565b600085815260208120601f198616915b82811015620003ea57888601518255948401946001909101908401620003c9565b5085821015620004095787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b611b1580620004296000396000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c80636352211e116100f9578063b88d4fde11610097578063d539139311610071578063d5391393146103e8578063d547741f1461040f578063db0a087c14610422578063e985e9c51461042a57600080fd5b8063b88d4fde146103af578063ba7aef43146103c2578063c87b56dd146103d557600080fd5b806391d14854116100d357806391d148541461037957806395d89b411461038c578063a217fddf14610394578063a22cb4651461039c57600080fd5b80636352211e146103205780636f4581331461033357806370a082311461036657600080fd5b8063248a9ca31161016657806336568abe1161014057806336568abe146102d457806342842e0e146102e757806342966c68146102fa5780634f6ccce71461030d57600080fd5b8063248a9ca31461028b5780632f2ff15d146102ae5780632f745c59146102c157600080fd5b8063081812fc116101a2578063081812fc14610226578063095ea7b31461025157806318160ddd1461026657806323b872dd1461027857600080fd5b806301ffc9a7146101c957806306fdde03146101f1578063080c98da14610206575b600080fd5b6101dc6101d73660046114d3565b61043d565b60405190151581526020015b60405180910390f35b6101f961044e565b6040516101e8919061153d565b61021961021436600461156c565b6104e0565b6040516101e89190611587565b61023961023436600461160a565b6106ba565b6040516001600160a01b0390911681526020016101e8565b61026461025f366004611623565b6106e3565b005b6008545b6040519081526020016101e8565b61026461028636600461164d565b6106f2565b61026a61029936600461160a565b6000908152600a602052604090206001015490565b6102646102bc366004611689565b61077d565b61026a6102cf366004611623565b6107a2565b6102646102e2366004611689565b610807565b6102646102f536600461164d565b61083f565b61026461030836600461160a565b61085a565b61026a61031b36600461160a565b61088d565b61023961032e36600461160a565b6108e6565b6101f96040518060400160405280601481526020017368747470733a2f2f617277656176652e6e65742f60601b81525081565b61026a61037436600461156c565b6108f1565b6101dc610387366004611689565b610939565b6101f9610964565b61026a600081565b6102646103aa3660046116b5565b610973565b6102646103bd36600461177d565b61097e565b6102646103d03660046117f9565b6109c6565b6101f96103e336600461160a565b610a45565b61026a7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b61026461041d366004611689565b610a70565b6101f9610a95565b6101dc610438366004611864565b610aa4565b600061044882610ad2565b92915050565b60606000805461045d9061188e565b80601f01602080910402602001604051908101604052809291908181526020018280546104899061188e565b80156104d65780601f106104ab576101008083540402835291602001916104d6565b820191906000526020600020905b8154815290600101906020018083116104b957829003601f168201915b5050505050905090565b606060006104ed836108f1565b9050600081116105375760405162461bcd60e51b815260206004820152601060248201526f6e6f20766572696669636174696f6e7360801b60448201526064015b60405180910390fd5b60008167ffffffffffffffff811115610552576105526116f1565b6040519080825280602002602001820160405280156105a757816020015b61059460405180606001604052806000815260200160608152602001600081525090565b8152602001906001900390816105705790505b50905060005b828110156106b25760006105c186836107a2565b90506040518060600160405280828152602001600d600084815260200190815260200160002080546105f29061188e565b80601f016020809104026020016040519081016040528092919081815260200182805461061e9061188e565b801561066b5780601f106106405761010080835404028352916020019161066b565b820191906000526020600020905b81548152906001019060200180831161064e57829003601f168201915b50505050508152602001600e60008481526020019081526020016000205481525083838151811061069e5761069e6118c8565b6020908102919091010152506001016105ad565b509392505050565b60006106c582610af7565b506000828152600460205260409020546001600160a01b0316610448565b6106ee828233610b30565b5050565b6001600160a01b03821661071c57604051633250574960e11b81526000600482015260240161052e565b6000610729838333610b3d565b9050836001600160a01b0316816001600160a01b031614610777576040516364283d7b60e01b81526001600160a01b038086166004830152602482018490528216604482015260640161052e565b50505050565b6000828152600a602052604090206001015461079881610c12565b6107778383610c1f565b60006107ad836108f1565b82106107de5760405163295f44f760e21b81526001600160a01b03841660048201526024810183905260440161052e565b506001600160a01b03919091166000908152600660209081526040808320938352929052205490565b6001600160a01b03811633146108305760405163334bd91960e11b815260040160405180910390fd5b61083a8282610cb3565b505050565b61083a8383836040518060200160405280600081525061097e565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a661088481610c12565b6106ee82610d20565b600061089860085490565b82106108c15760405163295f44f760e21b8152600060048201526024810183905260440161052e565b600882815481106108d4576108d46118c8565b90600052602060002001549050919050565b600061044882610af7565b60006001600160a01b03821661091d576040516322718ad960e21b81526000600482015260240161052e565b506001600160a01b031660009081526003602052604090205490565b6000918252600a602090815260408084206001600160a01b0393909316845291905290205460ff1690565b60606001805461045d9061188e565b6106ee338383610d5b565b60405162461bcd60e51b815260206004820152601860248201527f746f6b656e73206e6f6e2d7472616e736665727261626c650000000000000000604482015260640161052e565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a66109f081610c12565b6109fc84600b54610dfa565b600b546000908152600d60205260409020610a17848261192e565b50600b80546000908152600e6020526040812084905581549190610a3a83611a04565b919050555050505050565b6060610a5082610af7565b506040518060600160405280603f8152602001611aa1603f913992915050565b6000828152600a6020526040902060010154610a8b81610c12565b6107778383610cb3565b6060600c805461045d9061188e565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b60006001600160e01b03198216637965db0b60e01b1480610448575061044882610e14565b6000818152600260205260408120546001600160a01b03168061044857604051637e27328960e01b81526004810184905260240161052e565b61083a8383836001610e39565b600080610b4b858585610f3f565b90506001600160a01b038116610ba857610ba384600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b610bcb565b846001600160a01b0316816001600160a01b031614610bcb57610bcb8185611038565b6001600160a01b038516610be757610be2846110c9565b610c0a565b846001600160a01b0316816001600160a01b031614610c0a57610c0a8585611178565b949350505050565b610c1c81336111c8565b50565b6000610c2b8383610939565b610cab576000838152600a602090815260408083206001600160a01b03861684529091529020805460ff19166001179055610c633390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610448565b506000610448565b6000610cbf8383610939565b15610cab576000838152600a602090815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610448565b6000610d2f6000836000610b3d565b90506001600160a01b0381166106ee57604051637e27328960e01b81526004810183905260240161052e565b6001600160a01b038216610d8d57604051630b61174360e31b81526001600160a01b038316600482015260240161052e565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6106ee828260405180602001604052806000815250611201565b60006001600160e01b0319821663780e9d6360e01b1480610448575061044882611218565b8080610e4d57506001600160a01b03821615155b15610f0f576000610e5d84610af7565b90506001600160a01b03831615801590610e895750826001600160a01b0316816001600160a01b031614155b8015610e9c5750610e9a8184610aa4565b155b15610ec55760405163a9fbf51f60e01b81526001600160a01b038416600482015260240161052e565b8115610f0d5783856001600160a01b0316826001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b5050600090815260046020526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b6000828152600260205260408120546001600160a01b0390811690831615610f6c57610f6c818486611268565b6001600160a01b03811615610faa57610f89600085600080610e39565b6001600160a01b038116600090815260036020526040902080546000190190555b6001600160a01b03851615610fd9576001600160a01b0385166000908152600360205260409020805460010190555b60008481526002602052604080822080546001600160a01b0319166001600160a01b0389811691821790925591518793918516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4949350505050565b6000611043836108f1565b600083815260076020526040902054909150808214611096576001600160a01b03841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b5060009182526007602090815260408084208490556001600160a01b039094168352600681528383209183525290812055565b6008546000906110db90600190611a1d565b60008381526009602052604081205460088054939450909284908110611103576111036118c8565b906000526020600020015490508060088381548110611124576111246118c8565b600091825260208083209091019290925582815260099091526040808220849055858252812055600880548061115c5761115c611a30565b6001900381819060005260206000200160009055905550505050565b60006001611185846108f1565b61118f9190611a1d565b6001600160a01b039093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b6111d28282610939565b6106ee5760405163e2517d3f60e01b81526001600160a01b03821660048201526024810183905260440161052e565b61120b83836112cc565b61083a6000848484611331565b60006001600160e01b031982166380ac58cd60e01b148061124957506001600160e01b03198216635b5e139f60e01b145b8061044857506301ffc9a760e01b6001600160e01b0319831614610448565b61127383838361145a565b61083a576001600160a01b0383166112a157604051637e27328960e01b81526004810182905260240161052e565b60405163177e802f60e01b81526001600160a01b03831660048201526024810182905260440161052e565b6001600160a01b0382166112f657604051633250574960e11b81526000600482015260240161052e565b600061130483836000610b3d565b90506001600160a01b0381161561083a576040516339e3563760e11b81526000600482015260240161052e565b6001600160a01b0383163b1561077757604051630a85bd0160e11b81526001600160a01b0384169063150b7a0290611373903390889087908790600401611a46565b6020604051808303816000875af19250505080156113ae575060408051601f3d908101601f191682019092526113ab91810190611a83565b60015b611417573d8080156113dc576040519150601f19603f3d011682016040523d82523d6000602084013e6113e1565b606091505b50805160000361140f57604051633250574960e11b81526001600160a01b038516600482015260240161052e565b805181602001fd5b6001600160e01b03198116630a85bd0160e11b1461145357604051633250574960e11b81526001600160a01b038516600482015260240161052e565b5050505050565b60006001600160a01b03831615801590610c0a5750826001600160a01b0316846001600160a01b0316148061149457506114948484610aa4565b80610c0a5750506000908152600460205260409020546001600160a01b03908116911614919050565b6001600160e01b031981168114610c1c57600080fd5b6000602082840312156114e557600080fd5b81356114f0816114bd565b9392505050565b6000815180845260005b8181101561151d57602081850181015186830182015201611501565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006114f060208301846114f7565b80356001600160a01b038116811461156757600080fd5b919050565b60006020828403121561157e57600080fd5b6114f082611550565b600060208083018184528085518083526040925060408601915060408160051b87010184880160005b838110156115fc57603f19898403018552815160608151855288820151818a8701526115de828701826114f7565b928901519589019590955250948701949250908601906001016115b0565b509098975050505050505050565b60006020828403121561161c57600080fd5b5035919050565b6000806040838503121561163657600080fd5b61163f83611550565b946020939093013593505050565b60008060006060848603121561166257600080fd5b61166b84611550565b925061167960208501611550565b9150604084013590509250925092565b6000806040838503121561169c57600080fd5b823591506116ac60208401611550565b90509250929050565b600080604083850312156116c857600080fd5b6116d183611550565b9150602083013580151581146116e657600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff80841115611722576117226116f1565b604051601f8501601f19908116603f0116810190828211818310171561174a5761174a6116f1565b8160405280935085815286868601111561176357600080fd5b858560208301376000602087830101525050509392505050565b6000806000806080858703121561179357600080fd5b61179c85611550565b93506117aa60208601611550565b925060408501359150606085013567ffffffffffffffff8111156117cd57600080fd5b8501601f810187136117de57600080fd5b6117ed87823560208401611707565b91505092959194509250565b60008060006060848603121561180e57600080fd5b61181784611550565b9250602084013567ffffffffffffffff81111561183357600080fd5b8401601f8101861361184457600080fd5b61185386823560208401611707565b925050604084013590509250925092565b6000806040838503121561187757600080fd5b61188083611550565b91506116ac60208401611550565b600181811c908216806118a257607f821691505b6020821081036118c257634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603260045260246000fd5b601f82111561083a576000816000526020600020601f850160051c810160208610156119075750805b601f850160051c820191505b8181101561192657828155600101611913565b505050505050565b815167ffffffffffffffff811115611948576119486116f1565b61195c81611956845461188e565b846118de565b602080601f83116001811461199157600084156119795750858301515b600019600386901b1c1916600185901b178555611926565b600085815260208120601f198616915b828110156119c0578886015182559484019460019091019084016119a1565b50858210156119de5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052601160045260246000fd5b600060018201611a1657611a166119ee565b5060010190565b81810381811115610448576104486119ee565b634e487b7160e01b600052603160045260246000fd5b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090611a79908301846114f7565b9695505050505050565b600060208284031215611a9557600080fd5b81516114f0816114bd56fe68747470733a2f2f617277656176652e6e65742f51687343554c49537655703043325467564c61516c4250626e30306c37435a6f544e4773307a3539427a55a26469706673582212201c85f796fa53737026550d0d649d2adc1ca77489194d6975e4cf3be94224826e64736f6c63430008180033",
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

// SafeTransferFrom0 is a free data retrieval call binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) pure returns()
func (_VerificationV0 *VerificationV0Caller) SafeTransferFrom0(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) error {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "safeTransferFrom0", arg0, arg1, arg2, arg3)

	if err != nil {
		return err
	}

	return err

}

// SafeTransferFrom0 is a free data retrieval call binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) pure returns()
func (_VerificationV0 *VerificationV0Session) SafeTransferFrom0(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) error {
	return _VerificationV0.Contract.SafeTransferFrom0(&_VerificationV0.CallOpts, arg0, arg1, arg2, arg3)
}

// SafeTransferFrom0 is a free data retrieval call binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) pure returns()
func (_VerificationV0 *VerificationV0CallerSession) SafeTransferFrom0(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) error {
	return _VerificationV0.Contract.SafeTransferFrom0(&_VerificationV0.CallOpts, arg0, arg1, arg2, arg3)
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

// Verifications is a free data retrieval call binding the contract method 0x080c98da.
//
// Solidity: function verifications(address owner) view returns((uint256,string,uint256)[])
func (_VerificationV0 *VerificationV0Caller) Verifications(opts *bind.CallOpts, owner common.Address) ([]StructsVerificationData, error) {
	var out []interface{}
	err := _VerificationV0.contract.Call(opts, &out, "verifications", owner)

	if err != nil {
		return *new([]StructsVerificationData), err
	}

	out0 := *abi.ConvertType(out[0], new([]StructsVerificationData)).(*[]StructsVerificationData)

	return out0, err

}

// Verifications is a free data retrieval call binding the contract method 0x080c98da.
//
// Solidity: function verifications(address owner) view returns((uint256,string,uint256)[])
func (_VerificationV0 *VerificationV0Session) Verifications(owner common.Address) ([]StructsVerificationData, error) {
	return _VerificationV0.Contract.Verifications(&_VerificationV0.CallOpts, owner)
}

// Verifications is a free data retrieval call binding the contract method 0x080c98da.
//
// Solidity: function verifications(address owner) view returns((uint256,string,uint256)[])
func (_VerificationV0 *VerificationV0CallerSession) Verifications(owner common.Address) ([]StructsVerificationData, error) {
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

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_VerificationV0 *VerificationV0Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerificationV0.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_VerificationV0 *VerificationV0Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerificationV0.Contract.TransferFrom(&_VerificationV0.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_VerificationV0 *VerificationV0TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerificationV0.Contract.TransferFrom(&_VerificationV0.TransactOpts, from, to, tokenId)
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
