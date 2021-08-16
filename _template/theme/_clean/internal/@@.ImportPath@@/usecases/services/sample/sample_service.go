package sample

import (
	"@@.ImportPath@@/internal/@@.ImportPath@@/domain"
)

type Repository interface {
	GetSamples() ([]domain.Sample, error)
}

type SampleService struct {
	repository Repository
}

func NewSampleService(repository Repository) SampleService {
	return SampleService{repository: repository}
}

func (s SampleService) GetSamples() ([]domain.Sample, error) {
	return s.repository.GetSamples()
}
