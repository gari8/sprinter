package controllers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type response struct {
	Status int
	Result interface{}
}

func parseTemplate(dir string, fileName string) (*template.Template, error) {
	tmpl := template.New("")

	var layout string

	if err := filepath.Walk("interfaces/template/layout", func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.IsDir() && (strings.HasSuffix(path, ".html") || strings.HasSuffix(path, ".js")) {
			file, err := ioutil.ReadFile(path)

			if err != nil {
				return err
			}

			layout += string(file)
		}

		return nil
	}); err != nil {
		fmt.Println(err)
	}

	if err := filepath.Walk("interfaces/template/"+dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.IsDir() && (strings.HasSuffix(path, ".html") || strings.HasSuffix(path, ".js")) {
			file, err := ioutil.ReadFile(path)

			if err != nil {
				return err
			}

			filename := strings.Replace(path, "interfaces/template/"+dir, "", -1)

			if strings.Contains(filename, fileName) {
				tmpl = tmpl.New(filename)

				tmpl, err = tmpl.Parse(string(file) + layout)

				if err != nil {
					return err
				}
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return tmpl, nil
}
