package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/taciomcosta/chesstournament/internal/model"
)

func respond(w http.ResponseWriter, data interface{}) {
	json := mustJSON(data)
	w.Write(json)
}

func mustJSON(v interface{}) []byte {
	json, _ := json.Marshal(v)
	return json
}

func tryRespondWithError(w http.ResponseWriter, httpStatus int, err error) bool {
	if err == nil {
		return false
	}
	respondWithError(w, httpStatus, err)
	return true
}

func respondWithError(w http.ResponseWriter, httpStatus int, err error) {
	if err == model.UnknownError {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(httpStatus)
	}
	w.Write(errorResponse(err))
}

func errorResponse(err error) []byte {
	errString := fmt.Sprintf(`{"msg": "%s"}`, err)
	return []byte(errString)
}
