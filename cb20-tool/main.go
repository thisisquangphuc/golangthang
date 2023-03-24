package main

import (
	"context"
	"encoding/json"
	"fmt"

	Core "cb20-tool/core"

	"github.com/redis/go-redis/v9"
	"github.com/valyala/fastjson"
)

var ctx = context.Background()

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:32769",
		Password: "redispw", // no password set
		DB:       0,         // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
	sub := rdb.Subscribe(ctx, "hello")

	user := User{}

	for {

		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		if err := json.Valid([]byte(msg.Payload)); err != true {
			fmt.Println("Invalid JSON")
			err := fastjson.Validate(msg.Payload)
			fmt.Println(err)
		}
		if err := json.Unmarshal([]byte(msg.Payload), &user); err != nil {
			panic(err)
		}
		fmt.Println("Received message from " + msg.Channel + " channel, msg:" + msg.Payload)
		fmt.Printf("%+v\n", user)
		// ...
	}
}

func main() {
	Core.InitREDIS()
	Core.InitLogrus()

	fmt.Println("Hello, world!")
	fmt.Println("Hello 123")
}
