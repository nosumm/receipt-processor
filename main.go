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
					"PurchaseDate": "Date and time of purchase",
					"PurchaseTime": "Exact time of purchase",
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
		},
		"example": {
			"Process Receipt": { "curl -X POST http://localhost:8080/process-receipt -d '{\"Retailer\": \"My Retailer\",\"PurchaseDate\": \"2024-11-28T10:30:00Z\", \"PurchaseTime\": \"2024-11-28T10:30:00Z\", \"Items\": [{\"shortDescription\": \"Item 1\",\"price\": \"10.00\"},{\"shortDescription\": \"Item 2\",\"price\": \"20.50\"}], \"Total\": \"30.50\", \"RetailerId\": \"123\"}' -H \"Content-Type: application/json\"",
			"response": "{\"id\":\"22205988-20ad-42ce-9c43-804766bd019e\"}"
			},
			"Get Points": { "curl http://localhost:8080/get-points/22205988-20ad-42ce-9c43-804766bd019e",
			"response": "{\"points\":47}"
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
