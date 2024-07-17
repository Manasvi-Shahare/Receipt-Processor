package server

import (
	"receipt-processor/internal/api/http/handlers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/receipts/process", handlers.ProcessReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", handlers.GetPoints).Methods("GET")
	return router
}
