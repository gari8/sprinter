package conf

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DRIVER   = "mysql"
	HOSTNAME = "@@.ImportPath@@db"
	USER     = "mysql"
	DBNAME   = "mysql"
	PASSWORD = "mysql"
)


func NewDatabaseConnection() (*sql.DB, error) {
	source := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true", USER, PASSWORD, HOSTNAME, DBNAME)
	conn, err := sql.Open(DRIVER, source)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
