package injector

import (
	"@@.ImportPath@@/domain/repository"
	"@@.ImportPath@@/interfaces/controllers"
	"@@.ImportPath@@/usecase"
	"database/sql"
)

type (
	injector struct {
		conn *sql.DB
	}
	Injector interface {
		NewRepository() *Repository
		NewUseCase(r *Repository) *UseCase
		NewController(a *UseCase) *Controller
	}
	Repository struct {
		repository.SampleRepository
	}
	UseCase struct {
		usecase.SampleUseCase
	}
	Controller struct {
		controllers.SampleController
	}
)

func NewInjector(conn *sql.DB) Injector {
	return &injector{conn}
}

func (i *injector) NewRepository() *Repository {
	r := &Repository{}
	r.SampleRepository = repository.NewSampleRepository(i.conn)
	return r
}

func (i *injector) NewUseCase(r *Repository) *UseCase {
	a := &UseCase{}
	a.SampleUseCase = usecase.NewSampleUseCase(r.SampleRepository)
	return a
}

func (i *injector) NewController(a *UseCase) *Controller {
	h := &Controller{}
	h.SampleController = controllers.NewSampleController(a.SampleUseCase)
	return h
}
