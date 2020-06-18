package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taciomcosta/chesstournament/internal/config"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/chessclubs/{id}", GetChessclubDetailsHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(config.String("HOST"), nil)
}
