package main

import (
	"fmt"
	"log"

	"github.com/gochain/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())

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
