// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

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

// BlockhainMetaData contains all meta data concerning the Blockhain contract.
var BlockhainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_initialWallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_supply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620029da380380620029da8339818101604052810190620000379190620005fa565b838381600390816200004a9190620008eb565b5080600490816200005c9190620008eb565b5050506200007f62000073620000e060201b60201c565b620000e860201b60201c565b6000600560146101000a81548160ff021916908315150217905550620000d682620000af620001ae60201b60201c565b600a620000bd919062000b62565b83620000ca919062000bb3565b620001b760201b60201c565b5050505062000d72565b600033905090565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60006012905090565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160362000229576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002209062000c75565b60405180910390fd5b6200023d600083836200032460201b60201c565b806002600082825462000251919062000c97565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405162000304919062000ce3565b60405180910390a362000320600083836200035160201b60201c565b5050565b620003346200035660201b60201c565b6200034c838383620003ab60201b620007f51760201c565b505050565b505050565b62000366620003b060201b60201c565b15620003a9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620003a09062000d50565b60405180910390fd5b565b505050565b6000600560149054906101000a900460ff16905090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200043082620003e5565b810181811067ffffffffffffffff82111715620004525762000451620003f6565b5b80604052505050565b600062000467620003c7565b905062000475828262000425565b919050565b600067ffffffffffffffff821115620004985762000497620003f6565b5b620004a382620003e5565b9050602081019050919050565b60005b83811015620004d0578082015181840152602081019050620004b3565b60008484015250505050565b6000620004f3620004ed846200047a565b6200045b565b905082815260208101848484011115620005125762000511620003e0565b5b6200051f848285620004b0565b509392505050565b600082601f8301126200053f576200053e620003db565b5b815162000551848260208601620004dc565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062000587826200055a565b9050919050565b62000599816200057a565b8114620005a557600080fd5b50565b600081519050620005b9816200058e565b92915050565b6000819050919050565b620005d481620005bf565b8114620005e057600080fd5b50565b600081519050620005f481620005c9565b92915050565b60008060008060808587031215620006175762000616620003d1565b5b600085015167ffffffffffffffff811115620006385762000637620003d6565b5b620006468782880162000527565b945050602085015167ffffffffffffffff8111156200066a5762000669620003d6565b5b620006788782880162000527565b93505060406200068b87828801620005a8565b92505060606200069e87828801620005e3565b91505092959194509250565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620006fd57607f821691505b602082108103620007135762000712620006b5565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026200077d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200073e565b6200078986836200073e565b95508019841693508086168417925050509392505050565b6000819050919050565b6000620007cc620007c6620007c084620005bf565b620007a1565b620005bf565b9050919050565b6000819050919050565b620007e883620007ab565b62000800620007f782620007d3565b8484546200074b565b825550505050565b600090565b6200081762000808565b62000824818484620007dd565b505050565b5b818110156200084c57620008406000826200080d565b6001810190506200082a565b5050565b601f8211156200089b57620008658162000719565b62000870846200072e565b8101602085101562000880578190505b620008986200088f856200072e565b83018262000829565b50505b505050565b600082821c905092915050565b6000620008c060001984600802620008a0565b1980831691505092915050565b6000620008db8383620008ad565b9150826002028217905092915050565b620008f682620006aa565b67ffffffffffffffff811115620009125762000911620003f6565b5b6200091e8254620006e4565b6200092b82828562000850565b600060209050601f8311600181146200096357600084156200094e578287015190505b6200095a8582620008cd565b865550620009ca565b601f198416620009738662000719565b60005b828110156200099d5784890151825560018201915060208501945060208101905062000976565b86831015620009bd5784890151620009b9601f891682620008ad565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008160011c9050919050565b6000808291508390505b600185111562000a605780860481111562000a385762000a37620009d2565b5b600185161562000a485780820291505b808102905062000a588562000a01565b945062000a18565b94509492505050565b60008262000a7b576001905062000b4e565b8162000a8b576000905062000b4e565b816001811462000aa4576002811462000aaf5762000ae5565b600191505062000b4e565b60ff84111562000ac45762000ac3620009d2565b5b8360020a91508482111562000ade5762000add620009d2565b5b5062000b4e565b5060208310610133831016604e8410600b841016171562000b1f5782820a90508381111562000b195762000b18620009d2565b5b62000b4e565b62000b2e848484600162000a0e565b9250905081840481111562000b485762000b47620009d2565b5b81810290505b9392505050565b600060ff82169050919050565b600062000b6f82620005bf565b915062000b7c8362000b55565b925062000bab7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848462000a69565b905092915050565b600062000bc082620005bf565b915062000bcd83620005bf565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161562000c095762000c08620009d2565b5b828202905092915050565b600082825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b600062000c5d601f8362000c14565b915062000c6a8262000c25565b602082019050919050565b6000602082019050818103600083015262000c908162000c4e565b9050919050565b600062000ca482620005bf565b915062000cb183620005bf565b925082820190508082111562000ccc5762000ccb620009d2565b5b92915050565b62000cdd81620005bf565b82525050565b600060208201905062000cfa600083018462000cd2565b92915050565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b600062000d3860108362000c14565b915062000d458262000d00565b602082019050919050565b6000602082019050818103600083015262000d6b8162000d29565b9050919050565b611c588062000d826000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c806370a08231116100ad57806395d89b411161007157806395d89b41146102d2578063a457c2d7146102f0578063a9059cbb14610320578063dd62ed3e14610350578063f2fde38b1461038057610121565b806370a0823114610254578063715018a61461028457806379cc67901461028e5780638456cb59146102aa5780638da5cb5b146102b457610121565b8063313ce567116100f4578063313ce567146101c257806339509351146101e05780633f4ba83a1461021057806342966c681461021a5780635c975abb1461023657610121565b806306fdde0314610126578063095ea7b31461014457806318160ddd1461017457806323b872dd14610192575b600080fd5b61012e61039c565b60405161013b91906111e4565b60405180910390f35b61015e6004803603810190610159919061129f565b61042e565b60405161016b91906112fa565b60405180910390f35b61017c610451565b6040516101899190611324565b60405180910390f35b6101ac60048036038101906101a7919061133f565b61045b565b6040516101b991906112fa565b60405180910390f35b6101ca61048a565b6040516101d791906113ae565b60405180910390f35b6101fa60048036038101906101f5919061129f565b610493565b60405161020791906112fa565b60405180910390f35b6102186104ca565b005b610234600480360381019061022f91906113c9565b6104dc565b005b61023e6104f0565b60405161024b91906112fa565b60405180910390f35b61026e600480360381019061026991906113f6565b610507565b60405161027b9190611324565b60405180910390f35b61028c61054f565b005b6102a860048036038101906102a3919061129f565b610563565b005b6102b2610583565b005b6102bc610595565b6040516102c99190611432565b60405180910390f35b6102da6105bf565b6040516102e791906111e4565b60405180910390f35b61030a6004803603810190610305919061129f565b610651565b60405161031791906112fa565b60405180910390f35b61033a6004803603810190610335919061129f565b6106c8565b60405161034791906112fa565b60405180910390f35b61036a6004803603810190610365919061144d565b6106eb565b6040516103779190611324565b60405180910390f35b61039a600480360381019061039591906113f6565b610772565b005b6060600380546103ab906114bc565b80601f01602080910402602001604051908101604052809291908181526020018280546103d7906114bc565b80156104245780601f106103f957610100808354040283529160200191610424565b820191906000526020600020905b81548152906001019060200180831161040757829003601f168201915b5050505050905090565b6000806104396107fa565b9050610446818585610802565b600191505092915050565b6000600254905090565b6000806104666107fa565b90506104738582856109cb565b61047e858585610a57565b60019150509392505050565b60006012905090565b60008061049e6107fa565b90506104bf8185856104b085896106eb565b6104ba919061151c565b610802565b600191505092915050565b6104d2610ccd565b6104da610d4b565b565b6104ed6104e76107fa565b82610dae565b50565b6000600560149054906101000a900460ff16905090565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b610557610ccd565b6105616000610f7b565b565b6105758261056f6107fa565b836109cb565b61057f8282610dae565b5050565b61058b610ccd565b610593611041565b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600480546105ce906114bc565b80601f01602080910402602001604051908101604052809291908181526020018280546105fa906114bc565b80156106475780601f1061061c57610100808354040283529160200191610647565b820191906000526020600020905b81548152906001019060200180831161062a57829003601f168201915b5050505050905090565b60008061065c6107fa565b9050600061066a82866106eb565b9050838110156106af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a6906115c2565b60405180910390fd5b6106bc8286868403610802565b60019250505092915050565b6000806106d36107fa565b90506106e0818585610a57565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b61077a610ccd565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036107e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e090611654565b60405180910390fd5b6107f281610f7b565b50565b505050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610871576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610868906116e6565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d790611778565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516109be9190611324565b60405180910390a3505050565b60006109d784846106eb565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610a515781811015610a43576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3a906117e4565b60405180910390fd5b610a508484848403610802565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610ac6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610abd90611876565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610b35576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b2c90611908565b60405180910390fd5b610b408383836110a4565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610bc6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bbd9061199a565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610cb49190611324565b60405180910390a3610cc78484846110bc565b50505050565b610cd56107fa565b73ffffffffffffffffffffffffffffffffffffffff16610cf3610595565b73ffffffffffffffffffffffffffffffffffffffff1614610d49576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d4090611a06565b60405180910390fd5b565b610d536110c1565b6000600560146101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa610d976107fa565b604051610da49190611432565b60405180910390a1565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610e1d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1490611a98565b60405180910390fd5b610e29826000836110a4565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610eaf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ea690611b2a565b60405180910390fd5b8181036000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282540392505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610f629190611324565b60405180910390a3610f76836000846110bc565b505050565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61104961110a565b6001600560146101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861108d6107fa565b60405161109a9190611432565b60405180910390a1565b6110ac61110a565b6110b78383836107f5565b505050565b505050565b6110c96104f0565b611108576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110ff90611b96565b60405180910390fd5b565b6111126104f0565b15611152576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161114990611c02565b60405180910390fd5b565b600081519050919050565b600082825260208201905092915050565b60005b8381101561118e578082015181840152602081019050611173565b60008484015250505050565b6000601f19601f8301169050919050565b60006111b682611154565b6111c0818561115f565b93506111d0818560208601611170565b6111d98161119a565b840191505092915050565b600060208201905081810360008301526111fe81846111ab565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006112368261120b565b9050919050565b6112468161122b565b811461125157600080fd5b50565b6000813590506112638161123d565b92915050565b6000819050919050565b61127c81611269565b811461128757600080fd5b50565b60008135905061129981611273565b92915050565b600080604083850312156112b6576112b5611206565b5b60006112c485828601611254565b92505060206112d58582860161128a565b9150509250929050565b60008115159050919050565b6112f4816112df565b82525050565b600060208201905061130f60008301846112eb565b92915050565b61131e81611269565b82525050565b60006020820190506113396000830184611315565b92915050565b60008060006060848603121561135857611357611206565b5b600061136686828701611254565b935050602061137786828701611254565b92505060406113888682870161128a565b9150509250925092565b600060ff82169050919050565b6113a881611392565b82525050565b60006020820190506113c3600083018461139f565b92915050565b6000602082840312156113df576113de611206565b5b60006113ed8482850161128a565b91505092915050565b60006020828403121561140c5761140b611206565b5b600061141a84828501611254565b91505092915050565b61142c8161122b565b82525050565b60006020820190506114476000830184611423565b92915050565b6000806040838503121561146457611463611206565b5b600061147285828601611254565b925050602061148385828601611254565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806114d457607f821691505b6020821081036114e7576114e661148d565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061152782611269565b915061153283611269565b925082820190508082111561154a576115496114ed565b5b92915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b60006115ac60258361115f565b91506115b782611550565b604082019050919050565b600060208201905081810360008301526115db8161159f565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b600061163e60268361115f565b9150611649826115e2565b604082019050919050565b6000602082019050818103600083015261166d81611631565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b60006116d060248361115f565b91506116db82611674565b604082019050919050565b600060208201905081810360008301526116ff816116c3565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b600061176260228361115f565b915061176d82611706565b604082019050919050565b6000602082019050818103600083015261179181611755565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b60006117ce601d8361115f565b91506117d982611798565b602082019050919050565b600060208201905081810360008301526117fd816117c1565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b600061186060258361115f565b915061186b82611804565b604082019050919050565b6000602082019050818103600083015261188f81611853565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b60006118f260238361115f565b91506118fd82611896565b604082019050919050565b60006020820190508181036000830152611921816118e5565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b600061198460268361115f565b915061198f82611928565b604082019050919050565b600060208201905081810360008301526119b381611977565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006119f060208361115f565b91506119fb826119ba565b602082019050919050565b60006020820190508181036000830152611a1f816119e3565b9050919050565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b6000611a8260218361115f565b9150611a8d82611a26565b604082019050919050565b60006020820190508181036000830152611ab181611a75565b9050919050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b6000611b1460228361115f565b9150611b1f82611ab8565b604082019050919050565b60006020820190508181036000830152611b4381611b07565b9050919050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b6000611b8060148361115f565b9150611b8b82611b4a565b602082019050919050565b60006020820190508181036000830152611baf81611b73565b9050919050565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b6000611bec60108361115f565b9150611bf782611bb6565b602082019050919050565b60006020820190508181036000830152611c1b81611bdf565b905091905056fea2646970667358221220d2b3c9fd6251165e3d79329e5589a5a84db0a6e4fccd455aa46b68f50a9cd09d64736f6c63430008100033",
}

// BlockhainABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockhainMetaData.ABI instead.
var BlockhainABI = BlockhainMetaData.ABI

// BlockhainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BlockhainMetaData.Bin instead.
var BlockhainBin = BlockhainMetaData.Bin

// DeployBlockhain deploys a new Ethereum contract, binding an instance of Blockhain to it.
func DeployBlockhain(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, _initialWallet common.Address, _supply *big.Int) (common.Address, *types.Transaction, *Blockhain, error) {
	parsed, err := BlockhainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlockhainBin), backend, name, symbol, _initialWallet, _supply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Blockhain{BlockhainCaller: BlockhainCaller{contract: contract}, BlockhainTransactor: BlockhainTransactor{contract: contract}, BlockhainFilterer: BlockhainFilterer{contract: contract}}, nil
}

// Blockhain is an auto generated Go binding around an Ethereum contract.
type Blockhain struct {
	BlockhainCaller     // Read-only binding to the contract
	BlockhainTransactor // Write-only binding to the contract
	BlockhainFilterer   // Log filterer for contract events
}

// BlockhainCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockhainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockhainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockhainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockhainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockhainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockhainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockhainSession struct {
	Contract     *Blockhain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockhainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockhainCallerSession struct {
	Contract *BlockhainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BlockhainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockhainTransactorSession struct {
	Contract     *BlockhainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BlockhainRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockhainRaw struct {
	Contract *Blockhain // Generic contract binding to access the raw methods on
}

// BlockhainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockhainCallerRaw struct {
	Contract *BlockhainCaller // Generic read-only contract binding to access the raw methods on
}

// BlockhainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockhainTransactorRaw struct {
	Contract *BlockhainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockhain creates a new instance of Blockhain, bound to a specific deployed contract.
func NewBlockhain(address common.Address, backend bind.ContractBackend) (*Blockhain, error) {
	contract, err := bindBlockhain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blockhain{BlockhainCaller: BlockhainCaller{contract: contract}, BlockhainTransactor: BlockhainTransactor{contract: contract}, BlockhainFilterer: BlockhainFilterer{contract: contract}}, nil
}

// NewBlockhainCaller creates a new read-only instance of Blockhain, bound to a specific deployed contract.
func NewBlockhainCaller(address common.Address, caller bind.ContractCaller) (*BlockhainCaller, error) {
	contract, err := bindBlockhain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockhainCaller{contract: contract}, nil
}

// NewBlockhainTransactor creates a new write-only instance of Blockhain, bound to a specific deployed contract.
func NewBlockhainTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockhainTransactor, error) {
	contract, err := bindBlockhain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockhainTransactor{contract: contract}, nil
}

