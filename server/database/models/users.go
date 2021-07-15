package mysql

import (
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id int
	name string
}

type UserModel struct {
	Connection *sql.DB
}

func (model *UserModel) Verify(name string, password string) (error) {
	sql := "SELECT user_id, user_name, password FROM users WHERE user_id = ?"
	row := model.Connection.QueryRow(sql, name)
	user := &User{}
	pass := ""
	err := row.Scan(&user.id, &user.name, pass)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(pass))
	return err
}

func (model *UserModel) GetFromChat(chat string) ([]byte) {
	var users []User
	users_row, _ := model.Connection.Query("select user_id, user_name from user-chat inner join users on user-chat.user_id=users.user_id where chat_id=?", chat)
	defer users_row.Close()
	for users_row.Next() {
		var value User
		users_row.Scan(&value.id, &value.name)
		users = append(users, value)
	}
	js, _ := json.Marshal(users)
	return js
}

func (model *UserModel) Registration(name string, password string) (error) {
	sql := "INSERT INTO users (user_name, password) VALUES(?, ?)"
	_, err := model.Connection.Exec(sql, name, password)
	return err
}