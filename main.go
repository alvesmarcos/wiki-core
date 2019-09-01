package main

import (
	"wikilibras-core/src/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server.Start(true)
}
