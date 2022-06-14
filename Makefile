
mod:
	go mod download

build:
	go build -v -o alexstojda-ca ./cmd

run:
	go run cmd/main.go