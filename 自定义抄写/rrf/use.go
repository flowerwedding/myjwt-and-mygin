package rrf

import (
	"log"
	"net/http"
	"time"
)

func Logging() Middleware {
	return func(f http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()
			f.ServeHTTP(w, r)
		})
	}
}

type Middleware func(http.Handler) http.Handler

func Chain(f http.Handler, mmid ...[]Middleware) http.Handler {
	for _, mmap := range mmid {
		for _, m := range mmap {
			f = m(f)
		}
	}
	return f
}

var Group []Middleware

func (a *App) Use(middlewares ...Middleware) {
	Group = append(Group, middlewares...)
}