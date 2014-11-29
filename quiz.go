package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Set up Gorp
	dbmap := initDb()
	defer dbmap.Db.Close()

	// Set up static file server
	fileServer := http.FileServer(http.Dir("static"))

	// Set up Table services
	tableServices := initTableServices(dbmap)

	// Create controllers and add ervices to them
	users := &userController{services: &tableServices}

	// Set up router
	router := mux.NewRouter()
	router.StrictSlash(false)

	// Static route
	router.Handle("/", fileServer)

	// User routes
	router.Handle("/users", users.action(users.index)).Methods("GET")
	router.Handle("/users/{key}", users.action(users.get)).Methods("GET")
	router.Handle("/users", users.action(users.post)).Methods("POST")
	router.Handle("/users/{key}", users.action(users.put)).Methods("PUT")
	router.Handle("/users/{key}", users.action(users.delete)).Methods("DELETE")

	http.Handle("/", router)
	http.ListenAndServe(":3001", nil)
}