// NewBlockhainFilterer creates a new log filterer instance of Blockhain, bound to a specific deployed contract.
func NewBlockhainFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockhainFilterer, error) {
	contract, err := bindBlockhain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockhainFilterer{contract: contract}, nil
}

// bindBlockhain binds a generic wrapper to an already deployed contract.
func bindBlockhain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockhainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockhain *BlockhainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockhain.Contract.BlockhainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockhain *BlockhainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockhain.Contract.BlockhainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockhain *BlockhainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockhain.Contract.BlockhainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockhain *BlockhainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockhain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockhain *BlockhainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockhain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockhain *BlockhainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockhain.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Blockhain *BlockhainCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Blockhain.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Blockhain *BlockhainSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Blockhain.Contract.Allowance(&_Blockhain.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Blockhain *BlockhainCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Blockhain.Contract.Allowance(&_Blockhain.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Blockhain *BlockhainCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Blockhain.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Blockhain *BlockhainSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Blockhain.Contract.BalanceOf(&_Blockhain.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Blockhain *BlockhainCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Blockhain.Contract.BalanceOf(&_Blockhain.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Blockhain *BlockhainCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Blockhain.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Blockhain *BlockhainSession) Decimals() (uint8, error) {
	return _Blockhain.Contract.Decimals(&_Blockhain.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Blockhain *BlockhainCallerSession) Decimals() (uint8, error) {
	return _Blockhain.Contract.Decimals(&_Blockhain.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Blockhain *BlockhainCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Blockhain.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Blockhain *BlockhainSession) Name() (string, error) {
	return _Blockhain.Contract.Name(&_Blockhain.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Blockhain *BlockhainCallerSession) Name() (string, error) {
	return _Blockhain.Contract.Name(&_Blockhain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blockhain *BlockhainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blockhain.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blockhain *BlockhainSession) Owner() (common.Address, error) {
	return _Blockhain.Contract.Owner(&_Blockhain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blockhain *BlockhainCallerSession) Owner() (common.Address, error) {
	return _Blockhain.Contract.Owner(&_Blockhain.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Blockhain *BlockhainCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Blockhain.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Blockhain *BlockhainSession) Paused() (bool, error) {
	return _Blockhain.Contract.Paused(&_Blockhain.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Blockhain *BlockhainCallerSession) Paused() (bool, error) {
	return _Blockhain.Contract.Paused(&_Blockhain.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Blockhain *BlockhainCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Blockhain.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Blockhain *BlockhainSession) Symbol() (string, error) {
	return _Blockhain.Contract.Symbol(&_Blockhain.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Blockhain *BlockhainCallerSession) Symbol() (string, error) {
	return _Blockhain.Contract.Symbol(&_Blockhain.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Blockhain *BlockhainCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blockhain.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Blockhain *BlockhainSession) TotalSupply() (*big.Int, error) {
	return _Blockhain.Contract.TotalSupply(&_Blockhain.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Blockhain *BlockhainCallerSession) TotalSupply() (*big.Int, error) {
	return _Blockhain.Contract.TotalSupply(&_Blockhain.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Blockhain *BlockhainTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Blockhain.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Blockhain *BlockhainSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Blockhain.Contract.Approve(&_Blockhain.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Blockhain *BlockhainTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Blockhain.Contract.Approve(&_Blockhain.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Blockhain *BlockhainTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Blockhain.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Blockhain *BlockhainSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Blockhain.Contract.Burn(&_Blockhain.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Blockhain *BlockhainTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Blockhain.Contract.Burn(&_Blockhain.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Blockhain *BlockhainTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Blockhain.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Blockhain *BlockhainSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Blockhain.Contract.BurnFrom(&_Blockhain.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Blockhain *BlockhainTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Blockhain.Contract.BurnFrom(&_Blockhain.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Blockhain *BlockhainTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Blockhain.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Blockhain *BlockhainSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Blockhain.Contract.DecreaseAllowance(&_Blockhain.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Blockhain *BlockhainTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Blockhain.Contract.DecreaseAllowance(&_Blockhain.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Blockhain *BlockhainTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Blockhain.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Blockhain *BlockhainSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Blockhain.Contract.IncreaseAllowance(&_Blockhain.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Blockhain *BlockhainTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Blockhain.Contract.IncreaseAllowance(&_Blockhain.TransactOpts, spender, addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Blockhain *BlockhainTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockhain.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Blockhain *BlockhainSession) Pause() (*types.Transaction, error) {
	return _Blockhain.Contract.Pause(&_Blockhain.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Blockhain *BlockhainTransactorSession) Pause() (*types.Transaction, error) {
	return _Blockhain.Contract.Pause(&_Blockhain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blockhain *BlockhainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockhain.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blockhain *BlockhainSession) RenounceOwnership() (*types.Transaction, error) {
	return _Blockhain.Contract.RenounceOwnership(&_Blockhain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blockhain *BlockhainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Blockhain.Contract.RenounceOwnership(&_Blockhain.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Blockhain *BlockhainTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Blockhain.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Blockhain *BlockhainSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Blockhain.Contract.Transfer(&_Blockhain.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Blockhain *BlockhainTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Blockhain.Contract.Transfer(&_Blockhain.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Blockhain *BlockhainTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Blockhain.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Blockhain *BlockhainSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Blockhain.Contract.TransferFrom(&_Blockhain.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Blockhain *BlockhainTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Blockhain.Contract.TransferFrom(&_Blockhain.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blockhain *BlockhainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Blockhain.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blockhain *BlockhainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Blockhain.Contract.TransferOwnership(&_Blockhain.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blockhain *BlockhainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Blockhain.Contract.TransferOwnership(&_Blockhain.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Blockhain *BlockhainTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockhain.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Blockhain *BlockhainSession) Unpause() (*types.Transaction, error) {
	return _Blockhain.Contract.Unpause(&_Blockhain.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Blockhain *BlockhainTransactorSession) Unpause() (*types.Transaction, error) {
	return _Blockhain.Contract.Unpause(&_Blockhain.TransactOpts)
}

// BlockhainApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Blockhain contract.
type BlockhainApprovalIterator struct {
	Event *BlockhainApproval // Event containing the contract specifics and raw log

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
func (it *BlockhainApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockhainApproval)
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
		it.Event = new(BlockhainApproval)
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
func (it *BlockhainApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockhainApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockhainApproval represents a Approval event raised by the Blockhain contract.
type BlockhainApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Blockhain *BlockhainFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BlockhainApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Blockhain.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BlockhainApprovalIterator{contract: _Blockhain.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Blockhain *BlockhainFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BlockhainApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Blockhain.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockhainApproval)
				if err := _Blockhain.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Blockhain *BlockhainFilterer) ParseApproval(log types.Log) (*BlockhainApproval, error) {
	event := new(BlockhainApproval)
	if err := _Blockhain.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockhainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Blockhain contract.
type BlockhainOwnershipTransferredIterator struct {
	Event *BlockhainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BlockhainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockhainOwnershipTransferred)
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
		it.Event = new(BlockhainOwnershipTransferred)
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
func (it *BlockhainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockhainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockhainOwnershipTransferred represents a OwnershipTransferred event raised by the Blockhain contract.
type BlockhainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Blockhain *BlockhainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlockhainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Blockhain.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlockhainOwnershipTransferredIterator{contract: _Blockhain.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Blockhain *BlockhainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlockhainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Blockhain.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockhainOwnershipTransferred)
				if err := _Blockhain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Blockhain *BlockhainFilterer) ParseOwnershipTransferred(log types.Log) (*BlockhainOwnershipTransferred, error) {
	event := new(BlockhainOwnershipTransferred)
	if err := _Blockhain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockhainPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Blockhain contract.
type BlockhainPausedIterator struct {
	Event *BlockhainPaused // Event containing the contract specifics and raw log

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
func (it *BlockhainPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockhainPaused)
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
		it.Event = new(BlockhainPaused)
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
func (it *BlockhainPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockhainPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockhainPaused represents a Paused event raised by the Blockhain contract.
type BlockhainPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Blockhain *BlockhainFilterer) FilterPaused(opts *bind.FilterOpts) (*BlockhainPausedIterator, error) {

	logs, sub, err := _Blockhain.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BlockhainPausedIterator{contract: _Blockhain.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Blockhain *BlockhainFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BlockhainPaused) (event.Subscription, error) {

	logs, sub, err := _Blockhain.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockhainPaused)
				if err := _Blockhain.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Blockhain *BlockhainFilterer) ParsePaused(log types.Log) (*BlockhainPaused, error) {
	event := new(BlockhainPaused)
	if err := _Blockhain.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockhainTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Blockhain contract.
type BlockhainTransferIterator struct {
	Event *BlockhainTransfer // Event containing the contract specifics and raw log

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
func (it *BlockhainTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockhainTransfer)
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
		it.Event = new(BlockhainTransfer)
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
func (it *BlockhainTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockhainTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockhainTransfer represents a Transfer event raised by the Blockhain contract.
type BlockhainTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Blockhain *BlockhainFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BlockhainTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Blockhain.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BlockhainTransferIterator{contract: _Blockhain.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Blockhain *BlockhainFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BlockhainTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Blockhain.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockhainTransfer)
				if err := _Blockhain.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Blockhain *BlockhainFilterer) ParseTransfer(log types.Log) (*BlockhainTransfer, error) {
	event := new(BlockhainTransfer)
	if err := _Blockhain.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockhainUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Blockhain contract.
type BlockhainUnpausedIterator struct {
	Event *BlockhainUnpaused // Event containing the contract specifics and raw log

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
func (it *BlockhainUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockhainUnpaused)
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
		it.Event = new(BlockhainUnpaused)
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
func (it *BlockhainUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockhainUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockhainUnpaused represents a Unpaused event raised by the Blockhain contract.
type BlockhainUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Blockhain *BlockhainFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BlockhainUnpausedIterator, error) {

	logs, sub, err := _Blockhain.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BlockhainUnpausedIterator{contract: _Blockhain.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Blockhain *BlockhainFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BlockhainUnpaused) (event.Subscription, error) {

	logs, sub, err := _Blockhain.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockhainUnpaused)
				if err := _Blockhain.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Blockhain *BlockhainFilterer) ParseUnpaused(log types.Log) (*BlockhainUnpaused, error) {
	event := new(BlockhainUnpaused)
	if err := _Blockhain.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
