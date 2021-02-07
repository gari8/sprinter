package models

import "database/sql"

type Model struct {
	SampleModel
}

func NewModel(conn *sql.DB) *Model {
	m := &Model{}
	m.SampleModel = NewSampleModel(conn)
	return m
}
