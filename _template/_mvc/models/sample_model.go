package models

import "database/sql"

type sampleModel struct{
	conn *sql.DB
}

type SampleModel interface {
	Fetch() ([]*Sample, error)
}

func NewSampleModel(conn *sql.DB) SampleModel {
	return &sampleModel{conn}
}

type Sample struct {
	ID   int64
	Text string
}

func (s *sampleModel) Fetch() ([]*Sample, error) {
	var samples []*Sample
	rows, err := s.conn.Query("SELECT id, text FROM samples;")
	if rows == nil {
		return nil, err
	}
	for rows.Next() {
		sample := &Sample{}
		err = rows.Scan(&sample.ID, &sample.Text)
		if err == nil {
			samples = append(samples, sample)
		}
	}
	return samples, err
}
