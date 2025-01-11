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

	// Declare the same queue that the producer sends messages to
	_, err = channel.QueueDeclare(
		"email_queue", // Queue name
		true,          // Durable
		false,         // Delete when unused
		false,         // Exclusive
		false,         // No-wait
		nil,           // Arguments
	)
	if err != nil {
		return nil, nil, err
	}

	return conn, channel, nil
}

// Function to consume messages from RabbitMQ
func MQConsume(channel *amqp.Channel) error {
	// Start consuming messages from the queue
	msgs, err := channel.Consume(
		"email_queue", // Queue name
		"",            // Consumer name (empty means random)
		true,          // Auto-acknowledge
		false,         // Exclusive
		false,         // No-local
		false,         // No-wait
		nil,           // Arguments
	)
	if err != nil {
		return err
	}

	// Loop over messages and process them
	for msg := range msgs {
		var regData RegistrationData
		err := json.Unmarshal(msg.Body, &regData)
		if err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}

		// Log received message
		fmt.Println("Received message:")
		fmt.Printf("Email: %s\n", regData.Email)
		fmt.Printf("Message: %s\n", regData.Message)
	}

	return nil
}

func main() {
	// Connect to RabbitMQ
	conn, channel, err := MQConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	defer channel.Close()

	// Start consuming messages
	err = MQConsume(channel)
	if err != nil {
		log.Fatal(err)
	}
}
