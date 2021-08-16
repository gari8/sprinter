package sample

import (
	"context"
	"github.com/gari8/sprinter"
	sampleUseCase "@@.ImportPath@@/internal/@@.ImportPath@@/usecases/usecases/sample"
)

type SampleController interface {
	GetSamples(ctx context.Context) sprinter.Response
}

type sampleController struct {
	sampleUseCase.GetSampleInputPort
}

func NewSampleController(
	inputPort sampleUseCase.GetSampleInputPort,
	) SampleController {
	return sampleController{inputPort}
}

func (c sampleController) GetSamples(ctx context.Context) sprinter.Response {
	return c.GetSampleInputPort.GetSamples()
}
