package main

import (
	// "log"
	// "os"
	"log"
	"viniciusvasti/cerimonize/handler"

	// "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// if os.Getenv("ENV") == "dev" {
	// 	err := godotenv.Load()
	// 	if err != nil {
	// 		log.Panic("Error loading .env file")
	// 	}
	// }
	app := echo.New()

	landingHandler := handler.LandingHandler{}
	brideAndGroomChecklistHandler := handler.BrideAndGroomChecklistHandler{}
	suppliersHandler := handler.SuppliersHandler{}
	weddingsHandler := handler.WeddingsHandler{}
	agendaHandler := handler.AgendaHandler{}
	inspirationsHandler := handler.InspirationsHandler{}

	app.Static("/public", "public")
	app.Static("/public/img", "public/img")
	app.Static("/mock-up/public", "public")
	app.GET("/", landingHandler.HandleLanding)
	app.GET("/mock-up/checklist-noivos", brideAndGroomChecklistHandler.HandleChecklist)
	app.GET("/mock-up/fornecedores", suppliersHandler.Handle)
	app.GET("/mock-up/casamentos", weddingsHandler.Handle)
	app.GET("/mock-up/agenda-noivos", agendaHandler.HandleAgenda)
	app.GET("/mock-up/inspiracoes", inspirationsHandler.Handle)
	app.POST("/cadastrar", func(c echo.Context) error {
		newEmail := c.FormValue("email")
		log.Printf("New email: %s", newEmail)
		return c.Redirect(302, "/?registered=true")
	})
	app.Start(":3000")
}
