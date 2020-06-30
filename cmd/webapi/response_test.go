package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
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
	rr := httptest.NewRecorder()
	err := errors.New("some_error")

	wantBody := `{"msg": "some_error"}`
	wantStatus := http.StatusBadRequest
	gotOk := tryRespondWithError(rr, wantStatus, err)

	if rr.Body.String() != wantBody {
		t.Errorf("Body: want %s, got %s", wantBody, rr.Body.String())
	}

	if rr.Result().StatusCode != wantStatus {
		t.Errorf("Status: want %d, got %d", wantStatus, rr.Result().StatusCode)
	}

	wantOk := true
	if wantOk != gotOk {
		t.Errorf("OK: want %v, got %v", wantOk, gotOk)
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
