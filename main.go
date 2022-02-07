package main

import (
	"github.com/eslerkang/tjcoin/explorer"
	"github.com/eslerkang/tjcoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}