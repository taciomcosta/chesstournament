package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/taciomcosta/chesstournament/internal/chessclub"
	"github.com/taciomcosta/chesstournament/internal/repository"
)

var s *chessclub.Service

func init() {
	s = chessclub.NewService(repository.ChessClubRepository{})
}

func GetChessclubDetailsHandler(w http.ResponseWriter, r *http.Request) {
	id := getId(r)
	club, err := s.GetClubById(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorResponse(err))
	} else {
		json := mustJSON(*club)
		w.Write(json)
	}
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
