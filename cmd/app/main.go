package main

import (
	"wildberries/l0/pkg/logger"

	"wildberries/l0/internal/broker"
	"wildberries/l0/internal/domain"
	"wildberries/l0/internal/routes"

	"wildberries/l0/internal/view"

	"net/http"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	log := logger.NewLogger()
	log.Info("Starting server...")

	view := view.NewView(log)

	domain := domain.NewDomain(log)

	log.Info("Initializing router...")

	mux := routes.InitRouter(view)

	log.Info("Restoring cache...")

	err := view.RestoreCache()

	if err != nil {
		log.Error("Error restoring cache: " + err.Error())
	}

	log.Info("Cache restored...")

	log.Info("Server started on: " + os.Getenv("SERVER_ADDRESS"))

	go func() {
		log.Info("Starting broker...")
		oc := broker.NewOrderConsumer("broker:9092", "orders_group_id")
		err = oc.Listen("orders", func(m kafka.Message) {
			err := domain.HandleMessage(m)
			if err != nil {
				log.Error("Error handling message: " + err.Error())
			}
		})

		if err != nil {
			log.Error("Error listening to topic: " + err.Error())
		}
	}()

	if err := http.ListenAndServe(os.Getenv("SERVER_ADDRESS"), mux); err != nil {
		log.Error("Error starting server: " + err.Error())
	}

	// log := logger.NewLogger()
	// p, err := kafka.NewProducer(&kafka.ConfigMap{
	// 	"bootstrap.servers": "kafka:9092,localhost:9092",
	// 	"client.id":         "orders",
	// 	"acks":              "all",
	// })

	// if err != nil {
	// 	log.Error("Failed to create producer: " + err.Error())
	// }

	// deliverch := make(chan kafka.Event, 10000)
	// topic := "orders"
	// err = p.Produce(&kafka.Message{
	// 	TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	// 	Value:          []byte(`{"test": "test"}`),
	// }, deliverch)
	// if err != nil {
	// 	log.Error("Failed to produce message: " + err.Error())
	// }

}
