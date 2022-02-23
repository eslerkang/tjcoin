package blockchain

import (
	"fmt"
	"sync"

	"github.com/eslerkang/tjcoin/db"
	"github.com/eslerkang/tjcoin/utils"
)

type blockchain struct {
	NewestHash        string `json:"newestHash"`
	Height            int    `json:"height"`
	CurrentDifficulty int    `json:"currentdifficulty"`
}

var b *blockchain
var once sync.Once

const (
	defaultDifficulty      int = 2
	difficultyInterval     int = 5
	blockCreationInterval  int = 2
	timeDifferAllowedRange int = 2
)

func (b *blockchain) persist() {
	db.SaveInBucket(db.DATA_BUCKET, db.CHECKPOINT, utils.ToBytes(b))
}

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	b.persist()
}

func (b *blockchain) Blocks() []*Block {
	hashCursor := b.NewestHash
	var blocks []*Block
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func (b *blockchain) recalculateDifficulty() int {
	allBlocks := b.Blocks()
	newestBlock := allBlocks[0]
	lastRecalculatedIntervalBlock := allBlocks[difficultyInterval-1]
	actualTimeDiffer := (lastRecalculatedIntervalBlock.TimeStamp / 60) - (newestBlock.TimeStamp / 60)
	expectedTimeDiffer := difficultyInterval * blockCreationInterval

	if actualTimeDiffer < expectedTimeDiffer+timeDifferAllowedRange {
		return b.CurrentDifficulty + 1
	} else if actualTimeDiffer > expectedTimeDiffer-timeDifferAllowedRange {
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty
}

func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		return b.recalculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}

func BlockChain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{
				Height: 0,
			}
			// search for checkpoint on the db
			checkpoint := db.CheckPoint()
			if checkpoint == nil {
				b.AddBlock("Genesis Block")
			} else {
				// restore b from bytes
				b.restore(checkpoint)
			}
		})
	}
	fmt.Println(b.NewestHash)
	return b
}
