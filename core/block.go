package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)
type Block struct {
	Index int64 //  编号
	Timestamp int64 // 时间戳
	Prehash string
	Currhash string
	Data string
}

// 生成sha256的字符串
func generateHash(b Block) string {
	// 生成hash值不能用当前hash来进行计算
	blockData := string(b.Index) + string(b.Timestamp) + b.Prehash + b.Data
	// func Sum256(data []byte) [Size]byte
	sha := sha256.Sum256([]byte(blockData))
	// func EncodeToString(src []byte) string
	return hex.EncodeToString(sha[:])
}

// 生成区块
func GenerateBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Prehash = preBlock.Currhash
	newBlock.Data = data
	newBlock.Currhash = generateHash(newBlock)
	return newBlock
}

// 生成创世区块
func GenerateGenesisBlock() Block {
	newBlock := Block{}
	newBlock.Index = -1
	newBlock.Prehash = ""
	return GenerateBlock(newBlock, "this is GenesisBlock")
}