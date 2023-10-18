package controller

type (
	TransactionController interface {
		ProcessTransaction(
			fileName string,
			bucket string,
			to string,
		) error
	}
)
