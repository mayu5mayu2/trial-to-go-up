package main

import (
	"/Users/masahironakamura/Desktop/執筆関係/Axross/axross_go_app/webapp/internal/handlers/handlers.go"
	"net/http"

	"github.com/bmizerany/pat"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	return mux
}
