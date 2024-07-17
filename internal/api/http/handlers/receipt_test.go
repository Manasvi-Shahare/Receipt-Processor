package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"receipt-processor/pkg/receipt"
	"testing"

	"github.com/gorilla/mux"
)

func TestProcessReceipt(t *testing.T) {
	receiptData := receipt.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []receipt.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
		},
		Total: "18.74",
	}
	receiptJSON, _ := json.Marshal(receiptData)

	req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(receiptJSON))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ProcessReceipt)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response receipt.ProcessResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.ID == "" {
		t.Errorf("handler returned an empty ID")
	}
}

func TestGetPoints(t *testing.T) {
	receiptData := receipt.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []receipt.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
		},
		Total: "18.74",
	}
	receiptID := receipt.GenerateID()
	receipt.StoreReceipt(receiptID, receiptData)
	points := receipt.CalculatePoints(receiptData)
	receipt.StorePoints(receiptID, points)

	req, err := http.NewRequest("GET", "/receipts/"+receiptID+"/points", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/receipts/{id}/points", GetPoints)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response receipt.PointsResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	expectedPoints := receipt.CalculatePoints(receiptData)
	if response.Points != expectedPoints {
		t.Errorf("handler returned wrong points: got %v want %v", response.Points, expectedPoints)
	}
}

func TestInvalidReceipt(t *testing.T) {
	invalidReceipt := receipt.Receipt{
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []receipt.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
		},
		Total: "6.49",
	}
	invalidReceiptJSON, _ := json.Marshal(invalidReceipt)

	req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(invalidReceiptJSON))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ProcessReceipt)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
