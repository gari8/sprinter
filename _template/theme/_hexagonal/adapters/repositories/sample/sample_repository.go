package sample

import (
	"database/sql"

	"@@.ImportPath@@/application/domain"
)

type SampleRepository struct {
	Conn *sql.DB
}

func NewSampleRepository(Conn *sql.DB) SampleRepository {
	return SampleRepository{
		Conn: Conn,
	}
}

func (r SampleRepository) GetSamples() ([]domain.Sample, error) {
	var samples []domain.Sample
	rows, err := r.Conn.Query("SELECT id, text FROM samples ORDER BY id;")
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

func (r SampleRepository) CreateSample(sample domain.Sample) error {
	_, err := r.Conn.Exec("INSERT INTO samples(text) VALUES ($1)", sample.Text)
	if err != nil {
		return err
	}
	return nil
}
