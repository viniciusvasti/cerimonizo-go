package main

import (
	"log"
	"net/smtp"
	"os"
	"viniciusvasti/cerimonize/handler"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Panic("Error loading .env file")
		}
	}
	if os.Getenv("GMAIL_USERNAME") == "" {
		log.Panic("GMAIL_USERNAME not set")
	}
	if os.Getenv("GMAIL_PASSWORD") == "" {
		log.Panic("GMAIL_PASSWORD not set")
	}

	app := echo.New()

	landingHandler := handler.LandingHandler{}
	// brideAndGroomChecklistHandler := handler.BrideAndGroomChecklistHandler{}
	// suppliersHandler := handler.SuppliersHandler{}
	// weddingsHandler := handler.WeddingsHandler{}
	// agendaHandler := handler.AgendaHandler{}
	// inspirationsHandler := handler.InspirationsHandler{}

	app.Static("/public", "public")
	app.Static("/public/img", "public/img")
	// app.Static("/mock-up/public", "public")
	app.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})
	app.GET("/", landingHandler.HandleLanding)
	// app.GET("/mock-up/checklist-noivos", brideAndGroomChecklistHandler.HandleChecklist)
	// app.GET("/mock-up/fornecedores", suppliersHandler.Handle)
	// app.GET("/mock-up/casamentos", weddingsHandler.Handle)
	// app.GET("/mock-up/agenda-noivos", agendaHandler.HandleAgenda)
	// app.GET("/mock-up/inspiracoes", inspirationsHandler.Handle)
	app.POST("/cadastrar", func(c echo.Context) error {
		newEmail := c.FormValue("email")
		log.Printf("New email: %s", newEmail)
		sendEmail(newEmail)
		return c.Redirect(302, "/?registered=true")
	})
	app.Start(":3000")
}

func sendEmail(email string) error {
	// Set up authentication information.
	auth := smtp.PlainAuth("", os.Getenv("GMAIL_USERNAME"), os.Getenv("GMAIL_PASSWORD"), "smtp.gmail.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("GMAIL_USERNAME"), []string{os.Getenv("GMAIL_USERNAME")}, []byte("To: "+os.Getenv("GMAIL_USERNAME")+"\r\n"+
		"Subject: New subscription to Cerimonizo\r\n"+
		"\r\n"+
		email))
	if err != nil {
		log.Printf("Error sending email, %s: %s", email, err)
	}
	return err
}

