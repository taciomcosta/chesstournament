package chessclub

type ChessClub struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func GetClubById(id int) ChessClub {
	return mockChessClub
}
