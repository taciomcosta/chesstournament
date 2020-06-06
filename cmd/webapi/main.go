package main

import (
	"net/http"

	"github.com/gorilla/mux"
	webapi "github.com/taciomcosta/chesstournament/cmd/webapi/handlers"
)

var address string = ":8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/chessclubs/{id}", webapi.GetChessclubDetailsHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(address, nil)
}
