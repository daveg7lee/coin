package blockchain

import (
	"sync"

	"github.com/daveg7lee/kangaroocoin/db"
	"github.com/daveg7lee/kangaroocoin/utils"
)

const (
	defaultDifficulty  int = 2
	difficultyInterval int = 5
)

// blockchain struct
type blockchain struct {
	NewestHash        string `json:"newestHash"`
	Height            int    `json:"height"`
	CurrentDifficulty int    `json:"currentDifficulty"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
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

func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		// calculate difficulty
	} else {
		return b.CurrentDifficulty
	}
}

func Blockchain() *blockchain {
	// check blockchain is not nil
	if b == nil {
		// create blockchain but it occur only once
		once.Do(func() {
			b = &blockchain{
				Height: 0,
			}
			checkPoint := db.CheckPoint()
			if checkPoint == nil {
				// add genesis block
				b.AddBlock("Genesis Block!!")
			} else {
				b.restore(checkPoint)
			}
		})
	}
	// return blockchain
	return b
}
