package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/taciomcosta/chesstournament/internal/chessclub"
)

var service chessclub.Service

func init() {
	service = chessclub.New()
}

func GetChessclubDetailsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	club, err := service.GetClubById(id)

	if err != nil {
		http.Error(w, "resource not found", http.StatusNotFound)
		return
	}

	json := mustJSON(*club)
	w.Write(json)
}

func mustJSON(v interface{}) []byte {
	json, _ := json.Marshal(v)
	return json
}
