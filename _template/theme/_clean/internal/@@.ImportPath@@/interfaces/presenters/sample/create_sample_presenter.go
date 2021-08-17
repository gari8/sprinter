package sample

import (
	"github.com/gari8/sprinter"
	"net/http"
)

type CreateSamplePresenter struct {}

func NewCreateSamplePresenter() CreateSamplePresenter {
	return CreateSamplePresenter{}
}

func (p CreateSamplePresenter) CreateResponse() sprinter.Response {
	return sprinter.Response{
		Code: http.StatusOK,
		Object: nil,
	}
}