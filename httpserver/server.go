package main

import (
	"net/http"
	"encoding/json"
	"chain/core"
	"io"
)

var blockchain *core.Chain

func run() {
	// http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })
	http.HandleFunc("/blockmain/get", blockchainGet)
	http.HandleFunc("/blockmain/write", blockchainWrite)
	// func ListenAndServe(addr string, handler Handler) error
	http.ListenAndServe("localhost:8888", nil)
}


func blockchainGet(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(blockchain)
	if err != nil {
		// func Error(w ResponseWriter, error string, code int)
		http.Error(w, err.Error(), 500)
		return
	}
	// func WriteString(w Writer, s string) (n int, err error)
	io.WriteString(w, string(bytes))
}

func blockchainWrite(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	blockchain.SendData(data)
	blockchainGet(w, r)
}

func main() {
	blockchain = core.NewBlock()
	run()
}