package interactor

import (
	"database/sql"
	//"@@.ImportPath@@/application"
	//"@@.ImportPath@@/domain/repository"
	//"@@.ImportPath@@/presenter/handler"
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
	Repository struct {
		repository.SampleRepository
	}
	Application struct {
		application.SampleApplication
	}
	Handler struct {
		handler.SampleHandler
	}
)

func NewInteractor(conn *sql.DB) Interactor {
	return &interactor{conn}
}

func (i *interactor) NewRepository() *Repository {
	r := &Repository{}
	r.SampleRepository = repository.NewSampleRepository(i.conn)
	return r
}

func (i *interactor) NewApplication(r *Repository) *Application {
	a := &Application{}
	a.SampleApplication = application.NewSampleApplication(r.SampleRepository)
	return a
}

func (i *interactor) NewHandler(a *Application) *Handler {
	h := &Handler{}
	h.SampleHandler = handler.NewSampleHandler(a.SampleApplication)
	return h
}





