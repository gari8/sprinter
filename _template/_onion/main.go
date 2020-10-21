package main

import (
	"fmt"
	"net/http"
	@@ if .DataBase -@@
	"@@.ImportPath@@/infrastructure/mysql/conf"
	@@ else @@
	"@@.ImportPath@@/infrastructure/postgres/conf"
	@@ end @@
	"@@.ImportPath@@/interactor"
	"@@.ImportPath@@/presenter/middleware"
	"@@.ImportPath@@/presenter/router"
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

	fmt.Println(`HTML:	GET http://localhost:8080`)
	fmt.Println(`API:	GET http://localhost:8080/api/v1`)

	i := interactor.NewInteractor(conn)
	r := i.NewRepository()
	a := i.NewApplication(r)
	h := i.NewHandler(a)
	m := middleware.NewMiddleware()
	s := router.NewRouter()
	s.Router(h, m)

	_ = http.ListenAndServe(":8080", s.Route)
}
