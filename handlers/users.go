package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cjgk/quiz/storage"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type UserController struct {
	AppController
	Services *storage.Services
	Session  *sessions.CookieStore
}

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) error {
	var users []storage.User

	err := c.Services.User.RetrieveSet(&users)
	if err != nil {
		return err
	}

	jsonUsers, err := json.Marshal(users)
	if err != nil {
		return err
	}

	fmt.Fprint(w, string(jsonUsers))

	return nil
}

func (c *UserController) Get(w http.ResponseWriter, r *http.Request) error {
	user, err := c.getRequestedUser(r)
	if err != nil {
		return err
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		return err
	}

	fmt.Fprint(w, string(jsonUser))

	return nil
}

func (c *UserController) Post(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := storage.NewUser(name, email, password)
	if err != nil {
		return Err400
	}

	if err := c.Services.User.Save(&user); err != nil {
		if err == storage.ErrEmailExists {
			return Err409
		}

		return Err500
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		return Err500
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(jsonUser))

	return nil
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) error {
	user, err := c.getRequestedUser(r)
	if err != nil {
		return err
	}

	err = c.Services.User.Delete(&user)
	if err != nil {
		return Err500
	}

	return nil
}

func (c *UserController) Put(w http.ResponseWriter, r *http.Request) error {
	user, err := c.getRequestedUser(r)
	if err != nil {
		return err
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if len(name) > 0 {
		user.Name = name
	}

	if len(email) > 0 {
		user.Email = email
	}

	if len(password) > 0 {
		pwHash, err := storage.HashPw(password)
		if err != nil {
			return err
		}

		user.Password = pwHash
	}

	err = c.Services.User.Save(&user)
	if err != nil {
		return Err500
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		return Err500
	}

	fmt.Fprint(w, string(jsonUser))

	return nil
}

func (c *UserController) getRequestedUser(r *http.Request) (storage.User, error) {
	vars := mux.Vars(r)
	user := storage.User{}

	userId, err := strconv.Atoi(vars["key"])
	if err != nil {
		return user, Err400
	}

	err = c.Services.User.Retrieve(&user, userId)
	if err == storage.ErrNotFound {
		return user, Err404
	} else if err != nil {
		return user, Err500
	}

	return user, nil
}
