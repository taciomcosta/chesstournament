package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/taciomcosta/chesstournament/internal/model"
	"net/http"
	"strconv"
)

func id(r *http.Request) int {
	vars := mux.Vars(r)
	ID, _ := strconv.Atoi(vars["id"])
	return ID
}

func newFilter(r *http.Request) model.Filter {
	var f model.Filter
	r.ParseForm()
	schema.NewDecoder().Decode(&f, r.Form)
	return f
}
