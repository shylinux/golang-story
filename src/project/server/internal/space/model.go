package space

import "shylinux.com/x/golang-story/src/project/server/domain/model"

type Space struct {
	model.Common
	Name  string
	Email string
	Phone string
}

func (s Space) Table() string { return "space" }
