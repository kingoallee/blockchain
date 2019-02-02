package core
import (
	"log"
	"fmt"
)

type Chain struct {
	Blocks []*Block
}

func NewBlock() *Chain {
	genesiBlock := GenerateGenesisBlock()
	c := Chain{}
	c.AddBlock(&genesiBlock)
	return &c
}

// GenerateBlock(preBlock Block, data string)
func (c *Chain) SendData(data string) {
	preBlock := c.Blocks[len(c.Blocks) - 1]
	b := GenerateBlock(*preBlock, data)
	c.AddBlock(&b)
}

func (c *Chain) AddBlock(newBlock *Block) {
	if len(c.Blocks) == 0 {
		c.Blocks = append(c.Blocks, newBlock)
		return
	}
	if validate(*newBlock, *c.Blocks[len(c.Blocks) -1 ]) {
		c.Blocks = append(c.Blocks, newBlock)
	} else {
		log.Fatal("无效的区块")
	}
}

func validate(newBlock Block, oldBlock Block) bool {
	if newBlock.Index - 1 != oldBlock.Index {
		return false
	}

	if newBlock.Prehash != oldBlock.Currhash {
		return false
	}

	if generateHash(newBlock) != newBlock.Currhash {
		return false
	}

	return true
}

func (c *Chain) Print() {
	for _, block := range c.Blocks {
		fmt.Printf("Index= %d\n", block.Index)
		fmt.Printf("PreHash= %s\n", block.Prehash)
		fmt.Printf("CurrHash= %s\n", block.Currhash)
		fmt.Printf("Data= %s\n", block.Data)
		fmt.Printf("Timestamp= %d\n", block.Timestamp)
		fmt.Println()
	}
}