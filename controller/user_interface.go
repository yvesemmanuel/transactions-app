package controller

import "github.com/gin-gonic/gin"

type UserControllerInterface interface {
	GetUsers(g *gin.Context)
	GetUserByPhone(g *gin.Context)
	CreateUser(g *gin.Context)
}
