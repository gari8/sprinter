### OverView

- This file is a DI(Dependency Injection) container

- It has each layer of Repository, Application, and Presenter as a structure.
Each structure is initialized in the NewXXX method of that structure.

### How to use

- When adding a new model structure

```
ex... 
// When example is added to model
type Example struct {
    Name string
}
```

add example_repository.go in /domain/repository/ 

```
// example_repository.go
package repository

import (
	"database/sql"
	"$YourProjectName/domain/model"
)

type (
	exampleRepository struct{
		conn *sql.DB
	}
	ExampleRepository interface {
		Fetch() ([]*model.Example, error)
	}
)

func NewExampleRepository(Conn *sql.DB) ExampleRepository {
	return &exampleRepository{Conn}
}

func (s *exampleRepository) Fetch() ([]*model.Example, error) {
    // TODO
}
```

add example_application.go in /application/

```
package application

import (
	"$YourProjectName/domain/model"
	"$YourProjectName/domain/repository"
)

type (
	exampleApplication struct {
		repository.ExampleRepository
	}
	ExampleApplication interface {
		GetExamples() ([]*model.Example, error)
	}
)

func NewExampleApplication(rs repository.ExampleRepository) ExampleApplication {
	return &exampleApplication{rs}
}

func (s *exampleApplication) GetExamples() ([]*model.Example, error) {
	return s.ExampleRepository.Fetch()
}
```

add example_handler.go in /presenter/handler/

```
package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"$YourProjectName/application"
	"$YourProjectName/domain/model"
)

type(
	exampleHandler struct {
		application.ExampleApplication
	}
	ExampleHandler interface {
		ExampleIndex(w http.ResponseWriter, r *http.Request)
	}
	response struct {
		Status int
		Examples []*model.Example
	}
)

func NewExampleHandler(as application.ExampleApplication) ExampleHandler {
	return &exampleHandler{as}
}

func (s *exampleHandler) ExampleIndex(w http.ResponseWriter, r *http.Request) {
    // TODO
}
```

write additional interfaces in interactor.go

```
ex...

type Repository struct {
	repository.SampleRepository
    repository.ExampleRepository // additional codes
}

...

func (i *interactor) NewRepository() *Repository {
	r := &Repository{}
	r.SampleRepository = repository.NewSampleRepository(i.conn)
	r.ExampleRepository = repository.NewExampleRepository(i.conn)
	return r
}

...

```