package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/taciomcosta/chesstournament/internal/repository"
)

func TestRespond(t *testing.T) {
	s := struct {
		Hello string `json:"hello"`
	}{Hello: "world"}
	rr := httptest.NewRecorder()

	respond(rr, s)

	got := rr.Body.String()
	want := `{"hello":"world"}`
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestTryRespondWithError(t *testing.T) {
	t.Run("it should respond with error", testRespondWithError)
	t.Run("it should not respond to request", testDontRespondWithError)
}

func testRespondWithError(t *testing.T) {
	tests := []struct {
		body       string
		status     int
		wantStatus int
		err        error
		ok         bool
	}{
		{
			body:       `{"msg": "some_error"}`,
			status:     http.StatusBadRequest,
			wantStatus: http.StatusBadRequest,
			err:        errors.New("some_error"),
			ok:         true,
		},
		{
			body:       `{"msg": "An internal error has occurred"}`,
			status:     http.StatusBadRequest,
			wantStatus: http.StatusInternalServerError,
			err:        repository.InternalErr{},
			ok:         true,
		},
	}

	for _, tt := range tests {
		rr := httptest.NewRecorder()
		ok := tryRespondWithError(rr, tt.status, tt.err)

		if rr.Body.String() != tt.body {
			t.Errorf("Body: want %s, got %s", tt.body, rr.Body.String())
		}

		if rr.Result().StatusCode != tt.wantStatus {
			t.Errorf("Status: want %d, got %d", tt.wantStatus, rr.Result().StatusCode)
		}

		if tt.ok != ok {
			t.Errorf("OK: want %v, got %v", tt.ok, ok)
		}
	}

}

func testDontRespondWithError(t *testing.T) {
	rr := httptest.NewRecorder()
	got := tryRespondWithError(rr, http.StatusOK, nil)
	want := false
	if want != got {
		t.Errorf("OK: want %v, got %v", want, got)
	}
}
