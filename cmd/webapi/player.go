package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/taciomcosta/chesstournament/internal/model"
)

func CreatePlayerHandler(w http.ResponseWriter, r *http.Request) {
	p, err := readPlayerFromBody(r)
	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}

	_, err = s.CreatePlayer(p)

	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}

	w.WriteHeader(http.StatusCreated)
	respond(w, p)
}

func readPlayerFromBody(r *http.Request) (*model.Player, error) {
	p := new(model.Player)
	b, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	err := json.Unmarshal(b, p)
	return p, err
}
