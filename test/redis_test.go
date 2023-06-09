package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"testing"
)

func Test_redis(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "116.198.44.154:6379",
		Password: "admin123", // no password set
		DB:       0,          // uses default DB
		PoolSize: 1000,
	})
	ctx := context.Background()
	ping := client.Ping(ctx)
	if ping.String() == "ping: PONG" {
		//fmt.Println(ping.String())
		log.Println("连接redis 成功!")
	}
	result, err := client.Keys(ctx, "*").Result()
	r2, _ := client.Get(ctx, "e5a6071b-baaf-45aa-a587-784d0ff9a575").Result()
	r1, err := client.Exists(ctx, "e5a6071b-baaf-45aa-a587-784d0ff9a575").Result()
	if err != nil {
		return
	}
	if err != nil {
		return
	}

	fmt.Println(result, r1, r2)
}
