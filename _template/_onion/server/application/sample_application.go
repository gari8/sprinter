package application

import (
	"server/domain/model"
	"server/domain/repository"
)

type (
	sampleApplication struct {
		repository.SampleRepository
	}
	SampleApplication interface {
		GetSamples() ([]*model.Sample, error)
	}
)

func NewSampleApplication(rs repository.SampleRepository) SampleApplication {
	return &sampleApplication{rs}
}

func (s *sampleApplication) GetSamples() ([]*model.Sample, error) {
	return s.SampleRepository.Fetch()
}
