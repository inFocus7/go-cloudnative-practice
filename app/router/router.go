package router

import (
	"fabiangonz98/go-cloudnative-practice/app"
	"github.com/go-chi/chi"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", app.HandleIndex)

	return r
}
