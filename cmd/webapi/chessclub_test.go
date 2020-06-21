package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/taciomcosta/chesstournament/internal/chessclub"
	"github.com/taciomcosta/chesstournament/internal/repository"
)

func TestMain(m *testing.M) {
	s = chessclub.NewService(&repository.MockChessClub{})
	os.Exit(m.Run())
}

func TestGetChessclubDetails(t *testing.T) {
	var tests = []struct {
		vars   map[string]string
		status int
	}{
		{map[string]string{"id": "1"}, http.StatusOK},
		{map[string]string{"id": "unexistent"}, http.StatusNotFound},
	}

	for _, tt := range tests {
		w, _ := http.NewRequest("GET", "/v1/chessclubs/10000", nil)
		w = mux.SetURLVars(w, tt.vars)
		r := httptest.NewRecorder()

		GetChessclubDetailsHandler(r, w)

		if r.Code != tt.status {
			t.Errorf("want status %v, got %v", r.Code, tt.status)
		}
	}
}

func TestCreateChessclub(t *testing.T) {
	var tests = []struct {
		body   string
		status int
	}{
		{`{"name": "name", "address": "address"}`, http.StatusCreated},
		{`{invalid: json}`, http.StatusBadRequest},
		{`{"name": "", "address": ""}`, http.StatusBadRequest},
	}

	for _, tt := range tests {
		body := strings.NewReader(tt.body)
		w, _ := http.NewRequest("POST", "/v1/chessclubs", body)
		r := httptest.NewRecorder()

		CreateChessclubHandler(r, w)

		if r.Code != tt.status {
			t.Errorf("it should return status code BadRequest, got %v", r.Code)
		}
	}
}
