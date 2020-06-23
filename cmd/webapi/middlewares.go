package main

import (
	"fmt"
	"net/http"
	"strings"
)

func headersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, swaggerURLPath) {
			fmt.Println(r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}
