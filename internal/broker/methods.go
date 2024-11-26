package broker

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (oc *OrderConsumer) Listen(topic string, handleFunc func(kafka.Message)) error {
	err := oc.consumer.Subscribe(topic, nil)

	if err != nil {
		oc.log.Error("Error subscribing to topic: " + err.Error())
		return err
	}

	for {
		ev := oc.consumer.Poll(10)
		switch e := ev.(type) {
		case *kafka.Message:
			oc.log.Info(fmt.Sprintf("Message received. Key: %s, Value: %s. Partition: %d, Offset: %d\n", string(e.Key), string(e.Value), e.TopicPartition.Partition, e.TopicPartition.Offset))
			handleFunc(*e)
		case kafka.Error:
			oc.log.Error("Error from consumer: " + e.Error())
			return e
		}
	}

}
