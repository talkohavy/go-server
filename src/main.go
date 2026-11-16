package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

type Response = http.ResponseWriter
type Request = *http.Request

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

	mux.HandleFunc("/{$}", handleHello) // <--- this value: "/" also acts as the fallback route when no other route is matched. To avoid this, you can set the value to "/{$}". This means an exact match of the root path.
	mux.HandleFunc("/health-check", handleHealthCheck)
	mux.HandleFunc("/books", handleBooks)
	mux.HandleFunc("/books/{bookId}", handleBook)
	mux.HandleFunc("/api/login", handleLogin)

	log.Fatal(http.ListenAndServe(":8000", mux))
	// This line here would never be reached!
}

func handleHello(res Response, req Request) {
	slog.Info("[Hello] Incoming request!")

	helloResponse := "Hello, World!\n"

	wc, error := res.Write([]byte(helloResponse))

	if error != nil {
		slog.Error("Error writing response", "error", error)
		return
	}

	slog.Debug("bytes written", "bytes", wc)
}

func handleHealthCheck(res Response, req Request) {
	res.WriteHeader(http.StatusOK)

	okResponse := "OK"

	_, err := res.Write([]byte(okResponse))

	if err != nil {
		slog.Error("Error writing health check response", "error", err)
	}
}

func handleBooks(res Response, req Request) {
	params := req.URL.Query()

	page := params.Get(QueryKeys.Page)
	limit := params.Get(QueryKeys.Limit)

	if page == "" {
		page = "100"
	}

	if limit == "" {
		limit = "10"
	}

	slog.Info("[Books] Incoming request!", QueryKeys.Page, page, QueryKeys.Limit, limit)

	var booksResponse bytes.Buffer
	booksResponse.WriteString(fmt.Sprintf("Books. Page: %s, Limit: %s\n", page, limit))
	booksResponse.WriteString("Another line.!\n")
	booksResponse.WriteString("Done!\n")

	res.WriteHeader(http.StatusOK)

	_, err := res.Write(booksResponse.Bytes())

	if err != nil {
		slog.Error("Error writing books response", "error", err)
	}
}

func handleBook(res Response, req Request) {
	bookId := req.PathValue("bookId")

	if bookId == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	slog.Info("[Book] Incoming request!", "bookId", bookId)

	bookResponse := fmt.Sprintf("Book: %s\n", bookId)

	res.WriteHeader(http.StatusOK)

	_, err := res.Write([]byte(bookResponse))

	if err != nil {
		slog.Error("Error writing book response", "error", err)
	}
}

type UserData struct {
	name  string
	email string
	age   int
}

func handleLogin(res Response, req Request) {
	accessToken := req.Header.Get("cookie")

	slog.Info("[Login] incoming request!")

	if accessToken == "" {
		slog.Error("[Login] missing access token")

		http.Error(res, "access token required!", http.StatusBadRequest)

		return
	}

	slog.Info("[Login]", "cookie", accessToken)

	loginResponse := UserData{
		name:  "asd",
		email: "tal@gmail.com",
		age:   30,
	}

	res.Write([]byte(loginResponse.email))
}
