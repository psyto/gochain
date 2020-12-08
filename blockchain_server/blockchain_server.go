package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gochain/block"
	"github.com/gochain/wallet"
)

var cache map[string]*block.Blockchain = make(map[string]*block.Blockchain)

// BlockchainServer consists of port number.
type BlockchainServer struct {
	port uint16
}

// NewBlockchainServer creates a BlockchainServer with a port number.
func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}

// Port returns a port number.
func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

// GetBlockchain is an API to return a Blockchain.
func (bcs *BlockchainServer) GetBlockchain() *block.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minersWallet := wallet.NewWallet()
		bc = block.NewBlockchain(minersWallet.BlockchainAddress(), bcs.Port())
		cache["blockchain"] = bc
		log.Printf("private_key %v", minersWallet.PrivateKeyStr())
		log.Printf("public_key %v", minersWallet.PublicKeyStr())
		log.Printf("blockchain_address %v", minersWallet.BlockchainAddress())
	}
	return bc
}

// GetChain writes "Hello World!" over the network.
func (bcs *BlockchainServer) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockchain()
		m, _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}

// Run runs BlockchainServer.
func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", bcs.GetChain)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil))
}
