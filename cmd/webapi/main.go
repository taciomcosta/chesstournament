package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taciomcosta/chesstournament/internal/config"
	"github.com/taciomcosta/chesstournament/internal/data"
	"github.com/taciomcosta/chesstournament/internal/shared"
)

var service shared.Service = shared.NewService(data.Repository{}, data.ChessClubRepository{}, data.PlayerRepository{})
var swaggerURLPath = "/swagger"
var router *mux.Router = mux.NewRouter().PathPrefix("/v1").Subrouter()

func main() {
	addSwagger(router)
	addHandlers(router)
	addMiddlewares(router)
	serve()
}

func addSwagger(router *mux.Router) {
	fs := http.FileServer(http.Dir("./swagger/"))
	router.PathPrefix(swaggerURLPath).Handler(http.StripPrefix("/v1/swagger", fs))
}

func addHandlers(router *mux.Router) {
	router.HandleFunc("/chessclubs/{id}", GetChessclubDetailsHandler).Methods("GET")
	router.HandleFunc("/chessclubs", ListChessclubsHandler).Methods("GET")
	router.HandleFunc("/chessclubs", CreateChessclubHandler).Methods("POST")
	router.HandleFunc("/chessclubs/{id}", DeleteChessclubHandler).Methods("DELETE")
	router.HandleFunc("/chessclubs/{id}", EditChessclubHandler).Methods("PUT")
	router.HandleFunc("/players/{id}", GetPlayerDetailsHandler).Methods("GET")
	router.HandleFunc("/players", CreatePlayerHandler).Methods("POST")
}

func addMiddlewares(r *mux.Router) {
	r.Use(loggerMiddleware)
	r.Use(headersMiddleware)
}

func serve() {
	log.Printf("Server listening on %s\n", config.String("HOST"))
	http.Handle("/", router)
	http.ListenAndServe(config.String("HOST"), nil)
}
