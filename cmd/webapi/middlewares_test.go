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
		response := requestUsingMiddleware(tt.urlPath, headersMiddleware)
		thenAssertResponseContentTypeIs(t, response, tt.contentType)
	}
}

func requestUsingMiddleware(urlPath string, middleware Middleware) http.ResponseWriter {
	request, _ := http.NewRequest("GET", urlPath, nil)
	recorder := httptest.NewRecorder()
	handler := middleware(mockHandler{})
	handler.ServeHTTP(recorder, request)
	return recorder
}

func thenAssertResponseContentTypeIs(t *testing.T, response http.ResponseWriter, contentType string) {
	got := response.Header().Get("Content-Type")
	if got != contentType {
		t.Errorf("want %s, got %s", contentType, got)
	}
}

func TestLoggerMiddleware(t *testing.T) {
	buffer := getLogOutputBuffer()
	requestUsingMiddleware("/urlpath", loggerMiddleware)
	thenAssertStringWasLoggedIntoBuffer(t, "GET/urlpath", buffer)
}

func getLogOutputBuffer() *bytes.Buffer {
	b := new(bytes.Buffer)
	log.SetOutput(b)
	return b
}

func thenAssertStringWasLoggedIntoBuffer(t *testing.T, expected string, buffer *bytes.Buffer) {
	if !strings.Contains(buffer.String(), expected) {
		t.Errorf("want %s logged, got %s", expected, buffer.String())
	}
}
