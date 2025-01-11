package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

var (
	Conn    *amqp.Connection
	Channel *amqp.Channel
)

func InitRabbitMQ() {
	var err error
	Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	Channel, err = Conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}

	// Declare exchange
	err = Channel.ExchangeDeclare(
		"email_exchange", // name
		"direct",         // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %s", err)
	}
}
