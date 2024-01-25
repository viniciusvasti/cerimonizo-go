package application_test

import (
	"testing"
	"time"
	"viniciusvasti/cerimonize/application"
)

func Test_Wedding_Enable(t *testing.T) {
	wedding := application.Wedding{
		Name:   "John and Mary's",
		Status: application.DISABLED,
		Date:   time.Now().AddDate(0, 0, 1),
	}

	t.Run("Should enable a wedding", func(t *testing.T) {
		err := wedding.Enable()

		if err != nil {
			t.Errorf("Expected wedding to be enabled, but got error: %s", err.Error())
		}
	})

	t.Run("Should not enable a wedding with a past date", func(t *testing.T) {
		wedding.Date = time.Now().AddDate(0, 0, -1)

		err := wedding.Enable()

		if err == nil {
			t.Error("Expected an error when enabling a wedding with a past date")
		}
		if err.Error() != "The wedding date must be a future date" {
			t.Errorf("Expected error message to be 'The wedding date must be a future date', but got '%s'", err.Error())
		}
	})
}

func Test_Wedding_IsValid(t *testing.T) {
	wedding := application.Wedding{
		Name:   "John and Mary's",
		Status: application.ENABLED,
		Date:   time.Now().AddDate(0, 0, 1),
	}

	t.Run("Should validate a wedding", func(t *testing.T) {
		isValid, err := wedding.IsValid()

		if !isValid || err != nil {
			t.Errorf("Expected wedding to be valid, but got error: %s", err.Error())
		}
	})

	t.Run("Should not validate a wedding without a name", func(t *testing.T) {
		wedding.Name = ""

		isValid, err := wedding.IsValid()

		if isValid || err == nil {
			t.Error("Expected wedding to be invalid")
		}

		if err.Error() != "The wedding name is required" {
			t.Errorf("Expected error message to be 'The wedding name is required', but got '%s'", err.Error())
		}
	})

	t.Run("Should not validate a wedding without a date", func(t *testing.T) {
		wedding.Name = "John and Mary's"
		wedding.Date = time.Time{}

		isValid, err := wedding.IsValid()

		if isValid || err == nil {
			t.Error("Expected wedding to be invalid")
		}

		if err.Error() != "The wedding date is required" {
			t.Errorf("Expected error message to be 'The wedding date is required', but got '%s'", err.Error())
		}
	})

	t.Run("Should not validate a wedding with an invalid status", func(t *testing.T) {
		wedding.Date = time.Now().AddDate(0, 0, 1)
		wedding.Status = "invalid"

		isValid, err := wedding.IsValid()

		if isValid || err == nil {
			t.Error("Expected wedding to be invalid")
		}

		if err.Error() != "The wedding status is invalid" {
			t.Errorf("Expected error message to be 'The wedding status is invalid', but got '%s'", err.Error())
		}
	})

	t.Run("Should not validate a wedding with an invalid budget", func(t *testing.T) {
		wedding.Status = application.ENABLED
		wedding.Budget = -1

		isValid, err := wedding.IsValid()

		if isValid || err == nil {
			t.Error("Expected wedding to be invalid")
		}

		if err.Error() != "The wedding budget is invalid" {
			t.Errorf("Expected error message to be 'The wedding budget is invalid', but got '%s'", err.Error())
		}
	})
}
