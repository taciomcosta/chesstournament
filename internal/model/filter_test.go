package model

import "testing"

func TestFilter(t *testing.T) {
	t.Run("Empty Filter", testEmptyFilter)
	t.Run("Test.Limit", testLimit)
	t.Run("Test.Offset", testOffset)
}

func testEmptyFilter(t *testing.T) {
	var f Filter

	err := f.Validate()

	if err != nil {
		t.Error("empty filter should be valid")
	}
}

func testLimit(t *testing.T) {
	f := Filter{Limit: minLimit - 1}

	err := f.Validate()
	if err.Error() != "Limit must be between 0 and 20" {
		t.Error("should return error regarding Limit")
	}

	f.Limit = maxLimit + 1
	err = f.Validate()

	if err.Error() != "Limit must be between 0 and 20" {
		t.Error("should return error regarding Limit")
	}
}

func testOffset(t *testing.T) {
	f := Filter{Offset: minOffset - 1}
	err := f.Validate()
	if err.Error() != "Offset must be greater or equal to zero" {
		t.Error("should return error regarding Limit")
	}
}
