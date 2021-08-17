package sample

import (
	"@@.ImportPath@@/internal/@@.ImportPath@@/domain/model"
	sampleRepository "@@.ImportPath@@/internal/@@.ImportPath@@/domain/service/sample"
)

type CreateSampleInput struct {
	Text string `json:"text"`
}

type SampleApplicationService struct {
	repo sampleRepository.Repository
}

func NewSampleApplicationService(
	repo sampleRepository.Repository) SampleApplicationService {
	return SampleApplicationService{
		repo: repo,
	}
}

func (s SampleApplicationService) GetSamples() ([]model.Sample, error) {
	return s.repo.GetSamples()
}

func (s SampleApplicationService) CreateSample(input CreateSampleInput) error {
	var sample model.Sample
	sample.Text = input.Text
	return s.repo.CreateSample(sample)
}
