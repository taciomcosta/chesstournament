package main

import (
	"encoding/json"
	"fmt"
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
	id := getId(r)
	club, err := service.GetClubById(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(mustErrorJSON(err))
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

func mustErrorJSON(err error) []byte {
	errString := fmt.Sprintf(`{"code": "%T", "msg": "%s"}`, err, err)
	return []byte(errString)
}

func mustJSON(v interface{}) []byte {
	json, _ := json.Marshal(v)
	return json
}
