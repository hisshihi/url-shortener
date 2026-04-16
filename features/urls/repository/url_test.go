package repository

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_urlRepository_Create(t *testing.T) {
	t.Run("успешно создаёт url и возвращает alias", func(t *testing.T) {
		tx := newTx(t)
		repo := NewURLRepository(tx)

		alias, err := repo.Create(context.Background(), "https://example.com/long/url", "ex")

		require.NoError(t, err)
		assert.Equal(t, "ex", alias)
	})

	t.Run("дублирование alias", func(t *testing.T) {
		tx := newTx(t)
		repo := NewURLRepository(tx)
		_, _ = repo.Create(context.Background(), "https://example.com/long/url", "ex")

		_, err := repo.Create(context.Background(), "https://example.com/long/url", "ex")
		require.Error(t, err)

	})
}

func Test_urlRepository_SelectByAlias(t *testing.T) {
	t.Run("успешно найден long url по alias", func(t *testing.T) {
		tx := newTx(t)
		repo := NewURLRepository(tx)
		var err error
		var alias, url string
		alias, err = repo.Create(context.Background(), "https://example.com/long/url", "ex")
		require.NoError(t, err)

		url, err = repo.SelectByAlias(context.Background(), alias)
		require.NoError(t, err)
		require.Equal(t, "https://example.com/long/url", url)
	})

	t.Run("long url не найден", func(t *testing.T) {
		tx := newTx(t)
		repo := NewURLRepository(tx)
		var err error
		var alias string
		alias, err = repo.Create(context.Background(), "https://example.com/long/url", "ex")
		require.NoError(t, err)

		_, err = repo.SelectByAlias(context.Background(), alias+" fake")
		require.Error(t, pgx.ErrNoRows)
	})

	t.Run("ошибка бд", func(t *testing.T) {
		tx := newTx(t)
		repo := NewURLRepository(tx)
		var err error
		var nilString string
		_, err = repo.Create(context.Background(), "https://example.com/long/url", "ex")
		require.NoError(t, err)

		nilString, err = repo.SelectByAlias(context.Background(), "")
		require.Error(t, err)
		assert.Equal(t, "", nilString)
	})
}
