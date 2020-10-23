package controllers

import (
	"@@.ImportPath@@/domain/model"
	"database/sql"
)

type (
	sampleController struct {
		conn *sql.DB
	}
	SampleController interface {
		Fetch() ([]*model.Sample, error)
	}
)

func NewSampleController(Conn *sql.DB) SampleController {
	return &sampleController{Conn}
}

func (s *sampleController) Fetch() ([]*model.Sample, error) {
	var samples []*model.Sample
	rows, err := s.conn.Query("SELECT id, text FROM samples;")
	if rows == nil {
		return nil, err
	}
	for rows.Next() {
		sample := &model.Sample{}
		err = rows.Scan(&sample.ID, &sample.Text)
		if err == nil {
			samples = append(samples, sample)
		}
	}
	return samples, err
}
