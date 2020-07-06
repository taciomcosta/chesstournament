package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/taciomcosta/chesstournament/internal/data"
	"github.com/taciomcosta/chesstournament/internal/model"
	"github.com/taciomcosta/chesstournament/internal/shared"
)

var s *shared.Service

func init() {
	s = shared.NewService(data.ChessClubRepository{})
}

func GetChessclubDetailsHandler(w http.ResponseWriter, r *http.Request) {
	c, err := s.GetClubById(id(r))
	if ok := tryRespondWithError(w, http.StatusNotFound, err); ok {
		return
	}
	respond(w, c)
}

func CreateChessclubHandler(w http.ResponseWriter, r *http.Request) {
	c, err := readChessclubFromBody(r)
	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}

	_, err = s.CreateChessclub(c)

	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}

	w.WriteHeader(http.StatusCreated)
	respond(w, c)
}

func readChessclubFromBody(r *http.Request) (*model.ChessClub, error) {
	c := new(model.ChessClub)
	b, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	err := json.Unmarshal(b, c)
	return c, err
}

func EditChessclubHandler(w http.ResponseWriter, r *http.Request) {
	c, err := readChessclubFromBody(r)
	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}

	err = s.EditChessclub(id(r), c)

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
	c, err := s.DeleteClub(id(r))
	if ok := tryRespondWithError(w, http.StatusNotFound, err); ok {
		return
	}
	respond(w, c)
}
