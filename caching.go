package main

import (
	"context"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"log"
)

var ctx = context.Background()

type User struct {
	ID    int
	Name  string
	Email string
}

func initCache() *cache.Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return cache.New(&cache.Options{
		Redis: rdb,
	})
}

func main() {
	c := initCache()
	ctx := context.Background()
	userListKey := "user_list"

	users := []User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
		{ID: 3, Name: "Alice Johnson", Email: "alice@example.com"},
	}

	if err := c.Set(&cache.Item{
		Key:   userListKey,
		Value: users,
	}); err != nil {
		log.Fatalf("Could not cache user list: %v", err)
	}
	var cachedUsers []User
	if err := c.Get(ctx, userListKey, &cachedUsers); err != nil {
		log.Printf("Could not retrieve user list: %v", err)
	} else {
		log.Printf("Cached Users: %+v", cachedUsers)
	}

	userKey := "user:1"
	user := User{ID: 1, Name: "John Doe", Email: "john@example.com"}

	if err := c.Set(&cache.Item{
		Key:   userKey,
		Value: user,
	}); err != nil {
		log.Fatalf("Could not cache user data: %v", err)
	}

	var cachedUser User
	if err := c.Get(ctx, userKey, &cachedUser); err != nil {
		log.Printf("Could not retrieve user data: %v", err)
	} else {
		log.Printf("Cached User: %+v", cachedUser)
	}

	if err := c.Delete(ctx, userListKey); err != nil {
		log.Printf("Could not delete cached user list: %v", err)
	} else {
		log.Println("deleted cached user list.")
	}

	if err := c.Delete(ctx, userKey); err != nil {
		log.Printf("Could not delete cached user data: %v", err)
	} else {
		log.Println("deleted cached user data.")
	}
}
