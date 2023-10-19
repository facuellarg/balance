package service

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"net/smtp"
	"text/template"

	"github.com/facuellarg/stori/domain/entities"
)

type (
	EmailService struct {
		email    string
		password string
		tmpl     *template.Template
	}
)

//go:embed layout.html
var emailTemplate string

func NewEmailService(email string, password string) EmailService {

	// Create a new template and parse the email template
	tmpl, err := template.New("emailTemplate").Parse(emailTemplate)
	if err != nil {
		log.Fatalf("Error creating template: %s", err)
	}
	return EmailService{email, password, tmpl}
}

func (e EmailService) Send(to string, balance entities.Balance) error {
	// Buffer to store the executed template
	var tpl bytes.Buffer
	auth := smtp.PlainAuth("", e.email, e.password, "smtp.gmail.com")

	// Execute the template with the balance data
	err := e.tmpl.Execute(&tpl, balance)
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	message := []byte("Subject: testing stori\r\n" +
		"From:" + e.email + " \r\n" +
		"To:" + to + " \r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		tpl.String())

	return smtp.SendMail("smtp.gmail.com:587", auth, e.email, []string{to}, message)
}
