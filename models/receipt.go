// Receipt struct definitions

package models

type Receipt struct {
    Retailer       string `json:"retailer"`
    PurchaseDate   string `json:"purchaseDate"`
    PurchaseTime   string `json:"purchaseTime"`
    Items          []Item `json:"items"`
    Total          string `json:"total"`
}


