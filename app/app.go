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
	userController := controller.InstanceUserController(a.DB)
	scheduleController := controller.InstanceScheduleController(a.DB)

	routes.POST("/user", userController.CreateUser)
	routes.GET("/user/:id", userController.GetUserByPhone)
	routes.GET("/users", userController.GetUsers)

	routes.POST("/schedule", scheduleController.AddToQueue)
	routes.DELETE("/schedule", scheduleController.RemoveFromQueue)
	routes.GET("/schedule", scheduleController.GetQueue)

	a.Routes = routes
}

func (a *App) Run() {
	a.Routes.Run(":8080")
}
