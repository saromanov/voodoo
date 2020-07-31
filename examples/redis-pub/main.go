package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{})
	for i := 0; i < 100; i++ {
		go func(d int) {
			err := client.Publish("test", fmt.Sprintf("%ddata", d)).Err()
			if err != nil {
				log.Printf("redis failed with: %s", err)
			}
		}(i)

		go func(d int) {
			err := client.Publish("test2", fmt.Sprintf("%ddataNIL", d)).Err()
			if err != nil {
				log.Printf("redis failed with: %s", err)
			}
		}(i)
	}
	time.Sleep(3 * time.Second)
}
