package core_lib

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// I am VISIBLE OUTSIDE the FILE because I'm in capital case
func InitREDIS() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:32769",
		Password: "redispw", // no password set
		DB:       0,         // use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(pong)
	}
}
