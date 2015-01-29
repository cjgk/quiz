package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/gorilla/sessions"
)

type gameController struct {
	appController
	services *services
    session *sessions.CookieStore

}

func (c *gameController) index(w http.ResponseWriter, r *http.Request) error {
	var games []Game

	err := c.services.game.RetrieveSet(&games)
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

func (c *gameController) get(w http.ResponseWriter, r *http.Request) error {
	game, err := c.getRequestedGame(r)
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

func (c *gameController) post(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	userId := 1

	game := newGame(name, userId)

	if err := c.services.game.Save(&game); err != nil {
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

func (c *gameController) delete(w http.ResponseWriter, r *http.Request) error {
	game, err := c.getRequestedGame(r)
	if err != nil {
		return err
	}

	err = c.services.game.Delete(&game)
	if err != nil {
		return Err500
	}

	return nil
}

func (c *gameController) put(w http.ResponseWriter, r *http.Request) error {
	game, err := c.getRequestedGame(r)
	if err != nil {
		return err
	}

	name := r.FormValue("name")

	if len(name) > 0 {
		game.Name = name
	}

	err = c.services.game.Save(&game)
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

func (c *gameController) getRequestedGame(r *http.Request) (Game, error) {
	vars := mux.Vars(r)
	game := Game{}

	gameId, err := strconv.Atoi(vars["key"])
	if err != nil {
		return game, Err400
	}

	err = c.services.game.Retrieve(&game, gameId)
	if err == ErrNotFound {
		return game, Err404
	} else if err != nil {
		return game, Err500
	}

	return game, nil
}
