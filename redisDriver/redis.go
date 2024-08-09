package redisdriver

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func ConnectRedis() *redis.Client {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Replace with your Redis server address
		Password: "",               // No password for local development
		DB:       0,                // Default DB
	})

	// Ping the Redis server to check the connection
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	return client
}
