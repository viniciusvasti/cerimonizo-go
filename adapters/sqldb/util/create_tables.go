package util

import (
	"database/sql"
	"log"
)

func CreateTables(db *sql.DB) {
	table := `
		CREATE TABLE IF NOT EXISTS weddings (
			id TEXT NOT NULL PRIMARY KEY,
			name TEXT,
			date TEXT,
			budget REAL,
			status TEXT
		);
	`
	statement, err := db.Prepare(table)
	defer statement.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Weddings Table created")
}
