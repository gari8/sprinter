package router

import (
	mid "@@.ImportPath@@/infrastructure/middleware"
	"@@.ImportPath@@/interfaces/presenter/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	Route *chi.Mux
}

func NewRouter() *Server {
	return &Server{
		Route: chi.NewRouter(),
	}
}

func (s *Server) Router(h *handler.Handler, m mid.Middleware) {
	s.Route.Use(middleware.Logger)
	s.Route.Use(middleware.Recoverer)
	s.Route.Route("/", func(r chi.Router) {
		r.Get("/", h.SampleHTML)
	})
	s.Route.Route("/api/v1", func(r chi.Router) {
		r.Get("/", h.SampleIndex)
		// TODO
	})
}
