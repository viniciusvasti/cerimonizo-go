package application

import (
	"errors"
	"time"
)

type WeddingInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable()
	GetId() string
	GetName() string
	GetDate() time.Time
	GetStatus() string
	GetBudget() float64
}

const (
	ENABLED  = "enabled"
	DISABLED = "disabled"
)

type Wedding struct {
	ID     string
	Name   string
	Date   time.Time
	Budget float64
	Status string
}

func (w *Wedding) IsValid() (bool, error) {
	if w.Name == "" {
		return false, errors.New("The wedding name is required")
	}

	if w.Date.IsZero() {
		return false, errors.New("The wedding date is required")
	}

	if w.Status != ENABLED && w.Status != DISABLED {
		return false, errors.New("The wedding status is invalid")
	}

	if w.Budget < 0 {
		return false, errors.New("The wedding budget is invalid")
	}

	return true, nil
}

func (w *Wedding) Enable() error {
	if w.Date.Before(time.Now()) {
		return errors.New("The wedding date must be a future date")
	}
	w.Status = ENABLED
	return nil
}

func (w *Wedding) Disable() {
	w.Status = DISABLED
}

func (w *Wedding) GetId() string {
	return w.ID
}

func (w *Wedding) GetName() string {
	return w.Name
}

func (w *Wedding) GetDate() time.Time {
	return w.Date
}

func (w *Wedding) GetBudget() float64 {
	return w.Budget
}

func (w *Wedding) GetStatus() string {
	return w.Status
}
