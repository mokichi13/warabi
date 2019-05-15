package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Echo redis.")
	cli := ExampleNewClient()
	ExampleClient(cli)
}

// ExampleNewClient example create new client.
func ExampleNewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return client
}

// ExampleClient example use client.
func ExampleClient(client *redis.Client) {
	// value set.
	err := client.Set("key", "valueeeeeee", 0).Err()
	if err != nil {
		panic(err)
	}

	// value get.
	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
