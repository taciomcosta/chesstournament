package shared

var MockCreatePlayerDTO CreatePlayerDTO = CreatePlayerDTO{
	ClubId:    1,
	Ranking:   1,
	FirstName: "Tacio",
	LastName:  "Costa",
	Address:   "Somewhere",
	Phone:     "123456789",
	Email:     "tacio@email.com",
}

var MockCreatePlayerDTOWitId = CreatePlayerDTO{
	Id:        1,
	ClubId:    1,
	Ranking:   1,
	FirstName: "Tacio",
	LastName:  "Costa",
	Address:   "Somewhere",
	Phone:     "123456789",
	Email:     "tacio@email.com",
}

var MockCreatePlayerDTOInvalidClub CreatePlayerDTO = CreatePlayerDTO{
	ClubId:    0,
	Ranking:   1,
	FirstName: "Tacio",
	LastName:  "Costa",
	Address:   "Somewhere",
	Phone:     "123456789",
	Email:     "tacio@email.com",
}

var MockCreatePlayerDTOInvalid CreatePlayerDTO = CreatePlayerDTO{
	ClubId:    1,
	Ranking:   0,
	FirstName: "",
	LastName:  "",
	Address:   "",
	Phone:     "",
	Email:     "",
}
