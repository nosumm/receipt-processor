// Receipt struct definitions

package models
import "time"

type Item struct {
    ShortDescription string `json:"shortDescription"`
    Price           string `json:"price"`
}

type Receipt struct {
    Retailer       string `json:"retailer"`
    PurchaseDate   time.Time `json:"purchaseDate"`
    PurchaseTime   time.Time `json:"purchaseTime"`
    Items          []Item `json:"items"`
    Total          string `json:"total"`
    RetailerId     string `json:"retailerId"`

}


