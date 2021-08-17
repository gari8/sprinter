package server

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

type Server struct {
	Route *chi.Mux
	Conn *sql.DB
}

func NewServer(conn *sql.DB) Server {
	return Server{
		Route: chi.NewRouter(),
		Conn: conn,
	}
}

func (s Server) Serve() {
	s.Route.Use(middleware.Logger)
	s.Route.Use(middleware.Recoverer)
	s.Route.Route("/api/v1", func(r chi.Router) {
		r.Mount("/sample", SampleRouter(s))
	})
	err := http.ListenAndServe(":8080", s.Route)
	if err != nil {
		panic(err)
	}
}
