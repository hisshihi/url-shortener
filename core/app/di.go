package app

import (
	"log/slog"
	"os"

	"github.com/hisshihi/url-shortener/core/config"
	"github.com/hisshihi/url-shortener/core/database"
	"github.com/hisshihi/url-shortener/core/repository"
	"github.com/hisshihi/url-shortener/core/service"
)

type diContainer struct {
	cfg config.Config

	db database.DB

	urlRepo repository.URLRepo

	urlService service.URLService
}

func NewDIContainer(cfg config.Config) *diContainer {
	return &diContainer{cfg: cfg}
}

func (d *diContainer) DB() database.DB {
	if d.db == nil {
		dsn := d.cfg.DSN()
		db, err := database.New(dsn)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		d.db = db
	}
	return d.db
}

func (d *diContainer) URlRepo() repository.URLRepo {
	if d.urlRepo == nil {
		d.urlRepo = repository.NewURLRepository(d.DB())
	}

	return d.urlRepo
}

func (d *diContainer) URLService() service.URLService {
	if d.urlService == nil {
		d.urlService = service.NewURLService(d.URlRepo())
	}

	return d.urlService
}
