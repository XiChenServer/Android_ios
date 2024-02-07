package dao

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"sync"
)

var (
	producerMu sync.Mutex
	producer   *kafka.Producer
)

func initKafka() *kafka.Producer {
	producerMu.Lock()
	defer producerMu.Unlock()

	if producer == nil {
		var err error
		producer, err = createKafkaProducer()
		if err != nil {
			log.Fatal("Error creating Kafka producer: ", err)
		}
	}
	return producer
}

func createKafkaProducer() (*kafka.Producer, error) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "Herding_cattle_and_straying_horses",
	}

	p, err := kafka.NewProducer(configMap)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func ProduceMessage(topic string, message string) error {
	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}

	err := initKafka().Produce(kafkaMessage, nil)
	if err != nil {
		return err
	}

	return nil
}
