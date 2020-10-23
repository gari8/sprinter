package usecase

import "@@.ImportPath@@/interfaces/controllers"

type (
	UseCase struct {
		SampleUseCase
	}
)

func NewUseCase(r *controllers.Controller) *UseCase {
	a := &UseCase{}
	a.SampleUseCase = NewSampleUseCase(r.SampleController)
	return a
}