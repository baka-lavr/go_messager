package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", homePage)
	server.HandleFunc("/channel", channelShow)
	log.Println("Запуск сервера...")
	err := http.ListenAndServe(":4000", server)
	log.Fatal(err)
}