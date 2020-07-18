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
}

func NewService(c model.ChessClubRepository, p model.PlayerRepository) Service {
	return service{c, p}
}

type service struct {
	chessclubRepository model.ChessClubRepository
	playerRepository    model.PlayerRepository
}
