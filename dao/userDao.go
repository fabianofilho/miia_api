package dao

import (
	"github.com/joaopandolfi/blackwhale/remotes/mysql"
	"github.com/joaopandolfi/blackwhale/utils"
	"github.com/joaopandolfi/miia_api/models"
)

// Dao responsavel por acessar MySQL e persistir os dados do usuario

type UserDAO interface {
	NewUser(user models.User) (result models.User, success bool, err error)
	Login(user models.User) (result models.User, success bool, err error)
	CheckToken(user models.User) (success bool, err error)
}

type User struct {
}

// NewUser -
// Create new user on database
func (cc User) NewUser(user models.User) (result models.User, success bool, err error) {
	//cpf,nome,email,login,senha,tipo_usuario,idinstituicao
	id, err := mysql.Driver.ExecuteAndReturnLastId("SELECT new_user(?,?,?,?,?,?,?);", user.CPF, user.Name, user.Email, user.Username, user.Password, user.Level, user.Instution)
	user.Password = ""
	user.ID = id
	success = false

	if err == nil {
		result = user
		success = true
	}
	return
}

func (cc User) CheckToken(user models.User) (success bool, err error) {
	var users []models.MysqlUser
	success = false

	err = mysql.Driver.Execute(mysql.THE_CASE, &users,
		`SELECT * FROM user WHERE token  = ?`, user.Token)

	if err != nil {
		utils.CriticalError("[UserDAO][CheckToken] - Error on Token query", err.Error())
		return
	}

	success = len(users) > 0
	return
}

func (cc User) Login(user models.User) (result models.User, success bool, err error) {
	var users []models.MysqlUser
	success = false

	err = mysql.Driver.Execute(mysql.THE_CASE, &users,
		`SELECT * FROM user WHERE username  = ?`, user.Username)

	if err != nil {
		utils.CriticalError("[UserDAO][Login] - Error on login query", err.Error())
		return
	}

	if len(users) > 0 {
		if utils.CheckPasswordHash(user.Password, users[0].Password) {
			result = users[0].ParseSql()
			result.Password = ""
			success = true
			return
		}
	}

	return
}
