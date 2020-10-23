package conf

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	DRIVER   = os.Getenv("DRIVER")
	HOSTNAME = os.Getenv("HOSTNAME")
	USER     = os.Getenv("USER")
	DBNAME   = os.Getenv("DBNAME")
	PASSWORD = os.Getenv("PASSWORD")
	source   = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true", USER, PASSWORD, HOSTNAME, DBNAME)
)

func NewDatabaseConnection() (*sql.DB, error) {
	conn, err := sql.Open(DRIVER, source)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
