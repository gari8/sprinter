package server

import (
	"net/http"

	sampleController "@@.ImportPath@@/adapters/controllers/sample"
	sampleRepository "@@.ImportPath@@/adapters/repositories/sample"
	sampleApplication "@@.ImportPath@@/application/core/sample"
	"github.com/gari8/sprinter"
	"github.com/go-chi/chi"
)

func SampleRouter(s Server) http.Handler {
	r := chi.NewRouter()

	repo := sampleRepository.NewSampleRepository(s.Conn)
	application := sampleApplication.NewSampleApplication(repo)
	controller := sampleController.NewSampleController(application)
	r.Get("/", sprinter.Handle(controller.GetSamples))
	r.Post("/", sprinter.Handle(controller.CreateSample))
	return r
}
