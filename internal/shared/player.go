package shared

import (
	"github.com/taciomcosta/chesstournament/internal/model"
)

func (s service) GetPlayerById(id int) (*model.Player, error) {
	p, err := s.playerRepository.FindOne(&model.Player{Id: id})
	if err != nil {
		return nil, model.UnexistingError
	}
	return p, nil
}

func (s service) DeletePlayer(id int) (*model.Player, error) {
	p, err := s.GetPlayerById(id)
	s.playerRepository.Remove(p)
	return p, err
}

type CreatePlayerDTO struct {
	Id        int    `json:"id"`
	ClubId    int    `json:"clubId"`
	Ranking   int    `json:"rankingCode"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

func (s service) CreatePlayer(playerDTO *CreatePlayerDTO) (*CreatePlayerDTO, error) {
	club, err := s.chessclubRepository.GetById(playerDTO.ClubId)
	if err != nil {
		return nil, model.UnexistingError
	}

	player, err := newPlayer(playerDTO, club)
	if err != nil {
		return nil, err
	}

	s.playerRepository.Add(player)

	playerDTO.Id = player.Id
	return playerDTO, nil
}

func newPlayer(dto *CreatePlayerDTO, club *model.Club) (*model.Player, error) {
	player := new(model.Player)
	player.Ranking = model.Ranking(dto.Ranking)
	player.FirstName = dto.FirstName
	player.LastName = dto.LastName
	player.Address = dto.Address
	player.Phone = dto.Phone
	player.Email = dto.Email
	player.ClubId = club.Id
	return player, model.Validate(player)
}
