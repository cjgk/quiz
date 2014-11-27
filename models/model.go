package models

import (
	"database/sql"
	"errors"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type Services struct {
	User userServicer
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

	dbmap.AddTableWithName(UserFields{}, "users").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

// Initalize Services
func InitTableServices(dbmap *gorp.DbMap) Services {
	services := Services{
		User: NewUserService(dbmap),
	}

	return services
}

