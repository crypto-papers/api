package server

import (
	"context"
	"net/http"

	gophersql "github.com/graph-gophers/graphql-go"

	"github.com/crypto-papers/api/graphiql"
	"github.com/crypto-papers/api/handler"

	// "github.com/crypto-papers/api/loader"
	"github.com/crypto-papers/api/schema"
)

type query struct{}

// StartServer initiates a web-server at port 3000
func StartServer() {
	// Load static files for homepage
	fs := http.FileServer(http.Dir("static/js/dist"))

	// Set up graphiql playground
	graphiqlHandler, err := graphiql.EndpointHandler("/graphiql")
	if err != nil {
		panic(err)
	}

	// Set up GraphQL querying
	ctx := context.Background()

	graphqlSchema := gophersql.MustParseSchema(schema.GetParsableSchema(), &query{})

	graphqlHandler := &handler.GraphQL{
		Schema: graphqlSchema,
		// Loaders: loader.Initialize()
	}

	// Routing
	http.Handle("/", fs)
	http.Handle("/graphiql", graphiqlHandler)
	http.Handle("/query", handler.GraphQL{Schema: graphqlSchema})

	http.ListenAndServe(":3000", nil)
}
