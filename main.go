package main

import (
	"transactions-app/app"
)

func main() {
	var app app.App
	app.CreateConnection()
	app.CreateRoutes()
	app.Run()
}
