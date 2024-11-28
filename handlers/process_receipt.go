// HTTP route handler for processing receipts

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"receipt-processor/models"
	"receipt-processor/service"
	"receipt-processor/storage"
)

// processes incoming receipts
// uses a receipt store to save and manage receipts
type ProcessReceiptHandler struct {
	store *storage.ReceiptStore
}

// creates a new handler for processing receipts
// takes a receipt store as a parameter to manage receipt storage
func NewProcessReceiptHandler(store *storage.ReceiptStore) *ProcessReceiptHandler {
	return &ProcessReceiptHandler{
		store: store,
	}
}

// ServeHTTP implements the http.Handler interface for processing receipt requests
// It handles the following responsibilities:
// 1. Validate the HTTP method is POST
// 2. Decode the incoming JSON receipt
// 3. Validate the receipt data
// 4. Save the receipt to the store
// 5. Return the generated receipt ID
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
// checks for:
// 1. Non-nil receipt
// 2. Required retailer ID
// 3. Valid purchase date
// 4. Presence of items
// Returns an error if any validation fails
func validateReceipt(receipt *models.Receipt) error {
	// Check if receipt is nil
	if receipt == nil {
		return fmt.Errorf("receipt cannot be nil")
	}
	
	// Validate retailer ID is present
	if receipt.RetailerId == "" {
		return fmt.Errorf("retailer ID is required")
	}

	// Validate purchase date is not zero value
	if receipt.PurchaseDate.IsZero() {
		return fmt.Errorf("purchase date is required")
	}

	// Ensure receipt has at least one item
	if len(receipt.Items) == 0 {
		return fmt.Errorf("receipt must contain at least one item")
	}

	// Additional validation can be added here
	// For example, check for valid total amount, item prices, etc.

	return nil
}

