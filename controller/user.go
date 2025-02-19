package controller

import (
	"database/sql"
	"net/http"

	"transactions-app/model"
	"transactions-app/repository"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	DB *sql.DB
}

func InstanceUserController(db *sql.DB) UserControllerInterface {
	return &UserController{DB: db}
}

func (c *UserController) GetUsers(g *gin.Context) {
	db := c.DB
	repoUser := repository.InstanteUserRepository(db)
	get_user := repoUser.SelectUsers()
	if get_user == nil {
		g.JSON(http.StatusInternalServerError, gin.H{"status": "failed"})
		return
	}
	g.JSON(http.StatusOK, gin.H{"status": "success", "data": get_user})
}

func (c *UserController) GetUserByPhone(g *gin.Context) {
	db := c.DB
	repoUser := repository.InstanteUserRepository(db)

	phone := g.Param("phone")
	user, err := repoUser.SelectUserByPhone(phone)
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
		repoUser := repository.InstanteUserRepository(db)
		insert := repoUser.CreateUser(post)
		if insert {
			g.JSON(http.StatusOK, gin.H{"status": "success", "msg": "user created successfully"})
		} else {
			g.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "msg": "user not created"})
		}
	} else {
		g.JSON(http.StatusBadRequest, gin.H{"status": "failed", "msg": "invalid request", "error": err})
	}
}
