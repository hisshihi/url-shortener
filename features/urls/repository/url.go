package repository

import (
	"context"
	"errors"
	"log/slog"

	"github.com/hisshihi/url-shortener/internal/database"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type URLDb interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

type URLRepository struct {
	db URLDb
}

func NewURLRepository(db URLDb) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Create(ctx context.Context, longURL, alias string) (string, error) {
	rows, err := r.db.Exec(ctx, "INSERT INTO urls (long_url, alias) VALUES ($1, $2)", longURL, alias)
	if err != nil {
		slog.Error("ошибка при сохранении url", slog.String("long_url", longURL), slog.String("alias", alias), slog.Any("err", err))
		return "", err
	}

	if rows.RowsAffected() == 0 {
		return "", database.ErrNotInserted
	}

	return alias, nil
}

func (r *URLRepository) SelectByAlias(ctx context.Context, alias string) (string, error) {
	var longURL string
	err := r.db.QueryRow(ctx, "SELECT long_url FROM urls WHERE alias = $1", alias).Scan(&longURL)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", database.ErrURLNotFound
		}
		return "", err
	}

	return longURL, nil
}
