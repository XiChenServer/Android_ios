package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

func main() {
	// Kafka 配置
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}

	// 创建 Kafka 生产者
	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		return
	}

	defer producer.Close()

	topic := "myTopic"
	message := "Hello, Kafka!"
	for {
		// 发送消息到指定主题
		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(message),
		}, nil)
		if err != nil {
			fmt.Printf("Failed to produce message: %s\n", err)
			return
		}
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Message sent successfully")
}
