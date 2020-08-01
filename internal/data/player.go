package data

import (
	"errors"
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
	err := db.Select(criteria)
	if isNotFoundError(err) {
		return nil, errors.New("Player not found")
	}
	if err != nil {
		return nil, model.UnknownError
	}
	return criteria, nil
}

func isNotFoundError(err error) bool {
	return err != nil && err.Error() == "pg: no rows in result set"
}

func (r PlayerRepository) Remove(p *model.Player) error {
	return db.Delete(p)
}

func (r PlayerRepository) Count(criteria *model.Player) int {
	count, err := db.Model(criteria).Count()
	if err != nil {
		return 0
	}
	return count
}
