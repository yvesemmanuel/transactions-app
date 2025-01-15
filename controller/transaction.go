package controller

import (
	"database/sql"
	"net/http"
	"strconv"
	"transactions-app/model"
	"transactions-app/repository"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	DB *sql.DB
}

func NewTransactionController(db *sql.DB) TransactionControllerInterface {
	return &TransactionController{DB: db}
}

func (c *TransactionController) CreateTransaction(g *gin.Context) {
	db := c.DB
	var post model.PostTransaction
	if err := g.ShouldBindJSON(&post); err == nil {
		userRepo := repository.NewUserRepository(db)
		transactionRepo := repository.NewTransactionRepository(db)

		// Fraud detection: Check recent transactions
		maxTransactionsCount := 5
		transactionCount, err := transactionRepo.CountTransactionsLastHour(post.FromUserId)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": "failed to check transaction history"})
			return
		}

		if transactionCount >= maxTransactionsCount {
			g.JSON(http.StatusForbidden, gin.H{
				"status":  "fraud",
				"message": "transaction blocked due to suspected fraud",
			})
			return
		}

		sender, err := userRepo.SelectUserByID(post.FromUserId)
		if err != nil {
			g.JSON(http.StatusNotFound, gin.H{"status": "failed", "msg": "sender not found"})
			return
		}

		if sender.Amount < post.Amount {
			g.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": "insufficient balance"})
			return
		}

		err = transactionRepo.CreateTransaction(post.FromUserId, post.ToUserId, post.Amount)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": "failed to create transaction"})
			return
		}

		err = transactionRepo.UpdateUserBalance(post.FromUserId, -post.Amount)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": "failed to update sender's balance"})
			return
		}

		err = transactionRepo.UpdateUserBalance(post.ToUserId, post.Amount)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": "failed to update receiver's balance"})
			return
		}

		g.JSON(http.StatusOK, gin.H{"status": "success", "message": "transaction completed successfully"})

	} else {
		g.JSON(http.StatusBadRequest, gin.H{"status": "failed", "msg": "invalid request", "error": err.Error()})
	}
}

func (c *TransactionController) GetTransaction(g *gin.Context) {
	db := c.DB
	transaction_repo := repository.NewTransactionRepository(db)

	idStr := g.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": err})
		return
	}
	transaction, err := transaction_repo.SelectTransactionByID(uint(id))
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "msg": "failed to retrieve transaction"})
		return
	}
	g.JSON(http.StatusOK, gin.H{"status": "success", "data": transaction})
}
