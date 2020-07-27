package shared

import (
	"reflect"
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

func TestDeletePlayer(t *testing.T) {
	player, err := s.DeletePlayer(data.MockValidPlayer.Id)
	thenAssertErrorIsNil(t, err)
	thenAssertValueIs(t, *player, data.MockValidPlayer)
}

func thenAssertValueIsNil(t *testing.T, value interface{}) {
	if !reflect.ValueOf(value).IsNil() {
		t.Errorf("want value %v, got value %v", nil, value)
	}
}

func thenAssertValueIs(t *testing.T, value, expectedValue interface{}) {
	if value != expectedValue {
		t.Errorf("want value %v, got value %v", expectedValue, value)
	}
}

func thenAssertErrorIsNil(t *testing.T, err error) {
	if err != nil {
		t.Errorf("want error %v, got %v", nil, err)
	}
}

func thenAssertErrorIs(t *testing.T, err error, expectedErr error) {
	if err.Error() != expectedErr.Error() {
		t.Errorf("want error %v, got %v", expectedErr, err)
	}
}

func TestDeleteUnexistentPlayer(t *testing.T) {
	player, err := s.DeletePlayer(-1)
	thenAssertErrorIs(t, err, model.UnexistingError)
	thenAssertValueIsNil(t, player)
}
