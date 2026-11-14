package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHello(t *testing.T) {
	res := httptest.NewRecorder()

	handleHello(res, nil)

	actualCode := res.Code
	expectedCode := http.StatusOK
	actualBody := res.Body.String()
	expectedBody := "Hello, World!\n"

	if actualCode != expectedCode {
		t.Errorf("expected code %d, got %d\nBody: %s", expectedCode, actualCode, actualBody)
	}

	if actualBody != expectedBody {
		t.Errorf("expected body %s, got %s", expectedBody, actualBody)
	}
}
