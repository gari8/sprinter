package controllers

import "database/sql"

type (
	Controller struct {
		SampleController
	}
)

func NewController(conn *sql.DB) *Controller {
	r := &Controller{}
	r.SampleController = NewSampleController(conn)
	return r
}

