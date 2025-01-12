package main

import (
	// "context"
	// "taskflow/config"
	// "taskflow/db"
	// "fmt"
	"log"
	"taskmanage/consumers"
	"taskmanage/db"

	// "taskmanage/logger"

	// "taskmanage/handlers"
	// "taskmanage/logger"

	"taskmanage/models"

	// "taskmanage/rabbitmq"

	// "taskmanage/models"
	"taskmanage/routes"

	// "taskmanage/pkg/cache"

	// "taskflow/pkg/aws"
	// "taskflow/pkg/cache"
	// "taskflow/services"
	"taskmanage/config"

	"github.com/gin-gonic/gin"
)

// var ctx string=context.Background()

func main() {

	cfg := config.Loadconfig()
	// logger.Ok,_=logger.GetLogger("Service1")
	// handlers.Setok(ok)
	// ctx:=context.Background()
	// redisClient:=cache.ConnectRedis(ctx,cfg.Redisaddr)

	postconnection := db.ConnectPostgres(cfg.Postgresdsn)

	// Retrieve the underlying sql.DB object
	sqlDB, err := postconnection.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB: %v", err)
	}
	// Ping the database
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	} else {
		log.Println("Database connection successful")
	}
	models.SetDB(postconnection)

	go consumers.StartEmailConsumer()

	// Add this line
	router := gin.Default()
	routes.SetupRoutes(router)

	port := ":8080"
	router.Run(port)

}
