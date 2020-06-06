package webapi

import (
	"encoding/json"
	"net/http"

	"github.com/taciomcosta/chesstournament/internal/chessclub"
)

var service chessclub.Service

func init() {
	service = chessclub.New()
}

func GetChessclubDetailsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	club, _ := service.GetClubById(1)
	json := mustJSON(*club)
	w.Write(json)
}

func mustJSON(v interface{}) []byte {
	json, _ := json.Marshal(v)
	return json
}
