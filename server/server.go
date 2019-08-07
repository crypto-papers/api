package server

import (
	"database/sql"
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

	connectDb()

	http.Handle("/", handler.Playground("Cryptopapers", "/query"))
	http.Handle("/query", handler.GraphQL(gen.NewExecutableSchema(gen.Config{Resolvers: &res.Resolver{}})))

	log.Printf("Server is running on http://localhost:%s", port)

	err := http.ListenAndServe(":"+port, nil)
	handleErr(err)
}

func connectDb() {
	db, err := postgres.Connect()
	handleErr(err)

	initDB(db)
}

func initDB(db *sql.DB) {
	postgres.ExecQuery(db, "CREATE TABLE public.papers (id SERIAL PRIMARY KEY, title varchar(255))")
	postgres.ExecQuery(db, "INSERT INTO papers(title) VALUES('test')")
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
