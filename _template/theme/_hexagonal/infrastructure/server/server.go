package server

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"os"
)

const defaultPort = "8080"

type Server struct {
	Route *chi.Mux
	Conn  *sql.DB
}

func NewServer(conn *sql.DB) Server {
	return Server{
		Route: chi.NewRouter(),
		Conn:  conn,
	}
}

func (s Server) Serve() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	s.Route.Use(middleware.Logger)
	s.Route.Use(middleware.Recoverer)
	s.Route.Route("/api/v1", func(r chi.Router) {
		r.Mount("/sample", SampleRouter(s))
	})
	err := http.ListenAndServe(":"+port, s.Route)
	if err != nil {
		panic(err)
	}
}
