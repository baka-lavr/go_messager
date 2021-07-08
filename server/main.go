package main

import (
	"os"
	"log"
	"net/http"
	"bsu.ru/messenger/server/database"
)

func main() {
	dbObj, err := db.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	app := &Application {
		log.New(os.Stdout, "LOG", log.Ldate|log.Ltime),
		dbObj,
	}
	router := app.NewAuth(app.route())
	server := &http.Server {
		Addr: ":4000",
		ErrorLog: app.logger,
		Handler: router,
	}
	log.Println("Запуск сервера...")
	err = server.ListenAndServe()
	log.Fatal(err)
}

type Application struct{
	logger *log.Logger
	db *db.DataBase
}

