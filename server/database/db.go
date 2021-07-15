package db

import (
	"database/sql"
	"bsu.ru/messenger/server/database/models"
	_"github.com/go-sql-driver/mysql"
)

type DataBase struct {
	User *mysql.UserModel
	Message *mysql.MessageModel
	Chat *mysql.ChatModel
	Media *mysql.MediaModel
}

func OpenDB() (*DataBase, error) {
	db, err := sql.Open("mysql", "admin: supersecretpassword@/messenger")
	if err != nil {
		return nil, err
	}
	obj := &DataBase{
		&mysql.UserModel{
			Connection: db,
		},
		&mysql.MessageModel{
			Connection: db,
		},
		&mysql.ChatModel{
			Connection: db,
		},
		&mysql.MediaModel{
			Connection: db,
		},
	}
	return obj, nil
}