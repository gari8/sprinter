package sample

import (
	"github.com/gari8/sprinter"
)

type GetSampleInputPort interface {
	GetSamples () sprinter.Response
}
