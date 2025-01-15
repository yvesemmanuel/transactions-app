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
		transaction_repo := repository.NewTransactionRepository(db)

		err := transaction_repo.CreateTransaction(post.FromUserId, post.ToUserId, post.Amount)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"status": "success", "msg": "failed to create transaction"})
			return
		}

		err = transaction_repo.UpdateUserBalance(post.FromUserId, -post.Amount)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": "failed to update sender's balance"})
			return
		}

		err = transaction_repo.UpdateUserBalance(post.ToUserId, post.Amount)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": "failed to update receiver's balance"})
			return
		}

		g.JSON(http.StatusOK, gin.H{"status": "success", "message": "transaction completed successfully"})

	} else {
		g.JSON(http.StatusBadRequest, gin.H{"status": "failed", "msg": "invalid request", "error": err})
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
