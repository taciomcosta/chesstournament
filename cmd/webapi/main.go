package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taciomcosta/chesstournament/internal/config"
	"github.com/taciomcosta/chesstournament/internal/data"
	"github.com/taciomcosta/chesstournament/internal/shared"
)

var service shared.Service = shared.NewService(data.NewRepository())
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
	router.HandleFunc("/chessclubs/{id}", GetClubDetailsHandler).Methods("GET")
	router.HandleFunc("/chessclubs", ListClubsHandler).Methods("GET")
	router.HandleFunc("/chessclubs", CreateClubHandler).Methods("POST")
	router.HandleFunc("/chessclubs/{id}", DeleteClubHandler).Methods("DELETE")
	router.HandleFunc("/chessclubs/{id}", EditClubHandler).Methods("PUT")
	router.HandleFunc("/players/{id}", GetPlayerDetailsHandler).Methods("GET")
	router.HandleFunc("/players", CreatePlayerHandler).Methods("POST")
	router.HandleFunc("/players/{id}", DeletePlayerHandler).Methods("DELETE")
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
