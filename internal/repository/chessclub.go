package repository

import (
	"github.com/taciomcosta/chesstournament/internal/model"
)

type ChessClub interface {
	GetById(int) (*model.ChessClub, error)
}
