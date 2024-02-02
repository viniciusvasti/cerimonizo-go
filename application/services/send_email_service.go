package services

import (
	"log"
	"net/smtp"
	"os"
)

func SendEmail(email string) error {
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