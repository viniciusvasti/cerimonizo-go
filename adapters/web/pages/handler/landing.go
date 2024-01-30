package page_handler

import (
	"log"
	"viniciusvasti/cerimonize/view/landing"

	"github.com/labstack/echo/v4"
)

type LandingPageHandler struct {
}

func (lh LandingPageHandler) Handle(c echo.Context) error {
	log.Printf("new access from %s", c.RealIP())
	registered := c.QueryParam("registered") == "true"
	return render(c, landing.Show(registered))
}
