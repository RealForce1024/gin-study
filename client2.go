package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
	"time"
)

var redisdb *redis.Client

func init() {
	redisdb = redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
}

func main() {
	for {
		time.Sleep(time.Duration(1+rand.Intn(2)) * time.Second)
		redisdb.RPopLPush("task-queue", "tmp-queue").Val()
		if redisdb.LLen("tmp-queue").Val() != 0 {
			if rand.Intn(10)%3 == 0 {
				taskId := redisdb.RPopLPush("tmp-queue", "task-queue").Val()
				fmt.Printf("处理任务taskId: %s失败\n", taskId)
			} else {
				taskId := redisdb.RPop("tmp-queue").Val()
				fmt.Printf("处理任务taskId: %s成功\n", taskId)
			}
		}

	}
}
