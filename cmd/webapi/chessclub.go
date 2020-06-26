package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/taciomcosta/chesstournament/internal/chessclub"
	"github.com/taciomcosta/chesstournament/internal/model"
	"github.com/taciomcosta/chesstournament/internal/repository"
)

var s *chessclub.Service

func init() {
	s = chessclub.NewService(repository.ChessClubRepository{})
}

func GetChessclubDetailsHandler(w http.ResponseWriter, r *http.Request) {
	id := getId(r)

	club, err := s.GetClubById(id)

	if ok := tryRespondWithError(w, http.StatusNotFound, err); ok {
		return
	}

	json := mustJSON(club)
	w.Write(json)
}

func getId(r *http.Request) int {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	return id
}

func tryRespondWithError(w http.ResponseWriter, httpStatus int, err error) bool {
	if err == nil {
		return false
	}
	w.WriteHeader(httpStatus)
	w.Write(errorResponse(err))
	return true
}

func errorResponse(err error) []byte {
	errString := fmt.Sprintf(`{"msg": "%s"}`, err)
	return []byte(errString)
}

func mustJSON(v interface{}) []byte {
	json, _ := json.Marshal(v)
	return json
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

	json := mustJSON(c)
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
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

	err = s.EditChessclub(getId(r), c)

	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}

	json := mustJSON(c)
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func ListChessclubsHandler(w http.ResponseWriter, r *http.Request) {
	f := getFilter(r)

	cs, err := s.ListClubs(f)

	if ok := tryRespondWithError(w, http.StatusBadRequest, err); ok {
		return
	}

	json := mustJSON(cs)
	w.Write(json)
}

func getFilter(r *http.Request) repository.Filter {
	var f repository.Filter
	r.ParseForm()
	schema.NewDecoder().Decode(&f, r.Form)
	return f
}

func DeleteChessclubHandler(w http.ResponseWriter, r *http.Request) {
	id := getId(r)

	c, err := s.DeleteClub(id)

	if ok := tryRespondWithError(w, http.StatusNotFound, err); ok {
		return
	}

	json := mustJSON(*c)
	w.Write(json)
}
