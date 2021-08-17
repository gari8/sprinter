package sample

import (
	"@@.ImportPath@@/internal/@@.ImportPath@@/domain"
)

type Repository interface {
	GetSamples() ([]domain.Sample, error)
	CreateSample(sample domain.Sample) error
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

func (s SampleService) CreateSample(sample domain.Sample) error {
	err := s.repository.CreateSample(sample)
	if err != nil {
		return err
	}
	return nil
}
