package broker

import (
	"wildberries/l0/internal/models"
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

type Broker interface {
	Listen()
	PlaceOrder(models.Order)
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
