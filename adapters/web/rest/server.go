package rest

import (
	"database/sql"
	"log"
	"time"
	"viniciusvasti/cerimonize/adapters/sqldb"
	page_handler "viniciusvasti/cerimonize/adapters/web/pages/handler"
	rest_handler "viniciusvasti/cerimonize/adapters/web/rest/handler"
	"viniciusvasti/cerimonize/application"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
}

func (s *Server) Serve() {
	app := echo.New()
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
		return c.Redirect(302, "/?registered=true")
	})

	// REST API
	cerimonizoRoutes := app.Group("/api")
	makeWeddingRoutes(cerimonizoRoutes)

	err := app.Start(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

func makeWeddingRoutes(cerimonizoRoutes *echo.Group) {
	database, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	weddingRepository := sqldb.NewWeddingSQLRepository(database)
	weddingService := application.NewWeddingService(weddingRepository)
	weddingRoutes := cerimonizoRoutes.Group("/weddings")
	weddingHandler := rest_handler.WeddingRestHandler{
		Service: weddingService,
	}
	weddingRoutes.GET("", weddingHandler.HandleGetAll)
	weddingRoutes.GET("/:id", weddingHandler.HandleGet)
	weddingRoutes.POST("", weddingHandler.HandleCreate)
	weddingRoutes.PUT("/:id", weddingHandler.HandleUpdate)
}
