package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/cjgk/quiz/models"
	"github.com/coopernurse/gorp"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *mux.Router

func setup() {
	var dbmap *gorp.DbMap
	tableServices := models.InitTableServices(dbmap)

	users := &UserController{Services: &tableServices}

	router = mux.NewRouter()
	router.Handle("/users", users.Action(users.Index)).Methods("GET")
	router.Handle("/users/{key}", users.Action(users.Get)).Methods("GET")
	router.Handle("/users", users.Action(users.Post)).Methods("POST")
	router.Handle("/users/{key}", users.Action(users.Put)).Methods("PUT")
	router.Handle("/users/{key}", users.Action(users.Delete)).Methods("DELETE")

}

func TestUserGet(t *testing.T) {
	setup()

	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		fmt.Println(err)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("/users/5 did not return %v", http.StatusOK)
	}

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Could not read body")
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		t.Errorf("Could not unmarshal response to /users/5")
	}

	expectedId := 1
	if user.Id != expectedId {
		t.Errorf("user.Id is not %v", expectedId)
	}

	expectedEmail := "user1@example.com"
	if user.Email != expectedEmail {
		t.Error("user.email is not %s", expectedEmail)
	}

	expectedName := "User Numberone"
	if user.Name != expectedName {
		t.Error("user.name is not %s", expectedName)
	}
}

func TestUsersGet(t *testing.T) {
    setup()

    var users []models.User

    req, _ := http.NewRequest("GET", "/users", nil)
    rec := httptest.NewRecorder()
    router.ServeHTTP(rec, req)

    if rec.Code != http.StatusOK {
        t.Errorf("/users did not return %v", http.StatusOK)
    }

    body, _ := ioutil.ReadAll(rec.Body)
    _ = json.Unmarshal(body, &users)

    expectedLength := 2
    if len(users) != expectedLength {
        t.Errorf("length of response is not %v", expectedLength)
    }
}
