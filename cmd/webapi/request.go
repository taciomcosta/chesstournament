package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/taciomcosta/chesstournament/internal/repository"
	"net/http"
	"strconv"
)

func id(r *http.Request) int {
	vars := mux.Vars(r)
	ID, _ := strconv.Atoi(vars["id"])
	return ID
}

// TODO: find a better name or abstraction
// it's akward to have http concerns knowing about repository existence
func newFilter(r *http.Request) repository.Filter {
	var f repository.Filter
	r.ParseForm()
	schema.NewDecoder().Decode(&f, r.Form)
	return f
}
