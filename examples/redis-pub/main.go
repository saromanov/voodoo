package main

import (
	"log"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{})
	err := client.Publish("test", "data").Err()
	if err != nil {
		log.Printf("redis failed with: %s", err)
	}
}
