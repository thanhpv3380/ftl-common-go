package common

import (
	"fmt"

	"github.com/IBM/sarama"
)

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
	Data                interface{}          `json:"data"`
}

type MessageType string

const (
	MESSAGE          MessageType = "MESSAGE"
	REQUEST          MessageType = "REQUEST"
	RESPONSE         MessageType = "RESPONSE"
	REQUEST_RESPONSE MessageType = "REQUEST_RESPONSE"
)

type ResponseDestination struct {
	Topic string `json:"topic"`
	URI   string `json:"uri"`
}

type GeneralError struct {
	Code          string                  `json:"code"`
	MessageParams map[string]interface{}  `json:"messageParams"`
	Source        string                  `json:"source"`
	Params        map[string][]ParamError `json:"params"`
	IsSystemError bool                    `json:"isSystemError"`
}

func (e *GeneralError) Error() string {
	return fmt.Sprintf(e.Code)
}

// func (ge *GeneralError) UnmarshalJSON(data []byte) error {
// 	// Temporary struct to hold the unmarshalled data
// 	type Alias GeneralError
// 	aux := &struct {
// 		*Alias
// 	}{
// 		Alias: (*Alias)(ge),
// 	}

// 	// Attempt to unmarshal the JSON into the GeneralError fields
// 	if err := json.Unmarshal(data, &aux); err != nil {
// 		return fmt.Errorf("failed to unmarshal error data: %v", err)
// 	}

// 	// Handle custom deserialization of the Params and MessageParams if necessary
// 	// For example, if "params" is missing or malformed, you could set it to an empty map
// 	if ge.Params == nil {
// 		ge.Params = make(map[string][]ParamError)
// 	}
// 	if ge.MessageParams == nil {
// 		ge.MessageParams = make(map[string]interface{})
// 	}

// 	return nil
// }

type ParamError struct {
	Code          string                 `json:"code"`
	MessageParams map[string]interface{} `json:"messageParams,omitempty"`
}

type MessageResponse struct {
	Data   interface{}  `json:"data"`
	Status GeneralError `json:"status"`
}

type CommonResponse struct {
	Message string `json:"message"`
}

func CreateCommonResponse(message string) map[string]interface{} {
	if message == "" {
		message = "success"
	}

	return map[string]interface{}{
		"message": message,
	}
}
