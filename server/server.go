package server

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/crypto-papers/api/generated"
	"github.com/crypto-papers/api/resolver"
)

// StartServer initiates a web-server at port 3000
func StartServer() {
	http.Handle("/", handler.Playground("Cryptopapers", "/query"))
	
	http.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}})))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
