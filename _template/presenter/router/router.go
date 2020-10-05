package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"@@.ImportPath@@/interactor"
	mid "@@.Important@@/presenter/middleware"
)

type Server struct {
	Route *chi.Mux
}

func NewRouter() *Server {
	return &Server{
		Route: chi.NewRouter(),
	}
}

func (s *Server) Router(h *interactor.Handler, m mid.Middleware) {
	s.Route.Use(middleware.Logger)
	s.Route.Use(middleware.Recoverer)
	s.Route.Route("/api/v1", func(r chi.Router) {
		r.Get("/", h.SampleHandler.SampleIndex)
		// TODO
	})
}