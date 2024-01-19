package test

import (
	"bufio"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"os/signal"
	"strings"
	"testing"
)

func Test_kafka(t *testing.T) {
	// 创建 Kafka 生产者
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		return
	}
	defer producer.Close()

	// 创建 Kafka 消费者
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "my-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		return
	}
	defer consumer.Close()

	// 订阅主题
	topic := "123"
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Printf("Failed to subscribe to topic: %s\n", err)
		return
	}

	// 用于处理中断信号的通道
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	defer close(sigchan)

	// 创建输入读取器
	reader := bufio.NewReader(os.Stdin)

	// 启动循环
	for {
		select {
		case sig := <-sigchan:
			// 处理中断信号，停止循环
			fmt.Printf("Caught signal %v: stopping\n", sig)
			return

		default:
			// 从输入读取消息，发送到 Kafka
			fmt.Print("Enter message to produce (or 'exit' to stop): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "exit" {
				// 输入 exit 时退出循环
				fmt.Println("Exiting...")
				return
			}

			// 发送消息到 Kafka
			err := producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          []byte(input),
			}, nil)
			if err != nil {
				fmt.Printf("Failed to produce message: %s\n", err)
			}
		}

		// 从 Kafka 接收并打印消息
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received message: %s\n", string(msg.Value))
		} else {
			fmt.Printf("Error reading message: %v\n", err)
		}
	}
}
