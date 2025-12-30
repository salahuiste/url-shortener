package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres() (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), "postgres://user:password@localhost:5432/shortener")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to connect to DB: %w", err))
	}
	return db, nil
}
