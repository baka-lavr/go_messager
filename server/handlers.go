package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	files := []string{"./front/html/home.page.tmpl","./front/html/base.layout.tmpl"}
	templ, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("ERROR")
		//http.Error(w, "Error", 500)
	}
	err = templ.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error", 500)
	}
	fmt.Fprintf(w, "Канал %d", id)
}
func channelShow(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Fprintf(w, "Канал %d", id)
}