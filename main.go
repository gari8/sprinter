package main

import (
	"fmt"
	"net/http"
	"sprinter/infrastructure/postgres/conf"
	"sprinter/interactor"
	"sprinter/presenter/middleware"
	"sprinter/presenter/router"
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

	i := interactor.NewInteractor(conn)
	r := i.NewRepository()
	a := i.NewApplication(r)
	h := i.NewHandler(a)
	m := middleware.NewMiddleware()
	s := router.NewRouter()
	s.Router(h, m)

	_ = http.ListenAndServe(":8080", s.Route)
}
