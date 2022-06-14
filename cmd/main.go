package main

import (
	"github.com/alexstojda/alexstojda-ca/internal/http"
	"os"
)

func main() {
	spaPath := os.Getenv("SPA_PATH")
	if spaPath == "" {
		spaPath = "/app/build"
	}

	server := http.NewServer(spaPath)

	server.StartServer()
}
