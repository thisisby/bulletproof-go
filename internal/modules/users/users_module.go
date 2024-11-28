package users

import "bulletproof-go/internal/core/module"

type UsersModule struct {
	module.Module
}

func NewUsersModule() *UsersModule {
	usersService := NewUserService()
	usersController := NewUsersController(usersService, "/users")
	return &UsersModule{
		Module: *module.NewModule(
			"Users",
			[]interface{}{usersController},
			[]interface{}{usersService},
			[]module.IModule{},
		),
	}
}
