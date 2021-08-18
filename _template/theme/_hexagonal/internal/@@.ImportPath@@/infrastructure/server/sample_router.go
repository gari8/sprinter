package server

import (
	"github.com/gari8/sprinter"
	"github.com/go-chi/chi"
	sampleController "@@.ImportPath@@/internal/@@.ImportPath@@/adapters/controllers/sample"
	sampleRepository "@@.ImportPath@@/internal/@@.ImportPath@@/adapters/repositories/sample"
	sampleApplication "@@.ImportPath@@/internal/@@.ImportPath@@/application/core/sample"
	"net/http"
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
