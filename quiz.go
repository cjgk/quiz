package main

import (
	"github.com/cjgk/quiz/controllers"
	"github.com/cjgk/quiz/models"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Set up Gorp
	dbmap := models.InitDb()
	defer dbmap.Db.Close()

	// Set up Table services
	tableServices := models.InitTableServices(dbmap)

	// Create controllers and add Services to them
	users := &controllers.UserController{Services: &tableServices}

	// Set up router
	router := mux.NewRouter()
	router.StrictSlash(false)

    // User routes
	router.HandleFunc("/users", users.Action(users.Index)).Methods("GET")
	router.HandleFunc("/users/{key}", users.Action(users.Get)).Methods("GET")
	router.HandleFunc("/users", users.Action(users.Post)).Methods("POST")
	router.HandleFunc("/users/{key}", users.Action(users.Put)).Methods("PUT")
	router.HandleFunc("/users/{key}", users.Action(users.Delete)).Methods("DELETE")

	http.Handle("/", router)

	http.ListenAndServe(":3001", nil)
}
