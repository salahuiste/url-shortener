package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type URLRepository struct {
	DB *pgxpool.Pool
}

func NewURLRepository(db *pgxpool.Pool) *URLRepository {
	return &URLRepository{DB: db}
}

func (r *URLRepository) Save(ctx context.Context, original, code string) error {
	_, err := r.DB.Exec(ctx, "INSERT INTO urls (original_url, short_code) VALUES ($1,$2)", original, code)
	return err
}

func (r *URLRepository) Get(ctx context.Context, code string) (string, error) {
	var url string
	err := r.DB.QueryRow(ctx, "SELECT original_url FROM urls WHERE short_code=$1", code).Scan(&url)
	return url, err
}
