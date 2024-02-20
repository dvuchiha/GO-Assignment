package main

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

// Define a global Redis client
var client *redis.Client
var key string
var cacheExpirationSecondsStr string

// Initialize the Redis client once only
func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1, // database index 1 (temp_db)
	})
	key = "price_data"
	err := godotenv.Load()
	if err == nil {
		cacheExpirationSecondsStr = os.Getenv("CACHE_EXPIRATION_SECONDS")
	}
}

func retrieve() (interface{}, error) {
	ctx := context.Background() // used for signalling mechanisms/cancelling-processes

	exists, err := client.Exists(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if exists == 0 {
		return nil, nil
	} else {
		val, err := client.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		var result interface{}
		if err := json.Unmarshal([]byte(val), &result); err != nil {
			return nil, err
		}

		// fmt.Printf("Cache-hit. Value - %v", result)
		return result, nil
	}
}

func store(val interface{}) error {
	ctx := context.Background()
	jsonVal, err := json.Marshal(val)
	if err != nil {
		return err
	}
	cacheExpirationSeconds, err := strconv.Atoi(cacheExpirationSecondsStr)
	if err != nil {
		cacheExpirationSeconds = 0 // Default, no caching
	}
	err = client.Set(ctx, key, jsonVal, time.Duration(cacheExpirationSeconds)*time.Second).Err()
	if err != nil {
		return err
	}
	// fmt.Println("Cache-miss")
	return nil
}
