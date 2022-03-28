package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var REDIS *redis.Client

func InitRedis() {
	REDIS = redis_connect("default")
}

func CloseRedis() {
	REDIS.Close()
}

func redis_connect(project string) *redis.Client {
	addr := fmt.Sprintf("%s:%s", "144.123.173.6", "6379")
	//option := redis.DialPassword().String())
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "qiangbi123",
		DB:       0, // use default DB
	})
	log.Println("[GIN-Redis(" + project + ")] connected success")
	return c
}
