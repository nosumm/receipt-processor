package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"receipt-processor/handlers"
	"receipt-processor/storage"
)

func main() {
	// Create a new in-memory receipt store
	receiptStore := storage.NewReceiptStore()

	// Create a new router using gorilla/mux
	router := mux.NewRouter()

	// Set up the routes using the handlers package
	handlers.SetupRoutes(router, receiptStore)

	router.Handle("/process-receipt", handlers.NewProcessReceiptHandler(receiptStore))
	router.Handle("/get-points/{id}", handlers.NewGetPointsHandler(receiptStore))


	// Define the server address
	serverAddr := ":8080"

	// Log server startup
	log.Printf("Server starting on %s", serverAddr)

	// Start the HTTP server and log any fatal errors
	log.Fatal(http.ListenAndServe(serverAddr, router))
}
