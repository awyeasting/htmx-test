package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
)

func NewRootRouter() *chi.Mux {
	log.Info("Creating Root Router")
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		component := index("Alec yeastttttttt")
		component.Render(r.Context(), w)
	})

	staticFS := http.FileServer(http.Dir("./public"))
	r.Handle("/*", staticFS)

	return r
}
