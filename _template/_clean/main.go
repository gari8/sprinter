package main

import (
	@@ if .DataBase -@@
	"@@.ImportPath@@/infrastructure/database/mysql/conf"
	@@ else @@
	"@@.ImportPath@@/infrastructure/database/postgres/conf"
    @@ end @@
	"@@.ImportPath@@/interfaces/controllers"
	"@@.ImportPath@@/interfaces/presenter/handler"
	"@@.ImportPath@@/usecase"
	"fmt"
	"net/http"

	"@@.ImportPath@@/infrastructure/middleware"
	"@@.ImportPath@@/infrastructure/router"
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

	c := controllers.NewController(conn)
	u := usecase.NewUseCase(c)
	h := handler.NewHandler(u)
	m := middleware.NewMiddleware()
	s := router.NewRouter()
	s.Router(h, m)

	_ = http.ListenAndServe(":8080", s.Route)
}
