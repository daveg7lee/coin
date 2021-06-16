package blockchain

import (
	"sync"

	"github.com/daveg7lee/kangaroocoin/db"
	"github.com/daveg7lee/kangaroocoin/utils"
)

const (
	defaultDifficulty  int = 3
	difficultyInterval int = 5
	blockInterval      int = 2
	allowedRange       int = 2
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

func (b *blockchain) AddBlock() {
	block := createBlock(b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
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

func (b *blockchain) recalculateDifficulty() int {
	allBlocks := b.Blocks()
	newestBlock := allBlocks[0]
	recalculatedBlock := allBlocks[difficultyInterval-1]
	actualTime := (newestBlock.Timestamp / 60) - (recalculatedBlock.Timestamp / 60)
	expectedTime := difficultyInterval * blockInterval
	if actualTime <= (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime >= (expectedTime + allowedRange) {
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

func (b *blockchain) txOuts() []*TxOut {
	var txOuts []*TxOut
	blocks := b.Blocks()
	for _, block := range blocks {
		for _, tx := range block.Transactions {
			txOuts = append(txOuts, tx.TxOuts...)
		}
	}
	return txOuts
}

func (b *blockchain) TxOutsByAddress(address string) []*TxOut {
	var ownedTxOuts []*TxOut
	txOuts := b.txOuts()
	for _, txOut := range txOuts {
		if txOut.Owner == address {
			ownedTxOuts = append(ownedTxOuts, txOut)
		}
	}
	return ownedTxOuts
}

func (b *blockchain) BalanceByAddress(address string) int {
	txOuts := b.TxOutsByAddress(address)
	var amount int
	for _, txOut := range txOuts {
		amount += txOut.Amount
	}
	return amount
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
				b.AddBlock()
			} else {
				b.restore(checkPoint)
			}
		})
	}
	// return blockchain
	return b
}
