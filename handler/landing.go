package handler

import (
	"log"
	"viniciusvasti/cerimonize/view/landing"

	"github.com/labstack/echo/v4"
)

type LandingHandler struct {
}

func (lh LandingHandler) HandleLanding(c echo.Context) error {
	log.Printf("new access from %s", c.RealIP())
	registered := c.QueryParam("registered") == "true"
	return render(c, landing.Show(registered))
}