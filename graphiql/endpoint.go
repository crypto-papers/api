package graphiql

import (
	"bytes"
	"html/template"
	"net/http"
)

// TemplateHandler accepts a endpoint value and inserts it into te template
type TemplateHandler struct {
	template *template.Template
	Endpoint string
}

// EndpointHandler accepts an endpoint value and serves a graphiql playground at that endpoint
func EndpointHandler(endpoint string) (*TemplateHandler, error) {
	t, err := renderTemplate()
	if err != nil {
		return nil, err
	}

	return &TemplateHandler{Endpoint: endpoint, template: t}, nil
}

func (h *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	graphiql := new(bytes.Buffer)

	h.template.Execute(w, getVariables(h.Endpoint))

	w.Header().Set("Content-Type", "text/html")
	w.Write(graphiql.Bytes())
}
