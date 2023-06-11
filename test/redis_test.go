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
	if err != nil {
		return
	}
	//_, err = client.SAdd(ctx, "img", 1, 2, 3, 4).Result()
	//if err != nil {
	//	return
	//}
	//
	//f, err := client.SIsMember(ctx, "img", 1).Result()
	//if err != nil {
	//	log.Panicln("1", err)
	//
	//}
	//s, _ := client.SMembers(ctx, "img").Result()
	d, err := client.SRem(ctx, "img", 1, 2, 3, 4).Result()
	if err != nil {
		return
	}

	fmt.Println(result, d)

}
