package main

import (
	"assignment01bca"
	"fmt"
)

func main() {
	// Create genesis block
	genesisBlock := assignment01bca.NewBlock("Genesis Block", 0, "")
	assignment01bca.Blockchain = append(assignment01bca.Blockchain, *genesisBlock)

	// Add more blocks
	block1 := assignment01bca.NewBlock("Maheen to Alice", 1, genesisBlock.Hash)
	assignment01bca.Blockchain = append(assignment01bca.Blockchain, *block1)

	block2 := assignment01bca.NewBlock("alice to charlie", 2, block1.Hash)
	assignment01bca.Blockchain = append(assignment01bca.Blockchain, *block2)

	// List all blocks
	assignment01bca.ListBlocks()

	// Change a block transaction
	fmt.Println("Changing transaction of block 2...")
	assignment01bca.ChangeBlock(1, "Maheen to eve")

	// Verify the blockchain
	fmt.Println("Verifying blockchain...")
	assignment01bca.VerifyChain()

	// List blocks again to see changes
	assignment01bca.ListBlocks()
}
