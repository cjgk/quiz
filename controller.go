package main

import (
	"errors"
	"github.com/gorilla/sessions"
	"net/http"
)

type action func(rw http.ResponseWriter, r *http.Request) error

type appController struct{}

func (c *appController) action(a action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		c.runAction(a, rw, r)
	})
}

func (c *appController) authAction(a action, sess *sessions.CookieStore) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		session, _ := sess.Get(r, "login")
		userId := session.Values["id"]

		if userId == nil {
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
			return
		}

		c.runAction(a, rw, r)
	})
}

func (c *appController) runAction(a action, rw http.ResponseWriter, r *http.Request) {
	if err := a(rw, r); err != nil {
		switch err {
		case Err400:
			http.Error(rw, err.Error(), http.StatusBadRequest)
		case Err401:
			http.Error(rw, err.Error(), http.StatusUnauthorized)
		case Err404:
			http.Error(rw, err.Error(), http.StatusNotFound)
		default:
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}
}

var (
	Err400 = errors.New("Bad input")
	Err401 = errors.New("Unauthorized")
	Err404 = errors.New("Not found")
	Err500 = errors.New("Internal Server Error")
)
