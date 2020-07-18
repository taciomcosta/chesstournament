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

func GetPlayerDetailsHandler(w http.ResponseWriter, r *http.Request) {
	f := func(id int) (interface{}, error) { return s.GetPlayerById(id) }
	getDetails(w, r, f)
}

type getDetailsFunc func(id int) (interface{}, error)

func getDetails(w http.ResponseWriter, r *http.Request, getDetails getDetailsFunc) {
	v, err := getDetails(id(r))
	if ok := tryRespondWithError(w, http.StatusNotFound, err); ok {
		return
	}
	respond(w, v)
}
