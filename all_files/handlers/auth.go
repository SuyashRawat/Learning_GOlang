package handlers

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"

	// "os"
	"taskmanage/logger"
	"taskmanage/models"
	"taskmanage/pkg/utils"
	"taskmanage/rabbitmq"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

// Register handler
func Register(c *gin.Context) {

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		log.Printf("Recovered from panic: %v", r)
	// 		Regerror.Println("Recovered from panic:", r)
	// 	}
	// }()

	var user models.User

	// Parse and validate JSON input
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input"))
		return
	}

	log.Printf("Received user data: %+v", user)

	// Create user in the database
	userData, _ := json.Marshal(user)
	err := rabbitmq.Channel.Publish(
		"email_exchange",
		"email_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        userData,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enqueue email"})
		return
	}

	err = models.CreateUser(&user)
	if err != nil {
		log.Printf("Error in Register handler: %v", err)
		logger.Ok.Println("Could not complete registration: ", err.Error())
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to register user"))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse("User registered successfully"))
}

func Login(c *gin.Context) {

	var userlogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&userlogin); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("invalid input"))
		return

	}
	if userlogin.Email == "" || userlogin.Password == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("email and Password required"))
		return
	}

	user, err := models.Getuserbyemail(userlogin.Email)
	if err != nil || !utils.CheckPasswordHash(userlogin.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("invaild email or Password"))
		return
	}

	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("internal server error"))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   token,
	})

}

//register
//object of model
//map the incoming json to a struct which is the object in this case
//create the user in database
//keep logs everywhere to check

//login
//create empty struct and bind json to that struct to take incoming request
//now check if any required fields are empty or not
//then get the user from database using the fields provided
//match the password
//then generate token and send successresponse with token
//keeps logging errors at every step
