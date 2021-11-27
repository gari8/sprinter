package controller

import (
	"@@.ImportPath@@/model"
)

type Controller struct {
	SampleController
}

func NewController(model model.Model) Controller {
	c := Controller{}
	c.SampleController = NewSampleController(model)
	return c
}
