package main

import (
	"net/http"
	"path/filepath"
	"github.com/julienschmidt/httprouter"
)

type customFileSystem struct {
	fs http.FileSystem
}

func (cfs customFileSystem) Open(path string) (http.File, error) {
	f, err := cfs.fs.Open(path)
	if err != nil {
		return nil, err
	}
	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := cfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}
	return f, nil
}

func (app *Application) route() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", app.homePage)
	router.GET("/login", app.authPage)
	router.POST("/loginAttempt", app.loginAttempt)
	router.GET("/messenger", app.messenger)
	router.GET("/sse", app.noticeHandler)
	router.GET("/messenger/c/:channel", app.chat)
	router.POST("/send", app.messageCreate)
	router.POST("/showMessages", app.messagesShow)
	router.POST("/auth", app.authentical)
	router.POST("/register", app.registerUser)
	fileServer := customFileSystem{http.Dir("./front/static/")}
	router.ServeFiles("/static/*filepath", fileServer)
	return router
}