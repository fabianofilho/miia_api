package services

import (
	"github.com/joaopandolfi/blackwhale/utils"
	"github.com/joaopandolfi/miia_api/dao"
	"github.com/joaopandolfi/miia_api/models"
)

type UserService interface {
	Login(username string, password string) (user models.User, success bool, err error)
	NewUserClient(user models.User) (result models.User, success bool, err error)
}

type User struct {
	UserDAO dao.UserDAO
}

func (cc User) Login(username string, password string) (user models.User, success bool, err error) {
	return cc.UserDAO.Login(models.User{Username: username, Password: password})
}

// New basic client user
func (cc User) NewUserClient(user models.User) (result models.User, success bool, err error) {
	user.Level = models.USER_CLIENT
	user.Password, err = utils.HashPassword(user.Password)

	if err != nil {
		success = false
		user.Password = ""
		return
	}
	return cc.UserDAO.NewUser(user)
}
