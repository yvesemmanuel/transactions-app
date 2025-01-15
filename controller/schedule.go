package controller

import (
	"database/sql"
	"net/http"
	"transactions-app/model"
	"transactions-app/repository"

	"github.com/gin-gonic/gin"
)

type ScheduleController struct {
	DB *sql.DB
}

func InstanceScheduleController(db *sql.DB) ScheduleControllerInterface {
	return &ScheduleController{DB: db}
}

func (c *ScheduleController) AddToQueue(g *gin.Context) {
	db := c.DB
	repoSchedule := repository.InstanteScheduleRepository(db)

	var post model.PostSchedule
	if err := g.ShouldBindJSON(&post); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err.Error()})
		return
	}

	err := repoSchedule.AddToQueue(post.Phone, post.Priority)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": "failed to add user to queue"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"status": "success", "message": "user added to queue"})
}

func (c *ScheduleController) RemoveFromQueue(g *gin.Context) {
	db := c.DB
	repoSchedule := repository.InstanteScheduleRepository(db)
	schedule, err := repoSchedule.RemoveFromQueue()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": "failed to remove user from queue"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"status": "success", "data": schedule})
}

func (c *ScheduleController) GetQueue(g *gin.Context) {
	db := c.DB
	repoSchedule := repository.InstanteScheduleRepository(db)
	schedules, err := repoSchedule.GetQueue()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": "failed to fetch queue"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"status": "success", "data": schedules})
}
