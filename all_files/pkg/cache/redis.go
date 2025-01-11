package cache

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)
// var ctx = context.Background()
func ConnectRedis(ctx context.Context,addr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	return client
}


