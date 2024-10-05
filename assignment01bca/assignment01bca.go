package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// Color codes for pastel colors (light shades)
const (
	PastelGreen  = "\033[38;5;157m"
	PastelBlue   = "\033[38;5;117m"
	PastelPink   = "\033[38;5;218m"
	PastelPurple = "\033[38;5;141m"
	ResetColor   = "\033[0m"
)

// Block represents each 'block' in the blockchain
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

// Blockchain represents a series of blocks
var Blockchain []Block

// CalculateHash generates a SHA256 hash for a given string
func CalculateHash(stringToHash string) string {
	hash := sha256.New()
	hash.Write([]byte(stringToHash))
	return hex.EncodeToString(hash.Sum(nil))
}

// NewBlock creates a new block and returns a pointer to it
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.Hash = CalculateHash(transaction + string(nonce) + previousHash)
	return block
}

// ListBlocks prints all blocks in the blockchain with colorful output
func ListBlocks() {
	for i, block := range Blockchain {
		fmt.Printf(PastelGreen+"[]=> Block %d:\n"+ResetColor, i+1)
		fmt.Printf(PastelBlue+"[]=> Transaction: %s\n"+ResetColor, block.Transaction)
		fmt.Printf(PastelPurple+"[]=> Nonce: %d\n"+ResetColor, block.Nonce)
		fmt.Printf(PastelPink+"[]=> Previous Hash: %s\n"+ResetColor, block.PreviousHash)
		fmt.Printf(PastelGreen+"[]=> Hash: %s\n"+ResetColor, block.Hash)
		fmt.Println(strings.Repeat("-", 40))
	}
}

// ChangeBlock changes the transaction of the given block
func ChangeBlock(blockIndex int, newTransaction string) {
	if blockIndex >= 0 && blockIndex < len(Blockchain) {
		Blockchain[blockIndex].Transaction = newTransaction
		Blockchain[blockIndex].Hash = CalculateHash(newTransaction + string(Blockchain[blockIndex].Nonce) + Blockchain[blockIndex].PreviousHash)
	} else {
		fmt.Println(PastelPink + "Invalid block index." + ResetColor)
	}
}

// VerifyChain checks the validity of the entire blockchain
func VerifyChain() bool {
	for i := 1; i < len(Blockchain); i++ {
		if Blockchain[i].PreviousHash != Blockchain[i-1].Hash {
			fmt.Printf(PastelPink+"[]=> Blockchain is invalid at block %d.\n"+ResetColor, i+1)
			return false
		}
	}
	fmt.Println(PastelGreen + "[]=> Blockchain is valid." + ResetColor)
	return true
}
