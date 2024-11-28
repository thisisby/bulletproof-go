package users

import (
	"bulletproof-go/internal/core/service"
	"log"
)

type IUsersService interface{}

type UsersService struct {
	service.BaseService
}

func NewUserService() *UsersService {
	return &UsersService{}
}

func (es *UsersService) DoSomething() string {
	log.Print("Service Logic Executed")
	return "Service Logic Executed"
}
