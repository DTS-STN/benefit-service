package db

import (
	"database/sql"
	"os"

	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

func OpenDB() *sql.DB {
	connStr := os.Getenv("DB_CONN")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error(err)
	}

	return db
}
