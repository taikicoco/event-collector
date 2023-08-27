package repository

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func DBInit() (*sqlx.DB, error) {
	dsn := os.Getenv("DATASOURCE")
	if dsn == "" {
		return nil, errors.New("missing DATASOURCE environment variable")
	}

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := db.Ping(); err != nil {
		return nil, errors.WithStack(err)
	}
	return db, nil
}
