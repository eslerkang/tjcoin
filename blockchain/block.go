package blockchain

import (
	"crypto/sha256"
	"fmt"

	"github.com/eslerkang/tjcoin/db"
	"github.com/eslerkang/tjcoin/utils"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

func (b *Block) persist() {
	db.SaveInBucket(db.BlocksBucket, b.Hash, utils.ToBytes(b))
}

func createBlock(data string) *Block {
	block := Block{
		Data:     data,
		Hash:     "",
		PrevHash: BlockChain().NewestHash,
		Height:   BlockChain().Height + 1,
	}

	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	block.persist()
	return &block
}
