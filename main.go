// main application entry point

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"receipt-processor/handlers"
	"receipt-processor/storage"
)

// main is the entry point of the receipt processing application
// It sets up the server with the following key components:
// 1. In-memory receipt storage
// 2. Router for handling HTTP routes
// 3. Route configuration for receipt-related endpoints
func main() {
	// Create a new in-memory receipt store
	// This store will manage the storage and retrieval of receipts
	receiptStore := storage.NewReceiptStore()

	// Create a new router using gorilla/mux
	// Provides advanced routing capabilities like path variables and method-based routing
	router := mux.NewRouter()

	// Configure routes for receipt processing and point retrieval
	// Passes the receipt store to enable data persistence and retrieval
	handlers.SetupRoutes(router, receiptStore)

	// Configure server parameters
	serverAddr := ":8080"
	
	// Log server startup
	log.Printf("Server starting on %s", serverAddr)

	// Start the HTTP server
	// ListenAndServe blocks and runs the server until an error occurs
	log.Fatal(http.ListenAndServe(serverAddr, router))
}
