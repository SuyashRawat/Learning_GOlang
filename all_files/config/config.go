package config

import (
	// "log"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Postgresdsn string
	Redisaddr   string
}

func Loadconfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	return Config{
		Postgresdsn: os.Getenv("POSTGRES_DSN"),
		Redisaddr:   os.Getenv("REDIS_ADDR"),
	}
}
