package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/taciomcosta/chesstournament/internal/chessclub"
	"github.com/taciomcosta/chesstournament/internal/model"
)

func TestMain(m *testing.M) {
	service = chessclub.Mock()
}

func TestGetChessclubDetailsHandlerStatusOK(t *testing.T) {
	t.Error("must error")
	w, _ := http.NewRequest("GET", "/v1/chessclubs/1", nil)
	r := httptest.NewRecorder()

	GetChessclubDetailsHandler(r, w)

	if r.Code != http.StatusOK {
		t.Error("it should return status code OK")
	}

	if r.Header().Get("Content-Type") != "application/json" {
		t.Error(`it should set header "Content-Type: application/json"`)
	}

	expected := string(mustJSON(model.MockChessClub))
	if r.Body.String() != expected {
		t.Errorf("it should return json body %v, got %v", expected, r.Body.String())
	}
}

//func TestGetChessclubDetailsHandlerStatusNotFound(t *testing.T) {
//	w, _ := http.NewRequest("GET", "/v1/chessclubs/1", nil)
//	r := httptest.NewRecorder()
//
//	GetChessclubDetailsHandler(r, w)
//
//	if r.Code != http.StatusNotFound {
//		t.Error("it should return status code OK")
//	}
//
//}
