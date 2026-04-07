package app

import (
	"github.com/hisshihi/url-shortener/core/config"
)

type App struct {
	diContainer *diContainer
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
		//
	}
	for _, fn := range inits {
		fn()
	}
}
