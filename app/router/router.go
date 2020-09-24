package router

import (
	"fabiangonz98/go-cloudnative-practice/app"
	"fabiangonz98/go-cloudnative-practice/app/requestlog"
	"fabiangonz98/go-cloudnative-practice/app/router/middleware"
	"github.com/go-chi/chi"
)

func New(a *app.App) *chi.Mux {
	l := a.Logger()
	r := chi.NewRouter()

	r.Get("/healthz/liveness", app.HandleLive)
	r.Method("GET", "/", requestlog.NewHandler(a.HandleIndex, l))
	r.Method("GET", "/healthz/readiness", requestlog.NewHandler(a.HandleReady, l))

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.ContentTypeJSON)

		r.Method("GET", "/books", requestlog.NewHandler(a.HandleListBooks, l))
		r.Method("POST", "/books", requestlog.NewHandler(a.HandleCreateBook, l))
		r.Method("GET", "/books/{id}", requestlog.NewHandler(a.HandleReadBook, l))
		r.Method("PUT", "/books/{id}", requestlog.NewHandler(a.HandleUpdateBook, l))
		r.Method("DELETE", "/books/{id}", requestlog.NewHandler(a.HandleDeleteBook, l))
	})

	return r
}
