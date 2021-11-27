package sample

import (
	"github.com/gari8/sprinter"
)

type CreateSampleOutputPort interface {
	CreateResponse() sprinter.Response
}
