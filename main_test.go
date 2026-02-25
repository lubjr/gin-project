package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"hello/internal/router"
)

func TestHelloRoute(t *testing.T) {
	r := router.Setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 but got %d", w.Code)
	}
}

func TestHealthyRoute(t *testing.T) {
	r := router.Setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 but got %d", w.Code)
	}
}
