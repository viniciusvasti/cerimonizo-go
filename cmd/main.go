package main

import (
	"database/sql"
	"log"
	"os"

	"viniciusvasti/cerimonize/adapters/sqldb"
	"viniciusvasti/cerimonize/adapters/web/rest"
	"viniciusvasti/cerimonize/application/services"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Panic("Error loading .env file")
		}
	}
	if os.Getenv("GMAIL_USERNAME") == "" {
		log.Panic("GMAIL_USERNAME not set")
	}
	if os.Getenv("GMAIL_PASSWORD") == "" {
		log.Panic("GMAIL_PASSWORD not set")
	}

	// Database
	database, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	
	weddingRepository := sqldb.NewWeddingSQLRepository(database)
	weddingService := services.NewWeddingService(weddingRepository)

	server := rest.Server{}
	server.Serve(weddingService)
}
