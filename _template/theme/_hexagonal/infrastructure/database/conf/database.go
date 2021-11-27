package conf

import (
	"database/sql"
@@ if eq .DataBase "MySQL" -@@
_ "github.com/go-sql-driver/mysql"
@@ else @@
_ "github.com/lib/pq"
@@ end @@
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
