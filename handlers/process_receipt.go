package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"receipt-processor/models"
	"receipt-processor/storage"
)

// NewProcessReceiptHandler creates a new handler for processing receipts
// It takes a receipt store as a parameter to manage receipt storage
func NewProcessReceiptHandler(store *storage.ReceiptStore) http.Handler {
	return &ProcessReceiptHandler{
		store: store,
	}
}

// ProcessReceiptHandler manages the processing of receipts
// It uses a receipt store to save and manage receipts
type ProcessReceiptHandler struct {
	store *storage.ReceiptStore
}

// ServeHTTP implements the http.Handler interface for processing receipt requests
func (h *ProcessReceiptHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Ensure only POST method is allowed
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON receipt from the request body
	var receipt models.Receipt
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&receipt); err != nil {
		// Return an error if the JSON is invalid
		http.Error(w, "Invalid receipt data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Validate the receipt data
	if err := validateReceipt(&receipt); err != nil {
		// Return validation errors
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save the receipt to the store and get its unique ID
	id := h.store.SaveReceipt(&receipt)

	// Prepare the response with the generated ID
	response := struct {
		ID string `json:"id"`
	}{
		ID: id,
	}

	// Set response headers and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// validateReceipt performs comprehensive validation on the receipt
func validateReceipt(receipt *models.Receipt) error {
	// Ensure the receipt has valid data
	if receipt.RetailerId == "" {
		return fmt.Errorf("retailer ID is required")
	}

	// Validate that PurchaseDate is a valid time (not zero)
	if receipt.PurchaseDate.IsZero() {
		return fmt.Errorf("purchase date is required")
	}

	// Ensure the receipt contains at least one item
	if len(receipt.Items) == 0 {
		return fmt.Errorf("receipt must contain at least one item")
	}

	return nil
}
