package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func PopItem() []string {
	client := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})

	val, err := client.RPopCount(ctx, "asin", 2).Result()
	if err != nil {
		// There isn't enough data
		return make([]string, 0)
	}
	return val
}