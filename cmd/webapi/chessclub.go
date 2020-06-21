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

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorResponse(err))
		return
	}

	json := mustJSON(*club)
	w.Write(json)
}

func getId(r *http.Request) int {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	return id
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
	b, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	var c model.ChessClub
	if err := json.Unmarshal(b, &c); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorResponse(err))
		return
	}

	_, err := s.CreateChessclub(&c)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorResponse(err))
		return
	}

	json := mustJSON(c)
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}

func ListChessclubsHandler(w http.ResponseWriter, r *http.Request) {
	f := getFilter(r)

	cs, err := s.ListClubs(f)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorResponse(err))
		return
	}

	json := mustJSON(cs)
	w.Write(json)
}

func getFilter(r *http.Request) repository.ListFilter {
	var f repository.ListFilter
	r.ParseForm()
	schema.NewDecoder().Decode(&f, r.Form)
	return f
}
