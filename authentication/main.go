package main

import (
	"JWT/controllers"
	"JWT/database"
	"JWT/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	controllers.SetDatabase(database.DB)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	port := ":8090"
	if err := router.Run(port); err != nil {
		panic("Failed to start the server: " + err.Error())
	}
}
