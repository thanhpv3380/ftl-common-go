package kafka

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/thanhpv3380/ftl-common-go/common"
	"github.com/thanhpv3380/ftl-common-go/modules/logger"

	"github.com/IBM/sarama"
)

var messageId int = 0
var Producer *KafkaProducer

type KafkaProducerConfig struct {
	ClusterID string
	Brokers   []string
	NodeID    string
	Timeout   int
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
	topic string,
	message common.Message,
) (string, error) {
	value, convertErr := json.Marshal(message)
	if convertErr != nil {
		logger.Error("Error convert kafka message", convertErr)
		return message.MessageID, convertErr
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(message.MessageID),
		Value: sarama.ByteEncoder(value),
	}

	logger.Info("Send message to kafka", map[string]interface{}{
		"topic":   topic,
		"message": message,
	})

	_, _, err := Producer.producer.SendMessage(msg)
	if err != nil {
		logger.Error("Error send message to kafka", err)
		return message.MessageID, err
	}

	return message.MessageID, nil
}

func createMessageRequest(
	transactionID string,
	uri string,
	data interface{},
) common.Message {
	messageId += 1

	return common.Message{
		MessageType:   common.REQUEST,
		SourceID:      Producer.config.ClusterID,
		MessageID:     strconv.Itoa(messageId),
		TransactionID: transactionID,
		URI:           uri,
		Data:          data,
		ResponseDestination: &common.ResponseDestination{
			Topic: fmt.Sprintf("%s.response.%s", Producer.config.ClusterID, Producer.config.NodeID),
			URI:   string(common.REQUEST_RESPONSE),
		},
	}
}

func createMessageResponse(
	transactionID string,
	uri string,
	data interface{},
	messageId string,
) common.Message {
	var result interface{}

	switch v := data.(type) {
	case *common.GeneralError:
		result = map[string]interface{}{
			"status": v,
		}
	default:
		result = map[string]interface{}{
			"data": data,
		}
	}

	return common.Message{
		MessageType:   common.RESPONSE,
		SourceID:      Producer.config.ClusterID,
		MessageID:     messageId,
		TransactionID: transactionID,
		URI:           uri,
		Data:          result,
	}
}
