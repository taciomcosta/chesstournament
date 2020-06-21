package repository

import (
	"errors"

	"github.com/taciomcosta/chesstournament/internal/model"
)

var InternalDBErr = errors.New("Internal database error")

type ChessClub interface {
	GetById(int) (*model.ChessClub, error)
	Create(*model.ChessClub) (*model.ChessClub, error)
	ListClubs(Filter) ([]model.ChessClub, error)
}

type ChessClubRepository struct{}

func (r ChessClubRepository) GetById(id int) (*model.ChessClub, error) {
	club := &model.ChessClub{Id: id}
	if err := db.Select(club); err != nil {
		return nil, err
	}
	return club, nil
}

func (r ChessClubRepository) Create(c *model.ChessClub) (*model.ChessClub, error) {
	if err := db.Insert(c); err != nil {
		return nil, InternalDBErr
	}
	return c, nil
}

func (r ChessClubRepository) ListClubs(lr Filter) ([]model.ChessClub, error) {
	var cs []model.ChessClub
	if err := lr.validate(); err != nil {
		return []model.ChessClub{}, err
	}
	err := db.Model(&cs).OrderExpr(lr.OrderBy).Offset(lr.Offset).Limit(lr.Limit).Select()
	if cs == nil {
		return []model.ChessClub{}, err
	}
	return cs, err
}
