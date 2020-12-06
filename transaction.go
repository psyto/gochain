package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Transaction is a struct for transaction.
type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

// NewTransaction is a function to create a new transaction.
func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

// Print is a helper function to print transaction in pretty format.
func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("  sender_blockchain_address     %s\n", t.senderBlockchainAddress)
	fmt.Printf("  recipient_blockchain_address  %s\n", t.recipientBlockchainAddress)
	fmt.Printf("  value                         %.1f\n", t.value)
}

// MarshalJSON is an override of default function to marshal transaction properly.
func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}
