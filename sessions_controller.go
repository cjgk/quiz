package main

import (
	"github.com/gorilla/sessions"
	"net/http"
)

type sessionsController struct {
	appController
	services *services
	session  *sessions.CookieStore
}

func (c *sessionsController) post(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := User{}
	err := c.services.user.RetrieveByEmail(&user, email)
	if err != nil {
		return err
	}

	if err = validatePw(password, user.Password); err != nil {
		return Err401
	}

	session, err := c.session.Get(r, "login")
	if err != nil {
		return err
	}
	session.Values["id"] = user.Id
	session.Save(r, w)

	w.WriteHeader(http.StatusCreated)

	return nil
}

func (c *sessionsController) delete(w http.ResponseWriter, r *http.Request) error {
	session, err := c.session.Get(r, "login")
	if err != nil {
		return err
	}

	// Make session too old
	session.Options.MaxAge = -1
	session.Save(r, w)

	return nil
}

func (c *sessionsController) notImp(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNotImplemented)
	return nil
}
