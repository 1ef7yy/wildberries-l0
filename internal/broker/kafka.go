package broker

import (
	"fmt"
	"wildberries/l0/pkg/logger"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type OrderPlacer struct {
	log        logger.Logger
	producer   *kafka.Producer
	topic      string
	deliverych chan kafka.Event
}

type OrderConsumer struct {
	log      logger.Logger
	consumer *kafka.Consumer
}

func NewOrderPlacer(p *kafka.Producer, topic string) *OrderPlacer {
	return &OrderPlacer{
		log:        logger.NewLogger(),
		producer:   p,
		topic:      topic,
		deliverych: make(chan kafka.Event, 10000)}
}

func NewOrderConsumer(hosts, groupID string) *OrderConsumer {
	log := logger.NewLogger()
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": hosts,
		"group.id":          groupID,
		"auto.offset.reset": "smallest",
	})

	if err != nil {
		log.Error("Error creating consumer: " + err.Error())
	}

	log.Info("Successfully created consumer")
	return &OrderConsumer{
		log:      log,
		consumer: kafkaConsumer,
	}
}

func (op *OrderPlacer) PlaceOrder(orderType string, size int) error {

	var (
		format  = fmt.Sprintf("%s - %d", orderType, size)
		payload = []byte(format)
	)

	err := op.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &op.topic,
			Partition: kafka.PartitionAny,
		},
		Value: payload,
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
		}
	}

}
