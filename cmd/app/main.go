package main


import (
	"net/http"
	"os"

	"wildberries/l0/pkg/logger"

	"wildberries/l0/internal/routes"
	"wildberries/l0/internal/view"
)

func main() {
	log := logger.NewLogger()
	log.Info("Starting server...")

	view := view.NewView(log)

	log.Info("Initializing router...")

	mux := routes.InitRouter(view)

	log.Info("Server started on: " + os.Getenv("SERVER_ADDRESS"))

	if err := http.ListenAndServe(os.Getenv("SERVER_ADDRESS"), mux); err != nil {
		log.Error("Error starting server: " + err.Error())
	}

}
