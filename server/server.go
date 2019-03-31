package server

import (
	"net/http"

	// gophersql "github.com/graph-gophers/graphql-go"

	"github.com/crypto-papers/Cryptopapers_Graph_Api/graphiql"
	// "github.com/crypto-papers/Cryptopapers_Graph_Api/handler"
	// "github.com/crypto-papers/Cryptopapers_Graph_Api/loader"
	// "github.com/crypto-papers/Cryptopapers_Graph_Api/schema"
)

type query struct{}

// StartServer initiates a web-server at port 3000
func StartServer() {
	// Load static files for homepage
	fs := http.FileServer(http.Dir("static/js/dist"))

	// h := handler.GraphQL{
	// 	Schema: gophersql.MustParseSchema(schema.GetParsableSchema(), &query{}),
	// 	Loaders: loader.Initialize(c)
	// }

	graphiqlHandler, err := graphiql.EndpointHandler("/graphiql")
	if err != nil {
		panic(err)
	}

	http.Handle("/", fs)
	http.Handle("/graphiql", graphiqlHandler)
	// http.Handle("/query", h)
	http.ListenAndServe(":3000", nil)
}
