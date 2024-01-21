package main

import (
	"viniciusvasti/cerimonize/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	landingHandler := handler.LandingHandler{}
	app.Static("/public", "public")
	app.GET("/", landingHandler.HandleLanding)
	app.Start(":3000")
}