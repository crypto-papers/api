package server

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/crypto-papers/api/config"
	postgres "github.com/crypto-papers/api/db"
	gen "github.com/crypto-papers/api/generated"
	res "github.com/crypto-papers/api/resolver"
)

// StartServer initiates a web-server at port set in environment (4000 if no port provided)
func StartServer() {
	conf := config.New()
	port := conf.GQL.Port

	db, dberr := postgres.Connect()
	handleErr(dberr)

	postgres.CheckSchemaVersion(db)

	http.Handle("/", handler.Playground("Cryptopapers", "/query"))
	http.Handle("/query", handler.GraphQL(gen.NewExecutableSchema(res.NewRootResolvers(db))))

	log.Printf("Server is running on http://localhost:%s", port)

	err := http.ListenAndServe(":"+port, nil)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
