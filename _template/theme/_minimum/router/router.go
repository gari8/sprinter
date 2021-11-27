package router

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

type Server struct {
	Route *chi.Mux
	Conn *sql.DB
}

func NewRouter(conn *sql.DB) *Server {
	return &Server{
		Route: chi.NewRouter(),
		Conn: conn,
	}
}

type Sample struct {
	ID   int64 `json:"id" db:"id"`
	Text string `json:"text" db:"text"`
}

func (s *Server) Router() {
	s.Route.Use(middleware.Logger)
	s.Route.Use(middleware.Recoverer)
	s.Route.Route("/api/v1", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			var samples []*Sample
			rows, err := s.Conn.Query("SELECT id, text FROM samples;")
			if rows == nil {
				log.Fatalln(err)
			}
			for rows.Next() {
				sample := &Sample{}
				err = rows.Scan(&sample.ID, &sample.Text)
				if err == nil {
					samples = append(samples, sample)
				}
			}

			res, err := json.Marshal(samples)
			if err != nil {
				log.Fatal(err)
			}

			_, _ = w.Write(res)
		})
	})
}
