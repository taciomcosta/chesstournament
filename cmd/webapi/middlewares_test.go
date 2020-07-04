package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
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
		rr := requestWithMiddleware(headersMiddleware, tt.urlPath)
		if got := rr.Header().Get("Content-Type"); got != tt.contentType {
			t.Errorf("want %s, got %s", tt.contentType, got)
		}
	}
}

func requestWithMiddleware(m Middleware, urlPath string) *httptest.ResponseRecorder {
	r, _ := http.NewRequest("GET", urlPath, nil)
	rr := httptest.NewRecorder()
	h := m(mockHandler{})
	h.ServeHTTP(rr, r)
	return rr
}

func TestLoggerMiddleware(t *testing.T) {
	b := new(bytes.Buffer)
	log.SetOutput(b)

	want := "GET/urlpath"
	requestWithMiddleware(loggerMiddleware, "/urlpath")

	if !strings.Contains(b.String(), want) {
		t.Errorf("want %s, got %s", want, b.String())
	}
}
