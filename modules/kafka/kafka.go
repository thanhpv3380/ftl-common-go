package kafka

import (
	"github.com/IBM/sarama"
)

var producer sarama.SyncProducer

// InitKafka initializes the Kafka producer
func InitKafka(brokers []string) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	var err error
	producer, err = sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return err
	}
	return nil
}

// ProduceMessage sends a message to a Kafka topic
func ProduceMessage(topic string, key, value []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(value),
	}

	_, _, err := producer.SendMessage(msg)
	return err
}

// CloseKafka closes the Kafka producer
func CloseKafka() error {
	return producer.Close()
}
