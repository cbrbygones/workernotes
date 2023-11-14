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
	Stake     int    // The amount of stake required to validate each block
	// Proof-of-Stake may work best for this project.
}

type Blockchain struct {
	Blocks []*Block
}

func main() {
	fmt.Println("WorkerNotes", "\n", "Ok!")
	// fmt.Println("This is the header", "\n", "This is the body")
	CreateBlock("Labor", "voucher")
}

func CreateBlock(Header string, Body string) {
	fmt.Println(Header, "\n", Body)
}
