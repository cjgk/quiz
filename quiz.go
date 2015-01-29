package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
)

func main() {
	// Set up Gorp
	dbmap := initDb()
	defer dbmap.Db.Close()

	// Set up static file server
	fileServer := http.FileServer(http.Dir("./public/"))

	// Set up Sessions service
	sessionsStore := sessions.NewCookieStore([]byte("laksdjflöaskjdfölaskdjf"))

	// Set up Table services
	tableServices := initTableServices(dbmap)

	// Create controllers and add ervices to them
	users := &userController{services: &tableServices, session: sessionsStore}
	games := &gameController{services: &tableServices, session: sessionsStore}
	sessions := &sessionsController{services: &tableServices, session: sessionsStore}

	// Set up router
	router := mux.NewRouter()
	router.StrictSlash(true)

	// User routes
	router.Handle("/users", users.authAction(users.index, sessionsStore)).Methods("GET")
	router.Handle("/users/{key}", users.authAction(users.get, sessionsStore)).Methods("GET")
	router.Handle("/users", users.authAction(users.post, sessionsStore)).Methods("POST")
	router.Handle("/users/{key}", users.authAction(users.put, sessionsStore)).Methods("PUT")
	router.Handle("/users/{key}", users.authAction(users.delete, sessionsStore)).Methods("DELETE")

	// Session routes
	router.Handle("/sessions", sessions.action(sessions.post)).Methods("POST")
	router.Handle("/sessions/{key}", sessions.action(sessions.delete)).Methods("DELETE")
	router.Handle("/sessions.*", sessions.action(sessions.notImp)).Methods("GET", "PUT")

	// Game routes
	router.Handle("/games", games.action(games.index)).Methods("GET")
	router.Handle("/games/{key}", games.action(games.get)).Methods("GET")
	router.Handle("/games", games.action(games.post)).Methods("POST")
	router.Handle("/games/{key}", games.action(games.put)).Methods("PUT")
	router.Handle("/games/{key}", games.action(games.delete)).Methods("DELETE")

	// Static route
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	http.Handle("/", router)
	http.ListenAndServe(":3001", nil)
}
