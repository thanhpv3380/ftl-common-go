package main

import (
	"context"
	"ftl/kafi-common/common"
	"ftl/kafi-common/modules/kafka"
	"ftl/kafi-common/modules/logger"
	"os"
	"os/signal"
	"syscall"
)

// func startKafkaConsumer(ctx context.Context, wg *sync.WaitGroup) {
// 	wg.Add(1)

// 	go func() {
// 		defer wg.Done()

// 		kafkaUrls := []string{"192.168.4.129:9092"}

// 		handleMessage := func(msg *common.Message) error {
// 			logger.Info("handleMessage: xxx")
// 			return nil
// 		}

// 		consumerGroup, err := kafka.NewKafkaConsumer(&kafka.KafkaConsumerConfig{
// 			ClusterID: "test",
// 			Brokers:   kafkaUrls,
// 			GroupID:   "test1",
// 			Topic:     "test",
// 		}, handleMessage)

// 		if err != nil {
// 			return
// 		}

// 		defer func() {
// 			if err := consumerGroup.Close(); err != nil {
// 				log.Println("Lỗi khi dừng Kafka consumer:", err)
// 			}
// 		}()

// 		<-ctx.Done()
// 		log.Println("Dừng Kafka consumer...")
// 	}()
// }

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Đảm bảo hủy context khi main kết thúc

	logger.InitLogger(nil)

	kafkaUrls := []string{"192.168.4.129:9092"}

	handleMessage := func(msg *common.Message) error {
		logger.Info("handleMessage: xxx")
		return nil
	}

	_, err := kafka.NewKafkaConsumer(ctx, &kafka.KafkaConsumerConfig{
		ClusterID: "test",
		Brokers:   kafkaUrls,
		GroupID:   "test1",
		Topic:     "test",
	}, handleMessage)

	if err != nil {
		return
	}
	waitForShutdown(cancel)
	<-ctx.Done()
	logger.Info("Main: Context canceled, exiting...")
}

func waitForShutdown(cancel context.CancelFunc) {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm
	cancel()
}
