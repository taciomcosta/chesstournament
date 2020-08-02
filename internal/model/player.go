package model

type Player struct {
	Id        int     `json:"id"`
	ClubId    int     `json:"clubId"`
	Ranking   Ranking `validate:"required,min=1,max=3" json:"rankingCode"`
	FirstName string  `validate:"required" json:"firstName"`
	LastName  string  `validate:"required" json:"lastName"`
	Address   string  `validate:"required" json:"address"`
	Phone     string  `validate:"required" json:"phone"`
	Email     string  `validate:"required,email" json:"email"`
}

type Ranking int

const (
	RankingMaster  = 0
	RankingLearner = 1
	RankingNewbie  = 2
)
