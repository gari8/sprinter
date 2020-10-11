package conf

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	DRIVER = "postgres"
	HOSTNAME = "@@.ImportPath@@"
	USER = "postgres"
	DBNAME = "postgres"
	PASSWORD = "postgres"
)

func NewDatabaseConnection() (*sql.DB, error) {
	source := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", HOSTNAME, USER, DBNAME, PASSWORD)
	conn, err := sql.Open(DRIVER, source)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

