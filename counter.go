package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
)

// RedisPool is a Redis connection pool.
var RedisPool *redis.Pool

func main() {
	// Initialize Redis connection pool.
	RedisPool = &redis.Pool{
		MaxIdle:   5,
		MaxActive: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	// Set up HTTP endpoint for counter.
	http.HandleFunc("/counter", counterHandler)

	// Start HTTP server.
	log.Fatal(http.ListenAndServe(":18080", nil))
}

// counterHandler increments and returns the current value of the counter.
func counterHandler(w http.ResponseWriter, r *http.Request) {
	// Get a Redis connection from the pool.
	conn := RedisPool.Get()
	defer conn.Close()

	// Increment counter in Redis.
	_, err := conn.Do("INCR", "counter")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get current value of counter from Redis.
	count, err := redis.Int(conn.Do("GET", "counter"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write current value of counter to HTTP response.
	fmt.Fprintf(w, "Counter: %d", count)
}

