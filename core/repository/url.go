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
	Create(ctx context.Context, longURL, alias string) (string, error)
}

type urlRepository struct {
	db URLDb
}

func NewURLRepository(db URLDb) URLRepo {
	return &urlRepository{db: db}
}

func (r *urlRepository) Create(ctx context.Context, longURL, alias string) (string, error) {
	rows, err := r.db.Exec(ctx, "INSERT INTO urls (long_url, alias) VALUES ($1, $2)", longURL, alias)
	if err != nil {
		slog.Error("ошибка при сохранении url", slog.String("long_url", longURL), slog.String("alias", alias), slog.Any("err", err))
		return "", err
	}
	slog.Info("rows", slog.String("rows", rows.String()))
	return alias, nil
}
