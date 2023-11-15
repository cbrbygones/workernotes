package main

import (
	"fmt"
	"time"
)

type Block struct {
	Header       BlockHeader
	Transactions []Transactions
}

type BlockHeader struct {
	Timestamp    time.Time      // Time at which the block was created
	Transactions []Transactions // All transactions made on the chain
	PreviousHash string         // Hash of the previous block
	Hash         string         // Hash of the current block
}

type Transactions struct {
	From      string // The party issuing the labor voucher
	To        string // The party receiving the labor voucher
	Amount    int    // The amount in labor hours
	Signature string // Signature verifying the transaction
	// Proof-of-Authority may work best for this project.
}

type Blockchain struct {
	Blocks []*Block
}

func CreateBlock(Transactions []Transactions, PreviousHash string) *Block {
	// Creation of a block
	block := &Block{
		Header: BlockHeader{
			Timestamp:    time.Now(),
			Transactions: Transactions,
			PreviousHash: PreviousHash,
			Hash:         "calculate hash here",
		},
	}

	return block
}

func main() {
	// Prepares an empty blockchain
	blockchain := &Blockchain{
		Blocks: []*Block{},
	}

	// Prepares some transactions from issuer to recipient.
	//
	// Labor vouchers are represented in hours (i.e. 8 hours),
	// and they are destroyed upon use as to not be accumulated.
	transactions := []Transactions{
		{
			From:      "Issuer_a",
			To:        "Recipient_a",
			Amount:    8,
			Signature: "Signature_a",
		},
		{
			From:      "Issuer_b",
			To:        "Recipient_b",
			Amount:    8,
			Signature: "Signature_b",
		},
	}

	// Sets the previous hash, or the hash of the last block on the chain.
	PreviousHash := "PreviousHash"

	// Calls the CreateBlock func and generates a new block based on the transactions struct and PreviousHash variable
	newBlock := CreateBlock(transactions, PreviousHash)

	// This calculates the hash for the new block's header.
	newBlock.Header.Hash = calculateHash(newBlock.Header)

	// Adds the block to the block chain.
	blockchain.Blocks = append(blockchain.Blocks, newBlock)

	// Verifies that the blockchain creation was a success!
	printBlockchain(blockchain)

}

func calculateHash(header BlockHeader) string {
	return "actual_calculated_hash"
}

func printBlockchain(chain *Blockchain) {
	for _, block := range chain.Blocks {
		fmt.Printf("Timestamp: %v\n ", block.Header.Timestamp)
		fmt.Printf("Transactions: %+v\n ", block.Header.Transactions)
		fmt.Printf("Hash: %v\n ", block.Header.Hash)
		fmt.Printf("------")
	}
}
