package broker

import (
	"os"
	"wildberries/l0/internal/models"
	"wildberries/l0/pkg/logger"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	KafkaConn = os.Getenv("KAFKA_CONN")
)

type OrderConsumer struct {
	log      logger.Logger
	consumer *kafka.Consumer
}

type Broker interface {
	Listen()
	PlaceOrder(models.Order)
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
