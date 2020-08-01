package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/taciomcosta/chesstournament/internal/model"
	"github.com/taciomcosta/chesstournament/internal/shared"
)

func GetClubDetailsHandler(w http.ResponseWriter, r *http.Request) {
	v, err := service.GetClubById(getIdFromRequest(r))
	if ok := tryRespondWithError(w, http.StatusNotFound, err); ok {
		return
	}
	respond(w, v)
}

func CreateClubHandler(w http.ResponseWriter, r *http.Request) {
	club := readClubFromBody(r)
	_, err := service.CreateClub(club)
	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}
	w.WriteHeader(http.StatusCreated)
	respond(w, club)
}

func readClubFromBody(r *http.Request) *model.Club {
	club := new(model.Club)
	unmarshalJsonBody(r, club)
	return club
}

func unmarshalJsonBody(r *http.Request, v interface{}) {
	buffer, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	json.Unmarshal(buffer, v)
}

func EditClubHandler(w http.ResponseWriter, r *http.Request) {
	club := readClubFromBody(r)
	err := service.EditClub(getIdFromRequest(r), club)
	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}
	w.WriteHeader(http.StatusOK)
	respond(w, club)
}

func ListClubsHandler(w http.ResponseWriter, r *http.Request) {
	filter := newFilter(r)
	clubs, err := service.ListClubs(filter)
	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}
	respond(w, clubs)
}

func DeleteClubHandler(w http.ResponseWriter, r *http.Request) {
	club, err := service.DeleteClub(getIdFromRequest(r))
	if ok := tryRespondWithError(w, http.StatusNotFound, err); ok {
		return
	}
	respond(w, club)
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
	player := new(shared.CreatePlayerDTO)
	unmarshalJsonBody(r, player)
	return player
}

func GetPlayerDetailsHandler(w http.ResponseWriter, r *http.Request) {
	player, err := service.GetPlayerById(getIdFromRequest(r))
	if ok := tryRespondWithError(w, http.StatusNotFound, err); ok {
		return
	}
	respond(w, player)
}

func DeletePlayerHandler(w http.ResponseWriter, r *http.Request) {
	player, err := service.DeletePlayer(getIdFromRequest(r))
	if ok := tryRespondWithError(w, http.StatusNotFound, err); ok {
		return
	}
	respond(w, player)
}
