package main

import (
	"fmt"
	"@@.ImportPath@@/internal/@@.ImportPath@@/config/database/conf"
	"@@.ImportPath@@/internal/@@.ImportPath@@/router"
	"github.com/gari8/sprinter"
	"net/http"
	"os"
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

	s := router.NewRouter(conn)
	s.Router()

	_ = http.ListenAndServe(":"+port, s.Route)
}
