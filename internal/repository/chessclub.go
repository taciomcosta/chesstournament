package repository

import (
	"github.com/taciomcosta/chesstournament/internal/model"
)

type ChessClub interface {
	GetById(int) (*model.ChessClub, error)
	Add(*model.ChessClub) (*model.ChessClub, error)
	ListClubs(Filter) ([]model.ChessClub, error)
	Remove(*model.ChessClub) error
}

type ChessClubRepository struct{}

func (r ChessClubRepository) GetById(id int) (*model.ChessClub, error) {
	club := &model.ChessClub{Id: id}
	if err := db.Select(club); err != nil {
		return nil, err
	}
	return club, nil
}

func (r ChessClubRepository) Add(c *model.ChessClub) (*model.ChessClub, error) {
	_, err := db.Model(c).
		OnConflict("(id) DO UPDATE").
		Insert()
	if err != nil {
		return nil, InternalErr{}
	}
	return c, nil
}

func (r ChessClubRepository) ListClubs(lr Filter) ([]model.ChessClub, error) {
	cs := make([]model.ChessClub, 0)
	if err := lr.validate(); err != nil {
		return cs, err
	}
	err := db.Model(&cs).OrderExpr(lr.OrderBy).Offset(lr.Offset).Limit(lr.Limit).Select()
	return cs, err
}

func (r ChessClubRepository) Remove(c *model.ChessClub) error {
	return db.Delete(c)
}

type InternalErr struct{}

func (e InternalErr) Error() string {
	return "An internal error has occurred"
}
