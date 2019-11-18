package controllers

import (
	"github.com/joaopandolfi/miia_api/dao"
	"github.com/joaopandolfi/miia_api/services"
)

// NewUserService - Factory
func NewUserService() services.UserService {
	return services.User{
		UserDAO: dao.User{},
	}
}
