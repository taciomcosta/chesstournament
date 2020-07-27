package shared

import (
	"github.com/taciomcosta/chesstournament/internal/model"
)

type Service interface {
	GetClubById(id int) (*model.ChessClub, error)
	CreateChessclub(c *model.ChessClub) (*model.ChessClub, error)
	EditChessclub(id int, c *model.ChessClub) error
	ListClubs(r model.Filter) ([]model.ChessClub, error)
	DeleteClub(id int) (*model.ChessClub, error)
	CreatePlayer(p *model.Player) (*model.Player, error)
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
