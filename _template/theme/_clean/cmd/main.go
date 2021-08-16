package main

import (
	"@@.ImportPath@@/internal/@@.ImportPath@@/infrastructure/database/conf"
	"@@.ImportPath@@/internal/@@.ImportPath@@/infrastructure/server"
)

func main() {
	conn, err := conf.NewDatabaseConnection()
	if err != nil {
		panic(err)
	}
	s := server.NewServer(conn)
	s.Serve()
}
