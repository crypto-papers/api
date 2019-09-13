package server

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/rs/cors"

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

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:4000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		// Debug:            true,
	}).Handler)

	router.Handle("/", handler.Playground("Cryptopapers", "/query"))
	router.Handle("/query", handler.GraphQL(
		gen.NewExecutableSchema(res.NewRootResolvers(db)),
	))

	log.Printf("Server is running on http://localhost:%s", port)

	err := http.ListenAndServe(":"+port, router)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
