package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/satori/go.uuid"
	"math/rand"
	"time"
)

const (
	HOST         = "localhost"
	PORT         = ":6379"
	CONNECT_TYPE = "tcp"
)

func main() {
	c, err := redis.Dial(CONNECT_TYPE, HOST+PORT)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	fmt.Println("===>connect ok!")
	for {
		taskId := uuid.Must(uuid.NewV4(), nil)
		r := rand.Intn(3)
		time.Sleep(time.Duration(r) * time.Second)
		_, err = c.Do("lpush", "task-queue", taskId)
		fmt.Println("生成taskId: ", taskId)
	}
}

//总共用时: 0.911276685
//10000次,总共用时0.911276685
//QPS:=>10973.615549046994

//总共用时: 10.118333802
//100000次,总共用时10.118333802
//QPS:=>9883.050110506721
//func main() {
//	c, err := redis.Dial(CONNECT_TYPE, HOST+PORT)
//	if err != nil {
//		fmt.Println("Connect to redis error", err)
//		return
//	}
//	defer c.Close()
//	fmt.Println("===>connect ok!")
//
//	start := time.Now()
//	var count = 10000
//	for i := 1; i <= count; i++ {
//		in := strconv.Itoa(i)
//		// 写入值永不过期
//		_, err = c.Do("SET", "username"+in, "kobe"+in)
//		if err != nil {
//			fmt.Println("redis set failed:", err)
//		}
//		username, err := redis.String(c.Do("GET", "username"+in))
//		if err != nil {
//			fmt.Println("redis get failed:", err)
//		} else {
//			fmt.Printf("Got username%v, %v \n", in, username)
//		}
//	}
//
//	end := time.Since(start).Seconds()
//	fmt.Println("总共用时:", end)
//	fmt.Printf("%v次,总共用时%v\n", count, end)
//	fmt.Printf("QPS:=>%v\n", float64(count)/end)
//}
