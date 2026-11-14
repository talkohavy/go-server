package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("Starting server on port 8000")

	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHello)
	mux.HandleFunc("/health-check", handleHealthCheck)

	log.Fatal(http.ListenAndServe(":8000", mux))
	// This line here would never be reached!
}

func handleHello(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Incoming request!")

	wc, error := res.Write([]byte("Hello, World!\n"))

	if error != nil {
		slog.Error("Error writing response", "error", error)
		return
	}

	fmt.Printf("%d bytes written\n", wc)
}

func handleHealthCheck(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)

	_, err := res.Write([]byte("OK"))

	if err != nil {
		slog.Error("Error writing health check response", "error", err)
	}
}
