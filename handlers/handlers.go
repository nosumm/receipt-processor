package handlers

import (
	"net/http"
	"receipt-processor/storage"
	"github.com/gorilla/mux"
)

// SetupRoutes sets up the HTTP routes for the application
func SetupRoutes(router *mux.Router, receiptStore *storage.ReceiptStore) {
	// Define the route for processing receipts
	router.Handle("/process", NewProcessReceiptHandler(receiptStore)).Methods("POST")

	// Define the route for retrieving points
	router.Handle("/points/{id}", NewGetPointsHandler(receiptStore)).Methods("GET")

	// Health check route
	router.HandleFunc("/health", HealthCheckHandler).Methods("GET")
}

// HealthCheckHandler returns the status of the server
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running"))
}
