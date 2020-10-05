package router

import (
	"github.com/HAGARIHAYATO/sprinter/interactor"
	"github.com/HAGARIHAYATO/sprinter/presenter/middleware"
)

package router

import (
"sprinter/interactor"
"sprinter/presenter/handler"
"sprinter/presenter/middleware"
"github.com/go-chi/chi"
"github.com/go-chi/chi/middleware"
)

type Server struct {
	Route *chi.Mux
}

func NewRouter() *Server {
	return &Server {
		Route: chi.NewRouter(),
	}
}

func (s *Server) Router(h *interactor.Handler, m middleware.Middleware) {}