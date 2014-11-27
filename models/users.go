package models

import (
	"code.google.com/p/go.crypto/bcrypt"
	"database/sql"
	"github.com/coopernurse/gorp"
    "os"
)

type UserFields struct {
	Id       int    `json:"id"`
	Deleted  bool   `json:"-"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

type User struct {
	UserFields
}

func NewUser(name, email, password string) (User, error) {
	pwHash, err := HashPw(password)
	user := User{}
	if err != nil {
		return User{}, err
	}

	user = User{
		UserFields{
			Deleted:  false,
			Email:    email,
			Name:     name,
			Password: pwHash,
		},
	}

	return user, nil
}

func (u *User) Populate(uf UserFields) error {
	u.Id = uf.Id
	u.Deleted = uf.Deleted
	u.Email = uf.Email
	u.Name = uf.Name
	u.Password = uf.Password
	return nil
}

func (u *User) Extract() UserFields {
	return UserFields{
		Id:       u.Id,
		Deleted:  u.Deleted,
		Email:    u.Email,
		Name:     u.Name,
		Password: u.Password,
	}
}

type userServicer interface {
	Retrieve(user *User, id int) error
	RetrieveSet(users *[]User) error
	Save(user *User) error
	Delete(user *User) error
}

type userService struct {
	Db *gorp.DbMap
}

func NewUserService(dbmap *gorp.DbMap) userServicer {
    var environment string = os.Getenv("GOENV")

    if environment == "TEST" {
        return  mockUserService{}
    }

	return userService{Db: dbmap}
}

func (us userService) Retrieve(user *User, id int) error {
	query := "select * from users where deleted = 0 and id = ?"
	userFields := UserFields{}
	err := us.Db.SelectOne(&userFields, query, id)
	if err == sql.ErrNoRows {
		return ErrNotFound
	} else if err != nil {
		return err
	}

	user.Populate(userFields)

	return nil
}

func (us userService) RetrieveSet(users *[]User) error {
	query := "select * from users where deleted = 0"
	_, err := us.Db.Select(users, query)
	if err != nil {
		return err
	}

	return nil
}

func (us userService) Save(user *User) error {
	var err error

	userFields := user.Extract()
	if userFields.Id == 0 {
		err = us.Db.Insert(&userFields)
	} else {
		_, err = us.Db.Update(&userFields)
	}

	if err != nil {
		return err
	}

	user.Populate(userFields)

	return nil
}

func (us userService) Delete(user *User) error {
	userFields := user.Extract()
	userFields.Deleted = true
	if _, err := us.Db.Update(&userFields); err != nil {
		return err
	}
	user.Populate(userFields)

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

func validatePw(pass string, hash string) error {
	bytePass := []byte(pass)
	byteHash := []byte(hash)

	return bcrypt.CompareHashAndPassword(byteHash, bytePass)
}
