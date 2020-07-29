package shared

import (
	"github.com/taciomcosta/chesstournament/internal/model"
)

type Service interface {
	GetClubById(id int) (*model.Club, error)
	CreateClub(c *model.Club) (*model.Club, error)
	EditClub(id int, c *model.Club) error
	ListClubs(r model.Filter) ([]model.Club, error)
	DeleteClub(id int) (*model.Club, error)
	CreatePlayer(request *CreatePlayerDTO) (*CreatePlayerDTO, error)
	GetPlayerById(id int) (*model.Player, error)
	DeletePlayer(id int) (*model.Player, error)
}

func NewService(club model.ClubRepository, player model.PlayerRepository) Service {
	return service{club, player}
}

type service struct {
	chessclubRepository model.ClubRepository
	playerRepository    model.PlayerRepository
}
