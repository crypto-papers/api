package server

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	gen "github.com/crypto-papers/api/generated"
	res "github.com/crypto-papers/api/resolver"
)

const defaultPort = "4000"

// StartServer initiates a web-server at port set in environment (4000 if no port provided)
func StartServer() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("Cryptopapers", "/query"))
	http.Handle("/query", handler.GraphQL(gen.NewExecutableSchema(gen.Config{Resolvers: &res.Resolver{}})))

	log.Printf("Server is running on http://localhost:%s", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
