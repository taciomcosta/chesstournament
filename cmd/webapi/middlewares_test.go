package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockHandler struct{}

func (m mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func TestHeadersMiddleware(t *testing.T) {
	tests := []struct {
		urlPath     string
		contentType string
	}{
		{"/v1/chessclubs/10000", "application/json"},
		{"/swagger", ""},
		{"/swagger/", ""},
	}

	for _, tt := range tests {
		r, _ := http.NewRequest("GET", tt.urlPath, nil)
		rr := httptest.NewRecorder()
		h := headersMiddleware(mockHandler{})

		h.ServeHTTP(rr, r)

		if got := rr.Header().Get("Content-Type"); got != tt.contentType {
			t.Errorf("want %s, got %s", tt.contentType, got)
		}
	}
}
