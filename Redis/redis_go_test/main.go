package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password for local Docker Redis
		DB:       0,  // default database
	})

	// 1. Test connection
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Redis connection failed:", err)
	}
	fmt.Println("Redis Ping:", pong)

	// 2. SET data
	err = rdb.Set(ctx, "name", "Siam", 0).Err()
	if err != nil {
		log.Fatal("SET failed:", err)
	}
	fmt.Println("Data saved successfully")

	// 3. GET data
	name, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		log.Fatal("GET failed:", err)
	}
	fmt.Println("Name from Redis:", name)

	// 4. SET data with expiry
	err = rdb.Set(ctx, "otp", "123456", 10*time.Second).Err()
	if err != nil {
		log.Fatal("OTP SET failed:", err)
	}
	fmt.Println("OTP saved for 10 seconds")

	otp, err := rdb.Get(ctx, "otp").Result()
	if err != nil {
		log.Fatal("OTP GET failed:", err)
	}
	fmt.Println("OTP from Redis:", otp)

	// 5. DELETE data
	err = rdb.Del(ctx, "name").Err()
	if err != nil {
		log.Fatal("DELETE failed:", err)
	}
	fmt.Println("Name deleted successfully")
}