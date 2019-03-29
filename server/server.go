package server

import (
	"fmt"
	"net/http"

	"github.com/crypto-papers/Cryptopapers_Graph_Api/graphiql"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

// StartServer initiates a web-server at port 3000
func StartServer() {
	graphiqlHandler, err := graphiql.EndpointHandler("/graphiql")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", indexHandler)
	http.Handle("/graphiql", graphiqlHandler)
	http.ListenAndServe(":3000", nil)
}
