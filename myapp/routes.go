package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (a *application) routes() *chi.Mux {
	// middleware

	// add routes
	a.App.Routes.Get("/", a.Handers.Home)

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
