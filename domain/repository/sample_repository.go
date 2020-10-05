package repository

import (
	"database/sql"
	"sprinter/domain/model"
)

type (
	sampleRepository struct{
		conn *sql.DB
	}
	SampleRepository interface {
		Fetch() ([]*model.Sample, error)
	}
)

func NewSampleRepository(Conn *sql.DB) SampleRepository {
	return &sampleRepository{Conn}
}

func (s *sampleRepository) Fetch() ([]*model.Sample, error) {
	var samples []*model.Sample
	rows, err := s.conn.Query("SELECT id, text FROM samples;")
	if rows == nil { return nil, err }
	for rows.Next() {
		sample := &model.Sample{}
		err = rows.Scan(&sample.ID, &sample.Text)
		if err == nil {
			samples = append(samples, sample)
		}
	}
	return samples, err
}