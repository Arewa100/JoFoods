package models

type User struct {
	Username string
}

func (user *User) getName() string {
	return user.Username
}
