package webapi

import (
	"encoding/json"
	"net/http"

	"github.com/taciomcosta/chesstournament/internal/chessclub"
)

func GetChessclubDetailsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json := mustJSON(chessclub.GetClubById(0))
	w.Write(json)
}

func mustJSON(v interface{}) []byte {
	json, _ := json.Marshal(v)
	return json
}
