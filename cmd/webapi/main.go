package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taciomcosta/chesstournament/internal/config"
)

var swaggerURLPath = "/swagger"

func main() {
	r := mux.NewRouter()
	addSwagger(r)
	addHandlers(r)
	addMiddlewares(r)
	fmt.Printf("Server listening on %s\n", config.String("HOST"))
	http.Handle("/", r)
	http.ListenAndServe(config.String("HOST"), nil)
}

func addSwagger(r *mux.Router) {
	fs := http.FileServer(http.Dir("./swagger/"))
	r.PathPrefix(swaggerURLPath).Handler(http.StripPrefix(swaggerURLPath, fs))
}

func addHandlers(r *mux.Router) {
	r.HandleFunc("/v1/chessclubs/{id}", GetChessclubDetailsHandler).Methods("GET")
	r.HandleFunc("/v1/chessclubs", ListChessclubsHandler).Methods("GET")
	r.HandleFunc("/v1/chessclubs", CreateChessclubHandler).Methods("POST")
	r.HandleFunc("/v1/chessclubs/{id}", DeleteChessclubHandler).Methods("DELETE")
	r.HandleFunc("/v1/chessclubs/{id}", EditChessclubHandler).Methods("PUT")
}

func addMiddlewares(r *mux.Router) {
	r.Use(headersMiddleware)
}
