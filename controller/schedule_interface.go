package controller

import "github.com/gin-gonic/gin"

type ScheduleControllerInterface interface {
	AddToQueue(g *gin.Context)
	RemoveFromQueue(g *gin.Context)
	GetQueue(g *gin.Context)
}
