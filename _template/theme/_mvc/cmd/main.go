package main

import (
	"github.com/gari8/sprinter"
	"net/http"
	"os"

	"@@.ImportPath@@/internal/@@.ImportPath@@/config/database/conf"
	"@@.ImportPath@@/internal/@@.ImportPath@@/controller"
	"@@.ImportPath@@/internal/@@.ImportPath@@/model"
	"@@.ImportPath@@/internal/@@.ImportPath@@/router"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	conn, err := conf.NewDatabaseConnection()
	if err != nil {
		panic(err)
	}
	if conn == nil {
		panic(err)
	}
	defer func() {
		if conn != nil {
			if err := conn.Close(); err != nil {
				panic(err)
			}
		}
	}()

	sprinter.PrintLogo("API: GET http://localhost:8080/api/v1", "HTML: GET http://localhost:8080")

	m := model.NewModel(conn)
	c := controller.NewController(m)
	s := router.NewRouter()
	s.Router(c)

	_ = http.ListenAndServe(":"+port, s.Route)
}