package main

import (
	"net/http"
	"path/filepath"
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

func (app *Application) route() *http.ServeMux {
	server := http.NewServeMux()
	server.HandleFunc("/", app.homePage)
	server.HandleFunc("/register", app.registerUser)
	server.HandleFunc("/channel", app.channelShow)
	fileServer := http.FileServer(customFileSystem{http.Dir("./front/static/")})
	server.Handle("/static/", http.StripPrefix("/static", fileServer))
	return server
}