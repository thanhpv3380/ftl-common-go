package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/thanhpv3380/ftl-common-go/common"
	"github.com/thanhpv3380/ftl-common-go/errors"
	"github.com/thanhpv3380/ftl-common-go/modules/logger"
	"github.com/thanhpv3380/ftl-common-go/utils"
)

type KafkaConfig struct {
	ClusterID string
	NodeID    string
	Brokers   []string
	Timeout   int
}

type KafkaClient struct {
	producer *KafkaProducer
}

type CurrentResponse struct {
	MessageID string
	Data      interface{}
}

var currentResponse chan CurrentResponse

func NewKafkaClient(ctx context.Context, config *KafkaConfig) (*KafkaClient, error) {
	currentResponse = make(chan CurrentResponse)

	producerConfig := &KafkaProducerConfig{
		ClusterID: config.ClusterID,
		Brokers:   config.Brokers,
		NodeID:    config.NodeID,
		Timeout:   config.Timeout,
	}

	producer, producerErr := NewKafkaProducer(producerConfig)
	if producerErr != nil {
		return nil, producerErr
	}

	kafkaClient := &KafkaClient{
		producer: producer,
	}

	handleMessage := func(msg *common.Message) error {
		currentResponse <- CurrentResponse{msg.MessageID, msg.Data}
		return nil
	}

	_, consumerErr := NewKafkaConsumer(ctx, &KafkaConsumerConfig{
		ClusterID: config.ClusterID,
		Brokers:   config.Brokers,
		GroupID:   config.ClusterID,
		Topic:     fmt.Sprintf("%s.response.%s", config.ClusterID, config.NodeID),
	}, handleMessage)

	if consumerErr != nil {
		return nil, producerErr
	}

	return kafkaClient, nil
}

func SendRequest(
	transactionID string,
	topic string,
	uri string,
	data interface{},
	waitForResponse bool,
) (*interface{}, error) {
	message := createMessageRequest(transactionID, uri, data)
	messageID, _ := SendMessage(topic, message)

	if !waitForResponse {
		return nil, nil
	}

	timeout := make(chan struct{})
	done := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(Producer.config.Timeout) * time.Second)
		select {
		case <-done:
			return
		default:
			timeout <- struct{}{}
		}
	}()

	for {
		select {
		case <-timeout:
			close(timeout)
			close(done)
			return nil, errors.NewTimeoutError(nil, "", nil)
		case res := <-currentResponse:
			if res.MessageID == messageID {
				close(timeout)
				close(done)

				response, err := utils.ExtractResponse(res.Data)
				if err != nil {
					return nil, err
				}

				return response, nil
			}
		}
	}
}

func SendResponse(
	msg common.Message,
	responseData interface{},
) error {
	if msg.ResponseDestination.Topic == "" || msg.ResponseDestination.URI == "" {
		return nil
	}

	if responseData == nil {
		err := errors.NewUriNotFound("")
		logger.Error("", err)
		SendMessage(msg.ResponseDestination.Topic, createMessageResponse(msg.TransactionID, msg.ResponseDestination.URI, err, msg.MessageID))
		return nil
	}

	SendMessage(msg.ResponseDestination.Topic, createMessageResponse(msg.TransactionID, msg.ResponseDestination.URI, responseData, msg.MessageID))
	return nil
}
