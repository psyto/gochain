package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

// Block is a struct of Block.
type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

// NewBlock is a function to create a new block.
func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}

// Hash is a function to hash a block.
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256(m)
}

// MarshalJSON is an override of default function to marshal block properly.
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Timestamp    int64          `json:"timestamp"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Timestamp:    b.timestamp,
		Transactions: b.transactions,
	})
}

// Print is a helper function to print Block in pretty format.
func (b *Block) Print() {
	fmt.Printf("timestamp:        %d\n", b.timestamp)
	fmt.Printf("nonce:            %d\n", b.nonce)
	fmt.Printf("previous_hash:    %x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}
