package sample

import (
	"github.com/gari8/sprinter"
	"@@.ImportPath@@/internal/@@.ImportPath@@/domain"
	"net/http"
)

type GetSamplePresenter struct {

}

func NewGetSamplePresenter() GetSamplePresenter {
	return GetSamplePresenter{}
}

func (p GetSamplePresenter) CreateResponse(samples []domain.Sample) sprinter.Response {
	return sprinter.Response{
		Code: http.StatusOK,
		Object: samples,
	}
}