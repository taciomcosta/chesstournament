package shared

import (
	"testing"

	"github.com/taciomcosta/chesstournament/internal/data"
	"github.com/taciomcosta/chesstournament/internal/model"
)

func TestCreatePlayer(t *testing.T) {
	tests := []struct {
		p             *model.Player
		expectsErr    bool
		expectsPlayer bool
		description   string
	}{
		{
			p:             &data.MockValidPlayer,
			expectsErr:    false,
			expectsPlayer: true,
			description:   "should create player",
		},
		{
			p:             &data.MockInvalidPlayer,
			expectsErr:    true,
			expectsPlayer: false,
			description:   "should not create invalid player",
		},
		{
			p:             &data.MockNoClubPlayer,
			expectsErr:    true,
			expectsPlayer: false,
			description:   "should not create player for non-existing club",
		},
	}

	for _, tt := range tests {
		p, err := s.CreatePlayer(tt.p)
		if tt.expectsPlayer && p == nil {
			t.Error("expects player")
		}
		if tt.expectsErr && err == nil {
			t.Error(tt.description)
		}
	}
}

/*
[x] should return created player
[x] should not create an invalid player
[ ] should not create a player for a non-existing club
*/
