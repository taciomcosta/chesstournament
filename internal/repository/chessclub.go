package repository

import (
	"errors"
	"github.com/taciomcosta/chesstournament/internal/model"
)

type ChessClub interface {
	GetById(int) (*model.ChessClub, error)
}

type ChessClubRepository struct{}

func (r ChessClubRepository) GetById(id int) (*model.ChessClub, error) {
	club := &model.ChessClub{Id: id}
	err := db.Select(club)
	if err != nil {
		return nil, errors.New("club not found")
	}
	return club, nil
}
