package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cjgk/quiz/storage"
	"github.com/gorilla/sessions"
)

type SessionHandler struct {
	AppController
	Services *storage.Services
	Session  *sessions.CookieStore
}

func (h *SessionHandler) Post(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := storage.User{}
	err := h.Services.User.RetrieveByEmail(&user, email)
	if err != nil {
		return err
	}

	if err = storage.ValidatePw(password, user.Password); err != nil {
		return Err401
	}

	session, err := h.Session.Get(r, "login")
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

func (h *SessionHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	session, err := h.Session.Get(r, "login")
	if err != nil {
		return err
	}

	// Make session too old
	session.Options.MaxAge = -1
	session.Save(r, w)

	return nil
}

func (h *SessionHandler) NotImp(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNotImplemented)
	return nil
}
