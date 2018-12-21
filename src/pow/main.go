package main

import (
	"fmt"
)

func main() {

	// 创世区块
	blockchain := CreateBlockchainWithGenesisBlock()

	// 新区块
	blockchain.AddBlockToBlockchain("Send 100RMB To tom", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	blockchain.AddBlockToBlockchain("Send 200RMB To lily", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	blockchain.AddBlockToBlockchain("Send 300RMB To hanmeimei", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	blockchain.AddBlockToBlockchain("Send 50RMB To lucy", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	fmt.Println(blockchain)
	fmt.Println(blockchain.Blocks)
}
