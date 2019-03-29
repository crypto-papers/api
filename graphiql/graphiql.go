package graphiql

import "html/template"

// versionTmpl struct defines the imputs required by the graphiql template
type versionTmpl struct {
	Endpoint          string
	Es6PromiseVersion string
	FetchVersion      string
	ReactVersion      string
	GraphiqlVersion   string
}

const (
	templateName      = "graphiql"
	es6PromiseVersion = "4.2.6"
	fetchVersion      = "3.0.0"
	reactVersion      = "16.8.5"
	graphiqlVersion   = "0.13.0"
)

// GetVariables populates the versionTmpl struct with provided endpoint and values
func getVariables(endpoint string) *versionTmpl {
	return &versionTmpl{
		Endpoint:          endpoint,
		Es6PromiseVersion: es6PromiseVersion,
		FetchVersion:      fetchVersion,
		ReactVersion:      reactVersion,
		GraphiqlVersion:   graphiqlVersion,
	}
}

// RenderTemplate parses the required variables and plugs them into the graphiql template
func renderTemplate() (*template.Template, error) {
	t := template.New(templateName)
	t, err := t.Parse(graphiqlTemplate)

	if err != nil {
		return nil, err
	}

	return t, nil
}
