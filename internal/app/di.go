package app

import (
	"log/slog"
	"os"

	"github.com/hisshihi/url-shortener/internal/closer"
	"github.com/hisshihi/url-shortener/internal/config"
	"github.com/hisshihi/url-shortener/internal/database"
	"github.com/hisshihi/url-shortener/internal/handler"
	"github.com/hisshihi/url-shortener/internal/repository"
	"github.com/hisshihi/url-shortener/internal/service"
)

type diContainer struct {
	// config
	cfg config.Config

	// db
	db *database.DB

	// repos
	urlRepo *repository.URLRepository

	// services
	urlService service.UrlShorterService

	//	http
	urlHandler *handler.Handler
}

func NewDIContainer(cfg config.Config) *diContainer {
	return &diContainer{cfg: cfg}
}

func (d *diContainer) DB() *database.DB {
	if d.db == nil {
		dsn := d.cfg.DSN()
		db, err := database.New(dsn)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}

		closer.Add("база данных", db.Close)

		d.db = db
	}
	return d.db
}

func (d *diContainer) URLRepo() *repository.URLRepository {
	if d.urlRepo == nil {
		d.urlRepo = repository.NewURLRepository(d.DB())
	}

	return d.urlRepo
}

func (d *diContainer) URLService() service.UrlShorterService {
	if d.urlService == nil {
		d.urlService = service.NewURLService(d.URLRepo())
	}

	return d.urlService
}

func (d *diContainer) URLHandler() *handler.Handler {
	if d.urlHandler == nil {
		d.urlHandler = handler.NewUrlHandler(d.URLService())
	}
	return d.urlHandler
}
