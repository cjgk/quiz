package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Set up Gorp
	dbmap := InitDb()
	defer dbmap.Db.Close()

	// Set up Table services
	tableServices := InitTableServices(dbmap)

	// Create controllers and add Services to them
	users := &UserController{Services: &tableServices}

	// Set up router
	router := mux.NewRouter()
	router.StrictSlash(false)

	// User routes
	router.Handle("/users", users.Action(users.Index)).Methods("GET")
	router.Handle("/users/{key}", users.Action(users.Get)).Methods("GET")
	router.Handle("/users", users.Action(users.Post)).Methods("POST")
	router.Handle("/users/{key}", users.Action(users.Put)).Methods("PUT")
	router.Handle("/users/{key}", users.Action(users.Delete)).Methods("DELETE")

	http.Handle("/", router)

	http.ListenAndServe(":3001", nil)
}
