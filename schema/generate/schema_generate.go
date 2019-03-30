package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

var fs http.FileSystem = http.Dir("./schema/graphql")

func main() {
	err := vfsgen.Generate(fs, vfsgen.Options{
		Filename:     "schema/schema_vfsdata.go",
		PackageName:  "schema",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
