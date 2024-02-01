package services_test

import (
	"testing"
	"time"
	"viniciusvasti/cerimonize/application"
	"viniciusvasti/cerimonize/application/services"
)

// A trick to assert the count of calls to a method
type weddingRepositoryStub struct {
	GetCalls    int
	GetAllCalls int
	SaveCalls   int
}

// It's important to receive the weddingRepositoryStub as a pointer, so that the calls to the methods are counted
func (wrs *weddingRepositoryStub) Get(id string) (application.WeddingInterface, error) {
	wrs.GetCalls++
	return &application.Wedding{
		ID:     "123",
		Name:   "John and Mary's",
		Date:   time.Now().AddDate(0, 0, 1),
		Status: application.ENABLED,
		Budget: 10000,
	}, nil
}

func (wrs *weddingRepositoryStub) GetAll() ([]application.WeddingInterface, error) {
	wrs.GetAllCalls++
	return []application.WeddingInterface{
		&application.Wedding{
			ID:     "123",
			Name:   "John and Mary's",
			Date:   time.Now().AddDate(0, 0, 1),
			Status: application.ENABLED,
			Budget: 10000,
		},
		&application.Wedding{
			ID:     "456",
			Name:   "Peter and Jane's",
			Date:   time.Now().AddDate(0, 0, 2),
			Status: application.ENABLED,
			Budget: 20000,
		},
	}, nil
}

func (wrs *weddingRepositoryStub) Save(wedding application.WeddingInterface) (application.WeddingInterface, error) {
	wrs.SaveCalls++
	return &application.Wedding{
		ID:     "123",
		Name:   wedding.GetName(),
		Date:   wedding.GetDate(),
		Status: wedding.GetStatus(),
		Budget: wedding.GetBudget(),
	}, nil
}

func Test_WeddingService(t *testing.T) {
	t.Run("Should get a wedding", func(t *testing.T) {
		weddingRepository := weddingRepositoryStub{}
		weddingService := services.NewWeddingService(&weddingRepository)

		wedding, err := weddingService.Get("123")

		if err != nil {
			t.Errorf("Expected wedding to be retrieved, but got error: %s", err.Error())
		}
		if wedding == nil {
			t.Error("Expected wedding to be retrieved, but got nil")
		}
		if weddingRepository.GetCalls != 1 {
			t.Errorf("Expected wedding repository get to be called 1 time, but got %d", weddingRepository.GetCalls)
		}
	})

	t.Run("Should get all weddings", func(t *testing.T) {
		weddingRepository := weddingRepositoryStub{}
		weddingService := services.NewWeddingService(&weddingRepository)

		weddings, err := weddingService.GetAll()

		if err != nil {
			t.Errorf("Expected weddings to be retrieved, but got error: %s", err.Error())
		}
		if weddings == nil {
			t.Error("Expected weddings to be retrieved, but got nil")
		}
		if len(weddings) != 2 {
			t.Errorf("Expected weddings to have 2 items, but got %d", len(weddings))
		}
		if weddingRepository.GetAllCalls != 1 {
			t.Errorf("Expected wedding repository get all to be called 1 time, but got %d", weddingRepository.GetAllCalls)
		}
	})

	t.Run("Should create a wedding", func(t *testing.T) {
		weddingRepository := weddingRepositoryStub{}
		weddingService := services.NewWeddingService(&weddingRepository)

		wedding, err := weddingService.Create("John and Mary's", time.Now().AddDate(0, 0, 1), 10000)

		if err != nil {
			t.Errorf("Expected wedding to be created, but got error: %s", err.Error())
		}
		if wedding == nil {
			t.Error("Expected wedding to be created, but got nil")
		}
		if weddingRepository.SaveCalls != 1 {
			t.Errorf("Expected wedding repository save to be called 1 time, but got %d", weddingRepository.SaveCalls)
		}
	})

	t.Run("Should enable a wedding", func(t *testing.T) {
		weddingRepository := weddingRepositoryStub{}
		weddingService := services.NewWeddingService(&weddingRepository)

		wedding := &application.Wedding{
			ID:     "123",
			Name:   "John and Mary's",
			Date:   time.Now().AddDate(0, 0, 1),
			Status: application.DISABLED,
			Budget: 10000,
		}
		err := weddingService.Enable(wedding)

		if err != nil {
			t.Errorf("Expected wedding to be enabled, but got error: %s", err.Error())
		}
		if wedding.GetStatus() != application.ENABLED {
			t.Errorf("Expected wedding status to be '%s', but got '%s'", application.ENABLED, wedding.GetStatus())
		}
		if weddingRepository.SaveCalls != 1 {
			t.Errorf("Expected wedding repository save to be called 1 times, but got %d", weddingRepository.SaveCalls)
		}
	})
}
