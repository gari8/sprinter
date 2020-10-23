package handler

import "@@.ImportPath@@/usecase"

type (
	Handler struct {
		SampleHandler
	}
)

func NewHandler(a *usecase.UseCase) *Handler {
	h := &Handler{}
	h.SampleHandler = NewSampleHandler(a.SampleUseCase)
	return h
}

