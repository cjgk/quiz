package main

type mockUserService struct{}

func (mus mockUserService) Retrieve(user *User, id int) error {
	if id != 1 {
		return ErrNotFound
	}

	user.Id = 1
	user.Deleted = false
	user.Email = "user1@example.com"
	user.Name = "User Numberone"

	return nil
}

func (mus mockUserService) RetrieveSet(users *[]User) error {
	user1 := User{
		Id:      1,
		Deleted: false,
		Email:   "user1@example.com",
		Name:    "User Numberone",
	}

	user2 := User{
		Id:      2,
		Deleted: false,
		Email:   "user2@example.com",
		Name:    "User Numbertwo",
	}

	*users = append(*users, user1, user2)

	return nil
}

func (mus mockUserService) RetrieveByEmail(user *User, email string) error {
	return nil
}

func (mus mockUserService) Save(user *User) error {
	if user.Id == 0 {
		user.Id = 3
	}

	return nil
}

func (mus mockUserService) Delete(user *User) error {
	user.Deleted = true
	return nil
}
