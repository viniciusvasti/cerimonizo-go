package ports

import "viniciusvasti/cerimonize/application"

type WeddingRepositoryInterface interface {
	Get(id string) (application.WeddingInterface, error)
	GetAll() ([]application.WeddingInterface, error)
	Save(wedding application.WeddingInterface) (application.WeddingInterface, error)
}
