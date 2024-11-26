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

	domain := domain.NewDomain(log)

	view := view.NewView(log, domain)

	log.Info("Initializing router...")

	mux := routes.InitRouter(view)

	log.Info("Restoring cache...")

	err := view.RestoreCache()

	if err != nil {
		log.Error("Error restoring cache: " + err.Error())
	}

	log.Info("Cache restored...")

	kafkaHandleFunc := func(m kafka.Message) {
		err := domain.HandleMessage(m)
		if err != nil {
			log.Error("Error handling message: " + err.Error())
		}
	}
	oc := broker.NewOrderConsumer(os.Getenv("KAFKA_CONN"), "orders_group_id")

	go func() {
		log.Info("Starting broker...")
		err = oc.Listen("orders", kafkaHandleFunc)

		if err != nil {
			log.Error("Error listening to topic: " + err.Error())
		}
	}()

	log.Info("Server started on: " + os.Getenv("SERVER_ADDRESS"))
	if err := http.ListenAndServe(os.Getenv("SERVER_ADDRESS"), mux); err != nil {
		log.Error("Error starting server: " + err.Error())
	}

}
