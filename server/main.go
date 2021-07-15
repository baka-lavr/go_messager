package main

import (
	"os"
	"log"
	"net/http"
	"bsu.ru/messenger/server/database"
	"bsu.ru/messenger/server/database/models"
)

func main() {
	dbObj, err := db.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	sseChannel := Notifier{
		Clients: make(map[string](chan []byte)),
		MainChannel: make(chan mysql.ChatUpdateInfo),
	}
	done := make(chan interface{})
	defer close(done)
	go sseChannel.Broadcast(done)

	app := &Application {
		log.New(os.Stdout, "LOG", log.Ldate|log.Ltime),
		dbObj,
		sseChannel,
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
	notifier Notifier
}

