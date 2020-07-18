package main

import (
	"net/http"

	"github.com/taciomcosta/chesstournament/internal/model"
)

func GetChessclubDetailsHandler(w http.ResponseWriter, r *http.Request) {
	f := func(id int) (interface{}, error) { return s.GetClubById(id) }
	getDetails(w, r, f)
}

func CreateChessclubHandler(w http.ResponseWriter, r *http.Request) {
	c := readChessclubFromBody(r)
	_, err := s.CreateChessclub(c)
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

func EditChessclubHandler(w http.ResponseWriter, r *http.Request) {
	c := readChessclubFromBody(r)

	err := s.EditChessclub(getIdFromRequest(r), c)

	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}

	w.WriteHeader(http.StatusOK)
	respond(w, c)
}

func ListChessclubsHandler(w http.ResponseWriter, r *http.Request) {
	f := newFilter(r)
	cs, err := s.ListClubs(f)
	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}
	respond(w, cs)
}

func DeleteChessclubHandler(w http.ResponseWriter, r *http.Request) {
	c, err := s.DeleteClub(getIdFromRequest(r))
	if ok := tryRespondWithError(w, http.StatusNotFound, err); ok {
		return
	}
	respond(w, c)
}
