package model

import (
	"fmt"

	"shylinux.com/x/golang-story/src/project/server/domain/enums"
)

type User struct {
	Common
	UserID   int64  `gorm:"uniqueIndex:idx_userid"`
	Username string `gorm:"type:varchar(32);uniqueIndex:idx_username"`
	Password string
	Email    string
	Phone    string
}

func (s User) TableName() string { return enums.Table.User }
func (s User) GetKey() string    { return "user_id" }
func (s User) GetID() string     { return fmt.Sprintf("%d", s.UserID) }
