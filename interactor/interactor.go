package interactor

import (
	"database/sql"
)

type (
	interactor struct {
		conn *sql.DB
	}
	Interactor interface {
		NewRepository() *Repository
		NewApplication(r *Repository) *Application
		NewHandler(a *Application) *Handler
	}
	Repository struct {}
	Application struct {}
	Handler struct {}
)

func NewInteractor(conn *sql.DB) Interactor {
	return &interactor{conn}
}

func (i *interactor) NewRepository() *Repository {
	r := &Repository{}
	return r
}

func (i *interactor) NewApplication(r *Repository) *Application {
	a := &Application{}
	return a
}

func (i *interactor) NewHandler(a *Application) *Handler {
	h := &Handler{}
	return h
}




