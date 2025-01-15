package repository

import (
	"database/sql"
	"time"

	"transactions-app/model"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepositoryInterface {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) CreateTransaction(fromUserId uint, toUserId uint, amount float64) error {
	currentDate := time.Now()
	stmt, err := r.DB.Prepare("INSERT INTO transactions (from_user_id, to_user_id, amount, date_initiated, status) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(fromUserId, toUserId, amount, currentDate, "done")
	return err
}

func (r *TransactionRepository) UpdateUserBalance(userId uint, amount float64) error {
	stmt, err := r.DB.Prepare("UPDATE users SET amount = amount + $1 WHERE id = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(amount, userId)
	return err
}

func (r *TransactionRepository) SelectTransactionByID(id uint) (model.Transaction, error) {
	var transaction model.Transaction
	err := r.DB.QueryRow("SELECT * FROM transactions WHERE id = $1", id).Scan(&transaction.Id, &transaction.FromUserId, &transaction.ToUserId, &transaction.Amount, &transaction.DateInitiated, &transaction.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Transaction{}, nil
		}
		return model.Transaction{}, err
	}
	return transaction, nil
}
