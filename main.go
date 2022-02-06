package main

import (
	"fmt"

	"github.com/eslerkang/tjcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockChain()
	fmt.Println(chain)
}