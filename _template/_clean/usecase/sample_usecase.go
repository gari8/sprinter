package usecase

import (
	"@@.ImportPath@@/domain/model"
	"@@.ImportPath@@/interfaces/repository"
)

type (
	sampleUseCase struct {
		repository.SampleRepository
	}
	SampleUseCase interface {
		GetSamples() ([]*model.Sample, error)
	}
)

func NewSampleUseCase(rs repository.SampleRepository) SampleUseCase {
	return &sampleUseCase{rs}
}

func (s *sampleUseCase) GetSamples() ([]*model.Sample, error) {
	return s.SampleRepository.Fetch()
}
