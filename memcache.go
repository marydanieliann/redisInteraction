package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	fmt.Println("Without caching...")
	start := time.Now()
	getDataExpensive()
	elapsed := time.Since(start)
	fmt.Printf("Without caching took %s\n\n", elapsed)

	fmt.Println("With caching...")
	start = time.Now()
	getDataCached()
	elapsed = time.Since(start)
	fmt.Printf("With caching took %s\n", elapsed)
}

func getDataExpensive() {
	for i := 0; i < 3; i++ {
		fmt.Println("\tBefore query")
		result := databaseQuery()
		fmt.Printf("\tAfter query with result %s\n", result)
	}
}

func getDataCached() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	defer rdb.Close()

	for i := 0; i < 3; i++ {
		fmt.Println("\tBefore query")
		val, err := rdb.Get(ctx, "query").Result()
		if err != nil {
			val = databaseQuery()
			rdb.Set(ctx, "query", val, 0)
		}
		fmt.Printf("\tAfter query with result %s\n", val)
	}
}

func databaseQuery() string {
	fmt.Println("\tDatabase queried")
	time.Sleep(5 * time.Second)
	return "bar"
}
