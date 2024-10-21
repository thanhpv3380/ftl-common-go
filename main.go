package main

import (
	"ftl/kafi-common/common"
	"ftl/kafi-common/modules/kafka"
	"ftl/kafi-common/modules/logger"
)

func main() {
	logger.InitLogger()

	kafkaUrls := []string{"10.40.80.236:9092"}

	consumerConfig := &kafka.KafkaConsumerConfig{
		ClusterID: "test",
		Brokers:   kafkaUrls,
	}

	Consumer, consumerErr := kafka.NewKafkaConsumer(consumerConfig)
	if consumerErr != nil {
		return
	}

	// Cấu hình Kafka Producer
	producerConfig := &kafka.KafkaProducerConfig{
		ClusterID: "test",
		Brokers:   kafkaUrls,
	}

	_, producerErr := kafka.NewKafkaProducer(producerConfig)
	if producerErr != nil {
		return
	}

	// Bắt đầu tiêu thụ tin nhắn
	Consumer.ConsumeMessages("test")

	message := common.KafkaResponse{
		Data: map[string]interface{}{"key1": "value1", "key2": "value2"}, // Hoặc có thể là một struct khác
	}

	kafka.SendMessage("xxx", "test", "/update", message, common.REQUEST)
}
