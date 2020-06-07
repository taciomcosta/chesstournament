package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var address string = ":8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/chessclubs/{id}", GetChessclubDetailsHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(address, nil)
}
