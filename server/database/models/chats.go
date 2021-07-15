package mysql

import (
	"database/sql"
)

type ChatUpdateInfo struct {
	id int `json: "id"`
	name string `json: "name"`
	Users []string
}

type ChatModel struct {
	Connection *sql.DB
}

func (model *ChatModel) Update(id string) *ChatUpdateInfo {
	var users []string
	users_row, _ := model.Connection.Query("select user_id from user-chat where chat_id=?", id)
	defer users_row.Close()
	for users_row.Next() {
		var value string
		users_row.Scan(&value)
		users = append(users, value)
	}
	chat_row := model.Connection.QueryRow("select chat_id, chat_name from chat where chat_id=?", id)
	update_info := ChatUpdateInfo{
		Users: users,
	}
	chat_row.Scan(&update_info.id, &update_info.name)
	return &update_info
}