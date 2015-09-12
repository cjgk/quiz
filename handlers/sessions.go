package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cjgk/quiz/storage"
	"github.com/gorilla/sessions"
)

type SessionsController struct {
	AppController
	Services *storage.Services
	Session  *sessions.CookieStore
}

func (c *SessionsController) Post(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := storage.User{}
	err := c.Services.User.RetrieveByEmail(&user, email)
	if err != nil {
		return err
	}

	if err = storage.ValidatePw(password, user.Password); err != nil {
		return Err401
	}

	session, err := c.Session.Get(r, "login")
	if err != nil {
		return err
	}
	session.Values["id"] = user.Id
	session.Save(r, w)

	jsonUser, err := json.Marshal(user)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(jsonUser))

	return nil
}

func (c *SessionsController) Delete(w http.ResponseWriter, r *http.Request) error {
	session, err := c.Session.Get(r, "login")
	if err != nil {
		return err
	}

	// Make session too old
	session.Options.MaxAge = -1
	session.Save(r, w)

	return nil
}

func (c *SessionsController) NotImp(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNotImplemented)
	return nil
}
