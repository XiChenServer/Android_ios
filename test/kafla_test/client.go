package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"os/signal"
)

func main() {
	// Kafka 配置
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	}

	// 创建 Kafka 消费者
	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		return
	}

	defer consumer.Close()

	topic := "myTopic"

	// 订阅主题
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Printf("Failed to subscribe to topic: %s\n", err)
		return
	}

	// 处理信号
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("Message on %s: %s\n", e.TopicPartition, string(e.Value))
			case kafka.Error:
				fmt.Printf("Error: %v\n", e)
				run = false
			default:
				fmt.Printf("Ignored event: %s\n", e)
			}
		}
	}

	fmt.Println("Consumer closed")
}
