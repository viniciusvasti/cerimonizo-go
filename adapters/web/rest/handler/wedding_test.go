package rest_handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
	rest_handler "viniciusvasti/cerimonize/adapters/web/rest/handler"
	"viniciusvasti/cerimonize/application"

	"github.com/labstack/echo/v4"
)

var (
	wedding1, _ = application.NewWedding("Wedding Test", time.Now().Add(time.Hour*24*30), 1000)
	wedding2, _ = application.NewWedding("Wedding Test 2", time.Now().Add(time.Hour*24*30*2), 2000)
)

type WeddingServiceMock struct{}

func (wsm WeddingServiceMock) Get(id string) (application.WeddingInterface, error) {
	wedding1.ID = "1"

	if id != wedding1.ID {
		return nil, errors.New("sql: no rows in result set")
	}

	return wedding1, nil
}

func (wsm WeddingServiceMock) GetAll() ([]application.WeddingInterface, error) {
	wedding1.ID = "1"
	wedding2.ID = "2"
	return []application.WeddingInterface{wedding1, wedding2}, nil
}

func (wsm WeddingServiceMock) Create(name string, date time.Time, budget float64) (application.WeddingInterface, error) {
	return application.NewWedding(name, date, budget)
}

func (wsm WeddingServiceMock) Update(wedding application.WeddingInterface) (application.WeddingInterface, error) {
	return wedding, nil
}

func (wsm WeddingServiceMock) Enable(wedding application.WeddingInterface) error {
	return nil
}

func (wsm WeddingServiceMock) Disable(wedding application.WeddingInterface) error {
	return nil
}

func TestGet(t *testing.T) {
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/weddings/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/weddings/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		wh := rest_handler.WeddingRestHandler{Service: WeddingServiceMock{}}
		getHandler := wh.HandleGet

		if err := getHandler(c); err != nil {
			t.Errorf("Error: %v", err)
		}

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code 200, got %v", rec.Code)
		}

		budget := strconv.FormatFloat(wedding1.GetBudget(), 'f', -1, 64)
		bodyAsJson := `{"id":` + wedding1.ID + `,"name":` + wedding1.Name + `,"date":"` + wedding1.GetDate().String() + `","budget":` + budget + `,"status":` + wedding1.Status + `}`

		if reflect.DeepEqual(bodyAsJson, rec.Body.String()) {
			t.Errorf("Expected body %v, got %v", bodyAsJson, rec.Body.String())
		}
	})

	t.Run("Wedding not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/weddings/0", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/weddings/:id")
		c.SetParamNames("id")
		c.SetParamValues("0")
		wh := rest_handler.WeddingRestHandler{Service: WeddingServiceMock{}}
		getHandler := wh.HandleGet

		err := getHandler(c)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}

		if err.(*echo.HTTPError).Code != http.StatusNotFound {
			t.Errorf("Expected status code 404, got %v", rec.Code)
		}
	})
}

func TestGetAll(t *testing.T) {
	e := echo.New()
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/weddings", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		wh := rest_handler.WeddingRestHandler{Service: WeddingServiceMock{}}
		getAllHandler := wh.HandleGetAll

		if err := getAllHandler(c); err != nil {
			t.Errorf("Error: %v", err)
		}

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code 200, got %v", rec.Code)
		}

		budget1 := strconv.FormatFloat(wedding1.GetBudget(), 'f', -1, 64)
		budget2 := strconv.FormatFloat(wedding2.GetBudget(), 'f', -1, 64)
		bodyAsJson := `[{"id":` + wedding1.ID + `,"name":` + wedding1.Name + `,"date":"` + wedding1.GetDate().String() + `","budget":` + budget1 + `,"status":` + wedding1.Status + `},{"id":` + wedding2.ID + `,"name":` + wedding2.Name + `,"date":"` + wedding2.GetDate().String() + `","budget":` + budget2 + `,"status":` + wedding2.Status + `}]`

		if reflect.DeepEqual(bodyAsJson, rec.Body.String()) {
			t.Errorf("Expected body %v, got %v", bodyAsJson, rec.Body.String())
		}
	})
}

func TestCreate(t *testing.T) {
	e := echo.New()
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/weddings", strings.NewReader(`{"name":"Wedding Test","date":"2024-10-28 16:01:21-03:00","budget":1000}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		wh := rest_handler.WeddingRestHandler{Service: WeddingServiceMock{}}
		createHandler := wh.HandleCreate

		if err := createHandler(c); err != nil {
			t.Errorf("Error: %v", err)
		}

		if rec.Code != http.StatusCreated {
			t.Errorf("Expected status code 201, got %v", rec.Code)
		}
	})

	t.Run("Invalid request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/weddings", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		wh := rest_handler.WeddingRestHandler{Service: WeddingServiceMock{}}
		createHandler := wh.HandleCreate

		err := createHandler(c)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}

		if err.(*echo.HTTPError).Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code 422, got %v", rec.Code)
		}
	})
}

func TestUpdate(t *testing.T) {
	e := echo.New()
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/weddings", strings.NewReader(`{"id":"1","name":"Wedding Test","date":"2024-10-28 16:01:21-03:00","budget":1000,"status":"enabled"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		wh := rest_handler.WeddingRestHandler{Service: WeddingServiceMock{}}
		updateHandler := wh.HandleUpdate

		if err := updateHandler(c); err != nil {
			t.Errorf("Error: %v", err)
		}

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code 200, got %v", rec.Code)
		}
	})

	t.Run("Invalid request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/weddings", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		wh := rest_handler.WeddingRestHandler{Service: WeddingServiceMock{}}
		updateHandler := wh.HandleUpdate

		err := updateHandler(c)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}

		if err.(*echo.HTTPError).Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code 422, got %v", rec.Code)
		}
	})
}
