package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":8000", nil)
	// This line here would never be reached!
}
