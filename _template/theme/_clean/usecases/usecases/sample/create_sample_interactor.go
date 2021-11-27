package sample

import (
	"net/http"

	"@@.ImportPath@@/domain"
	"@@.ImportPath@@/usecases/services/sample"
	"github.com/gari8/sprinter"
)

type CreateSampleInteractor struct {
	outputPort    CreateSampleOutputPort
	sampleService sample.SampleService
}

func NewCreateSampleInteractor(
	outputPort CreateSampleOutputPort,
	sampleService sample.SampleService,
) CreateSampleInteractor {
	return CreateSampleInteractor{
		outputPort:    outputPort,
		sampleService: sampleService,
	}
}

func (i CreateSampleInteractor) CreateSample(input CreateSampleInput) sprinter.Response {
	var sample domain.Sample
	sample.Text = input.Text
	err := i.sampleService.CreateSample(sample)
	if err != nil {
		return sprinter.Response{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}
	return i.outputPort.CreateResponse()
}
