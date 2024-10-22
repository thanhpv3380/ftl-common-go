package main

import (
	"ftl/kafi-common/common"
	"ftl/kafi-common/modules/kafka"
	"ftl/kafi-common/modules/logger"
	"os"
	"os/signal"
	"syscall"
)

func run() {
	kafkaUrls := []string{"192.168.4.129:9092"}

	handleMessage := func(msg *common.Message) error {
		logger.Info("handleMessage: xxx")
		return nil
	}

	// Tạo kênh tín hiệu dừng
	stopChan := make(chan struct{})

	consumer, consumerErr := kafka.NewKafkaConsumer(&kafka.KafkaConsumerConfig{
		ClusterID: "test",
		Brokers:   kafkaUrls,
		GroupID:   "test1",
		Topic:     "test2",
	}, handleMessage, stopChan)
	if consumerErr != nil {
		return
	}

	defer consumer.Close()

	producerConfig := &kafka.KafkaProducerConfig{
		ClusterID: "test",
		Brokers:   kafkaUrls,
	}

	_, producerErr := kafka.NewKafkaProducer(producerConfig)
	if producerErr != nil {
		return
	}

	message := common.KafkaResponse{
		Data: "xxx", // Hoặc có thể là một struct khác
	}
	kafka.SendMessage("xxx", "test2", "/update", message, common.REQUEST)

	// Lắng nghe tín hiệu dừng
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan

	// Gửi tín hiệu dừng đến consumer
	close(stopChan)
}

func main() {
	logger.InitLogger(nil)

	logger.Info("start main", map[string]interface{}{"name": "test"})

	// run()
}
