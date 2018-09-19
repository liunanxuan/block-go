package core

import (
	"fmt"
	"log"
)

type BlockChain struct {
	Blocks []*Block
}

func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		//fmt.Println("1 %d,%d.", newBlock.Index-1, oldBlock.Index)
		return false
	}
	if newBlock.Preblockhash != oldBlock.Hash {
		//fmt.Println("2 %s,%s.", newBlock.Preblockhash, oldBlock.Hash)
		return false
	}
	if GenerateHash(newBlock) != newBlock.Hash {
		//fmt.Println("3 %s,%s.", GenerateHash(newBlock), newBlock.Hash)
		return false
	}
	return true
}
func (bc *BlockChain) ApendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) == true {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block")
	}

}
func NewBlockchain() *BlockChain {
	genesisBlock := GenerateSisBlock()
	blockchain := BlockChain{}
	blockchain.ApendBlock(&genesisBlock)
	return &blockchain
}
func (bc *BlockChain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.ApendBlock(&newBlock)
}
func (bc *BlockChain) Blockprint() {
	for _, block := range bc.Blocks {
		fmt.Println("Index:", block.Index)
		fmt.Println("Prev.Hash:", block.Preblockhash)
		fmt.Println("Curr.Hash:", block.Hash)
		fmt.Println("Data:", block.Data)
		fmt.Println("timestamp:", block.Timestamp)
		fmt.Println("----------------------------------")

	}
}
