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
	Username  string `json:"username"`
	Token     string `json:"token"`
	Picture   string `json:"foto"`
	Password  string `json:"password"`
	ID        int64  `json:"iduser"`
	Level     int    `json:"level"`
	Instution int    `json:"idcompany"`
	Specialty int
}

type MysqlUser struct {
	User
	ID        string `json:"iduser"`
	Level     string `json:"level"`
	Instution string `json:"idcompany"`
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
