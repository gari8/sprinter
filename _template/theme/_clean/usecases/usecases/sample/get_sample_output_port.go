package sample

import (
	"@@.ImportPath@@/domain"
	"github.com/gari8/sprinter"
)

type GetSampleOutputPort interface {
	CreateResponse(samples []domain.Sample) sprinter.Response
}
