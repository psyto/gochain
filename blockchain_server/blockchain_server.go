package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

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

// HelloWorld writes "Hello World!" over the network.
func HelloWorld(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!")
}

// Run runs BlockchainServer.
func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil))
}
