package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/hisshihi/url-shortener/internal/closer"
	"github.com/hisshihi/url-shortener/internal/config"
)

type App struct {
	diContainer *diContainer
	httpServer  *http.Server
}

func New(cfg config.Config) *App {
	a := &App{
		diContainer: NewDIContainer(cfg),
	}

	a.initDeps()
	return a
}

func (a *App) initDeps() {
	inits := []func(){
		a.initHTTPServer,
	}
	for _, fn := range inits {
		fn()
	}
}

func (a *App) initHTTPServer() {
	a.httpServer = &http.Server{
		Addr:         a.diContainer.cfg.HTTPAddr,
		Handler:      a.diContainer.URLHandler().Routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

func (a *App) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	slog.Info("сервер запущен", "addr", a.diContainer.cfg.HTTPAddr)

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("ошибка сервера", "err", err)
		}
	}()

	<-ctx.Done()
	slog.Info("получен сигнал, завершаем...")

	// Паттер "двойной Ctrl + C": снимаем custom handler. Второй Ctrl + C сразу убъёт процесс
	stop()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	if err := a.httpServer.Shutdown(shutdownCtx); err != nil {
		slog.Error("ошибка при остановке сервера", "err", err)
	}

	slog.Info("сервер остановлен")

	closerCtx, closerCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer closerCancel()

	if err := closer.CloseAll(closerCtx); err != nil {
		slog.Error("ошибки при закрытии ресурсов", "err", err)
	}

	return nil
}
