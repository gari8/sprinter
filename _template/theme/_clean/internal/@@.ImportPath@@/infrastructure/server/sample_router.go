package server

import (
	"github.com/gari8/sprinter"
	"github.com/go-chi/chi"
	sampleController "@@.ImportPath@@/internal/@@.ImportPath@@/interfaces/controllers/sample"
	samplePresenter "@@.ImportPath@@/internal/@@.ImportPath@@/interfaces/presenters/sample"
	sampleRepository "@@.ImportPath@@/internal/@@.ImportPath@@/interfaces/repositories/sample"
	sampleService "@@.ImportPath@@/internal/@@.ImportPath@@/usecases/services/sample"
	sampleUseCase "@@.ImportPath@@/internal/@@.ImportPath@@/usecases/usecases/sample"
	"net/http"
)

func SampleRouter(s Server) http.Handler {
	r := chi.NewRouter()

	repo := sampleRepository.NewSampleRepository(s.Conn)
	service := sampleService.NewSampleService(repo)
	getPresenter := samplePresenter.NewGetSamplePresenter()
	createPresenter := samplePresenter.NewCreateSamplePresenter()
	getInteractor := sampleUseCase.NewGetSampleInteractor(getPresenter, service)
	createInteractor := sampleUseCase.NewCreateSampleInteractor(createPresenter, service)
	controller := sampleController.NewSampleController(getInteractor, createInteractor)

	r.Get("/", sprinter.Handle(controller.GetSamples))
	r.Post("/", sprinter.Handle(controller.CreateSample))
	return r
}