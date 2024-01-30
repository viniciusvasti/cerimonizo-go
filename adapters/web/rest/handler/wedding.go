package rest_handler

import (
	"net/http"
	"viniciusvasti/cerimonize/adapters/dto"
	"viniciusvasti/cerimonize/application"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type WeddingRestHandler struct {
	Service application.WeddingServiceInterface
}

func (wh WeddingRestHandler) HandleGet(c echo.Context) error {
	wedding, err := wh.Service.Get(c.Param("id"))
	if err != nil {
		// TODO: Find a better way to handle this error
		if err.Error() == "sql: no rows in result set" {
			return echo.NewHTTPError(http.StatusNotFound, "Wedding not found")
		}
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	weddingDTO := dto.WeddingDTO{}
	weddingDTO.Bind(wedding)
	return c.JSON(http.StatusOK, weddingDTO)
}

func (wh WeddingRestHandler) HandleGetAll(c echo.Context) error {
	result, err := wh.Service.GetAll()
	if err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	weddingsDTO := dto.BindAll(result)
	return c.JSON(http.StatusOK, weddingsDTO)
}

func (wh WeddingRestHandler) HandleCreate(c echo.Context) error {
	weddingDTO := dto.WeddingDTO{}
	if err := c.Bind(&weddingDTO); err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid request body")
	}

	wedding, err := weddingDTO.ConvertToEntity()
	if err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid request body")
	}

	createdWedding, err := wh.Service.Create(wedding.GetName(), wedding.GetDate(), wedding.GetBudget())
	if err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	createdDTO := dto.WeddingDTO{}
	createdDTO.Bind(createdWedding)
	return c.JSON(http.StatusCreated, createdDTO)
}

func (wh WeddingRestHandler) HandleUpdate(c echo.Context) error {
	weddingDTO := dto.WeddingDTO{}
	if err := c.Bind(&weddingDTO); err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid request body")
	}

	wedding, err := weddingDTO.ConvertToEntity()
	if err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid request body")
	}

	updatedWedding, err := wh.Service.Update(wedding)
	if err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	updatedDTO := dto.WeddingDTO{}
	updatedDTO.Bind(updatedWedding)
	return c.JSON(http.StatusOK, updatedDTO)
}
