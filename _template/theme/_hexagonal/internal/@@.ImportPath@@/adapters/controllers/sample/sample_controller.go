package sample

import (
	"context"
	"github.com/gari8/sprinter"
	sampleApplication "@@.ImportPath@@/internal/@@.ImportPath@@/application/core/sample"
	"net/http"
)

type SampleController struct {
	application sampleApplication.SampleApplication
}

func NewSampleController (
	application sampleApplication.SampleApplication) SampleController {
	return SampleController{
		application: application,
	}
}

func (c SampleController) GetSamples(ctx context.Context, r *http.Request) sprinter.Response {
	return c.application.GetSamples()
}

func (c SampleController) CreateSample(ctx context.Context, r *http.Request) sprinter.Response {
	var createSampleInput sampleApplication.CreateSampleInput
	sprinter.GetInputByJson(r.Body, &createSampleInput)
	return c.application.CreateSample(createSampleInput)
}
