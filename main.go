package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"time"
)

// Block represents each 'item' in the blockchain
// Block struct contains basic block properties
// Index: position of the data record in the blockchain
// Timestamp: time of block creation
// Data: actual data stored in the block
// PreviousHash: hash of the previous block
// Hash: hash of the current block

type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PreviousHash string
	Hash         string
}

// calculateHash generates a SHA256 hash of a block's contents
func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.Data, block.PreviousHash)
	h := sha256.New()
	h.Write([]byte(record))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

// Blockchain is a series of validated Blocks
var Blockchain []Block

// createGenesisBlock creates the first block in the blockchain
func createGenesisBlock() Block {
	return Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Data:         "Genesis Block",
		PreviousHash: "",
		Hash:         "",
	}
}

// generateBlock creates a new block using previous block's hash
func generateBlock(oldBlock Block, data string) Block {
	newBlock := Block{
		Index:        oldBlock.Index + 1,
		Timestamp:    time.Now().String(),
		Data:         data,
		PreviousHash: oldBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func printBlockchain() {
	for _, block := range Blockchain {
		fmt.Printf("Index: %d\nTimestamp: %s\nData: %s\nPreviousHash: %s\nHash: %s\n\n",
			block.Index, block.Timestamp, block.Data, block.PreviousHash, block.Hash)
	}
}

func main() {
	genesisBlock := createGenesisBlock()
	genesisBlock.Hash = calculateHash(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Blockchain CLI:")
		fmt.Println("1. Add new block")
		fmt.Println("2. View blockchain")
		fmt.Println("3. Exit")
		fmt.Print("Enter choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter block data: ")
			scanner.Scan()
			data := scanner.Text()
			newBlock := generateBlock(Blockchain[len(Blockchain)-1], data)
			Blockchain = append(Blockchain, newBlock)
			fmt.Println("Block added!\n")
		case "2":
			printBlockchain()
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
