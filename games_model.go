package main

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	"os"
	"time"
)

type Game struct {
	Id      int       `db:"id"      json:"id"`
	Deleted bool      `db:"deleted" json:"-"`
	Created time.Time `db:"created" json:"created"`
	UserId  int       `db:"user_id" json:"-"`
	Status  string    `db:"status"  json:"status"`
	Name    string    `db:"name"    json:"name"`
}

func newGame(name string, userId int) Game {
	return Game{
		Deleted: false,
		Created: time.Now().UTC(),
		UserId:  userId,
		Status:  "new",
		Name:    name,
	}
}

type gameServicer interface {
	Retrieve(game *Game, id int) error
	RetrieveSet(games *[]Game) error
	Save(game *Game) error
	Delete(game *Game) error
}

type gameService struct {
	Db *gorp.DbMap
}

func newGameService(dbmap *gorp.DbMap) gameServicer {
	var environment string = os.Getenv("GOENV")

	if environment == "TEST" {
		//			return mockGameService{}
	}

	return gameService{Db: dbmap}
}

func (us gameService) Retrieve(game *Game, id int) error {
	query := "select * from games where deleted = 0 and id = ?"
	err := us.Db.SelectOne(&game, query, id)
	if err == sql.ErrNoRows {
		return ErrNotFound
	} else if err != nil {
		return err
	}

	return nil
}

func (us gameService) RetrieveSet(games *[]Game) error {
	query := "select * from games where deleted = 0"
	_, err := us.Db.Select(games, query)
	if err != nil {
		return err
	}

	return nil
}

func (us gameService) Save(game *Game) error {
	var err error

	if game.Id == 0 {
		err = us.Db.Insert(game)
	} else {
		_, err = us.Db.Update(game)
	}

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (us gameService) Delete(game *Game) error {
	game.Deleted = true
	if _, err := us.Db.Update(game); err != nil {
		return err
	}

	return nil
}