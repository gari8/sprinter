package main

import (
	"fmt"
	"net/http"
	//"@@.ImportPath@@/infrastructure/postgres/conf"
	//"@@.ImportPath@@/interactor"
	//"@@.ImportPath@@/presenter/middleware"
	//"@@.ImportPath@@/presenter/router"
)

func main() {
	conn, err := conf.NewDatabaseConnection()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
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
    `)

	fmt.Println(`	GET http://localhost:8080/api/v1`)
	i := interactor.NewInteractor(conn)
	r := i.NewRepository()
	a := i.NewApplication(r)
	h := i.NewHandler(a)
	m := middleware.NewMiddleware()
	s := router.NewRouter()
	s.Router(h, m)

	_ = http.ListenAndServe(":8080", s.Route)
}
