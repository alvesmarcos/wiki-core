package main

import (
	"log"
	"wikilibras-core/src/config"
	"wikilibras-core/src/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.LoadConfig()

	server.Start()
}
