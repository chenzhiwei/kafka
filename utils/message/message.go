package message

import (
	"fmt"

	kafka "github.com/segmentio/kafka-go"

	v1 "github.com/chenzhiwei/kafka/api/v1"
)

func ToKafkaMessages(data *v1.Data) ([]kafka.Message, error) {
	var kMessages []kafka.Message
	for _, message := range data.Messages {
		var kMessage kafka.Message

		kMessage.Topic = message.Topic
		if kMessage.Topic == "" {
			kMessage.Topic = data.Topic
		}
		if kMessage.Topic == "" {
			return nil, fmt.Errorf("no topic found for message %v", message)
		}

		kMessage.Key = []byte(message.Key)
		kMessage.Value = []byte(message.Value)

		for k, v := range message.Headers {
			kMessage.Headers = append(kMessage.Headers, kafka.Header{
				Key:   k,
				Value: []byte(v),
			})
		}

		kMessages = append(kMessages, kMessage)
	}

	return kMessages, nil
}
