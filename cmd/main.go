package main

import (
	// "log"
	// "os"
	"viniciusvasti/cerimonize/handler"

	// "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	landingHandler := handler.LandingHandler{}
	app.Static("/public", "public")
	app.Static("/public/img", "public/img")
	app.GET("/", landingHandler.HandleLanding)
	app.Start(":3000")
}
