package kafka

import (
	"encoding/json"
	"ftl/kafi-common/common"
	"ftl/kafi-common/modules/logger"
	"strconv"

	"github.com/IBM/sarama"
)

var messageId int = 0
var Producer *KafkaProducer

type KafkaProducerConfig struct {
	ClusterID string
	Brokers   []string
}

type KafkaProducer struct {
	config   *KafkaProducerConfig
	producer sarama.SyncProducer
}

func NewKafkaProducer(config *KafkaProducerConfig) (*KafkaProducer, error) {
	producer, err := sarama.NewSyncProducer(config.Brokers, nil)
	if err != nil {
		logger.Error("Error create producer", err)
		return nil, err
	}

	logger.Info("Producer created successfully")
	Producer = &KafkaProducer{
		config:   config,
		producer: producer,
	}

	return Producer, nil
}

func SendMessage(
	transactionID string,
	topic string,
	uri string,
	data common.KafkaResponse,
	messageType common.MessageType,
) error {
	kafkaMessage := createMessage(transactionID, topic, uri, data, messageType)

	value, err := json.Marshal(kafkaMessage)
	if err != nil {
		logger.Error("Error convert kafka message", err)
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: kafkaMessage.Topic,
		Key:   sarama.StringEncoder(kafkaMessage.Message.MessageID),
		Value: sarama.ByteEncoder(value),
	}

	partition, offset, err := Producer.producer.SendMessage(msg)
	if err != nil {
		logger.Error("Error send message to kafka", err)
		return err
	}

	logger.Info("Send message to kafka success", map[string]interface{}{
		"topic":     topic,
		"partition": partition,
		"offset":    offset,
		"message":   kafkaMessage,
	})

	return nil
}

func createMessage(
	transactionID string,
	topic string,
	uri string,
	data common.KafkaResponse,
	messageType common.MessageType,
) common.KafkaMessage {
	messageId += 1

	return common.KafkaMessage{
		Topic: topic,
		Message: common.Message{
			MessageType:   messageType,
			SourceID:      Producer.config.ClusterID,
			MessageID:     strconv.Itoa(messageId),
			TransactionID: transactionID,
			URI:           uri,
			Data:          data,
		},
	}
}
