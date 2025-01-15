package app

import (
	"database/sql"
	"fmt"
	"log"
	"transactions-app/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type App struct {
	DB     *sql.DB
	Routes *gin.Engine
}

func (a *App) CreateConnection() {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", UNAMEDB, PASSDB, HOSTDB, DBNAME)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	a.DB = db
}

func (a *App) CreateRoutes() {
	routes := gin.Default()
	user_controller := controller.NewUserController(a.DB)
	transaction_controller := controller.NewTransactionController(a.DB)

	routes.POST("/user", user_controller.CreateUser)
	routes.GET("/user/:id", user_controller.GetUserByID)
	routes.GET("/users", user_controller.GetUsers)
	routes.POST("/transactions/", transaction_controller.CreateTransaction)

	a.Routes = routes
}

func (a *App) Run() {
	a.Routes.Run(":8080")
}
