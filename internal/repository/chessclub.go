package repository

import (
	"github.com/taciomcosta/chesstournament/internal/model"
)

type ChessClub interface {
	GetById(int) (*model.ChessClub, error)
}

type ChessClubRepository struct{}

func (r ChessClubRepository) GetById(id int) (*model.ChessClub, error) {
	club := &model.ChessClub{Id: id}
	if err := db.Select(club); err != nil {
		return nil, err
	}
	return club, nil
}
