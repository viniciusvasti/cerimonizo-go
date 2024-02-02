package services

import (
	"time"
	"viniciusvasti/cerimonize/application"
	"viniciusvasti/cerimonize/application/ports"
)

type WeddingService struct {
	Repository ports.WeddingRepositoryInterface
}

func NewWeddingService(repository ports.WeddingRepositoryInterface) WeddingService {
	return WeddingService{
		Repository: repository,
	}
}

func (ws WeddingService) Get(id string) (application.WeddingInterface, error) {
	wedding, err := ws.Repository.Get(id)
	if err != nil {
		return nil, err
	}
	return wedding, nil
}

func (ws WeddingService) GetAll() ([]application.WeddingInterface, error) {
	weddings, err := ws.Repository.GetAll()
	if err != nil {
		return nil, err
	}
	return weddings, nil
}

func (ws WeddingService) Create(name string, date time.Time, budget float64) (application.WeddingInterface, error) {
	wedding, err := application.NewWedding(name, date, budget)
	if err != nil {
		return nil, err
	}
	createdWedding, err := ws.Repository.Save(wedding)
	if err != nil {
		return nil, err
	}
	return createdWedding, nil
}

func (ws WeddingService) Update(wedding application.WeddingInterface) (application.WeddingInterface, error) {
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

func (ws WeddingService) Enable(wedding application.WeddingInterface) error {
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

func (ws WeddingService) Disable(wedding application.WeddingInterface) error {
	wedding.Disable()
	_, err := ws.Repository.Save(wedding)
	if err != nil {
		return err
	}
	return nil
}
