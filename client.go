package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"time"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	fmt.Println("===>connect ok!")
	for {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		taskId, _ := c.Do("rpoplpush", "task-queue", "tmp-queue")

		tmp_queue_len, _ := c.Do("llen", "tmp-queue");
		task_queue_len, _ := c.Do("llen", "task-queue");
		fmt.Printf("待处理任务数: %v\n", tmp_queue_len)
		fmt.Printf("总任务队列: %v\n", task_queue_len)
		if tmp_queue_len != int64(0) {
			if rand.Intn(10)%3 == 0 {
				//time.Sleep(time.Second + 2) //模拟处理失败的情况,宕机后并未完成下面的操作,失败的也停留在tmp-queue中
				taskId, _ = c.Do("rpoplpush", "tmp-queue", "task-queue")
				fmt.Printf("taskId处理失败: %s\n", taskId)
			} else {
				taskId, _ = c.Do("rpop", "tmp-queue")
				fmt.Printf("taskId处理成功: %s\n", taskId)
			}
		}

	}
}
