package repository

import "transactions-app/model"

type TransactionRepositoryInterface interface {
	CreateTransaction(fromUserId uint, toUserId uint, amount float64) error
	UpdateUserBalance(userId uint, amount float64) error
	SelectTransactionByID(id uint) (model.Transaction, error)
	CountTransactionsLastHour(userId uint) (int, error)
}
