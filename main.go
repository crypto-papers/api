//go:generate go run schema/generate/schema_generate.go

package main

import "github.com/crypto-papers/Cryptopapers_Graph_Api/server"

func main() {
	server.StartServer()
}
