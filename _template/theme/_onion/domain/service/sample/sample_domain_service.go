package sample

import "@@.ImportPath@@/domain/model"

type Repository interface {
	GetSamples() ([]model.Sample, error)
	CreateSample(sample model.Sample) error
	GetSample(id uint32) (model.Sample, error)
}
