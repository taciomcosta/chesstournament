package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taciomcosta/chesstournament/internal/config"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	addHandlers(r)
	addMiddlewares(r)
	fmt.Printf("Server listening on %s\n", config.String("HOST"))
	http.ListenAndServe(config.String("HOST"), nil)
}

func addHandlers(r *mux.Router) {
	r.HandleFunc("/v1/chessclubs/{id}", GetChessclubDetailsHandler).Methods("GET")
	r.HandleFunc("/v1/chessclubs", CreateChessclubHandler).Methods("POST")
}

func addMiddlewares(r *mux.Router) {
	r.Use(headersMiddleware)
}
