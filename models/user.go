package models

import "strconv"

const USER_ROOT int = 1
const USER_ADMIN int = 2
const USER_ESPECIALIST int = 3
const USER_DOCTOR int = 4
const USER_SECRETARY int = 5
const USER_CLIENT int = 6

const SESSION_VALUE_INSTITUTION string = "institution"
const SESSION_VALUE_TOKEN string = "token"
const SESSION_VALUE_LEVEL string = "level"
const SESSION_VALUE_USERNAME string = "username"
const SESSION_VALUE_ID string = "user_id"
const SESSION_VALUE_NAME string = "name"
const SESSION_VALUE_LOGGED string = "logged"
const SESSION_VALUE_SPECIALTY string = "specialty"

type User struct {
	People
	Email     string `json:"email"`
	Username  string `json:"login"`
	Token     string `json:"token"`
	Picture   string `json:"foto"`
	Password  string `json:"senha"`
	ID        int64  `json:"iduser"`
	Level     int    `json:"tipo_usuario"`
	Instution int    `json:"idinstituicao"`
	Specialty int
}

type MysqlUser struct {
	User
	ID        string `json:"iduser"`
	Level     string `json:"tipo_usuario"`
	Instution string `json:"idinstituicao"`
}

func (mysql MysqlUser) ParseSql() (u User) {
	intId, _ := strconv.Atoi(mysql.ID)
	intLevel, _ := strconv.Atoi(mysql.Level)
	intInst, _ := strconv.Atoi(mysql.Instution)
	u = mysql.User
	u.ID = int64(intId)
	u.Level = intLevel
	u.Instution = intInst
	return
}
