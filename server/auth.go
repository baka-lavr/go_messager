package main

import (
	"net/http"
	"fmt"
	"errors"
	"time"
	//"golang.org/x/crypto/bcrypt"
	"github.com/patrickmn/go-cache"
)

type Auth struct {
	handler http.Handler
	cache *cache.Cache
}

func (app *Application) NewAuth(handler http.Handler) *Auth {
	auth := Auth {
		handler: handler,
		cache: cache.New(time.Hour, 10*time.Minute),
	}
	return &auth
}

func (auth *Auth) checkUser(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return "", err
	}
	token := cookie.Value
	value, found := auth.cache.Get(token)
	if found {
		return value.(string), nil
	} else {
		return "", errors.New("ErrNotAuth")
	}
}

func (auth *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("protected") != "" {
		token, err := auth.checkUser(r)
		if err == http.ErrNoCookie {
			http.Redirect(w,r,"/login",http.StatusSeeOther)
		} else if err != nil {
			client := &http.Client{}
			req, _ := http.NewRequest(http.MethodPost, "/auth", nil)
			req.Header.Add("Authorization", fmt.Sprintf("auth_token=%s", token))
			_, err := client.Do(req)
			if err != nil {
				http.SetCookie(w, &http.Cookie{
					Name: "session_token",
					Value: "",
					Expires: time.Now(),
				})
				http.Redirect(w,r,"/login",http.StatusSeeOther)
			}
		}
	}
	auth.handler.ServeHTTP(w,r)
	fmt.Fprintf(w, "Show")
}