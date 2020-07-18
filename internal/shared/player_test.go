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

func TestGetPlayerById(t *testing.T) {
	f := func(id int) (interface{}, error) { return s.GetPlayerById(id) }
	testGetById(f, t)
}

type getFunc func(int) (interface{}, error)

func testGetById(get getFunc, t *testing.T) {
	var tests = []struct {
		id            int
		expectsPlayer bool
		expectsErr    bool
	}{
		{1, true, false},
		{-1, false, true},
	}

	for _, tt := range tests {
		c, err := get(tt.id)
		if tt.expectsPlayer && c == nil {
			t.Error("it should return a Player")
		}
		if tt.expectsErr && err == nil {
			t.Error("it should return an error")
		}
	}

}
