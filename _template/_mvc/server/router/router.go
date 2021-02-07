package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"server/controllers"
)

type Server struct {
	Route *chi.Mux
}

func NewRouter() *Server {
	return &Server{
		Route: chi.NewRouter(),
	}
}


func (s *Server) Router(c *controllers.Controller) {
	s.Route.Use(middleware.Logger)
	s.Route.Use(middleware.Recoverer)
	s.Route.Route("/", func(r chi.Router) {
		r.Get("/", c.SampleController.SampleHTML)
	})
	s.Route.Route("/server/v1", func(r chi.Router) {
		r.Get("/",c.SampleController.SampleIndex)
		// TODO
	})
}
