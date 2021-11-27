package sample

import "github.com/gari8/sprinter"

type CreateSampleInput struct {
	Text string `json:"text,omitempty" db:"text"`
}

type CreateSampleInputPort interface {
	CreateSample (input CreateSampleInput) sprinter.Response
}
