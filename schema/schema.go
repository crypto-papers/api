//go:generate go-bindata -ignore=\.go -pkg=schema -o=schema_bindata.go ./...
package schema

import "bytes"

// GetParsableSchema converts the graphql schema into a string to be used by graphql-go
func GetParsableSchema() string {
	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		b := MustAsset(name)
		buf.Write(b)

		// Add a newline if the file does not end in a newline.
		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}
