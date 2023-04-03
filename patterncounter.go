package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Set the pattern and starting value of the counter
	pattern := "example_counter:*"
	startValue := 0

	// Get all keys that match the pattern
	keys, err := rdb.Keys(ctx, pattern).Result()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the counter to the starting value
	counter := startValue

	// Loop through each matching key and increment the counter
	for _, key := range keys {
		val, err := rdb.Get(ctx, key).Int()
		if err != nil {
			log.Fatal(err)
		}
		counter += val
	}

	// Print the final value of the counter
	fmt.Printf("The counter for pattern %q is %d\n", pattern, counter)
}

