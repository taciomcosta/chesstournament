package webapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/taciomcosta/chesstournament/pkg/chessclub"
)

func TestGetChessclubDetailsHandlerStatusOK(t *testing.T) {
	w, _ := http.NewRequest("GET", "/chessclub/{id}", nil)
	r := httptest.NewRecorder()

	GetChessclubDetailsHandler(r, w)

	if r.Code != http.StatusOK {
		t.Error("it should return status code OK")
	}

	if r.Header().Get("Content-Type") != "application/json" {
		t.Error(`it should set header "Content-Type: application/json"`)
	}

	expected := string(mustJSON(chessclub.GetClubById(0)))
	if r.Body.String() != expected {
		t.Errorf("it should return json body %v, got %v", expected, r.Body.String())
	}
}
