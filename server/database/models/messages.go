package mysql

import (
	"database/sql"
	"encoding/json"
)

type Message struct {
	text string `json: "text"`
	user string `json: "user"`
	media string `json: "media"`
	//date string `json: "date"`
}

type MessageModel struct {
	Connection *sql.DB
}

func (model *MessageModel) Create(chat string, user string, text string, media int) error {
	tr, err := model.Connection.Begin()
	if err != nil {
		return err
	}
	content, err := tr.Exec("insert into contents(mess_contents_content) values(?)", text)
	if err != nil {
		tr.Rollback()
		return err
	}
	content, err = tr.Exec("insert into messages(user_id, mess_contents_id, mess_created, media_file_id) values(?, ?, NOW(), ?)", user, content, media)
	if err != nil {
		tr.Rollback()
		return err
	}
	_, err = tr.Exec("insert into message-chat(mess_id,chat_id) values(?, ?)", content, chat)
	if err != nil {
		tr.Rollback()
		return err
	}
	err = tr.Commit()
	return err
}

func (model *MessageModel) GetFromChat(chat string) []byte {
	var mess []Message
	row, _ := model.Connection.Query("select mess_contents_content, user_id, media_file_link from `message-chat` inner join messages on `message-chat`.mess_id=messages.mess_id inner join contents on messages.mess_id=contents.mess_id inner join media on messages.media_file_id=media.media_file_id where chat_id=?", chat)
	defer row.Close()
	for row.Next() {
		var value Message
		row.Scan(&value.text, &value.user, &value.media)
		mess = append(mess, value)
	}
	js, _ := json.Marshal(mess)
	return js
}