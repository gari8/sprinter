package main

import (
	"fmt"
	"net/http"
	"server/config/database/conf"
	"server/interactor"
	"server/presenter/middleware"
	"server/presenter/router"
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

	fmt.Println(`
    * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * 
  *        ####    #####    #####     ####    ##  ##   ######   ######   #####  *
  *      ##  ##   ##  ##   ##  ##     ##     ### ##     ##     ##       ##  ##  *
  *     ##       ##  ##   ##  ##     ##     ######     ##     ##       ##  ##   *
  *     ####    #####    #####      ##     ######     ##     ####     #####     *
  *       ##   ##       ####       ##     ## ###     ##     ##       ####       *
  *  ##  ##   ##       ## ##      ##     ##  ##     ##     ##       ## ##       *
  *  ####    ##       ##  ##    ####    ##  ##     ##     ######   ##  ##       *
    * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

	HTML:	GET http://localhost:8080
    API:	GET http://localhost:8080/api/v1
    `)

	i := interactor.NewInteractor(conn)
	r := i.NewRepository()
	a := i.NewApplication(r)
	h := i.NewHandler(a)
	m := middleware.NewMiddleware()
	s := router.NewRouter()
	s.Router(h, m)

	_ = http.ListenAndServe(":"+port, s.Route)
}