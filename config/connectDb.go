package connectDb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

const dbUser = "root"
const dbPassword = ""
const dbName = "goCrud"
const ip = "127.0.0.1"
const dbPort = "3306"

func DbConnection() (*sql.DB, error) {
	dsn := dbUser + ":" + dbPassword + "@tcp(" + ip + ":" + dbPort + ")/" + dbName
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return db, nil
}
