package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/taciomcosta/chesstournament/internal/model"
)

func GetChessclubDetailsHandler(w http.ResponseWriter, r *http.Request) {
	f := func(id int) (interface{}, error) { return service.GetClubById(id) }
	getDetails(w, r, f)
}

type getDetailsFunc func(id int) (interface{}, error)

func getDetails(w http.ResponseWriter, r *http.Request, get getDetailsFunc) {
	v, err := get(getIdFromRequest(r))
	if ok := tryRespondWithError(w, http.StatusNotFound, err); ok {
		return
	}
	respond(w, v)
}

func CreateChessclubHandler(w http.ResponseWriter, r *http.Request) {
	c := readChessclubFromBody(r)
	_, err := service.CreateChessclub(c)
	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}
	w.WriteHeader(http.StatusCreated)
	respond(w, c)
}

func readChessclubFromBody(r *http.Request) *model.ChessClub {
	c := new(model.ChessClub)
	unmarshalJsonBody(r, c)
	return c
}

func unmarshalJsonBody(r *http.Request, v interface{}) {
	b, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	json.Unmarshal(b, v)
}

func EditChessclubHandler(w http.ResponseWriter, r *http.Request) {
	c := readChessclubFromBody(r)
	err := service.EditChessclub(getIdFromRequest(r), c)
	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}
	w.WriteHeader(http.StatusOK)
	respond(w, c)
}

func ListChessclubsHandler(w http.ResponseWriter, r *http.Request) {
	f := newFilter(r)
	cs, err := service.ListClubs(f)
	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}
	respond(w, cs)
}

func DeleteChessclubHandler(w http.ResponseWriter, r *http.Request) {
	c, err := service.DeleteClub(getIdFromRequest(r))
	if ok := tryRespondWithError(w, http.StatusNotFound, err); ok {
		return
	}
	respond(w, c)
}

func CreatePlayerHandler(w http.ResponseWriter, r *http.Request) {
	p := readPlayerFromBody(r)
	_, err := service.CreatePlayer(p)
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

func GetPlayerDetailsHandler(w http.ResponseWriter, r *http.Request) {
	f := func(id int) (interface{}, error) { return service.GetPlayerById(id) }
	getDetails(w, r, f)
}

func DeletePlayerHandler(w http.ResponseWriter, r *http.Request) {
	player, err := service.DeletePlayer(getIdFromRequest(r))
	if ok := tryRespondWithError(w, http.StatusNotFound, err); ok {
		return
	}
	respond(w, player)
}
