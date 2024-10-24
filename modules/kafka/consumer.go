package kafka

import (
	"context"
	"encoding/json"
	"ftl/kafi-common/common"
	"ftl/kafi-common/modules/logger"
	"sync"

	"github.com/IBM/sarama"
)

type KafkaConsumerConfig struct {
	ClusterID     string
	Brokers       []string
	GroupID       string
	Topic         string
	UseConcurrent bool // Xác định xử lý song song hay không
}

type KafkaConsumerGroupHandler struct {
	consumerGroup sarama.ConsumerGroup
	handleMessage func(*common.Message) error
	config        *KafkaConsumerConfig
	wg            *sync.WaitGroup
}

func NewKafkaConsumer(ctx context.Context, config *KafkaConsumerConfig, handleMessage func(*common.Message) error) (*KafkaConsumerGroupHandler, error) {
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

	kc := &KafkaConsumerGroupHandler{
		consumerGroup: consumerGroup,
		wg:            &sync.WaitGroup{},
		handleMessage: handleMessage,
	}

	kc.wg.Add(1)

	go func() {
		defer kc.wg.Done()

		for {
			err := consumerGroup.Consume(ctx, []string{config.Topic}, &KafkaConsumerGroupHandler{
				handleMessage: handleMessage,
				config:        config,
			})

			if err != nil {
				logger.Error("Error from consumer", err)
				return
			}

			if ctx.Err() != nil {
				logger.Info("xxx1")
				return
			}
		}
	}()

	return kc, nil
}

func (h *KafkaConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *KafkaConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *KafkaConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		messageParsed, err := h.parseMessage(msg)
		if err != nil {
			continue
		}

		if h.config.UseConcurrent {
			go func(msg *sarama.ConsumerMessage) {
				if err := h.handleMessage(messageParsed); err != nil {
					logger.Error("Error handling message", err)
				}
				session.MarkMessage(msg, "")
			}(msg)
		} else {
			if err := h.handleMessage(messageParsed); err != nil {
				logger.Error("Error handling message", err)
			}
			session.MarkMessage(msg, "")
		}
	}

	return nil
}

func (h *KafkaConsumerGroupHandler) parseMessage(msg *sarama.ConsumerMessage) (*common.Message, error) {
	var kafkaMessage common.Message
	if err := json.Unmarshal(msg.Value, &kafkaMessage); err != nil {
		logger.Error("Error parse kafka message", err)
		return nil, err
	}

	jsonData, errParseData := json.Marshal(kafkaMessage.Data)
	if errParseData != nil {
		logger.Error("Error parse data in kafka", errParseData)
	}

	logger.Info("Received message from Kafka", map[string]interface{}{
		"topic":         msg.Topic,
		"partition":     msg.Partition,
		"offset":        msg.Offset,
		"sourceID":      kafkaMessage.SourceID,
		"transactionID": kafkaMessage.TransactionID,
		"uri":           kafkaMessage.URI,
		"data":          string(jsonData),
	})

	return &kafkaMessage, nil
}
