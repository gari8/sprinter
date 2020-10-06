# sprinter

### OverView

- this command line tool makes it easier for us to design the architecture .

```
go get github.com/HAGARIHAYATO/sprinter
```

### Architecture

- Onion Architecture

### How to use

```
sprinter -path=<application-name>

ex... sprinter -path=go-app

      cd go-app
      
      tree .
      
      .
      ├── Dockerfile
      ├── README.md
      ├── application
      │   └── sample_application.go
      ├── build.sh
      ├── docker-compose.yml
      ├── domain
      │   ├── model
      │   │   └── sample_model.go
      │   └── repository
      │       └── sample_repository.go
      ├── go.mod
      ├── infrastructure
      │   └── postgres
      │       ├── conf
      │       │   └── database.go
      │       └── init
      │           └── 1_init.sql
      ├── interactor
      │   └── interactor.go
      ├── main.go
      └── presenter
          ├── handler
          │   └── sample_handler.go
          ├── middleware
          │   └── main.go
          └── router
              ├── router.go
              └── router_test.go
```

```
docker-compose build

docker-compose up

opne http://localhost:8080/api/v1
```
