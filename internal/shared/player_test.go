package shared

import (
	"reflect"
	"testing"

	"github.com/taciomcosta/chesstournament/internal/data"
	"github.com/taciomcosta/chesstournament/internal/model"
)

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

func thenAssertValueIsNotNil(t *testing.T, err error) {
	if err == nil {
		t.Errorf("want error, got %v", nil)
	}
}

func TestDeleteUnexistentPlayer(t *testing.T) {
	player, err := s.DeletePlayer(-1)
	thenAssertErrorIs(t, err, model.UnexistingError)
	thenAssertValueIsNil(t, player)
}

func TestCreatePlayer(t *testing.T) {
	response, err := s.CreatePlayer(&MockCreatePlayerDTO)
	thenAssertValueIs(t, *response, MockCreatePlayerDTOWitId)
	thenAssertErrorIsNil(t, err)
}

func TestCreateInvalidPlayer(t *testing.T) {
	response, err := s.CreatePlayer(&MockCreatePlayerDTOInvalid)
	thenAssertValueIsNil(t, response)
	thenAssertValueIsNotNil(t, err)
}

func TestCreatePlayerInvalidClub(t *testing.T) {
	response, err := s.CreatePlayer(&MockCreatePlayerDTOInvalidClub)
	thenAssertValueIsNil(t, response)
	thenAssertValueIsNotNil(t, err)
}
