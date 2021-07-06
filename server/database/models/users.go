package mysql

import (
	"database/sql"
)

type UserModel struct {
	Connection *sql.DB
}

func (model *UserModel) Registration(name string, password string) (error) {
	sql := "INSERT INTO users (user_name, password) VALUES(?, ?)"
	_, err := model.Connection.Exec(sql, name, password)
	return err
}