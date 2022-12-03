package internal

import (
	"crypto/tls"
	"fmt"

	"github.com/ysfglmzz/e-shop-microservices/email/config"
	"github.com/ysfglmzz/e-shop-microservices/email/internal/event"
	gomail "gopkg.in/mail.v2"
)

func SendEmail(cfg config.SystemConfig, userCreatedEvent event.UserCreatedEvent) error {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", cfg.FromEmail)

	// Set E-Mail receivers
	m.SetHeader("To", userCreatedEvent.UserEmail)

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	verifyInfo := fmt.Sprintf("Verify link: http://localhost:5001/user/%d", userCreatedEvent.UserId)
	m.SetBody("text/plain", verifyInfo)

	// Settings for SMTP server
	d := gomail.NewDialer(cfg.EmailHost, cfg.EmailPort, cfg.FromEmail, cfg.FromEmailPassword)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
