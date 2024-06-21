package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

// log template

type MyFormatter struct{}

func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")

	var logLevel string
	switch entry.Level {
	case logrus.DebugLevel:
		logLevel = "\033[1;35mDEBUG\033[0m"
	case logrus.InfoLevel:
		logLevel = "\033[1;32mINFO\033[0m"
	case logrus.WarnLevel:
		logLevel = "\033[1;33mWARN\033[0m"
	case logrus.ErrorLevel:
		logLevel = "\033[1;31mERROR\033[0m"
	case logrus.FatalLevel:
		logLevel = "\033[1;31mFATAL\033[0m"
	case logrus.PanicLevel:
		logLevel = "\033[1;31mPANIC\033[0m"
	default:
		logLevel = fmt.Sprintf("[%s]", entry.Level)
	}

	var newLog string

	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("[%s][%s][%s:%d] %s\n",
			logLevel, timestamp, fName, entry.Caller.Line, entry.Message)
	} else {
		newLog = fmt.Sprintf("[%s][%s] %s\n", logLevel, timestamp, entry.Message)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}

func initLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&MyFormatter{})
	logger.SetOutput(os.Stderr)
	logger.SetReportCaller(true)

	return logger
}

// log template

var ctx = context.Background() // context for redis

func PublishingMessage(rdb *redis.ClusterClient, log *logrus.Logger, i int) {
	//Publishing message to stream
	_, err := rdb.XAdd(ctx, &redis.XAddArgs{ // add message to a new stream, if stream not exist, create a new one
		Stream: os.Getenv("STREAM_NAME"),
		Values: map[string]interface{}{
			"message": fmt.Sprintf("Message ID: %d", i),
		},
	}).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Send Message: \"Message ID: %d\"", i)
}

func Producer(log *logrus.Logger) {
	//parameters for connecting to redis cluster
	options := redis.ClusterOptions{
		Addrs:    []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	//connect to redis cluster
	rdb := redis.NewClusterClient(&options)

	for i := 0; i < 1000; i++ {
		PublishingMessage(rdb, log, i)
	}
}

func AutoClaim(log *logrus.Logger) {
	//parameters for connecting to redis cluster
	options := redis.ClusterOptions{
		Addrs:    []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	//connect to redis cluster
	rdb := redis.NewClusterClient(&options)

	//Creating a consumer group, if exists, we can ignore the error
	_, err := rdb.XGroupCreateMkStream(ctx, os.Getenv("STREAM_NAME"), os.Getenv("CUSTOMER_GROUPNAME"), "$").Result()
	if err != nil {
		log.Error(err)
	}

	start := "0-0" // start from the beginning
	//Auto claim messages that have been idle for 300 seconds
	for {
		messages, nextStart, err := rdb.XAutoClaim(ctx, &redis.XAutoClaimArgs{
			Stream:   os.Getenv("STREAM_NAME"),
			Group:    os.Getenv("CUSTOMER_GROUPNAME"),
			Consumer: "testConsumer",
			MinIdle:  300000 * time.Millisecond, // claim messages that have been idle for 300 seconds
			Start:    start,                     // start from the last message
			Count:    100,                       // claim 100 messages at a time
		}).Result()

		if err != nil {
			log.Fatal(err)
		}

		for _, event := range messages {
			log.Warnf("claim Message: \"%s\"", event.Values["message"])

			// Acknowledge the message
			_, err := rdb.XAck(ctx, os.Getenv("STREAM_NAME"), os.Getenv("CUSTOMER_GROUPNAME"), event.ID).Result()
			if err != nil {
				log.Fatal(err)
			}
		}

		start = nextStart
	}
}

func Consumer(log *logrus.Logger) {
	//parameters for connecting to redis cluster
	options := redis.ClusterOptions{
		Addrs:    []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	//connect to redis cluster
	rdb := redis.NewClusterClient(&options)

	//Creating a consumer group, if exists, we can ignore the error
	_, err := rdb.XGroupCreateMkStream(ctx, os.Getenv("STREAM_NAME"), os.Getenv("CUSTOMER_GROUPNAME"), "$").Result()
	if err != nil {
		log.Error(err)
	}

	//Reading messages from stream
	for {
		messages, err := rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    os.Getenv("CUSTOMER_GROUPNAME"),
			Consumer: "testConsumer",
			Streams:  []string{os.Getenv("STREAM_NAME"), ">"},
			Block:    0,     // 0 means block until a new message arrives
			Count:    1,     // read 1 message at a time
			NoAck:    false, // set to false to enable message acknowledgment
		}).Result()
		if err != nil {
			log.Fatal(err)
		}

		for _, message := range messages {
			for _, event := range message.Messages {
				log.Infof("Receive Message: \"%s\"", event.Values["message"])

				// Acknowledge the message
				_, err := rdb.XAck(ctx, os.Getenv("STREAM_NAME"), os.Getenv("CUSTOMER_GROUPNAME"), event.ID).Result()
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func main() {
	log := initLogger()
	log.Info("producer start!")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("cnanot load .env file")
	}

	//parameters for connecting to redis cluster
	options := redis.ClusterOptions{
		Addrs:    []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	//connect to redis cluster
	rdb := redis.NewClusterClient(&options)

	//check connection
	pingResult, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	// PONG
	log.Info(pingResult)

	go Producer(log)  // start producer
	go AutoClaim(log) // start auto claim, auto claim will claim messages that have been idle for 300 seconds
	Consumer(log)     // start consumer
}
