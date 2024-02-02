package ports

import (
	"time"
	"viniciusvasti/cerimonize/application"
)

type WeddingServiceInterface interface {
	Get(id string) (application.WeddingInterface, error)
	GetAll() ([]application.WeddingInterface, error)
	Create(name string, date time.Time, budget float64) (application.WeddingInterface, error)
	Update(wedding application.WeddingInterface) (application.WeddingInterface, error)
	Enable(wedding application.WeddingInterface) error
	Disable(wedding application.WeddingInterface) error
}
