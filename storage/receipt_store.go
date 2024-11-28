// thread safe in-memory receipt storage solution

package storage

import (
    "sync"
    "github.com/google/uuid"
    "receipt-processor/models"
    "receipt-processor/service"
)

type ReceiptStore struct {
	mu sync.RWMutex // a read write mutex to safely handle concurrent access to the receipts map
	receipts map[string]receiptEntry // receipts map - stores receipt entries with string keys (UUIDs) and receiptEntry values
}


type receiptEntry struct {
	Receipt *models.Receipt // a pointer to the receipt model
	Points int 				// num of points calculated for this receipt
}

// constructor function - creates a new ReceiptStore
// initializes the receipts map to prevent nil map errors
func NewReceiptStore() *ReceiptStore {
    return &ReceiptStore{
        receipts: make(map[string]receiptEntry),
    }
}

// This thread safe method saves a receipt to the store and returns generated ID
func (s *ReceiptStore) SaveReceipt(receipt *models.Receipt) string {
    s.mu.Lock() 			// prevents concurrent writes 
    defer s.mu.Unlock()		// release the lock when method exits

    // Generate a unique ID
    id := uuid.New().String()
    
    // Calculate points
    points := service.CalculatePoints(receipt)

    // Store the receipt with its points
    s.receipts[id] = receiptEntry{
        Receipt: receipt,
        Points:  points,
    }

    return id
}


// This thread safe method retrieves points for a specific receipt
func (s *ReceiptStore) GetReceiptPoints(id string) (int, bool) {
    s.mu.RLock()			// for concurrent read access
    defer s.mu.RUnlock()	// release read lock when method exists 

	// check if the receipt exists in the map
    entry, exists := s.receipts[id]
    if !exists {
        return 0, false
    }
    return entry.Points, true
}



