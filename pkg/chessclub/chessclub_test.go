package chessclub

import (
	"testing"
)

func TestGetClubById(t *testing.T) {
	club := GetClubById(1)
	if club != mockChessClub {
		t.Error("it should get club by id")
	}
}
