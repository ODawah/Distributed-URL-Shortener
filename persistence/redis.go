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

func RedisGet(id string) *models.URL {

	val, _ := RedisClient.Get(context.Background(), id).Result()
	if val == "" {
		return nil
	}

	return &models.URL{
		URL: val,
		ID:  id,
	}
}
