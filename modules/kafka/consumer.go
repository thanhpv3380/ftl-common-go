package kafka

import (
	"encoding/json"
	"ftl/kafi-common/common"
	"ftl/kafi-common/modules/logger"

	"github.com/IBM/sarama"
)

type KafkaConsumerConfig struct {
	ClusterID string
	Brokers   []string
}

type KafkaConsumer struct {
	config   *KafkaConsumerConfig
	consumer sarama.Consumer
}

func NewKafkaConsumer(config *KafkaConsumerConfig) (*KafkaConsumer, error) {
	consumer, err := sarama.NewConsumer(config.Brokers, nil)
	if err != nil {
		logger.Error("Error create consumer", err)
		return nil, err
	}

	logger.Info("Consumer created successfully")
	return &KafkaConsumer{
		config:   config,
		consumer: consumer,
	}, nil
}

func (c *KafkaConsumer) ConsumeMessages(topic string) {
	partitionList, err := c.consumer.Partitions(topic)
	if err != nil {
		logger.Error("Error get the list of partitions", err)
		return
	}

	for _, partition := range partitionList {
		pc, err := c.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			logger.Error("Error start consuming partition", err)
			continue
		}

		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				c.processMessage(msg)
			}
		}(pc)
	}
}

func (c *KafkaConsumer) processMessage(msg *sarama.ConsumerMessage) (*common.Message, error) {
	var kafkaMessage common.KafkaMessage
	if err := json.Unmarshal(msg.Value, &kafkaMessage); err != nil {
		logger.Error("Error convert message raw to kafka message", err)
		return nil, err
	}

	logger.Info("Received message from Kafka", map[string]interface{}{
		"topic":         msg.Topic,
		"partition":     msg.Partition,
		"offset":        msg.Offset,
		"sourceID":      kafkaMessage.Message.SourceID,
		"transactionID": kafkaMessage.Message.TransactionID,
		"uri":           kafkaMessage.Message.URI,
		"data":          kafkaMessage.Message.Data,
	})

	return &kafkaMessage.Message, nil
}

func (c *KafkaConsumer) Close() {
	if err := c.consumer.Close(); err != nil {
		logger.Error("Error close consumer", err)
	}
}
