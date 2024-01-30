package main

import (
	// "log"
	// "os"

	"viniciusvasti/cerimonize/adapters/web/rest"
	// "github.com/joho/godotenv"
)

func main() {
	// if os.Getenv("ENV") == "dev" {
	// 	err := godotenv.Load()
	// 	if err != nil {
	// 		log.Panic("Error loading .env file")
	// 	}
	// }
	server := rest.Server{}
	server.Serve()
}
