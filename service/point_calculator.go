// Calculate Receipt Points

package service

import (
    "math"
    "strconv"
    "strings"
    "time"
    "unicode"

    "receipt-processor/models" // import path
)

// calc points logic
func CalculatePoints(receipt *models.Receipt) int {
	// initialize point counter at 0
	points := 0
	// add points for every alphanumeric character in retailer name
	points += countAlphanumericChars(receipt.Retailer)
	// add 50 points for round dollar total
	if isRoundDollarAmount(receipt.Total){
		points +=50
	}
	// add 25 points if total is a multiple of 0.25
	if isMultipleof025(receipt.Total){
		points += 25
	}
	
	// add 5 points for every 2 items
	points += (len(receipt.Items) / 2) *5

	// item description points
	for _, item := range receipt.Items {
		points += calculateDescriptionPoints(item)
	}

	// add 6 points if purchage day is odd
	if isPurchaseDayOdd(receipt.PurchaseDate){
		points += 6
	}

	// add 10 points if purchase time is between 2-4 PM
	if isPurchaseTimeInRange(receipt.PurchaseTime){
		points += 10
	}

	return points

}

// HELPER FUNCTIONS

// counts and returns number of alphannumeric chars in a string
func countAlphanumericChars(s string) int {
    count := 0
    for _, ch := range s {
        if IsLetterorNumber(ch) {
            count++
        }
    }
    return count
}

func IsLetterorNumber(r rune) bool {
    return unicode.IsLetter(r) || unicode.IsNumber(r)
}

func isRoundDollarAmount(total string) bool {
    amount, err := strconv.ParseFloat(total, 64)
    if err != nil {
        return false
    }
    return math.Mod(amount, 1.0) == 0
}

func isMultipleof025(total string) bool {
    amount, err := strconv.ParseFloat(total, 64)
    if err != nil {
        return false
    }
    return math.Mod(amount*4, 1.0) == 0
}


func calculateDescriptionPoints(item models.Item) int {
    // Remove leading/trailing whitespace from description and check length
    description := strings.TrimSpace(item.ShortDescription)
    if len(description)%3 == 0 {
        price, err := strconv.ParseFloat(item.Price, 64)
        if err != nil {
            return 0
        }
        return int(math.Ceil(price * 0.2))
    }
    return 0
}

/*
func isPurchaseDayOdd(dateStr string) bool {
    date, err := time.Parse("2006-01-02", dateStr)
    if err != nil {
        return false
    }
    return date.Day()%2 != 0
}
*/
func isPurchaseDayOdd(purchaseDate time.Time) bool {
    day := purchaseDate.Day() // Get the day of the month
    return day%2 != 0
}

/*
func isPurchaseTimeInRange(timeStr string) bool {
    t, err := time.Parse("15:04", timeStr)
    if err != nil {
        return false
    }
    
    startTime, _ := time.Parse("15:04", "14:00")
    endTime, _ := time.Parse("15:04", "16:00")
    
    return t.After(startTime) && t.Before(endTime)
}
*/


// isPurchaseTimeInRange checks if the given time is within the specified range
func isPurchaseTimeInRange(purchaseTime time.Time) bool {
	// Define the start and end times (only time part matters here)
	startTime := time.Date(1, 1, 1, 14, 0, 0, 0, time.UTC) // 14:00
	endTime := time.Date(1, 1, 1, 16, 0, 0, 0, time.UTC)   // 16:00
	
	// Check if the purchase time is after the start time and before the end time
	return purchaseTime.After(startTime) && purchaseTime.Before(endTime)
}