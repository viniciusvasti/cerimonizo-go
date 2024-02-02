package sqldb_test

import (
	"database/sql"
	"log"
	"testing"
	"time"
	"viniciusvasti/cerimonize/adapters/sqldb"
	"viniciusvasti/cerimonize/application"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func setup() {
	var err error
	database, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err.Error())
	}
	createTable(database)
	createWedding(database)
}

func createTable(db *sql.DB) {
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

	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
}

func createWedding(db *sql.DB) {
	insert := `
		INSERT INTO weddings (id, name, date, budget, status)
		VALUES ('1', 'Wedding 1', '2024-10-11 19:45:52-03:00', 1000.00, 'disabled');
	`
	statement, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

func TestGet(t *testing.T) {
	setup()
	defer database.Close()

	weddingDb := sqldb.NewWeddingSQLRepository(database)
	wedding, err := weddingDb.Get("1")

	if err != nil {
		t.Fatalf("Expected wedding to be retrieved, but got error: %s", err.Error())
	}
	if wedding == nil {
		t.Fatalf("Expected wedding to be retrieved, but got nil")
	}
	assertEqual(t, "1", wedding.GetId())
	assertEqual(t, "Wedding 1", wedding.GetName())
	assertEqual(t, "2024-10-11", wedding.GetDate().Format("2006-01-02"))
	assertEqual(t, 1000.00, wedding.GetBudget())
	assertEqual(t, "disabled", wedding.GetStatus())
}

func TestSave(t *testing.T) {
	setup()
	defer database.Close()

	weddingDb := sqldb.NewWeddingSQLRepository(database)
	wedding := &application.Wedding{
		ID:     "2",
		Name:   "Wedding 2",
		Date:   time.Date(2024, 10, 11, 19, 45, 52, 0, time.UTC),
		Budget: 1000.00,
		Status: "disabled",
	}
	_, err := weddingDb.Save(wedding)
	if err != nil {
		t.Fatalf("Expected wedding to be saved, but got error: %s", err.Error())
	}

	createdWedding, err := weddingDb.Get("2")
	if err != nil {
		t.Fatalf("Expected the created wedding to be retrieved, but got error: %s", err.Error())
	}

	if createdWedding == nil {
		t.Fatalf("Expected get the created wedding, but got nil")
	}
	assertEqual(t, "2", createdWedding.GetId())
	assertEqual(t, "Wedding 2", createdWedding.GetName())
	assertEqual(t, "2024-10-11", createdWedding.GetDate().Format("2006-01-02"))
	assertEqual(t, 1000.00, createdWedding.GetBudget())
	assertEqual(t, "disabled", createdWedding.GetStatus())

	wedding.Status = "enabled"
	_, err = weddingDb.Save(wedding)

	if err != nil {
		t.Fatalf("Expected wedding to be saved, but got error: %s", err.Error())
	}
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
