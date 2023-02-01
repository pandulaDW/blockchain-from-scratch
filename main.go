package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	blocks []Block
}

func InitBlockChain() BlockChain {
	return BlockChain{blocks: []Block{Genesis()}}
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := &chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

type Block struct {
	Data     []byte
	Hash     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	// combine the data and prev-hash
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})

	// compute the hash
	hash := sha256.Sum256(info)

	// set the hash on the block
	b.Hash = hash[:]
}

func Genesis() Block {
	return CreateBlock("Genesis", []byte{})
}

func CreateBlock(data string, prevHash []byte) Block {
	b := Block{Data: []byte(data), PrevHash: prevHash}
	b.DeriveHash()
	return b
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First block after Genesis")
	chain.AddBlock("second block after Genesis")
	chain.AddBlock("Third block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
