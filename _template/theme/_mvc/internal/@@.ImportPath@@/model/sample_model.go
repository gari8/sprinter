package model

import "database/sql"

type SampleModel struct {
	conn *sql.DB
}

func NewSampleModel(conn *sql.DB) SampleModel {
	return SampleModel{conn}
}

type Sample struct {
	ID   int64 `json:"id" db:"id"`
	Text string `json:"text" db:"text"`
}

func (s *SampleModel) Fetch() ([]*Sample, error) {
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
