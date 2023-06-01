package model

import "shylinux.com/x/golang-story/src/project/server/domain/enums"

type User struct {
	Common
	Name  string
	Email string
	Phone string
}

func (s User) TableName() string { return enums.Table.User }
