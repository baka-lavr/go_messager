package main

import (
	"fmt"
	"log"
	"time"
	"strings"
	"net/http"
	"html/template"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	//"bsu.ru/messenger/server/database/models"
)
func (app *Application) authentical(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Auth")
	user := r.Header.Get("User")
	token := r.Header.Get("Token")
	err := app.db.User.Verify(user, token)
	if err != nil {
		http.Error(w, "Error", 500)
	}
}

func (app *Application) loginAttempt(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	id := r.Form.Get("id")
	password := r.Form.Get("password")
	token, _ := bcrypt.GenerateFromPassword([]byte(password), 6)
	http.SetCookie(w, &http.Cookie{
		Name: "user",
		Value: id,
		Expires: time.Now().Add(time.Hour),
	})
	http.SetCookie(w, &http.Cookie{
		Name: "session_token",
		Value: string(token),
		Expires: time.Now().Add(time.Hour),
	})
	fmt.Fprintf(w, "Auth")
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
func (app *Application) homePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/messenger", http.StatusSeeOther)
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

func (app *Application) noticeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming error", http.StatusBadRequest)
		return
	}
	userCache := "test"
	channel := make(chan []byte)
	app.notifier.Clients[userCache] = channel
	
	d := make(chan interface{})
	defer close(d)
	for {
		select {
		case <-d:
			close(channel)
			return
		case data := <-channel:
			app.logger.Println(data)
			fmt.Fprintf(w, string(data))
			flusher.Flush()
		}
	}
}

func (app *Application) chat(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	chat_id := r.Form.Get("chat")
	us := app.db.User.GetFromChat(chat_id)
	fmt.Fprintf(w, string(us))
}
func (app *Application) messagesShow(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	chat_id := r.Form.Get("chat")
	me := app.db.Message.GetFromChat(chat_id)
	fmt.Fprintf(w, string(me))
}
func (app *Application) messageCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseMultipartForm(32<<20)
	chat := r.Form.Get("chat")
	user := r.Form.Get("user")
	message := r.Form.Get("message")
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "Error")
		return
	}
	var media int64
	if file != nil {
		defer file.Close()
		path := strings.Split(header.Filename, ".")[0]
		media, err = app.db.Media.Upload(file, path)
		if err != nil {
			fmt.Fprintf(w, "Error")
			return
		}
	}
	err = app.db.Message.Create(chat,user,message, int(media))
	if err != nil {
		return
	}
	data := app.db.Chat.Update(chat)
	app.notifier.MainChannel <- *data
}
