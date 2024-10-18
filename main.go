package main

import (
	"fmt"
	"ftl/kafi-common/modules/kafka"
	"log"
)

func main() {
	brokers := []string{"10.40.80.236:9092"}
	err := kafka.InitKafka(brokers)
	if err != nil {
		log.Fatalf("Failed to initialize Kafka: %v", err)
	}
	defer kafka.CloseKafka()

	topic := "notification"
	key := []byte("key")
	value := []byte("message value")

	err = kafka.ProduceMessage(topic, key, value)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	fmt.Println("Message sent successfully")
}
