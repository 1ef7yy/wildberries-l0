package broker

import (
	"fmt"
	"wildberries/l0/internal/models"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (op *OrderPlacer) PlaceOrder(order models.Order) error {
	err := op.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &op.topic,
			Partition: kafka.PartitionAny,
		},
		Key:   []byte(order.OrderUid),
		Value: order.Data,
	}, op.deliverych)

	if err != nil {
		op.log.Error("Error producing message: " + err.Error())
		return err
	}

	return nil

}

func (oc *OrderConsumer) Listen(topic string, handleFunc func(kafka.Message)) error {
	err := oc.consumer.Subscribe(topic, nil)

	if err != nil {
		oc.log.Error("Error subscribing to topic: " + err.Error())
		return err
	}

	for {
		ev := oc.consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			oc.log.Info(fmt.Sprintf("Message received. Key: %s, Value: %s. Partition: %d, Offset: %d\n", string(e.Key), string(e.Value), e.TopicPartition.Partition, e.TopicPartition.Offset))
			handleFunc(*e)
		case kafka.Error:
			oc.log.Error("Error from consumer: " + e.Error())
			return e

		default:
			oc.log.Info("Ignored event: " + fmt.Sprint(e))
		}
	}

}
