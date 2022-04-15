package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingStatusOK(t *testing.T) {
	r := getRouter()

	routes(r)

	req, _ := http.NewRequest("GET", "/health/ping", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		return statusOK
	})
}

func TestStatusPageOK(t *testing.T) {
	r := getRouter()

	templateFunctions(r)
	routes(r)

	req, _ := http.NewRequest("GET", "/health/status", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		return statusOK
	})
}

func TestDBStatusOK(t *testing.T) {
	r := getRouter()

	templateFunctions(r)
	routes(r)
	Connect()

	req, _ := http.NewRequest("GET", "/health/db", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		return statusOK
	})
}
