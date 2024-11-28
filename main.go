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

	// Define the root route to provide API info
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Provide basic API info
		apiInfo := `
		{
			"message": "Welcome to the Receipt Processor API!",
			"endpoints": {
				"/process-receipt": {
					"method": "POST",
					"description": "Processes a new receipt.",
					"body": {
						"Retailer": "Name of the retailer",
						"PurchaseDate": "Date and time of purchase (ISO 8601)",
						"PurchaseTime": "Exact time of purchase (ISO 8601)",
						"Items": [
							{
								"shortDescription": "Name of the item",
								"price": "Price of the item (string)"
							}
						],
						"Total": "Total amount of the receipt",
						"RetailerId": "Unique ID for the retailer"
					}
				},
				"/get-points/{id}": {
					"method": "GET",
					"description": "Retrieves points for a specific receipt by ID."
				}
			}
		}
		`
		w.Write([]byte(apiInfo))
	}).Methods("GET")

	// Set up the routes using the handlers package
	handlers.SetupRoutes(router, receiptStore)

	// Define the routes for processing receipts and getting points
	router.Handle("/process-receipt", handlers.NewProcessReceiptHandler(receiptStore)).Methods("POST")
	router.Handle("/get-points/{id}", handlers.NewGetPointsHandler(receiptStore)).Methods("GET")

	// Define the server address
	serverAddr := ":8080"

	// Start the HTTP server and log any fatal errors
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(serverAddr, router))
}
