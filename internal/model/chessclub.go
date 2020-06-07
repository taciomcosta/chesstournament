package model

type ChessClub struct {
	Id      int    `json:"id" pg:",pk"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
