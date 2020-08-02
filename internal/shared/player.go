package shared

import (
	"errors"

	"github.com/taciomcosta/chesstournament/internal/model"
)

func (s service) GetPlayerById(id int) (*model.Player, error) {
	p, err := s.repository.Player.FindOne(&model.Player{Id: id})
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s service) DeletePlayer(id int) (*model.Player, error) {
	p, err := s.GetPlayerById(id)
	s.repository.Player.Remove(p)
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
	if !s.clubExists(playerDTO.ClubId) {
		return nil, errors.New("Player does not exist")
	}

	player, err := newPlayer(playerDTO)
	if err != nil {
		return nil, err
	}

	s.repository.Player.Add(player)

	playerDTO.Id = player.Id
	return playerDTO, nil
}

func (s service) clubExists(clubId int) bool {
	_, err := s.repository.Club.GetById(clubId)
	return err == nil
}

func newPlayer(dto *CreatePlayerDTO) (*model.Player, error) {
	player := new(model.Player)
	player.Ranking = model.Ranking(dto.Ranking)
	player.FirstName = dto.FirstName
	player.LastName = dto.LastName
	player.Address = dto.Address
	player.Phone = dto.Phone
	player.Email = dto.Email
	player.ClubId = dto.ClubId
	return player, model.Validate(player)
}
