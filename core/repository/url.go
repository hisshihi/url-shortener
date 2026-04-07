package repository

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type URLDb interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

type URLRepo interface {
	Create(ctx context.Context, longURL string) (string, error)
}

type URLRepository struct {
	db URLDb
}

func NewURLRepository(db URLDb) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Create(ctx context.Context, longURL string) (string, error) {
	rows, err := r.db.Exec(ctx, "INSERT INTO urls (long_url) VALUES ($1)", longURL)
	if err != nil {
		slog.Error("ошибка при сохранении url", longURL, err)
		return "", err
	}
	slog.Info("rows", rows)
	return "", nil
}
