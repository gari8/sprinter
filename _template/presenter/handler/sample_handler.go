package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"@@.ImportPath@@/application"
	"@@.ImportPath@@/domain/model"
)

type(
	sampleHandler struct {
		application.SampleApplication
	}
	SampleHandler interface {
		SampleIndex(w http.ResponseWriter, r *http.Request)
	}
	response struct {
		Status int
		Samples []*model.Sample
	}
)

func NewSampleHandler(as application.SampleApplication) SampleHandler {
	return &sampleHandler{as}
}

func (s *sampleHandler) SampleIndex(w http.ResponseWriter, r *http.Request) {
	samples, err := s.SampleApplication.GetSamples()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	resp := &response{
		Status: http.StatusOK,
		Samples: samples,
	}

	res, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	_ , _ = w.Write(res)
}
