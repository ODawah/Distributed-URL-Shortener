// persistence/redis.go
package persistence

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ODawah/Distributed-URL-Shortener/models"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

const MonthlyExpiration = (24 * time.Hour) * 30

func ConnectToRedis() error {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	fmt.Println("Connected to Redis: ", pong)
	RedisClient = client
	return nil
}

func RedisInsert(data *models.URL) error {
	err := RedisClient.Set(context.Background(), data.ID, data.URL, MonthlyExpiration).Err()
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to set key: %v", err))
	}

	return nil
}

func RedisGet(data *models.URL) error {
	fmt.Println("Looking up Redis for ID:", data.ID) // Debugging statement

	val, err := RedisClient.Get(context.Background(), data.ID).Result()
	if err != nil {
		fmt.Println("Error fetching from Redis:", err) // Debugging statement
		return errors.New(fmt.Sprintf("Failed to get key: %v", err))
	}

	fmt.Println("Found in Redis, URL:", val) // Debugging statement
	data.URL = val
	return nil
}
