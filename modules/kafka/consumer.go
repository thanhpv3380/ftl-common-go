package kafka

import (
	"context"
	"encoding/json"
	"github/thanhpv3380/ftl-common-go/common"
	"github/thanhpv3380/ftl-common-go/modules/logger"

	"github.com/IBM/sarama"
)

type KafkaConsumerConfig struct {
	ClusterID     string
	Brokers       []string
	GroupID       string
	Topic         string
	UseConcurrent bool // Xác định xử lý song song hay không
}

func NewKafkaConsumer(config *KafkaConsumerConfig, handleMessage func(*common.Message) error, stopChan <-chan struct{}) (sarama.ConsumerGroup, error) {
	saramaConfig := sarama.NewConfig()

	consumerGroup, err := sarama.NewConsumerGroup(config.Brokers, config.GroupID, saramaConfig)
	if err != nil {
		logger.Error("Error creating consumer group", err)
		return nil, err
	}

	logger.Info("Consumer created successfully", map[string]interface{}{
		"clusterId": config.ClusterID,
		"groupId":   config.GroupID,
		"topic":     config.Topic,
	})

	go func() {
		for {
			select {
			case <-stopChan:
				logger.Info("Stopping consumer...")
				return
			default:
				err := consumerGroup.Consume(context.Background(), []string{config.Topic}, &KafkaConsumerGroupHandler{
					handleMessage: handleMessage,
					config:        config,
				})
				if err != nil {
					logger.Error("Error from consumer", err)
					return
				}
			}
		}
	}()

	return consumerGroup, nil
}

type KafkaConsumerGroupHandler struct {
	handleMessage func(*common.Message) error
	config        *KafkaConsumerConfig
}

// Setup được gọi khi consumer group bắt đầu
func (h *KafkaConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	// Thực hiện các thiết lập cần thiết nếu có
	logger.Info("start")
	return nil
}

// Cleanup được gọi khi consumer group kết thúc
func (h *KafkaConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	// Dọn dẹp tài nguyên nếu cần
	logger.Info("end")
	return nil
}

func (h *KafkaConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	logger.Info("xxx")
	for msg := range claim.Messages() {
		logger.Info("receieve")
		messageParsed, err := h.parseMessage(msg)
		if err != nil {
			continue
		}

		if h.config.UseConcurrent {
			go func(msg *sarama.ConsumerMessage) {
				if err := h.handleMessage(messageParsed); err != nil {
					logger.Error("Error handling message", err)
				}
				logger.Info("1mark")
				session.MarkMessage(msg, "")
			}(msg)
		} else {
			if err := h.handleMessage(messageParsed); err != nil {
				logger.Error("Error handling message", err)
			}
			logger.Info("mark")
			session.MarkMessage(msg, "")
		}
	}

	return nil
}

func (h *KafkaConsumerGroupHandler) parseMessage(msg *sarama.ConsumerMessage) (*common.Message, error) {
	var kafkaMessage common.KafkaMessage
	if err := json.Unmarshal(msg.Value, &kafkaMessage); err != nil {
		logger.Error("Error parse kafka message", err)
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
