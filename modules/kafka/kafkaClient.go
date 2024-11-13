package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/thanhpv3380/ftl-common-go/common"
	"github.com/thanhpv3380/ftl-common-go/errors"
)

type KafkaConfig struct {
	ClusterID string
	NodeId    string
	Brokers   []string
}

type KafkaClient struct {
	producer *KafkaProducer
}

type CurrentResponse struct {
	MessageID int
	Data      interface{}
}

var currentResponse chan CurrentResponse

func NewKafkaClient(ctx context.Context, config *KafkaConfig) (*KafkaClient, error) {
	currentResponse = make(chan CurrentResponse)

	producerConfig := &KafkaProducerConfig{
		ClusterID: config.ClusterID,
		Brokers:   config.Brokers,
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
		Topic:     fmt.Sprintf("%s.response.%s", config.ClusterID, config.NodeId),
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
	messageType common.MessageType,
	waitForResponse bool,
) (*interface{}, error) {
	msgId, _ := SendMessage(transactionID, topic, uri, data, messageType)

	if waitForResponse {
		var response *interface{}

		timeout := make(chan struct{})

		go func() {
			time.Sleep(10 * time.Second)
			timeout <- struct{}{}
		}()

		for {
			select {
			case <-timeout:
				return nil, errors.NewTimeoutError(nil, "", nil)
			case res := <-currentResponse:
				if res.MessageID == msgId {
					response = &res.Data
					close(timeout)
					return response, nil
				}
			}
		}
	}

	return nil, nil
}
