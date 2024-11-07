package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/thanhpv3380/ftl-common-go/common"
	"github.com/thanhpv3380/ftl-common-go/modules/kafka"
	"github.com/thanhpv3380/ftl-common-go/modules/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger.InitLogger(nil)

	kafkaUrls := []string{"10.40.80.236:9092"}

	handleMessage := func(msg *common.Message) error {
		logger.Info("handleMessage: xxx", map[string]interface{}{
			"msg": msg.URI,
		})
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
