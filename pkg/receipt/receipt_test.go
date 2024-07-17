package receipt

import (
	"testing"
)

func TestGenerateID(t *testing.T) {
	id := GenerateID()
	if id == "" {
		t.Errorf("Generated ID should not be empty")
	}
}

func TestCalculatePoints(t *testing.T) {
	receiptData := Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
		},
		Total: "18.74",
	}
	points := CalculatePoints(receiptData)
	expectedPoints := 20 // Points breakdown: 6 (retailer) + 5 (items) + 6 (odd day) + 3 (item description multiple of 3)

	if points != expectedPoints {
		t.Errorf("Calculated points are incorrect: got %v, want %v", points, expectedPoints)
	}
}

func TestStoreAndGetReceipt(t *testing.T) {
	receiptData := Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
		},
		Total: "18.74",
	}
	id := GenerateID()
	StoreReceipt(id, receiptData)
	storedReceipt, exists := receipts[id]

	if !exists {
		t.Errorf("Receipt was not stored correctly")
	}

	// Manually compare the fields
	if storedReceipt.Retailer != receiptData.Retailer ||
		storedReceipt.PurchaseDate != receiptData.PurchaseDate ||
		storedReceipt.PurchaseTime != receiptData.PurchaseTime ||
		storedReceipt.Total != receiptData.Total ||
		len(storedReceipt.Items) != len(receiptData.Items) {
		t.Errorf("Stored receipt does not match original")
	}

	for i := range storedReceipt.Items {
		if storedReceipt.Items[i] != receiptData.Items[i] {
			t.Errorf("Stored receipt items do not match original items")
		}
	}
}

func TestStoreAndGetPoints(t *testing.T) {
	receiptData := Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
		},
		Total: "18.74",
	}
	id := GenerateID()
	points := CalculatePoints(receiptData)
	StorePoints(id, points)
	storedPoints, err := GetPoints(id)

	if err != nil {
		t.Errorf("Error retrieving points: %v", err)
	}

	if storedPoints != points {
		t.Errorf("Stored points do not match calculated points: got %v, want %v", storedPoints, points)
	}
}
