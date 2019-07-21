package main

import (
	"fmt"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"time"
)

var r *redis.Client

func init() {
	r = redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
}
func main() {
	fmt.Println(r.Ping().Val())
	for {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		random := uuid.Must(uuid.NewV4(), nil).String()
		fmt.Printf("生成taskId:  %v\n", random)
		fmt.Println("生成taskId", random)
		r.LPush("task-queue", random)
		//taskId, _ := task.Result()
	}

}
