package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

type QueryKeysType struct {
	Page  string
	Limit string
}

var QueryKeys = QueryKeysType{
	Page:  "page",
	Limit: "limit",
}

func main() {
	slog.Info("Starting server on port 8000")

	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHello)
	mux.HandleFunc("/health-check", handleHealthCheck)
	mux.HandleFunc("/books", handleBooks)

	log.Fatal(http.ListenAndServe(":8000", mux))
	// This line here would never be reached!
}

func handleHello(res http.ResponseWriter, req *http.Request) {
	slog.Info("Incoming request!")

	wc, error := res.Write([]byte("Hello, World!\n"))

	if error != nil {
		slog.Error("Error writing response", "error", error)
		return
	}

	slog.Debug("bytes written", "bytes", wc)
}

func handleHealthCheck(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)

	_, err := res.Write([]byte("OK"))

	if err != nil {
		slog.Error("Error writing health check response", "error", err)
	}
}

func handleBooks(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()

	page := params.Get(QueryKeys.Page)
	limit := params.Get(QueryKeys.Limit)
	answer := fmt.Sprintf("Books. Page: %s, Limit: %s", page, limit)

	slog.Debug("books query", QueryKeys.Page, page, QueryKeys.Limit, limit)

	res.WriteHeader(http.StatusOK)

	_, err := res.Write([]byte(answer))

	if err != nil {
		slog.Error("Error writing books response", "error", err)
	}
}
