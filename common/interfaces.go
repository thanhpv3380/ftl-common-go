package common

import "github.com/IBM/sarama"

type ParamError struct {
	Code          string                 `json:"code"`
	MessageParams map[string]interface{} `json:"messageParams,omitempty"`
}

type Status struct {
	Code          string                  `json:"code"`
	MessageParams map[string]interface{}  `json:"messageParams,omitempty"`
	Params        map[string][]ParamError `json:"params,omitempty"`
}

type KafkaConfigOptions struct {
	ClientID              string
	ProducerRequiredAcks  sarama.RequiredAcks
	ProducerRetryMax      int
	ProducerReturnSuccess bool
	ConsumerReturnErrors  bool
}

type KafkaMessage struct {
	Topic   string  `json:"topic"`
	Message Message `json:"message"`
}

type Message struct {
	MessageType         MessageType          `json:"messageType"`
	SourceID            string               `json:"sourceId"`
	MessageID           string               `json:"messageId"`
	TransactionID       string               `json:"transactionId"`
	URI                 string               `json:"uri"`
	ResponseDestination *ResponseDestination `json:"responseDestination,omitempty"`
	Data                KafkaResponse        `json:"data"`
}

type MessageType string

const (
	MESSAGE  MessageType = "MESSAGE"
	REQUEST  MessageType = "REQUEST"
	RESPONSE MessageType = "RESPONSE"
)

type ResponseDestination struct {
	Topic string `json:"topic"`
	URI   string `json:"uri"`
}

type KafkaResponse struct {
	Status *Status     `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
