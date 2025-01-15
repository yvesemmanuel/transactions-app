package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"transactions-app/model"
	"transactions-app/repository"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	DB *sql.DB
}

func NewUserController(db *sql.DB) UserControllerInterface {
	return &UserController{DB: db}
}

func (c *UserController) GetUsers(g *gin.Context) {
	db := c.DB
	repo_user := repository.NewUserRepository(db)
	get_user := repo_user.SelectUsers()
	if get_user == nil {
		g.JSON(http.StatusInternalServerError, gin.H{"status": "failed"})
		return
	}
	g.JSON(http.StatusOK, gin.H{"status": "success", "data": get_user})
}

func (c *UserController) GetUserByID(g *gin.Context) {
	db := c.DB
	repo_user := repository.NewUserRepository(db)

	idStr := g.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": err})
		return
	}
	user, err := repo_user.SelectUserByID(uint(id))
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"status": "failed", "msg": "user not found"})
		return
	}
	g.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

func (c *UserController) CreateUser(g *gin.Context) {
	db := c.DB
	var post model.PostUser
	if err := g.ShouldBindJSON(&post); err == nil {
		repo_user := repository.NewUserRepository(db)
		insert := repo_user.CreateUser(post)
		if insert {
			g.JSON(http.StatusOK, gin.H{"status": "success", "msg": "user created successfully"})
		} else {
			g.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "msg": "user not created"})
		}
	} else {
		g.JSON(http.StatusBadRequest, gin.H{"status": "failed", "msg": "invalid request", "error": err})
	}
}
