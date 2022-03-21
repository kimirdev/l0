package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var (
	username = "postgres"
	password = "123456"
	host     = "localhost"
	port     = "5432"
	dbname   = "postgres"
	sslmode  = "disable"
)

func NewPostgresDB() (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, username, dbname, password, sslmode)

	db, err := sqlx.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
