package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
	"strconv"
)

type userController struct {
	appController
	services *services
    session *sessions.CookieStore
}

func (c *userController) index(w http.ResponseWriter, r *http.Request) error {
	var users []User

	err := c.services.user.RetrieveSet(&users)
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

func (c *userController) get(w http.ResponseWriter, r *http.Request) error {
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

func (c *userController) post(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := newUser(name, email, password)
	if err != nil {
		return Err400
	}

	if err := c.services.user.Save(&user); err != nil {
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

func (c *userController) delete(w http.ResponseWriter, r *http.Request) error {
	user, err := c.getRequestedUser(r)
	if err != nil {
		return err
	}

	err = c.services.user.Delete(&user)
	if err != nil {
		return Err500
	}

	return nil
}

func (c *userController) put(w http.ResponseWriter, r *http.Request) error {
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
		pwHash, err := HashPw(password)
		if err != nil {
			return err
		}

		user.Password = pwHash
	}

	err = c.services.user.Save(&user)
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

func (c *userController) getRequestedUser(r *http.Request) (User, error) {
	vars := mux.Vars(r)
	user := User{}

	userId, err := strconv.Atoi(vars["key"])
	if err != nil {
		return user, Err400
	}

	err = c.services.user.Retrieve(&user, userId)
	if err == ErrNotFound {
		return user, Err404
	} else if err != nil {
		return user, Err500
	}

	return user, nil
}
