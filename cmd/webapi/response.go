package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func respond(w http.ResponseWriter, data interface{}) {
	json := mustJSON(data)
	w.Write(json)
}

func mustJSON(v interface{}) []byte {
	json, _ := json.Marshal(v)
	return json
}

func errorResponse(err error) []byte {
	errString := fmt.Sprintf(`{"msg": "%s"}`, err)
	return []byte(errString)
}

func tryRespondWithError(w http.ResponseWriter, httpStatus int, err error) bool {
	if err == nil {
		return false
	}
	w.WriteHeader(httpStatus)
	w.Write(errorResponse(err))
	return true
}