package controller

import "github.com/gin-gonic/gin"

type UserControllerInterface interface {
	GetUsers(g *gin.Context)
	GetUserByID(g *gin.Context)
	CreateUser(g *gin.Context)
}
