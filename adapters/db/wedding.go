package db

import (
	"database/sql"
	"log"
	"time"
	"viniciusvasti/cerimonize/application"
)

type WeddingDB struct {
	db *sql.DB
}

func NewWeddingDB(db *sql.DB) *WeddingDB {
	return &WeddingDB{db: db}
}

func (p *WeddingDB) Get(id string) (application.WeddingInterface, error) {
	var wedding application.Wedding

	// Prepare statement for preventing SQL injection
	statement, err := p.db.Prepare("SELECT id, name, date, budget, status FROM weddings WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	dateString := ""
	err = statement.QueryRow(id).Scan(&wedding.ID, &wedding.Name, &dateString, &wedding.Budget, &wedding.Status)
	if err != nil {
		return nil, err
	}

	log.Println(dateString)
	wedding.Date, err = time.Parse("2006-01-02 15:04:05-07:00", dateString)
	if err != nil {
		return nil, err
	}

	return &wedding, nil
}

func (p *WeddingDB) Save(wedding application.WeddingInterface) (application.WeddingInterface, error) {
	var rows int
	p.db.QueryRow("SELECT COUNT(*) FROM weddings WHERE id = ?", wedding.GetId()).Scan(&rows)

	if rows == 0 {
		_, err := p.create(wedding)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(wedding)
		if err != nil {
			return nil, err
		}
	}

	return wedding, nil
}

func (p *WeddingDB) create(wedding application.WeddingInterface) (application.WeddingInterface, error) {
	statement, err := p.db.Prepare("INSERT INTO weddings (id, name, date, budget, status) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	_, err = statement.Exec(wedding.GetId(), wedding.GetName(), wedding.GetDate(), wedding.GetBudget(), wedding.GetStatus())
	if err != nil {
		return nil, err
	}

	return wedding, nil
}

func (p *WeddingDB) update(wedding application.WeddingInterface) (application.WeddingInterface, error) {
	statement, err := p.db.Prepare("UPDATE weddings SET name = ?, date = ?, budget = ?, status = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	_, err = statement.Exec(wedding.GetName(), wedding.GetDate(), wedding.GetBudget(), wedding.GetStatus(), wedding.GetId())
	if err != nil {
		return nil, err
	}

	return wedding, nil
}
