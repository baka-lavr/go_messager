package mysql

import (
	"io"
	"os"
	"database/sql"
	"mime/multipart"
	//"golang.org/x/crypto/bcrypt"
)

type Media struct {
	id int
	name string
}

type MediaModel struct {
	Connection *sql.DB
}

func (model *MediaModel) Upload(file multipart.File, path string) (int64, error) {
	f, err := os.OpenFile("./front/static/media/"+path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	id, err := model.Connection.Exec("insert into media(media_type, media_file_link) values(`DEFAULT`, ?)", path)
	if err != nil {
		return 0, err
	}
	io.Copy(f, file)
	return id.LastInsertId()
}