package conf

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
)

var (
	DRIVER = "postgres"
	HOSTNAME = "db"
	USER = "postgres"
	DBNAME = "devdb"
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

