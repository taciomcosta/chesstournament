package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/taciomcosta/chesstournament/internal/model"
)

func getIdFromRequest(r *http.Request) int {
	vars := mux.Vars(r)
	ID, _ := strconv.Atoi(vars["id"])
	return ID
}

func newFilter(r *http.Request) model.Filter {
	var f model.Filter
	r.ParseForm()
	schema.NewDecoder().Decode(&f, r.Form)
	return f
}

type requestBuilder struct {
	request *http.Request
}

func newRequestBuilder() *requestBuilder {
	builder := new(requestBuilder)
	builder.request, _ = http.NewRequest("GET", "/", strings.NewReader(""))
	builder.request = mux.SetURLVars(builder.request, map[string]string{})
	return builder
}

func (b *requestBuilder) withBody(body string) *requestBuilder {
	reader := strings.NewReader(body)
	b.request.Body = ioutil.NopCloser(reader)
	return b
}

func (b *requestBuilder) withPathVar(key string, value string) *requestBuilder {
	vars := mux.Vars(b.request)
	vars[key] = value
	b.request = mux.SetURLVars(b.request, vars)
	return b
}

func (b *requestBuilder) withQueryParam(key string, value string) *requestBuilder {
	query := b.request.URL.Query()
	query.Set(key, value)
	b.request.URL.RawQuery = query.Encode()
	return b
}

func (b *requestBuilder) build() *http.Request {
	return b.request
}
