package main

import (
	"fmt"
	"net/http"
	@@ if .DataBase -@@
	"@@.ImportPath@@/infrastructure/database/mysql/conf"
	@@ else @@
	"@@.ImportPath@@/infrastructure/database/postgres/conf"
    @@ end @@
	"@@.ImportPath@@/infrastructure/middleware"
	"@@.ImportPath@@/infrastructure/router"
	"@@.ImportPath@@/injector"
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

	i := injector.NewInjector(conn)
	r := i.NewRepository()
	a := i.NewUseCase(r)
	h := i.NewController(a)
	m := middleware.NewMiddleware()
	s := router.NewRouter()
	s.Router(h, m)

	_ = http.ListenAndServe(":8080", s.Route)
}
