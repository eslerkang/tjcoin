package main

import (
	"fmt"

	"github.com/eslerkang/tjcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockChain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")
	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data: %s\nHash: %s\nPrevHash: %s\n\n", block.Data, block.Hash, block.PrevHash)
	}
}