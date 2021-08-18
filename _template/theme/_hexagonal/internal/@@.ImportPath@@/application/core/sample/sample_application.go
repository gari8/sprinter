package sample

import (
	"github.com/gari8/sprinter"
	"@@.ImportPath@@/internal/@@.ImportPath@@/application/domain"
	"net/http"
)

type CreateSampleInput struct {
	Text string `json:"text,omitempty"`
}

type Repository interface {
	GetSamples() ([]domain.Sample, error)
	CreateSample(sample domain.Sample) error
}

type SampleApplication struct {
	repo Repository
}

func NewSampleApplication (
	repo Repository) SampleApplication {
	return SampleApplication{
		repo:       repo,
	}
}

func (a SampleApplication) GetSamples() sprinter.Response {
	samples, err := a.repo.GetSamples()
	if err != nil {
		return sprinter.Response{
			Code: http.StatusInternalServerError,
			Err: err,
		}
	}
	return sprinter.Response{
		Code:        http.StatusOK,
		Object:      samples,
	}
}

func (a SampleApplication) CreateSample(input CreateSampleInput) sprinter.Response {
	var sample domain.Sample
	sample.Text = input.Text
	err := a.repo.CreateSample(sample)
	if err != nil {
		return sprinter.Response{
			Code: http.StatusInternalServerError,
			Err: err,
		}
	}
	return sprinter.Response{
		Code: http.StatusOK,
	}
}

