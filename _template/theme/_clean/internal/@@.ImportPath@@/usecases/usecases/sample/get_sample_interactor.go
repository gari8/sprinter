package sample

import (
	"github.com/gari8/sprinter"
	"@@.ImportPath@@/internal/@@.ImportPath@@/usecases/services/sample"
	"net/http"
)

type GetSampleInteractor struct {
	outputPort GetSampleOutputPort
	sampleService sample.SampleService
}

func NewGetSampleInteractor(
	outputPort GetSampleOutputPort,
	sampleService sample.SampleService,
	) GetSampleInteractor {
	return GetSampleInteractor{
		outputPort: outputPort,
		sampleService: sampleService,
	}
}

func (i GetSampleInteractor) GetSamples() sprinter.Response {
	samples, err := i.sampleService.GetSamples()
	if err != nil {
		return sprinter.Response{
			Code: http.StatusInternalServerError,
			Err: err,
		}
	}
	return i.outputPort.CreateResponse(samples)
}
