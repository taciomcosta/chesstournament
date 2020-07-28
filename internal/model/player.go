package model

type Player struct {
	Id        int
	ClubId    int
	Ranking   Ranking `validate:"required,min=1,max=3"`
	FirstName string  `validate:"required"`
	LastName  string  `validate:"required"`
	Address   string  `validate:"required"`
	Phone     string  `validate:"required"`
	Email     string  `validate:"required,email"`
}

type Ranking int

const (
	RankingMaster  = 1
	RankingLearner = 2
	RankingNewbie  = 3
)

type PlayerRepository interface {
	Add(*Player) (*Player, error)
	FindOne(*Player) (*Player, error)
	Remove(*Player) error
}
