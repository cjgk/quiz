package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

type homeController struct {
	appController
	services *services
	session  *sessions.CookieStore
}

func (c *homeController) index(w http.ResponseWriter, r *http.Request) error {

	indexFile, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		log.Fatal(err)
		return Err500
	}

	fmt.Fprint(w, string(indexFile))

	return nil
}
