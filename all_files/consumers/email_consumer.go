package consumers

import (
    "encoding/json"
    "log"
    // "net/smtp"
    "taskmanage/rabbitmq"
)

type User struct {
    Username string `json:"username"`
    Email    string `json:"email"`
}

func StartEmailConsumer() {
    rabbitmq.InitRabbitMQ()
    defer rabbitmq.Conn.Close()
    defer rabbitmq.Channel.Close()

    q, err := rabbitmq.Channel.QueueDeclare(
        "email_queue",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        log.Fatalf("Failed to declare a queue: %s", err)
    }

    err = rabbitmq.Channel.QueueBind(
        q.Name,
        "email_queue",
        "email_exchange",
        false,
        nil,
    )
    if err != nil {
        log.Fatalf("Failed to bind queue: %s", err)
    }

    msgs, err := rabbitmq.Channel.Consume(
        q.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        log.Fatalf("Failed to register a consumer: %s", err)
    }

    go func() {
        for d := range msgs {
            var user User
            json.Unmarshal(d.Body, &user)
            // sendWelcomeEmail(user)
			if err != nil {
				log.Printf("Error unmarshalling user data: %s", err)
				continue
			}
			// Print the user details to the console
			log.Printf("Received user data: %+v", user)
		
        }
    }()

    log.Printf("Waiting for messages. To exit press CTRL+C")
    select {}
}

