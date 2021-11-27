package server

import (
	"net/http"

	sampleController "@@.ImportPath@@/interfaces/controllers/sample"
	samplePresenter "@@.ImportPath@@/interfaces/presenters/sample"
	sampleRepository "@@.ImportPath@@/interfaces/repositories/sample"
	sampleService "@@.ImportPath@@/usecases/services/sample"
	sampleUseCase "@@.ImportPath@@/usecases/usecases/sample"
	"github.com/gari8/sprinter"
	"github.com/go-chi/chi"
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
