package main

import (
	"../core"
	"encoding/json"
	"io"
	"net/http"
)

var blockchain *core.BlockChain

func run() {
	http.HandleFunc("/blockchain/get", handleGetRequest)
	http.HandleFunc("/blockchain/write", handleWriteRequest)
	http.ListenAndServe(":12345", nil)
}
func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	bytes, error := json.Marshal(blockchain)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}
func handleWriteRequest(w http.ResponseWriter, r *http.Request) {
	blockdata := r.URL.Query().Get("data")
	blockchain.SendData(blockdata)
	handleGetRequest(w, r)
}
func main() {
	blockchain = core.NewBlockchain()
	run()
}
