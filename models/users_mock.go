package models

type mockUserService struct {}

func (mus mockUserService) Retrieve(user *User, id int) error {
    user.Id = 1
    user.Deleted = false
    user.Email = "user1@example.com"
    user.Name = "User Numberone"

    return nil
}

func (mus mockUserService) RetrieveSet (users *[]User) error {
    return nil
}

func (mus mockUserService) Save (user *User) error {
    return nil
}

func (mus mockUserService) Delete (user *User) error {
    return nil
}
