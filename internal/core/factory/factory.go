package factory

import (
	"bulletproof-go/internal/core/config"
	"bulletproof-go/internal/core/di"
	"bulletproof-go/internal/core/router"
)

// App represents the main application structure
type App struct {
	Router    *router.Router
	Container *di.Container
}

// NewApp creates a new App instance
func NewApp(router *router.Router, container *di.Container) *App {
	return &App{
		Router:    router,
		Container: container,
	}
}

// Start the HTTP server on the specified address
func (app *App) Start(addr string, httpConfig *config.HttpConfig) error {
	return app.Router.Listen(addr, httpConfig)
}
