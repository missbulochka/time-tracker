package app

import (
	"net/http"
	httpapp "time-tracker/internal/handler/http"
)

func (a *App) MustRunHTTPServer() {
	a.setupServer()
	if err := a.HTTPSrv.Start(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func (a *App) setupServer() {
	router := httpapp.NewRouter()
	router.WithMiddlewares().AddRoutes(a.log, a.useCase)

	a.HTTPSrv.RegisterRouts(router.GetRouter())
}
