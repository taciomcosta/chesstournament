package main

import (
	"log"
	"net/http"
	"strings"
)

type Middleware func(http.Handler) http.Handler

func headersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, swaggerURLPath) {
			w.Header().Set("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.Method + r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
