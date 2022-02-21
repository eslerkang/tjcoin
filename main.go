package main

import (
	"github.com/eslerkang/tjcoin/blockchain"
)


func main() {
	// cli.Start()
	blockchain.BlockChain().AddBlock("First")
	blockchain.BlockChain().AddBlock("Second")
	blockchain.BlockChain().AddBlock("Third")
}