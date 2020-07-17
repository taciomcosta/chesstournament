package model

type Player struct {
	Id        int     `json:"id" pg:",pk"`
	ClubId    int     `validate:"required" json:"clubId"`
	Ranking   Ranking `validate:"required,min=1,max=3" json:"rankingCode"`
	FirstName string  `validate:"required" json:"firstName"`
	LastName  string  `validate:"required" json:"lastName"`
	Address   string  `validate:"required" json:"address"`
	Phone     string  `validate:"required" json:"phone"`
	Email     string  `validate:"required,email" json:"email"`
}

type Ranking int

const (
	RankingMaster  = 1
	RankingLearner = 2
	RankingNewbie  = 3
)

type PlayerRepository interface {
	Add(*Player) (*Player, error)
}
