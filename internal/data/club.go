package data

import (
	"errors"

	"github.com/taciomcosta/chesstournament/internal/model"
)

type ClubRepository struct{}

func (r ClubRepository) GetById(id int) (*model.Club, error) {
	club := &model.Club{Id: id}
	err := db.Select(club)
	if isRecordNotFoundError(err) {
		return nil, errors.New("Club not found")
	}
	if err != nil {
		return nil, model.UnknownError
	}
	return club, nil
}

func (r ClubRepository) Add(c *model.Club) (*model.Club, error) {
	_, err := db.Model(c).OnConflict("(id) DO UPDATE").Insert()
	if err != nil {
		return nil, model.UnknownError
	}
	return c, nil
}

func (r ClubRepository) ListClubs(lr model.Filter) ([]model.Club, error) {
	cs := make([]model.Club, 0)
	if err := lr.Validate(); err != nil {
		return cs, err
	}
	err := db.Model(&cs).OrderExpr(lr.OrderBy).Offset(lr.Offset).Limit(lr.Limit).Select()
	return cs, err
}

func (r ClubRepository) Remove(c *model.Club) error {
	return db.Delete(c)
}
