package main

import (
	"fmt"
	"log"

	"github.com/gochain/block"
	"github.com/gochain/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	// Wallet
	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(),
		walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)

	// Blockchain
	blockchain := block.NewBlockchain(walletM.BlockchainAddress())
	isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0,
		walletA.PublicKey(), t.GenerateSignature())
	fmt.Println("Added? ", isAdded)

	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress()))

	/*
		myBlockchainAddress := "my_blockchain_address"

		blockchain := block.NewBlockchain(myBlockchainAddress)
		blockchain.Print()

		blockchain.AddTransaction("A", "B", 1.0)
		blockchain.Mining()
		blockchain.Print()

		blockchain.AddTransaction("C", "D", 2.0)
		blockchain.AddTransaction("X", "Y", 3.0)
		blockchain.Mining()
		blockchain.Print()

		fmt.Printf("my %.1f\n", blockchain.CalculateTotalAmount("my_blockchain_address"))
		fmt.Printf("C %.1f\n", blockchain.CalculateTotalAmount("C"))
		fmt.Printf("D %.1f\n", blockchain.CalculateTotalAmount("D"))
	*/
}
