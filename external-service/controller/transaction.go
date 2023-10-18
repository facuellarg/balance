package controller

type (
	TransactionController interface {
		ProcessTransaction(
			fileName string,
			to string,
		) error
	}
)
