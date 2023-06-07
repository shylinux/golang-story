package space

import (
	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/domain/model"
)

type Space struct {
	model.Common
	Name  string
	Email string
	Phone string
}

func (s Space) TableName() string { return enums.Service.Space }