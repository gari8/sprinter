package main

import (
	"@@.ImportPath@@/infrastructure/database/conf"
	"@@.ImportPath@@/presentation/server"

	"github.com/gari8/sprinter"
)

func main() {
	conn, err := conf.NewDatabaseConnection()
	if err != nil {
		panic(err)
	}
	s := server.NewServer(conn)
	sprinter.PrintLogo("GET http://localhost:8080/api/v1/sample", "POST http://localhost:8080/api/v1/sample")
	s.Serve()
}
