package main

import (
	"fmt"

	"github.com/daveg7lee/kangaroocoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	blocks := chain.AllBlocks()
	for _, block := range blocks {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Prev Hash: %s\n", block.PrevHash)
	}
}
