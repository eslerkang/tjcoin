package blockchain

import (
	"sync"

	"github.com/eslerkang/tjcoin/db"
	"github.com/eslerkang/tjcoin/utils"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) persist() {
	db.SaveInBucket(db.Databucket, "checkpoint", utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func BlockChain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}