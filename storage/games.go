package storage

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/coopernurse/gorp"
)

type Game struct {
	Id      int       `db:"id"      json:"id"`
	Deleted bool      `db:"deleted" json:"-"`
	Created time.Time `db:"created" json:"created"`
	UserId  int       `db:"user_id" json:"-"`
	Status  string    `db:"status"  json:"status"`
	Name    string    `db:"name"    json:"name"`
}

func NewGame(name string, userId int) Game {
	return Game{
		Deleted: false,
		Created: time.Now().UTC(),
		UserId:  userId,
		Status:  "new",
		Name:    name,
	}
}

type GameServicer interface {
	Retrieve(game *Game, id int) error
	RetrieveSet(games *[]Game) error
	Save(game *Game) error
	Delete(game *Game) error
}

type GameService struct {
	Db *gorp.DbMap
}

func NewGameService(dbmap *gorp.DbMap) GameServicer {
	var environment string = os.Getenv("GOENV")

	if environment == "TEST" {
		//			return mockGameService{}
	}

	return GameService{Db: dbmap}
}

func (us GameService) Retrieve(game *Game, id int) error {
	query := "select * from games where deleted = 0 and id = ?"
	err := us.Db.SelectOne(&game, query, id)
	if err == sql.ErrNoRows {
		return ErrNotFound
	} else if err != nil {
		return err
	}

	return nil
}

func (us GameService) RetrieveSet(games *[]Game) error {
	query := "select * from games where deleted = 0"
	_, err := us.Db.Select(games, query)
	if err != nil {
		return err
	}

	return nil
}

func (us GameService) Save(game *Game) error {
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

func (us GameService) Delete(game *Game) error {
	game.Deleted = true
	if _, err := us.Db.Update(game); err != nil {
		return err
	}

	return nil
}
