package kafka

import (
	"ftl/kafi-common/common"

	"github.com/IBM/sarama"
)

func NewKafkaConfig(opts *common.KafkaConfigOptions) *sarama.Config {
	config := sarama.NewConfig()

	// ClientID: chuỗi định danh cho Kafka client
	if opts != nil && opts.ClientID != "" {
		config.ClientID = opts.ClientID
	} else {
		config.ClientID = "common"
	}

	// Producer.RequiredAcks: Xác định số lượng acknowledgment (xác nhận) mà Kafka producer yêu cầu từ Kafka broker sau khi gửi một message
	if opts != nil && opts.ProducerRequiredAcks != 0 {
		config.Producer.RequiredAcks = opts.ProducerRequiredAcks
	} else {
		config.Producer.RequiredAcks = sarama.WaitForAll
	}

	// Producer.Retry.Max: Số lần retry send message nếu gửi lỗi
	if opts != nil && opts.ProducerRetryMax != 0 {
		config.Producer.Retry.Max = opts.ProducerRetryMax
	} else {
		config.Producer.Retry.Max = 5
	}

	// Producer.Return.Successes: Xác định xem producer có trả về kết quả thành công sau khi gửi message hay không
	if opts != nil {
		config.Producer.Return.Successes = opts.ProducerReturnSuccess
	} else {
		config.Producer.Return.Successes = true
	}

	// Consumer.Return.Errors: Cho phép Kafka consumer trả về các lỗi xảy ra trong quá trình tiêu thụ message
	if opts != nil {
		config.Consumer.Return.Errors = opts.ConsumerReturnErrors
	} else {
		config.Consumer.Return.Errors = true
	}

	return config
}
