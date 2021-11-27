package server

import (
	sampleApplicationService "@@.ImportPath@@/application/service/sample"
	sampleRepository "@@.ImportPath@@/infrastructure/repositories/sample"
	"@@.ImportPath@@/presentation/controllers/sample"

	"github.com/gari8/sprinter"
	"github.com/go-chi/chi"

	"net/http"
)

func SampleRouter(s Server) http.Handler {
	r := chi.NewRouter()

	repository := sampleRepository.NewSampleRepository(s.Conn)
	applicationService := sampleApplicationService.NewSampleApplicationService(repository)
	controller := sample.NewSampleController(applicationService)
	r.Get("/", sprinter.Handle(controller.GetSamples))
	r.Post("/", sprinter.Handle(controller.CreateSample))
	return r
}
