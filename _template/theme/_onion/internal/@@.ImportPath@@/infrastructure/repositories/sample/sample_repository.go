package sample

import (
	"database/sql"
	"@@.ImportPath@@/internal/@@.ImportPath@@/domain/model"
	sampleService "@@.ImportPath@@/internal/@@.ImportPath@@/domain/service/sample"
)

type SampleRepository struct {
	Conn *sql.DB
}

func NewSampleRepository(Conn *sql.DB) sampleService.Repository {
	return SampleRepository{Conn}
}

func (s SampleRepository) GetSamples() ([]model.Sample, error) {
	var samples []model.Sample
	rows, err := s.Conn.Query("SELECT id, text FROM samples;")
	if rows == nil {
		return nil, err
	}
	for rows.Next() {
		var sample model.Sample
		err = rows.Scan(&sample.ID, &sample.Text)
		if err == nil {
			samples = append(samples, sample)
		}
	}
	return samples, err
}

func (s SampleRepository) GetSample(id uint32) (model.Sample, error) {
	var sample model.Sample
	err := s.Conn.QueryRow("SELECT id, text FROM samples WHERE id = $1", id).Scan(&sample.ID, &sample.Text)
	if err != nil {
		return sample, err
	}
	return sample, nil
}

func (s SampleRepository) CreateSample(sample model.Sample) error {
	_, err := s.Conn.Exec("INSERT INTO samples(text) VALUES ($1)", sample.Text)
	if err != nil {
		return err
	}
	return nil
}