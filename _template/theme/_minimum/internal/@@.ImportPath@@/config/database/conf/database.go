package conf

import (
	"database/sql"

	_ "github.com/lib/pq"

	"os"
)

var (
	source = os.Getenv("DATABASE_URL")
	driver = os.Getenv("DRIVER")
)

func NewDatabaseConnection() (*sql.DB, error) {
	conn, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
