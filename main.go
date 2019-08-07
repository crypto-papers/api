package main

import (
	"log"

	"github.com/crypto-papers/api/server"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	server.StartServer()
}
