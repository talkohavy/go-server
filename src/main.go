package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", nil))
	// This line here would never be reached!
}
