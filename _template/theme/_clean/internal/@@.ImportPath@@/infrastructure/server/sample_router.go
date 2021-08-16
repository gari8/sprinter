package server

import (
	"fmt"
	"github.com/gari8/sprinter"
	"github.com/go-chi/chi"
	sampleController "@@.ImportPath@@/internal/@@.ImportPath@@/interfaces/controllers/sample"
	samplePresenter "@@.ImportPath@@/internal/@@.ImportPath@@/interfaces/presenters/sample"
	sampleRepository "@@.ImportPath@@/internal/@@.ImportPath@@/interfaces/repositories/sample"
	sampleService "@@.ImportPath@@/internal/@@.ImportPath@@/usecases/services/sample"
	sampleUseCase "@@.ImportPath@@/internal/@@.ImportPath@@/usecases/usecases/sample"
	"net/http"
)

func sampleRouter(s Server) http.Handler {
	r := chi.NewRouter()

	repo := sampleRepository.NewSampleRepository(s.Conn)
	service := sampleService.NewSampleService(repo)
	presenter := samplePresenter.NewGetSamplePresenter()
	interactor := sampleUseCase.NewGetSampleInteractor(presenter, service)
	controller := sampleController.NewSampleController(interactor)
	r.Get("/", sprinter.Handle(controller.GetSamples))
	r.Post("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("=======hi")
	})
	return r
}