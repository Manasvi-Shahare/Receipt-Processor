package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor/internal/util"
	"receipt-processor/pkg/receipt"
	"github.com/gorilla/mux"
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var rec receipt.Receipt
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil || !util.ValidateReceipt(rec) {
		http.Error(w, "Invalid receipt", http.StatusBadRequest)
		return
	}
	id := receipt.GenerateID()
	receipt.StoreReceipt(id, rec)
	points := receipt.CalculatePoints(rec)
	receipt.StorePoints(id, points)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(receipt.ProcessResponse{ID: id})
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	points, err := receipt.GetPoints(id)
	if err != nil {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(receipt.PointsResponse{Points: points})
}
