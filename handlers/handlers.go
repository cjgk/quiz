package handlers

import (
	"errors"
	"net/http"

	"github.com/gorilla/sessions"
)

type Action func(rw http.ResponseWriter, r *http.Request) error

type AppController struct{}

func (c *AppController) Action(a Action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		c.RunAction(a, rw, r)
	})
}

func (c *AppController) AuthAction(a Action, sess *sessions.CookieStore) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		session, _ := sess.Get(r, "login")
		userId := session.Values["id"]

		if userId == nil {
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
			return
		}

		c.RunAction(a, rw, r)
	})
}

func (c *AppController) RunAction(a Action, rw http.ResponseWriter, r *http.Request) {
	if err := a(rw, r); err != nil {
		switch err {
		case Err400:
			http.Error(rw, err.Error(), http.StatusBadRequest)
		case Err401:
			http.Error(rw, err.Error(), http.StatusUnauthorized)
		case Err404:
			http.Error(rw, err.Error(), http.StatusNotFound)
		case Err409:
			http.Error(rw, err.Error(), http.StatusConflict)
		default:
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}
}

var (
	Err400 = errors.New("Bad input")
	Err401 = errors.New("Unauthorized")
	Err404 = errors.New("Not found")
	Err409 = errors.New("Conflict")
	Err500 = errors.New("Internal Server Error")
)
