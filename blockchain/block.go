package blockchain

type BlockChain struct {
	Blocks []Block
}

// InitBlockChain initializes a new blockchain
func InitBlockChain() BlockChain {
	return BlockChain{Blocks: []Block{Genesis()}}
}

// AddBlock adds a new block to the chain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := &chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

type Block struct {
	Data     []byte
	Hash     []byte
	PrevHash []byte
	Nonce    int
}

// Genesis creates the new genesis block
func Genesis() Block {
	return CreateBlock("Genesis", []byte{})
}

// CreateBlock creates a new block and returns it
func CreateBlock(data string, prevHash []byte) Block {
	block := Block{Data: []byte(data), PrevHash: prevHash}

	// run the proof of work for the block
	pow := NewProof(&block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash

	return block
}
