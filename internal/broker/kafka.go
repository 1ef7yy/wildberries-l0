package broker

import (
	"fmt"
	"wildberries/l0/pkg/logger"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	log := logger.NewLogger()
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "orders",
		"acks":              "all",
	})

	if err != nil {
		log.Error("Failed to create producer: " + err.Error())
	}

	fmt.Printf("$v+\n", p)
}
