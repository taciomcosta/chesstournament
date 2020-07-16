package data

import (
	"log"

	"github.com/taciomcosta/chesstournament/internal/model"
)

type PlayerRepository struct{}

func (r PlayerRepository) Add(p *model.Player) (*model.Player, error) {
	_, err := db.Model(p).OnConflict("(id) DO UPDATE").Insert()
	if err != nil {
		log.Println(err)
		return nil, model.UnknownError
	}
	return p, nil
}
