package entities

type Transaction struct {
	ID     int     `json:"id"`
	Amount float32 `json:"amount"`
	Month  int     `json:"month"`
	Day    int     `json:"day"`
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
