package rest

import (
	"log"
	"time"
	page_handler "viniciusvasti/cerimonize/adapters/web/pages/handler"
	rest_handler "viniciusvasti/cerimonize/adapters/web/rest/handler"
	"viniciusvasti/cerimonize/application/ports"
	"viniciusvasti/cerimonize/application/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
}

func (s *Server) Serve(weddingService ports.WeddingServiceInterface) {
	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.HideBanner = true
	app.Server.ReadTimeout = time.Second * 10
	app.Static("/public", "public")

	// Web App
	landingHandler := page_handler.LandingPageHandler{}
	app.GET("/", landingHandler.Handle)
	app.POST("/cadastrar", func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/json")
		newEmail := c.FormValue("email")
		log.Printf("New email: %s", newEmail)
		services.SendEmail(newEmail)
		return c.Redirect(302, "/?registered=true")
	})

	// REST API
	weddingHandler := rest_handler.WeddingRestHandler{
		Service: weddingService,
	}
	cerimonizoRoutes := app.Group("/api")
	makeWeddingRoutes(cerimonizoRoutes, weddingHandler)

	err := app.Start(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

func makeWeddingRoutes(cerimonizoRoutes *echo.Group, weddingHandler rest_handler.WeddingRestHandler) {
	weddingRoutes := cerimonizoRoutes.Group("/weddings")
	weddingRoutes.GET("", weddingHandler.HandleGetAll)
	weddingRoutes.GET("/:id", weddingHandler.HandleGet)
	weddingRoutes.POST("", weddingHandler.HandleCreate)
	weddingRoutes.PUT("/:id", weddingHandler.HandleUpdate)
}
