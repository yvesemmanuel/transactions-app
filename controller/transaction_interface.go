package controller

import "github.com/gin-gonic/gin"

type TransactionControllerInterface interface {
	CreateTransaction(g *gin.Context)
	GetTransaction(g *gin.Context)
}
