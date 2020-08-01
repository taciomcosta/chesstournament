package main

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"github.com/taciomcosta/chesstournament/internal/model"

	"strconv"
)

func TestGetIdFromRequest(t *testing.T) {
	want := 5
	request := newRequestBuilder().withPathVar("id", strconv.Itoa(want)).build()
	got := getIdFromRequest(request)
	if got != want {
		t.Errorf("should get id from request")
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

func TestRequestBuilderBody(t *testing.T) {
	wantBody := `{"id":"id"}`
	request := newRequestBuilder().withBody(wantBody).build()
	if wantBody != readBodyAsString(request) {
		t.Errorf("expected request to have body %s", wantBody)
	}
}

func readBodyAsString(request *http.Request) string {
	bytes, _ := ioutil.ReadAll(request.Body)
	request.Body.Close()
	return string(bytes)
}

func TestRequestBuilderPathVars(t *testing.T) {
	wantVars := map[string]string{"1": "1", "2": "2"}
	request := newRequestBuilder().withPathVar("1", "1").withPathVar("2", "2").build()
	if !reflect.DeepEqual(wantVars, mux.Vars(request)) {
		t.Errorf("expected request to have path vars %v", wantVars)
	}
}

func TestRequestBuilderQueryParams(t *testing.T) {
	wantParams := map[string][]string{"1": {"1"}, "2": {"2"}}
	request := newRequestBuilder().withQueryParam("1", "1").withQueryParam("2", "2").build()
	for key, value := range wantParams {
		if !reflect.DeepEqual(value, request.URL.Query()[key]) {
			t.Errorf("expected request to have query params %v, %v", wantParams, request.URL.Query())
		}
	}
}
