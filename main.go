package main

import (
	"github.com/eslerkang/tjcoin/blockchain"
	"github.com/eslerkang/tjcoin/cli"
	"github.com/eslerkang/tjcoin/db"
)

func main() {
	defer db.Close()
	blockchain.BlockChain()
	cli.Start()
}
