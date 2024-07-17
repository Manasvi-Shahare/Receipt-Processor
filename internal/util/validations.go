package util

import (
	"regexp"
	"time"
	"receipt-processor/pkg/receipt"
)

func ValidateReceipt(rec receipt.Receipt) bool {
	retailerRegex := regexp.MustCompile(`^[\w\s\-&]+$`)
	if !retailerRegex.MatchString(rec.Retailer) {
		return false
	}

	if _, err := time.Parse("2006-01-02", rec.PurchaseDate); err != nil {
		return false
	}

	if _, err := time.Parse("15:04", rec.PurchaseTime); err != nil {
		return false
	}

	totalRegex := regexp.MustCompile(`^\d+\.\d{2}$`)
	if !totalRegex.MatchString(rec.Total) {
		return false
	}

	itemRegex := regexp.MustCompile(`^[\w\s\-]+$`)
	priceRegex := regexp.MustCompile(`^\d+\.\d{2}$`)
	for _, item := range rec.Items {
		if !itemRegex.MatchString(item.ShortDescription) || !priceRegex.MatchString(item.Price) {
			return false
		}
	}

	return true
}
