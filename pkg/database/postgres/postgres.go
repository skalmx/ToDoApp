package postgres

import (
	"database/sql"
	"fmt"
)

func NewPostgresConnection() (*sql.DB, error) {
	connection := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s", "localhost", 5432, "postgres", "postgres", "disable", "123")
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}