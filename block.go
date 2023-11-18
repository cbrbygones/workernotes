package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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

func calculateHash(header BlockHeader) string {
	// Serializes the header field into a byte slice.
	headerBytes, err := json.Marshal(header)
	if err != nil {
		panic(err)
	}

	// Creates a new SHA256 instance and writes a slice (in bytes) of string data.
	hash := sha256.New()
	hash.Write(headerBytes)
	hashInBytes := hash.Sum(nil)

	// Converts hashInBytes to hexidecimal
	hashString := hex.EncodeToString(hashInBytes)

	return hashString
}

func CreateBlock(Transactions []Transactions, PreviousHash string) *Block {
	// Creates a new SHA256 instance/writes byte slice for the previous hash.
	previousHashHash := sha256.New()
	previousHashHash.Write([]byte(PreviousHash))
	previousHashInBytes := previousHashHash.Sum(nil)
	PreviousHashEncoded := hex.EncodeToString(previousHashInBytes)

	// Creation of a block
	block := &Block{
		Header: BlockHeader{
			Timestamp:    time.Now(),
			Transactions: Transactions,
			PreviousHash: PreviousHashEncoded,
		},
	}

	block.Header.Hash = calculateHash(block.Header)

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

	// Adds the block to the block chain.
	blockchain.Blocks = append(blockchain.Blocks, newBlock)

	// Verifies that the blockchain creation was a success!
	printBlockchain(blockchain)

}

func printBlockchain(chain *Blockchain) {
	for _, block := range chain.Blocks {
		fmt.Printf("Timestamp: %v\n ", block.Header.Timestamp)
		fmt.Printf("Transactions: %+v\n ", block.Header.Transactions)
		fmt.Printf("Hash: %v\n ", block.Header.Hash)
		fmt.Printf("------")
	}
}
