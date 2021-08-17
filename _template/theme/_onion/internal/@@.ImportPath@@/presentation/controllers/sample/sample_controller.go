package sample

import (
	"context"
	sampleApplicationService "@@.ImportPath@@/internal/@@.ImportPath@@/application/service/sample"
	"net/http"

	"github.com/gari8/sprinter"
)

type SampleController interface {
	GetSamples(ctx context.Context, r *http.Request) sprinter.Response
	CreateSample(ctx context.Context, r *http.Request) sprinter.Response
}

type sampleController struct {
	applicationService sampleApplicationService.SampleApplicationService
}

func NewSampleController(
	applicationService sampleApplicationService.SampleApplicationService) SampleController {
	return sampleController{
		applicationService: applicationService,
	}
}

func (c sampleController) GetSamples(ctx context.Context, r *http.Request) sprinter.Response {
	samples, err := c.applicationService.GetSamples()
	if err != nil {
		return sprinter.Response{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}
	return sprinter.Response{
		Code:   http.StatusOK,
		Object: samples,
	}
}

func (c sampleController) CreateSample(ctx context.Context, r *http.Request) sprinter.Response {
	var createSampleInput sampleApplicationService.CreateSampleInput
	sprinter.GetInputByJson(r.Body, &createSampleInput)
	c.applicationService.CreateSample(createSampleInput)
	return sprinter.Response{
		Code:   http.StatusOK,
		Object: nil,
	}
}
