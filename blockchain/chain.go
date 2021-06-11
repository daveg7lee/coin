package blockchain

import (
	"sync"
)

// blockchain struct
type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height)
	b.NewestHash = block.Hash
	b.Height = block.Height
}

func Blockchain() *blockchain {
	// check blockchain is not nil
	if b == nil {
		// create blockchain but it occur only once
		once.Do(func() { b = &blockchain{"", 0} })
		// add genesis block
		b.AddBlock("Genesis Block!!")
	}
	// return blockchain
	return b
}
