package sample

import (
	"github.com/gari8/sprinter"
	"@@.ImportPath@@/internal/@@.ImportPath@@/domain"
)

type GetSampleOutputPort interface {
	CreateResponse(samples []domain.Sample) sprinter.Response
}
