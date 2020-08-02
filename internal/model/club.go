package model

type Club struct {
	Id      int    `json:"id" pg:",pk"`
	Name    string `validate:"required" json:"name"`
	Address string `validate:"required" json:"address"`
}
