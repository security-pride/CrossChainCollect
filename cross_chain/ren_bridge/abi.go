package renbridge

const (
	renAbiStr = `
	[
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": false,
					"internalType": "bytes",
					"name": "_to",
					"type": "bytes"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "_amount",
					"type": "uint256"
				},
				{
					"indexed": true,
					"internalType": "uint256",
					"name": "_n",
					"type": "uint256"
				},
				{
					"indexed": true,
					"internalType": "bytes",
					"name": "_indexedTo",
					"type": "bytes"
				}
			],
			"name": "LogBurn",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "_to",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "_amount",
					"type": "uint256"
				},
				{
					"indexed": true,
					"internalType": "uint256",
					"name": "_n",
					"type": "uint256"
				},
				{
					"indexed": true,
					"internalType": "bytes32",
					"name": "_nHash",
					"type": "bytes32"
				}
			],
			"name": "LogMint",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "_newMintAuthority",
					"type": "address"
				}
			],
			"name": "LogMintAuthorityUpdated",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "previousOwner",
					"type": "address"
				},
				{
					"indexed": true,
					"internalType": "address",
					"name": "newOwner",
					"type": "address"
				}
			],
			"name": "OwnershipTransferred",
			"type": "event"
		},
		{
			"constant": false,
			"inputs": [
				{
					"internalType": "address",
					"name": "newOwner",
					"type": "address"
				}
			],
			"name": "_directTransferOwnership",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [
				{
					"internalType": "bytes32",
					"name": "_pHash",
					"type": "bytes32"
				},
				{
					"internalType": "uint256",
					"name": "_amount",
					"type": "uint256"
				},
				{
					"internalType": "address",
					"name": "_to",
					"type": "address"
				},
				{
					"internalType": "bytes32",
					"name": "_nHash",
					"type": "bytes32"
				}
			],
			"name": "_legacy_hashForSignature",
			"outputs": [
				{
					"internalType": "bytes32",
					"name": "",
					"type": "bytes32"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"internalType": "address",
					"name": "_token",
					"type": "address"
				}
			],
			"name": "blacklistRecoverableToken",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"internalType": "bytes",
					"name": "_to",
					"type": "bytes"
				},
				{
					"internalType": "uint256",
					"name": "_amount",
					"type": "uint256"
				}
			],
			"name": "burn",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "",
					"type": "uint256"
				}
			],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [],
			"name": "burnFee",
			"outputs": [
				{
					"internalType": "uint16",
					"name": "",
					"type": "uint16"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [],
			"name": "claimOwnership",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [],
			"name": "claimTokenOwnership",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [],
			"name": "feeRecipient",
			"outputs": [
				{
					"internalType": "address",
					"name": "",
					"type": "address"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [
				{
					"internalType": "uint256",
					"name": "_n",
					"type": "uint256"
				}
			],
			"name": "getBurn",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "_blocknumber",
					"type": "uint256"
				},
				{
					"internalType": "bytes",
					"name": "_to",
					"type": "bytes"
				},
				{
					"internalType": "uint256",
					"name": "_amount",
					"type": "uint256"
				},
				{
					"internalType": "string",
					"name": "_chain",
					"type": "string"
				},
				{
					"internalType": "bytes",
					"name": "_payload",
					"type": "bytes"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [
				{
					"internalType": "bytes32",
					"name": "_pHash",
					"type": "bytes32"
				},
				{
					"internalType": "uint256",
					"name": "_amount",
					"type": "uint256"
				},
				{
					"internalType": "address",
					"name": "_to",
					"type": "address"
				},
				{
					"internalType": "bytes32",
					"name": "_nHash",
					"type": "bytes32"
				}
			],
			"name": "hashForSignature",
			"outputs": [
				{
					"internalType": "bytes32",
					"name": "",
					"type": "bytes32"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"internalType": "contract RenERC20LogicV1",
					"name": "_token",
					"type": "address"
				},
				{
					"internalType": "address",
					"name": "_feeRecipient",
					"type": "address"
				},
				{
					"internalType": "address",
					"name": "_mintAuthority",
					"type": "address"
				},
				{
					"internalType": "uint16",
					"name": "_mintFee",
					"type": "uint16"
				},
				{
					"internalType": "uint16",
					"name": "_burnFee",
					"type": "uint16"
				},
				{
					"internalType": "uint256",
					"name": "_minimumBurnAmount",
					"type": "uint256"
				}
			],
			"name": "initialize",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"internalType": "address",
					"name": "_nextOwner",
					"type": "address"
				}
			],
			"name": "initialize",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [],
			"name": "isOwner",
			"outputs": [
				{
					"internalType": "bool",
					"name": "",
					"type": "bool"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [],
			"name": "minimumBurnAmount",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "",
					"type": "uint256"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"internalType": "bytes32",
					"name": "_pHash",
					"type": "bytes32"
				},
				{
					"internalType": "uint256",
					"name": "_amountUnderlying",
					"type": "uint256"
				},
				{
					"internalType": "bytes32",
					"name": "_nHash",
					"type": "bytes32"
				},
				{
					"internalType": "bytes",
					"name": "_sig",
					"type": "bytes"
				}
			],
			"name": "mint",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "",
					"type": "uint256"
				}
			],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [],
			"name": "mintAuthority",
			"outputs": [
				{
					"internalType": "address",
					"name": "",
					"type": "address"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [],
			"name": "mintFee",
			"outputs": [
				{
					"internalType": "uint16",
					"name": "",
					"type": "uint16"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [],
			"name": "nextN",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "",
					"type": "uint256"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [],
			"name": "owner",
			"outputs": [
				{
					"internalType": "address",
					"name": "",
					"type": "address"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [],
			"name": "pendingOwner",
			"outputs": [
				{
					"internalType": "address",
					"name": "",
					"type": "address"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"internalType": "address",
					"name": "_token",
					"type": "address"
				}
			],
			"name": "recoverTokens",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [],
			"name": "renounceOwnership",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [],
			"name": "selectorHash",
			"outputs": [
				{
					"internalType": "bytes32",
					"name": "",
					"type": "bytes32"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [
				{
					"internalType": "bytes32",
					"name": "",
					"type": "bytes32"
				}
			],
			"name": "status",
			"outputs": [
				{
					"internalType": "bool",
					"name": "",
					"type": "bool"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [],
			"name": "token",
			"outputs": [
				{
					"internalType": "contract RenERC20LogicV1",
					"name": "",
					"type": "address"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"internalType": "address",
					"name": "newOwner",
					"type": "address"
				}
			],
			"name": "transferOwnership",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"internalType": "contract MintGatewayLogicV2",
					"name": "_nextTokenOwner",
					"type": "address"
				}
			],
			"name": "transferTokenOwnership",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"internalType": "uint16",
					"name": "_nextBurnFee",
					"type": "uint16"
				}
			],
			"name": "updateBurnFee",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"internalType": "address",
					"name": "_nextFeeRecipient",
					"type": "address"
				}
			],
			"name": "updateFeeRecipient",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"internalType": "uint16",
					"name": "_nextMintFee",
					"type": "uint16"
				},
				{
					"internalType": "uint16",
					"name": "_nextBurnFee",
					"type": "uint16"
				}
			],
			"name": "updateFees",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"internalType": "uint256",
					"name": "_minimumBurnAmount",
					"type": "uint256"
				}
			],
			"name": "updateMinimumBurnAmount",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"internalType": "address",
					"name": "_nextMintAuthority",
					"type": "address"
				}
			],
			"name": "updateMintAuthority",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"internalType": "uint16",
					"name": "_nextMintFee",
					"type": "uint16"
				}
			],
			"name": "updateMintFee",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"internalType": "bytes32",
					"name": "_selectorHash",
					"type": "bytes32"
				}
			],
			"name": "updateSelectorHash",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"internalType": "string",
					"name": "symbol",
					"type": "string"
				}
			],
			"name": "updateSymbol",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [
				{
					"internalType": "bytes32",
					"name": "_sigHash",
					"type": "bytes32"
				},
				{
					"internalType": "bytes",
					"name": "_sig",
					"type": "bytes"
				}
			],
			"name": "verifySignature",
			"outputs": [
				{
					"internalType": "bool",
					"name": "",
					"type": "bool"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		}
	]
	`
)
