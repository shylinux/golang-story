package model

type User struct {
	Common
	Name  string
	Email string
	Phone string
}

func (s User) Table() string { return "user" }
