package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/cjgk/quiz/storage"

	"github.com/gorilla/sessions"
)

type HomeHandler struct {
	AppController
	Services *storage.Services
	Session  *sessions.CookieStore
}

func (h *HomeHandler) Index(w http.ResponseWriter, r *http.Request) error {

	indexFile, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		log.Fatal(err)
		return Err500
	}

	fmt.Fprint(w, string(indexFile))

	return nil
}
