package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/cjgk/quiz/storage"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type GameHandler struct {
	AppController
	Services *storage.Services
	Session  *sessions.CookieStore
}

func (h *GameHandler) Index(w http.ResponseWriter, r *http.Request) error {
	var games []storage.Game
	var userId int

	sess, err := h.Session.Get(r, "login")
	log.Print(sess)

	if id, ok := sess.Values["id"]; ok {
		userId = id.(int)
	} else {
		userId = 0
	}

	err = h.Services.Game.RetrieveSet(&games, userId)
	if err != nil {
		return err
	}

	jsonGames, err := json.Marshal(games)
	if err != nil {
		return err
	}

	fmt.Fprint(w, string(jsonGames))

	return nil
}

func (h *GameHandler) Get(w http.ResponseWriter, r *http.Request) error {
	game, err := h.getRequestedGame(r)
	if err != nil {
		return err
	}

	jsonGame, err := json.Marshal(game)
	if err != nil {
		return err
	}

	fmt.Fprint(w, string(jsonGame))

	return nil
}

func (h *GameHandler) Post(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	userId := 1

	game := storage.NewGame(name, userId)

	if err := h.Services.Game.Save(&game); err != nil {
		return Err500
	}

	jsonGame, err := json.Marshal(game)
	if err != nil {
		return Err500
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(jsonGame))

	return nil
}

func (h *GameHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	game, err := h.getRequestedGame(r)
	if err != nil {
		return err
	}

	err = h.Services.Game.Delete(&game)
	if err != nil {
		return Err500
	}

	return nil
}

func (h *GameHandler) Put(w http.ResponseWriter, r *http.Request) error {
	game, err := h.getRequestedGame(r)
	if err != nil {
		return err
	}

	name := r.FormValue("name")

	if len(name) > 0 {
		game.Name = name
	}

	err = h.Services.Game.Save(&game)
	if err != nil {
		return Err500
	}

	jsonGame, err := json.Marshal(game)
	if err != nil {
		return Err500
	}

	fmt.Fprint(w, string(jsonGame))

	return nil
}

func (h *GameHandler) getRequestedGame(r *http.Request) (storage.Game, error) {
	vars := mux.Vars(r)
	game := storage.Game{}

	gameId, err := strconv.Atoi(vars["key"])
	if err != nil {
		return game, Err400
	}

	err = h.Services.Game.Retrieve(&game, gameId)
	if err == storage.ErrNotFound {
		return game, Err404
	} else if err != nil {
		return game, Err500
	}

	return game, nil
}
