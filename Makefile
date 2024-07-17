.PHONY: build run test

build:
	go build -o receipt-processor ./cmd/app

run: build
	./receipt-processor

test:
	go test ./...
