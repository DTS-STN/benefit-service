package db

import (
	"database/sql"
	"fmt"

	"github.com/DTS-STN/benefit-service/config"
	_ "github.com/lib/pq"
)

// Connect returns an open database connection
func Connect(config *config.Database) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", config.Host, config.Port, config.DB, config.User, config.Password)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
