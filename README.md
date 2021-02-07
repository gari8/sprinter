# sprinter

### OverView

- this command line tool makes it easier for us to design the architecture .

```
go get github.com/gari8/sprinter
```

### Architecture

- Onion Architecture
- MVC Pattern

### How to use

```
sprinter -new
or
sprinter -n

      sprinter -new
      <conversation>

      cd go-app
      
      tree .
      
      .
      ├── Dockerfile
      ├── README.md
      ├── application
      │   ├── README.md
      │   └── sample_application.go
      ├── build.sh
      ├── docker-compose.yml
      ├── domain
      │   ├── model
      │   │   └── sample_model.go
      │   └── repository
      │       ├── README.md
      │       └── sample_repository.go
      ├── go.mod
      ├── infrastructure
      │   ├── mysql
      │   │   ├── conf
      │   │   │   └── database.go
      │   │   └── init
      │   │       └── 1_init.sql
      │   └── postgres
      │       ├── conf
      │       │   └── database.go
      │       └── init
      │           └── 1_init.sql
      ├── interactor
      │   ├── README.md
      │   └── interactor.go
      ├── main.go
      └── presenter
          ├── handler
          │   ├── README.md
          │   ├── handler_util.go
          │   └── sample_handler.go
          ├── middleware
          │   └── main.go
          ├── router
          │   ├── router.go
          │   └── router_test.go
          └── template
              ├── layout
              │   ├── _footer.html
              │   └── _header.html
              └── sample
                  └── index.html
```

```
docker-compose build

docker-compose up

open http://localhost:8080/api/v1 or http://localhost:8080
```
