package handler

import (
	"encoding/json"
	"net/http"

	gophersql "github.com/graph-gophers/graphql-go"

	"github.com/crypto-papers/api/loader"
)

// GraphQL handler handles GraphQL API requests over HTTP.
type GraphQL struct {
	Schema  *gophersql.Schema
	Loaders loader.Collection
}

func (h *GraphQL) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var query struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}

	if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := h.Loaders.Attach(r.Context())

	response := h.Schema.Exec(ctx, query.Query, query.OperationName, query.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
