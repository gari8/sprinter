package main

import (
	"fmt"
	"github.com/HAGARIHAYATO/sprinter/infrastructure/db/conf"
	"github.com/HAGARIHAYATO/sprinter/interactor"
	"github.com/HAGARIHAYATO/sprinter/presenter/middleware"
	"github.com/HAGARIHAYATO/sprinter/presenter/router"
	"net/http"
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
	repository := i.NewRepository()
	application := i.NewApplication(repository)
	handler := i.NewHandler(application)
	middleware := middleware.NewMiddleware()
	server := router.NewRouter()
	server.Router(handler, middleware)
	http.ListenAndServe(":8080", server.Route)
}
