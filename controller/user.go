package controller

import (
	"database/sql"
	"net/http"

	"fintech-app/model"
	"fintech-app/repository"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	DB *sql.DB
}

func NewUserController(db *sql.DB) UserControllerInterface {
	return &UserController{DB: db}
}

func (m *UserController) GetUsers(g *gin.Context) {
	db := m.DB
	repo_user := repository.NewUserRepository(db)
	get_user := repo_user.SelectUsers()
	g.JSON(http.StatusOK, gin.H{"status": "success", "data": get_user, "msg": "get user successfully"})
}

func (m *UserController) GetUserByID(g *gin.Context) {
	db := m.DB
	id := g.Param("id")
	repo_user := repository.NewUserRepository(db)
	user, err := repo_user.SelectUserByID(id)
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"status": "failed", "msg": "user not found"})
		return
	}
	g.JSON(http.StatusOK, user)
}

func (m *UserController) CreateUser(g *gin.Context) {
	db := m.DB
	var post model.PostUser
	if err := g.ShouldBindJSON(&post); err == nil {
		repo_user := repository.NewUserRepository(db)
		insert := repo_user.CreateUser(post)
		if insert {
			g.JSON(http.StatusOK, gin.H{"status": "success"})
		} else {
			g.JSON(http.StatusInternalServerError, gin.H{"status": "failed"})
		}
	} else {
		g.JSON(http.StatusBadRequest, gin.H{"status": "failed", "msg": err})
	}
}
