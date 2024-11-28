package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"receipt-processor/storage"
)

// NewGetPointsHandler creates a new handler for retrieving receipt points
// It takes a receipt store as a parameter to access stored receipts
func NewGetPointsHandler(store *storage.ReceiptStore) http.Handler {
	return &GetPointsHandler{
		store: store,
	}
}

// GetPointsHandler manages the retrieval of points for a specific receipt
// It uses a receipt store to look up receipt points
type GetPointsHandler struct {
	store *storage.ReceiptStore
}

// ServeHTTP implements the http.Handler interface for retrieving receipt points
func (h *GetPointsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Ensure only GET method is allowed
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the receipt ID from the URL parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Attempt to retrieve points for the specified receipt
	points, found := h.store.GetReceiptPoints(id)
	if !found {
		// Return a not found error if the receipt doesn't exist
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	// Prepare the response with the retrieved points
	response := struct {
		Points int `json:"points"`
	}{
		Points: points,
	}

	// Set response headers and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
