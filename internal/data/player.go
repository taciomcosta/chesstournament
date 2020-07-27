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

func (r PlayerRepository) FindOne(criteria *model.Player) (*model.Player, error) {
	if err := db.Select(criteria); err != nil {
		return nil, err
	}
	return criteria, nil
}

func (r PlayerRepository) Remove(p *model.Player) error {
	return db.Delete(p)
}
