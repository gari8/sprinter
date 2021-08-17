package sample

import (
	"context"
	"github.com/gari8/sprinter"
	sampleUseCase "@@.ImportPath@@/internal/@@.ImportPath@@/usecases/usecases/sample"
	"net/http"
)

type SampleController interface {
	GetSamples(ctx context.Context, r *http.Request) sprinter.Response
	CreateSample(ctx context.Context, r *http.Request) sprinter.Response
}

type sampleController struct {
	getSampleInputPort sampleUseCase.GetSampleInputPort
	createSampleInputPort sampleUseCase.CreateSampleInputPort
}

func NewSampleController(
	getSampleInputPort sampleUseCase.GetSampleInputPort,
	createSampleInputPort sampleUseCase.CreateSampleInputPort) SampleController {
	return sampleController{
		getSampleInputPort: getSampleInputPort,
		createSampleInputPort: createSampleInputPort,
	}
}

func (c sampleController) GetSamples(ctx context.Context, r *http.Request) sprinter.Response {
	return c.getSampleInputPort.GetSamples()
}

func (c sampleController) CreateSample(ctx context.Context, r *http.Request) sprinter.Response {
	var createSampleInput sampleUseCase.CreateSampleInput
	sprinter.GetInputByJson(r.Body, &createSampleInput)
	return c.createSampleInputPort.CreateSample(createSampleInput)
}
