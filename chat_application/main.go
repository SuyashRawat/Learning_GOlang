package main

import (
	"chat_application/models"
	"chat_application/trace"
<<<<<<< HEAD
	"os"
	// "chat_application/rooms"
=======
>>>>>>> 1f07827734644fd605fe5b3740ab0d4a63861888
	"log"
	"os"

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
	if err := router.Run(":8070"); err != nil {
		log.Fatal("Gin server failed to start:", err)
	}

}
