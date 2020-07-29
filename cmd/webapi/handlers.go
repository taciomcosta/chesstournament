package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/taciomcosta/chesstournament/internal/model"
	"github.com/taciomcosta/chesstournament/internal/shared"
)

func GetClubDetailsHandler(w http.ResponseWriter, r *http.Request) {
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

func CreateClubHandler(w http.ResponseWriter, r *http.Request) {
	c := readClubFromBody(r)
	_, err := service.CreateClub(c)
	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}
	w.WriteHeader(http.StatusCreated)
	respond(w, c)
}

func readClubFromBody(r *http.Request) *model.Club {
	c := new(model.Club)
	unmarshalJsonBody(r, c)
	return c
}

func unmarshalJsonBody(r *http.Request, v interface{}) {
	b, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	json.Unmarshal(b, v)
}

func EditClubHandler(w http.ResponseWriter, r *http.Request) {
	c := readClubFromBody(r)
	err := service.EditClub(getIdFromRequest(r), c)
	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}
	w.WriteHeader(http.StatusOK)
	respond(w, c)
}

func ListClubsHandler(w http.ResponseWriter, r *http.Request) {
	f := newFilter(r)
	cs, err := service.ListClubs(f)
	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}
	respond(w, cs)
}

func DeleteClubHandler(w http.ResponseWriter, r *http.Request) {
	c, err := service.DeleteClub(getIdFromRequest(r))
	if ok := tryRespondWithError(w, http.StatusNotFound, err); ok {
		return
	}
	respond(w, c)
}

func CreatePlayerHandler(w http.ResponseWriter, r *http.Request) {
	dto := readCreatePlayerDTO(r)
	_, err := service.CreatePlayer(dto)
	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}
	w.WriteHeader(http.StatusCreated)
	respond(w, dto)
}

func readCreatePlayerDTO(r *http.Request) *shared.CreatePlayerDTO {
	p := new(shared.CreatePlayerDTO)
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
