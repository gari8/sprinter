package sample

import (
	"database/sql"
	"@@.ImportPath@@/internal/@@.ImportPath@@/domain"
)

type SampleRepository struct {
	conn *sql.DB
}

func NewSampleRepository(conn *sql.DB) SampleRepository {
	return SampleRepository{conn}
}

func (s SampleRepository) GetSamples() ([]domain.Sample, error) {
	var samples []domain.Sample
	rows, err := s.conn.Query("SELECT id, text FROM samples;")
	if rows == nil {
		return nil, err
	}
	for rows.Next() {
		var sample domain.Sample
		err = rows.Scan(&sample.ID, &sample.Text)
		if err == nil {
			samples = append(samples, sample)
		}
	}
	return samples, err
}

func (s SampleRepository) GetSample(id uint32) (domain.Sample, error) {
	var sample domain.Sample
	err := s.conn.QueryRow("SELECT id, text FROM samples WHERE id = $1", id).Scan(&sample.ID, &sample.Text)
	if err != nil {
		return sample, err
	}
	return sample, nil
}

func (s SampleRepository) CreateSample(sample domain.Sample) error {
	_, err := s.conn.Exec("INSERT INTO samples(text) VALUES ($1)", sample.Text)
	if err != nil {
		return err
	}
	return nil
}
