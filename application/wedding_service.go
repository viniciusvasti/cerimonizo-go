package application

import "time"

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

func (ws WeddingService) Create(name string, date time.Time) (WeddingInterface, error) {
	wedding, err := NewWedding(name, date)
	if err != nil {
		return nil, err
	}
	createdWedding, creatingErr := ws.Repository.Save(wedding)
	if creatingErr != nil {
		return nil, creatingErr
	}
	return createdWedding, nil
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
