package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/taciomcosta/chesstournament/internal/data"
	"github.com/taciomcosta/chesstournament/internal/shared"
)

func TestMain(m *testing.M) {
	s = shared.NewService(&data.MockChessClubRepository{})
	os.Exit(m.Run())
}

func TestGetChessclubDetails(t *testing.T) {
	tests := []struct {
		vars   map[string]string
		status int
	}{
		{map[string]string{"id": "1"}, http.StatusOK},
		{map[string]string{"id": "unexistent"}, http.StatusNotFound},
	}

	for _, tt := range tests {
		w, _ := http.NewRequest("GET", "/chessclubs/10000", nil)
		w = mux.SetURLVars(w, tt.vars)
		r := httptest.NewRecorder()

		GetChessclubDetailsHandler(r, w)

		if r.Code != tt.status {
			t.Errorf("want status %v, got %v", tt.status, r.Code)
		}
	}
}

func TestCreateChessclub(t *testing.T) {
	tests := []struct {
		body   string
		status int
	}{
		{`{"name": "name", "address": "address"}`, http.StatusCreated},
		{`{invalid: json}`, http.StatusBadRequest},
		{`{"name": "", "address": ""}`, http.StatusBadRequest},
	}

	for _, tt := range tests {
		body := strings.NewReader(tt.body)
		w, _ := http.NewRequest("POST", "/chessclubs", body)
		r := httptest.NewRecorder()

		CreateChessclubHandler(r, w)

		if r.Code != tt.status {
			t.Errorf("it should return status code BadRequest, got %v", r.Code)
		}
	}
}

func TestEditChessclub(t *testing.T) {
	tests := []struct {
		body   string
		id     string
		status int
	}{
		{`{"name": "name", "address": "address"}`, "1", http.StatusOK},
		{`{invalid: json}`, "1", http.StatusBadRequest},
		{`{"name": "", "address": ""}`, "1", http.StatusBadRequest},
	}

	for _, tt := range tests {
		body := strings.NewReader(tt.body)
		w, _ := http.NewRequest("PUT", "/chessclubs/1", body)
		w = mux.SetURLVars(w, map[string]string{"id": "1"})
		r := httptest.NewRecorder()

		EditChessclubHandler(r, w)

		if r.Code != tt.status {
			t.Errorf("want status %d, got %d", tt.status, r.Code)
		}
	}

}

func TestListChessclubs(t *testing.T) {
	tests := []struct {
		queryParams string
		status      int
	}{
		{"", http.StatusOK},
		{"$orderBy=invalid", http.StatusBadRequest},
	}

	for _, tt := range tests {
		url := "/chessclubs?" + tt.queryParams
		w, _ := http.NewRequest("GET", url, nil)
		r := httptest.NewRecorder()

		ListChessclubsHandler(r, w)

		if r.Code != tt.status {
			t.Errorf("want status %v, got %v", tt.status, r.Code)
		}
	}
}

func TestDeleteChessclub(t *testing.T) {
	tests := []struct {
		vars   map[string]string
		status int
	}{
		{map[string]string{"id": "1"}, http.StatusOK},
		{map[string]string{"id": "-1"}, http.StatusNotFound},
	}

	for _, tt := range tests {
		w, _ := http.NewRequest("DELETE", "/chessclubs/10000", nil)
		w = mux.SetURLVars(w, tt.vars)
		r := httptest.NewRecorder()

		DeleteChessclubHandler(r, w)

		if r.Code != tt.status {
			t.Errorf("want status %v, got %v", r.Code, tt.status)
		}
	}
}
