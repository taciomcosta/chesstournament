package model

type Repository struct {
	Club   ClubRepository
	Player PlayerRepository
}

type ClubRepository interface {
	GetById(int) (*Club, error)
	Add(*Club) (*Club, error)
	ListClubs(Filter) ([]Club, error)
	Remove(*Club) error
}

type PlayerRepository interface {
	Add(*Player) (*Player, error)
	FindOne(*Player) (*Player, error)
	Remove(*Player) error
	Count(*Player) int
}
