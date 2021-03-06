package storage

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"strings"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/coopernurse/gorp"
)

var ErrEmailExists = errors.New("users: email exists")

type User struct {
	Id       int    `db:"id"       json:"id"`
	Deleted  bool   `db:"deleted"  json:"-"`
	Email    string `db:"email"    json:"email"`
	Name     string `db:"name"     json:"name"`
	Password string `db:"password" json:"-"`
}

func NewUser(name, email, password string) (User, error) {
	user := User{}

	pwHash, err := HashPw(password)
	if err != nil {
		return user, err
	}

	user = User{
		Deleted:  false,
		Email:    email,
		Name:     name,
		Password: pwHash,
	}

	return user, nil
}

type UserServicer interface {
	Retrieve(user *User, id int) error
	RetrieveSet(users *[]User) error
	Save(user *User) error
	Delete(user *User) error
	RetrieveByEmail(user *User, email string) error
}

type UserService struct {
	Db *gorp.DbMap
}

func NewUserService(dbmap *gorp.DbMap) UserServicer {
	var environment string = os.Getenv("GOENV")

	if environment == "TEST" {
		return MockUserService{}
	}

	return UserService{Db: dbmap}
}

func (us UserService) Retrieve(user *User, id int) error {
	query := "select * from users where deleted = 0 and id = ?"
	err := us.Db.SelectOne(&user, query, id)
	if err == sql.ErrNoRows {
		return ErrNotFound
	} else if err != nil {
		return err
	}

	return nil
}

func (us UserService) RetrieveSet(users *[]User) error {
	query := "select * from users where deleted = 0"
	_, err := us.Db.Select(users, query)
	if err != nil {
		return err
	}

	return nil
}

func (us UserService) RetrieveByEmail(user *User, email string) error {
	query := "select * from users where deleted = 0 and email = ?"
	err := us.Db.SelectOne(&user, query, email)
	if err == sql.ErrNoRows {
		return ErrNotFound
	} else if err != nil {
		return err
	}

	return nil
}

func (us UserService) Save(user *User) error {
	var err error

	if user.Id == 0 {
		err = us.Db.Insert(user)
	} else {
		_, err = us.Db.Update(user)
	}

	if err != nil {
		if strings.Index(err.Error(), "UNIQUE") == 0 {
			err = ErrEmailExists
		}
		log.Print(err)
		return err
	}

	return nil
}

func (us UserService) Delete(user *User) error {
	user.Deleted = true
	if _, err := us.Db.Update(user); err != nil {
		return err
	}

	return nil
}

func HashPw(pass string) (string, error) {
	bytePass := []byte(pass)
	pwHash, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	strHashPass := string(pwHash)

	return strHashPass, nil
}

func ValidatePw(pass string, hash string) error {
	bytePass := []byte(pass)
	byteHash := []byte(hash)

	return bcrypt.CompareHashAndPassword(byteHash, bytePass)
}
