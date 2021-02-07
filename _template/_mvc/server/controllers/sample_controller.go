package controllers

import (
	"server/models"
	"encoding/json"
	"log"
	"net/http"
)

type sampleController struct {
	models.SampleModel
}

type SampleController interface {
	SampleIndex(w http.ResponseWriter, r *http.Request)
	SampleHTML(w http.ResponseWriter, r *http.Request)
}

func NewSampleController(sm models.SampleModel) SampleController {
	return &sampleController{sm}
}

func (s *sampleController) SampleIndex(w http.ResponseWriter, r *http.Request) {
	samples, err := s.SampleModel.Fetch()

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

	_ , _ = w.Write(res)
}

func (s *sampleController) SampleHTML(w http.ResponseWriter, r *http.Request) {
	tmpl, err := parseTemplate("sample", "index")

	if err != nil {
		log.Fatal("err :", err)
	}

	if err := tmpl.Execute(w, nil); err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}

