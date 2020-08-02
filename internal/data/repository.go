package data

import (
	"errors"
	"log"

	"github.com/taciomcosta/chesstournament/internal/model"
)

func NewRepository() model.Repository {
	base := baseRepository{}
	return model.Repository{
		Club:   ClubRepository{base},
		Player: PlayerRepository{base},
	}
}

type baseRepository struct{}

func (_ baseRepository) Add(value interface{}) error {
	_, err := db.Model(value).OnConflict("(id) DO UPDATE").Insert()
	if err != nil {
		log.Println(err)
		return model.UnknownError
	}
	return nil
}

func (_ baseRepository) Remove(value interface{}) error {
	return db.Delete(value)
}

func (_ baseRepository) FindOne(value interface{}) error {
	err := db.Select(value)
	if isRecordNotFoundError(err) {
		return errors.New("Player not found")
	}
	if err != nil {
		return model.UnknownError
	}
	return nil
}

func isRecordNotFoundError(err error) bool {
	return err != nil && err.Error() == "pg: no rows in result set"
}
