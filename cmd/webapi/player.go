package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/taciomcosta/chesstournament/internal/model"
)

func CreatePlayerHandler(w http.ResponseWriter, r *http.Request) {
	p := readPlayerFromBody(r)
	_, err := s.CreatePlayer(p)
	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}
	w.WriteHeader(http.StatusCreated)
	respond(w, p)
}

func readPlayerFromBody(r *http.Request) *model.Player {
	p := new(model.Player)
	unmarshalJsonBody(r, p)
	return p
}

func unmarshalJsonBody(r *http.Request, v interface{}) {
	b, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	json.Unmarshal(b, v)
}
