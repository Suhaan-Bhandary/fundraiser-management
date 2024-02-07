package repository

import (
	"context"
	"os"

	"database/sql"

	_ "github.com/lib/pq"
)

func InitializeDatabase(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("PSQL_INFO"))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db, nil
}
