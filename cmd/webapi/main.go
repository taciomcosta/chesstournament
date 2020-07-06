package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taciomcosta/chesstournament/internal/config"
)

var swaggerURLPath = "/swagger"

func main() {
	r := mux.NewRouter().PathPrefix("/v1").Subrouter()
	addSwagger(r)
	addHandlers(r)
	addMiddlewares(r)
	log.Printf("Server listening on %s\n", config.String("HOST"))
	http.Handle("/", r)
	http.ListenAndServe(config.String("HOST"), nil)
}

func addSwagger(r *mux.Router) {
	fs := http.FileServer(http.Dir("./swagger/"))
	r.PathPrefix(swaggerURLPath).Handler(http.StripPrefix("/v1/swagger", fs))
}

func addHandlers(r *mux.Router) {
	r.HandleFunc("/chessclubs/{id}", GetChessclubDetailsHandler).Methods("GET")
	r.HandleFunc("/chessclubs", ListChessclubsHandler).Methods("GET")
	r.HandleFunc("/chessclubs", CreateChessclubHandler).Methods("POST")
	r.HandleFunc("/chessclubs/{id}", DeleteChessclubHandler).Methods("DELETE")
	r.HandleFunc("/chessclubs/{id}", EditChessclubHandler).Methods("PUT")
}

func addMiddlewares(r *mux.Router) {
	r.Use(loggerMiddleware)
	r.Use(headersMiddleware)
}
