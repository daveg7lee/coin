package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

// block struct
type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
}

// blockchain struct
type blockchain struct {
	blocks []*Block
}

var b *blockchain
var once sync.Once

func (b *Block) calculateHash() {
	// calculate hash of block's data + prev block's data
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	// change fomat of hash to '%x'
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	// get length of blockchain
	totalBlocks := len(GetBlockchain().blocks)
	// return nothing when block is genesis block
	if totalBlocks == 0 {
		return ""
	}
	// return last block's hash
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

func createBlock(data string) *Block {
	// make new block
	newBlock := Block{data, "", getLastHash()}
	// calculate Hash
	newBlock.calculateHash()
	// return block
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	// append block to blockchain
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	// check blockchain is not nil
	if b == nil {
		// create blockchain but it occur only once
		once.Do(func() { b = &blockchain{} })
		// add genesis block
		b.AddBlock("Genesis Block!!")
	}
	// return blockchain
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	// return all blocks
	return b.blocks
}
