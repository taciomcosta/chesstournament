package repository

import (
	"errors"

	"github.com/taciomcosta/chesstournament/internal/model"
)

var InternalDBErr = errors.New("Internal database error")

type ChessClub interface {
	GetById(int) (*model.ChessClub, error)
	Create(*model.ChessClub) (*model.ChessClub, error)
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
