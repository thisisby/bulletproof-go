package root

import (
	"bulletproof-go/internal/core/di"
	"bulletproof-go/internal/core/module"
	"bulletproof-go/internal/core/router"
	"bulletproof-go/internal/modules/users"
)

type RootModule struct {
	module.Module
}

func NewRootModule(r *router.Router) *RootModule {
	// Initialize and collect all modules
	modules := []module.IModule{
		users.NewUsersModule(),
	}

	// Register all modules
	di.RegisterModules(r, modules)
	service := NewRootService()
	controller := NewRootController(service, "/")

	return &RootModule{
		Module: *module.NewModule(
			"Root",
			[]interface{}{controller},
			[]interface{}{service},
			modules,
		),
	}
}
