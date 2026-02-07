package app

import "net/http"

type App struct {
	server *http.Server
}