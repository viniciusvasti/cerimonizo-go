package sqldb

import (
	"database/sql"
	"time"
	"viniciusvasti/cerimonize/adapters/sqldb/util"
	"viniciusvasti/cerimonize/application"
)

type WeddingSQLRepository struct {
	db *sql.DB
}

func NewWeddingSQLRepository(db *sql.DB) *WeddingSQLRepository {
	util.CreateTables(db)
	return &WeddingSQLRepository{db: db}
}

func (p *WeddingSQLRepository) Get(id string) (application.WeddingInterface, error) {
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

	wedding.Date, err = time.Parse("2006-01-02 15:04:05-07:00", dateString)
	if err != nil {
		return nil, err
	}

	return &wedding, nil
}

func (p *WeddingSQLRepository) GetAll() ([]application.WeddingInterface, error) {
	var weddings []application.WeddingInterface = make([]application.WeddingInterface, 0)

	// Prepare statement for preventing SQL injection
	statement, err := p.db.Prepare("SELECT id, name, date, budget, status FROM weddings")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var wedding application.Wedding
		dateString := ""
		err = rows.Scan(&wedding.ID, &wedding.Name, &dateString, &wedding.Budget, &wedding.Status)
		if err != nil {
			return nil, err
		}

		wedding.Date, err = time.Parse("2006-01-02 15:04:05-07:00", dateString)
		if err != nil {
			return nil, err
		}

		weddings = append(weddings, &wedding)
	}

	return weddings, nil
}

func (p *WeddingSQLRepository) Save(wedding application.WeddingInterface) (application.WeddingInterface, error) {
	var rows int
	statement, err := p.db.Prepare("SELECT COUNT(*) FROM weddings WHERE id = ?")
	statement.QueryRow(wedding.GetId()).Scan(&rows)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

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

func (p *WeddingSQLRepository) create(wedding application.WeddingInterface) (application.WeddingInterface, error) {
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

func (p *WeddingSQLRepository) update(wedding application.WeddingInterface) (application.WeddingInterface, error) {
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
