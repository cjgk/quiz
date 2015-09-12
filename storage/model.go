package storage

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"time"

	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
)

type Services struct {
	User UserServicer
	Game GameServicer
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

var (
	ErrNotFound = errors.New("Could not find entity")
)

// Initialize gorp for struct mapping
func InitDb() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "/tmp/quiz.sqlite")
	checkErr(err, "DB INIT")

	err = db.Ping()
	checkErr(err, "DB PING")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	usersMap := dbmap.AddTableWithName(User{}, "users")
	usersMap.SetKeys(true, "Id")
	usersMap.ColMap("email").Unique = true
	dbmap.AddTableWithName(Game{}, "games").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

// Initalize services
func InitTableServices(dbmap *gorp.DbMap) Services {
	services := Services{
		User: NewUserService(dbmap),
		Game: NewGameService(dbmap),
	}

	return services
}

// Database helpers
func dateTimeToDbDateTime(dateTime time.Time) string {
	return dateTime.UTC().Format("2006-01-02 15:04:05")
}
