package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/taciomcosta/chesstournament/internal/chessclub"
	"github.com/taciomcosta/chesstournament/internal/repository"
)

func TestMain(m *testing.M) {
	s = chessclub.NewService(&repository.MockChessClub{})
	os.Exit(m.Run())
}

func TestGetChessclubDetailsHandlerStatusOK(t *testing.T) {
	w, _ := http.NewRequest("GET", "/v1/chessclubs/1", nil)
	w = mux.SetURLVars(w, map[string]string{
		"id": "1",
	})
	r := httptest.NewRecorder()

	GetChessclubDetailsHandler(r, w)

	if r.Code != http.StatusOK {
		t.Error("it should return status code OK")
	}

	if r.Header().Get("Content-Type") != "application/json" {
		t.Error(`it should set header "Content-Type: application/json"`)
	}
}

func TestGetChessclubDetailsHandlerStatusNotFound(t *testing.T) {
	w, _ := http.NewRequest("GET", "/v1/chessclubs/10000", nil)
	r := httptest.NewRecorder()

	GetChessclubDetailsHandler(r, w)

	if r.Code != http.StatusNotFound {
		t.Errorf("it should return status code Not Found, got %v", r.Code)
	}
}
