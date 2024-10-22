package kafka

import (
	"encoding/json"
	"github/thanhpv3380/ftl-common-go/common"
	"github/thanhpv3380/ftl-common-go/modules/logger"
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

	value, convertErr := json.Marshal(kafkaMessage)
	if convertErr != nil {
		logger.Error("Error convert kafka message", convertErr)
		return convertErr
	}

	msg := &sarama.ProducerMessage{
		Topic: kafkaMessage.Topic,
		Key:   sarama.StringEncoder(kafkaMessage.Message.MessageID),
		Value: sarama.ByteEncoder(value),
	}

	logger.Info("Send message to kafka", map[string]interface{}{
		"topic":   topic,
		"message": kafkaMessage,
	})

	_, _, err := Producer.producer.SendMessage(msg)
	if err != nil {
		logger.Error("Error send message to kafka", err)
		return err
	}

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
