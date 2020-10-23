package usecase

import (
	"@@.ImportPath@@/domain/model"
	"@@.ImportPath@@/interfaces/controllers"
)

type (
	sampleUseCase struct {
		controllers.SampleController
	}
	SampleUseCase interface {
		GetSamples() ([]*model.Sample, error)
	}
)

func NewSampleUseCase(rs controllers.SampleController) SampleUseCase {
	return &sampleUseCase{rs}
}

func (s *sampleUseCase) GetSamples() ([]*model.Sample, error) {
	return s.Fetch()
}
