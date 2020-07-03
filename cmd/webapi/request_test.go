package main

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/taciomcosta/chesstournament/internal/model"

	"strconv"
)

func TestId(t *testing.T) {
	want := 5
	w, _ := http.NewRequest("GET", "/", nil)
	w = mux.SetURLVars(w, map[string]string{"id": strconv.Itoa(want)})
	got := id(w)
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestNewFilter(t *testing.T) {
	tests := []struct {
		queryParams string
		want        model.Filter
	}{
		{
			queryParams: "/?$top=1&$offset=2&$orderBy=p1 asc, p2 desc",
			want:        model.Filter{Limit: 1, Offset: 2, OrderBy: "p1 asc, p2 desc"},
		},
		{
			queryParams: "/?$orderBy=p1 asc, p2 desc",
			want:        model.Filter{OrderBy: "p1 asc, p2 desc"},
		},
	}

	for _, tt := range tests {
		w, _ := http.NewRequest("GET", tt.queryParams, nil)
		got := newFilter(w)
		if got != tt.want {
			t.Errorf("want %v, got %v", tt.want, got)
		}
	}
}
