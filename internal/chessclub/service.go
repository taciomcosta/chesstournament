package chessclub

import (
	"github.com/taciomcosta/chesstournament/internal/model"
	"github.com/taciomcosta/chesstournament/internal/repository"
)

var chessclubRepository repository.ChessClub

func init() {
	chessclubRepository = repository.ChessClubRepository{}
}

type Service interface {
	GetClubById(id int) (*model.ChessClub, error)
}

func New() Service {
	return service{}
}

type service struct{}

func (s service) GetClubById(id int) (*model.ChessClub, error) {
	return chessclubRepository.GetById(id)
}
