package service

import (
	"bytes"
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

const emailTemplate = `
	<html>
	<head>
	<style>
		table {
			font-family: arial, sans-serif;
			border-collapse: collapse;
			width: 100%;
		}
		td, th {
			border: 1px solid #dddddd;
			text-align: left;
			padding: 8px;
		}
	</style>
	</head>
	<body>
	<img src="https://upload.wikimedia.org/wikipedia/commons/e/e3/Stori_logo_vertical.png" alt="Local Image" style="width:200px;height:300px;">
	<h1>Balance Information:</h1>
	<p>Total Balance: ${{.TotalBalance}}</p>
	<p>Average Debit Amount: ${{.AverageDebitAmount | printf "%.2f"}}</p>
	<p>Average Credit Amount: ${{.AverageCreditAmount | printf "%.2f"}}</p>
	
	<h2>Transactions Per Month:</h2>
	<table>
	  <tr>
		<th>Month</th>
		<th>Total</th>
		<th>Credit Average</th>
		<th>Debit Average</th>
	  </tr>
	  {{range $key, $value := .TransactionsPerMonth}}
	  <tr> 
		<td>{{$key}}</td>
		<td>{{$value.Total}}</td>
		<td>{{$value.CreditAverage}}</td>
		<td>{{$value.DebitAverage}}</td>
	  </tr>
	  {{end}}
	</table>
	</body>
	</html>
	`

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
