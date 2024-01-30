package application

import (
	"time"
)

type WeddingService struct {
	Repository WeddingRepositoryInterface
}

func NewWeddingService(repository WeddingRepositoryInterface) WeddingService {
	return WeddingService{
		Repository: repository,
	}
}

func (ws WeddingService) Get(id string) (WeddingInterface, error) {
	wedding, err := ws.Repository.Get(id)
	if err != nil {
		return nil, err
	}
	return wedding, nil
}

func (ws WeddingService) GetAll() ([]WeddingInterface, error) {
	weddings, err := ws.Repository.GetAll()
	if err != nil {
		return nil, err
	}
	return weddings, nil
}

func (ws WeddingService) Create(name string, date time.Time, budget float64) (WeddingInterface, error) {
	wedding, err := NewWedding(name, date, budget)
	if err != nil {
		return nil, err
	}
	createdWedding, err := ws.Repository.Save(wedding)
	if err != nil {
		return nil, err
	}
	return createdWedding, nil
}

func (ws WeddingService) Update(wedding WeddingInterface) (WeddingInterface, error) {
	_, err := wedding.IsValid()
	if err != nil {
		return nil, err
	}
	
	updatedWedding, err := ws.Repository.Save(wedding)
	if err != nil {
		return nil, err
	}
	return updatedWedding, nil
}

func (ws WeddingService) Enable(wedding WeddingInterface) error {
	enablingError := wedding.Enable()
	if enablingError != nil {
		return enablingError
	}
	_, err := ws.Repository.Save(wedding)
	if err != nil {
		return err
	}
	return nil
}

func (ws WeddingService) Disable(wedding WeddingInterface) error {
	wedding.Disable()
	_, err := ws.Repository.Save(wedding)
	if err != nil {
		return err
	}
	return nil
}
