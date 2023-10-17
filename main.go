package main

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"text/template"
)

type Transaction struct {
	ID     int
	Amount float32
	Month  int
	Day    int
}

type TransactionMonthInfo struct {
	Total         int
	CreditAverage float32
	DebitAverage  float32
}

type Balance struct {
	TotalBalance         float32
	TransactionsPerMonth map[string]TransactionMonthInfo
	AverageDebitAmount   float32
	AverageCreditAmount  float32
}

var (
	pass  = os.Getenv("STORI_PASSWORD")
	email = os.Getenv("STORI_EMAIL")
	to    = os.Getenv("STORI_TO")
)

func main() {

	auth := smtp.PlainAuth("", email, email, "smtp.gmail.com")

	loader := CSVLoaderTransformer{FileName: "data.csv"}
	transactions := loader.Load()
	fmt.Printf("transactions: %+v\n", transactions)
	balance := CreateBalance(transactions)
	fmt.Printf("balance: %+v\n", balance)
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
	// Create a new template and parse the email template
	tmpl, err := template.New("emailTemplate").Parse(emailTemplate)
	if err != nil {
		log.Fatalf("Error creating template: %s", err)
	}

	// Buffer to store the executed template
	var tpl bytes.Buffer

	// Execute the template with the balance data
	err = tmpl.Execute(&tpl, balance)
	if err != nil {
		log.Fatalf("Error executing template: %s", err)
	}

	message := []byte("Subject: testing stori\r\n" +
		"From:" + email + " \r\n" +
		"To:" + to + " \r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		tpl.String())
	if err := smtp.SendMail("smtp.gmail.com:587", auth, email, []string{to}, message); err != nil {
		log.Fatal(err)
	}
	// Print the executed template
	fmt.Println(tpl.String())

}

func CreateBalance(transactions []Transaction) Balance {

	balance := Balance{
		TotalBalance:         0,
		TransactionsPerMonth: make(map[string]TransactionMonthInfo),
		AverageDebitAmount:   0,
		AverageCreditAmount:  0,
	}

	balanceTotalDebit := 0
	balanceTotalCredit := 0
	transactionsPerMonth := make([][]Transaction, 12)
	for _, transaction := range transactions {
		transactionsPerMonth[transaction.Month] = append(transactionsPerMonth[transaction.Month], transaction)
	}

	for i, monthTransactions := range transactionsPerMonth {

		if len(monthTransactions) == 0 {
			continue
		}

		info := TransactionMonthInfo{}
		totalDebit := 0
		totalCredit := 0
		info.Total = len(monthTransactions)
		month := mapperIntToMonth(i)
		for _, transaction := range monthTransactions {
			balance.TotalBalance += transaction.Amount
			if transaction.Amount < 0 {
				totalDebit += 1
				balanceTotalDebit += 1
				info.DebitAverage += transaction.Amount
				balance.AverageDebitAmount += transaction.Amount
			} else {
				totalCredit += 1
				balanceTotalCredit += 1
				info.CreditAverage += transaction.Amount
				balance.AverageCreditAmount += transaction.Amount
			}
		}
		if totalDebit > 0 {
			info.DebitAverage = info.DebitAverage / float32(totalDebit)
		}
		if totalCredit > 0 {
			info.CreditAverage = info.CreditAverage / float32(totalCredit)
		}
		balance.TransactionsPerMonth[month] = info
	}

	if balanceTotalDebit > 0 {
		balance.AverageDebitAmount = balance.AverageDebitAmount / float32(balanceTotalDebit)
	}

	if balanceTotalCredit > 0 {
		balance.AverageCreditAmount = balance.AverageCreditAmount / float32(balanceTotalCredit)
	}

	return balance
}

func mapperIntToMonth(month int) string {
	switch month {
	case 1:
		return "Jan"
	case 2:
		return "Feb"
	case 3:
		return "Mar"
	case 4:
		return "Apr"
	case 5:
		return "May"
	case 6:
		return "Jun"
	case 7:
		return "Jul"
	case 8:
		return "Aug"
	case 9:
		return "Sep"
	case 10:
		return "Oct"
	case 11:
		return "Nov"
	case 12:
		return "Dec"
	}
	return ""
}
