package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	redisStreamKey = "exampleStream"
	groupName      = "exampleGroup"
	consumerName   = "exampleConsumer"
)

var ctx = context.Background() //建立一個一個空的背景上下文，可以做平行處理時的溝通之用

func main() {
	//配置連結至資料庫所需的參數
	options := redis.Options{
		Addr:     "localhost:6379", //資料庫所在位址
		Password: "",               // 密碼
		DB:       0,                // 使用的database，0代表預設的database
	}

	//建立程式與資料庫的連結，rdb便是連結的橋樑
	rdb := redis.NewClient(&options)

	// 创建消费者组
	if err := createConsumerGroup(ctx, rdb, redisStreamKey, groupName); err != nil {
		log.Fatalf("Error creating consumer group: %v\n", err)
	}

	// 启动消费者
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go consumeMessages(ctx, rdb, redisStreamKey, groupName, consumerName, wg)

	// 启动生产者
	go produceMessages(ctx, rdb, redisStreamKey)

	wg.Wait()
}

// 创建消费者组
func createConsumerGroup(ctx context.Context, rdb *redis.Client, streamKey, groupName string) error {
	_, err := rdb.XGroupCreate(ctx, streamKey, groupName, "$").Result()
	return err
}

// 消费消息
func consumeMessages(ctx context.Context, rdb *redis.Client, streamKey, groupName, consumerName string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		// 读取消息
		streams, err := rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    groupName,
			Consumer: consumerName,
			Streams:  []string{streamKey, ">"},
			Block:    0,  // 永久阻塞直到有新消息
			Count:    10, // 每次读取的最大消息数
			NoAck:    false,
		}).Result()
		if err != nil {
			log.Printf("Error reading from stream: %v\n", err)
			continue
		}

		// 处理消息
		for _, stream := range streams {
			for _, message := range stream.Messages {
				// 这里可以根据需要对消息进行处理
				fmt.Printf("Received message with ID %s: %v\n", message.ID, message.Values)
			}
		}
	}
}

// 生产消息
func produceMessages(ctx context.Context, rdb *redis.Client, streamKey string) {
	for i := 1; ; i++ {
		// 生产消息
		message := map[string]interface{}{
			"index":   i,
			"message": fmt.Sprintf("Message %d", i),
			"time":    time.Now().Unix(),
		}

		// 将消息写入流
		_, err := rdb.XAdd(ctx, &redis.XAddArgs{
			Stream: streamKey,
			Values: message,
		}).Result()
		if err != nil {
			log.Printf("Error producing message: %v\n", err)
			continue
		}

		fmt.Printf("Produced message %d\n", i)

		// 模拟产生消息的间隔
		time.Sleep(1 * time.Second)
	}
}
