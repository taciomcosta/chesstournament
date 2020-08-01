package shared

import (
	"errors"

	"github.com/taciomcosta/chesstournament/internal/model"
)

func (s service) GetPlayerById(id int) (*model.Player, error) {
	p, err := s.playerRepository.FindOne(&model.Player{Id: id})
	if err != nil {
		return nil, err
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
	if !s.clubExists(playerDTO.ClubId) {
		return nil, errors.New("Player does not exist")
	}

	player, err := newPlayer(playerDTO)
	if err != nil {
		return nil, err
	}

	s.playerRepository.Add(player)

	playerDTO.Id = player.Id
	return playerDTO, nil
}

func (s service) clubExists(clubId int) bool {
	if _, err := s.chessclubRepository.GetById(clubId); err != nil {
		return false
	}
	return true
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
