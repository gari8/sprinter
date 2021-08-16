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
