package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"@@.ImportPath@@/model"
)

type SampleController struct {
	model.Model
}

func NewSampleController(sm model.Model) SampleController {
	return SampleController{sm}
}

func (s SampleController) SampleIndex(w http.ResponseWriter, r *http.Request) {
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

	_, _ = w.Write(res)
}

func (s SampleController) SampleHTML(w http.ResponseWriter, r *http.Request) {
	tmpl, err := parseTemplate("sample", "index")

	if err != nil {
		log.Fatal("err :", err)
	}

	if err := tmpl.Execute(w, nil); err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}
