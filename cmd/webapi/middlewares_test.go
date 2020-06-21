package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockHandler struct{}

func (m mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func TestHeadersMiddleware(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/chessclubs/10000", nil)
	rr := httptest.NewRecorder()
	h := headersMiddleware(mockHandler{})

	h.ServeHTTP(rr, r)

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Error(`it should set header "Content-Type: application/json"`)
	}
}
