package main

import (
	"wildberries/l0/pkg/logger"

	"wildberries/l0/internal/routes"

	"wildberries/l0/internal/view"

	"net/http"
	"os"
)

func main() {
	log := logger.NewLogger()
	log.Info("Starting server...")

	view := view.NewView(log)

	log.Info("Initializing router...")

	mux := routes.InitRouter(view)

	log.Info("Restoring cache...")

	err := view.RestoreCache()

	if err != nil {
		log.Error("Error restoring cache: " + err.Error())
	}

	log.Info("Cache restored...")

	log.Info("Server started on: " + os.Getenv("SERVER_ADDRESS"))

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
