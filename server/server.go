package server

import (
	"fmt"
	"net/http"
	// "github.com/crypto-papers/Crytopapers_Graph_Api/schema"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func StartServer() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":3000", nil)
}
