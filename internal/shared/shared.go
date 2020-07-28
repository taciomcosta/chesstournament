package shared

import (
	"github.com/taciomcosta/chesstournament/internal/model"
)

type Service interface {
	GetClubById(id int) (*model.Club, error)
	CreateChessclub(c *model.Club) (*model.Club, error)
	EditChessclub(id int, c *model.Club) error
	ListClubs(r model.Filter) ([]model.Club, error)
	DeleteClub(id int) (*model.Club, error)
	CreatePlayer(request *CreatePlayerDTO) (*CreatePlayerDTO, error)
	GetPlayerById(id int) (*model.Player, error)
	DeletePlayer(id int) (*model.Player, error)
}

func NewService(
	repository model.Repository, chessclub model.ChessClubRepository, player model.PlayerRepository,
) Service {
	return service{repository, chessclub, player}
}

type service struct {
	repository          model.Repository
	chessclubRepository model.ChessClubRepository
	playerRepository    model.PlayerRepository
}
