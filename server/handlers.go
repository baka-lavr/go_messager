package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

func (app *Application) registerUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		app.logger.Println("GET not allowed")
		return
	}
	r.ParseForm()
	name := r.Form.Get("username")
	password := r.Form.Get("password")
	err := app.db.User.Registration(name, password)
	if err != nil {
		app.logger.Println("Registration is failed")
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func (app *Application) homePage(w http.ResponseWriter, r *http.Request) {
	//id := r.URL.Query().Get("id")
	files := []string{"./front/html/home.page.tmpl","./front/html/base.layout.tmpl"}
	templ, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("ERROR")
		http.Error(w, "Error", 500)
	}
	err = templ.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error", 500)
	}
	//fmt.Fprintf(w, "Канал %d", id)
}
func (app *Application) channelShow(w http.ResponseWriter, r *http.Request) {
	app.messagesShow(w,r)
}
func (app *Application) messagesShow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Show")
}
func (app *Application) messageCreate(w http.ResponseWriter, r *http.Request) {
	chat := r.URL.Query().Get("chat")
	fmt.Fprintf(w, chat)
}
