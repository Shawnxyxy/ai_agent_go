package database

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client
var Ctx = context.Background()

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:		"127.0.0.1:6379",
		Password: 	"",
		DB:			0,
	})
	// 连接redis
	_, err := RDB.Ping(Ctx).Result()
	if err != nil {
		panic("failed to connect redis:" + err.Error())
	}
	fmt.Println("Redis connected successfully")
}