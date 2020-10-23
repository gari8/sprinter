package handler

import (
	"@@.ImportPath@@/usecase"
	"encoding/json"
	"log"
	"net/http"
)

type (
	sampleHandler struct {
		usecase.SampleUseCase
	}
	SampleHandler interface {
		SampleIndex(w http.ResponseWriter, r *http.Request)
		SampleHTML(w http.ResponseWriter, r *http.Request)
	}
)

func NewSampleHandler(as usecase.SampleUseCase) SampleHandler {
	return &sampleHandler{as}
}

func (s *sampleHandler) SampleIndex(w http.ResponseWriter, r *http.Request) {
	samples, err := s.SampleUseCase.GetSamples()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	resp := &response{
		Status: http.StatusOK,
		Result: samples,
	}

	res, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	_, _ = w.Write(res)
}

func (s *sampleHandler) SampleHTML(w http.ResponseWriter, r *http.Request) {
	tmpl, err := parseTemplate("sample", "index")

	if err != nil {
		log.Fatal("err :", err)
	}

	if err := tmpl.Execute(w, nil); err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}
