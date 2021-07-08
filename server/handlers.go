package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"github.com/julienschmidt/httprouter"
)
func (app *Application) authentical(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Auth")
}

func (app *Application) registerUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
func (app *Application) authPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//id := r.URL.Query().Get("id")
	files := []string{"./front/html/auth.html"}
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
func (app *Application) homePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
}
func (app *Application) messenger(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	files := []string{"./front/html/index.html"}
	templ, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("ERROR")
		http.Error(w, "Error", 500)
	}
	err = templ.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error", 500)
	}
}
func (app *Application) channelShow(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	app.messagesShow(w,r)
}
func (app *Application) messagesShow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Show")
}
func (app *Application) messageCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	chat := r.URL.Query().Get("chat")
	fmt.Fprintf(w, chat)
}
