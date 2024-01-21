package handler

import (
	"viniciusvasti/cerimonize/view/landing"

	"github.com/labstack/echo/v4"
)

type LandingHandler struct {
}

func (lh LandingHandler) HandleLanding(c echo.Context) error {
	return landing.Show().Render(c.Request().Context(), c.Response())
}
