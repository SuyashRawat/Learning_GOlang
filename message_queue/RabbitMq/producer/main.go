package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// Define the message structure
type RegistrationData struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

// Function to connect to RabbitMQ
func MQConnect() (*amqp.Connection, *amqp.Channel, error) {
	// Connect to RabbitMQ
	url := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, nil, err
	}

	// Create a channel
	channel, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	// Declare a queue to send messages to
	_, err = channel.QueueDeclare(
		"email_queue", // Queue name
		true,          // Durable (survive restarts)
		false,         // Delete when unused
		false,         // Exclusive (only this connection can use it)
		false,         // No-wait
		nil,           // Arguments
	)
	if err != nil {
		return nil, nil, err
	}

	return conn, channel, nil
}

// Function to publish a message to RabbitMQ
func MQPublish(channel *amqp.Channel, message []byte) error {
	err := channel.Publish(
		"",            // Exchange ("" means default)
		"email_queue", // Routing key (name of the queue)
		false,         // Mandatory
		false,         // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	return err
}

func main() {
	// Connect to RabbitMQ
	conn, channel, err := MQConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	defer channel.Close()

	// Create a sample registration data
	regData := RegistrationData{
		Email:   "user@example.com",
		Message: "Welcome to our CMS!",
	}

	// Marshal the data into JSON
	message, err := json.Marshal(regData)
	if err != nil {
		log.Fatal(err)
	}

	// Send the message to RabbitMQ
	err = MQPublish(channel, message)
	if err != nil {
		log.Fatal(err)
	}

	// Log successful message sending
	fmt.Println("Message sent:", string(message))
}
