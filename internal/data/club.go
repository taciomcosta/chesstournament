package data

import (
	"github.com/taciomcosta/chesstournament/internal/model"
)

type ClubRepository struct {
	base baseRepository
}

func (r ClubRepository) GetById(id int) (*model.Club, error) {
	club := &model.Club{Id: id}
	return club, r.base.FindOne(club)
}

func (r ClubRepository) Add(club *model.Club) (*model.Club, error) {
	return club, r.base.Add(club)
}

func (r ClubRepository) ListClubs(filter model.Filter) ([]model.Club, error) {
	clubs := make([]model.Club, 0)
	if err := filter.Validate(); err != nil {
		return clubs, err
	}
	err := db.Model(&clubs).
		OrderExpr(filter.OrderBy).
		Offset(filter.Offset).
		Limit(filter.Limit).Select()
	return clubs, err
}

func (r ClubRepository) Remove(club *model.Club) error {
	return r.base.Remove(club)
}
