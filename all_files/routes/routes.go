package routes

import (
	"taskmanage/handlers"
	"taskmanage/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine){

	api:=router.Group("/api")
	api.POST("/register",handlers.Register)
	api.POST("/login",handlers.Login)

	auth:=api.Group("/")
	auth.Use(middleware.AuthMiddleware())
	router.POST("/projects", handlers.CreateProject)
	router.GET("/projects", handlers.GetProjects)
	router.GET("/projects/:id", handlers.GetProject)

	router.POST("/tasks", handlers.CreateTask)
	router.GET("/tasks", handlers.GetTasks)
	router.GET("/tasks/:id", handlers.GetTask)
	router.PUT("/tasks/:id", handlers.UpdateTask)





}