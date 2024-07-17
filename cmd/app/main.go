package main

import (
	"log"
	"net/http"
	"receipt-processor/internal/api/http/server"
)

func main() {
	router := server.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
