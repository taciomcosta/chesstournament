package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreatePlayer(t *testing.T) {
	tests := []struct {
		body   string
		status int
	}{
		{
			body: `
				{
					"id": 1,
					"clubId": 1,
					"rankingCode": 1,
					"firstName": "Tacio",
					"lastName": "Costa",
					"address": "Somewhere",
					"phone": "12341234",
					"email": "tacio@email.com"
				}
			`,
			status: http.StatusCreated,
		},
		{
			body:   `{invalid: json}`,
			status: http.StatusBadRequest,
		},
		{
			body:   `{"id": 0}`,
			status: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		body := strings.NewReader(tt.body)
		w, _ := http.NewRequest("POST", "/players", body)
		r := httptest.NewRecorder()

		CreatePlayerHandler(r, w)

		if r.Code != tt.status {
			t.Errorf("want http status %v, got %v", tt.status, r.Code)
		}
	}
}
