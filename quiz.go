package main

import (
	"net/http"

	"github.com/cjgk/quiz/handlers"
	"github.com/cjgk/quiz/storage"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func main() {
	// Set up Gorp
	dbmap := storage.InitDb()
	defer dbmap.Db.Close()

	// Set up static file server
	fileServer := http.FileServer(http.Dir("./public/"))

	// Set up Sessions service
	sessionsStore := sessions.NewCookieStore([]byte("laksdjflöaskjdfölaskdjf"))

	// Set up Table services
	tableServices := storage.InitTableServices(dbmap)

	// Create handlers and add services to them
	home := &handlers.HomeHandler{Services: &tableServices, Session: sessionsStore}
	users := &handlers.UserHandler{Services: &tableServices, Session: sessionsStore}
	games := &handlers.GameHandler{Services: &tableServices, Session: sessionsStore}
	sessions := &handlers.SessionHandler{Services: &tableServices, Session: sessionsStore}

	// Set up router
	router := mux.NewRouter()
	router.StrictSlash(true)

	// Home route
	router.Handle("/", home.Action(home.Index)).Methods("GET")

	// User routes
	router.Handle("/users", users.AuthAction(users.Index, sessionsStore)).Methods("GET")
	router.Handle("/users/{key}", users.AuthAction(users.Get, sessionsStore)).Methods("GET")
	router.Handle("/users", users.Action(users.Post)).Methods("POST")
	router.Handle("/users/{key}", users.AuthAction(users.Put, sessionsStore)).Methods("PUT")
	router.Handle("/users/{key}", users.AuthAction(users.Delete, sessionsStore)).Methods("DELETE")

	// Session routes
	router.Handle("/sessions", sessions.Action(sessions.Post)).Methods("POST")
	router.Handle("/sessions", sessions.AuthAction(sessions.Delete, sessionsStore)).Methods("DELETE")
	router.Handle("/sessions.*", sessions.Action(sessions.NotImp)).Methods("GET", "PUT")

	// Game routes
	router.Handle("/games", games.AuthAction(games.Index, sessionsStore)).Methods("GET")
	router.Handle("/games/{key}", games.AuthAction(games.Get, sessionsStore)).Methods("GET")
	router.Handle("/games", games.AuthAction(games.Post, sessionsStore)).Methods("POST")
	router.Handle("/games/{key}", games.AuthAction(games.Put, sessionsStore)).Methods("PUT")
	router.Handle("/games/{key}", games.AuthAction(games.Delete, sessionsStore)).Methods("DELETE")

	// Static route
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	http.Handle("/", router)
	http.ListenAndServe(":3000", nil)
}
