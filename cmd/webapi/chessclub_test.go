package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/taciomcosta/chesstournament/internal/data"
	"github.com/taciomcosta/chesstournament/internal/shared"
)

func TestMain(m *testing.M) {
	s = shared.NewService(&data.MockRepository{}, &data.MockChessClubRepository{}, data.MockPlayerRepository{})
	os.Exit(m.Run())
}

func TestGetChessclubDetails(t *testing.T) {
	testGetDetailsHandler(t, GetChessclubDetailsHandler)
}

func TestGetUnexistingChessclubDetails(t *testing.T) {
	testGetUnexistentDetailsHandler(t, GetChessclubDetailsHandler)
}

func testGetDetailsHandler(t *testing.T, handle http.HandlerFunc) {
	request := newRequestBuilder().withPathVar("id", "1").build()
	recorder := httptest.NewRecorder()

	handle(recorder, request)

	thenAssertRecorderStatusIs(t, recorder, http.StatusOK)
}

func testGetUnexistentDetailsHandler(t *testing.T, handle http.HandlerFunc) {
	request := newRequestBuilder().withPathVar("id", "unexistent").build()
	recorder := httptest.NewRecorder()

	handle(recorder, request)

	thenAssertRecorderStatusIs(t, recorder, http.StatusNotFound)
}

func thenAssertRecorderStatusIs(t *testing.T, recorder *httptest.ResponseRecorder, status int) {
	if recorder.Code != status {
		t.Errorf("want status %v, got %v", status, recorder.Code)
	}
}

func TestCreateChessclub(t *testing.T) {
	request := newRequestBuilder().withBody(toJSONString(data.MockValidChessClub)).build()
	recorder := httptest.NewRecorder()

	CreateChessclubHandler(recorder, request)

	thenAssertRecorderStatusIs(t, recorder, http.StatusCreated)
}

func TestCreateInvalidChessclub(t *testing.T) {
	request := newRequestBuilder().withBody(`{"invalid": "body"}`).build()
	recorder := httptest.NewRecorder()

	CreateChessclubHandler(recorder, request)

	thenAssertRecorderStatusIs(t, recorder, http.StatusBadRequest)
}

func TestEditChessclub(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := newRequestBuilder().
		withBody(toJSONString(data.MockValidChessClub)).
		withPathVar("id", "1").build()

	EditChessclubHandler(recorder, request)

	thenAssertRecorderStatusIs(t, recorder, http.StatusOK)
}

func TestEditChessclubInvalidBody(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := newRequestBuilder().
		withBody(toJSONString(`{"id":"1"}`)).
		withPathVar("id", "1").build()

	EditChessclubHandler(recorder, request)

	thenAssertRecorderStatusIs(t, recorder, http.StatusBadRequest)
}

func TestListChessclubs(t *testing.T) {
	request := newRequestBuilder().build()
	recorder := httptest.NewRecorder()

	ListChessclubsHandler(recorder, request)

	thenAssertRecorderStatusIs(t, recorder, http.StatusOK)
}

func TestListChessclubInvalidFilter(t *testing.T) {
	request := newRequestBuilder().withQueryParam("$orderBy", "invalid").build()
	recorder := httptest.NewRecorder()

	ListChessclubsHandler(recorder, request)

	thenAssertRecorderStatusIs(t, recorder, http.StatusBadRequest)
}

func TestDeleteChessclub(t *testing.T) {
	request := newRequestBuilder().withPathVar("id", "1").build()
	recorder := httptest.NewRecorder()

	DeleteChessclubHandler(recorder, request)

	thenAssertRecorderStatusIs(t, recorder, http.StatusOK)
}

func TestDeleteUnexistentChessclub(t *testing.T) {
	request := newRequestBuilder().withPathVar("id", "-1").build()
	recorder := httptest.NewRecorder()

	DeleteChessclubHandler(recorder, request)

	thenAssertRecorderStatusIs(t, recorder, http.StatusNotFound)
}
