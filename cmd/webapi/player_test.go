package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/taciomcosta/chesstournament/internal/data"
)

func TestGetPlayerDetails(t *testing.T) {
	testGetDetailsHandler(t, GetPlayerDetailsHandler)
}

func TestGetUnexistinPlayerDetails(t *testing.T) {
	testGetUnexistentDetailsHandler(t, GetPlayerDetailsHandler)
}

func TestCreatePlayer(t *testing.T) {
	request := newRequestBuilder().withBody(toJSONString(data.MockValidPlayer)).build()
	recorder := httptest.NewRecorder()

	CreatePlayerHandler(recorder, request)

	thenAssertRecorderStatusIs(t, recorder, http.StatusCreated)
}

func TestCreateInvalidPlayer(t *testing.T) {
	request := newRequestBuilder().withBody(`{"invalid": "body"}`).build()
	recorder := httptest.NewRecorder()

	CreatePlayerHandler(recorder, request)

	thenAssertRecorderStatusIs(t, recorder, http.StatusBadRequest)
}
