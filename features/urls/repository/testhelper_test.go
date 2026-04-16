package repository

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// testPool - одна на весь пакет, поднимается в TestMain
var testPool *pgxpool.Pool

func TestMain(m *testing.M) {
	ctx := context.Background()

	ctr, err := postgres.Run(ctx,
		"postgres:18.3-alpine3.23",
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("test"),
		postgres.WithPassword("test"),
		testcontainers.WithWaitStrategy(
			// Ждём пока postgres реально готов принимать соединения
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2),
		),
	)
	if err != nil {
		log.Fatalf("не удалось поднять контейнер: %v", err)
	}
	defer func() {
		if err := ctr.Terminate(ctx); err != nil {
			log.Fatalf("ошибка при остановке контейнера: %v", err)
		}
	}()

	connStr, err := ctr.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatalf("не удалось получить connection string: %v", err)
	}

	testPool, err = pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("не удалось создать пул: %v", err)
	}
	defer testPool.Close()

	if err := runMigrations(ctx, testPool); err != nil {
		log.Fatalf("не удалось накатить миграции: %v", err)
	}

	// m.Run() — запускает все Test* функции в пакете
	// os.Exit нужен чтобы код выхода корректно пробросился
	code := m.Run()
	os.Exit(code)
}

func runMigrations(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS urls (
	    id BIGSERIAL PRIMARY KEY ,
    alias VARCHAR(16) NOT NULL UNIQUE ,
    long_url TEXT NOT NULL ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
	)
`)
	return err
}

// newTx создаёт транзакцию для одного теста.
// После теста делает Rollback — следующий тест получает чистую БД.
// Это быстрее чем TRUNCATE и не требует порядка выполнения тестов.
func newTx(t *testing.T) pgx.Tx {
	t.Helper()

	tx, err := testPool.Begin(context.Background())
	if err != nil {
		t.Fatalf("не удалось начать транзакцию: %v", err)
	}

	// t.Cleanup вызывается автоматически когда тест завершается
	t.Cleanup(func() {
		_ = tx.Rollback(context.Background())
	})

	return tx
}
