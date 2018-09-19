package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index        int64
	Timestamp    int64
	Preblockhash string
	Hash         string
	Data         string
}

func GenerateHash(newBlock Block) string {
	valueStr := string(newBlock.Index) + string(newBlock.Timestamp) + newBlock.Preblockhash + newBlock.Data
	hashInByte := sha256.Sum256([]byte(valueStr))
	hashInStr := hex.EncodeToString(hashInByte[:])
	//hashInStr := fmt.Sprintf("%x", hashInByte)
	return hashInStr
}

func GenerateNewBlock(Preblock Block, Data string) Block {
	newBlock := Block{}
	newBlock.Index = Preblock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Preblockhash = Preblock.Hash
	newBlock.Data = Data
	newBlock.Hash = GenerateHash(newBlock)
	return newBlock
}
func GenerateSisBlock() Block {
	Preblock := Block{}
	Preblock.Index = -1
	Preblock.Hash = ""
	Data := "this is first block"
	genersisblock := GenerateNewBlock(Preblock, Data)
	return genersisblock
}
