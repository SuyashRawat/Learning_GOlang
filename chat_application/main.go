package main

import (
	"chat_application/models"
	"chat_application/trace"
	"os"
	// "chat_application/rooms"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := models.NewRoom()
	r.Tracer = trace.New(os.Stdout)
	router := gin.Default()

	router.StaticFile("/", "templates/chat.html")
	router.GET("/room", func(ctx *gin.Context) {
		r.ServeNewRoom(ctx)
	})
	go r.Run()
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Gin server failed to start:", err)
	}

}
