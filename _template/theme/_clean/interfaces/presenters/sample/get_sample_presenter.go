package sample

import (
	"net/http"

	"@@.ImportPath@@/domain"
	"github.com/gari8/sprinter"
)

type GetSamplePresenter struct{}

func NewGetSamplePresenter() GetSamplePresenter {
	return GetSamplePresenter{}
}

func (p GetSamplePresenter) CreateResponse(samples []domain.Sample) sprinter.Response {
	return sprinter.Response{
		Code:   http.StatusOK,
		Object: samples,
	}
}
