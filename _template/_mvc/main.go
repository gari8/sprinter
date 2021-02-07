package main

import (
	"@@.ImportPath@@/controllers"
	"@@.ImportPath@@/models"
	"fmt"
	"net/http"
	@@ if .DataBase -@@
	"@@.ImportPath@@/database/mysql/conf"
	@@ else @@
	"@@.ImportPath@@/database/database/conf"
	@@ end @@
	"@@.ImportPath@@/router"
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
	fmt.Println(`API:	GET http://localhost:8080/server/v1`)

	m := models.NewModel(conn)
	c := controllers.NewController(m)
	s := router.NewRouter()
	s.Router(c)

	_ = http.ListenAndServe(":8080", s.Route)
}
