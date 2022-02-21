package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
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

func (b *blockchain) restore(data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	utils.HandleError(decoder.Decode(b))
}

func (b *blockchain) persist() {
	db.SaveInBucket(db.DATA_BUCKET, db.CHECKPOINT, utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data)
	b.NewestHash = block.Hash
	b.Height = block.Height + 1
	b.persist()
}

func BlockChain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			fmt.Printf("NewestHash: %s\nHeight: %d", b.NewestHash, b.Height)
			// search for checkpoint on the db
			checkpoint := db.CheckPoint()
			if checkpoint == nil {
				b.AddBlock("Genesis Block")
			} else {
				// restore b from bytes
				fmt.Println("Restoring")
				b.restore(checkpoint)

				fmt.Printf("NewestHash: %s\nHeight: %d", b.NewestHash, b.Height)
			}
		})
	}
	return b
}
