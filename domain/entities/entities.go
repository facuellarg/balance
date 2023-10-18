package entities

type (
	Transaction struct {
		ID     int     `json:"id"`
		Amount float32 `json:"amount"`
		Month  int     `json:"month"`
		Day    int     `json:"day"`
	}

	TransactionMonthInfo struct {
		Total         int
		CreditAverage float32
		DebitAverage  float32
	}

	Balance struct {
		TotalBalance         float32
		TransactionsPerMonth map[string]TransactionMonthInfo
		AverageDebitAmount   float32
		AverageCreditAmount  float32
	}

	Transactions []Transaction
)

func (t Transactions) GetBalance() Balance {

	balance := Balance{
		TotalBalance:         0,
		TransactionsPerMonth: make(map[string]TransactionMonthInfo),
		AverageDebitAmount:   0,
		AverageCreditAmount:  0,
	}

	balanceTotalDebit := 0
	balanceTotalCredit := 0
	transactionsPerMonth := make([][]Transaction, 12)
	for _, transaction := range t {
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
