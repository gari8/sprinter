package conf

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var (
	DRIVER   = os.Getenv("DRIVER")
	HOSTNAME = os.Getenv("HOSTNAME")
	USER     = os.Getenv("USER")
	DBNAME   = os.Getenv("DBNAME")
	PASSWORD = os.Getenv("PASSWORD")
	source   = fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", HOSTNAME, USER, DBNAME, PASSWORD)
)

func NewDatabaseConnection() (*sql.DB, error) {
	conn, err := sql.Open(DRIVER, source)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
